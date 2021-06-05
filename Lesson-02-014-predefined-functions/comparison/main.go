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
	match1 := struct {
		Score1 int
		Score2 int
	}{
		20,
		34,
	}

	err := tpl.Execute(os.Stdout, match1)
	if err != nil {
		log.Fatal(err)
	}
}
