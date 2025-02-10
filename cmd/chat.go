package main

import (
	"fmt"

	"github.com/modelflux/modelflux/pkg/model"
	"github.com/spf13/cobra"
)

var Model string

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Send a message to the model. This is just a test command.",
	Run: func(cmd *cobra.Command, args []string) {
		var input = args[0]
		var m model.Model

		if Model == "azure" {
			m = &model.AzureOpenAIModel{}
		} else if Model == "openai" {
			m = &model.OpenAIModel{}
		} else {
			fmt.Printf("model %s not found", Model)
			return
		}

		if err := m.New(Config); err != nil {
			fmt.Printf("error initializing model: %v", err)
			return
		}

		resp, err := m.Generate(input)
		if err != nil {
			fmt.Printf("error generating response: %v", err)
			return
		}
		fmt.Println(resp)
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(chatCmd)

	chatCmd.Flags().StringVarP(&Model, "model", "m", "azure", "Model to use (required)")
}
