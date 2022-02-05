# Hasura Go API

An unofficial API based on [Hasura API Reference](https://hasura.io/docs/latest/graphql/core/api-reference/index.html).

## Use Example

```go
client, err := hasura_api.NewHasuraClient()
if err != nil {
    //...
}

res, err = client.Metadata.PgTrackTable(&metadata.PgTrackTableArgs{
    Table:  metadata.TableName("users"),
    Source: "public",
})
if err != nil {
     //...
}


```

## API Implementation Status

**Priority**: Low | Medium | High

| API               | Endpoint            | Access           | Priority | Status |
| ----------------- | ------------------- | ---------------- | -------- | ------ |
| GraphQL           | /v1/graphql         | Permission rules | Medium   | 0%     |
| Relay             | /v1beta1/relay      | Permission rules | Low      | 0%     |
| Schema (> v2.0)   | /v2/query           | Admin only       | High     | 0%     |
| Metadata (> v2.0) | /v1/metadata        | Admin only       | High     | 10%    |
| Restified GQL     | /api/rest           | GQL REST Routes  | Low      | 0%     |
| Version           | /v1/version         | Public           | Low      | 0%     |
| Health            | /healthz            | Public           | Low      | 0%     |
| PG Dump           | /v1alpha1/pg_dump   | Admin only       | Low      | 0%     |
| Config            | /v1alpha1/config    | Admin only       | Medium   | 0%     |
| Explain           | /v1/graphql/explain | Admin only       | Low      | 0%     |

### Metadata API Status

| Query                          | Status | Comments |
| ------------------------------ | ------ | -------- |
| Bulk                           | 100%   |          |
| PgAddSource                    |        |          |
| PgDropSource                   |        |          |
| PgTrackTable                   | 100%   |          |
| PgUntrackTable                 | 100%   |          |
| PgSetTableCustomization        | 100%   |          |
| PgSetTableIsEnum               |        |          |
| PgTrackFunction                |        |          |
| PgUntrackFunction              |        |          |
| PgSetFunctionCustomization     |        |          |
| PgCreateFunctionPermission     |        |          |
| PgDropFunctionPermission       |        |          |
| PgCreateObjectRelationship     |        |          |
| PgCreateArrayRelationship      |        |          |
| PgDropRelationship             |        |          |
| PgRenameRelationship           |        |          |
| PgSetRelationshipComment       |        |          |
| PgAddComputedField             |        |          |
| PgDropComputedField            |        |          |
| PgCreateInsertPermission       | 100%   |          |
| PgDropInsertPermission         | 100%   |          |
| PgCreateSelectPermission       | 100%   |          |
| PgDropSelectPermission         | 100%   |          |
| PgCreateUpdatePermission       | 100%   |          |
| PgDropUpdatePermission         | 100%   |          |
| PgCreateDeletePermission       | 100%   |          |
| PgDropDeletePermission         | 100%   |          |
| PgSetPermissionComment         |        |          |
| PgCreateEventTrigger           |        |          |
| PgDeleteEventTrigger           |        |          |
| PgRedeliverEvent               |        |          |
| PgInvokeEventTrigger           |        |          |
| BigqueryTrackTable             |        |          |
| BigqueryUntrackTable           |        |          |
| BigquerySetTableCustomization  |        |          |
| MssqlAddSource                 |        |          |
| MssqlDropSource                |        |          |
| MssqlTrackTable                |        |          |
| MssqlUntrackTable              |        |          |
| MssqlCreateObjectRelationship  |        |          |
| MssqlCreateArrayRelationship   |        |          |
| MssqlDropRelationship          |        |          |
| MssqlRenameRelationship        |        |          |
| MssqlSetRelationshipComment    |        |          |
| MssqlSetTableCustomization     |        |          |
| MssqlCreateInsertPermission    |        |          |
| MssqlDropInsertPermission      |        |          |
| MssqlCreateSelectPermission    |        |          |
| MssqlDropSelectPermission      |        |          |
| MssqlCreateUpdatePermission    |        |          |
| MssqlDropUpdatePermission      |        |          |
| MssqlCreateDeletePermission    |        |          |
| MssqlDropDeletePermission      |        |          |
| MssqlSetPermissionComment      |        |          |
| CreateCronTrigger              |        |          |
| DeleteCronTrigger              |        |          |
| GetCronTriggers                |        |          |
| CreateScheduledEvent           |        |          |
| DeleteScheduledEvent           |        |          |
| AddRemoteSchema                |        |          |
| UpdateRemoteSchema             |        |          |
| RemoveRemoteSchema             |        |          |
| ReloadRemoteSchema             |        |          |
| AddRemoteSchemaPermissions     |        |          |
| DropRemoteSchemaPermissions    |        |          |
| PgCreateRemoteRelationship     |        |          |
| PgUpdateRemoteRelationship     |        |          |
| PgDeleteRemoteRelationship     |        |          |
| ExportMetadata                 |        |          |
| ReplaceMetadata                |        |          |
| ReloadMetadata                 |        |          |
| ClearMetadata                  |        |          |
| GetInconsistentMetadata        |        |          |
| DropInconsistentMetadata       |        |          |
| CreateQueryCollection          |        |          |
| DropQueryCollection            |        |          |
| AddQueryToCollection           |        |          |
| DropQueryFromCollection        |        |          |
| AddCollectionToAllowlist       |        |          |
| DropCollectionFromAllowlist    |        |          |
| SetCustomTypes                 |        |          |
| CreateAction                   |        |          |
| DropAction                     |        |          |
| UpdateAction                   |        |          |
| CreateActionPermission         |        |          |
| DropActionPermission           |        |          |
| CreateRestEndpoint             |        |          |
| DropRestEndpoint               |        |          |
| AddInheritedRole               |        |          |
| DropInheritedRole              |        |          |
| SetGraphqlIntrospectionOptions |        |          |
| AddHostToTlsAllowlist          |        |          |
| DropHostFromTlsAllowlist       |        |          |
