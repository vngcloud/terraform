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

type SecgroupRule struct {
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	Description      string    `json:"description,omitempty"`
	Direction        string    `json:"direction,omitempty"`
	DisplayCreatedAt string    `json:"displayCreatedAt,omitempty"`
	EtherType        string    `json:"etherType,omitempty"`
	Id               string    `json:"id,omitempty"`
	PortRangeMax     int32     `json:"portRangeMax,omitempty"`
	PortRangeMin     int32     `json:"portRangeMin,omitempty"`
	Protocol         string    `json:"protocol,omitempty"`
	RemoteGroupId    string    `json:"remoteGroupId,omitempty"`
	RemoteGroupName  string    `json:"remoteGroupName,omitempty"`
	RemoteIpPrefix   string    `json:"remoteIpPrefix,omitempty"`
	Status           string    `json:"status,omitempty"`
}
