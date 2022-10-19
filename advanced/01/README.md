# Definitions

A goroutine is a lightweight thread managed by the Go runtime. Goroutines run in the same address space, so access to shared memory must be synchronized.

A race condition or race hazard is the condition of an electronics, software, or other system where the system's substantive behavior is dependent on the sequence or timing of other uncontrollable events.

Golang has in-build mechanism to detect race conditions:

    go run -race main.go

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

# Exercises:

**1.go**

1. The output of the program.
2. What happens if we run the program with a race detector?
3. How can we serialize access to variable `i`?

**2.go**

- ordering is not guaranteed
- the loop is not guaranteed to finish

Many compilers (at compile time) and CPU processors (at run time) often make some optimizations by adjusting the instruction orders, so that the instruction execution orders may differ from the orders presented in code. Instruction ordering is also often called memory ordering.

Surely, instruction reordering can't be arbitrary. The basic requirement for a reordering inside a specified goroutine is the reordering must not be detectable by the goroutine itself if the goroutine doesn't share data with other goroutines. In other words, from the perspective of such a goroutine, it can think its instruction execution order is always the same as the order specified by code, even if instruction reordering really happens inside it.

Source: https://go101.org/article/memory-model.html

# Homework

1. create a pipeline

You have an array of integers as input [5, 7, 8]. To each number in the array you need to apply two operations: (1) raise in power of 2 and then (2) multiply by 2

i.e. [5, 7, 8] -> [25, 49, 64] -> [50, 98, 128]

Do this processing in two ways:
* use arrays as input parameters and operate on the array
* speed up the execution by creating a pipeline. (if you are stuck see file `4.go` and use channels)

2. create a program that has a data race.
