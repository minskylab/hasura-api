package metadata

type PgAddSourceArgs struct{}

type PgDropSourceArgs struct{}

type PgTrackTableArgs struct {
	Table         ITableName          `json:"table"`
	Configuration *TableConfiguration `json:"configuration,omitempty"`
	Source        *SourceName         `json:"source,omitempty"`
}

type PgUntrackTableArgs struct {
	Table   ITableName  `json:"table"`
	Cascade bool        `json:"cascade,omitempty"`
	Source  *SourceName `json:"source,omitempty"`
}

type PgSetTableCustomizationArgs struct {
	Table         ITableName          `json:"table"`
	Configuration *TableConfiguration `json:"configuration"`
	Source        *SourceName         `json:"source,omitempty"`
}

type PgSetTableIsEnumArgs struct{}

type PgTrackFunctionArgs struct{}

type PgUntrackFunctionArgs struct{}

type PgSetFunctionCustomizationArgs struct{}

type PgCreateFunctionPermissionArgs struct{}

type PgDropFunctionPermissionArgs struct{}

type PgCreateObjectRelationshipArgs struct {
	Table   ITableName       `json:"table"`
	Name    RelationshipName `json:"name"`
	Using   ObjRelUsing      `json:"using"`
	Comment string           `json:"comment,omitempty"`
	Source  *SourceName      `json:"source,omitempty"`
}

type PgCreateArrayRelationshipArgs struct {
	Table   ITableName       `json:"table"`
	Name    RelationshipName `json:"name"`
	Using   ArrRelUsing      `json:"using"`
	Comment string           `json:"comment,omitempty"`
	Source  *SourceName      `json:"source,omitempty"`
}

type PgDropRelationshipArgs struct {
	Table        ITableName       `json:"table"`
	Relationship RelationshipName `json:"relationship"`
	Cascade      bool             `json:"cascade,omitempty"`
	Source       *SourceName      `json:"source,omitempty"`
}

type PgRenameRelationshipArgs struct{}

type PgSetRelationshipCommentArgs struct{}

type PgAddComputedFieldArgs struct{}

type PgDropComputedFieldArgs struct{}

type PgCreateInsertPermissionArgs struct {
	Table      ITableName       `json:"table"`
	Role       RoleName         `json:"role"`
	Permission InsertPermission `json:"permission"`
	Comment    string           `json:"comment,omitempty"`
	Source     *SourceName      `json:"source,omitempty"`
}

type PgDropInsertPermissionArgs struct{}

type PgCreateSelectPermissionArgs struct{}

type PgDropSelectPermissionArgs struct{}

type PgCreateUpdatePermissionArgs struct{}

type PgDropUpdatePermissionArgs struct{}

type PgCreateDeletePermissionArgs struct{}

type PgDropDeletePermissionArgs struct{}

type PgSetPermissionCommentArgs struct{}

type PgCreateEventTriggerArgs struct{}

type PgDeleteEventTriggerArgs struct{}

type PgRedeliverEventArgs struct{}

type PgInvokeEventTriggerArgs struct{}

type BigqueryTrackTableArgs struct{}

type BigqueryUntrackTableArgs struct{}

type BigquerySetTableCustomizationArgs struct{}

type MssqlAddSourceArgs struct{}

type MssqlDropSourceArgs struct{}

type MssqlTrackTableArgs struct{}

type MssqlUntrackTableArgs struct{}

type MssqlCreateObjectRelationshipArgs struct{}

type MssqlCreateArrayRelationshipArgs struct{}

type MssqlDropRelationshipArgs struct{}

type MssqlRenameRelationshipArgs struct{}

type MssqlSetRelationshipCommentArgs struct{}

type MssqlSetTableCustomizationArgs struct{}

type MssqlCreateInsertPermissionArgs struct{}

type MssqlDropInsertPermissionArgs struct{}

type MssqlCreateSelectPermissionArgs struct{}

type MssqlDropSelectPermissionArgs struct{}

type MssqlCreateUpdatePermissionArgs struct{}

type MssqlDropUpdatePermissionArgs struct{}

type MssqlCreateDeletePermissionArgs struct{}

type MssqlDropDeletePermissionArgs struct{}

type MssqlSetPermissionCommentArgs struct{}

type CreateCronTriggerArgs struct{}

type DeleteCronTriggerArgs struct{}

type GetCronTriggersArgs struct{}

type CreateScheduledEventArgs struct{}

type DeleteScheduledEventArgs struct{}

type AddRemoteSchemaArgs struct{}

type UpdateRemoteSchemaArgs struct{}

type RemoveRemoteSchemaArgs struct{}

type ReloadRemoteSchemaArgs struct{}

type AddRemoteSchemaPermissionsArgs struct{}

type DropRemoteSchemaPermissionsArgs struct{}

type PgCreateRemoteRelationshipArgs struct{}

type PgUpdateRemoteRelationshipArgs struct{}

type PgDeleteRemoteRelationshipArgs struct{}

type ExportMetadataArgs struct{}

type ReplaceMetadataArgs struct{}

type ReloadMetadataArgs struct{}

type ClearMetadataArgs struct{}

type GetInconsistentMetadataArgs struct{}

type DropInconsistentMetadataArgs struct{}

type CreateQueryCollectionArgs struct{}

type DropQueryCollectionArgs struct{}

type AddQueryToCollectionArgs struct{}

type DropQueryFromCollectionArgs struct{}

type AddCollectionToAllowlistArgs struct{}

type DropCollectionFromAllowlistArgs struct{}

type SetCustomTypesArgs struct{}

type CreateActionArgs struct{}

type DropActionArgs struct{}

type UpdateActionArgs struct{}

type CreateActionPermissionArgs struct{}

type DropActionPermissionArgs struct{}

type CreateRestEndpointArgs struct{}

type DropRestEndpointArgs struct{}

type AddInheritedRoleArgs struct{}

type DropInheritedRoleArgs struct{}

type SetGraphqlIntrospectionOptionsArgs struct{}

type AddHostToTlsAllowlistArgs struct{}

type DropHostFromTlsAllowlistArgs struct{}
