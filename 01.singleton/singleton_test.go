// Package _1_singleton DESC https://github.com/stretchr/testify testify框架
package _1_singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	result := GetInstance()
	assert.ObjectsAreEqual(Singleton{}, result)
}
