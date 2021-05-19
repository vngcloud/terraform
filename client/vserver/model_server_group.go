/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

import (
	"time"
)

type ServerGroup struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	Policy string `json:"policy,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}