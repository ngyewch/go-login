package generate

import (
	"encoding/json"
	"github.com/ngyewch/go-login/resources"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Generator struct {
	inputDirectory  string
	outputDirectory string
	templates       *template.Template
	config          Config
}

type Config struct {
	AppName     string        `json:"appName"`
	Actions     []*Action     `json:"actions"`
	Backgrounds []*Background `json:"backgrounds"`
}

type Action struct {
	Text string `json:"text"`
	Link string `json:"link"`
	Icon string `json:"icon"`
}

type Background struct {
	Name            string `json:"name"`
	ProfileUrl      string `json:"profileUrl"`
	Url             string `json:"url"`
	BackgroundImage string `json:"backgroundImage"`
}

func NewGenerator(inputDirectory string, outputDirectory string) (*Generator, error) {
	tmpl, err := template.ParseFS(resources.TemplatesFS, "templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	var config Config
	configFile, err := os.Open(filepath.Join(inputDirectory, "config.json"))
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	configBytes, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}

	return &Generator{
		inputDirectory:  inputDirectory,
		outputDirectory: outputDirectory,
		templates:       tmpl,
		config:          config,
	}, nil
}

func (g *Generator) Generate() error {
	err := g.executeTemplate("index.html.tmpl", g.config, "index.html")
	if err != nil {
		return err
	}

	for _, background := range g.config.Backgrounds {
		if !strings.HasPrefix(background.BackgroundImage, "https://") && !strings.HasPrefix(background.BackgroundImage, "http://") {
			err := g.copyFile(background.BackgroundImage)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Generator) executeTemplate(templateName string, data any, relativePath string) error {
	err := os.MkdirAll(g.outputDirectory, 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(g.outputDirectory, relativePath))
	if err != nil {
		return err
	}
	defer f.Close()

	err = g.templates.ExecuteTemplate(f, templateName, data)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) copyFile(relativePath string) error {
	r, err := os.Open(filepath.Join(g.inputDirectory, relativePath))
	if err != nil {
		return err
	}
	defer r.Close()

	outputPath := filepath.Join(g.outputDirectory, relativePath)
	outputDir := filepath.Dir(outputPath)

	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		return err
	}

	w, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	return nil
}
