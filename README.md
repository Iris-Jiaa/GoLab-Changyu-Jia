# Go-Lab
**Name**: Changyu Jia <br>
**Student ID**: C00292876 <br>

## Overview
The examples in this repository showcase different approaches to handling concurrency in Go, including barriers, mutexes, semaphores, atomic operations, and channel-based synchronization.

## Go Code
### [barrier_exercise](barrier_exercise/barrier.go)
- Barrier implementation using mutexes and semaphores for coordinating goroutines.
### [barrier_reuseable](barrier_reuseable/barrier2.go)
- Reusable barrier implementation using atomic operations and channels.
### [mutex](./mutex/mutex.go)
- A simple concurrent counter using mutexes and wait groups for synchronization.
### [rendezvous](./rendezvous/rendezvous.go)
- Rendezvous pattern implementation using channels and mutexes for synchronization.
### [dinPhil](./dinPhil_disruptCycle/dinPhil.go)
- Dining Philosophers problem: Philosophers alternate between thinking and eating while sharing limited forks, with a semaphore used to manage resource access.
### [gol](./go_gol/gol.go)
- Conway's Game of Life simulation using coroutines for parallel processing.

## Getting Started

### Prerequisites
- Go 1.18 or later

### Running the Examples
```bash
# Run any code
go run barrier.go
go run barrier2.go
go run mutex.go
go run rendezvous.go
go run dinPhil.go
go run gol.go
