// Code generated by struct2interface; DO NOT EDIT.

package redis

// RedisInterface ...
//
type RedisInterface interface {
	// todo 采用 redis 的原子操作，这里目前是为了先实现功能
	GetNodeId() (int64, error)
}
