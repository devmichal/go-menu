package quote

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/devmichal/foods/internal/domain"
	"github.com/devmichal/foods/internal/interface/connect"
	"github.com/gofiber/fiber/v2"
	"html/template"
	"net/http"
)

type Handler struct {
}

func NewPdfMenu() *Handler {
	return &Handler{}
}

func (h *Handler) PdfMenu(c *fiber.Ctx) []byte {
	var menu domain.Menu
	db := connect.Connect()
	_ = db.First(&menu, 1)

	c.GetReqHeaders()

	pdf, err := generatePDF(menu)

	if err == nil {
		//return c.JSON("Error created pdf")
	}

	c.Accepts("Content-Disposition", "attachment; filename=kittens.pdf")
	c.Accepts("Content-Type", "application/pdf")
	c.Status(http.StatusOK)
	c.Write(pdf)

	//return c.JSON(menu.Menu)
	return c.Response().Body()
}

func generatePDF(menu domain.Menu) ([]byte, error) {
	var templ *template.Template
	var err error

	// use Go's default HTML template generation tools to generate your HTML
	if templ, err = template.ParseFiles("label.html"); err != nil {
		return nil, err
	}

	// apply the parsed HTML template data and keep the result in a Buffer
	var body bytes.Buffer
	if err = templ.Execute(&body, menu); err != nil {
		return nil, err
	}

	// initalize a wkhtmltopdf generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// read the HTML page as a PDF page
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

	// enable this if the HTML file contains local references such as images, CSS, etc.
	page.EnableLocalFileAccess.Set(true)

	// add the page to your generator
	pdfg.AddPage(page)

	// manipulate page attributes as needed
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	// magic
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
