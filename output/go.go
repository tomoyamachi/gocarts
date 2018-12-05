package output

import (
	"bytes"
	"github.com/tomoyamachi/gocarts/models"
	"text/template"
)

func GenerateCveDictJP(cves map[string][]models.Alert) (string, error) {
	return GenerateCveDict(cves, templateCveJP)
}

func GenerateCveDictUS(cves map[string][]models.Alert) (string, error) {
	return GenerateCveDict(cves, templateCveUS)
}

func GenerateAlertDictJP(alerts []models.Alert) (string, error) {
	return GenerateAlertDict(alerts, templateAlertJP)
}

func GenerateAlertDictUS(alerts []models.Alert) (string, error) {
	return GenerateAlertDict(alerts, templateAlertUS)
}

// Generate go const definition
func GenerateCveDict(cves map[string][]models.Alert, tmplstr string) (string, error) {
	tmpl, err := template.New("detail").Parse(tmplstr)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil) // create empty buffer
	if err := tmpl.Execute(buf, cves); err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

func GenerateAlertDict(alerts []models.Alert, tmplstr string) (string, error) {
	tmpl, err := template.New("detail").Parse(tmplstr)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil) // create empty buffer
	if err := tmpl.Execute(buf, alerts); err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

const templateCveUS = `package alert

// CveDictEn has CVE-ID key which included USCERT alerts
var CveDictEn = map[string][]string{
{{range $cveID, $alerts := .}}
    "{{$cveID}}" : { {{range $alert := . -}}"{{$alert.URL}}",{{end}}},{{end}}
}
`

const templateCveJP = `package alert

// CveDictJa has CVE-ID key which included JPCERT alerts
var CveDictJa = map[string][]string{
{{range $cveID, $alerts := .}}
    "{{$cveID}}" : { {{range $alert := . -}}"{{$alert.URL}}",{{end}} },{{end}}
}
`

const templateAlertJP = `package alert

// Alert has XCERT alert information
type Alert struct {
	URL         string
	Title       string
	Team        string
}

// AlertDictJa has JPCERT alerts
var AlertDictJa = map[string]Alert{
{{range $alert := . -}}
    "{{$alert.URL}}" : { 
        URL         : "{{$alert.URL}}",
        Title       :  ` + "`" + `{{$alert.Title}}` + "`" + `,
	    Team        : "jp",
    },
{{end}}
}
`

const templateAlertUS = `package alert

// AlertDictEn has USCERT alerts
var AlertDictEn = map[string]Alert{
{{range $alert := . -}}
    "{{$alert.URL}}" : {
        URL         : "{{$alert.URL}}",
        Title       :  ` + "`" + `{{$alert.Title}}` + "`" + `,
	    Team        : "us",
    },
{{end}}
}
`
