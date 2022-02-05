package hasura_api

import "github.com/minskylab/hasura-api/metadata"

func (r *HasuraClient) BulkQuery(args metadata.BulkArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.BulkType,
		Args: args,
	}
}

func (r *HasuraClient) ExecBulk(body []metadata.MetadataQuery) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.BulkQuery(metadata.BulkArgs(body)))
}

func (r *HasuraClient) PgAddSourceQuery(args *metadata.PgAddSourceArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgAddSource,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgAddSource(body *metadata.PgAddSourceArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgAddSourceQuery(body))
}

func (r *HasuraClient) PgDropSourceQuery(args *metadata.PgDropSourceArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgDropSource,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgDropSource(body *metadata.PgDropSourceArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgDropSourceQuery(body))
}

func (r *HasuraClient) PgTrackTableQuery(args *metadata.PgTrackTableArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgTrackTable,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgTrackTable(body *metadata.PgTrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgTrackTableQuery(body))
}

func (r *HasuraClient) PgUntrackTableQuery(args *metadata.PgUntrackTableArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgUntrackTable,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgUntrackTable(body *metadata.PgUntrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgUntrackTableQuery(body))
}

func (r *HasuraClient) PgSetTableCustomizationQuery(args *metadata.PgSetTableCustomizationArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgSetTableCustomization,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgSetTableCustomization(body *metadata.PgSetTableCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgSetTableCustomizationQuery(body))
}

func (r *HasuraClient) PgSetTableIsEnumQuery(args *metadata.PgSetTableIsEnumArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgSetTableIsEnum,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgSetTableIsEnum(body *metadata.PgSetTableIsEnumArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgSetTableIsEnumQuery(body))
}

func (r *HasuraClient) PgTrackFunctionQuery(args *metadata.PgTrackFunctionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgTrackFunction,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgTrackFunction(body *metadata.PgTrackFunctionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgTrackFunctionQuery(body))
}

func (r *HasuraClient) PgUntrackFunctionQuery(args *metadata.PgUntrackFunctionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgUntrackFunction,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgUntrackFunction(body *metadata.PgUntrackFunctionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgUntrackFunctionQuery(body))
}

func (r *HasuraClient) PgSetFunctionCustomizationQuery(args *metadata.PgSetFunctionCustomizationArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgSetFunctionCustomization,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgSetFunctionCustomization(body *metadata.PgSetFunctionCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgSetFunctionCustomizationQuery(body))
}

func (r *HasuraClient) PgCreateFunctionPermissionQuery(args *metadata.PgCreateFunctionPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgCreateFunctionPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgCreateFunctionPermission(body *metadata.PgCreateFunctionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgCreateFunctionPermissionQuery(body))
}

func (r *HasuraClient) PgDropFunctionPermissionQuery(args *metadata.PgDropFunctionPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgDropFunctionPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgDropFunctionPermission(body *metadata.PgDropFunctionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgDropFunctionPermissionQuery(body))
}

func (r *HasuraClient) PgCreateObjectRelationshipQuery(args *metadata.PgCreateObjectRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgCreateObjectRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgCreateObjectRelationship(body *metadata.PgCreateObjectRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgCreateObjectRelationshipQuery(body))
}

func (r *HasuraClient) PgCreateArrayRelationshipQuery(args *metadata.PgCreateArrayRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgCreateArrayRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgCreateArrayRelationship(body *metadata.PgCreateArrayRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgCreateArrayRelationshipQuery(body))
}

func (r *HasuraClient) PgDropRelationshipQuery(args *metadata.PgDropRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgDropRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgDropRelationship(body *metadata.PgDropRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgDropRelationshipQuery(body))
}

func (r *HasuraClient) PgRenameRelationshipQuery(args *metadata.PgRenameRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgRenameRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgRenameRelationship(body *metadata.PgRenameRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgRenameRelationshipQuery(body))
}

func (r *HasuraClient) PgSetRelationshipCommentQuery(args *metadata.PgSetRelationshipCommentArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgSetRelationshipComment,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgSetRelationshipComment(body *metadata.PgSetRelationshipCommentArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgSetRelationshipCommentQuery(body))
}

func (r *HasuraClient) PgAddComputedFieldQuery(args *metadata.PgAddComputedFieldArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgAddComputedField,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgAddComputedField(body *metadata.PgAddComputedFieldArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgAddComputedFieldQuery(body))
}

