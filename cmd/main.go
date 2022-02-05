package main

import (
	"encoding/json"

	"github.com/minskylab/hasura-api/metadata"
)

func main() {
	args1 := metadata.PgTrackTableArgs{
		Table:  metadata.TableName("xxx"),
		Source: "xxx",
	}

	args2 := metadata.PgTrackTableArgs{
		Table: metadata.QualifiedTableName{
			Name:   "yyy",
			Schema: "yyy",
		},
		Source: "yyy",
	}

	//
	switch table := args1.Table.GetTableName().(type) {
	case metadata.TableName:
		print(table)
	case metadata.QualifiedTableName:
		print(table)

	}

	d, _ := json.Marshal(args1)
	print(string(d))

	d, _ = json.Marshal(args2)
	print(string(d))
}
