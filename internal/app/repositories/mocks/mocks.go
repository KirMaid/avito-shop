// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -source=interfaces.go -destination=./mocks/mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "avitoshop/internal/app/entities"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockInventoryRepository is a mock of InventoryRepository interface.
type MockInventoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryRepositoryMockRecorder
	isgomock struct{}
}

// MockInventoryRepositoryMockRecorder is the mock recorder for MockInventoryRepository.
type MockInventoryRepositoryMockRecorder struct {
	mock *MockInventoryRepository
}

// NewMockInventoryRepository creates a new mock instance.
func NewMockInventoryRepository(ctrl *gomock.Controller) *MockInventoryRepository {
	mock := &MockInventoryRepository{ctrl: ctrl}
	mock.recorder = &MockInventoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryRepository) EXPECT() *MockInventoryRepositoryMockRecorder {
	return m.recorder
}

// GetByUser mocks base method.
func (m *MockInventoryRepository) GetByUser(ctx context.Context, userID int) ([]entities.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUser", ctx, userID)
	ret0, _ := ret[0].([]entities.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUser indicates an expected call of GetByUser.
func (mr *MockInventoryRepositoryMockRecorder) GetByUser(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUser", reflect.TypeOf((*MockInventoryRepository)(nil).GetByUser), ctx, userID)
}

// InsertOrUpdate mocks base method.
func (m *MockInventoryRepository) InsertOrUpdate(ctx context.Context, inventory *entities.Inventory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrUpdate", ctx, inventory)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOrUpdate indicates an expected call of InsertOrUpdate.
func (mr *MockInventoryRepositoryMockRecorder) InsertOrUpdate(ctx, inventory any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrUpdate", reflect.TypeOf((*MockInventoryRepository)(nil).InsertOrUpdate), ctx, inventory)
}

// MockRedisInventoryRepository is a mock of RedisInventoryRepository interface.
type MockRedisInventoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisInventoryRepositoryMockRecorder
	isgomock struct{}
}

// MockRedisInventoryRepositoryMockRecorder is the mock recorder for MockRedisInventoryRepository.
type MockRedisInventoryRepositoryMockRecorder struct {
	mock *MockRedisInventoryRepository
}

// NewMockRedisInventoryRepository creates a new mock instance.
func NewMockRedisInventoryRepository(ctrl *gomock.Controller) *MockRedisInventoryRepository {
	mock := &MockRedisInventoryRepository{ctrl: ctrl}
	mock.recorder = &MockRedisInventoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisInventoryRepository) EXPECT() *MockRedisInventoryRepositoryMockRecorder {
	return m.recorder
}

// DeleteByUser mocks base method.
func (m *MockRedisInventoryRepository) DeleteByUser(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUser", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByUser indicates an expected call of DeleteByUser.
func (mr *MockRedisInventoryRepositoryMockRecorder) DeleteByUser(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUser", reflect.TypeOf((*MockRedisInventoryRepository)(nil).DeleteByUser), ctx, userID)
}

// GetByUser mocks base method.
func (m *MockRedisInventoryRepository) GetByUser(ctx context.Context, userID int) ([]entities.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUser", ctx, userID)
	ret0, _ := ret[0].([]entities.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUser indicates an expected call of GetByUser.
func (mr *MockRedisInventoryRepositoryMockRecorder) GetByUser(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUser", reflect.TypeOf((*MockRedisInventoryRepository)(nil).GetByUser), ctx, userID)
}

// SetByUser mocks base method.
func (m *MockRedisInventoryRepository) SetByUser(ctx context.Context, userID int, inventory []entities.Inventory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetByUser", ctx, userID, inventory)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetByUser indicates an expected call of SetByUser.
func (mr *MockRedisInventoryRepositoryMockRecorder) SetByUser(ctx, userID, inventory any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetByUser", reflect.TypeOf((*MockRedisInventoryRepository)(nil).SetByUser), ctx, userID, inventory)
}

// MockGoodRepository is a mock of GoodRepository interface.
type MockGoodRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGoodRepositoryMockRecorder
	isgomock struct{}
}

// MockGoodRepositoryMockRecorder is the mock recorder for MockGoodRepository.
type MockGoodRepositoryMockRecorder struct {
	mock *MockGoodRepository
}

// NewMockGoodRepository creates a new mock instance.
func NewMockGoodRepository(ctrl *gomock.Controller) *MockGoodRepository {
	mock := &MockGoodRepository{ctrl: ctrl}
	mock.recorder = &MockGoodRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGoodRepository) EXPECT() *MockGoodRepositoryMockRecorder {
	return m.recorder
}

// GetByName mocks base method.
func (m *MockGoodRepository) GetByName(ctx context.Context, name string) (*entities.Good, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(*entities.Good)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockGoodRepositoryMockRecorder) GetByName(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockGoodRepository)(nil).GetByName), ctx, name)
}

// GetList mocks base method.
func (m *MockGoodRepository) GetList(ctx context.Context) ([]entities.Good, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx)
	ret0, _ := ret[0].([]entities.Good)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockGoodRepositoryMockRecorder) GetList(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockGoodRepository)(nil).GetList), ctx)
}

