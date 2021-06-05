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
	xs := []string{"jul", "aug", "sep", "oct", "nov", "dec"}

	data := struct {
		Months []string
		Year   int
	}{
		xs,
		2021,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
