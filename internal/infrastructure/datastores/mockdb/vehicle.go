// Code generated by MockGen. DO NOT EDIT.
// Source: parkwise/internal/domain/definition (interfaces: VehicleRepository)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination internal/infrastructure/datastores/mockdb/vehicle.go parkwise/internal/domain/definition VehicleRepository
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	bo "parkwise/internal/domain/bo"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockVehicleRepository is a mock of VehicleRepository interface.
type MockVehicleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVehicleRepositoryMockRecorder
}

// MockVehicleRepositoryMockRecorder is the mock recorder for MockVehicleRepository.
type MockVehicleRepositoryMockRecorder struct {
	mock *MockVehicleRepository
}

// NewMockVehicleRepository creates a new mock instance.
func NewMockVehicleRepository(ctrl *gomock.Controller) *MockVehicleRepository {
	mock := &MockVehicleRepository{ctrl: ctrl}
	mock.recorder = &MockVehicleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVehicleRepository) EXPECT() *MockVehicleRepositoryMockRecorder {
	return m.recorder
}

// Park mocks base method.
func (m *MockVehicleRepository) Park(arg0 context.Context, arg1 bo.Park) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Park", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Park indicates an expected call of Park.
func (mr *MockVehicleRepositoryMockRecorder) Park(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Park", reflect.TypeOf((*MockVehicleRepository)(nil).Park), arg0, arg1)
}

// UnPark mocks base method.
func (m *MockVehicleRepository) UnPark(arg0 context.Context, arg1 bo.UnPark) (bo.Charge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnPark", arg0, arg1)
	ret0, _ := ret[0].(bo.Charge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnPark indicates an expected call of UnPark.
func (mr *MockVehicleRepositoryMockRecorder) UnPark(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnPark", reflect.TypeOf((*MockVehicleRepository)(nil).UnPark), arg0, arg1)
}
