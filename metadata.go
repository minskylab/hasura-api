package hasura_api

import "github.com/minskylab/hasura-api/metadata"

func (r *MetadataClient) Exec(body metadata.MetadataQuery) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(body)
}

func (r *MetadataClient) Bulk(body []metadata.MetadataQuery) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.BulkQuery(metadata.BulkArgs(body)))
}

func (r *MetadataClient) PgAddSource(body *metadata.PgAddSourceArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgAddSourceQuery(body))
}

func (r *MetadataClient) PgDropSource(body *metadata.PgDropSourceArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgDropSourceQuery(body))
}

func (r *MetadataClient) PgTrackTable(body *metadata.PgTrackTableArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgTrackTableQuery(body))
}

func (r *MetadataClient) PgUntrackTable(body *metadata.PgUntrackTableArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgUntrackTableQuery(body))
}

func (r *MetadataClient) PgSetTableCustomization(body *metadata.PgSetTableCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgSetTableCustomizationQuery(body))
}

func (r *MetadataClient) PgSetTableIsEnum(body *metadata.PgSetTableIsEnumArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgSetTableIsEnumQuery(body))
}

func (r *MetadataClient) PgTrackFunction(body *metadata.PgTrackFunctionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgTrackFunctionQuery(body))
}

func (r *MetadataClient) PgUntrackFunction(body *metadata.PgUntrackFunctionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgUntrackFunctionQuery(body))
}

func (r *MetadataClient) PgSetFunctionCustomization(body *metadata.PgSetFunctionCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgSetFunctionCustomizationQuery(body))
}

func (r *MetadataClient) PgCreateFunctionPermission(body *metadata.PgCreateFunctionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgCreateFunctionPermissionQuery(body))
}

func (r *MetadataClient) PgDropFunctionPermission(body *metadata.PgDropFunctionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgDropFunctionPermissionQuery(body))
}

func (r *MetadataClient) PgCreateObjectRelationship(body *metadata.PgCreateObjectRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgCreateObjectRelationshipQuery(body))
}

func (r *MetadataClient) PgCreateArrayRelationship(body *metadata.PgCreateArrayRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgCreateArrayRelationshipQuery(body))
}

func (r *MetadataClient) PgDropRelationship(body *metadata.PgDropRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgDropRelationshipQuery(body))
}

func (r *MetadataClient) PgRenameRelationship(body *metadata.PgRenameRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgRenameRelationshipQuery(body))
}

func (r *MetadataClient) PgSetRelationshipComment(body *metadata.PgSetRelationshipCommentArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgSetRelationshipCommentQuery(body))
}

func (r *MetadataClient) PgAddComputedField(body *metadata.PgAddComputedFieldArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgAddComputedFieldQuery(body))
}

func (r *MetadataClient) PgDropComputedField(body *metadata.PgDropComputedFieldArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgDropComputedFieldQuery(body))
}

func (r *MetadataClient) PgCreateInsertPermission(body *metadata.PgCreateInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgCreateInsertPermissionQuery(body))
}

func (r *MetadataClient) PgDropInsertPermission(body *metadata.PgDropInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgDropInsertPermissionQuery(body))
}

func (r *MetadataClient) PgCreateSelectPermission(body *metadata.PgCreateSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgCreateSelectPermissionQuery(body))
}

func (r *MetadataClient) PgDropSelectPermission(body *metadata.PgDropSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgDropSelectPermissionQuery(body))
}

func (r *MetadataClient) PgCreateUpdatePermission(body *metadata.PgCreateUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgCreateUpdatePermissionQuery(body))
}

func (r *MetadataClient) PgDropUpdatePermission(body *metadata.PgDropUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgDropUpdatePermissionQuery(body))
}

func (r *MetadataClient) PgCreateDeletePermission(body *metadata.PgCreateDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgCreateDeletePermissionQuery(body))
}

func (r *MetadataClient) PgDropDeletePermission(body *metadata.PgDropDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgDropDeletePermissionQuery(body))
}

func (r *MetadataClient) PgSetPermissionComment(body *metadata.PgSetPermissionCommentArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgSetPermissionCommentQuery(body))
}

func (r *MetadataClient) PgCreateEventTrigger(body *metadata.PgCreateEventTriggerArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgCreateEventTriggerQuery(body))
}

func (r *MetadataClient) PgDeleteEventTrigger(body *metadata.PgDeleteEventTriggerArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgDeleteEventTriggerQuery(body))
}

func (r *MetadataClient) PgRedeliverEvent(body *metadata.PgRedeliverEventArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgRedeliverEventQuery(body))
}

