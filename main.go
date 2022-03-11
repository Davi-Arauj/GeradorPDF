package main

import (
	"fmt"
	"geradorPDF/htmlParser"
	"geradorPDF/pdfGenerator"
	"os"
)

type Data struct {
	Name string
}

func main() {

	dataHTML := Data{
		Name: "Davi",
	}

	h := htmlParser.New("tmp")
	wk := pdfGenerator.NewWkHtmlToPDF("tmp")

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer os.Remove(htmlGenerated)
	fmt.Println("HTML gerado", htmlGenerated)

	filePDFName, err := wk.Create(htmlGenerated)
	if err != nil {
		return
	}

	fmt.Println("PDF gerado", filePDFName)

}
