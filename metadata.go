package hasura_api

import "github.com/minskylab/hasura-api/metadata"

func (r *HasuraClient) Bulk(body []metadata.MetadataQuery) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.BulkType,
		Args: body,
	})
}

func (r *HasuraClient) PgAddSource(body *metadata.PgAddSourceArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgAddSource,
		Args: body,
	})
}

func (r *HasuraClient) PgDropSource(body *metadata.PgDropSourceArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgDropSource,
		Args: body,
	})
}

func (r *HasuraClient) PgTrackTable(body *metadata.PgTrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgTrackTable,
		Args: body,
	})
}

func (r *HasuraClient) PgUntrackTable(body *metadata.PgUntrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgUntrackTable,
		Args: body,
	})
}

func (r *HasuraClient) PgSetTableCustomization(body *metadata.PgSetTableCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgSetTableCustomization,
		Args: body,
	})
}

func (r *HasuraClient) PgSetTableIsEnum(body *metadata.PgSetTableIsEnumArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgSetTableIsEnum,
		Args: body,
	})
}

func (r *HasuraClient) PgTrackFunction(body *metadata.PgTrackFunctionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgTrackFunction,
		Args: body,
	})
}

func (r *HasuraClient) PgUntrackFunction(body *metadata.PgUntrackFunctionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgUntrackFunction,
		Args: body,
	})
}

func (r *HasuraClient) PgSetFunctionCustomization(body *metadata.PgSetFunctionCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgSetFunctionCustomization,
		Args: body,
	})
}

func (r *HasuraClient) PgCreateFunctionPermission(body *metadata.PgCreateFunctionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgCreateFunctionPermission,
		Args: body,
	})
}

func (r *HasuraClient) PgDropFunctionPermission(body *metadata.PgDropFunctionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgDropFunctionPermission,
		Args: body,
	})
}

func (r *HasuraClient) PgCreateObjectRelationship(body *metadata.PgCreateObjectRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgCreateObjectRelationship,
		Args: body,
	})
}

func (r *HasuraClient) PgCreateArrayRelationship(body *metadata.PgCreateArrayRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgCreateArrayRelationship,
		Args: body,
	})
}

func (r *HasuraClient) PgDropRelationship(body *metadata.PgDropRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgDropRelationship,
		Args: body,
	})
}

func (r *HasuraClient) PgRenameRelationship(body *metadata.PgRenameRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgRenameRelationship,
		Args: body,
	})
}

func (r *HasuraClient) PgSetRelationshipComment(body *metadata.PgSetRelationshipCommentArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgSetRelationshipComment,
		Args: body,
	})
}

func (r *HasuraClient) PgAddComputedField(body *metadata.PgAddComputedFieldArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgAddComputedField,
		Args: body,
	})
}

func (r *HasuraClient) PgDropComputedField(body *metadata.PgDropComputedFieldArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgDropComputedField,
		Args: body,
	})
}

func (r *HasuraClient) PgCreateInsertPermission(body *metadata.PgCreateInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgCreateInsertPermission,
		Args: body,
	})
}

func (r *HasuraClient) PgDropInsertPermission(body *metadata.PgDropInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgDropInsertPermission,
		Args: body,
	})
}

func (r *HasuraClient) PgCreateSelectPermission(body *metadata.PgCreateSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgCreateSelectPermission,
		Args: body,
	})
}

func (r *HasuraClient) PgDropSelectPermission(body *metadata.PgDropSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgDropSelectPermission,
		Args: body,
	})
}

func (r *HasuraClient) PgCreateUpdatePermission(body *metadata.PgCreateUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgCreateUpdatePermission,
		Args: body,
	})
}

func (r *HasuraClient) PgDropUpdatePermission(body *metadata.PgDropUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgDropUpdatePermission,
		Args: body,
	})
}

func (r *HasuraClient) PgCreateDeletePermission(body *metadata.PgCreateDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgCreateDeletePermission,
		Args: body,
	})
}

func (r *HasuraClient) PgDropDeletePermission(body *metadata.PgDropDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgDropDeletePermission,
		Args: body,
	})
}

func (r *HasuraClient) PgSetPermissionComment(body *metadata.PgSetPermissionCommentArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgSetPermissionComment,
		Args: body,
	})
}

func (r *HasuraClient) PgCreateEventTrigger(body *metadata.PgCreateEventTriggerArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgCreateEventTrigger,
		Args: body,
	})
}

func (r *HasuraClient) PgDeleteEventTrigger(body *metadata.PgDeleteEventTriggerArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgDeleteEventTrigger,
		Args: body,
	})
}

func (r *HasuraClient) PgRedeliverEvent(body *metadata.PgRedeliverEventArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgRedeliverEvent,
		Args: body,
	})
}

func (r *HasuraClient) PgInvokeEventTrigger(body *metadata.PgInvokeEventTriggerArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgInvokeEventTrigger,
		Args: body,
	})
}

