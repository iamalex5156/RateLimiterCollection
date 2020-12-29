package rlc

import (
	"testing"
	"time"
)

// go test -bench=. -benchmem
func BenchmarkRLCShortAskStep(b *testing.B) {
	offset := NewStepOffset(96, 1)
	eq := NewRLC(96, 16, offset)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eq.Ask()
	}
}

func BenchmarkRLCShortAskRandom(b *testing.B) {
	offset := NewRandomOffset(96)
	eq := NewRLC(96, 16, offset)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eq.Ask()
	}
}

func BenchmarkRLCShortNotify(b *testing.B) {
	offset := NewStepOffset(96, 1)
	eq := NewRLC(96, 16, offset)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eq.Notify(false, 1)
	}
}

func BenchmarkRLCLongAskStep(b *testing.B) {
	offset := NewStepOffset(1048576, 1)
	eq := NewRLC(1048576, 16, offset)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eq.Ask()
	}
}

func BenchmarkRLCLongAskRandom(b *testing.B) {
	offset := NewRandomOffset(1048576)
	eq := NewRLC(1048576, 16, offset)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eq.Ask()
	}
}

func BenchmarkRLCLongNotify(b *testing.B) {
	offset := NewStepOffset(1048576, 1)
	eq := NewRLC(1048576, 16, offset)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eq.Notify(false, 1)
	}
}

func BenchmarkSliderShortWindow(b *testing.B) {
	slider := NewSlider(time.Millisecond*100, time.Millisecond*10, 32)
	for i := 0; i < b.N; i++ {
		slider.Ask()
	}
}

func BenchmarkSliderLongerWindow(b *testing.B) {
	slider := NewSlider(time.Second, time.Millisecond*100, 32)
	for i := 0; i < b.N; i++ {
		slider.Ask()
	}
}

func BenchmarkTokenBucketDenseRefill(b *testing.B) {
	tokenBucket := NewTokenBucket(32, time.Millisecond*10)
	for i := 0; i < b.N; i++ {
		tokenBucket.Ask()
	}
}

func BenchmarkTokenBucketSparseRefill(b *testing.B) {
	tokenBucket := NewTokenBucket(32, time.Second)
	for i := 0; i < b.N; i++ {
		tokenBucket.Ask()
	}
}
