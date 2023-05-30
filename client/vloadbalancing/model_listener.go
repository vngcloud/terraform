/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancing

import (
	"github.com/vngcloud/terraform-provider-vngcloud/client/utils"
)

type Listener struct {
	AllowedCidrs                    string                `json:"allowedCidrs"`
	ClientCertificateAuthentication string                `json:"clientCertificateAuthentication"`
	ConnectionLimit                 int32                 `json:"connectionLimit"`
	CreatedAt                       utils.JsonParsingDate `json:"createdAt"`
	DefaultPoolId                   string                `json:"defaultPoolId"`
	DefaultCertificateAuthority     string                `json:"defaultCertificateAuthority"`
	CertificateAuthorities          []string              `json:"certificateAuthorities"`
	Description                     string                `json:"description"`
	Headers                         []string              `json:"headers"`
	Name                            string                `json:"name"`
	Protocol                        string                `json:"protocol"`
	ProtocolPort                    int32                 `json:"protocolPort"`
	Status                          string                `json:"displayStatus"`
	ProgressStatus                  string                `json:"progressStatus"`
	TimeoutClient                   int32                 `json:"timeoutClient"`
	TimeoutConnection               int32                 `json:"timeoutConnection"`
	TimeoutMember                   int32                 `json:"timeoutMember"`
	UpdatedAt                       utils.JsonParsingDate `json:"updatedAt,omitempty"`
	Uuid                            string                `json:"uuid"`
}
