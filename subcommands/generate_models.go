package subcommands

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/bxcodec/gclean/subcommands/models"
	_ "github.com/go-sql-driver/mysql" //mysql driver
	// "html/template"
)

func (s *Subs) generateModels(dataSend *models.DataGenerator) {

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
