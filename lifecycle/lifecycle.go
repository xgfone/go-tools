// The manager of the lifecycle of some apps in a program.
package lifecycle

// LifeCycleManager manage the lifecycle of some apps in a program.
type LifeCycleManager struct {
	callbacks []func()
}

// NewLifeCycleManager returns a new LifeCycleManager.
func NewLifeCycleManager() *LifeCycleManager {
	return &LifeCycleManager{
		callbacks: make([]func(), 0, 8),
	}
}

// RegisterChannel is same as Register, but using the channel, not the callback.
//
// The parameter in is used to notice the app to end. And out is used to notice
// the manager that the app has cleaned and ended successfully.
func (self *LifeCycleManager) RegisterChannel(in chan<- interface{}, out <-chan interface{}) {
	self.Register(func() {
		in <- true
		<-out
	})
}

// Register registers a callback function for the app.
//
// When calling Stop(), the callback function will be called in turn
// by the order that they are registered.
func (self *LifeCycleManager) Register(f func()) *LifeCycleManager {
	self.callbacks = append(self.callbacks, f)
	return self
}

// Stop terminates and cleans all the apps.
//
// This method will be blocked until all the apps finish the clean.
// If the cleaning function of a certain app panics, ignore it and continue to
// call the cleaning function of the next app.
func (self LifeCycleManager) Stop() {
	for _, f := range self.callbacks {
		// f()
		callFuncAndIgnorePanic(f)
	}
}

func callFuncAndIgnorePanic(f func()) {
	defer func() {
		recover()
	}()
	f()
}
