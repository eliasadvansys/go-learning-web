package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "Hotel 1",
			Address: "Hotel Street 1",
			City:    "Hotel City 1",
			Zip:     "0001",
			Region:  "Hotel Region 1",
		},
		hotel{
			Name:    "Hotel 2",
			Address: "Hotel Street 2",
			City:    "Hotel City 2",
			Zip:     "0002",
			Region:  "Hotel Region 2",
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}

}
