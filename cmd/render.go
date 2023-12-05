/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/helmfile/helmfile/pkg/filesystem"
	"github.com/helmfile/helmfile/pkg/tmpl"
	"github.com/spf13/cobra"
)

var (
	params           map[string]string
	templateFilePath string
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render a template",
	Long:  `Render a template`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := filepath.Abs(templateFilePath)

		if err != nil {
			log.Fatal(err)
		}

		if _, err := os.Stat(path); err != nil {
			log.Fatal(err)
		}

		data := make(map[string]interface{})
		for k, v := range params {
			data[k] = v
		}

		c := &tmpl.Context{}
		c.SetBasePath(".")
		c.SetFileSystem(filesystem.DefaultFileSystem())

		// register helmfile custom functions
		t := template.New(filepath.Base(path)).Funcs(c.CreateFuncMap())

		t, err = t.ParseFiles(path)

		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Failed to execute template:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)

	renderCmd.Flags().StringToStringVarP(&params, "param", "p", nil, "Template parameters")

	renderCmd.Flags().StringVarP(&templateFilePath, "file", "f", "", "File to render")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
