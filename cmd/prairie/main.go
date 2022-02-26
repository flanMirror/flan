package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/Joker/jade"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/minify"
)

var (
	misskeyRoot string
	outputDir   string
	workingDir  string
	wipe        bool
)

func init() {
	if wd, err := os.Getwd(); err != nil {
		fmt.Println("error getting working directory:", err)
		os.Exit(1)
	} else {
		workingDir = wd
	}
}

func main() {
	flag.StringVar(&misskeyRoot, "i", "", "path to Misskey root")
	flag.StringVar(&outputDir, "o", "assets/template", "path to output directory")
	flag.BoolVar(&wipe, "w", false, "wipe output directory if exists")
	flag.Parse()
	if misskeyRoot == "" {
		fmt.Println("must specify Misskey root")
		os.Exit(1)
	}

	// read input directory
	const templateDirRelative = "packages/backend/src/server/web/views/"
	templateDir := path.Join(misskeyRoot, templateDirRelative)
	files, err := os.ReadDir(templateDir)
	if err != nil {
		fmt.Println("error while reading directory:", err)
		os.Exit(1)
	}
	if err = os.Chdir(templateDir); err != nil {
		fmt.Println("error switching to template directory:", err)
		os.Exit(1)
	}

	// create output directory
	outputDir = path.Join(workingDir, outputDir)
create:
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		if !os.IsExist(err) {
			fmt.Println("error creating output directory:", err)
			os.Exit(1)
		} else {
			if !wipe {
				fmt.Println("directory already exists and wipe is not set")
				os.Exit(1)
			} else {
				if err = os.RemoveAll(outputDir); err != nil {
					fmt.Println("error wiping output dir:", err)
					os.Exit(1)
				} else {
					goto create
				}
			}
		}
	}

	html.DefaultMinifier.KeepComments = true

	for _, file := range files {
		fmt.Println("translating", file.Name())

		// read Pug file
		content, err := os.ReadFile(file.Name())
		if err != nil {
			fmt.Println("error while reading file:", err)
			os.Exit(1)
		}

		// translate Pug to html/template
		tmpl, err := jade.Parse(file.Name(), content)
		if err != nil {
			fmt.Println("error while parsing template:", err)
			os.Exit(1)
		}

		// .pug -> .tmpl
		filename := file.Name()[:len(file.Name())-3] + "tmpl"

		// minify
		min, err := minify.HTML(tmpl)
		if err != nil {
			fmt.Println("error minifying", filename+":", err)
		}

		// write to file
		if err = os.WriteFile(outputDir+"/"+filename, []byte(min), 0644); err != nil {
			fmt.Println("error while writing file:", err)
			os.Exit(1)
		}
	}
}
