package main

import (
	"encoding/json"

	"github.com/minskylab/hasura-api/metadata"
)

func main() {
	source := metadata.SourceName("xxx")

	args1 := metadata.PgTrackTableArgs{
		Table:  metadata.TableName("xxx"),
		Source: &source,
	}

	args2 := metadata.PgTrackTableArgs{
		Table: metadata.QualifiedTableName{
			Name:   "yyy",
			Schema: "yyy",
		},
		// Source: "yyy",
	}

	args3 := metadata.PgCreateInsertPermissionArgs{
		Table: metadata.TableName("table_1"),
		Permission: metadata.InsertPermission{
			Check: metadata.ColumnExp{
				"organization_id": map[metadata.Operator]interface{}{
					"_eq": "X-Hasura-User-Id",
				},
			},
		},
	}

	//
	// switch table := args1.Table.GetTableName().(type) {
	// case metadata.TableName:
	// 	print(table)
	// case metadata.QualifiedTableName:
	// 	print(table)

	// }

	d, _ := json.Marshal(args1)
	print(string(d))

	d, _ = json.Marshal(args2)
	print(string(d))

	d, _ = json.Marshal(args3)
	print(string(d))
}
