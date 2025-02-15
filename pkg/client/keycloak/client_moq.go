// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package keycloak

import (
	"github.com/Nerzal/gocloak/v8"
	"sync"
)

// Ensure, that KcClientMock does implement KcClient.
// If this is not the case, regenerate this file with moq.
var _ KcClient = &KcClientMock{}

// KcClientMock is a mock implementation of KcClient.
//
// 	func TestSomethingThatUsesKcClient(t *testing.T) {
//
// 		// make and configure a mocked KcClient
// 		mockedKcClient := &KcClientMock{
// 			AddRealmRoleToUserFunc: func(accessToken string, userId string, role gocloak.Role) error {
// 				panic("mock out the AddRealmRoleToUser method")
// 			},
// 			ClientConfigFunc: func(client ClientRepresentation) gocloak.Client {
// 				panic("mock out the ClientConfig method")
// 			},
// 			CreateClientFunc: func(client gocloak.Client, accessToken string) (string, error) {
// 				panic("mock out the CreateClient method")
// 			},
// 			CreateProtocolMapperConfigFunc: func(s string) []gocloak.ProtocolMapperRepresentation {
// 				panic("mock out the CreateProtocolMapperConfig method")
// 			},
// 			CreateRealmRoleFunc: func(accessToken string, roleName string) (*gocloak.Role, error) {
// 				panic("mock out the CreateRealmRole method")
// 			},
// 			DeleteClientFunc: func(internalClientID string, accessToken string) error {
// 				panic("mock out the DeleteClient method")
// 			},
// 			GetCachedTokenFunc: func(tokenKey string) (string, error) {
// 				panic("mock out the GetCachedToken method")
// 			},
// 			GetClientFunc: func(clientId string, accessToken string) (*gocloak.Client, error) {
// 				panic("mock out the GetClient method")
// 			},
// 			GetClientByIdFunc: func(id string, accessToken string) (*gocloak.Client, error) {
// 				panic("mock out the GetClientById method")
// 			},
// 			GetClientSecretFunc: func(internalClientId string, accessToken string) (string, error) {
// 				panic("mock out the GetClientSecret method")
// 			},
// 			GetClientServiceAccountFunc: func(accessToken string, internalClient string) (*gocloak.User, error) {
// 				panic("mock out the GetClientServiceAccount method")
// 			},
// 			GetClientsFunc: func(accessToken string, first int, max int, attribute string) ([]*gocloak.Client, error) {
// 				panic("mock out the GetClients method")
// 			},
// 			GetConfigFunc: func() *KeycloakConfig {
// 				panic("mock out the GetConfig method")
// 			},
// 			GetRealmConfigFunc: func() *KeycloakRealmConfig {
// 				panic("mock out the GetRealmConfig method")
// 			},
// 			GetRealmRoleFunc: func(accessToken string, roleName string) (*gocloak.Role, error) {
// 				panic("mock out the GetRealmRole method")
// 			},
// 			GetTokenFunc: func() (string, error) {
// 				panic("mock out the GetToken method")
// 			},
// 			IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
// 				panic("mock out the IsClientExist method")
// 			},
// 			IsOwnerFunc: func(client *gocloak.Client, userId string) bool {
// 				panic("mock out the IsOwner method")
// 			},
// 			IsSameOrgFunc: func(client *gocloak.Client, orgId string) bool {
// 				panic("mock out the IsSameOrg method")
// 			},
// 			RegenerateClientSecretFunc: func(accessToken string, id string) (*gocloak.CredentialRepresentation, error) {
// 				panic("mock out the RegenerateClientSecret method")
// 			},
// 			UpdateServiceAccountUserFunc: func(accessToken string, serviceAccountUser gocloak.User) error {
// 				panic("mock out the UpdateServiceAccountUser method")
// 			},
// 			UserHasRealmRoleFunc: func(accessToken string, userId string, roleName string) (*gocloak.Role, error) {
// 				panic("mock out the UserHasRealmRole method")
// 			},
// 		}
//
// 		// use mockedKcClient in code that requires KcClient
// 		// and then make assertions.
//
// 	}
type KcClientMock struct {
	// AddRealmRoleToUserFunc mocks the AddRealmRoleToUser method.
	AddRealmRoleToUserFunc func(accessToken string, userId string, role gocloak.Role) error

	// ClientConfigFunc mocks the ClientConfig method.
	ClientConfigFunc func(client ClientRepresentation) gocloak.Client

	// CreateClientFunc mocks the CreateClient method.
	CreateClientFunc func(client gocloak.Client, accessToken string) (string, error)

	// CreateProtocolMapperConfigFunc mocks the CreateProtocolMapperConfig method.
	CreateProtocolMapperConfigFunc func(s string) []gocloak.ProtocolMapperRepresentation

	// CreateRealmRoleFunc mocks the CreateRealmRole method.
	CreateRealmRoleFunc func(accessToken string, roleName string) (*gocloak.Role, error)

	// DeleteClientFunc mocks the DeleteClient method.
	DeleteClientFunc func(internalClientID string, accessToken string) error

	// GetCachedTokenFunc mocks the GetCachedToken method.
	GetCachedTokenFunc func(tokenKey string) (string, error)

	// GetClientFunc mocks the GetClient method.
	GetClientFunc func(clientId string, accessToken string) (*gocloak.Client, error)

	// GetClientByIdFunc mocks the GetClientById method.
	GetClientByIdFunc func(id string, accessToken string) (*gocloak.Client, error)

	// GetClientSecretFunc mocks the GetClientSecret method.
	GetClientSecretFunc func(internalClientId string, accessToken string) (string, error)

	// GetClientServiceAccountFunc mocks the GetClientServiceAccount method.
	GetClientServiceAccountFunc func(accessToken string, internalClient string) (*gocloak.User, error)

	// GetClientsFunc mocks the GetClients method.
	GetClientsFunc func(accessToken string, first int, max int, attribute string) ([]*gocloak.Client, error)

	// GetConfigFunc mocks the GetConfig method.
	GetConfigFunc func() *KeycloakConfig

	// GetRealmConfigFunc mocks the GetRealmConfig method.
	GetRealmConfigFunc func() *KeycloakRealmConfig

	// GetRealmRoleFunc mocks the GetRealmRole method.
	GetRealmRoleFunc func(accessToken string, roleName string) (*gocloak.Role, error)

	// GetTokenFunc mocks the GetToken method.
	GetTokenFunc func() (string, error)

	// IsClientExistFunc mocks the IsClientExist method.
	IsClientExistFunc func(clientId string, accessToken string) (string, error)

	// IsOwnerFunc mocks the IsOwner method.
	IsOwnerFunc func(client *gocloak.Client, userId string) bool

	// IsSameOrgFunc mocks the IsSameOrg method.
	IsSameOrgFunc func(client *gocloak.Client, orgId string) bool

	// RegenerateClientSecretFunc mocks the RegenerateClientSecret method.
	RegenerateClientSecretFunc func(accessToken string, id string) (*gocloak.CredentialRepresentation, error)

	// UpdateServiceAccountUserFunc mocks the UpdateServiceAccountUser method.
	UpdateServiceAccountUserFunc func(accessToken string, serviceAccountUser gocloak.User) error

	// UserHasRealmRoleFunc mocks the UserHasRealmRole method.
	UserHasRealmRoleFunc func(accessToken string, userId string, roleName string) (*gocloak.Role, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddRealmRoleToUser holds details about calls to the AddRealmRoleToUser method.
		AddRealmRoleToUser []struct {
			// AccessToken is the accessToken argument value.
			AccessToken string
			// UserId is the userId argument value.
			UserId string
			// Role is the role argument value.
			Role gocloak.Role
		}
		// ClientConfig holds details about calls to the ClientConfig method.
		ClientConfig []struct {
			// Client is the client argument value.
			Client ClientRepresentation
		}
		// CreateClient holds details about calls to the CreateClient method.
		CreateClient []struct {
			// Client is the client argument value.
			Client gocloak.Client
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// CreateProtocolMapperConfig holds details about calls to the CreateProtocolMapperConfig method.
		CreateProtocolMapperConfig []struct {
			// S is the s argument value.
			S string
		}
		// CreateRealmRole holds details about calls to the CreateRealmRole method.
		CreateRealmRole []struct {
			// AccessToken is the accessToken argument value.
			AccessToken string
			// RoleName is the roleName argument value.
			RoleName string
		}
		// DeleteClient holds details about calls to the DeleteClient method.
		DeleteClient []struct {
			// InternalClientID is the internalClientID argument value.
			InternalClientID string
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// GetCachedToken holds details about calls to the GetCachedToken method.
		GetCachedToken []struct {
			// TokenKey is the tokenKey argument value.
			TokenKey string
		}
		// GetClient holds details about calls to the GetClient method.
		GetClient []struct {
			// ClientId is the clientId argument value.
			ClientId string
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// GetClientById holds details about calls to the GetClientById method.
		GetClientById []struct {
			// ID is the id argument value.
			ID string
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// GetClientSecret holds details about calls to the GetClientSecret method.
		GetClientSecret []struct {
			// InternalClientId is the internalClientId argument value.
			InternalClientId string
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// GetClientServiceAccount holds details about calls to the GetClientServiceAccount method.
		GetClientServiceAccount []struct {
			// AccessToken is the accessToken argument value.
			AccessToken string
			// InternalClient is the internalClient argument value.
			InternalClient string
		}
		// GetClients holds details about calls to the GetClients method.
		GetClients []struct {
			// AccessToken is the accessToken argument value.
			AccessToken string
			// First is the first argument value.
			First int
			// Max is the max argument value.
			Max int
			// Attribute is the attribute argument value.
			Attribute string
		}
		// GetConfig holds details about calls to the GetConfig method.
		GetConfig []struct {
		}
		// GetRealmConfig holds details about calls to the GetRealmConfig method.
		GetRealmConfig []struct {
		}
		// GetRealmRole holds details about calls to the GetRealmRole method.
		GetRealmRole []struct {
			// AccessToken is the accessToken argument value.
			AccessToken string
			// RoleName is the roleName argument value.
			RoleName string
		}
		// GetToken holds details about calls to the GetToken method.
		GetToken []struct {
		}
		// IsClientExist holds details about calls to the IsClientExist method.
		IsClientExist []struct {
			// ClientId is the clientId argument value.
			ClientId string
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// IsOwner holds details about calls to the IsOwner method.
		IsOwner []struct {
			// Client is the client argument value.
			Client *gocloak.Client
			// UserId is the userId argument value.
			UserId string
		}
		// IsSameOrg holds details about calls to the IsSameOrg method.
		IsSameOrg []struct {
			// Client is the client argument value.
			Client *gocloak.Client
			// OrgId is the orgId argument value.
			OrgId string
		}
		// RegenerateClientSecret holds details about calls to the RegenerateClientSecret method.
		RegenerateClientSecret []struct {
			// AccessToken is the accessToken argument value.
			AccessToken string
			// ID is the id argument value.
			ID string
		}
		// UpdateServiceAccountUser holds details about calls to the UpdateServiceAccountUser method.
		UpdateServiceAccountUser []struct {
			// AccessToken is the accessToken argument value.
			AccessToken string
			// ServiceAccountUser is the serviceAccountUser argument value.
			ServiceAccountUser gocloak.User
		}
		// UserHasRealmRole holds details about calls to the UserHasRealmRole method.
		UserHasRealmRole []struct {
			// AccessToken is the accessToken argument value.
			AccessToken string
			// UserId is the userId argument value.
			UserId string
			// RoleName is the roleName argument value.
			RoleName string
		}
	}
	lockAddRealmRoleToUser         sync.RWMutex
	lockClientConfig               sync.RWMutex
	lockCreateClient               sync.RWMutex
	lockCreateProtocolMapperConfig sync.RWMutex
	lockCreateRealmRole            sync.RWMutex
	lockDeleteClient               sync.RWMutex
	lockGetCachedToken             sync.RWMutex
	lockGetClient                  sync.RWMutex
	lockGetClientById              sync.RWMutex
	lockGetClientSecret            sync.RWMutex
	lockGetClientServiceAccount    sync.RWMutex
	lockGetClients                 sync.RWMutex
	lockGetConfig                  sync.RWMutex
	lockGetRealmConfig             sync.RWMutex
	lockGetRealmRole               sync.RWMutex
	lockGetToken                   sync.RWMutex
	lockIsClientExist              sync.RWMutex
	lockIsOwner                    sync.RWMutex
	lockIsSameOrg                  sync.RWMutex
	lockRegenerateClientSecret     sync.RWMutex
	lockUpdateServiceAccountUser   sync.RWMutex
	lockUserHasRealmRole           sync.RWMutex
}

// AddRealmRoleToUser calls AddRealmRoleToUserFunc.
func (mock *KcClientMock) AddRealmRoleToUser(accessToken string, userId string, role gocloak.Role) error {
	if mock.AddRealmRoleToUserFunc == nil {
		panic("KcClientMock.AddRealmRoleToUserFunc: method is nil but KcClient.AddRealmRoleToUser was just called")
	}
	callInfo := struct {
		AccessToken string
		UserId      string
		Role        gocloak.Role
	}{
		AccessToken: accessToken,
		UserId:      userId,
		Role:        role,
	}
	mock.lockAddRealmRoleToUser.Lock()
	mock.calls.AddRealmRoleToUser = append(mock.calls.AddRealmRoleToUser, callInfo)
	mock.lockAddRealmRoleToUser.Unlock()
	return mock.AddRealmRoleToUserFunc(accessToken, userId, role)
}

// AddRealmRoleToUserCalls gets all the calls that were made to AddRealmRoleToUser.
// Check the length with:
//     len(mockedKcClient.AddRealmRoleToUserCalls())
func (mock *KcClientMock) AddRealmRoleToUserCalls() []struct {
	AccessToken string
	UserId      string
	Role        gocloak.Role
} {
	var calls []struct {
		AccessToken string
		UserId      string
		Role        gocloak.Role
	}
	mock.lockAddRealmRoleToUser.RLock()
	calls = mock.calls.AddRealmRoleToUser
	mock.lockAddRealmRoleToUser.RUnlock()
	return calls
}

// ClientConfig calls ClientConfigFunc.
func (mock *KcClientMock) ClientConfig(client ClientRepresentation) gocloak.Client {
	if mock.ClientConfigFunc == nil {
		panic("KcClientMock.ClientConfigFunc: method is nil but KcClient.ClientConfig was just called")
	}
	callInfo := struct {
		Client ClientRepresentation
	}{
		Client: client,
	}
	mock.lockClientConfig.Lock()
	mock.calls.ClientConfig = append(mock.calls.ClientConfig, callInfo)
	mock.lockClientConfig.Unlock()
	return mock.ClientConfigFunc(client)
}

// ClientConfigCalls gets all the calls that were made to ClientConfig.
// Check the length with:
//     len(mockedKcClient.ClientConfigCalls())
func (mock *KcClientMock) ClientConfigCalls() []struct {
	Client ClientRepresentation
} {
	var calls []struct {
		Client ClientRepresentation
	}
	mock.lockClientConfig.RLock()
	calls = mock.calls.ClientConfig
	mock.lockClientConfig.RUnlock()
	return calls
}

// CreateClient calls CreateClientFunc.
func (mock *KcClientMock) CreateClient(client gocloak.Client, accessToken string) (string, error) {
	if mock.CreateClientFunc == nil {
		panic("KcClientMock.CreateClientFunc: method is nil but KcClient.CreateClient was just called")
	}
	callInfo := struct {
		Client      gocloak.Client
		AccessToken string
	}{
		Client:      client,
		AccessToken: accessToken,
	}
	mock.lockCreateClient.Lock()
	mock.calls.CreateClient = append(mock.calls.CreateClient, callInfo)
	mock.lockCreateClient.Unlock()
	return mock.CreateClientFunc(client, accessToken)
}

// CreateClientCalls gets all the calls that were made to CreateClient.
// Check the length with:
//     len(mockedKcClient.CreateClientCalls())
func (mock *KcClientMock) CreateClientCalls() []struct {
	Client      gocloak.Client
	AccessToken string
} {
	var calls []struct {
		Client      gocloak.Client
		AccessToken string
	}
	mock.lockCreateClient.RLock()
	calls = mock.calls.CreateClient
	mock.lockCreateClient.RUnlock()
	return calls
}

// CreateProtocolMapperConfig calls CreateProtocolMapperConfigFunc.
func (mock *KcClientMock) CreateProtocolMapperConfig(s string) []gocloak.ProtocolMapperRepresentation {
	if mock.CreateProtocolMapperConfigFunc == nil {
		panic("KcClientMock.CreateProtocolMapperConfigFunc: method is nil but KcClient.CreateProtocolMapperConfig was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockCreateProtocolMapperConfig.Lock()
	mock.calls.CreateProtocolMapperConfig = append(mock.calls.CreateProtocolMapperConfig, callInfo)
	mock.lockCreateProtocolMapperConfig.Unlock()
	return mock.CreateProtocolMapperConfigFunc(s)
}

// CreateProtocolMapperConfigCalls gets all the calls that were made to CreateProtocolMapperConfig.
// Check the length with:
//     len(mockedKcClient.CreateProtocolMapperConfigCalls())
func (mock *KcClientMock) CreateProtocolMapperConfigCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockCreateProtocolMapperConfig.RLock()
	calls = mock.calls.CreateProtocolMapperConfig
	mock.lockCreateProtocolMapperConfig.RUnlock()
	return calls
}

// CreateRealmRole calls CreateRealmRoleFunc.
func (mock *KcClientMock) CreateRealmRole(accessToken string, roleName string) (*gocloak.Role, error) {
	if mock.CreateRealmRoleFunc == nil {
		panic("KcClientMock.CreateRealmRoleFunc: method is nil but KcClient.CreateRealmRole was just called")
	}
	callInfo := struct {
		AccessToken string
		RoleName    string
	}{
		AccessToken: accessToken,
		RoleName:    roleName,
	}
	mock.lockCreateRealmRole.Lock()
	mock.calls.CreateRealmRole = append(mock.calls.CreateRealmRole, callInfo)
	mock.lockCreateRealmRole.Unlock()
	return mock.CreateRealmRoleFunc(accessToken, roleName)
}

// CreateRealmRoleCalls gets all the calls that were made to CreateRealmRole.
// Check the length with:
//     len(mockedKcClient.CreateRealmRoleCalls())
func (mock *KcClientMock) CreateRealmRoleCalls() []struct {
	AccessToken string
	RoleName    string
} {
	var calls []struct {
		AccessToken string
		RoleName    string
	}
	mock.lockCreateRealmRole.RLock()
	calls = mock.calls.CreateRealmRole
	mock.lockCreateRealmRole.RUnlock()
	return calls
}

// DeleteClient calls DeleteClientFunc.
func (mock *KcClientMock) DeleteClient(internalClientID string, accessToken string) error {
	if mock.DeleteClientFunc == nil {
		panic("KcClientMock.DeleteClientFunc: method is nil but KcClient.DeleteClient was just called")
	}
	callInfo := struct {
		InternalClientID string
		AccessToken      string
	}{
		InternalClientID: internalClientID,
		AccessToken:      accessToken,
	}
	mock.lockDeleteClient.Lock()
	mock.calls.DeleteClient = append(mock.calls.DeleteClient, callInfo)
	mock.lockDeleteClient.Unlock()
	return mock.DeleteClientFunc(internalClientID, accessToken)
}

// DeleteClientCalls gets all the calls that were made to DeleteClient.
// Check the length with:
//     len(mockedKcClient.DeleteClientCalls())
func (mock *KcClientMock) DeleteClientCalls() []struct {
	InternalClientID string
	AccessToken      string
} {
	var calls []struct {
		InternalClientID string
		AccessToken      string
	}
	mock.lockDeleteClient.RLock()
	calls = mock.calls.DeleteClient
	mock.lockDeleteClient.RUnlock()
	return calls
}

// GetCachedToken calls GetCachedTokenFunc.
func (mock *KcClientMock) GetCachedToken(tokenKey string) (string, error) {
	if mock.GetCachedTokenFunc == nil {
		panic("KcClientMock.GetCachedTokenFunc: method is nil but KcClient.GetCachedToken was just called")
	}
	callInfo := struct {
		TokenKey string
	}{
		TokenKey: tokenKey,
	}
	mock.lockGetCachedToken.Lock()
	mock.calls.GetCachedToken = append(mock.calls.GetCachedToken, callInfo)
	mock.lockGetCachedToken.Unlock()
	return mock.GetCachedTokenFunc(tokenKey)
}

// GetCachedTokenCalls gets all the calls that were made to GetCachedToken.
// Check the length with:
//     len(mockedKcClient.GetCachedTokenCalls())
func (mock *KcClientMock) GetCachedTokenCalls() []struct {
	TokenKey string
} {
	var calls []struct {
		TokenKey string
	}
	mock.lockGetCachedToken.RLock()
	calls = mock.calls.GetCachedToken
	mock.lockGetCachedToken.RUnlock()
	return calls
}

// GetClient calls GetClientFunc.
func (mock *KcClientMock) GetClient(clientId string, accessToken string) (*gocloak.Client, error) {
	if mock.GetClientFunc == nil {
		panic("KcClientMock.GetClientFunc: method is nil but KcClient.GetClient was just called")
	}
	callInfo := struct {
		ClientId    string
		AccessToken string
	}{
		ClientId:    clientId,
		AccessToken: accessToken,
	}
	mock.lockGetClient.Lock()
	mock.calls.GetClient = append(mock.calls.GetClient, callInfo)
	mock.lockGetClient.Unlock()
	return mock.GetClientFunc(clientId, accessToken)
}

// GetClientCalls gets all the calls that were made to GetClient.
// Check the length with:
//     len(mockedKcClient.GetClientCalls())
func (mock *KcClientMock) GetClientCalls() []struct {
	ClientId    string
	AccessToken string
} {
	var calls []struct {
		ClientId    string
		AccessToken string
	}
	mock.lockGetClient.RLock()
	calls = mock.calls.GetClient
	mock.lockGetClient.RUnlock()
	return calls
}

// GetClientById calls GetClientByIdFunc.
func (mock *KcClientMock) GetClientById(id string, accessToken string) (*gocloak.Client, error) {
	if mock.GetClientByIdFunc == nil {
		panic("KcClientMock.GetClientByIdFunc: method is nil but KcClient.GetClientById was just called")
	}
	callInfo := struct {
		ID          string
		AccessToken string
	}{
		ID:          id,
		AccessToken: accessToken,
	}
	mock.lockGetClientById.Lock()
	mock.calls.GetClientById = append(mock.calls.GetClientById, callInfo)
	mock.lockGetClientById.Unlock()
	return mock.GetClientByIdFunc(id, accessToken)
}

// GetClientByIdCalls gets all the calls that were made to GetClientById.
// Check the length with:
//     len(mockedKcClient.GetClientByIdCalls())
func (mock *KcClientMock) GetClientByIdCalls() []struct {
	ID          string
	AccessToken string
} {
	var calls []struct {
		ID          string
		AccessToken string
	}
	mock.lockGetClientById.RLock()
	calls = mock.calls.GetClientById
	mock.lockGetClientById.RUnlock()
	return calls
}

// GetClientSecret calls GetClientSecretFunc.
func (mock *KcClientMock) GetClientSecret(internalClientId string, accessToken string) (string, error) {
	if mock.GetClientSecretFunc == nil {
		panic("KcClientMock.GetClientSecretFunc: method is nil but KcClient.GetClientSecret was just called")
	}
	callInfo := struct {
		InternalClientId string
		AccessToken      string
	}{
		InternalClientId: internalClientId,
		AccessToken:      accessToken,
	}
	mock.lockGetClientSecret.Lock()
	mock.calls.GetClientSecret = append(mock.calls.GetClientSecret, callInfo)
	mock.lockGetClientSecret.Unlock()
	return mock.GetClientSecretFunc(internalClientId, accessToken)
}

// GetClientSecretCalls gets all the calls that were made to GetClientSecret.
// Check the length with:
//     len(mockedKcClient.GetClientSecretCalls())
func (mock *KcClientMock) GetClientSecretCalls() []struct {
	InternalClientId string
	AccessToken      string
} {
	var calls []struct {
		InternalClientId string
		AccessToken      string
	}
	mock.lockGetClientSecret.RLock()
	calls = mock.calls.GetClientSecret
	mock.lockGetClientSecret.RUnlock()
	return calls
}

// GetClientServiceAccount calls GetClientServiceAccountFunc.
func (mock *KcClientMock) GetClientServiceAccount(accessToken string, internalClient string) (*gocloak.User, error) {
	if mock.GetClientServiceAccountFunc == nil {
		panic("KcClientMock.GetClientServiceAccountFunc: method is nil but KcClient.GetClientServiceAccount was just called")
	}
	callInfo := struct {
		AccessToken    string
		InternalClient string
	}{
		AccessToken:    accessToken,
		InternalClient: internalClient,
	}
	mock.lockGetClientServiceAccount.Lock()
	mock.calls.GetClientServiceAccount = append(mock.calls.GetClientServiceAccount, callInfo)
	mock.lockGetClientServiceAccount.Unlock()
	return mock.GetClientServiceAccountFunc(accessToken, internalClient)
}

// GetClientServiceAccountCalls gets all the calls that were made to GetClientServiceAccount.
// Check the length with:
//     len(mockedKcClient.GetClientServiceAccountCalls())
func (mock *KcClientMock) GetClientServiceAccountCalls() []struct {
	AccessToken    string
	InternalClient string
} {
	var calls []struct {
		AccessToken    string
		InternalClient string
	}
	mock.lockGetClientServiceAccount.RLock()
	calls = mock.calls.GetClientServiceAccount
	mock.lockGetClientServiceAccount.RUnlock()
	return calls
}

// GetClients calls GetClientsFunc.
func (mock *KcClientMock) GetClients(accessToken string, first int, max int, attribute string) ([]*gocloak.Client, error) {
	if mock.GetClientsFunc == nil {
		panic("KcClientMock.GetClientsFunc: method is nil but KcClient.GetClients was just called")
	}
	callInfo := struct {
		AccessToken string
		First       int
		Max         int
		Attribute   string
	}{
		AccessToken: accessToken,
		First:       first,
		Max:         max,
		Attribute:   attribute,
	}
	mock.lockGetClients.Lock()
	mock.calls.GetClients = append(mock.calls.GetClients, callInfo)
	mock.lockGetClients.Unlock()
	return mock.GetClientsFunc(accessToken, first, max, attribute)
}

// GetClientsCalls gets all the calls that were made to GetClients.
// Check the length with:
//     len(mockedKcClient.GetClientsCalls())
func (mock *KcClientMock) GetClientsCalls() []struct {
	AccessToken string
	First       int
	Max         int
	Attribute   string
} {
	var calls []struct {
		AccessToken string
		First       int
		Max         int
		Attribute   string
	}
	mock.lockGetClients.RLock()
	calls = mock.calls.GetClients
	mock.lockGetClients.RUnlock()
	return calls
}

// GetConfig calls GetConfigFunc.
func (mock *KcClientMock) GetConfig() *KeycloakConfig {
	if mock.GetConfigFunc == nil {
		panic("KcClientMock.GetConfigFunc: method is nil but KcClient.GetConfig was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetConfig.Lock()
	mock.calls.GetConfig = append(mock.calls.GetConfig, callInfo)
	mock.lockGetConfig.Unlock()
	return mock.GetConfigFunc()
}

// GetConfigCalls gets all the calls that were made to GetConfig.
// Check the length with:
//     len(mockedKcClient.GetConfigCalls())
func (mock *KcClientMock) GetConfigCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetConfig.RLock()
	calls = mock.calls.GetConfig
	mock.lockGetConfig.RUnlock()
	return calls
}

// GetRealmConfig calls GetRealmConfigFunc.
func (mock *KcClientMock) GetRealmConfig() *KeycloakRealmConfig {
	if mock.GetRealmConfigFunc == nil {
		panic("KcClientMock.GetRealmConfigFunc: method is nil but KcClient.GetRealmConfig was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetRealmConfig.Lock()
	mock.calls.GetRealmConfig = append(mock.calls.GetRealmConfig, callInfo)
	mock.lockGetRealmConfig.Unlock()
	return mock.GetRealmConfigFunc()
}

// GetRealmConfigCalls gets all the calls that were made to GetRealmConfig.
// Check the length with:
//     len(mockedKcClient.GetRealmConfigCalls())
func (mock *KcClientMock) GetRealmConfigCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetRealmConfig.RLock()
	calls = mock.calls.GetRealmConfig
	mock.lockGetRealmConfig.RUnlock()
	return calls
}

// GetRealmRole calls GetRealmRoleFunc.
func (mock *KcClientMock) GetRealmRole(accessToken string, roleName string) (*gocloak.Role, error) {
	if mock.GetRealmRoleFunc == nil {
		panic("KcClientMock.GetRealmRoleFunc: method is nil but KcClient.GetRealmRole was just called")
	}
	callInfo := struct {
		AccessToken string
		RoleName    string
	}{
		AccessToken: accessToken,
		RoleName:    roleName,
	}
	mock.lockGetRealmRole.Lock()
	mock.calls.GetRealmRole = append(mock.calls.GetRealmRole, callInfo)
	mock.lockGetRealmRole.Unlock()
	return mock.GetRealmRoleFunc(accessToken, roleName)
}

// GetRealmRoleCalls gets all the calls that were made to GetRealmRole.
// Check the length with:
//     len(mockedKcClient.GetRealmRoleCalls())
func (mock *KcClientMock) GetRealmRoleCalls() []struct {
	AccessToken string
	RoleName    string
} {
	var calls []struct {
		AccessToken string
		RoleName    string
	}
	mock.lockGetRealmRole.RLock()
	calls = mock.calls.GetRealmRole
	mock.lockGetRealmRole.RUnlock()
	return calls
}

// GetToken calls GetTokenFunc.
func (mock *KcClientMock) GetToken() (string, error) {
	if mock.GetTokenFunc == nil {
		panic("KcClientMock.GetTokenFunc: method is nil but KcClient.GetToken was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetToken.Lock()
	mock.calls.GetToken = append(mock.calls.GetToken, callInfo)
	mock.lockGetToken.Unlock()
	return mock.GetTokenFunc()
}

// GetTokenCalls gets all the calls that were made to GetToken.
// Check the length with:
//     len(mockedKcClient.GetTokenCalls())
func (mock *KcClientMock) GetTokenCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetToken.RLock()
	calls = mock.calls.GetToken
	mock.lockGetToken.RUnlock()
	return calls
}

// IsClientExist calls IsClientExistFunc.
func (mock *KcClientMock) IsClientExist(clientId string, accessToken string) (string, error) {
	if mock.IsClientExistFunc == nil {
		panic("KcClientMock.IsClientExistFunc: method is nil but KcClient.IsClientExist was just called")
	}
	callInfo := struct {
		ClientId    string
		AccessToken string
	}{
		ClientId:    clientId,
		AccessToken: accessToken,
	}
	mock.lockIsClientExist.Lock()
	mock.calls.IsClientExist = append(mock.calls.IsClientExist, callInfo)
	mock.lockIsClientExist.Unlock()
	return mock.IsClientExistFunc(clientId, accessToken)
}

// IsClientExistCalls gets all the calls that were made to IsClientExist.
// Check the length with:
//     len(mockedKcClient.IsClientExistCalls())
func (mock *KcClientMock) IsClientExistCalls() []struct {
	ClientId    string
	AccessToken string
} {
	var calls []struct {
		ClientId    string
		AccessToken string
	}
	mock.lockIsClientExist.RLock()
	calls = mock.calls.IsClientExist
	mock.lockIsClientExist.RUnlock()
	return calls
}

// IsOwner calls IsOwnerFunc.
func (mock *KcClientMock) IsOwner(client *gocloak.Client, userId string) bool {
	if mock.IsOwnerFunc == nil {
		panic("KcClientMock.IsOwnerFunc: method is nil but KcClient.IsOwner was just called")
	}
	callInfo := struct {
		Client *gocloak.Client
		UserId string
	}{
		Client: client,
		UserId: userId,
	}
	mock.lockIsOwner.Lock()
	mock.calls.IsOwner = append(mock.calls.IsOwner, callInfo)
	mock.lockIsOwner.Unlock()
	return mock.IsOwnerFunc(client, userId)
}

// IsOwnerCalls gets all the calls that were made to IsOwner.
// Check the length with:
//     len(mockedKcClient.IsOwnerCalls())
func (mock *KcClientMock) IsOwnerCalls() []struct {
	Client *gocloak.Client
	UserId string
} {
	var calls []struct {
		Client *gocloak.Client
		UserId string
	}
	mock.lockIsOwner.RLock()
	calls = mock.calls.IsOwner
	mock.lockIsOwner.RUnlock()
	return calls
}

// IsSameOrg calls IsSameOrgFunc.
func (mock *KcClientMock) IsSameOrg(client *gocloak.Client, orgId string) bool {
	if mock.IsSameOrgFunc == nil {
		panic("KcClientMock.IsSameOrgFunc: method is nil but KcClient.IsSameOrg was just called")
	}
	callInfo := struct {
		Client *gocloak.Client
		OrgId  string
	}{
		Client: client,
		OrgId:  orgId,
	}
	mock.lockIsSameOrg.Lock()
	mock.calls.IsSameOrg = append(mock.calls.IsSameOrg, callInfo)
	mock.lockIsSameOrg.Unlock()
	return mock.IsSameOrgFunc(client, orgId)
}

// IsSameOrgCalls gets all the calls that were made to IsSameOrg.
// Check the length with:
//     len(mockedKcClient.IsSameOrgCalls())
func (mock *KcClientMock) IsSameOrgCalls() []struct {
	Client *gocloak.Client
	OrgId  string
} {
	var calls []struct {
		Client *gocloak.Client
		OrgId  string
	}
	mock.lockIsSameOrg.RLock()
	calls = mock.calls.IsSameOrg
	mock.lockIsSameOrg.RUnlock()
	return calls
}

// RegenerateClientSecret calls RegenerateClientSecretFunc.
func (mock *KcClientMock) RegenerateClientSecret(accessToken string, id string) (*gocloak.CredentialRepresentation, error) {
	if mock.RegenerateClientSecretFunc == nil {
		panic("KcClientMock.RegenerateClientSecretFunc: method is nil but KcClient.RegenerateClientSecret was just called")
	}
	callInfo := struct {
		AccessToken string
		ID          string
	}{
		AccessToken: accessToken,
		ID:          id,
	}
	mock.lockRegenerateClientSecret.Lock()
	mock.calls.RegenerateClientSecret = append(mock.calls.RegenerateClientSecret, callInfo)
	mock.lockRegenerateClientSecret.Unlock()
	return mock.RegenerateClientSecretFunc(accessToken, id)
}

// RegenerateClientSecretCalls gets all the calls that were made to RegenerateClientSecret.
// Check the length with:
//     len(mockedKcClient.RegenerateClientSecretCalls())
func (mock *KcClientMock) RegenerateClientSecretCalls() []struct {
	AccessToken string
	ID          string
} {
	var calls []struct {
		AccessToken string
		ID          string
	}
	mock.lockRegenerateClientSecret.RLock()
	calls = mock.calls.RegenerateClientSecret
	mock.lockRegenerateClientSecret.RUnlock()
	return calls
}

// UpdateServiceAccountUser calls UpdateServiceAccountUserFunc.
func (mock *KcClientMock) UpdateServiceAccountUser(accessToken string, serviceAccountUser gocloak.User) error {
	if mock.UpdateServiceAccountUserFunc == nil {
		panic("KcClientMock.UpdateServiceAccountUserFunc: method is nil but KcClient.UpdateServiceAccountUser was just called")
	}
	callInfo := struct {
		AccessToken        string
		ServiceAccountUser gocloak.User
	}{
		AccessToken:        accessToken,
		ServiceAccountUser: serviceAccountUser,
	}
	mock.lockUpdateServiceAccountUser.Lock()
	mock.calls.UpdateServiceAccountUser = append(mock.calls.UpdateServiceAccountUser, callInfo)
	mock.lockUpdateServiceAccountUser.Unlock()
	return mock.UpdateServiceAccountUserFunc(accessToken, serviceAccountUser)
}

// UpdateServiceAccountUserCalls gets all the calls that were made to UpdateServiceAccountUser.
// Check the length with:
//     len(mockedKcClient.UpdateServiceAccountUserCalls())
func (mock *KcClientMock) UpdateServiceAccountUserCalls() []struct {
	AccessToken        string
	ServiceAccountUser gocloak.User
} {
	var calls []struct {
		AccessToken        string
		ServiceAccountUser gocloak.User
	}
	mock.lockUpdateServiceAccountUser.RLock()
	calls = mock.calls.UpdateServiceAccountUser
	mock.lockUpdateServiceAccountUser.RUnlock()
	return calls
}

// UserHasRealmRole calls UserHasRealmRoleFunc.
func (mock *KcClientMock) UserHasRealmRole(accessToken string, userId string, roleName string) (*gocloak.Role, error) {
	if mock.UserHasRealmRoleFunc == nil {
		panic("KcClientMock.UserHasRealmRoleFunc: method is nil but KcClient.UserHasRealmRole was just called")
	}
	callInfo := struct {
		AccessToken string
		UserId      string
		RoleName    string
	}{
		AccessToken: accessToken,
		UserId:      userId,
		RoleName:    roleName,
	}
	mock.lockUserHasRealmRole.Lock()
	mock.calls.UserHasRealmRole = append(mock.calls.UserHasRealmRole, callInfo)
	mock.lockUserHasRealmRole.Unlock()
	return mock.UserHasRealmRoleFunc(accessToken, userId, roleName)
}

// UserHasRealmRoleCalls gets all the calls that were made to UserHasRealmRole.
// Check the length with:
//     len(mockedKcClient.UserHasRealmRoleCalls())
func (mock *KcClientMock) UserHasRealmRoleCalls() []struct {
	AccessToken string
	UserId      string
	RoleName    string
} {
	var calls []struct {
		AccessToken string
		UserId      string
		RoleName    string
	}
	mock.lockUserHasRealmRole.RLock()
	calls = mock.calls.UserHasRealmRole
	mock.lockUserHasRealmRole.RUnlock()
	return calls
}
