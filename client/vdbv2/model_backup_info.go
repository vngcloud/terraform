/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BackupInfo struct {
	Id string `json:"id,omitempty"`
	BackendId int64 `json:"backendId,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DbInstanceId string `json:"dbInstanceId,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
	InstanceName string `json:"instanceName,omitempty"`
	Type_ string `json:"type,omitempty"`
	Parent string `json:"parent,omitempty"`
	ParentName string `json:"parentName,omitempty"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	BackupType string `json:"backupType,omitempty"`
	BackupDuration int32 `json:"backupDuration,omitempty"`
	Status string `json:"status,omitempty"`
	DatastoreType string `json:"datastoreType,omitempty"`
	DatastoreVersion string `json:"datastoreVersion,omitempty"`
	StorageType string `json:"storageType,omitempty"`
	StorageSize int32 `json:"storageSize,omitempty"`
	Ram int32 `json:"ram,omitempty"`
	Vcpu int32 `json:"vcpu,omitempty"`
	PackageId string `json:"packageId,omitempty"`
	Username string `json:"username,omitempty"`
	ConfigId string `json:"configId,omitempty"`
	ConfigName string `json:"configName,omitempty"`
	NetIds []string `json:"netIds,omitempty"`
	IsRestoring bool `json:"isRestoring,omitempty"`
	PriceKey string `json:"priceKey,omitempty"`
	Size float32 `json:"size,omitempty"`
	BackupTier string `json:"backupTier,omitempty"`
	EngineGroup int32 `json:"engineGroup,omitempty"`
	SharedBy string `json:"sharedBy,omitempty"`
	NetName string `json:"netName,omitempty"`
	SharedActions []string `json:"sharedActions,omitempty"`
}
