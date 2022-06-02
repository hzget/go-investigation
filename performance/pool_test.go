package lperformance

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"io"
	"runtime"
	"sync"
	"testing"
)

/*
 The Pool has its performance cost. Using the Pool is
 much slower than the initialization for simple data.

BenchmarkSimpleDataWithPool    	190975378	         6.287 ns/op	       0 B/op	       0 allocs/op
BenchmarkSimpleDataWithoutPool 	1000000000	         0.7004 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkSimpleDataWithPool(b *testing.B) {
	var p sync.Pool
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})
}

func BenchmarkSimpleDataWithoutPool(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := 0
			i = i
		}
	})
}

/*
 Items in the pool will be cleared after 2 GC circles.
 Subsequent Get() or Put() method will trigger pool to
 recreate objects. Thus the performance will get down.

BenchmarkPoolNoGC              	  799978	      1486 ns/op	      82 B/op	       0 allocs/op
BenchmarkPoolWithGC            	     105	  11929488 ns/op	66357720 B/op	     114 allocs/op
*/
func BenchmarkPoolNoGC(b *testing.B) {
	b.ReportAllocs()
	var a, z [100]*flate.Writer
	p := sync.Pool{New: func() interface{} { return &flate.Writer{} }}
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(a); j++ {
			a[j] = p.Get().(*flate.Writer)
		}
		for j := 0; j < len(a); j++ {
			p.Put(a[j])
		}
		a = z
		// runtime.GC()
	}
}

func BenchmarkPoolWithGC(b *testing.B) {
	b.ReportAllocs()
	var a, z [100]*flate.Writer
	p := sync.Pool{New: func() interface{} { return &flate.Writer{} }}
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(a); j++ {
			a[j] = p.Get().(*flate.Writer)
		}
		for j := 0; j < len(a); j++ {
			p.Put(a[j])
		}
		a = z
		runtime.GC()
		runtime.GC()
	}
}

/*
 If an object is expensive in the creation and
 you have to create it quite often,
 the overhead of using the Pool is much smaller than
 the benefits we have in exchange.

BenchmarkWriteGzipWithPool     	 1321502	       847.6 ns/op	      56 B/op	       2 allocs/op
BenchmarkWriteGzipWithoutPool  	   20839	     55449 ns/op	  807017 B/op	      16 allocs/op
*/
func BenchmarkWriteGzipWithPool(b *testing.B) {
	writerGzipPool := sync.Pool{
		New: func() interface{} {
			return gzip.NewWriter(io.Discard)
		},
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		data := bytes.NewReader([]byte("aohayo."))
		writer := writerGzipPool.Get().(*gzip.Writer)
		writer.Flush()
		io.Copy(writer, data)
		writerGzipPool.Put(writer)
	}
}

func BenchmarkWriteGzipWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		data := bytes.NewReader([]byte("aohayo."))
		writer := gzip.NewWriter(io.Discard)
		io.Copy(writer, data)
	}
}
