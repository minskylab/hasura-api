package metadata

type MetadataQuery struct {
	Type MetadataRequestType `json:"type"`
	Args interface{}         `json:"args"`
}
