# Policies

<b> Policies are the algorithms that give the OS behavior and make decisions. These are the scheduling algorithms for now. </b>

Two key algorithms are included in this design

1. Mutli Level Feedback Queue (MLFQ)
1. Lottery Schedulings

Policy design structure

A design written in Golang

### Mutli Level Feedback Queue (MLFQ)

<b> Parameters </b>

1. Time Slice - time given single execution window
1. Allotment - Time given before priority is re-assesed
1. Queue Count - Number of queues used

#### Rules

**Rule 1:** If Priority(A) > Priority(B), A runs (B doesn’t).

**Rule 2:** If Priority(A) = Priority(B), A & B run in round-robin fashion using the time slice (quantum length) of the given queue.

**Rule 3:** When a job enters the system, it is placed at the highest priority (the topmost queue)

**Rule 4:** Once a job uses up its time allotment at a given level (re-gardless of how many times it has given up the CPU), its priority is reduced (i.e., it moves down one queue).

**Rule 5:** After some time period S, move all the jobs in the system to the topmost queue.

#### Q's to clarify

1. How does execution begin... Does the head of each queue get a core? Does queue B only run while there are no jobs in queue A?

### Lottery Scheduling

#### Rules
