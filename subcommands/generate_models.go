package subcommands

import (
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/Masterminds/sprig"

	// "html/template"
	"time"
)

type ModelGenerator struct {
	TimeStamp  time.Time
	ModelName  string
	Attributes []*Attribute
}

type Attribute struct {
	Name string
	Type string
}

func (s *Subs) generateModels(dataSend *ModelGenerator, k int) {

	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles("template/models.tpl")

	if err != nil {
		fmt.Println("GALGAL", err)
		os.Exit(0)
	}

	pathP := "models/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}
	f, err := os.Create(pathP + "sample" + strconv.Itoa(k) + ".go")
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
