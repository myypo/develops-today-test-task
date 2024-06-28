//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Targets = newTargetsTable("public", "targets", "")

type targetsTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnString
	Name      postgres.ColumnString
	Country   postgres.ColumnString
	Notes     postgres.ColumnString
	Status    postgres.ColumnString
	MissionID postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TargetsTable struct {
	targetsTable

	EXCLUDED targetsTable
}

// AS creates new TargetsTable with assigned alias
func (a TargetsTable) AS(alias string) *TargetsTable {
	return newTargetsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TargetsTable with assigned schema name
func (a TargetsTable) FromSchema(schemaName string) *TargetsTable {
	return newTargetsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TargetsTable with assigned table prefix
func (a TargetsTable) WithPrefix(prefix string) *TargetsTable {
	return newTargetsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TargetsTable with assigned table suffix
func (a TargetsTable) WithSuffix(suffix string) *TargetsTable {
	return newTargetsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTargetsTable(schemaName, tableName, alias string) *TargetsTable {
	return &TargetsTable{
		targetsTable: newTargetsTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newTargetsTableImpl("", "excluded", ""),
	}
}

func newTargetsTableImpl(schemaName, tableName, alias string) targetsTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		NameColumn      = postgres.StringColumn("name")
		CountryColumn   = postgres.StringColumn("country")
		NotesColumn     = postgres.StringColumn("notes")
		StatusColumn    = postgres.StringColumn("status")
		MissionIDColumn = postgres.StringColumn("mission_id")
		allColumns      = postgres.ColumnList{IDColumn, NameColumn, CountryColumn, NotesColumn, StatusColumn, MissionIDColumn}
		mutableColumns  = postgres.ColumnList{NameColumn, CountryColumn, NotesColumn, StatusColumn, MissionIDColumn}
	)

	return targetsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Name:      NameColumn,
		Country:   CountryColumn,
		Notes:     NotesColumn,
		Status:    StatusColumn,
		MissionID: MissionIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
