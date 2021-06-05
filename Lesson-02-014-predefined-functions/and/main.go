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

type User struct {
	UserName string
	Admin    bool
}

func main() {
	u1 := User{
		UserName: "johndoe",
		Admin:    true,
	}

	u2 := User{
		UserName: "janedoe",
		Admin:    false,
	}

	u3 := User{
		Admin: false,
	}

	users := []User{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatal(err)
	}
}
