// Package main is a small utility program to parse and print benchmark results in a table format.
package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"

	"github.com/JavierZunzunegui/zerrors/internal/benchmark"
	"golang.org/x/tools/benchmark/parse"
)

func main() {
	if err := runMain(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runMain() error {
	set, err := parse.ParseSet(os.Stdin)
	if err != nil {
		return err
	}

	const depthPrefix = "depth_"
	re := regexp.MustCompile(depthPrefix + "[0-9]+")

	rawBenchmarks := make(map[string]map[int]*parse.Benchmark)

	for name, bench := range set {
		if len(bench) != 1 {
			return fmt.Errorf("len(benchmark %s): expecting 1, got %d", name, len(bench))
		}

		v := re.FindString(path.Base(bench[0].Name))
		if len(v) < len(depthPrefix) || v[:len(depthPrefix)] != depthPrefix {
			return fmt.Errorf("unrecognised benchmark prefix(1): expecting %s prefix, got %s", depthPrefix, v)
		}
		depth, err := strconv.Atoi(v[len(depthPrefix):])
		if err != nil {
			return fmt.Errorf("unrecognised benchmark format: expecting an integer, got %s", v[len(depthPrefix):])
		}

		b := path.Dir(bench[0].Name)
		d := rawBenchmarks[b]
		if d == nil {
			d = make(map[int]*parse.Benchmark)
			rawBenchmarks[b] = d
		}
		d[depth] = bench[0]
	}

	benchmarks := make([]benchSet, 0)

	for k, v := range rawBenchmarks {
		if len(v) != benchmark.ScenarioCount {
			return fmt.Errorf("benchmark %s: expecting %d benchmark depths, got %d", k, benchmark.ScenarioCount, len(v))
		}

		var bs [benchmark.ScenarioCount]*parse.Benchmark
		for i, depth := range benchmark.Scenarios() {
			bench, ok := v[depth]
			if !ok {
				return fmt.Errorf("benchmark %s: expecting a benchmark for %d depth", k, depth)
			}
			bs[i] = bench
		}

		benchmarks = append(benchmarks, benchSet{k, bs})
	}

	sort.Slice(benchmarks, func(i, j int) bool {
		return benchmarks[i].name < benchmarks[j].name
	})

	fmt.Println("# ns/op")
	fmt.Println("\n", table(benchmarks, "ns/op", nsPerOp))

	fmt.Println("\n", "# B/op")
	fmt.Println("\n", table(benchmarks, "B/op", allocedBytesPerOp))

	fmt.Println("\n", "# allocs/op")
	fmt.Println("\n", table(benchmarks, "allocs/op", allocsPerOp))

	return nil
}

type benchSet struct {
	name       string
	benchmarks [benchmark.ScenarioCount]*parse.Benchmark
}

func nsPerOp(b *parse.Benchmark) string {
	return fmt.Sprintf("%.2f", b.NsPerOp)
}

func allocedBytesPerOp(b *parse.Benchmark) string {
	return fmt.Sprintf("%d", b.AllocedBytesPerOp)
}

func allocsPerOp(b *parse.Benchmark) string {
	return fmt.Sprintf(" %d", b.AllocsPerOp)
}

func table(benchmarks []benchSet, unit string, f func(*parse.Benchmark) string) string {
	var buf bytes.Buffer

	buf.WriteString("Benchmark (" + unit + ")")
	for i, v := range benchmark.Scenarios() {
		buf.WriteString(" | ")
		if i == 0 {
			buf.WriteString("depth=")
		}
		buf.WriteString(strconv.Itoa(v))
	}

	buf.WriteString("\n---")
	for range benchmark.Scenarios() {
		buf.WriteString(" | ---")
	}

	for _, bench := range benchmarks {
		buf.WriteString("\n")
		buf.WriteString(bench.name)
		for _, v := range bench.benchmarks {
			buf.WriteString(" | ")
			buf.WriteString(f(v))
		}
	}

	return buf.String()
}
