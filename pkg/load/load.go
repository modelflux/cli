package load

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type WorkflowSchema struct {
	Name   string                        `yaml:"name"`
	Models map[string]ModelConfiguration `yaml:"models"`
	Tools  map[string]Tool               `yaml:"tools"`
	Task   Task                          `yaml:"task"`
}

type ModelConfiguration struct {
	Type  string `yaml:"type"`
	Model string `yaml:"model"`
}

type Tool struct {
	Source      string            `yaml:"source"`
	ToolOptions map[string]string `yaml:"toolOptions"`
}

type Task struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name       string            `yaml:"name"`
	ID         string            `yaml:"id,omitempty"`
	Model      string            `yaml:"model,omitempty"`
	Tool       string            `yaml:"tool,omitempty"`
	Parameters map[string]string `yaml:"parameters,omitempty"`
	Output     string            `yaml:"output,omitempty"`
}

func loadYAML(filePath string) (*WorkflowSchema, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var workflow WorkflowSchema
	err = yaml.Unmarshal(data, &workflow)
	if err != nil {
		return nil, err
	}

	return &workflow, nil
}

func Load(workflowName string) *WorkflowSchema {
	workflowPath := fmt.Sprintf("workflows/%s.yaml", workflowName)
	fmt.Println("Loading workflow:", workflowPath)
	workflow, err := loadYAML(workflowPath)
	if err != nil {
		fmt.Println("Error loading YAML:", err)
		return nil
	}
	return workflow
}
