package lifecycle_test

import (
	"fmt"

	"github.com/xgfone/go-tools/lifecycle"
)

func ExampleLifeCycleManager() {
	lcm := lifecycle.NewLifeCycleManager()

	in := make(chan interface{})
	out := make(chan interface{})
	all := make(chan interface{})

	fmt.Println("Register the clean functions of apps")
	lcm.Register(func() {
		fmt.Println("Clean and wait app1 to end")
	}).Register(func() {
		fmt.Println("Clean and wait app2 to end")
	}).RegisterChannel(in, out).RegisterChannel(all, all)

	fmt.Println("Apps do something ...")
	go func() {
		<-in // Block until the program exits, that's, calling the method lcm.Stop()
		fmt.Println("Clean and wait app3 to end")
		out <- true // Inform the main goruntine that the app has cleaned and ended.
	}()

	go func() {
		<-all // Block until the program exits, that's, calling the method lcm.Stop()
		fmt.Println("Clean and wait app4 to end")
		all <- true // Inform the main goruntine that the app has cleaned and ended.
	}()

	fmt.Println("The program is ready to exit ...")
	lcm.Stop() // Apps clean
	fmt.Println("The program exited")

	// Output:
	// Register the clean functions of apps
	// Apps do something ...
	// The program is ready to exit ...
	// Clean and wait app1 to end
	// Clean and wait app2 to end
	// Clean and wait app3 to end
	// Clean and wait app4 to end
	// The program exited
}