func (r *HasuraClient) BigqueryTrackTable(body *metadata.BigqueryTrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.BigqueryTrackTable,
		Args: body,
	})
}

func (r *HasuraClient) BigqueryUntrackTable(body *metadata.BigqueryUntrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.BigqueryUntrackTable,
		Args: body,
	})
}

func (r *HasuraClient) BigquerySetTableCustomization(body *metadata.BigquerySetTableCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.BigquerySetTableCustomization,
		Args: body,
	})
}

func (r *HasuraClient) MssqlAddSource(body *metadata.MssqlAddSourceArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlAddSource,
		Args: body,
	})
}

func (r *HasuraClient) MssqlDropSource(body *metadata.MssqlDropSourceArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlDropSource,
		Args: body,
	})
}

func (r *HasuraClient) MssqlTrackTable(body *metadata.MssqlTrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlTrackTable,
		Args: body,
	})
}

func (r *HasuraClient) MssqlUntrackTable(body *metadata.MssqlUntrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlUntrackTable,
		Args: body,
	})
}

func (r *HasuraClient) MssqlCreateObjectRelationship(body *metadata.MssqlCreateObjectRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlCreateObjectRelationship,
		Args: body,
	})
}

func (r *HasuraClient) MssqlCreateArrayRelationship(body *metadata.MssqlCreateArrayRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlCreateArrayRelationship,
		Args: body,
	})
}

func (r *HasuraClient) MssqlDropRelationship(body *metadata.MssqlDropRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlDropRelationship,
		Args: body,
	})
}

func (r *HasuraClient) MssqlRenameRelationship(body *metadata.MssqlRenameRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlRenameRelationship,
		Args: body,
	})
}

func (r *HasuraClient) MssqlSetRelationshipComment(body *metadata.MssqlSetRelationshipCommentArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlSetRelationshipComment,
		Args: body,
	})
}

func (r *HasuraClient) MssqlSetTableCustomization(body *metadata.MssqlSetTableCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlSetTableCustomization,
		Args: body,
	})
}

func (r *HasuraClient) MssqlCreateInsertPermission(body *metadata.MssqlCreateInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlCreateInsertPermission,
		Args: body,
	})
}

func (r *HasuraClient) MssqlDropInsertPermission(body *metadata.MssqlDropInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlDropInsertPermission,
		Args: body,
	})
}

func (r *HasuraClient) MssqlCreateSelectPermission(body *metadata.MssqlCreateSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlCreateSelectPermission,
		Args: body,
	})
}

func (r *HasuraClient) MssqlDropSelectPermission(body *metadata.MssqlDropSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlDropSelectPermission,
		Args: body,
	})
}

func (r *HasuraClient) MssqlCreateUpdatePermission(body *metadata.MssqlCreateUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlCreateUpdatePermission,
		Args: body,
	})
}

func (r *HasuraClient) MssqlDropUpdatePermission(body *metadata.MssqlDropUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlDropUpdatePermission,
		Args: body,
	})
}

func (r *HasuraClient) MssqlCreateDeletePermission(body *metadata.MssqlCreateDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlCreateDeletePermission,
		Args: body,
	})
}

func (r *HasuraClient) MssqlDropDeletePermission(body *metadata.MssqlDropDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlDropDeletePermission,
		Args: body,
	})
}

func (r *HasuraClient) MssqlSetPermissionComment(body *metadata.MssqlSetPermissionCommentArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.MssqlSetPermissionComment,
		Args: body,
	})
}

func (r *HasuraClient) CreateCronTrigger(body *metadata.CreateCronTriggerArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.CreateCronTrigger,
		Args: body,
	})
}

func (r *HasuraClient) DeleteCronTrigger(body *metadata.DeleteCronTriggerArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DeleteCronTrigger,
		Args: body,
	})
}

func (r *HasuraClient) GetCronTriggers(body *metadata.GetCronTriggersArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.GetCronTriggers,
		Args: body,
	})
}

func (r *HasuraClient) CreateScheduledEvent(body *metadata.CreateScheduledEventArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.CreateScheduledEvent,
		Args: body,
	})
}

func (r *HasuraClient) DeleteScheduledEvent(body *metadata.DeleteScheduledEventArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DeleteScheduledEvent,
		Args: body,
	})
}

func (r *HasuraClient) AddRemoteSchema(body *metadata.AddRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.AddRemoteSchema,
		Args: body,
	})
}

func (r *HasuraClient) UpdateRemoteSchema(body *metadata.UpdateRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.UpdateRemoteSchema,
		Args: body,
	})
}

func (r *HasuraClient) RemoveRemoteSchema(body *metadata.RemoveRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.RemoveRemoteSchema,
		Args: body,
	})
}

