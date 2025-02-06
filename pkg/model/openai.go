package model

import (
	"context"
	"fmt"

	"github.com/modelflux/cli/pkg/util"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/spf13/viper"
)

type openAIModelOptions struct {
	APIKey string `yaml:"api_key"`
}

type openAImodelParameters struct {
}

type OpenAIModel struct {
	options    openAIModelOptions
	parameters openAImodelParameters
}

func (m *OpenAIModel) ValidateAndSetOptions(uOptions map[string]interface{}, cfg *viper.Viper) error {
	// Create a struct from the map using the util package.
	options, err := util.CreateStruct[openAIModelOptions](uOptions)

	if err != nil {
		return err
	}

	m.options = options

	if m.options.APIKey == "" {
		mcfg := cfg.GetStringMapString("model")
		m.options.APIKey = mcfg["key"]
	}
	if m.options.APIKey == "" {
		return fmt.Errorf("missing api_key for openai model")
	}

	return nil

}

func (m *OpenAIModel) ValidateAndSetParameters(uParams map[string]interface{}) error {
	m.parameters = openAImodelParameters{}
	return nil
}

func (m *OpenAIModel) New() error {
	return nil
}

func (m *OpenAIModel) Generate(input string) (string, error) {
	client := openai.NewClient(option.WithAPIKey(m.options.APIKey))
	resp, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(input),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err.Error())
	}
	return resp.Choices[0].Message.Content, nil
}
