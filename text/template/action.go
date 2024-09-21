package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	templateText := "Template start\nAction:{{.}}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(templateText)
	check(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)
	err = tmpl.Execute(os.Stdout, 42)
	check(err)
	err = tmpl.Execute(os.Stdout, true)
	check(err)
	//使用“if”action使模板的某些部分可选
	templateText = "start {{if .}}Dot is true{{end}} finish\n"
	tmpl, err = template.New("test").Parse(templateText)
	err = tmpl.Execute(os.Stdout, true)
	check(err)
	err = tmpl.Execute(os.Stdout, false)
	check(err)
}
