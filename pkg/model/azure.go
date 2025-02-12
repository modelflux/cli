package model

import (
	"context"
	"fmt"

	"github.com/modelflux/modelflux/pkg/util"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/azure"
	"github.com/spf13/viper"
)

type azureOpenAIModelOptions struct {
	APIKey     string `yaml:"api_key"`
	Endpoint   string `yaml:"endpoint"`
	Deployment string `yaml:"deployment"`
	Version    string `yaml:"version"`
}

type AzureOpenAIModel struct {
	options azureOpenAIModelOptions
}

func (m *AzureOpenAIModel) ValidateAndSetOptions(uOptions map[string]interface{}, cfg *viper.Viper) error {
	// Create a struct from the map using the util package.
	options, err := util.BuildStruct[azureOpenAIModelOptions](uOptions)

	if err != nil {
		return err
	}

	if options.APIKey == "" || options.Endpoint == "" || options.Deployment == "" || options.Version == "" {
		mcfg := cfg.GetStringMapString("model")
		m.options.APIKey = mcfg["key"]
		m.options.Endpoint = mcfg["endpoint"]
		m.options.Deployment = mcfg["deployment"]
		m.options.Version = mcfg["version"]
	} else {
		m.options = options
	}

	if m.options.APIKey == "" || m.options.Endpoint == "" || m.options.Deployment == "" || m.options.Version == "" {
		return fmt.Errorf("missing required api_key, endpoint, deployment, or version for azure model")
	}

	return nil
}

func (m *AzureOpenAIModel) Init() error {
	return nil
}

func (m *AzureOpenAIModel) Generate(input string) (string, error) {
	client := openai.NewClient(
		azure.WithEndpoint(m.options.Endpoint, m.options.Version),
		azure.WithAPIKey(m.options.APIKey),
	)

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
