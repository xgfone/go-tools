package lifecycle

var defaultManager *Manager

func init() {
	defaultManager = NewManager()
}

// Register registers the argument to the global default one.
func Register(f func()) *Manager {
	return defaultManager.Register(f)
}

// RegisterChannel registers the argument to the global default one.
func RegisterChannel(in chan<- interface{}, out <-chan interface{}) *Manager {
	return defaultManager.RegisterChannel(in, out)
}

// Stop stops the global default one.
func Stop() {
	defaultManager.Stop()
}

// GetDefaultManager returns the default global Manager.
func GetDefaultManager() *Manager {
	return defaultManager
}

// IsStop returns true if the default global manager, or false.
func IsStop() bool {
	return defaultManager.IsStop()
}
