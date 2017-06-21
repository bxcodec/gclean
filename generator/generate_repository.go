package generator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/Masterminds/sprig"
	"github.com/bxcodec/gclean/generator/models"
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
func ToCamelCase(str string) string {
	if len(str) == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	var r0, r1 rune
	var size int

	// leading '_' will appear in output.
	for len(str) > 0 {
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if r0 != '_' {
			break
		}

		buf.WriteRune(r0)
	}

	if len(str) == 0 {
		return buf.String()
	}

	buf.WriteRune(unicode.ToUpper(r0))
	r0, size = utf8.DecodeRuneInString(str)
	str = str[size:]

	for len(str) > 0 {
		r1 = r0
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if r1 == '_' && r0 != '_' {
			r0 = unicode.ToUpper(r0)
		} else {
			buf.WriteRune(r1)
		}
	}

	buf.WriteRune(r0)
	return buf.String()
}

func (s *Subs) generateMocksRepository(data *models.DataGenerator) {

	name := ToCamelCase(data.ModelName) + "Repository"
	cmnd := exec.Command("/Users/iman/go/bin/mockery", "-name="+name)

	w, _ := os.Getwd()

	var out bytes.Buffer

	// set the output to our variable
	cmnd.Stdout = &out
	cmnd.Dir = w + "/repository"

	err := cmnd.Run()
	if err != nil {
		fmt.Println("err : ", err)
	}

	fmt.Println(out.String())
}
