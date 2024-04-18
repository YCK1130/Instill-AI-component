package website

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/instill-ai/component/pkg/base"
)

const (
	taskScrapeWebsite = "TASK_SCRAPE_WEBSITE"
)

var (
	//go:embed config/definition.json
	definitionJSON []byte
	//go:embed config/tasks.json
	tasksJSON []byte

	once      sync.Once
	connector base.IConnector
)

type Connector struct {
	base.Connector
}

type Execution struct {
	base.Execution
}

func Init(logger *zap.Logger, usageHandler base.UsageHandler) base.IConnector {
	once.Do(func() {
		connector = &Connector{
			Connector: base.Connector{
				Component: base.Component{Logger: logger, UsageHandler: usageHandler},
			},
		}
		err := connector.LoadConnectorDefinition(definitionJSON, tasksJSON, nil)
		if err != nil {
			logger.Fatal(err.Error())
		}
	})
	return connector
}

func (c *Connector) CreateExecution(defUID uuid.UUID, task string, connection *structpb.Struct, logger *zap.Logger) (base.IExecution, error) {
	e := &Execution{}
	e.Execution = base.CreateExecutionHelper(e, c, defUID, task, connection, logger)
	return e, nil
}

func (e *Execution) Execute(inputs []*structpb.Struct) ([]*structpb.Struct, error) {
	outputs := []*structpb.Struct{}

	for _, input := range inputs {
		switch e.Task {
		case taskScrapeWebsite:
			inputStruct := ScrapeWebsiteInput{}
			err := base.ConvertFromStructpb(input, &inputStruct)
			if err != nil {
				return nil, err
			}

			outputStruct, err := Scrape(inputStruct)
			if err != nil {
				return nil, err
			}
			output, err := base.ConvertToStructpb(outputStruct)
			if err != nil {
				return nil, err
			}
			outputs = append(outputs, output)
		default:
			return nil, fmt.Errorf("not supported task: %s", e.Task)
		}
	}

	return outputs, nil
}

func (c *Connector) Test(defUID uuid.UUID, config *structpb.Struct, logger *zap.Logger) error {
	return nil
}
