package service_test

import (
	"context"
	"testing"
	"withpattern/mocks"
	"withpattern/service"

	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T) {
	repoMock := &mocks.Repository{}
	service := service.NewService(repoMock)
	repoMock.Mock.On("GetTodo", context.Background()).Return(nil, nil)
	_, err := service.GetTodo(context.Background())
	assert.NoError(t, err)
}
