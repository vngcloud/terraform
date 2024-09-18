/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type WrapContentVolumeTypeGatewayResponse struct {
	Code int32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data *VolumeTypeGatewayResponse `json:"data,omitempty"`
}
