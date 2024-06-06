// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mock

import (
	"io"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// CommandRunnerMock implements document.commandRunner
type CommandRunnerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCombinedOutput          func() (ba1 []byte, err error)
	inspectFuncCombinedOutput   func()
	afterCombinedOutputCounter  uint64
	beforeCombinedOutputCounter uint64
	CombinedOutputMock          mCommandRunnerMockCombinedOutput

	funcStdinPipe          func() (w1 io.WriteCloser, err error)
	inspectFuncStdinPipe   func()
	afterStdinPipeCounter  uint64
	beforeStdinPipeCounter uint64
	StdinPipeMock          mCommandRunnerMockStdinPipe
}

// NewCommandRunnerMock returns a mock for document.commandRunner
func NewCommandRunnerMock(t minimock.Tester) *CommandRunnerMock {
	m := &CommandRunnerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CombinedOutputMock = mCommandRunnerMockCombinedOutput{mock: m}

	m.StdinPipeMock = mCommandRunnerMockStdinPipe{mock: m}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mCommandRunnerMockCombinedOutput struct {
	mock               *CommandRunnerMock
	defaultExpectation *CommandRunnerMockCombinedOutputExpectation
	expectations       []*CommandRunnerMockCombinedOutputExpectation
}

// CommandRunnerMockCombinedOutputExpectation specifies expectation struct of the commandRunner.CombinedOutput
type CommandRunnerMockCombinedOutputExpectation struct {
	mock *CommandRunnerMock

	results *CommandRunnerMockCombinedOutputResults
	Counter uint64
}

// CommandRunnerMockCombinedOutputResults contains results of the commandRunner.CombinedOutput
type CommandRunnerMockCombinedOutputResults struct {
	ba1 []byte
	err error
}

// Expect sets up expected params for commandRunner.CombinedOutput
func (mmCombinedOutput *mCommandRunnerMockCombinedOutput) Expect() *mCommandRunnerMockCombinedOutput {
	if mmCombinedOutput.mock.funcCombinedOutput != nil {
		mmCombinedOutput.mock.t.Fatalf("CommandRunnerMock.CombinedOutput mock is already set by Set")
	}

	if mmCombinedOutput.defaultExpectation == nil {
		mmCombinedOutput.defaultExpectation = &CommandRunnerMockCombinedOutputExpectation{}
	}

	return mmCombinedOutput
}

// Inspect accepts an inspector function that has same arguments as the commandRunner.CombinedOutput
func (mmCombinedOutput *mCommandRunnerMockCombinedOutput) Inspect(f func()) *mCommandRunnerMockCombinedOutput {
	if mmCombinedOutput.mock.inspectFuncCombinedOutput != nil {
		mmCombinedOutput.mock.t.Fatalf("Inspect function is already set for CommandRunnerMock.CombinedOutput")
	}

	mmCombinedOutput.mock.inspectFuncCombinedOutput = f

	return mmCombinedOutput
}

// Return sets up results that will be returned by commandRunner.CombinedOutput
func (mmCombinedOutput *mCommandRunnerMockCombinedOutput) Return(ba1 []byte, err error) *CommandRunnerMock {
	if mmCombinedOutput.mock.funcCombinedOutput != nil {
		mmCombinedOutput.mock.t.Fatalf("CommandRunnerMock.CombinedOutput mock is already set by Set")
	}

	if mmCombinedOutput.defaultExpectation == nil {
		mmCombinedOutput.defaultExpectation = &CommandRunnerMockCombinedOutputExpectation{mock: mmCombinedOutput.mock}
	}
	mmCombinedOutput.defaultExpectation.results = &CommandRunnerMockCombinedOutputResults{ba1, err}
	return mmCombinedOutput.mock
}

// Set uses given function f to mock the commandRunner.CombinedOutput method
func (mmCombinedOutput *mCommandRunnerMockCombinedOutput) Set(f func() (ba1 []byte, err error)) *CommandRunnerMock {
	if mmCombinedOutput.defaultExpectation != nil {
		mmCombinedOutput.mock.t.Fatalf("Default expectation is already set for the commandRunner.CombinedOutput method")
	}

	if len(mmCombinedOutput.expectations) > 0 {
		mmCombinedOutput.mock.t.Fatalf("Some expectations are already set for the commandRunner.CombinedOutput method")
	}

	mmCombinedOutput.mock.funcCombinedOutput = f
	return mmCombinedOutput.mock
}

// CombinedOutput implements document.commandRunner
func (mmCombinedOutput *CommandRunnerMock) CombinedOutput() (ba1 []byte, err error) {
	mm_atomic.AddUint64(&mmCombinedOutput.beforeCombinedOutputCounter, 1)
	defer mm_atomic.AddUint64(&mmCombinedOutput.afterCombinedOutputCounter, 1)

	if mmCombinedOutput.inspectFuncCombinedOutput != nil {
		mmCombinedOutput.inspectFuncCombinedOutput()
	}

	if mmCombinedOutput.CombinedOutputMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCombinedOutput.CombinedOutputMock.defaultExpectation.Counter, 1)

		mm_results := mmCombinedOutput.CombinedOutputMock.defaultExpectation.results
		if mm_results == nil {
			mmCombinedOutput.t.Fatal("No results are set for the CommandRunnerMock.CombinedOutput")
		}
		return (*mm_results).ba1, (*mm_results).err
	}
	if mmCombinedOutput.funcCombinedOutput != nil {
		return mmCombinedOutput.funcCombinedOutput()
	}
	mmCombinedOutput.t.Fatalf("Unexpected call to CommandRunnerMock.CombinedOutput.")
	return
}

