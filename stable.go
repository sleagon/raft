package raft

// StableStore is used to provide stable storage
// of key configurations to ensure safety.
// StableStore提供的是稳定数据的存储，需要保证稳定存储以后可以稳定取出。
type StableStore interface {
	Set(key []byte, val []byte) error

	// Get returns the value for key, or an empty byte slice if key was not found.
	Get(key []byte) ([]byte, error)

	SetUint64(key []byte, val uint64) error

	// GetUint64 returns the uint64 value for key, or 0 if key was not found.
	GetUint64(key []byte) (uint64, error)
}
