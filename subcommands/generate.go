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
	models := mysqlExtractor.ExtractModel(data)
	for _, v := range models {

		s.generateModels(&v)

		s.fixingImportsRepo(&v)
		s.generateRepository(&v)

		s.fixingImportsRepoImpl(&v)
		s.generateRepositoryImpl(&v)
	}
	// s.generateRepositoryImpl("mysql", "article")
	s.generateUsecaseTmp("article")
	s.generateDelivery()

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

func (s *Subs) AddGenerate(root *cobra.Command) {

	var cmdDemo = &cobra.Command{
		Use:   "generate ",
		Short: "Generate your Golang projects",
		Run:   s.generate,
	}

	root.AddCommand(cmdDemo)
}
