package mock

// Code generated by http://github.com/gojuno/minimock (3.0.8). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/gojuno/minimock/v3"
)

// CloudWatchMock implements cloudmetrics.CloudWatch
type CloudWatchMock struct {
	t minimock.Tester

	funcPutMetricData          func(input *cloudwatch.PutMetricDataInput) (pp1 *cloudwatch.PutMetricDataOutput, err error)
	inspectFuncPutMetricData   func(input *cloudwatch.PutMetricDataInput)
	afterPutMetricDataCounter  uint64
	beforePutMetricDataCounter uint64
	PutMetricDataMock          mCloudWatchMockPutMetricData
}

// NewCloudWatchMock returns a mock for cloudmetrics.CloudWatch
func NewCloudWatchMock(t minimock.Tester) *CloudWatchMock {
	m := &CloudWatchMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.PutMetricDataMock = mCloudWatchMockPutMetricData{mock: m}
	m.PutMetricDataMock.callArgs = []*CloudWatchMockPutMetricDataParams{}

	return m
}

type mCloudWatchMockPutMetricData struct {
	mock               *CloudWatchMock
	defaultExpectation *CloudWatchMockPutMetricDataExpectation
	expectations       []*CloudWatchMockPutMetricDataExpectation

	callArgs []*CloudWatchMockPutMetricDataParams
	mutex    sync.RWMutex
}

// CloudWatchMockPutMetricDataExpectation specifies expectation struct of the CloudWatch.PutMetricData
type CloudWatchMockPutMetricDataExpectation struct {
	mock    *CloudWatchMock
	params  *CloudWatchMockPutMetricDataParams
	results *CloudWatchMockPutMetricDataResults
	Counter uint64
}

// CloudWatchMockPutMetricDataParams contains parameters of the CloudWatch.PutMetricData
type CloudWatchMockPutMetricDataParams struct {
	input *cloudwatch.PutMetricDataInput
}

// CloudWatchMockPutMetricDataResults contains results of the CloudWatch.PutMetricData
type CloudWatchMockPutMetricDataResults struct {
	pp1 *cloudwatch.PutMetricDataOutput
	err error
}

// Expect sets up expected params for CloudWatch.PutMetricData
func (mmPutMetricData *mCloudWatchMockPutMetricData) Expect(input *cloudwatch.PutMetricDataInput) *mCloudWatchMockPutMetricData {
	if mmPutMetricData.mock.funcPutMetricData != nil {
		mmPutMetricData.mock.t.Fatalf("CloudWatchMock.PutMetricData mock is already set by Set")
	}

	if mmPutMetricData.defaultExpectation == nil {
		mmPutMetricData.defaultExpectation = &CloudWatchMockPutMetricDataExpectation{}
	}

	mmPutMetricData.defaultExpectation.params = &CloudWatchMockPutMetricDataParams{input}
	for _, e := range mmPutMetricData.expectations {
		if minimock.Equal(e.params, mmPutMetricData.defaultExpectation.params) {
			mmPutMetricData.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPutMetricData.defaultExpectation.params)
		}
	}

	return mmPutMetricData
}

// Inspect accepts an inspector function that has same arguments as the CloudWatch.PutMetricData
func (mmPutMetricData *mCloudWatchMockPutMetricData) Inspect(f func(input *cloudwatch.PutMetricDataInput)) *mCloudWatchMockPutMetricData {
	if mmPutMetricData.mock.inspectFuncPutMetricData != nil {
		mmPutMetricData.mock.t.Fatalf("Inspect function is already set for CloudWatchMock.PutMetricData")
	}

	mmPutMetricData.mock.inspectFuncPutMetricData = f

	return mmPutMetricData
}

// Return sets up results that will be returned by CloudWatch.PutMetricData
func (mmPutMetricData *mCloudWatchMockPutMetricData) Return(pp1 *cloudwatch.PutMetricDataOutput, err error) *CloudWatchMock {
	if mmPutMetricData.mock.funcPutMetricData != nil {
		mmPutMetricData.mock.t.Fatalf("CloudWatchMock.PutMetricData mock is already set by Set")
	}

	if mmPutMetricData.defaultExpectation == nil {
		mmPutMetricData.defaultExpectation = &CloudWatchMockPutMetricDataExpectation{mock: mmPutMetricData.mock}
	}
	mmPutMetricData.defaultExpectation.results = &CloudWatchMockPutMetricDataResults{pp1, err}
	return mmPutMetricData.mock
}

//Set uses given function f to mock the CloudWatch.PutMetricData method
func (mmPutMetricData *mCloudWatchMockPutMetricData) Set(f func(input *cloudwatch.PutMetricDataInput) (pp1 *cloudwatch.PutMetricDataOutput, err error)) *CloudWatchMock {
	if mmPutMetricData.defaultExpectation != nil {
		mmPutMetricData.mock.t.Fatalf("Default expectation is already set for the CloudWatch.PutMetricData method")
	}

	if len(mmPutMetricData.expectations) > 0 {
		mmPutMetricData.mock.t.Fatalf("Some expectations are already set for the CloudWatch.PutMetricData method")
	}

	mmPutMetricData.mock.funcPutMetricData = f
	return mmPutMetricData.mock
}