// MockRedisGoodRepository is a mock of RedisGoodRepository interface.
type MockRedisGoodRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisGoodRepositoryMockRecorder
	isgomock struct{}
}

// MockRedisGoodRepositoryMockRecorder is the mock recorder for MockRedisGoodRepository.
type MockRedisGoodRepositoryMockRecorder struct {
	mock *MockRedisGoodRepository
}

// NewMockRedisGoodRepository creates a new mock instance.
func NewMockRedisGoodRepository(ctrl *gomock.Controller) *MockRedisGoodRepository {
	mock := &MockRedisGoodRepository{ctrl: ctrl}
	mock.recorder = &MockRedisGoodRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisGoodRepository) EXPECT() *MockRedisGoodRepositoryMockRecorder {
	return m.recorder
}

// GetByName mocks base method.
func (m *MockRedisGoodRepository) GetByName(ctx context.Context, name string) (*entities.Good, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(*entities.Good)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockRedisGoodRepositoryMockRecorder) GetByName(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockRedisGoodRepository)(nil).GetByName), ctx, name)
}

// SetByName mocks base method.
func (m *MockRedisGoodRepository) SetByName(ctx context.Context, name string, good *entities.Good) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetByName", ctx, name, good)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetByName indicates an expected call of SetByName.
func (mr *MockRedisGoodRepositoryMockRecorder) SetByName(ctx, name, good any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetByName", reflect.TypeOf((*MockRedisGoodRepository)(nil).SetByName), ctx, name, good)
}

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
	isgomock struct{}
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// GetReceivedTransactions mocks base method.
func (m *MockTransactionRepository) GetReceivedTransactions(ctx context.Context, userID int) ([]entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceivedTransactions", ctx, userID)
	ret0, _ := ret[0].([]entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceivedTransactions indicates an expected call of GetReceivedTransactions.
func (mr *MockTransactionRepositoryMockRecorder) GetReceivedTransactions(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceivedTransactions", reflect.TypeOf((*MockTransactionRepository)(nil).GetReceivedTransactions), ctx, userID)
}

// GetSentTransactions mocks base method.
func (m *MockTransactionRepository) GetSentTransactions(ctx context.Context, userID int) ([]entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSentTransactions", ctx, userID)
	ret0, _ := ret[0].([]entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSentTransactions indicates an expected call of GetSentTransactions.
func (mr *MockTransactionRepositoryMockRecorder) GetSentTransactions(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSentTransactions", reflect.TypeOf((*MockTransactionRepository)(nil).GetSentTransactions), ctx, userID)
}

// Insert mocks base method.
func (m *MockTransactionRepository) Insert(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, transaction)
	ret0, _ := ret[0].(*entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockTransactionRepositoryMockRecorder) Insert(ctx, transaction any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTransactionRepository)(nil).Insert), ctx, transaction)
}

// MockRedisTransactionRepository is a mock of RedisTransactionRepository interface.
type MockRedisTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisTransactionRepositoryMockRecorder
	isgomock struct{}
}

// MockRedisTransactionRepositoryMockRecorder is the mock recorder for MockRedisTransactionRepository.
type MockRedisTransactionRepositoryMockRecorder struct {
	mock *MockRedisTransactionRepository
}

// NewMockRedisTransactionRepository creates a new mock instance.
func NewMockRedisTransactionRepository(ctrl *gomock.Controller) *MockRedisTransactionRepository {
	mock := &MockRedisTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockRedisTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisTransactionRepository) EXPECT() *MockRedisTransactionRepositoryMockRecorder {
	return m.recorder
}

