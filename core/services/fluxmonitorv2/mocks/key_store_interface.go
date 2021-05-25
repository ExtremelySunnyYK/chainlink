// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"
)

// KeyStoreInterface is an autogenerated mock type for the KeyStoreInterface type
type KeyStoreInterface struct {
	mock.Mock
}

// GetRoundRobinAddress provides a mock function with given fields: _a0
func (_m *KeyStoreInterface) GetRoundRobinAddress(_a0 ...common.Address) (common.Address, error) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(...common.Address) common.Address); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...common.Address) error); ok {
		r1 = rf(_a0...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendingKeys provides a mock function with given fields:
func (_m *KeyStoreInterface) SendingKeys() ([]models.Key, error) {
	ret := _m.Called()

	var r0 []models.Key
	if rf, ok := ret.Get(0).(func() []models.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
