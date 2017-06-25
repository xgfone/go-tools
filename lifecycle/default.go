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

// ResetDefaultManager resets the default global manager.
func ResetDefaultManager(m *Manager) {
	if m == nil {
		panic("The argument is nil")
	}
	defaultManager = m
}

// IsStop returns true if the default global manager, or false.
func IsStop() bool {
	return defaultManager.IsStop()
}

// RunForever calls the method RunForever of the default global manager.
func RunForever() {
	defaultManager.RunForever()
}

// Wait calls the method Wait of the default global manager.
func Wait() {
	defaultManager.Wait()
}