func (r *HasuraClient) PgDropComputedFieldQuery(args *metadata.PgDropComputedFieldArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgDropComputedField,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgDropComputedField(body *metadata.PgDropComputedFieldArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgDropComputedFieldQuery(body))
}

func (r *HasuraClient) PgCreateInsertPermissionQuery(args *metadata.PgCreateInsertPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgCreateInsertPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgCreateInsertPermission(body *metadata.PgCreateInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgCreateInsertPermissionQuery(body))
}

func (r *HasuraClient) PgDropInsertPermissionQuery(args *metadata.PgDropInsertPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgDropInsertPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgDropInsertPermission(body *metadata.PgDropInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgDropInsertPermissionQuery(body))
}

func (r *HasuraClient) PgCreateSelectPermissionQuery(args *metadata.PgCreateSelectPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgCreateSelectPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgCreateSelectPermission(body *metadata.PgCreateSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgCreateSelectPermissionQuery(body))
}

func (r *HasuraClient) PgDropSelectPermissionQuery(args *metadata.PgDropSelectPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgDropSelectPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgDropSelectPermission(body *metadata.PgDropSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgDropSelectPermissionQuery(body))
}

func (r *HasuraClient) PgCreateUpdatePermissionQuery(args *metadata.PgCreateUpdatePermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgCreateUpdatePermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgCreateUpdatePermission(body *metadata.PgCreateUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgCreateUpdatePermissionQuery(body))
}

func (r *HasuraClient) PgDropUpdatePermissionQuery(args *metadata.PgDropUpdatePermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgDropUpdatePermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgDropUpdatePermission(body *metadata.PgDropUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgDropUpdatePermissionQuery(body))
}

func (r *HasuraClient) PgCreateDeletePermissionQuery(args *metadata.PgCreateDeletePermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgCreateDeletePermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgCreateDeletePermission(body *metadata.PgCreateDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgCreateDeletePermissionQuery(body))
}

func (r *HasuraClient) PgDropDeletePermissionQuery(args *metadata.PgDropDeletePermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgDropDeletePermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgDropDeletePermission(body *metadata.PgDropDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgDropDeletePermissionQuery(body))
}

func (r *HasuraClient) PgSetPermissionCommentQuery(args *metadata.PgSetPermissionCommentArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgSetPermissionComment,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgSetPermissionComment(body *metadata.PgSetPermissionCommentArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgSetPermissionCommentQuery(body))
}

func (r *HasuraClient) PgCreateEventTriggerQuery(args *metadata.PgCreateEventTriggerArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgCreateEventTrigger,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgCreateEventTrigger(body *metadata.PgCreateEventTriggerArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgCreateEventTriggerQuery(body))
}

func (r *HasuraClient) PgDeleteEventTriggerQuery(args *metadata.PgDeleteEventTriggerArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgDeleteEventTrigger,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgDeleteEventTrigger(body *metadata.PgDeleteEventTriggerArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgDeleteEventTriggerQuery(body))
}

func (r *HasuraClient) PgRedeliverEventQuery(args *metadata.PgRedeliverEventArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgRedeliverEvent,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgRedeliverEvent(body *metadata.PgRedeliverEventArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgRedeliverEventQuery(body))
}

func (r *HasuraClient) PgInvokeEventTriggerQuery(args *metadata.PgInvokeEventTriggerArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgInvokeEventTrigger,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgInvokeEventTrigger(body *metadata.PgInvokeEventTriggerArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgInvokeEventTriggerQuery(body))
}

func (r *HasuraClient) BigqueryTrackTableQuery(args *metadata.BigqueryTrackTableArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.BigqueryTrackTable,
		Args: args,
	}
}

func (r *HasuraClient) ExecBigqueryTrackTable(body *metadata.BigqueryTrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.BigqueryTrackTableQuery(body))
}

func (r *HasuraClient) BigqueryUntrackTableQuery(args *metadata.BigqueryUntrackTableArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.BigqueryUntrackTable,
		Args: args,
	}
}

func (r *HasuraClient) ExecBigqueryUntrackTable(body *metadata.BigqueryUntrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.BigqueryUntrackTableQuery(body))
}

func (r *HasuraClient) BigquerySetTableCustomizationQuery(args *metadata.BigquerySetTableCustomizationArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.BigquerySetTableCustomization,
		Args: args,
	}
}

