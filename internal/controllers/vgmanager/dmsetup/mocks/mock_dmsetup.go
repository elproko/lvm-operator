// Code generated by mockery v2.36.0. DO NOT EDIT.

package dmsetup

import mock "github.com/stretchr/testify/mock"

// MockDmsetup is an autogenerated mock type for the Dmsetup type
type MockDmsetup struct {
	mock.Mock
}

type MockDmsetup_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDmsetup) EXPECT() *MockDmsetup_Expecter {
	return &MockDmsetup_Expecter{mock: &_m.Mock}
}

// Remove provides a mock function with given fields: deviceName
func (_m *MockDmsetup) Remove(deviceName string) error {
	ret := _m.Called(deviceName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(deviceName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDmsetup_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type MockDmsetup_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - deviceName string
func (_e *MockDmsetup_Expecter) Remove(deviceName interface{}) *MockDmsetup_Remove_Call {
	return &MockDmsetup_Remove_Call{Call: _e.mock.On("Remove", deviceName)}
}

func (_c *MockDmsetup_Remove_Call) Run(run func(deviceName string)) *MockDmsetup_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDmsetup_Remove_Call) Return(_a0 error) *MockDmsetup_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDmsetup_Remove_Call) RunAndReturn(run func(string) error) *MockDmsetup_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDmsetup creates a new instance of MockDmsetup. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDmsetup(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDmsetup {
	mock := &MockDmsetup{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
