// Package _1_singleton DESC https://github.com/stretchr/testify testify框架
package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	result := GetInstance()
	assert.ObjectsAreEqual(Singleton{}, result)
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
