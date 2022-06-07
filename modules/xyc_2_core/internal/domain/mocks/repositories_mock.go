// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/interfaces/repositories.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/interfaces"
)

// MockContactDTO is a mock of ContactDTO interface.
type MockContactDTO struct {
	ctrl     *gomock.Controller
	recorder *MockContactDTOMockRecorder
}

// MockContactDTOMockRecorder is the mock recorder for MockContactDTO.
type MockContactDTOMockRecorder struct {
	mock *MockContactDTO
}

// NewMockContactDTO creates a new mock instance.
func NewMockContactDTO(ctrl *gomock.Controller) *MockContactDTO {
	mock := &MockContactDTO{ctrl: ctrl}
	mock.recorder = &MockContactDTOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContactDTO) EXPECT() *MockContactDTOMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockContactDTO) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockContactDTOMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockContactDTO)(nil).Name))
}

// Phone mocks base method.
func (m *MockContactDTO) Phone() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Phone")
	ret0, _ := ret[0].(string)
	return ret0
}

// Phone indicates an expected call of Phone.
func (mr *MockContactDTOMockRecorder) Phone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Phone", reflect.TypeOf((*MockContactDTO)(nil).Phone))
}

// MockContactsRepository is a mock of ContactsRepository interface.
type MockContactsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockContactsRepositoryMockRecorder
}

// MockContactsRepositoryMockRecorder is the mock recorder for MockContactsRepository.
type MockContactsRepositoryMockRecorder struct {
	mock *MockContactsRepository
}

// NewMockContactsRepository creates a new mock instance.
func NewMockContactsRepository(ctrl *gomock.Controller) *MockContactsRepository {
	mock := &MockContactsRepository{ctrl: ctrl}
	mock.recorder = &MockContactsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContactsRepository) EXPECT() *MockContactsRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockContactsRepository) Add(arg0 interfaces.ContactDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockContactsRepositoryMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockContactsRepository)(nil).Add), arg0)
}

// AddAll mocks base method.
func (m *MockContactsRepository) AddAll(arg0 interfaces.ContactDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAll indicates an expected call of AddAll.
func (mr *MockContactsRepositoryMockRecorder) AddAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAll", reflect.TypeOf((*MockContactsRepository)(nil).AddAll), arg0)
}

// GetByPhone mocks base method.
func (m *MockContactsRepository) GetByPhone(arg0 string) (interfaces.ContactDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPhone", arg0)
	ret0, _ := ret[0].(interfaces.ContactDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPhone indicates an expected call of GetByPhone.
func (mr *MockContactsRepositoryMockRecorder) GetByPhone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPhone", reflect.TypeOf((*MockContactsRepository)(nil).GetByPhone), arg0)
}
