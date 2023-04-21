/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

import (
	"time"
)

type Server struct {
	BootVolumeId       string                      `json:"bootVolumeId,omitempty"`
	CreatedAt          time.Time                   `json:"createdAt,omitempty"`
	EncryptionVolume   bool                        `json:"encryptionVolume,omitempty"`
	ExternalInterfaces []InterfaceNetworkInterface `json:"externalInterfaces,omitempty"`
	Flavor             *Flavor                     `json:"flavor,omitempty"`
	Image              *OsImage                    `json:"image,omitempty"`
	InternalInterfaces []InterfaceNetworkInterface `json:"internalInterfaces,omitempty"`
	Licence            bool                        `json:"licence,omitempty"`
	Name               string                      `json:"name,omitempty"`
	SecGroups          []ServerSecGroup            `json:"secGroups,omitempty"`
	ServerGroupId      string                      `json:"serverGroupId,omitempty"`
	ServerGroupName    string                      `json:"serverGroupName,omitempty"`
	SshKeyName         string                      `json:"sshKeyName,omitempty"`
	Status             string                      `json:"status,omitempty"`
	Uuid               string                      `json:"uuid,omitempty"`
	ZoneId             string                      `json:"zoneId,omitempty"`
}
