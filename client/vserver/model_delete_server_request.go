/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

// Delete Server Request
type DeleteServerRequest struct {
	// Skip step move into the trash
	ForceDelete bool `json:"forceDelete,omitempty"`
	// Id of server
	ServerId string `json:"serverId"`
	// Id of user
	UserId int32 `json:"userId,omitempty"`
}