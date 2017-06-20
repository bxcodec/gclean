package subcommands

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bxcodec/gclean/subcommands/models"
	mysqlExc "github.com/bxcodec/gclean/subcommands/mysql"
)

type Subs struct {
}

func (s *Subs) generate(cmd *cobra.Command, args []string) {

	dbHost := "127.0.0.1"
	dbPort := "33060"
	dbUser := "root"
	dbPass := "password"
	dbName := "article"
	mysqlConfig := &models.MysqlConnection{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPass,
		DBName:   dbName,
	}

	dbConn, err := sql.Open(`mysql`, mysqlConfig.Dsn())
	if err != nil {
		fmt.Println(err)
	}
	defer dbConn.Close()

	mysqlExtractor := mysqlExc.MysqlExtractor{DBCon: dbConn}

	data, err := mysqlExtractor.FetchSchema("article")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	m := mysqlExtractor.ExtractModel(data)
	mapImport := make(map[string]models.Import)
	for i := 0; i < len(m); i++ {

		s.generateModels(&m[i])

		s.fixingImportsRepo(&m[i])
		s.generateRepository(&m[i])

		s.fixingImportsRepoImpl(&m[i])
		s.generateRepositoryImpl(&m[i])

		s.fixingImportsUsecase(&m[i])
		s.generateUsecaseTmp(&m[i])

		s.fixingImportDeliveryHandler(&m[i])
		s.generateHandler(&m[i])

		m[i] = s.fixingImportDelivery(m[i], mapImport)

		// fmt.Println(" : ", v.Imports)

	}

	dlv := &models.DeliveryGenerator{}
	dlv.Data = m
	//
	// mm := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	// t := models.Import{Alias: "time", Path: "time"}
	// ss := models.Import{Alias: "sql", Path: "database/sql"}
	// r := models.Import{Alias: "repository", Path: "github.com/bxcodec/gclean/repository"}
	// a := models.Import{Alias: "articleUcase", Path: "github.com/bxcodec/gclean/delivery/http/article"}
	// mapImport["models"] = mm
	// mapImport["time"] = t
	// mapImport["sql"] = ss
	// mapImport["repository"] = r
	// mapImport["article"] = a
	framework := models.Import{Alias: "echo", Path: "github.com/labstack/echo"}
	mapImport["framework"] = framework
	delete(mapImport, "models")

	dlv.Imports = mapImport
	s.generateDelivery(dlv)

}

func (s *Subs) fixingImportDeliveryHandler(m *models.DataGenerator) {

	mapImport := make(map[string]models.Import)

	aliasModel := m.ModelName + "Ucase"
	// mm := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}

	// ss := models.Import{Alias: "sql", Path: "database/sql"}
	// r := models.Import{Alias: "repository", Path: "github.com/bxcodec/gclean/repository"}
	a := models.Import{Alias: aliasModel, Path: "github.com/bxcodec/gclean/usecase/" + m.ModelName}
	// mapImport["models"] = mm
	// mapImport["time"] = t
	// mapImport["sql"] = ss
	// mapImport["repository"] = r
	framework := models.Import{Alias: "echo", Path: "github.com/labstack/echo"}

	mapImport["framework"] = framework

	mapImport[m.ModelName+"usecase"] = a

	m.Imports = mapImport

}

func (s *Subs) fixingImportDelivery(m models.DataGenerator, mapImport map[string]models.Import) models.DataGenerator {

	aliasModel := m.ModelName + "Ucase"
	mm := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	// t := models.Import{Alias: "time", Path: "time"}
	// ss := models.Import{Alias: "sql", Path: "database/sql"}
	// r := models.Import{Alias: "repository", Path: "github.com/bxcodec/gclean/repository"}
	a := models.Import{Alias: aliasModel, Path: "github.com/bxcodec/gclean/usecase/" + m.ModelName}
	h := models.Import{Alias: m.ModelName + "Handler", Path: "github.com/bxcodec/gclean/delivery/http/" + m.ModelName}
	mapImport["models"] = mm
	// mapImport["time"] = t
	// mapImport["sql"] = ss
	// mapImport["repository"] = r
	mapImport[m.ModelName+"usecase"] = a
	mapImport[m.ModelName+"handler"] = h

	m.Imports = mapImport
	return m

}

func (s *Subs) fixingImportsRepo(m *models.DataGenerator) {
	model := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	mapIp := make(map[string]models.Import)
	mapIp["models"] = model
	m.Imports = mapIp

}

func (s *Subs) fixingImportsRepoImpl(m *models.DataGenerator) {
	mapIp := make(map[string]models.Import)

	model := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	mapIp["models"] = model

	sq := models.Import{Alias: "sql", Path: "database/sql"}
	mapIp["sql"] = sq

	r := models.Import{Alias: "repository", Path: "github.com/bxcodec/gclean/repository"}
	mapIp["repository"] = r

	m.Imports = mapIp

}
func (s *Subs) fixingImportsUsecase(m *models.DataGenerator) {
	mapIp := make(map[string]models.Import)

	model := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	mapIp["models"] = model

	r := models.Import{Alias: "repository", Path: "github.com/bxcodec/gclean/repository"}
	mapIp["repository"] = r

	m.Imports = mapIp

}

func (s *Subs) AddGenerate(root *cobra.Command) {

	var cmdDemo = &cobra.Command{
		Use:   "generate ",
		Short: "Generate your Golang projects",
		Run:   s.generate,
	}

	root.AddCommand(cmdDemo)
}
