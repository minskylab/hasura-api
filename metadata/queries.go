package metadata

func BulkQuery(args BulkArgs) MetadataQuery {
	return MetadataQuery{
		Type: BulkType,
		Args: args,
	}
}

func PgAddSourceQuery(args *PgAddSourceArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgAddSource,
		Args: args,
	}
}

func PgDropSourceQuery(args *PgDropSourceArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgDropSource,
		Args: args,
	}
}

func PgTrackTableQuery(args *PgTrackTableArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgTrackTable,
		Args: args,
	}
}

func PgUntrackTableQuery(args *PgUntrackTableArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgUntrackTable,
		Args: args,
	}
}

func PgSetTableCustomizationQuery(args *PgSetTableCustomizationArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgSetTableCustomization,
		Args: args,
	}
}

func PgSetTableIsEnumQuery(args *PgSetTableIsEnumArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgSetTableIsEnum,
		Args: args,
	}
}

func PgTrackFunctionQuery(args *PgTrackFunctionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgTrackFunction,
		Args: args,
	}
}

func PgUntrackFunctionQuery(args *PgUntrackFunctionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgUntrackFunction,
		Args: args,
	}
}

func PgSetFunctionCustomizationQuery(args *PgSetFunctionCustomizationArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgSetFunctionCustomization,
		Args: args,
	}
}

func PgCreateFunctionPermissionQuery(args *PgCreateFunctionPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgCreateFunctionPermission,
		Args: args,
	}
}

func PgDropFunctionPermissionQuery(args *PgDropFunctionPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgDropFunctionPermission,
		Args: args,
	}
}

func PgCreateObjectRelationshipQuery(args *PgCreateObjectRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgCreateObjectRelationship,
		Args: args,
	}
}

func PgCreateArrayRelationshipQuery(args *PgCreateArrayRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgCreateArrayRelationship,
		Args: args,
	}
}

func PgDropRelationshipQuery(args *PgDropRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgDropRelationship,
		Args: args,
	}
}

func PgRenameRelationshipQuery(args *PgRenameRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgRenameRelationship,
		Args: args,
	}
}

func PgSetRelationshipCommentQuery(args *PgSetRelationshipCommentArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgSetRelationshipComment,
		Args: args,
	}
}

func PgAddComputedFieldQuery(args *PgAddComputedFieldArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgAddComputedField,
		Args: args,
	}
}

func PgDropComputedFieldQuery(args *PgDropComputedFieldArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgDropComputedField,
		Args: args,
	}
}

func PgCreateInsertPermissionQuery(args *PgCreateInsertPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgCreateInsertPermission,
		Args: args,
	}
}

func PgDropInsertPermissionQuery(args *PgDropInsertPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgDropInsertPermission,
		Args: args,
	}
}

func PgCreateSelectPermissionQuery(args *PgCreateSelectPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgCreateSelectPermission,
		Args: args,
	}
}

func PgDropSelectPermissionQuery(args *PgDropSelectPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgDropSelectPermission,
		Args: args,
	}
}

func PgCreateUpdatePermissionQuery(args *PgCreateUpdatePermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgCreateUpdatePermission,
		Args: args,
	}
}

func PgDropUpdatePermissionQuery(args *PgDropUpdatePermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgDropUpdatePermission,
		Args: args,
	}
}

func PgCreateDeletePermissionQuery(args *PgCreateDeletePermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgCreateDeletePermission,
		Args: args,
	}
}

func PgDropDeletePermissionQuery(args *PgDropDeletePermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgDropDeletePermission,
		Args: args,
	}
}

func PgSetPermissionCommentQuery(args *PgSetPermissionCommentArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgSetPermissionComment,
		Args: args,
	}
}

func PgCreateEventTriggerQuery(args *PgCreateEventTriggerArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgCreateEventTrigger,
		Args: args,
	}
}

func PgDeleteEventTriggerQuery(args *PgDeleteEventTriggerArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgDeleteEventTrigger,
		Args: args,
	}
}

func PgRedeliverEventQuery(args *PgRedeliverEventArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgRedeliverEvent,
		Args: args,
	}
}

func PgInvokeEventTriggerQuery(args *PgInvokeEventTriggerArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgInvokeEventTrigger,
		Args: args,
	}
}

func BigqueryTrackTableQuery(args *BigqueryTrackTableArgs) MetadataQuery {
	return MetadataQuery{
		Type: BigqueryTrackTable,
		Args: args,
	}
}

func BigqueryUntrackTableQuery(args *BigqueryUntrackTableArgs) MetadataQuery {
	return MetadataQuery{
		Type: BigqueryUntrackTable,
		Args: args,
	}
}

func BigquerySetTableCustomizationQuery(args *BigquerySetTableCustomizationArgs) MetadataQuery {
	return MetadataQuery{
		Type: BigquerySetTableCustomization,
		Args: args,
	}
}

func MssqlAddSourceQuery(args *MssqlAddSourceArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlAddSource,
		Args: args,
	}
}