func (r *HasuraClient) ReloadRemoteSchema(body *metadata.ReloadRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.ReloadRemoteSchema,
		Args: body,
	})
}

func (r *HasuraClient) AddRemoteSchemaPermissions(body *metadata.AddRemoteSchemaPermissionsArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.AddRemoteSchemaPermissions,
		Args: body,
	})
}

func (r *HasuraClient) DropRemoteSchemaPermissions(body *metadata.DropRemoteSchemaPermissionsArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DropRemoteSchemaPermissions,
		Args: body,
	})
}

func (r *HasuraClient) PgCreateRemoteRelationship(body *metadata.PgCreateRemoteRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgCreateRemoteRelationship,
		Args: body,
	})
}

func (r *HasuraClient) PgUpdateRemoteRelationship(body *metadata.PgUpdateRemoteRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgUpdateRemoteRelationship,
		Args: body,
	})
}

func (r *HasuraClient) PgDeleteRemoteRelationship(body *metadata.PgDeleteRemoteRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.PgDeleteRemoteRelationship,
		Args: body,
	})
}

func (r *HasuraClient) ExportMetadata(body *metadata.ExportMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.ExportMetadata,
		Args: body,
	})
}

func (r *HasuraClient) ReplaceMetadata(body *metadata.ReplaceMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.ReplaceMetadata,
		Args: body,
	})
}

func (r *HasuraClient) ReloadMetadata(body *metadata.ReloadMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.ReloadMetadata,
		Args: body,
	})
}

func (r *HasuraClient) ClearMetadata(body *metadata.ClearMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.ClearMetadata,
		Args: body,
	})
}

func (r *HasuraClient) GetInconsistentMetadata(body *metadata.GetInconsistentMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.GetInconsistentMetadata,
		Args: body,
	})
}

func (r *HasuraClient) DropInconsistentMetadata(body *metadata.DropInconsistentMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DropInconsistentMetadata,
		Args: body,
	})
}

func (r *HasuraClient) CreateQueryCollection(body *metadata.CreateQueryCollectionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.CreateQueryCollection,
		Args: body,
	})
}

func (r *HasuraClient) DropQueryCollection(body *metadata.DropQueryCollectionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DropQueryCollection,
		Args: body,
	})
}

func (r *HasuraClient) AddQueryToCollection(body *metadata.AddQueryToCollectionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.AddQueryToCollection,
		Args: body,
	})
}

func (r *HasuraClient) DropQueryFromCollection(body *metadata.DropQueryFromCollectionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DropQueryFromCollection,
		Args: body,
	})
}

func (r *HasuraClient) AddCollectionToAllowlist(body *metadata.AddCollectionToAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.AddCollectionToAllowlist,
		Args: body,
	})
}

func (r *HasuraClient) DropCollectionFromAllowlist(body *metadata.DropCollectionFromAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DropCollectionFromAllowlist,
		Args: body,
	})
}

func (r *HasuraClient) SetCustomTypes(body *metadata.SetCustomTypesArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.SetCustomTypes,
		Args: body,
	})
}

func (r *HasuraClient) CreateAction(body *metadata.CreateActionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.CreateAction,
		Args: body,
	})
}

func (r *HasuraClient) DropAction(body *metadata.DropActionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DropAction,
		Args: body,
	})
}

func (r *HasuraClient) UpdateAction(body *metadata.UpdateActionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.UpdateAction,
		Args: body,
	})
}

func (r *HasuraClient) CreateActionPermission(body *metadata.CreateActionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.CreateActionPermission,
		Args: body,
	})
}

func (r *HasuraClient) DropActionPermission(body *metadata.DropActionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DropActionPermission,
		Args: body,
	})
}

func (r *HasuraClient) CreateRestEndpoint(body *metadata.CreateRestEndpointArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.CreateRestEndpoint,
		Args: body,
	})
}

func (r *HasuraClient) DropRestEndpoint(body *metadata.DropRestEndpointArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DropRestEndpoint,
		Args: body,
	})
}

func (r *HasuraClient) AddInheritedRole(body *metadata.AddInheritedRoleArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.AddInheritedRole,
		Args: body,
	})
}

func (r *HasuraClient) DropInheritedRole(body *metadata.DropInheritedRoleArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DropInheritedRole,
		Args: body,
	})
}

func (r *HasuraClient) SetGraphqlIntrospectionOptions(body *metadata.SetGraphqlIntrospectionOptionsArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.SetGraphqlIntrospectionOptions,
		Args: body,
	})
}

func (r *HasuraClient) AddHostToTlsAllowlist(body *metadata.AddHostToTlsAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.AddHostToTlsAllowlist,
		Args: body,
	})
}

func (r *HasuraClient) DropHostFromTlsAllowlist(body *metadata.DropHostFromTlsAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(metadata.MetadataQuery{
		Type: metadata.DropHostFromTlsAllowlist,
		Args: body,
	})
}
