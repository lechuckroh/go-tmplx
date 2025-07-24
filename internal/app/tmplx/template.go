package tmplx

import (
	"log"
	"os"
	"text/template"
)

func ConvertTemplate(inputFilename string, outputFilename string, envMap map[string]string) {
	// Parse template file
	tmpl, err := template.ParseFiles(inputFilename)
	if err != nil {
		log.Fatalf("Failed to parse template file: %v", err)
	}

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}

	if err := tmpl.Execute(outputFile, envMap); err != nil {
		log.Fatalf("Failed to convert template: %v", err)
	}
}
