## 代码实现
单例模式采用了 饿汉式 和 懒汉式 两种实现，个人其实更倾向于饿汉式的实现，简单，并且可以将问题及早暴露，懒汉式虽然支持延迟加载，但是这只是把冷启动时间
放到了第一次使用的时候，并没有本质上解决问题，并且为了实现懒汉式还不可避免的需要加锁。

### 饿汉式 
#### 代码实现
```go
package singleton

// Singleton 饿汉式单例
type Singleton struct {}
var singleton *Singleton

func init()  {
    singleton = &Singleton{}
}

func GetInstance() *Singleton  {
	return singleton
}
```
### 懒汉式
#### 代码实现

```go
package singleton

import "sync"

// Singleton 懒汉式单例
type Singleton struct{}

var (
	lazySingleton *Singleton
	once           = &sync.Once{}
)

func GetLazySingleton() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
            lazySingleton = &Singleton{}
		})
    }
	return lazySingleton
}
```
### 测试结果对比
并发环境下饿汉式性能要好一些的
```
goos: darwin
goarch: amd64
pkg: designPattern-go/singleton
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkGetInstance
BenchmarkGetInstance-12        	1000000000	         0.1539 ns/op
BenchmarkGetLazyInstance
BenchmarkGetLazyInstance-12    	1000000000	         0.6045 ns/op
PASS
ok  	designPattern-go/singleton	0.842s
```