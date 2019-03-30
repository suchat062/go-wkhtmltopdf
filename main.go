package main

import (
	"fmt"
	"go-wkhtmltopdf/html"
)

func main()  {
	err := html.Example("./layout.html", "./pdf_demo.pdf")
	if err != nil {
		fmt.Println(err.Error())
	}
}