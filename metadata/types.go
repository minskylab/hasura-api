package metadata

type CustomRootFields struct {
	Insert          string `json:"insert,omitempty"`
	SelectAggregate string `json:"select_aggregate,omitempty"`
	InsertOne       string `json:"insert_one,omitempty"`
	SelectByPk      string `json:"select_by_pk,omitempty"`
	Select          string `json:"select,omitempty"`
	Delete          string `json:"delete,omitempty"`
	Update          string `json:"update,omitempty"`
	DeleteByPk      string `json:"delete_by_pk,omitempty"`
	UpdateByPk      string `json:"update_by_pk,omitempty"`
}

type TableConfiguration struct {
	Identifier        string            `json:"identifier,omitempty"` // TO REVIEW: https://hasura.io/docs/latest/graphql/core/api-reference/syntax-defs.html#table-config
	CustomRootFields  *CustomRootFields `json:"custom_root_fields,omitempty"`
	CustomName        string            `json:"custom_name,omitempty"`
	CustomColumnNames map[string]string `json:"custom_column_names,omitempty"`
}

type Using struct {
	// ForeignKeyConstraintOn string `json:"foreign_key_constraint_on"`
	ForeignKeyConstraintOn interface{} `json:"foreign_key_constraint_on,omitempty"`
}

func (u Using) GetObjRelUsingChoice() interface{} {
	return u.ForeignKeyConstraintOn
}

type ObjectRelationship struct {
	Name  string `json:"name"`
	Using Using  `json:"using"`
}

type ID struct {
	Eq string `json:"_eq"`
}

type InsertPermission struct {
	Check       IBoolExp               `json:"check"`
	Set         map[string]interface{} `json:"set,omitempty"`
	Columns     PermissionColumns      `json:"columns,omitempty"`
	BackendOnly bool                   `json:"backend_only,omitempty"`
}

type SelectPermission struct {
	Columns           PermissionColumns `json:"columns,omitempty"`
	ComputedFields    []string          `json:"computed_fields,omitempty"`
	Filter            IBoolExp          `json:"filter"`
	Limit             *int              `json:"limit,omitempty"`
	AllowAggregations bool              `json:"allow_aggregations,omitempty"`
}

type UpdatePermission struct {
	Columns PermissionColumns      `json:"columns"`
	Filter  IBoolExp               `json:"filter"`
	Check   IBoolExp               `json:"check,omitempty"`
	Set     map[string]interface{} `json:"set,omitempty"`
}

type DeletePermission struct {
	Filter IBoolExp `json:"filter"`
}

type ForeignKeyConstraintOn struct {
	Column          string     `json:"column,omitempty"`
	ExcludedColumns []string   `json:"excluded_columns,omitempty"`
	Table           ITableName `json:"table,omitempty"`
}

type UsingArray struct {
	ForeignKeyConstraintOn ForeignKeyConstraintOn `json:"foreign_key_constraint_on"`
}

// func (u UsingArray) Get

type ArrayRelationship struct {
	Name  string     `json:"name"`
	Using UsingArray `json:"using"`
}

type TableDefinition struct {
	Table               ITableName            `json:"table"`
	Configuration       *TableConfiguration   `json:"configuration,omitempty"`
	ObjectRelationships []*ObjectRelationship `json:"object_relationships,omitempty"`
	InsertPermissions   []*InsertPermission   `json:"insert_permissions,omitempty"`
	SelectPermissions   []*SelectPermission   `json:"select_permissions,omitempty"`
	UpdatePermissions   []*UpdatePermission   `json:"update_permissions,omitempty"`
	DeletePermissions   []*DeletePermission   `json:"delete_permissions,omitempty"`
	ArrayRelationships  []*ArrayRelationship  `json:"array_relationships,omitempty"`
}

type DatabaseURL struct {
	FromEnv string `json:"from_env"`
}

type PoolSettings struct {
	ConnectionLifetime int `json:"connection_lifetime"`
	Retries            int `json:"retries"`
	IdleTimeout        int `json:"idle_timeout"`
	MaxConnections     int `json:"max_connections"`
}

type ConnectionInfo struct {
	UsePreparedStatements bool         `json:"use_prepared_statements"`
	DatabaseURL           DatabaseURL  `json:"database_url"`
	IsolationLevel        string       `json:"isolation_level"`
	PoolSettings          PoolSettings `json:"pool_settings"`
}

type ConfigurationSource struct {
	ConnectionInfo ConnectionInfo `json:"connection_info"`
}

type Source struct {
	Name          string              `json:"name"`
	Kind          string              `json:"kind"`
	Tables        []*TableDefinition  `json:"tables"`
	Configuration ConfigurationSource `json:"configuration"`
}

