package main

import (
	"log"
	"os"
	"text/template"
)

type customer struct {
	Name string
}

type car struct {
	Make  string
	Model string
}

type carOwner struct {
	Owner []customer
	Cars  []car
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	john := customer{
		Name: "John Doe",
	}

	jane := customer{
		Name: "Jane Doe",
	}

	jCar1 := car{
		Make:  "Toyota",
		Model: "RAV4",
	}

	jCar2 := car{
		Make:  "Ford",
		Model: "Everest",
	}

	custs := []customer{john, jane}
	jCars := []car{jCar1, jCar2}

	carOwners := carOwner{
		Owner: custs,
		Cars:  jCars,
	}

	err := tpl.Execute(os.Stdout, carOwners)
	if err != nil {
		log.Fatalln(err)
	}
}
