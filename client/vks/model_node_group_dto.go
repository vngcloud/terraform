/*
 * vks-api API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0-SNAPSHOT
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vks

type NodeGroupDto struct {
	Id        string `json:"id,omitempty"`
	ClusterId string `json:"clusterId,omitempty"`
	Name      string `json:"name,omitempty"`
	Status    string `json:"status,omitempty"`
	NumNodes  int64  `json:"numNodes,omitempty"`
	ImageId   string `json:"imageId,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
