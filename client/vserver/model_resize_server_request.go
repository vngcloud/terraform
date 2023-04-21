/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

// Resize Server Request
type ResizeServerRequest struct {
	// Id of flavor
	FlavorId string `json:"flavorId"`
	// Id of server
	ServerId string `json:"serverId"`
	// Zone of flavor. You can choose if having multiple zone
	ZoneId string `json:"zoneId,omitempty"`
}
