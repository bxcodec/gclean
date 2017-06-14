package subcommands

import (
	"strconv"
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

	for i := 1; i <= 5; i++ {

		dataList = append(dataList, &ModelGenerator{
			ModelName:  "Model" + strconv.Itoa(i),
			Attributes: []*Attribute{id, title, content},
			TimeStamp:  time.Now(),
		})
	}

	for k, v := range dataList {
		s.generateModels(v, k)

	}
	s.generateRepository()
	s.generateRepositoryImpl("mysql", "article")
	s.generateUsecase("article")

}

func (s *Subs) AddGenerate(root *cobra.Command) {

	var cmdDemo = &cobra.Command{
		Use:   "generate ",
		Short: "Generate your Golang projects",
		Run:   s.generate,
	}

	root.AddCommand(cmdDemo)
}
