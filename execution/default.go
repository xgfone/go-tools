package execution

import "sync"

var globalLock = new(sync.Mutex)

var (
	// OnceCMD is the commond to be execute once.
	// The default use the global lock, but you can set OnceCMD.IsLock to false
	// to cancel it.
	OnceCMD = new(Execution)

	// RedoCMD is the commond to be execute .
	// The default use the global lock, but you can set OnceCMD.IsLock to false
	// to cancel it.
	RedoCMD = new(Execution) // Repeat to execute until failure

	// RetryCMD is the commond to be execute once.
	// The default use the global lock, but you can set OnceCMD.IsLock to false
	// to cancel it.
	RetryCMD = new(Execution) // Repeat to execute until success
)
