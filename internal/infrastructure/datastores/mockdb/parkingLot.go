// Code generated by MockGen. DO NOT EDIT.
// Source: parkwise/internal/domain/definition (interfaces: ParkingLotRepository)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination internal/infrastructure/datastores/mockdb/parkingLot.go parkwise/internal/domain/definition ParkingLotRepository
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	bo "parkwise/internal/domain/bo"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockParkingLotRepository is a mock of ParkingLotRepository interface.
type MockParkingLotRepository struct {
	ctrl     *gomock.Controller
	recorder *MockParkingLotRepositoryMockRecorder
}

// MockParkingLotRepositoryMockRecorder is the mock recorder for MockParkingLotRepository.
type MockParkingLotRepositoryMockRecorder struct {
	mock *MockParkingLotRepository
}

// NewMockParkingLotRepository creates a new mock instance.
func NewMockParkingLotRepository(ctrl *gomock.Controller) *MockParkingLotRepository {
	mock := &MockParkingLotRepository{ctrl: ctrl}
	mock.recorder = &MockParkingLotRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParkingLotRepository) EXPECT() *MockParkingLotRepositoryMockRecorder {
	return m.recorder
}

// CreateParkingLot mocks base method.
func (m *MockParkingLotRepository) CreateParkingLot(arg0 context.Context, arg1 *bo.ParkingLot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateParkingLot", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateParkingLot indicates an expected call of CreateParkingLot.
func (mr *MockParkingLotRepositoryMockRecorder) CreateParkingLot(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateParkingLot", reflect.TypeOf((*MockParkingLotRepository)(nil).CreateParkingLot), arg0, arg1)
}

// DailyReport mocks base method.
func (m *MockParkingLotRepository) DailyReport(arg0 context.Context, arg1 string) (bo.DailyReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DailyReport", arg0, arg1)
	ret0, _ := ret[0].(bo.DailyReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DailyReport indicates an expected call of DailyReport.
func (mr *MockParkingLotRepositoryMockRecorder) DailyReport(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyReport", reflect.TypeOf((*MockParkingLotRepository)(nil).DailyReport), arg0, arg1)
}

// Maintenance mocks base method.
func (m *MockParkingLotRepository) Maintenance(arg0 context.Context, arg1 bo.Maintenance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Maintenance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Maintenance indicates an expected call of Maintenance.
func (mr *MockParkingLotRepositoryMockRecorder) Maintenance(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Maintenance", reflect.TypeOf((*MockParkingLotRepository)(nil).Maintenance), arg0, arg1)
}

// ParkingLotStatus mocks base method.
func (m *MockParkingLotRepository) ParkingLotStatus(arg0 context.Context, arg1 int64) (bo.ParkingLotStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParkingLotStatus", arg0, arg1)
	ret0, _ := ret[0].(bo.ParkingLotStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParkingLotStatus indicates an expected call of ParkingLotStatus.
func (mr *MockParkingLotRepositoryMockRecorder) ParkingLotStatus(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParkingLotStatus", reflect.TypeOf((*MockParkingLotRepository)(nil).ParkingLotStatus), arg0, arg1)
}
