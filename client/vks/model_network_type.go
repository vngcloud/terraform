/*
 * vks-api API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0-SNAPSHOT
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vks

type NetworkType string

// List of NetworkType
const (
	CALICO_NetworkType NetworkType = "CALICO"
	CILIUM_NetworkType NetworkType = "CILIUM"
)
