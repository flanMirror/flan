package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Joker/hpp"
	"github.com/Joker/jade"
)

var (
	misskeyRoot string
	outputDir   string
)

func main() {
	flag.StringVar(&misskeyRoot, "path", "", "path to Misskey root")
	flag.StringVar(&outputDir, "output", "prairie", "output directory")
	flag.Parse()
	if misskeyRoot == "" {
		fmt.Println("error: must specify Misskey root")
		os.Exit(1)
	}

	// create output directory
	outputDir := fmt.Sprintf("%s/%s", os.Getenv("PWD"), outputDir)
	_ = os.Mkdir(outputDir, 0755)

	// read input directory
	templateDir := misskeyRoot + "/packages/backend/src/server/web/views/"
	files, err := os.ReadDir(templateDir)
	if err != nil {
		fmt.Println("error while reading directory:", err)
		os.Exit(1)
	}
	os.Chdir(templateDir)


	for _, file := range files {
		fmt.Println("translating", file.Name())

		// read Pug file
		content, err := os.ReadFile(templateDir + file.Name())
		if err != nil {
			fmt.Println("error while reading file:", err)
			os.Exit(1)
		}

		// translate Pug to html/template
		tmpl, err := jade.Parse("jade", content)
		if err != nil {
			fmt.Println("error while parsing template:", err)
			os.Exit(1)
		}

		// .pug -> .tmpl
		filename := file.Name()[:len(file.Name())-3] + "tmpl"

		// write to file
		err = os.WriteFile(outputDir+"/"+filename, []byte(hpp.PrPrint(tmpl)), 0644)
		if err != nil {
			fmt.Println("error while writing file:", err)
			os.Exit(1)
		}
	}
}