func MssqlDropSourceQuery(args *MssqlDropSourceArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlDropSource,
		Args: args,
	}
}

func MssqlTrackTableQuery(args *MssqlTrackTableArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlTrackTable,
		Args: args,
	}
}

func MssqlUntrackTableQuery(args *MssqlUntrackTableArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlUntrackTable,
		Args: args,
	}
}

func MssqlCreateObjectRelationshipQuery(args *MssqlCreateObjectRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlCreateObjectRelationship,
		Args: args,
	}
}

func MssqlCreateArrayRelationshipQuery(args *MssqlCreateArrayRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlCreateArrayRelationship,
		Args: args,
	}
}

func MssqlDropRelationshipQuery(args *MssqlDropRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlDropRelationship,
		Args: args,
	}
}

func MssqlRenameRelationshipQuery(args *MssqlRenameRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlRenameRelationship,
		Args: args,
	}
}

func MssqlSetRelationshipCommentQuery(args *MssqlSetRelationshipCommentArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlSetRelationshipComment,
		Args: args,
	}
}

func MssqlSetTableCustomizationQuery(args *MssqlSetTableCustomizationArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlSetTableCustomization,
		Args: args,
	}
}

func MssqlCreateInsertPermissionQuery(args *MssqlCreateInsertPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlCreateInsertPermission,
		Args: args,
	}
}

func MssqlDropInsertPermissionQuery(args *MssqlDropInsertPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlDropInsertPermission,
		Args: args,
	}
}

func MssqlCreateSelectPermissionQuery(args *MssqlCreateSelectPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlCreateSelectPermission,
		Args: args,
	}
}

func MssqlDropSelectPermissionQuery(args *MssqlDropSelectPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlDropSelectPermission,
		Args: args,
	}
}

func MssqlCreateUpdatePermissionQuery(args *MssqlCreateUpdatePermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlCreateUpdatePermission,
		Args: args,
	}
}

func MssqlDropUpdatePermissionQuery(args *MssqlDropUpdatePermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlDropUpdatePermission,
		Args: args,
	}
}

func MssqlCreateDeletePermissionQuery(args *MssqlCreateDeletePermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlCreateDeletePermission,
		Args: args,
	}
}

func MssqlDropDeletePermissionQuery(args *MssqlDropDeletePermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlDropDeletePermission,
		Args: args,
	}
}

func MssqlSetPermissionCommentQuery(args *MssqlSetPermissionCommentArgs) MetadataQuery {
	return MetadataQuery{
		Type: MssqlSetPermissionComment,
		Args: args,
	}
}

func CreateCronTriggerQuery(args *CreateCronTriggerArgs) MetadataQuery {
	return MetadataQuery{
		Type: CreateCronTrigger,
		Args: args,
	}
}

func DeleteCronTriggerQuery(args *DeleteCronTriggerArgs) MetadataQuery {
	return MetadataQuery{
		Type: DeleteCronTrigger,
		Args: args,
	}
}

func GetCronTriggersQuery(args *GetCronTriggersArgs) MetadataQuery {
	return MetadataQuery{
		Type: GetCronTriggers,
		Args: args,
	}
}

func CreateScheduledEventQuery(args *CreateScheduledEventArgs) MetadataQuery {
	return MetadataQuery{
		Type: CreateScheduledEvent,
		Args: args,
	}
}

func DeleteScheduledEventQuery(args *DeleteScheduledEventArgs) MetadataQuery {
	return MetadataQuery{
		Type: DeleteScheduledEvent,
		Args: args,
	}
}

func AddRemoteSchemaQuery(args *AddRemoteSchemaArgs) MetadataQuery {
	return MetadataQuery{
		Type: AddRemoteSchema,
		Args: args,
	}
}

func UpdateRemoteSchemaQuery(args *UpdateRemoteSchemaArgs) MetadataQuery {
	return MetadataQuery{
		Type: UpdateRemoteSchema,
		Args: args,
	}
}

func RemoveRemoteSchemaQuery(args *RemoveRemoteSchemaArgs) MetadataQuery {
	return MetadataQuery{
		Type: RemoveRemoteSchema,
		Args: args,
	}
}

func ReloadRemoteSchemaQuery(args *ReloadRemoteSchemaArgs) MetadataQuery {
	return MetadataQuery{
		Type: ReloadRemoteSchema,
		Args: args,
	}
}

func AddRemoteSchemaPermissionsQuery(args *AddRemoteSchemaPermissionsArgs) MetadataQuery {
	return MetadataQuery{
		Type: AddRemoteSchemaPermissions,
		Args: args,
	}
}

func DropRemoteSchemaPermissionsQuery(args *DropRemoteSchemaPermissionsArgs) MetadataQuery {
	return MetadataQuery{
		Type: DropRemoteSchemaPermissions,
		Args: args,
	}
}

func PgCreateRemoteRelationshipQuery(args *PgCreateRemoteRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgCreateRemoteRelationship,
		Args: args,
	}
}

func PgUpdateRemoteRelationshipQuery(args *PgUpdateRemoteRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgUpdateRemoteRelationship,
		Args: args,
	}
}

