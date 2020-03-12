module github.com/JavierZunzunegui/zerrors/internal/benchmark

go 1.13

// this go.mod just exists to not force the import of some third party libraries just because of some benchmarks.
replace github.com/JavierZunzunegui/zerrors => ../../

require (
	github.com/JavierZunzunegui/zerrors v0.1.0
	github.com/pkg/errors v0.9.1
	golang.org/x/tools v0.0.0-20200304193943-95d2e580d8eb
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
)