func (r *HasuraClient) ExecBigquerySetTableCustomization(body *metadata.BigquerySetTableCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.BigquerySetTableCustomizationQuery(body))
}

func (r *HasuraClient) MssqlAddSourceQuery(args *metadata.MssqlAddSourceArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlAddSource,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlAddSource(body *metadata.MssqlAddSourceArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlAddSourceQuery(body))
}

func (r *HasuraClient) MssqlDropSourceQuery(args *metadata.MssqlDropSourceArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlDropSource,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlDropSource(body *metadata.MssqlDropSourceArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlDropSourceQuery(body))
}

func (r *HasuraClient) MssqlTrackTableQuery(args *metadata.MssqlTrackTableArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlTrackTable,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlTrackTable(body *metadata.MssqlTrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlTrackTableQuery(body))
}

func (r *HasuraClient) MssqlUntrackTableQuery(args *metadata.MssqlUntrackTableArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlUntrackTable,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlUntrackTable(body *metadata.MssqlUntrackTableArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlUntrackTableQuery(body))
}

func (r *HasuraClient) MssqlCreateObjectRelationshipQuery(args *metadata.MssqlCreateObjectRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlCreateObjectRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlCreateObjectRelationship(body *metadata.MssqlCreateObjectRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlCreateObjectRelationshipQuery(body))
}

func (r *HasuraClient) MssqlCreateArrayRelationshipQuery(args *metadata.MssqlCreateArrayRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlCreateArrayRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlCreateArrayRelationship(body *metadata.MssqlCreateArrayRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlCreateArrayRelationshipQuery(body))
}

func (r *HasuraClient) MssqlDropRelationshipQuery(args *metadata.MssqlDropRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlDropRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlDropRelationship(body *metadata.MssqlDropRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlDropRelationshipQuery(body))
}

func (r *HasuraClient) MssqlRenameRelationshipQuery(args *metadata.MssqlRenameRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlRenameRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlRenameRelationship(body *metadata.MssqlRenameRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlRenameRelationshipQuery(body))
}

func (r *HasuraClient) MssqlSetRelationshipCommentQuery(args *metadata.MssqlSetRelationshipCommentArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlSetRelationshipComment,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlSetRelationshipComment(body *metadata.MssqlSetRelationshipCommentArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlSetRelationshipCommentQuery(body))
}

func (r *HasuraClient) MssqlSetTableCustomizationQuery(args *metadata.MssqlSetTableCustomizationArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlSetTableCustomization,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlSetTableCustomization(body *metadata.MssqlSetTableCustomizationArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlSetTableCustomizationQuery(body))
}

func (r *HasuraClient) MssqlCreateInsertPermissionQuery(args *metadata.MssqlCreateInsertPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlCreateInsertPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlCreateInsertPermission(body *metadata.MssqlCreateInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlCreateInsertPermissionQuery(body))
}

func (r *HasuraClient) MssqlDropInsertPermissionQuery(args *metadata.MssqlDropInsertPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlDropInsertPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlDropInsertPermission(body *metadata.MssqlDropInsertPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlDropInsertPermissionQuery(body))
}

func (r *HasuraClient) MssqlCreateSelectPermissionQuery(args *metadata.MssqlCreateSelectPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlCreateSelectPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlCreateSelectPermission(body *metadata.MssqlCreateSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlCreateSelectPermissionQuery(body))
}

func (r *HasuraClient) MssqlDropSelectPermissionQuery(args *metadata.MssqlDropSelectPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlDropSelectPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlDropSelectPermission(body *metadata.MssqlDropSelectPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlDropSelectPermissionQuery(body))
}

func (r *HasuraClient) MssqlCreateUpdatePermissionQuery(args *metadata.MssqlCreateUpdatePermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlCreateUpdatePermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlCreateUpdatePermission(body *metadata.MssqlCreateUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlCreateUpdatePermissionQuery(body))
}

func (r *HasuraClient) MssqlDropUpdatePermissionQuery(args *metadata.MssqlDropUpdatePermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlDropUpdatePermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlDropUpdatePermission(body *metadata.MssqlDropUpdatePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlDropUpdatePermissionQuery(body))
}

func (r *HasuraClient) MssqlCreateDeletePermissionQuery(args *metadata.MssqlCreateDeletePermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlCreateDeletePermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlCreateDeletePermission(body *metadata.MssqlCreateDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlCreateDeletePermissionQuery(body))
}

