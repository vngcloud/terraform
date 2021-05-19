/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

// Detach Volume Request
type DetachVolumeRequest struct {
	// Id of server
	InstanceId string `json:"instanceId"`
	// Id of user
	UserId int32 `json:"userId,omitempty"`
	// Id of volume
	VolumeId string `json:"volumeId"`
}