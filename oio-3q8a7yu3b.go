// Package main contains a simple timer implementation in Go.
package main

import (
	"fmt"
	"time"
)

// Timer represents a simple timer.
type Timer struct {
	startTime time.Time
}

// NewTimer returns a new timer instance.
func NewTimer() *Timer {
	return &Timer{
		startTime: time.Now(),
	}
}

// Start starts the timer.
func (t *Timer) Start() {
	t.startTime = time.Now()
}

// Stop stops the timer and returns the elapsed time.
func (t *Timer) Stop() time.Duration {
	return time.Since(t.startTime)
}

// Example usage:
func main() {
	timer := NewTimer()
	fmt.Println("Timer started. Waiting for 5 seconds...")
	time.Sleep(5 * time.Second)
	elapsed := timer.Stop()
	fmt.Printf("Timer stopped. Elapsed time: %v\n", elapsed)
}