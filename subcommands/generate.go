package subcommands

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type Subs struct {
}

func (s *Subs) generate(cmd *cobra.Command, args []string) {
	dataList := make([]*ModelGenerator, 0)

	id := &Attribute{
		Name: "ID",
		Type: "int64",
	}
	title := &Attribute{
		Name: "Title",
		Type: "string",
	}
	content := &Attribute{
		Name: "Content",
		Type: "string",
	}

	dataList = append(dataList, &ModelGenerator{
		ModelName:  "Article",
		Attributes: []Attribute{*id, *title, *content},
		TimeStamp:  time.Now(),
	})

	data, err := FetchSchema("article")
	if err != nil {
		fmt.Println(err)
		panic(err)
		// os.Exit(0)
	}
	models := ExtractModel(data)
	for _, v := range models {

		s.generateModels(&v)

	}
	s.generateRepository()
	s.generateRepositoryImpl("mysql", "article")
	s.generateUsecase("article")
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
