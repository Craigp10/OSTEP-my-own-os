package virtualization

import "operatingsystem/virtualization/process"

type OperatingSystem interface {
	Create()
	Destory()
	Wait()
	Status()
	// Other control functions we could implement
	Suspend() // Pause a process for a while... This may take it out of memory and back to the disk w/o killing it.
}

// Hmmm not sure what to do with this one. We'll want this used within our processeses. some how.
type SystemCaller interface {
	Fork()
	Exec()
	Wait()

	// Others...
	// Pipe()
}

// ProcessControlBlock Will represent the C structure that contains info about each process.
// Knows which ones are running etc...
type ProcessControlBlock struct {
	// Rough design, likely to changes
	ActiveProcesses    []process.Process
	NonActiveProcesses []process.Process
}
