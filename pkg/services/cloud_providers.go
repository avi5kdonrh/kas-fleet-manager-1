package services

import (
	"strings"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/clusters"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/clusters/types"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/db"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"github.com/patrickmn/go-cache"
)

const keyCloudProvidersWithRegions = "cloudProviderWithRegions"

var cloudPoviderIdToDisplayNameMapping map[string]string = map[string]string{
	"aws":   "Amazon Web Services",
	"azure": "Microsoft Azure",
	"gcp":   "Google Cloud Platform",
}

//go:generate moq -out cloud_providers_moq.go . CloudProvidersService
type CloudProvidersService interface {
	GetCloudProvidersWithRegions() ([]CloudProviderWithRegions, *errors.ServiceError)
	GetCachedCloudProvidersWithRegions() ([]CloudProviderWithRegions, *errors.ServiceError)
	ListCloudProviders() ([]api.CloudProvider, *errors.ServiceError)
	ListCloudProviderRegions(id string) ([]api.CloudRegion, *errors.ServiceError)
}

func NewCloudProvidersService(providerFactory clusters.ProviderFactory, connectionFactory *db.ConnectionFactory) CloudProvidersService {
	return &cloudProvidersService{
		providerFactory:   providerFactory,
		connectionFactory: connectionFactory,
		cache:             cache.New(5*time.Minute, 10*time.Minute),
	}
}

type cloudProvidersService struct {
	providerFactory   clusters.ProviderFactory
	connectionFactory *db.ConnectionFactory
	cache             *cache.Cache
}

type CloudProviderWithRegions struct {
	ID         string
	RegionList *types.CloudProviderRegionInfoList
}

func (p cloudProvidersService) GetCloudProvidersWithRegions() ([]CloudProviderWithRegions, *errors.ServiceError) {
	var cloudProviderWithRegions []CloudProviderWithRegions
	// TODO: when there are multiple provider types, we can list them from the db first and then call out to each of the implementation to get their provider and region info
	provider, err := p.providerFactory.GetProvider(api.ClusterProviderOCM)
	if err != nil {
		return nil, errors.NewWithCause(errors.ErrorGeneral, err, "failed to find implementation")
	}
	providerList, err := provider.GetCloudProviders()
	if err != nil {
		return nil, errors.NewWithCause(errors.ErrorGeneral, err, "failed to retrieve cloud provider list")
	}
	for _, cp := range providerList.Items {
		regions, regionErr := provider.GetCloudProviderRegions(cp)
		if regionErr != nil {
			return nil, errors.NewWithCause(errors.ErrorGeneral, err, "failed to retrieve cloud regions")
		}
		cloudProviderWithRegions = append(cloudProviderWithRegions, CloudProviderWithRegions{
			ID:         cp.ID,
			RegionList: regions,
		})
	}

	return cloudProviderWithRegions, nil
}

func (p cloudProvidersService) GetCachedCloudProvidersWithRegions() ([]CloudProviderWithRegions, *errors.ServiceError) {
	cachedCloudProviderWithRegions, cached := p.cache.Get(keyCloudProvidersWithRegions)
	if cached {
		return convertToCloudProviderWithRegionsType(cachedCloudProviderWithRegions)
	}
	cloudProviderWithRegions, err := p.GetCloudProvidersWithRegions()
	if err != nil {
		return nil, err
	}
	p.cache.Set(keyCloudProvidersWithRegions, cloudProviderWithRegions, cache.DefaultExpiration)
	return cloudProviderWithRegions, nil
}

func convertToCloudProviderWithRegionsType(cachedCloudProviderWithRegions interface{}) ([]CloudProviderWithRegions, *errors.ServiceError) {
	cloudProviderWithRegions, ok := cachedCloudProviderWithRegions.([]CloudProviderWithRegions)
	if ok {
		return cloudProviderWithRegions, nil
	}
	return nil, nil
}

func (p cloudProvidersService) ListCloudProviders() ([]api.CloudProvider, *errors.ServiceError) {
	type Cluster struct {
		ProviderType api.ClusterProviderType `json:"provider_type"`
	}

	dbConn := p.connectionFactory.New().
		Model(&Cluster{}).
		Distinct("provider_type").
		Where("status NOT IN (?)", api.ClusterDeletionStatuses)

	var results []Cluster
	err := dbConn.Find(&results).Error
	if err != nil {
		return nil, errors.NewWithCause(errors.ErrorGeneral, err, "Failed to list clusters providers")
	}

	alreadyVisitedCloudProviders := map[string]bool{}
	cloudProviderList := []api.CloudProvider{}
	for _, result := range results {
		provider, err := p.providerFactory.GetProvider(result.ProviderType)
		if err != nil {
			return nil, errors.NewWithCause(errors.ErrorGeneral, err, "failed to find implementation")
		}
		providerList, err := provider.GetCloudProviders()
		if err != nil {
			return nil, errors.NewWithCause(errors.ErrorGeneral, err, "failed to retrieve cloud provider list")
		}
		for _, cp := range providerList.Items {
			_, cloudProviderAlreadyCollected := alreadyVisitedCloudProviders[cp.ID]
			if cloudProviderAlreadyCollected {
				continue
			}
			cloudProviderList = append(cloudProviderList, api.CloudProvider{
				Id:          cp.ID,
				Name:        cp.Name,
				DisplayName: setDisplayName(cp.ID, cp.DisplayName),
			})
			alreadyVisitedCloudProviders[cp.ID] = true
		}
	}

	return cloudProviderList, nil
}

func (p cloudProvidersService) ListCloudProviderRegions(id string) ([]api.CloudRegion, *errors.ServiceError) {
	cloudRegionList := []api.CloudRegion{}
	cloudProviders, err := p.GetCloudProvidersWithRegions()
	if err != nil {
		return nil, errors.NewWithCause(errors.ErrorGeneral, err, "failed to retrieve cloud provider regions")
	}

	for _, cloudProvider := range cloudProviders {
		if cloudProvider.ID == id {
			for _, r := range cloudProvider.RegionList.Items {
				cloudRegionList = append(cloudRegionList, api.CloudRegion{
					Id:            r.ID,
					CloudProvider: r.CloudProviderID,
					DisplayName:   r.DisplayName,
				})
			}
			break
		}
	}

	return cloudRegionList, nil
}

func setDisplayName(providerId string, defaultDisplayName string) string {
	displayName, ok := cloudPoviderIdToDisplayNameMapping[strings.ToLower(providerId)]
	if ok {
		return displayName
	}

	return defaultDisplayName
}
