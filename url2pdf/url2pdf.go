package url2pdf

import html2pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"

func Url2PDF(url string) (filename string, err error) {
	filename = "./tmp.pdf"

	// Create new PDF generator
	pdfFile, err := html2pdf.NewPDFGenerator()
	if err != nil {
		return "", err
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
		return "", err
	}

	// Write buffer contents to file on disk
	err = pdfFile.WriteFile(filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}
