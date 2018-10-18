package concatenation

import (
	"testing"
)

var test10words []string
var test100words []string
var test1kWords []string
var test10kWords []string
var test100kWords []string

func init() {
	for i := 0; i < 100000; i++ {
		test100kWords = append(test100kWords, "word")
	}
	test10kWords = test100kWords[:10000]
	test1kWords = test10kWords[:1000]
	test100words = test1kWords[:100]
	test10words = test100words[:10]
}

func doBenchmark(b *testing.B, words []string, function func([]string) string) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		function(words)
	}
}

func Benchmark10StringConcat(b *testing.B)   { doBenchmark(b, test10words, Concat) }
func Benchmark10StringsJoin(b *testing.B)    { doBenchmark(b, test10words, Join) }
func Benchmark10StringsBuilder(b *testing.B) { doBenchmark(b, test10words, Builder) }
func Benchmark10StringBuffer(b *testing.B)   { doBenchmark(b, test10words, Buffer) }

func Benchmark100StringConcat(b *testing.B)   { doBenchmark(b, test100words, Concat) }
func Benchmark100StringsJoin(b *testing.B)    { doBenchmark(b, test100words, Join) }
func Benchmark100StringsBuilder(b *testing.B) { doBenchmark(b, test100words, Builder) }
func Benchmark100StringBuffer(b *testing.B)   { doBenchmark(b, test100words, Buffer) }

func Benchmark1000StringConcat(b *testing.B)   { doBenchmark(b, test1kWords, Concat) }
func Benchmark1000StringsJoin(b *testing.B)    { doBenchmark(b, test1kWords, Join) }
func Benchmark1000StringsBuilder(b *testing.B) { doBenchmark(b, test1kWords, Builder) }
func Benchmark1000StringBuffer(b *testing.B)   { doBenchmark(b, test1kWords, Buffer) }

func Benchmark10000StringConcat(b *testing.B)   { doBenchmark(b, test10kWords, Concat) }
func Benchmark10000StringsJoin(b *testing.B)    { doBenchmark(b, test10kWords, Join) }
func Benchmark10000StringsBuilder(b *testing.B) { doBenchmark(b, test10kWords, Builder) }
func Benchmark10000StringBuffer(b *testing.B)   { doBenchmark(b, test10kWords, Buffer) }

func Benchmark100000StringConcat(b *testing.B)   { doBenchmark(b, test100kWords, Concat) }
func Benchmark100000StringsJoin(b *testing.B)    { doBenchmark(b, test100kWords, Join) }
func Benchmark100000StringsBuilder(b *testing.B) { doBenchmark(b, test100kWords, Builder) }
func Benchmark100000StringBuffer(b *testing.B)   { doBenchmark(b, test100kWords, Buffer) }
