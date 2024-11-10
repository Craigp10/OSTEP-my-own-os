package virtualization

type process_switch_behavior string
type io_finished_behavior string

var (

	// process switch behavior
	SCHED_SWITCH_ON_IO  process_switch_behavior = "SWITCH_ON_IO"
	SCHED_SWITCH_ON_END process_switch_behavior = "SWITCH_ON_END"

	//1 io finished behavior
	IO_RUN_LATER     = "IO_RUN_LATER"
	IO_RUN_IMMEDIATE = "IO_RUN_IMMEDIATE"

	// process states
	STATE_RUNNING = "RUNNING"
	STATE_READY   = "READY"
	STATE_DONE    = "DONE"
	STATE_WAIT    = "WAITING"

	// members of process structure
	PROC_CODE  = "code_"
	PROC_PC    = "pc_"
	PROC_ID    = "pid_"
	PROC_STATE = "proc_state_"

	// things a process can do
	DO_COMPUTE       = "cpu"
	DO_IO            = "io"
	DO_PROGRAMMED_IO = "p_io"
)

type Scheduler struct {
	proc_info               map[string]string
	process_switch_behavior process_switch_behavior
	io_done_behavior        string //io_done_behavior
	io_length               int    // Limit to allow for a I/O blocking operation to run
	interrupt_overhead      func() error
}
