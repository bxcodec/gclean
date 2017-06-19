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
		s.generateRepository(&v)
	}
	s.generateRepositoryImpl("mysql", "article")
	s.generateUsecaseTmp("article")
	s.generateDelivery()

}

func (s *Subs) AddGenerate(root *cobra.Command) {

	var cmdDemo = &cobra.Command{
		Use:   "generate ",
		Short: "Generate your Golang projects",
		Run:   s.generate,
	}

	root.AddCommand(cmdDemo)
}
