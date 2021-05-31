package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	months := map[string]string{
		"01": "Jan",
		"02": "Feb",
		"03": "Mar",
		"04": "Apr",
	}

	err := tpl.Execute(os.Stdout, months)

	if err != nil {
		log.Fatalln(err)
	}
}
