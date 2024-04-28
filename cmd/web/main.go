package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func render (writer http.ResponseWriter, templateName string) {
	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", templateName))

	templateSlice = append(templateSlice, partials...)

	parseTemplate, err := template.ParseFiles(templateSlice...)
	if (err != nil) {
		fmt.Println("Server error occurred")
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	if err := parseTemplate.Execute(writer, nil); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func main () {
	http.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request) {
		render(writer, "test.page.gohtml")
	})


	fmt.Println("Starting front-end service on Port 2000")
	err := http.ListenAndServe(":2000", nil)

	if err != nil {
		fmt.Println("Panic error occurred")
		log.Panic(err)
	}
}
