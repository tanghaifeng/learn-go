package systemtype

type counter int

// New 一个简单的工厂方法
func New(count int) counter {
	return counter(count)
}