func (r *HasuraClient) MssqlDropDeletePermissionQuery(args *metadata.MssqlDropDeletePermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlDropDeletePermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlDropDeletePermission(body *metadata.MssqlDropDeletePermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlDropDeletePermissionQuery(body))
}

func (r *HasuraClient) MssqlSetPermissionCommentQuery(args *metadata.MssqlSetPermissionCommentArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.MssqlSetPermissionComment,
		Args: args,
	}
}

func (r *HasuraClient) ExecMssqlSetPermissionComment(body *metadata.MssqlSetPermissionCommentArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.MssqlSetPermissionCommentQuery(body))
}

func (r *HasuraClient) CreateCronTriggerQuery(args *metadata.CreateCronTriggerArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.CreateCronTrigger,
		Args: args,
	}
}

func (r *HasuraClient) ExecCreateCronTrigger(body *metadata.CreateCronTriggerArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.CreateCronTriggerQuery(body))
}

func (r *HasuraClient) DeleteCronTriggerQuery(args *metadata.DeleteCronTriggerArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DeleteCronTrigger,
		Args: args,
	}
}

func (r *HasuraClient) ExecDeleteCronTrigger(body *metadata.DeleteCronTriggerArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DeleteCronTriggerQuery(body))
}

func (r *HasuraClient) GetCronTriggersQuery(args *metadata.GetCronTriggersArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.GetCronTriggers,
		Args: args,
	}
}

func (r *HasuraClient) ExecGetCronTriggers(body *metadata.GetCronTriggersArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.GetCronTriggersQuery(body))
}

func (r *HasuraClient) CreateScheduledEventQuery(args *metadata.CreateScheduledEventArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.CreateScheduledEvent,
		Args: args,
	}
}

func (r *HasuraClient) ExecCreateScheduledEvent(body *metadata.CreateScheduledEventArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.CreateScheduledEventQuery(body))
}

func (r *HasuraClient) DeleteScheduledEventQuery(args *metadata.DeleteScheduledEventArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DeleteScheduledEvent,
		Args: args,
	}
}

func (r *HasuraClient) ExecDeleteScheduledEvent(body *metadata.DeleteScheduledEventArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DeleteScheduledEventQuery(body))
}

func (r *HasuraClient) AddRemoteSchemaQuery(args *metadata.AddRemoteSchemaArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.AddRemoteSchema,
		Args: args,
	}
}

func (r *HasuraClient) ExecAddRemoteSchema(body *metadata.AddRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.AddRemoteSchemaQuery(body))
}

func (r *HasuraClient) UpdateRemoteSchemaQuery(args *metadata.UpdateRemoteSchemaArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.UpdateRemoteSchema,
		Args: args,
	}
}

func (r *HasuraClient) ExecUpdateRemoteSchema(body *metadata.UpdateRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.UpdateRemoteSchemaQuery(body))
}

func (r *HasuraClient) RemoveRemoteSchemaQuery(args *metadata.RemoveRemoteSchemaArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.RemoveRemoteSchema,
		Args: args,
	}
}

func (r *HasuraClient) ExecRemoveRemoteSchema(body *metadata.RemoveRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.RemoveRemoteSchemaQuery(body))
}

func (r *HasuraClient) ReloadRemoteSchemaQuery(args *metadata.ReloadRemoteSchemaArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.ReloadRemoteSchema,
		Args: args,
	}
}

func (r *HasuraClient) ExecReloadRemoteSchema(body *metadata.ReloadRemoteSchemaArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.ReloadRemoteSchemaQuery(body))
}

func (r *HasuraClient) AddRemoteSchemaPermissionsQuery(args *metadata.AddRemoteSchemaPermissionsArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.AddRemoteSchemaPermissions,
		Args: args,
	}
}

func (r *HasuraClient) ExecAddRemoteSchemaPermissions(body *metadata.AddRemoteSchemaPermissionsArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.AddRemoteSchemaPermissionsQuery(body))
}

func (r *HasuraClient) DropRemoteSchemaPermissionsQuery(args *metadata.DropRemoteSchemaPermissionsArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DropRemoteSchemaPermissions,
		Args: args,
	}
}

func (r *HasuraClient) ExecDropRemoteSchemaPermissions(body *metadata.DropRemoteSchemaPermissionsArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DropRemoteSchemaPermissionsQuery(body))
}

