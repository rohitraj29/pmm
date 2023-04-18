// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type kubernetesClusterTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *kubernetesClusterTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("kubernetes_clusters").
func (v *kubernetesClusterTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *kubernetesClusterTableType) Columns() []string {
	return []string{
		"id",
		"kubernetes_cluster_name",
		"kube_config",
		"ready",
		"pxc",
		"proxysql",
		"haproxy",
		"mongod",
		"postgresql",
		"pgbouncer",
		"pgbackrest",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *kubernetesClusterTableType) NewStruct() reform.Struct {
	return new(KubernetesCluster)
}

// NewRecord makes a new record for that table.
func (v *kubernetesClusterTableType) NewRecord() reform.Record {
	return new(KubernetesCluster)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *kubernetesClusterTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// KubernetesClusterTable represents kubernetes_clusters view or table in SQL database.
var KubernetesClusterTable = &kubernetesClusterTableType{
	s: parse.StructInfo{
		Type:    "KubernetesCluster",
		SQLName: "kubernetes_clusters",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "string", Column: "id"},
			{Name: "KubernetesClusterName", Type: "string", Column: "kubernetes_cluster_name"},
			{Name: "KubeConfig", Type: "string", Column: "kube_config"},
			{Name: "IsReady", Type: "bool", Column: "ready"},
			{Name: "PXC", Type: "*Component", Column: "pxc"},
			{Name: "ProxySQL", Type: "*Component", Column: "proxysql"},
			{Name: "HAProxy", Type: "*Component", Column: "haproxy"},
			{Name: "Mongod", Type: "*Component", Column: "mongod"},
			{Name: "Postgresql", Type: "*Component", Column: "postgresql"},
			{Name: "Pgbouncer", Type: "*Component", Column: "pgbouncer"},
			{Name: "Pgbackrest", Type: "*Component", Column: "pgbackrest"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"},
		},
		PKFieldIndex: 0,
	},
	z: new(KubernetesCluster).Values(),
}

// String returns a string representation of this struct or record.
func (s KubernetesCluster) String() string {
	res := make([]string, 13)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "KubernetesClusterName: " + reform.Inspect(s.KubernetesClusterName, true)
	res[2] = "KubeConfig: " + reform.Inspect(s.KubeConfig, true)
	res[3] = "IsReady: " + reform.Inspect(s.IsReady, true)
	res[4] = "PXC: " + reform.Inspect(s.PXC, true)
	res[5] = "ProxySQL: " + reform.Inspect(s.ProxySQL, true)
	res[6] = "HAProxy: " + reform.Inspect(s.HAProxy, true)
	res[7] = "Mongod: " + reform.Inspect(s.Mongod, true)
	res[8] = "Postgresql: " + reform.Inspect(s.Postgresql, true)
	res[9] = "Pgbouncer: " + reform.Inspect(s.Pgbouncer, true)
	res[10] = "Pgbackrest: " + reform.Inspect(s.Pgbackrest, true)
	res[11] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[12] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *KubernetesCluster) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.KubernetesClusterName,
		s.KubeConfig,
		s.IsReady,
		s.PXC,
		s.ProxySQL,
		s.HAProxy,
		s.Mongod,
		s.Postgresql,
		s.Pgbouncer,
		s.Pgbackrest,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *KubernetesCluster) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.KubernetesClusterName,
		&s.KubeConfig,
		&s.IsReady,
		&s.PXC,
		&s.ProxySQL,
		&s.HAProxy,
		&s.Mongod,
		&s.Postgresql,
		&s.Pgbouncer,
		&s.Pgbackrest,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *KubernetesCluster) View() reform.View {
	return KubernetesClusterTable
}

// Table returns Table object for that record.
func (s *KubernetesCluster) Table() reform.Table {
	return KubernetesClusterTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *KubernetesCluster) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *KubernetesCluster) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *KubernetesCluster) HasPK() bool {
	return s.ID != KubernetesClusterTable.z[KubernetesClusterTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *KubernetesCluster) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = KubernetesClusterTable
	_ reform.Struct = (*KubernetesCluster)(nil)
	_ reform.Table  = KubernetesClusterTable
	_ reform.Record = (*KubernetesCluster)(nil)
	_ fmt.Stringer  = (*KubernetesCluster)(nil)
)

func init() {
	parse.AssertUpToDate(&KubernetesClusterTable.s, new(KubernetesCluster))
}
