package url2pdf

import (
	html2pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
)

func Url2PDF(url, tempPath string) (err error) {
	log.Printf("Url to be rendered: %s", url)

	// Create new PDF generator
	pdfFile, err := html2pdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	// Set global options
	pdfFile.Dpi.Set(600)
	pdfFile.Orientation.Set(html2pdf.OrientationLandscape)
	pdfFile.Grayscale.Set(false)

	// Create a new input page from URL
	page := html2pdf.NewPage(url)

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(0)
	page.Zoom.Set(0.5)

	// Add to document
	pdfFile.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfFile.Create()
	if err != nil {
		return err
	}

	// Write buffer contents to file on disk
	err = pdfFile.WriteFile(tempPath)
	if err != nil {
		return err
	}

	return nil
}