func (r *HasuraClient) PgCreateRemoteRelationshipQuery(args *metadata.PgCreateRemoteRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgCreateRemoteRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgCreateRemoteRelationship(body *metadata.PgCreateRemoteRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgCreateRemoteRelationshipQuery(body))
}

func (r *HasuraClient) PgUpdateRemoteRelationshipQuery(args *metadata.PgUpdateRemoteRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgUpdateRemoteRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgUpdateRemoteRelationship(body *metadata.PgUpdateRemoteRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgUpdateRemoteRelationshipQuery(body))
}

func (r *HasuraClient) PgDeleteRemoteRelationshipQuery(args *metadata.PgDeleteRemoteRelationshipArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.PgDeleteRemoteRelationship,
		Args: args,
	}
}

func (r *HasuraClient) ExecPgDeleteRemoteRelationship(body *metadata.PgDeleteRemoteRelationshipArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.PgDeleteRemoteRelationshipQuery(body))
}

func (r *HasuraClient) ExportMetadataQuery(args *metadata.ExportMetadataArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.ExportMetadata,
		Args: args,
	}
}

func (r *HasuraClient) ExecExportMetadata(body *metadata.ExportMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.ExportMetadataQuery(body))
}

func (r *HasuraClient) ReplaceMetadataQuery(args *metadata.ReplaceMetadataArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.ReplaceMetadata,
		Args: args,
	}
}

func (r *HasuraClient) ExecReplaceMetadata(body *metadata.ReplaceMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.ReplaceMetadataQuery(body))
}

func (r *HasuraClient) ReloadMetadataQuery(args *metadata.ReloadMetadataArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.ReloadMetadata,
		Args: args,
	}
}

func (r *HasuraClient) ExecReloadMetadata(body *metadata.ReloadMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.ReloadMetadataQuery(body))
}

func (r *HasuraClient) ClearMetadataQuery(args *metadata.ClearMetadataArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.ClearMetadata,
		Args: args,
	}
}

func (r *HasuraClient) ExecClearMetadata(body *metadata.ClearMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.ClearMetadataQuery(body))
}

func (r *HasuraClient) GetInconsistentMetadataQuery(args *metadata.GetInconsistentMetadataArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.GetInconsistentMetadata,
		Args: args,
	}
}

func (r *HasuraClient) ExecGetInconsistentMetadata(body *metadata.GetInconsistentMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.GetInconsistentMetadataQuery(body))
}

func (r *HasuraClient) DropInconsistentMetadataQuery(args *metadata.DropInconsistentMetadataArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DropInconsistentMetadata,
		Args: args,
	}
}

func (r *HasuraClient) ExecDropInconsistentMetadata(body *metadata.DropInconsistentMetadataArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DropInconsistentMetadataQuery(body))
}

func (r *HasuraClient) CreateQueryCollectionQuery(args *metadata.CreateQueryCollectionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.CreateQueryCollection,
		Args: args,
	}
}

func (r *HasuraClient) ExecCreateQueryCollection(body *metadata.CreateQueryCollectionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.CreateQueryCollectionQuery(body))
}

func (r *HasuraClient) DropQueryCollectionQuery(args *metadata.DropQueryCollectionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DropQueryCollection,
		Args: args,
	}
}

func (r *HasuraClient) ExecDropQueryCollection(body *metadata.DropQueryCollectionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DropQueryCollectionQuery(body))
}

func (r *HasuraClient) AddQueryToCollectionQuery(args *metadata.AddQueryToCollectionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.AddQueryToCollection,
		Args: args,
	}
}

func (r *HasuraClient) ExecAddQueryToCollection(body *metadata.AddQueryToCollectionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.AddQueryToCollectionQuery(body))
}

func (r *HasuraClient) DropQueryFromCollectionQuery(args *metadata.DropQueryFromCollectionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DropQueryFromCollection,
		Args: args,
	}
}

func (r *HasuraClient) ExecDropQueryFromCollection(body *metadata.DropQueryFromCollectionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DropQueryFromCollectionQuery(body))
}

func (r *HasuraClient) AddCollectionToAllowlistQuery(args *metadata.AddCollectionToAllowlistArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.AddCollectionToAllowlist,
		Args: args,
	}
}

func (r *HasuraClient) ExecAddCollectionToAllowlist(body *metadata.AddCollectionToAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.AddCollectionToAllowlistQuery(body))
}

func (r *HasuraClient) DropCollectionFromAllowlistQuery(args *metadata.DropCollectionFromAllowlistArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DropCollectionFromAllowlist,
		Args: args,
	}
}

