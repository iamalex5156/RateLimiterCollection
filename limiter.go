package rlc

type token struct{}

// Limiter represents a rate limiter.
type Limiter interface {
	Ask() bool
	Take()
}
