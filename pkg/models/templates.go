package models

import (
	"errors"
	"html/template"
	"io"

	"github.com/ensamblaTec/learning/week2/problema4/pkg/utils"
	"github.com/labstack/echo/v4"
)

var (
	errCannotConvertTemplate = errors.New("cannot create template")
	templatesRoute           = "./web/templates/*.html"
	htmlRoute                = "./web/html/*.html"
	// staticRoute              = ""
)

type Template struct {
	templates *template.Template
}

func (template *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return template.templates.ExecuteTemplate(w, name, data)
}

func CreateItem(route string) *template.Template {
	return template.Must(template.ParseFiles(route))
}

func ConvertTemplate(templates ...string) (*template.Template, error) {
	parse, err := template.ParseFiles(templates...)
	if err != nil {
		return nil, errCannotConvertTemplate
	}

	tmplParse := template.Must(parse, nil)

	return tmplParse, nil
}

func CreateTemplate(template *template.Template) *Template {
	return &Template{
		templates: template,
	}
}

func Init() *Template {
	var htmlFiles []string
	var err error
	htmlFiles, err = utils.GetFilesFromRoute(htmlRoute)
	if err != nil {
		utils.PrintErrorMessage("models/template htmlFiles", err)
	}

	var templateFiles []string
	templateFiles, err = utils.GetFilesFromRoute(templatesRoute)
	if err != nil {
		utils.PrintErrorMessage("models/template templateFiles", err)
		return nil
	}
	templateFiles = append(htmlFiles, templateFiles...)

	templates, err := ConvertTemplate(templateFiles...)
	if err != nil {
		utils.PrintErrorMessage("models/template parse", err)
		return nil
	}
	return &Template{
		templates: templates,
	}
}
