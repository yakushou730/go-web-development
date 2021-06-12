package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name    string
		Phone   string
		Points  []int
		Mapping map[string]string
	}{
		Name:   "<script>alert('Howdy!')</script>",
		Phone:  "3345678",
		Points: []int{1, 2, 3, 4, 5},
		Mapping: map[string]string{
			"A": "100",
			"B": "200",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
