package lifecycle

var defaultManager *LifeCycleManager

func init() {
	defaultManager = NewLifeCycleManager()
}

// Register registers the argument to the global default one.
func Register(f func()) *LifeCycleManager {
	return defaultManager.Register(f)
}

// RegisterChannel registers the argument to the global default one.
func RegisterChannel(in chan<- interface{}, out <-chan interface{}) *LifeCycleManager {
	return defaultManager.RegisterChannel(in, out)
}

// Stop stops the global default one.
func Stop() {
	defaultManager.Stop()
}
