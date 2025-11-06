package job

// Job represents a unit of work to be processed.
type Job struct {
	ID      string
	Handler func() error
}
