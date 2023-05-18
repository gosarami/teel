/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	params      map[string]string
	templateStr string
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tmpl := templateStr

		data := make(map[string]interface{})
		for k, v := range params {
			data[k] = v
		}

		t, err := template.New("mytemplate").Parse(tmpl)
		if err != nil {
			fmt.Println("Failed to parse template:", err)
			return
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

	renderCmd.Flags().StringVarP(&templateStr, "template", "t", "", "Template")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
