# To run `calling_fork.c` call:
clang calling_fork.c -o calling_fork
./calling_fork

notice the different pids, and which path the pids are calling
The forked process, runs on the same function from the part of init.
The exection of these are not deterministic, the child may be prioritized by
the scheduler and complete first.