// DeleteReceivedTransactions mocks base method.
func (m *MockRedisTransactionRepository) DeleteReceivedTransactions(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReceivedTransactions", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReceivedTransactions indicates an expected call of DeleteReceivedTransactions.
func (mr *MockRedisTransactionRepositoryMockRecorder) DeleteReceivedTransactions(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReceivedTransactions", reflect.TypeOf((*MockRedisTransactionRepository)(nil).DeleteReceivedTransactions), ctx, userID)
}

// DeleteSentTransactions mocks base method.
func (m *MockRedisTransactionRepository) DeleteSentTransactions(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSentTransactions", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSentTransactions indicates an expected call of DeleteSentTransactions.
func (mr *MockRedisTransactionRepositoryMockRecorder) DeleteSentTransactions(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSentTransactions", reflect.TypeOf((*MockRedisTransactionRepository)(nil).DeleteSentTransactions), ctx, userID)
}

// GetReceivedTransactions mocks base method.
func (m *MockRedisTransactionRepository) GetReceivedTransactions(ctx context.Context, userID int) ([]entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceivedTransactions", ctx, userID)
	ret0, _ := ret[0].([]entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceivedTransactions indicates an expected call of GetReceivedTransactions.
func (mr *MockRedisTransactionRepositoryMockRecorder) GetReceivedTransactions(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceivedTransactions", reflect.TypeOf((*MockRedisTransactionRepository)(nil).GetReceivedTransactions), ctx, userID)
}

// GetSentTransactions mocks base method.
func (m *MockRedisTransactionRepository) GetSentTransactions(ctx context.Context, userID int) ([]entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSentTransactions", ctx, userID)
	ret0, _ := ret[0].([]entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSentTransactions indicates an expected call of GetSentTransactions.
func (mr *MockRedisTransactionRepositoryMockRecorder) GetSentTransactions(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSentTransactions", reflect.TypeOf((*MockRedisTransactionRepository)(nil).GetSentTransactions), ctx, userID)
}

// SetReceivedTransactions mocks base method.
func (m *MockRedisTransactionRepository) SetReceivedTransactions(ctx context.Context, userID int, transactions []entities.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReceivedTransactions", ctx, userID, transactions)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReceivedTransactions indicates an expected call of SetReceivedTransactions.
func (mr *MockRedisTransactionRepositoryMockRecorder) SetReceivedTransactions(ctx, userID, transactions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReceivedTransactions", reflect.TypeOf((*MockRedisTransactionRepository)(nil).SetReceivedTransactions), ctx, userID, transactions)
}

// SetSentTransactions mocks base method.
func (m *MockRedisTransactionRepository) SetSentTransactions(ctx context.Context, userID int, transactions []entities.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSentTransactions", ctx, userID, transactions)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSentTransactions indicates an expected call of SetSentTransactions.
func (mr *MockRedisTransactionRepositoryMockRecorder) SetSentTransactions(ctx, userID, transactions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSentTransactions", reflect.TypeOf((*MockRedisTransactionRepository)(nil).SetSentTransactions), ctx, userID, transactions)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockUserRepository) GetByID(ctx context.Context, userID int) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, userID)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserRepositoryMockRecorder) GetByID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepository)(nil).GetByID), ctx, userID)
}

// GetByIDs mocks base method.
func (m *MockUserRepository) GetByIDs(ctx context.Context, userIDs []int) ([]entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, userIDs)
	ret0, _ := ret[0].([]entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockUserRepositoryMockRecorder) GetByIDs(ctx, userIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockUserRepository)(nil).GetByIDs), ctx, userIDs)
}

// GetByUsername mocks base method.
func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", ctx, username)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockUserRepositoryMockRecorder) GetByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockUserRepository)(nil).GetByUsername), ctx, username)
}

// Insert mocks base method.
func (m *MockUserRepository) Insert(ctx context.Context, user *entities.User) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, user)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockUserRepositoryMockRecorder) Insert(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserRepository)(nil).Insert), ctx, user)
}

// UpdateBalance mocks base method.
func (m *MockUserRepository) UpdateBalance(ctx context.Context, userID, balance int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalance", ctx, userID, balance)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockUserRepositoryMockRecorder) UpdateBalance(ctx, userID, balance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockUserRepository)(nil).UpdateBalance), ctx, userID, balance)
}

// MockRedisUserRepository is a mock of RedisUserRepository interface.
type MockRedisUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisUserRepositoryMockRecorder
	isgomock struct{}
}

// MockRedisUserRepositoryMockRecorder is the mock recorder for MockRedisUserRepository.
type MockRedisUserRepositoryMockRecorder struct {
	mock *MockRedisUserRepository
}

// NewMockRedisUserRepository creates a new mock instance.
func NewMockRedisUserRepository(ctrl *gomock.Controller) *MockRedisUserRepository {
	mock := &MockRedisUserRepository{ctrl: ctrl}
	mock.recorder = &MockRedisUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisUserRepository) EXPECT() *MockRedisUserRepositoryMockRecorder {
	return m.recorder
}

// GetByIDs mocks base method.
func (m *MockRedisUserRepository) GetByIDs(ctx context.Context, userIDs []int) ([]entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, userIDs)
	ret0, _ := ret[0].([]entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockRedisUserRepositoryMockRecorder) GetByIDs(ctx, userIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockRedisUserRepository)(nil).GetByIDs), ctx, userIDs)
}

// GetByUsername mocks base method.
func (m *MockRedisUserRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", ctx, username)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockRedisUserRepositoryMockRecorder) GetByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockRedisUserRepository)(nil).GetByUsername), ctx, username)
}

// GetUsernamesByIDs mocks base method.
func (m *MockRedisUserRepository) GetUsernamesByIDs(ctx context.Context, userIDs []int) (map[int]string, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsernamesByIDs", ctx, userIDs)
	ret0, _ := ret[0].(map[int]string)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsernamesByIDs indicates an expected call of GetUsernamesByIDs.
func (mr *MockRedisUserRepositoryMockRecorder) GetUsernamesByIDs(ctx, userIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsernamesByIDs", reflect.TypeOf((*MockRedisUserRepository)(nil).GetUsernamesByIDs), ctx, userIDs)
}

// SetByUsername mocks base method.
func (m *MockRedisUserRepository) SetByUsername(ctx context.Context, username string, user *entities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetByUsername", ctx, username, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetByUsername indicates an expected call of SetByUsername.
func (mr *MockRedisUserRepositoryMockRecorder) SetByUsername(ctx, username, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetByUsername", reflect.TypeOf((*MockRedisUserRepository)(nil).SetByUsername), ctx, username, user)
}
