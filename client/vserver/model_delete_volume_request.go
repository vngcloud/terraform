/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

// Delete Volume Request
type DeleteVolumeRequest struct {
	// Id of project
	ProjectId string `json:"projectId,omitempty"`
	// Id of user
	UserId int32 `json:"userId,omitempty"`
	// Id of volume
	VolumeId string `json:"volumeId"`
}
