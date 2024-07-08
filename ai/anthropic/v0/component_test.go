package anthropic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/instill-ai/component/base"
	"github.com/instill-ai/component/internal/util/httpclient"
	"github.com/instill-ai/x/errmsg"
)

const (
	apiKey  = "123"
	errResp = `
{
  "error": {
    "message": "Incorrect API key provided."
  }
}`
)

func TestComponent_Execute(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()

	bc := base.Component{}
	connector := Init(bc)

	testcases := []struct {
		name        string
		task        string
		path        string
		contentType string
	}{
		{
			name:        "text generation",
			task:        TextGenerationTask,
			path:        messagesPath,
			contentType: httpclient.MIMETypeJSON,
		},
	}

	// TODO we'll likely want to have a test function per task and test at
	// least OK, NOK. For now, only errors are tested in order to verify
	// end-user messages.
	// 2024-06-21 summer intern An-Che: Implemented text generation test case
	for _, tc := range testcases {
		c.Run("nok - "+tc.name+" 401", func(c *qt.C) {
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				c.Check(r.Method, qt.Equals, http.MethodPost)
				c.Check(r.URL.Path, qt.Equals, tc.path)
				c.Check(r.Header.Get("X-Api-Key"), qt.Equals, apiKey)

				c.Check(r.Header.Get("Content-Type"), qt.Matches, tc.contentType)

				w.Header().Set("Content-Type", httpclient.MIMETypeJSON)
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, errResp)
			})

			anthropicServer := httptest.NewServer(h)
			c.Cleanup(anthropicServer.Close)

			setup, err := structpb.NewStruct(map[string]any{
				"base-path": anthropicServer.URL,
				"api-key":   apiKey,
			})
			c.Assert(err, qt.IsNil)

			exec, err := connector.CreateExecution(nil, setup, tc.task)
			c.Assert(err, qt.IsNil)

			pbIn := new(structpb.Struct)
			_, err = exec.Execution.Execute(ctx, []*structpb.Struct{pbIn})
			c.Check(err, qt.IsNotNil)

			want := "Anthropic responded with a 401 status code. Incorrect API key provided."
			c.Check(errmsg.Message(err), qt.Equals, want)
		})
	}
	c.Run("nok - unsupported task", func(c *qt.C) {
		task := "FOOBAR"

		_, err := connector.CreateExecution(nil, nil, task)
		c.Check(err, qt.ErrorMatches, "unsupported task")
	})
}

func TestComponent_Connection(t *testing.T) {
	c := qt.New(t)

	bc := base.Component{}
	connector := Init(bc)

	c.Run("nok - error", func(c *qt.C) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Check(r.Method, qt.Equals, http.MethodGet)

			w.Header().Set("Content-Type", httpclient.MIMETypeJSON)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, errResp)
		})

		anthropicServer := httptest.NewServer(h)
		c.Cleanup(anthropicServer.Close)

		_, err := structpb.NewStruct(map[string]any{
			"base-path": anthropicServer.URL,
		})
		c.Assert(err, qt.IsNil)
	})

	c.Run("ok - disconnected", func(c *qt.C) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Check(r.Method, qt.Equals, http.MethodGet)

			w.Header().Set("Content-Type", httpclient.MIMETypeJSON)
			fmt.Fprintln(w, `{}`)
		})

		anthropicServer := httptest.NewServer(h)
		c.Cleanup(anthropicServer.Close)

		_, err := structpb.NewStruct(map[string]any{
			"base-path": anthropicServer.URL,
		})
		c.Assert(err, qt.IsNil)
	})

	c.Run("ok - connected", func(c *qt.C) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Check(r.Method, qt.Equals, http.MethodGet)

			w.Header().Set("Content-Type", httpclient.MIMETypeJSON)
			fmt.Fprintln(w, `{"data": [{}]}`)
		})

		anthropicServer := httptest.NewServer(h)
		c.Cleanup(anthropicServer.Close)

		setup, err := structpb.NewStruct(map[string]any{
			"base-path": anthropicServer.URL,
		})
		c.Assert(err, qt.IsNil)

		err = connector.Test(nil, setup)
		c.Check(err, qt.IsNil)
	})
}

type MockAnthropicClient struct{}

func (m *MockAnthropicClient) generateTextChat(request messagesReq) (messagesResp, error) {

	messageCount := len(request.Messages)
	message := fmt.Sprintf("Hi! My name is Claude. (messageCount: %d)", messageCount)
	resp := messagesResp{
		ID:         "msg_013Zva2CMHLNnXjNJJKqJ2EF",
		Type:       "message",
		Role:       "assistant",
		Content:    []content{{Text: message, Type: "text"}},
		Model:      "claude-3-5-sonnet-20240620",
		StopReason: "end_turn",
		Usage:      usage{InputTokens: 10, OutputTokens: 25},
	}

	return resp, nil
}

func TestComponent_Generation(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	bc := base.Component{}
	connector := Init(bc)

	mockHistory := []message{
		{Role: "user", Content: []content{{Type: "text", Text: "Answer the following question in traditional chinses"}}},
		{Role: "assistant", Content: []content{{Type: "text", Text: "沒問題"}}},
	}

	tc := struct {
		input    map[string]any
		wantResp MessagesOutput
	}{
		input: map[string]any{"prompt": "Hi! What's your name?", "chat-history": mockHistory},
		wantResp: MessagesOutput{
			Text: "Hi! My name is Claude. (messageCount: 3)",
			Usage: messagesUsage{
				InputTokens:  10,
				OutputTokens: 25,
			},
		},
	}

	c.Run("ok - generation", func(c *qt.C) {
		setup, err := structpb.NewStruct(map[string]any{
			"api-key": apiKey,
		})
		c.Assert(err, qt.IsNil)

		e := &execution{
			ComponentExecution: base.ComponentExecution{Component: connector, SystemVariables: nil, Setup: setup, Task: TextGenerationTask},
			client:             &MockAnthropicClient{},
		}
		e.execute = e.generateText
		exec := &base.ExecutionWrapper{Execution: e}

		pbIn, err := base.ConvertToStructpb(tc.input)
		c.Assert(err, qt.IsNil)

		got, err := exec.Execution.Execute(ctx, []*structpb.Struct{pbIn})
		c.Assert(err, qt.IsNil)

		wantJSON, err := json.Marshal(tc.wantResp)
		c.Assert(err, qt.IsNil)
		c.Check(wantJSON, qt.JSONEquals, got[0].AsMap())
	})
}