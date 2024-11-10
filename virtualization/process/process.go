package process

type STATE string

// Process states following the xv6 Proc Structure
const (
	UNUSED   STATE = "UNUSED"
	EMBRYO   STATE = "EMBRYO"
	SLEEPING STATE = "SLEEPING"
	BLOCKED  STATE = "BLOCKED"
	RUNNABLE STATE = "RUNNABLE" // may rename to PENDING
	ZOMBIE   STATE = "ZOMBIE"
	RUNNING  STATE = "RUNNING"
)

// Context structure following the xv6 Proc
type Context struct {
	eip int
	esp int
	ebx int
	ecx int
	edx int
	esi int
	edi int
	ebp int
}

// Process structure following the xv6 Proc Structure
type Process struct {
	PID    int
	State  STATE
	Parent *Process
	killed int
	// chan -- not sure?
	// file File
	// cwd *inode -- hmm may not need for my impl
	context Context // context switch criteria
	Size    uint    // size of process memory -- may not need
	// trapFrame Trap -- Trap for the current interupt
	AddressSpace []byte // Thoguht we needed?
	instructions []byte // thought we needed?
}

func (p *Process) Fork()
func (p *Process) Wait()
func (p *Process) Exec()
