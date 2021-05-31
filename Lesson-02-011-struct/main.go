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
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	jan := month{
		Name: "January",
		Days: 31,
	}

	err := tpl.Execute(os.Stdout, jan)

	if err != nil {
		log.Fatalln(err)
	}
}
