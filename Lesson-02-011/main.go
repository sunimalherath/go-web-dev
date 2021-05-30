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
	vehicleMakes := []string{"Toyota", "Hyundai", "Ford", "Kia"}

	err := tpl.Execute(os.Stdout, vehicleMakes)

	if err != nil {
		log.Fatalln(err)
	}
}
