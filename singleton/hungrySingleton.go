// Package singleton DESC
package singleton

// Singleton 饿汉式单例
type Singleton struct{}

//singleton 单例
var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}
