package lperformance

import (
	"bytes"
	"compress/gzip"
	"io"
	"sync"
	"testing"
)

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
