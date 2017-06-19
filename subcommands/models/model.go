package models

import (
	"database/sql"
	"time"
)

type DataGenerator struct {
	TimeStamp  time.Time
	Type       string
	ModelName  string
	Attributes []Attribute
	Imports    map[string]Import
}

type Import struct {
	Alias string
	Path  string
}

type Attribute struct {
	Name string
	Type string
}

type ColumnSchema struct {
	TableName              string
	ColumnName             string
	IsNullable             string
	DataType               string
	CharacterMaximumLength sql.NullInt64
	NumericPrecision       sql.NullInt64
	NumericScale           sql.NullInt64
	ColumnType             string
	ColumnKey              string
}

type DBConnection interface {
	Dsn() string
}
type MysqlConnection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (m *MysqlConnection) Dsn() string {
	dsn := m.User + `:` + m.Password + `@tcp(` + m.Host + `:` + m.Port + `)/` +
		m.DBName + `?parseTime=1&loc=Asia%2FJakarta`
	return dsn
}
