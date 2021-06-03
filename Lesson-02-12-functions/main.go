package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type customer struct {
	FullName string
	Age      int
}

var tpl *template.Template

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	c := customer{
		FullName: "John Doe",
		Age:      42,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", c)
	if err != nil {
		log.Fatalln(err)
	}
}