func (r *MetadataClient) PgInvokeEventTrigger(body *metadata.PgInvokeEventTriggerArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgInvokeEventTriggerQuery(body))
}

func (r *MetadataClient) BigqueryTrackTable(body *metadata.BigqueryTrackTableArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.BigqueryTrackTableQuery(body))
}

func (r *MetadataClient) BigqueryUntrackTable(body *metadata.BigqueryUntrackTableArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.BigqueryUntrackTableQuery(body))
}

func (r *MetadataClient) BigquerySetTableCustomization(body *metadata.BigquerySetTableCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.BigquerySetTableCustomizationQuery(body))
}

func (r *MetadataClient) MssqlAddSource(body *metadata.MssqlAddSourceArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlAddSourceQuery(body))
}

func (r *MetadataClient) MssqlDropSource(body *metadata.MssqlDropSourceArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlDropSourceQuery(body))
}

func (r *MetadataClient) MssqlTrackTable(body *metadata.MssqlTrackTableArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlTrackTableQuery(body))
}

func (r *MetadataClient) MssqlUntrackTable(body *metadata.MssqlUntrackTableArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlUntrackTableQuery(body))
}

func (r *MetadataClient) MssqlCreateObjectRelationship(body *metadata.MssqlCreateObjectRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlCreateObjectRelationshipQuery(body))
}

func (r *MetadataClient) MssqlCreateArrayRelationship(body *metadata.MssqlCreateArrayRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlCreateArrayRelationshipQuery(body))
}

func (r *MetadataClient) MssqlDropRelationship(body *metadata.MssqlDropRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlDropRelationshipQuery(body))
}

func (r *MetadataClient) MssqlRenameRelationship(body *metadata.MssqlRenameRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlRenameRelationshipQuery(body))
}

func (r *MetadataClient) MssqlSetRelationshipComment(body *metadata.MssqlSetRelationshipCommentArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlSetRelationshipCommentQuery(body))
}

func (r *MetadataClient) MssqlSetTableCustomization(body *metadata.MssqlSetTableCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlSetTableCustomizationQuery(body))
}

func (r *MetadataClient) MssqlCreateInsertPermission(body *metadata.MssqlCreateInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlCreateInsertPermissionQuery(body))
}

func (r *MetadataClient) MssqlDropInsertPermission(body *metadata.MssqlDropInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlDropInsertPermissionQuery(body))
}

func (r *MetadataClient) MssqlCreateSelectPermission(body *metadata.MssqlCreateSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlCreateSelectPermissionQuery(body))
}

func (r *MetadataClient) MssqlDropSelectPermission(body *metadata.MssqlDropSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlDropSelectPermissionQuery(body))
}

func (r *MetadataClient) MssqlCreateUpdatePermission(body *metadata.MssqlCreateUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlCreateUpdatePermissionQuery(body))
}

func (r *MetadataClient) MssqlDropUpdatePermission(body *metadata.MssqlDropUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlDropUpdatePermissionQuery(body))
}

func (r *MetadataClient) MssqlCreateDeletePermission(body *metadata.MssqlCreateDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlCreateDeletePermissionQuery(body))
}

func (r *MetadataClient) MssqlDropDeletePermission(body *metadata.MssqlDropDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlDropDeletePermissionQuery(body))
}

func (r *MetadataClient) MssqlSetPermissionComment(body *metadata.MssqlSetPermissionCommentArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.MssqlSetPermissionCommentQuery(body))
}

func (r *MetadataClient) CreateCronTrigger(body *metadata.CreateCronTriggerArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.CreateCronTriggerQuery(body))
}

func (r *MetadataClient) DeleteCronTrigger(body *metadata.DeleteCronTriggerArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DeleteCronTriggerQuery(body))
}

func (r *MetadataClient) GetCronTriggers(body *metadata.GetCronTriggersArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.GetCronTriggersQuery(body))
}

func (r *MetadataClient) CreateScheduledEvent(body *metadata.CreateScheduledEventArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.CreateScheduledEventQuery(body))
}

func (r *MetadataClient) DeleteScheduledEvent(body *metadata.DeleteScheduledEventArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DeleteScheduledEventQuery(body))
}

func (r *MetadataClient) AddRemoteSchema(body *metadata.AddRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.AddRemoteSchemaQuery(body))
}

func (r *MetadataClient) UpdateRemoteSchema(body *metadata.UpdateRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.UpdateRemoteSchemaQuery(body))
}

func (r *MetadataClient) RemoveRemoteSchema(body *metadata.RemoveRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.RemoveRemoteSchemaQuery(body))
}

