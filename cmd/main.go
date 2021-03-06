package main

import (
	"fmt"

	hasura_api "github.com/minskylab/hasura-api"
	"github.com/minskylab/hasura-api/metadata"
)

func main() {
	client, err := hasura_api.NewHasuraClient()
	if err != nil {
		panic(err)
	}

	query := metadata.PgTrackTableQuery(&metadata.PgTrackTableArgs{
		Table: metadata.QualifiedTableName{
			Name:   "users",
			Schema: "public",
		},
	})

	var res metadata.MetadataResponse

	res, err = client.Metadata.Exec(query)
	if err != nil {
		panic(err)
	}

	switch res := res.(type) {
	case metadata.ObjectResponse:
		fmt.Println(res)
	case metadata.BadRequestResponse,
		metadata.UnauthorizedResponse,
		metadata.InternalServerErrorResponse:
		panic(res)
	}

	res, err = client.Metadata.PgTrackTable(&metadata.PgTrackTableArgs{
		Table:  metadata.TableName("users"),
		Source: "public",
	})
	if err != nil {
		panic(err)
	}

	switch res := res.(type) {
	case metadata.ObjectResponse:
		fmt.Println(res)
	case metadata.BadRequestResponse,
		metadata.UnauthorizedResponse,
		metadata.InternalServerErrorResponse:
		panic(res)
	}
	// source := metadata.SourceName("xxx")

	// args1 := metadata.PgTrackTableArgs{
	// 	Table:  metadata.TableName("xxx"),
	// 	Source: &source,
	// }

	// args2 := metadata.PgTrackTableArgs{
	// 	Table: metadata.QualifiedTableName{
	// 		Name:   "yyy",
	// 		Schema: "yyy",
	// 	},
	// 	// Source: "yyy",
	// }

	// args3 := metadata.PgCreateInsertPermissionArgs{
	// 	Table: metadata.TableName("table_1"),
	// 	Permission: metadata.InsertPermission{
	// 		Check: metadata.ColumnExp{
	// 			"organization_id": map[metadata.Operator]interface{}{
	// 				"_eq": "X-Hasura-User-Id",
	// 			},
	// 		},
	// 	},
	// }

	//
	// switch table := args1.Table.GetTableName().(type) {
	// case metadata.TableName:
	// 	print(table)
	// case metadata.QualifiedTableName:
	// 	print(table)

	// }

	// d, _ := json.Marshal(args1)
	// print(string(d))

	// d, _ = json.Marshal(args2)
	// print(string(d))

	// d, _ = json.Marshal(args3)
	// print(string(d))
}