type Definition struct {
	URL                  string `json:"url"`
	TimeoutSeconds       int    `json:"timeout_seconds"`
	ForwardClientHeaders bool   `json:"forward_client_headers"`
}

type RemoteSchemas struct {
	Name       string     `json:"name"`
	Definition Definition `json:"definition"`
}

// HAND-CRAFTED IMPLEMENTATIONS
type PermissionColumns interface {
	GetColumns() interface{}
}

type All struct{}

func (All) GetColumns() interface{} {
	return "*"
}

type PGColumns []string

func (c PGColumns) GetColumns() interface{} {
	return c
}

type IBoolExp interface {
	GetBoolExp() interface{}
}

// AndExp | OrExp | NotExp | ExistsExp | TrueExp | ColumnExp
type AndExp struct {
	And []IBoolExp `json:"$and"`
}

type OrExp struct {
	Or []IBoolExp `json:"$or"`
}

type NotExp struct {
	Not IBoolExp `json:"$not"`
}

type ExistsExp struct {
	Exists struct {
		Table ITableName `json:"_table"`
		Where IBoolExp   `json:"_where"`
	} `json:"$exists"`
}

type TrueExp struct{}

type Operator string

const (
	EqOperator  Operator = "$eq"
	NeqOperator Operator = "$ne"
	LtOperator  Operator = "$lt"
	LteOperator Operator = "$lte"
	GtOperator  Operator = "$gt"
	GteOperator Operator = "$gte"
	InOperator  Operator = "$in"
	NinOperator Operator = "$nin"

	LikeOperator     Operator = "$like"
	NLikeOperator    Operator = "$nlike"
	ILikeOperator    Operator = "$ilike"
	NILikeOperator   Operator = "$nilike"
	SimilarOperator  Operator = "$similar"
	NSimilarOperator Operator = "$nsimilar"
	RegexOperator    Operator = "$regex"
	NRegexOperator   Operator = "$nregex"
	IRegexOperator   Operator = "$iregex"
	NIRegexOperator  Operator = "$niregex"
)

type ColumnExp map[string]map[Operator]interface{}

func (a AndExp) GetBoolExp() interface{} {
	return a
}

func (a OrExp) GetBoolExp() interface{} {
	return a
}

func (n NotExp) GetBoolExp() interface{} {
	return n
}

func (e ExistsExp) GetBoolExp() interface{} {
	return e
}

func (t TrueExp) GetBoolExp() interface{} {
	return t
}

func (c ColumnExp) GetBoolExp() interface{} {
	return c
}

type (
	TableName   string
	SameTable   string
	InsertOrder string
)

type ITableName interface {
	GetTableName() interface{}
}

func (l TableName) GetTableName() interface{} {
	return l
}

type QualifiedTableName struct {
	Schema string `json:"schema,omitempty"`
	Name   string `json:"name"`
}

func (q QualifiedTableName) GetTableName() interface{} {
	return q
}

type IObjRelUsingChoice interface {
	GetObjRelUsingChoice() interface{}
}

func (s SameTable) GetObjRelUsingChoice() interface{} {
	return s
}

type RemoteTable struct {
	Table  ITableName `json:"table"`
	Column string     `json:"column"`
}

func (r RemoteTable) GetObjRelUsingChoice() interface{} {
	return r
}

const (
	BeforeParent InsertOrder = "before_parent"
	AfterParent  InsertOrder = "after_parent"
)

type ObjRelUsingManualMapping struct {
	RemoteTable    ITableName        `json:"remote_table"`
	ColumnMapping  map[string]string `json:"column_mapping"`
	InsertionOrder *InsertOrder      `json:"insertion_order,omitempty"`
}

type ObjRelUsing struct {
	ForeignKeyConstraintOn IObjRelUsingChoice        `json:"foreign_key_constraint_on,omitempty"`
	ManualConfiguration    *ObjRelUsingManualMapping `json:"manual_configuration,omitempty"`
}

type ArrRelUsingFKeyOn struct {
	Table  ITableName `json:"table"`
	Column string     `json:"column"`
}

type ArrRelUsingManualMapping struct {
	RemoteTable   ITableName        `json:"remote_table"`
	ColumnMapping map[string]string `json:"column_mapping"`
}

type ArrRelUsing struct {
	ForeignKeyConstraintOn *ArrRelUsingFKeyOn        `json:"foreign_key_constraint_on,omitempty"`
	ManualConfiguration    *ArrRelUsingManualMapping `json:"manual_configuration,omitempty"`
}
