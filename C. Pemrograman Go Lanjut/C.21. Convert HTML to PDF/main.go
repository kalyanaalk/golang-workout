package main

import (
	"log"
	// "os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	// pdfg, err := wkhtmltopdf.NewPDFGenerator()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// f, err := os.Open("./input.html")
	// if f != nil {
	// 	defer f.Close()
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	// pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	// pdfg.Dpi.Set(300)

	// err = pdfg.Create()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = pdfg.WriteFile("./output.pdf")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Done")

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	page := wkhtmltopdf.NewPage("http://localhost:9000")
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("./output.pdf")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")
}