func PgDeleteRemoteRelationshipQuery(args *PgDeleteRemoteRelationshipArgs) MetadataQuery {
	return MetadataQuery{
		Type: PgDeleteRemoteRelationship,
		Args: args,
	}
}

func ExportMetadataQuery(args *ExportMetadataArgs) MetadataQuery {
	return MetadataQuery{
		Type: ExportMetadata,
		Args: args,
	}
}

func ReplaceMetadataQuery(args *ReplaceMetadataArgs) MetadataQuery {
	return MetadataQuery{
		Type: ReplaceMetadata,
		Args: args,
	}
}

func ReloadMetadataQuery(args *ReloadMetadataArgs) MetadataQuery {
	return MetadataQuery{
		Type: ReloadMetadata,
		Args: args,
	}
}

func ClearMetadataQuery(args *ClearMetadataArgs) MetadataQuery {
	return MetadataQuery{
		Type: ClearMetadata,
		Args: args,
	}
}

func GetInconsistentMetadataQuery(args *GetInconsistentMetadataArgs) MetadataQuery {
	return MetadataQuery{
		Type: GetInconsistentMetadata,
		Args: args,
	}
}

func DropInconsistentMetadataQuery(args *DropInconsistentMetadataArgs) MetadataQuery {
	return MetadataQuery{
		Type: DropInconsistentMetadata,
		Args: args,
	}
}

func CreateQueryCollectionQuery(args *CreateQueryCollectionArgs) MetadataQuery {
	return MetadataQuery{
		Type: CreateQueryCollection,
		Args: args,
	}
}

func DropQueryCollectionQuery(args *DropQueryCollectionArgs) MetadataQuery {
	return MetadataQuery{
		Type: DropQueryCollection,
		Args: args,
	}
}

func AddQueryToCollectionQuery(args *AddQueryToCollectionArgs) MetadataQuery {
	return MetadataQuery{
		Type: AddQueryToCollection,
		Args: args,
	}
}

func DropQueryFromCollectionQuery(args *DropQueryFromCollectionArgs) MetadataQuery {
	return MetadataQuery{
		Type: DropQueryFromCollection,
		Args: args,
	}
}

func AddCollectionToAllowlistQuery(args *AddCollectionToAllowlistArgs) MetadataQuery {
	return MetadataQuery{
		Type: AddCollectionToAllowlist,
		Args: args,
	}
}

func DropCollectionFromAllowlistQuery(args *DropCollectionFromAllowlistArgs) MetadataQuery {
	return MetadataQuery{
		Type: DropCollectionFromAllowlist,
		Args: args,
	}
}

func SetCustomTypesQuery(args *SetCustomTypesArgs) MetadataQuery {
	return MetadataQuery{
		Type: SetCustomTypes,
		Args: args,
	}
}

func CreateActionQuery(args *CreateActionArgs) MetadataQuery {
	return MetadataQuery{
		Type: CreateAction,
		Args: args,
	}
}

func DropActionQuery(args *DropActionArgs) MetadataQuery {
	return MetadataQuery{
		Type: DropAction,
		Args: args,
	}
}

func UpdateActionQuery(args *UpdateActionArgs) MetadataQuery {
	return MetadataQuery{
		Type: UpdateAction,
		Args: args,
	}
}

func CreateActionPermissionQuery(args *CreateActionPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: CreateActionPermission,
		Args: args,
	}
}

func DropActionPermissionQuery(args *DropActionPermissionArgs) MetadataQuery {
	return MetadataQuery{
		Type: DropActionPermission,
		Args: args,
	}
}

func CreateRestEndpointQuery(args *CreateRestEndpointArgs) MetadataQuery {
	return MetadataQuery{
		Type: CreateRestEndpoint,
		Args: args,
	}
}

func DropRestEndpointQuery(args *DropRestEndpointArgs) MetadataQuery {
	return MetadataQuery{
		Type: DropRestEndpoint,
		Args: args,
	}
}

func AddInheritedRoleQuery(args *AddInheritedRoleArgs) MetadataQuery {
	return MetadataQuery{
		Type: AddInheritedRole,
		Args: args,
	}
}

func DropInheritedRoleQuery(args *DropInheritedRoleArgs) MetadataQuery {
	return MetadataQuery{
		Type: DropInheritedRole,
		Args: args,
	}
}

func SetGraphqlIntrospectionOptionsQuery(args *SetGraphqlIntrospectionOptionsArgs) MetadataQuery {
	return MetadataQuery{
		Type: SetGraphqlIntrospectionOptions,
		Args: args,
	}
}

func AddHostToTlsAllowlistQuery(args *AddHostToTlsAllowlistArgs) MetadataQuery {
	return MetadataQuery{
		Type: AddHostToTlsAllowlist,
		Args: args,
	}
}

func DropHostFromTlsAllowlistQuery(args *DropHostFromTlsAllowlistArgs) MetadataQuery {
	return MetadataQuery{
		Type: DropHostFromTlsAllowlist,
		Args: args,
	}
}
