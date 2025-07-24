package main

import (
	"log"

	"github.com/alexflint/go-arg"
	"github.com/lechuckroh/go-tmplx/internal/app/tmplx"
)

type args struct {
	Input     string `arg:"positional,required" help:"Input template file path"`
	Output    string `arg:"positional,required" help:"Output file path"`
	EnvFile   string `arg:"env:ENV_FILE,-e,--env" help:".env file to load"`
	EnvPrefix string `arg:"-p,--env-prefix" help:"Environment variable prefix"`
}

var version = "(devel)"

func (args) Version() string {
	return "kong-template " + version
}

func main() {
	var args args
	arg.MustParse(&args)

	// Load environment variables
	envMap := tmplx.LoadEnv(args.EnvFile, args.EnvPrefix)
	for key, value := range envMap {
		log.Printf("%s=%s", key, value)
	}

	// Convert template
	tmplx.ConvertTemplate(args.Input, args.Output, envMap)
	log.Printf("Generated %s", args.Output)
}
