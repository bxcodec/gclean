package generator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/bxcodec/gclean/generator/models"
)

func (s *Subs) generateUcaseInterface(data *models.DataGenerator) {
	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles("template/usecaseInterface.tpl")

	if err != nil {
		fmt.Println("GALGAL", err)
		os.Exit(0)
	}

	pathP := "usecase/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}

	f, err := os.Create(pathP + strings.ToLower(data.ModelName) + "_usecase.go")
	if err != nil {
		fmt.Println("Erorr")
	}

	defer f.Close()
	err = temp.ExecuteTemplate(f, "usecaseInterface.tpl", data)

	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(0)
	}
}

func (s *Subs) generateMocksUsecase(data *models.DataGenerator) {

	name := ToCamelCase(data.ModelName) + "Usecase"
	cmnd := exec.Command("/Users/iman/go/bin/mockery", "-name="+name)

	w, _ := os.Getwd()

	var out bytes.Buffer

	// set the output to our variable
	cmnd.Stdout = &out
	cmnd.Dir = w + "/usecase"

	err := cmnd.Run()
	if err != nil {
		fmt.Println("err : ", err)
	}

	fmt.Println(out.String())
}