// CombinedOutputAfterCounter returns a count of finished CommandRunnerMock.CombinedOutput invocations
func (mmCombinedOutput *CommandRunnerMock) CombinedOutputAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCombinedOutput.afterCombinedOutputCounter)
}

// CombinedOutputBeforeCounter returns a count of CommandRunnerMock.CombinedOutput invocations
func (mmCombinedOutput *CommandRunnerMock) CombinedOutputBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCombinedOutput.beforeCombinedOutputCounter)
}

// MinimockCombinedOutputDone returns true if the count of the CombinedOutput invocations corresponds
// the number of defined expectations
func (m *CommandRunnerMock) MinimockCombinedOutputDone() bool {
	for _, e := range m.CombinedOutputMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CombinedOutputMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCombinedOutputCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCombinedOutput != nil && mm_atomic.LoadUint64(&m.afterCombinedOutputCounter) < 1 {
		return false
	}
	return true
}

// MinimockCombinedOutputInspect logs each unmet expectation
func (m *CommandRunnerMock) MinimockCombinedOutputInspect() {
	for _, e := range m.CombinedOutputMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CommandRunnerMock.CombinedOutput")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CombinedOutputMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCombinedOutputCounter) < 1 {
		m.t.Error("Expected call to CommandRunnerMock.CombinedOutput")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCombinedOutput != nil && mm_atomic.LoadUint64(&m.afterCombinedOutputCounter) < 1 {
		m.t.Error("Expected call to CommandRunnerMock.CombinedOutput")
	}
}

type mCommandRunnerMockStdinPipe struct {
	mock               *CommandRunnerMock
	defaultExpectation *CommandRunnerMockStdinPipeExpectation
	expectations       []*CommandRunnerMockStdinPipeExpectation
}

// CommandRunnerMockStdinPipeExpectation specifies expectation struct of the commandRunner.StdinPipe
type CommandRunnerMockStdinPipeExpectation struct {
	mock *CommandRunnerMock

	results *CommandRunnerMockStdinPipeResults
	Counter uint64
}

// CommandRunnerMockStdinPipeResults contains results of the commandRunner.StdinPipe
type CommandRunnerMockStdinPipeResults struct {
	w1  io.WriteCloser
	err error
}

// Expect sets up expected params for commandRunner.StdinPipe
func (mmStdinPipe *mCommandRunnerMockStdinPipe) Expect() *mCommandRunnerMockStdinPipe {
	if mmStdinPipe.mock.funcStdinPipe != nil {
		mmStdinPipe.mock.t.Fatalf("CommandRunnerMock.StdinPipe mock is already set by Set")
	}

	if mmStdinPipe.defaultExpectation == nil {
		mmStdinPipe.defaultExpectation = &CommandRunnerMockStdinPipeExpectation{}
	}

	return mmStdinPipe
}

// Inspect accepts an inspector function that has same arguments as the commandRunner.StdinPipe
func (mmStdinPipe *mCommandRunnerMockStdinPipe) Inspect(f func()) *mCommandRunnerMockStdinPipe {
	if mmStdinPipe.mock.inspectFuncStdinPipe != nil {
		mmStdinPipe.mock.t.Fatalf("Inspect function is already set for CommandRunnerMock.StdinPipe")
	}

	mmStdinPipe.mock.inspectFuncStdinPipe = f

	return mmStdinPipe
}

