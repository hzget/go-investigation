package main

import (
	"testing"
)

func BenchmarkIncVisitCount(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//IncVisitCountWithMutex()
			IncVisitCount()
		}
	})
}
