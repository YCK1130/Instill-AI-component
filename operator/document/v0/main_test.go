package document

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"code.sajari.com/docconv"
	"github.com/frankban/quicktest"
	"github.com/instill-ai/component/base"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestOperator(t *testing.T) {
	c := quicktest.New(t)

	fileContent, _ := os.ReadFile("testdata/test.txt")
	base64DataURI := fmt.Sprintf("data:%s;base64,%s", docconv.MimeTypeByExtension("testdata/test.txt"), base64.StdEncoding.EncodeToString(fileContent))

	testcases := []struct {
		name  string
		task  string
		input structpb.Struct
	}{
		{
			name: "convert to text",
			task: "TASK_CONVERT_TO_TEXT",
			input: structpb.Struct{
				Fields: map[string]*structpb.Value{
					"doc": {Kind: &structpb.Value_StringValue{StringValue: base64DataURI}},
				},
			},
		},
	}
	bc := base.Component{}
	ctx := context.Background()
	for i := range testcases {
		tc := &testcases[i]
		c.Run(tc.name, func(c *quicktest.C) {
			component := Init(bc)
			c.Assert(component, quicktest.IsNotNil)

			execution, err := component.CreateExecution(map[string]any{}, nil, tc.task)
			c.Assert(err, quicktest.IsNil)
			c.Assert(execution, quicktest.IsNotNil)

			input := []*structpb.Struct{&tc.input}

			outputs, err := execution.Execution.Execute(ctx, input)

			c.Assert(err, quicktest.IsNil)
			c.Assert(outputs, quicktest.HasLen, 1)

		})
	}
}