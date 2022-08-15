// Package main DESC
package main

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockCrawler struct {
	mock.Mock
}

// 实现GetUserList()方法时，需要调用Mock.Called()方法，传入参数（示例中无参数）
func (m *MockCrawler) GetUserList() ([]*User, error) {
	// Called()会返回一个mock.Arguments对象，该对象中保存着返回的值。它提供了对基本类型和error的获取方法Int()/String()/Bool()/Error()，
	// 和通用的获取方法Get()，通用方法返回interface{}，需要类型断言为具体类型，它们都接受一个表示索引的参数。
	args := m.Called()
	return args.Get(0).([]*User), args.Error(1)
}

var MockUsers = []*User{}

func init() {
	MockUsers = append(MockUsers, &User{"dj", 18})
	MockUsers = append(MockUsers, &User{"zs", 34})
}

func TestGetUserList(t *testing.T) {
	crawler := new(MockCrawler)
	crawler.On("GetUserList").Return(MockUsers, nil)
	GetAndPrintUsers(crawler)
	crawler.AssertExpectations(t)
}