// When sets expectation for the CloudWatch.PutMetricData which will trigger the result defined by the following
// Then helper
func (mmPutMetricData *mCloudWatchMockPutMetricData) When(input *cloudwatch.PutMetricDataInput) *CloudWatchMockPutMetricDataExpectation {
	if mmPutMetricData.mock.funcPutMetricData != nil {
		mmPutMetricData.mock.t.Fatalf("CloudWatchMock.PutMetricData mock is already set by Set")
	}

	expectation := &CloudWatchMockPutMetricDataExpectation{
		mock:   mmPutMetricData.mock,
		params: &CloudWatchMockPutMetricDataParams{input},
	}
	mmPutMetricData.expectations = append(mmPutMetricData.expectations, expectation)
	return expectation
}

// Then sets up CloudWatch.PutMetricData return parameters for the expectation previously defined by the When method
func (e *CloudWatchMockPutMetricDataExpectation) Then(pp1 *cloudwatch.PutMetricDataOutput, err error) *CloudWatchMock {
	e.results = &CloudWatchMockPutMetricDataResults{pp1, err}
	return e.mock
}

// PutMetricData implements cloudmetrics.CloudWatch
func (mmPutMetricData *CloudWatchMock) PutMetricData(input *cloudwatch.PutMetricDataInput) (pp1 *cloudwatch.PutMetricDataOutput, err error) {
	mm_atomic.AddUint64(&mmPutMetricData.beforePutMetricDataCounter, 1)
	defer mm_atomic.AddUint64(&mmPutMetricData.afterPutMetricDataCounter, 1)

	if mmPutMetricData.inspectFuncPutMetricData != nil {
		mmPutMetricData.inspectFuncPutMetricData(input)
	}

	mm_params := &CloudWatchMockPutMetricDataParams{input}

	// Record call args
	mmPutMetricData.PutMetricDataMock.mutex.Lock()
	mmPutMetricData.PutMetricDataMock.callArgs = append(mmPutMetricData.PutMetricDataMock.callArgs, mm_params)
	mmPutMetricData.PutMetricDataMock.mutex.Unlock()

	for _, e := range mmPutMetricData.PutMetricDataMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp1, e.results.err
		}
	}

	if mmPutMetricData.PutMetricDataMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPutMetricData.PutMetricDataMock.defaultExpectation.Counter, 1)
		mm_want := mmPutMetricData.PutMetricDataMock.defaultExpectation.params
		mm_got := CloudWatchMockPutMetricDataParams{input}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPutMetricData.t.Errorf("CloudWatchMock.PutMetricData got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPutMetricData.PutMetricDataMock.defaultExpectation.results
		if mm_results == nil {
			mmPutMetricData.t.Fatal("No results are set for the CloudWatchMock.PutMetricData")
		}
		return (*mm_results).pp1, (*mm_results).err
	}
	if mmPutMetricData.funcPutMetricData != nil {
		return mmPutMetricData.funcPutMetricData(input)
	}
	mmPutMetricData.t.Fatalf("Unexpected call to CloudWatchMock.PutMetricData. %v", input)
	return
}

// PutMetricDataAfterCounter returns a count of finished CloudWatchMock.PutMetricData invocations
func (mmPutMetricData *CloudWatchMock) PutMetricDataAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPutMetricData.afterPutMetricDataCounter)
}

// PutMetricDataBeforeCounter returns a count of CloudWatchMock.PutMetricData invocations
func (mmPutMetricData *CloudWatchMock) PutMetricDataBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPutMetricData.beforePutMetricDataCounter)
}

// Calls returns a list of arguments used in each call to CloudWatchMock.PutMetricData.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPutMetricData *mCloudWatchMockPutMetricData) Calls() []*CloudWatchMockPutMetricDataParams {
	mmPutMetricData.mutex.RLock()

	argCopy := make([]*CloudWatchMockPutMetricDataParams, len(mmPutMetricData.callArgs))
	copy(argCopy, mmPutMetricData.callArgs)

	mmPutMetricData.mutex.RUnlock()

	return argCopy
}

// MinimockPutMetricDataDone returns true if the count of the PutMetricData invocations corresponds
// the number of defined expectations
func (m *CloudWatchMock) MinimockPutMetricDataDone() bool {
	for _, e := range m.PutMetricDataMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PutMetricDataMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPutMetricDataCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPutMetricData != nil && mm_atomic.LoadUint64(&m.afterPutMetricDataCounter) < 1 {
		return false
	}
	return true
}

// MinimockPutMetricDataInspect logs each unmet expectation
func (m *CloudWatchMock) MinimockPutMetricDataInspect() {
	for _, e := range m.PutMetricDataMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CloudWatchMock.PutMetricData with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PutMetricDataMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPutMetricDataCounter) < 1 {
		if m.PutMetricDataMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CloudWatchMock.PutMetricData")
		} else {
			m.t.Errorf("Expected call to CloudWatchMock.PutMetricData with params: %#v", *m.PutMetricDataMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPutMetricData != nil && mm_atomic.LoadUint64(&m.afterPutMetricDataCounter) < 1 {
		m.t.Error("Expected call to CloudWatchMock.PutMetricData")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CloudWatchMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockPutMetricDataInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CloudWatchMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CloudWatchMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockPutMetricDataDone()
}
