/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

type PageObject struct {
	TotalPages    int32 `json:"totalPages,omitempty"`
	TotalElements int32 `json:"totalElements,omitempty"`
	Size          int32 `json:"size,omitempty"`
	Number        int32 `json:"number,omitempty"`
}