func (r *MetadataClient) ReloadRemoteSchema(body *metadata.ReloadRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.ReloadRemoteSchemaQuery(body))
}

func (r *MetadataClient) AddRemoteSchemaPermissions(body *metadata.AddRemoteSchemaPermissionsArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.AddRemoteSchemaPermissionsQuery(body))
}

func (r *MetadataClient) DropRemoteSchemaPermissions(body *metadata.DropRemoteSchemaPermissionsArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DropRemoteSchemaPermissionsQuery(body))
}

func (r *MetadataClient) PgCreateRemoteRelationship(body *metadata.PgCreateRemoteRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgCreateRemoteRelationshipQuery(body))
}

func (r *MetadataClient) PgUpdateRemoteRelationship(body *metadata.PgUpdateRemoteRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgUpdateRemoteRelationshipQuery(body))
}

func (r *MetadataClient) PgDeleteRemoteRelationship(body *metadata.PgDeleteRemoteRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.PgDeleteRemoteRelationshipQuery(body))
}

func (r *MetadataClient) ExportMetadata(body *metadata.ExportMetadataArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.ExportMetadataQuery(body))
}

func (r *MetadataClient) ReplaceMetadata(body *metadata.ReplaceMetadataArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.ReplaceMetadataQuery(body))
}

func (r *MetadataClient) ReloadMetadata(body *metadata.ReloadMetadataArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.ReloadMetadataQuery(body))
}

func (r *MetadataClient) ClearMetadata(body *metadata.ClearMetadataArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.ClearMetadataQuery(body))
}

func (r *MetadataClient) GetInconsistentMetadata(body *metadata.GetInconsistentMetadataArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.GetInconsistentMetadataQuery(body))
}

func (r *MetadataClient) DropInconsistentMetadata(body *metadata.DropInconsistentMetadataArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DropInconsistentMetadataQuery(body))
}

func (r *MetadataClient) CreateQueryCollection(body *metadata.CreateQueryCollectionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.CreateQueryCollectionQuery(body))
}

func (r *MetadataClient) DropQueryCollection(body *metadata.DropQueryCollectionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DropQueryCollectionQuery(body))
}

func (r *MetadataClient) AddQueryToCollection(body *metadata.AddQueryToCollectionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.AddQueryToCollectionQuery(body))
}

func (r *MetadataClient) DropQueryFromCollection(body *metadata.DropQueryFromCollectionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DropQueryFromCollectionQuery(body))
}

func (r *MetadataClient) AddCollectionToAllowlist(body *metadata.AddCollectionToAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.AddCollectionToAllowlistQuery(body))
}

func (r *MetadataClient) DropCollectionFromAllowlist(body *metadata.DropCollectionFromAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DropCollectionFromAllowlistQuery(body))
}

func (r *MetadataClient) SetCustomTypes(body *metadata.SetCustomTypesArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.SetCustomTypesQuery(body))
}

func (r *MetadataClient) CreateAction(body *metadata.CreateActionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.CreateActionQuery(body))
}

func (r *MetadataClient) DropAction(body *metadata.DropActionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DropActionQuery(body))
}

func (r *MetadataClient) UpdateAction(body *metadata.UpdateActionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.UpdateActionQuery(body))
}

func (r *MetadataClient) CreateActionPermission(body *metadata.CreateActionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.CreateActionPermissionQuery(body))
}

func (r *MetadataClient) DropActionPermission(body *metadata.DropActionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DropActionPermissionQuery(body))
}

func (r *MetadataClient) CreateRestEndpoint(body *metadata.CreateRestEndpointArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.CreateRestEndpointQuery(body))
}

func (r *MetadataClient) DropRestEndpoint(body *metadata.DropRestEndpointArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DropRestEndpointQuery(body))
}

func (r *MetadataClient) AddInheritedRole(body *metadata.AddInheritedRoleArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.AddInheritedRoleQuery(body))
}

func (r *MetadataClient) DropInheritedRole(body *metadata.DropInheritedRoleArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DropInheritedRoleQuery(body))
}

func (r *MetadataClient) SetGraphqlIntrospectionOptions(body *metadata.SetGraphqlIntrospectionOptionsArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.SetGraphqlIntrospectionOptionsQuery(body))
}

func (r *MetadataClient) AddHostToTlsAllowlist(body *metadata.AddHostToTlsAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.AddHostToTlsAllowlistQuery(body))
}

func (r *MetadataClient) DropHostFromTlsAllowlist(body *metadata.DropHostFromTlsAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.Exec(metadata.DropHostFromTlsAllowlistQuery(body))
}
