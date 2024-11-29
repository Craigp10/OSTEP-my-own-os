package process

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type STATE string

// Process states following the xv6 Proc Structure
const (
	UNUSED   STATE = "UNUSED"
	EMBRYO   STATE = "EMBRYO"
	SLEEPING STATE = "SLEEPING"
	BLOCKED  STATE = "BLOCKED"
	RUNNABLE STATE = "RUNNABLE"
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

func getNextPID() uuid.UUID {
	return uuid.New()
}

// Process structure following the xv6 Proc Structure
type Process struct {
	PID    uuid.UUID
	State  STATE
	Parent *Process
	Child  *Process
	// killed int // ?
	// chan -- not sure?
	// file File
	// cwd *inode -- hmm may not need for my impl
	// context Context // context switch criteria
	Size uint // size of process memory -- may not need
	// trapFrame Trap -- Trap for the current interupt
	AddressSpace []byte // Memory space allotted for the process...
	instructions []byte // For now... This may be a func that just executes a timer...
}

// Fork spawns a new process from the current process' state ... Almost feel liek fork should be higher level, not on a process
func (p *Process) fork() *Process {
	child := &Process{
		PID:          getNextPID(), // Function to generate unique PIDs
		State:        EMBRYO,
		Parent:       p,
		AddressSpace: append([]byte{}, p.AddressSpace...), // Copy memory
		instructions: append([]byte{}, p.instructions...), // Copy instructions
	}
	child.State = RUNNABLE
	return child
}

// Wait is called by a parent process to wait until a child process completes.
func (p *Process) Wait() {
	if p.Child == nil {
		return
	}

	// Only works if child is spun up in separate routine...
	for p.Child.State != ZOMBIE {
		time.Sleep(500 * time.Millisecond)
	}

	p.Parent.RemoveChild()
}

// func (p *Process) Exec()

func (p *Process) RemoveChild() {
	p.Child = nil
}

func (p *Process) CreateChild() error {
	if p.Child != nil {
		return fmt.Errorf("process already has child process")
	}

	child := p.fork()
	p.Child = child

	return nil
}
