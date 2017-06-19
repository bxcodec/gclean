package subcommands

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/bxcodec/gclean/subcommands/models"
)

func (s *Subs) generateRepository(data *models.DataGenerator) {
	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles("template/repositoryInterface.tpl")

	if err != nil {
		fmt.Println("GALGAL", err)
		os.Exit(0)
	}

	pathP := "repository/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}

	f, err := os.Create(pathP + strings.ToLower(data.ModelName) + "_repository.go")
	if err != nil {
		fmt.Println("Erorr")
	}

	defer f.Close()
	err = temp.ExecuteTemplate(f, "repositoryInterface.tpl", data)

	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(0)
	}
}
