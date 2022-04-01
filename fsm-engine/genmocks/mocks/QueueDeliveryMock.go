// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// QueueDeliveryMock is an autogenerated mock type for the QueueDeliveryMock type
type QueueDeliveryMock struct {
	mock.Mock
}

// Ack provides a mock function with given fields: ctx
func (_m *QueueDeliveryMock) Ack(ctx context.Context) {
	_m.Called(ctx)
}

// GetBody provides a mock function with given fields:
func (_m *QueueDeliveryMock) GetBody() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Reject provides a mock function with given fields: ctx
func (_m *QueueDeliveryMock) Reject(ctx context.Context) {
	_m.Called(ctx)
}
