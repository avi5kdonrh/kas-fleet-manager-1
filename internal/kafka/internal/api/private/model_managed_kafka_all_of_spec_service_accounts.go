/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager APIs that are used by internal services e.g kas-fleetshard operators.
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ManagedKafkaAllOfSpecServiceAccounts struct for ManagedKafkaAllOfSpecServiceAccounts
type ManagedKafkaAllOfSpecServiceAccounts struct {
	Name      string `json:"name"`
	Principal string `json:"principal"`
	Password  string `json:"password"`
}
