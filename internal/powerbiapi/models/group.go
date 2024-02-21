package models

type Group struct {
	CapacityId                  string                      `json:"capacityId"`
	DataflowStorageId           string                      `json:"dataflowStorageId"`
	DefaultDatasetStorageFormat DefaultDatasetStorageFormat `json:"defaultDatasetStorageFormat"`
	Id                          string                      `json:"id"`
	IsOnDedicatedCapacity       bool                        `json:"isOnDedicatedCapacity"`
	IsReadOnly                  bool                        `json:"isReadOnly"`
	LogAnalyticsWorkspace       string                      `json:"logAnalyticsWorkspace"`
	Name                        string                      `json:"name"`
}
