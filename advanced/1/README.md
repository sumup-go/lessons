# Definitions

A goroutine is a lightweight thread managed by the Go runtime. Goroutines run in the same address space, so access to shared memory must be synchronized.

A race condition or race hazard is the condition of an electronics, software, or other system where the system's substantive behavior is dependent on the sequence or timing of other uncontrollable events.

Golang has in-build mechanism to detect race conditions:

    go run -race main.go

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
