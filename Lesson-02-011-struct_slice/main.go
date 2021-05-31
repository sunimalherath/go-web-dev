package main

import (
	"log"
	"os"
	"text/template"
)

type month struct {
	Name string
	Days int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	jan := month{
		Name: "January",
		Days: 31,
	}

	feb := month{
		Name: "February",
		Days: 28,
	}

	months := []month{jan, feb}

	err := tpl.Execute(os.Stdout, months)
	if err != nil {
		log.Fatalln(err)
	}
}
