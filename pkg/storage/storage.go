package storage

// Storage defines persistence interface (for future extension).
type Storage interface {
	Save(data []byte) error
	Load() ([]byte, error)
}