func (r *HasuraClient) ExecDropCollectionFromAllowlist(body *metadata.DropCollectionFromAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DropCollectionFromAllowlistQuery(body))
}

func (r *HasuraClient) SetCustomTypesQuery(args *metadata.SetCustomTypesArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.SetCustomTypes,
		Args: args,
	}
}

func (r *HasuraClient) ExecSetCustomTypes(body *metadata.SetCustomTypesArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.SetCustomTypesQuery(body))
}

func (r *HasuraClient) CreateActionQuery(args *metadata.CreateActionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.CreateAction,
		Args: args,
	}
}

func (r *HasuraClient) ExecCreateAction(body *metadata.CreateActionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.CreateActionQuery(body))
}

func (r *HasuraClient) DropActionQuery(args *metadata.DropActionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DropAction,
		Args: args,
	}
}

func (r *HasuraClient) ExecDropAction(body *metadata.DropActionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DropActionQuery(body))
}

func (r *HasuraClient) UpdateActionQuery(args *metadata.UpdateActionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.UpdateAction,
		Args: args,
	}
}

func (r *HasuraClient) ExecUpdateAction(body *metadata.UpdateActionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.UpdateActionQuery(body))
}

func (r *HasuraClient) CreateActionPermissionQuery(args *metadata.CreateActionPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.CreateActionPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecCreateActionPermission(body *metadata.CreateActionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.CreateActionPermissionQuery(body))
}

func (r *HasuraClient) DropActionPermissionQuery(args *metadata.DropActionPermissionArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DropActionPermission,
		Args: args,
	}
}

func (r *HasuraClient) ExecDropActionPermission(body *metadata.DropActionPermissionArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DropActionPermissionQuery(body))
}

func (r *HasuraClient) CreateRestEndpointQuery(args *metadata.CreateRestEndpointArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.CreateRestEndpoint,
		Args: args,
	}
}

func (r *HasuraClient) ExecCreateRestEndpoint(body *metadata.CreateRestEndpointArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.CreateRestEndpointQuery(body))
}

func (r *HasuraClient) DropRestEndpointQuery(args *metadata.DropRestEndpointArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DropRestEndpoint,
		Args: args,
	}
}

func (r *HasuraClient) ExecDropRestEndpoint(body *metadata.DropRestEndpointArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DropRestEndpointQuery(body))
}

func (r *HasuraClient) AddInheritedRoleQuery(args *metadata.AddInheritedRoleArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.AddInheritedRole,
		Args: args,
	}
}

func (r *HasuraClient) ExecAddInheritedRole(body *metadata.AddInheritedRoleArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.AddInheritedRoleQuery(body))
}

func (r *HasuraClient) DropInheritedRoleQuery(args *metadata.DropInheritedRoleArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DropInheritedRole,
		Args: args,
	}
}

func (r *HasuraClient) ExecDropInheritedRole(body *metadata.DropInheritedRoleArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DropInheritedRoleQuery(body))
}

func (r *HasuraClient) SetGraphqlIntrospectionOptionsQuery(args *metadata.SetGraphqlIntrospectionOptionsArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.SetGraphqlIntrospectionOptions,
		Args: args,
	}
}

func (r *HasuraClient) ExecSetGraphqlIntrospectionOptions(body *metadata.SetGraphqlIntrospectionOptionsArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.SetGraphqlIntrospectionOptionsQuery(body))
}

func (r *HasuraClient) AddHostToTlsAllowlistQuery(args *metadata.AddHostToTlsAllowlistArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.AddHostToTlsAllowlist,
		Args: args,
	}
}

func (r *HasuraClient) ExecAddHostToTlsAllowlist(body *metadata.AddHostToTlsAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.AddHostToTlsAllowlistQuery(body))
}

func (r *HasuraClient) DropHostFromTlsAllowlistQuery(args *metadata.DropHostFromTlsAllowlistArgs) metadata.MetadataQuery {
	return metadata.MetadataQuery{
		Type: metadata.DropHostFromTlsAllowlist,
		Args: args,
	}
}

func (r *HasuraClient) ExecDropHostFromTlsAllowlist(body *metadata.DropHostFromTlsAllowlistArgs) (metadata.MetadataResponse, error) {
	return r.genericHasuraQuery(r.DropHostFromTlsAllowlistQuery(body))
}
