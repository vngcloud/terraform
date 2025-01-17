/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

type CreateBackupResponse struct {
	ErrorMsg     string `json:"errorMsg,omitempty"`
	Code         int32  `json:"code,omitempty"`
	Success      bool   `json:"success,omitempty"`
	Total        int32  `json:"total,omitempty"`
	BackupId     string `json:"backupId,omitempty"`
	ProjectId    string `json:"projectId,omitempty"`
	DbInstanceId string `json:"dbInstanceId,omitempty"`
}
