//go:generate compogen readme ./config ./README.mdx
package redis

import (
	"context"
	_ "embed"
	"fmt"
	"sync"

	"google.golang.org/protobuf/types/known/structpb"

	"github.com/instill-ai/component/base"
)

const (
	taskWriteChatMessage           = "TASK_WRITE_CHAT_MESSAGE"
	taskWriteMultiModalChatMessage = "TASK_WRITE_MULTI_MODAL_CHAT_MESSAGE"
	taskRetrieveChatHistory        = "TASK_RETRIEVE_CHAT_HISTORY"
)

var (
	//go:embed config/definition.json
	definitionJSON []byte
	//go:embed config/setup.json
	setupJSON []byte
	//go:embed config/tasks.json
	tasksJSON []byte

	once sync.Once
	comp *component
)

type component struct {
	base.Component
}

type execution struct {
	base.ComponentExecution
}

func Init(bc base.Component) *component {
	once.Do(func() {
		comp = &component{Component: bc}
		err := comp.LoadDefinition(definitionJSON, setupJSON, tasksJSON, nil)
		if err != nil {
			panic(err)
		}
	})
	return comp
}

func (c *component) CreateExecution(sysVars map[string]any, setup *structpb.Struct, task string) (*base.ExecutionWrapper, error) {
	return &base.ExecutionWrapper{Execution: &execution{
		ComponentExecution: base.ComponentExecution{Component: c, SystemVariables: sysVars, Setup: setup, Task: task},
	}}, nil
}

func (e *execution) Execute(ctx context.Context, inputs []*structpb.Struct) ([]*structpb.Struct, error) {
	outputs := []*structpb.Struct{}

	client, err := NewClient(e.Setup)
	if err != nil {
		return outputs, err
	}
	defer client.Close()

	for _, input := range inputs {
		var output *structpb.Struct
		switch e.Task {
		case taskWriteChatMessage:
			inputStruct := ChatMessageWriteInput{}
			err := base.ConvertFromStructpb(input, &inputStruct)
			if err != nil {
				return nil, err
			}
			outputStruct := WriteMessage(client, inputStruct)
			output, err = base.ConvertToStructpb(outputStruct)
			if err != nil {
				return nil, err
			}
		case taskWriteMultiModalChatMessage:
			inputStruct := ChatMultiModalMessageWriteInput{}
			err := base.ConvertFromStructpb(input, &inputStruct)
			if err != nil {
				return nil, err
			}
			outputStruct := WriteMultiModelMessage(client, inputStruct)
			output, err = base.ConvertToStructpb(outputStruct)
			if err != nil {
				return nil, err
			}
		case taskRetrieveChatHistory:
			inputStruct := ChatHistoryRetrieveInput{}
			err := base.ConvertFromStructpb(input, &inputStruct)
			if err != nil {
				return nil, err
			}
			outputStruct := RetrieveSessionMessages(client, inputStruct)
			output, err = base.ConvertToStructpb(outputStruct)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported task: %s", e.Task)
		}
		outputs = append(outputs, output)
	}
	return outputs, nil
}

func (c *component) Test(sysVars map[string]any, setup *structpb.Struct) error {
	client, err := NewClient(setup)
	if err != nil {
		return err
	}
	defer client.Close()

	// Ping the Redis server to check the setup
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}
