package generator

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bxcodec/gclean/generator/models"
	mysqlExc "github.com/bxcodec/gclean/generator/mysql"
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
		fmt.Println("Generating ", m[i].ModelName, " Repository")
		s.generateRepository(&m[i])
		s.generateMocksRepository(&m[i])

		s.fixingImportsRepoImpl(&m[i])

		fmt.Println("Generating ", m[i].ModelName, " Repository Implement")
		s.generateRepositoryImpl(&m[i])

		fmt.Println("Generating ", m[i].ModelName, " Usecase ")
		s.fixingImportsUsecaseInterface(&m[i])
		s.generateUcaseInterface(&m[i])
		s.generateMocksUsecase(&m[i])

		s.fixingImportsUsecase(&m[i])
		fmt.Println("Generating ", m[i].ModelName, " Usecase Implement")
		s.generateUsecaseTmp(&m[i])
		s.fixingImportsUsecaseTest(&m[i])
		s.generateUsecaseTest(&m[i])

		s.fixingImportDeliveryHandler(&m[i])

		fmt.Println("Generating ", m[i].ModelName, " HttpHandler")
		s.generateHandler(&m[i])
		s.fixingImportDeliveryTest(&m[i])
		s.generateDeliveryTest(&m[i])

		m[i] = s.fixingImportDelivery(m[i], mapImport)

	}

	dlv := &models.DeliveryGenerator{}
	dlv.Data = m

	framework := models.Import{Alias: "echo", Path: "github.com/labstack/echo"}
	mapImport["framework"] = framework
	delete(mapImport, "models")

	dlv.Imports = mapImport

	fmt.Println("Generating  All Http Handler Repository")
	s.generateDelivery(dlv)

}

func (s *Subs) fixingImportDeliveryHandler(m *models.DataGenerator) {

	mapImport := make(map[string]models.Import)

	a := models.Import{Alias: "usecase", Path: "github.com/bxcodec/gclean/usecase"}
	framework := models.Import{Alias: "echo", Path: "github.com/labstack/echo"}

	mapImport["framework"] = framework
	mapImport["usecase"] = a

	m.Imports = mapImport

}

func (s *Subs) fixingImportDeliveryTest(m *models.DataGenerator) {

	mapImport := make(map[string]models.Import)

	mocks := models.Import{Alias: "mocks", Path: "github.com/bxcodec/gclean/usecase/mocks"}
	framework := models.Import{Alias: "echo", Path: "github.com/labstack/echo"}

	h := models.Import{Alias: "handler", Path: "github.com/bxcodec/gclean/delivery/http/" + m.ModelName}
	model := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}

	mapImport["framework"] = framework
	mapImport["mocks"] = mocks
	mapImport["models"] = model
	mapImport["handler"] = h

	m.Imports = mapImport

}

func (s *Subs) fixingImportDelivery(m models.DataGenerator, mapImport map[string]models.Import) models.DataGenerator {

	mm := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}

	a := models.Import{Alias: "usecase", Path: "github.com/bxcodec/gclean/usecase"}
	h := models.Import{Alias: m.ModelName + "Handler", Path: "github.com/bxcodec/gclean/delivery/http/" + m.ModelName}

	mapImport["models"] = mm
	mapImport["usecase"] = a
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
func (s *Subs) fixingImportsUsecaseTest(m *models.DataGenerator) {
	mapIp := make(map[string]models.Import)

	model := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	mocks := models.Import{Alias: "mocks", Path: "github.com/bxcodec/gclean/repository/mocks"}
	u := models.Import{Alias: "usecase", Path: "github.com/bxcodec/gclean/usecase/" + m.ModelName}
	mapIp["usecase"] = u
	mapIp["models"] = model
	mapIp["mocks"] = mocks
	m.Imports = mapIp

}
func (s *Subs) fixingImportsUsecase(m *models.DataGenerator) {
	mapIp := make(map[string]models.Import)

	model := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	mapIp["models"] = model

	r := models.Import{Alias: "repository", Path: "github.com/bxcodec/gclean/repository"}
	mapIp["repository"] = r
	u := models.Import{Alias: "usecase", Path: "github.com/bxcodec/gclean/usecase"}
	mapIp["usecase"] = u
	m.Imports = mapIp

}
func (s *Subs) fixingImportsUsecaseInterface(m *models.DataGenerator) {
	mapIp := make(map[string]models.Import)

	model := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	mapIp["models"] = model

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
