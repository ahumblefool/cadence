// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: transmissionTaskHandler.go
//
// Generated by this command:
//
//	mockgen -package domain -source transmissionTaskHandler.go -destination transmissionTaskHandler_mock.go
//

// Package domain is a generated GoMock package.
package domain

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	persistence "github.com/uber/cadence/common/persistence"
	types "github.com/uber/cadence/common/types"
)

// MockReplicator is a mock of Replicator interface.
type MockReplicator struct {
	ctrl     *gomock.Controller
	recorder *MockReplicatorMockRecorder
	isgomock struct{}
}

// MockReplicatorMockRecorder is the mock recorder for MockReplicator.
type MockReplicatorMockRecorder struct {
	mock *MockReplicator
}

// NewMockReplicator creates a new mock instance.
func NewMockReplicator(ctrl *gomock.Controller) *MockReplicator {
	mock := &MockReplicator{ctrl: ctrl}
	mock.recorder = &MockReplicatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicator) EXPECT() *MockReplicatorMockRecorder {
	return m.recorder
}

// HandleTransmissionTask mocks base method.
func (m *MockReplicator) HandleTransmissionTask(ctx context.Context, domainOperation types.DomainOperation, info *persistence.DomainInfo, config *persistence.DomainConfig, replicationConfig *persistence.DomainReplicationConfig, configVersion, failoverVersion, previousFailoverVersion int64, isGlobalDomainEnabled bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleTransmissionTask", ctx, domainOperation, info, config, replicationConfig, configVersion, failoverVersion, previousFailoverVersion, isGlobalDomainEnabled)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleTransmissionTask indicates an expected call of HandleTransmissionTask.
func (mr *MockReplicatorMockRecorder) HandleTransmissionTask(ctx, domainOperation, info, config, replicationConfig, configVersion, failoverVersion, previousFailoverVersion, isGlobalDomainEnabled any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTransmissionTask", reflect.TypeOf((*MockReplicator)(nil).HandleTransmissionTask), ctx, domainOperation, info, config, replicationConfig, configVersion, failoverVersion, previousFailoverVersion, isGlobalDomainEnabled)
}
