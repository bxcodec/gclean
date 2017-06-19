package subcommands

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	_ "github.com/go-sql-driver/mysql" //mysql driver
	// "html/template"
	"time"
)

type ModelGenerator struct {
	TimeStamp  time.Time
	ModelName  string
	Attributes []Attribute
	Imports    map[string]*Import
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

func fetch(query string, args ...interface{}) ([]*ColumnSchema, error) {

	dbHost := "127.0.0.1"
	dbPort := "33060"
	dbUser := "root"
	dbPass := "password"
	dbName := "article"
	dsn := dbUser + `:` + dbPass + `@tcp(` + dbHost + `:` + dbPort + `)/` + dbName + `?parseTime=1&loc=Asia%2FJakarta`
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		fmt.Println(err)
	}
	defer dbConn.Close()

	rows, err := dbConn.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*ColumnSchema, 0)
	for rows.Next() {

		t := new(ColumnSchema)
		err = rows.Scan(
			&t.TableName,
			&t.ColumnName,
			&t.IsNullable,
			&t.DataType,
			&t.CharacterMaximumLength,
			&t.NumericPrecision,
			&t.NumericScale,
			&t.ColumnType,
			&t.ColumnKey,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func FetchSchema(tableName string) ([]*ColumnSchema, error) {

	query := `SELECT TABLE_NAME, COLUMN_NAME, IS_NULLABLE, DATA_TYPE,
		CHARACTER_MAXIMUM_LENGTH, NUMERIC_PRECISION, NUMERIC_SCALE, COLUMN_TYPE,
		COLUMN_KEY FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = ? ORDER BY TABLE_NAME, ORDINAL_POSITION`

	return fetch(query, tableName)

}

func ExtractModel(schemaList []*ColumnSchema) []ModelGenerator {
	last := schemaList[0].TableName

	var model ModelGenerator
	model.ModelName = last
	var modelList []ModelGenerator
	var attrList []Attribute
	imports := make(map[string]*Import)
	for i, schema := range schemaList {

		if last != schema.TableName {
			model.Attributes = attrList
			model.Imports = imports

			attrList = nil
			modelList = append(modelList, model)

			model.ModelName = schema.TableName
			last = schema.TableName
		} else {
			tipeData := ""

			switch schema.DataType {
			case "char", "varchar", "enum", "set", "text", "longtext", "mediumtext", "tinytext":
				tipeData = "string"
				break
			case "blob", "mediumblob", "longblob", "varbinary", "binary":
				tipeData = "[]byte"
				break
			case "date", "time", "datetime", "timestamp":

				imports["time"] = &Import{
					Alias: "time",
					Path:  "time",
				}
				tipeData = "time.Time"
				break
			case "bit", "tinyint", "smallint", "int", "mediumint", "bigint":
				tipeData = "int64"
				break
			case "float", "decimal", "double":
				tipeData = "float64"
				break
			}

			a := Attribute{
				Name: schema.ColumnName,
				Type: tipeData,
			}
			attrList = append(attrList, a)
		}

		if i == len(schemaList)-1 {
			model.Imports = imports
			model.ModelName = last
			model.Attributes = attrList
			modelList = append(modelList, model)
		}
	}
	return modelList
}

func (s *Subs) generateModels(dataSend *ModelGenerator) {

	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles("template/models.tpl")

	if err != nil {
		fmt.Println("GALGAL", err)
		os.Exit(0)
	}

	pathP := "models/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}
	f, err := os.Create(pathP + strings.ToLower(dataSend.ModelName) + ".go")
	if err != nil {
		fmt.Println("Erorr")
	}

	defer f.Close()
	err = temp.ExecuteTemplate(f, "models.tpl", dataSend)

	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(0)
	}
}