// Return sets up results that will be returned by commandRunner.StdinPipe
func (mmStdinPipe *mCommandRunnerMockStdinPipe) Return(w1 io.WriteCloser, err error) *CommandRunnerMock {
	if mmStdinPipe.mock.funcStdinPipe != nil {
		mmStdinPipe.mock.t.Fatalf("CommandRunnerMock.StdinPipe mock is already set by Set")
	}

	if mmStdinPipe.defaultExpectation == nil {
		mmStdinPipe.defaultExpectation = &CommandRunnerMockStdinPipeExpectation{mock: mmStdinPipe.mock}
	}
	mmStdinPipe.defaultExpectation.results = &CommandRunnerMockStdinPipeResults{w1, err}
	return mmStdinPipe.mock
}

// Set uses given function f to mock the commandRunner.StdinPipe method
func (mmStdinPipe *mCommandRunnerMockStdinPipe) Set(f func() (w1 io.WriteCloser, err error)) *CommandRunnerMock {
	if mmStdinPipe.defaultExpectation != nil {
		mmStdinPipe.mock.t.Fatalf("Default expectation is already set for the commandRunner.StdinPipe method")
	}

	if len(mmStdinPipe.expectations) > 0 {
		mmStdinPipe.mock.t.Fatalf("Some expectations are already set for the commandRunner.StdinPipe method")
	}

	mmStdinPipe.mock.funcStdinPipe = f
	return mmStdinPipe.mock
}

// StdinPipe implements document.commandRunner
func (mmStdinPipe *CommandRunnerMock) StdinPipe() (w1 io.WriteCloser, err error) {
	mm_atomic.AddUint64(&mmStdinPipe.beforeStdinPipeCounter, 1)
	defer mm_atomic.AddUint64(&mmStdinPipe.afterStdinPipeCounter, 1)

	if mmStdinPipe.inspectFuncStdinPipe != nil {
		mmStdinPipe.inspectFuncStdinPipe()
	}

	if mmStdinPipe.StdinPipeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStdinPipe.StdinPipeMock.defaultExpectation.Counter, 1)

		mm_results := mmStdinPipe.StdinPipeMock.defaultExpectation.results
		if mm_results == nil {
			mmStdinPipe.t.Fatal("No results are set for the CommandRunnerMock.StdinPipe")
		}
		return (*mm_results).w1, (*mm_results).err
	}
	if mmStdinPipe.funcStdinPipe != nil {
		return mmStdinPipe.funcStdinPipe()
	}
	mmStdinPipe.t.Fatalf("Unexpected call to CommandRunnerMock.StdinPipe.")
	return
}

// StdinPipeAfterCounter returns a count of finished CommandRunnerMock.StdinPipe invocations
func (mmStdinPipe *CommandRunnerMock) StdinPipeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStdinPipe.afterStdinPipeCounter)
}

// StdinPipeBeforeCounter returns a count of CommandRunnerMock.StdinPipe invocations
func (mmStdinPipe *CommandRunnerMock) StdinPipeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStdinPipe.beforeStdinPipeCounter)
}

// MinimockStdinPipeDone returns true if the count of the StdinPipe invocations corresponds
// the number of defined expectations
func (m *CommandRunnerMock) MinimockStdinPipeDone() bool {
	for _, e := range m.StdinPipeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StdinPipeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStdinPipeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStdinPipe != nil && mm_atomic.LoadUint64(&m.afterStdinPipeCounter) < 1 {
		return false
	}
	return true
}

// MinimockStdinPipeInspect logs each unmet expectation
func (m *CommandRunnerMock) MinimockStdinPipeInspect() {
	for _, e := range m.StdinPipeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CommandRunnerMock.StdinPipe")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StdinPipeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStdinPipeCounter) < 1 {
		m.t.Error("Expected call to CommandRunnerMock.StdinPipe")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStdinPipe != nil && mm_atomic.LoadUint64(&m.afterStdinPipeCounter) < 1 {
		m.t.Error("Expected call to CommandRunnerMock.StdinPipe")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CommandRunnerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCombinedOutputInspect()

			m.MinimockStdinPipeInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CommandRunnerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *CommandRunnerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCombinedOutputDone() &&
		m.MinimockStdinPipeDone()
}