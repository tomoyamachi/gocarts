package output

import (
	"bytes"
	"github.com/tomoyamachi/gocarts/models"
	"text/template"
)

func GenerateJP(alerts map[string][]models.Alert) (string, error) {
	return Generate(alerts, templateJP)
}

func GenerateUS(alerts map[string][]models.Alert) (string, error) {
	return Generate(alerts, templateUS)
}

// Generate go const definition
func Generate(alerts map[string][]models.Alert, tmplstr string) (string, error) {
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

const templateUS = `package alert

var AlertDictUS = map[string][]AlertGolang {
{{range $cveID, $alerts := .}}
    "{{$cveID}}" : { {{range $alert := . -}}
        {
	        CveID       : "{{$cveID}}",
	        URL         : "{{$alert.URL}}",
            Title       : "{{$alert.Title}}",
	        Team        : "us",
        },{{end}}
    },{{end}}
}
`

const templateJP = `package alert

// Alert has XCERT alert information
type AlertGolang struct {
	CveID       string
	URL         string
	Title       string
	Team        string
}

var AlertDictJP = map[string][]AlertGolang {
{{range $cveID, $alerts := .}}
    "{{$cveID}}" : { {{range $alert := . -}}
        {
	        CveID       : "{{$cveID}}",
	        URL         : "{{$alert.URL}}",
            Title       : "{{$alert.Title}}",
	        Team        : "jp",
        },{{end}}
    },{{end}}
}
`
