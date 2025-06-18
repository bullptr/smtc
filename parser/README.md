# Benchmarks

The following benchmark results compare the performance of [TreeSitter](https://tree-sitter.github.io/tree-sitter/) and [Antlr4](https://github.com/antlr4-go/antlr) SMTC parser implementations.

Environment:

```bash
goos: linux
goarch: amd64
pkg: github.com/smtc/benchmarks
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx
```

To run the benchmarks, use:

```bash
make bench
```

## Performance Analysis: TreeSitter vs Antlr

Based on the benchmark results, TreeSitter generally performs better than Antlr across both small and large inputs. Here's a detailed breakdown:

### Results

| Parser Implementation | Operations/sec | Average Time | Memory | Allocations |
| --------------------- | -------------- | ------------ | ------ | ----------- |
| TreeSitter (Small)    | 13,664         | 86.4μs       | 1.1KB  | 11          |
| Antlr (Small)         | 12,742         | 96.2μs       | 18.9KB | 206         |
| TreeSitter (Large)    | 85             | 13.5ms       | 55KB   | 11          |
| Antlr (Large)         | 81             | 14.3ms       | 3.2MB  | 40,446      |

### Small Input Performance

- **Execution Speed:** TreeSitter is about 10% faster (86,426 ns/op vs 96,222 ns/op)
- **Memory Usage:** TreeSitter uses significantly less memory (1,088 B/op vs 18,900 B/op)
- **Allocation Efficiency:** TreeSitter makes fewer allocations (11 vs 206 allocs/op)

### Large Input Performance

- **Execution Speed:** TreeSitter is about 5.6% faster (13.5ms vs 14.3ms)
- **Memory Usage:** TreeSitter is vastly more memory-efficient (55KB vs 3.2MB)
- **Allocation Efficiency:** TreeSitter maintains consistent allocations (11 vs 40,446 allocs/op)

### Conclusion

TreeSitter is the better performer overall, showing:

- Better speed performance in both small and large inputs
- Significantly lower memory consumption
- More consistent and efficient memory allocation patterns
- Better scalability with larger inputs

The most striking advantage is TreeSitter's memory efficiency, particularly with large inputs where it uses only 1.7% of the memory that Antlr requires.

### Raw Output

```bash
goos: linux
goarch: amd64
pkg: github.com/smtc/benchmarks
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkParsers
BenchmarkParsers/TreeSitter_Small
BenchmarkParsers/TreeSitter_Small-8                13664             86426 ns/op            1088 B/op         11 allocs/op
BenchmarkParsers/Antlr_Small
BenchmarkParsers/Antlr_Small-8                     12742             96222 ns/op           18900 B/op        206 allocs/op
BenchmarkParsers/TreeSitter_Large
BenchmarkParsers/TreeSitter_Large-8                   85          13539014 ns/op           55008 B/op         11 allocs/op
BenchmarkParsers/Antlr_Large
BenchmarkParsers/Antlr_Large-8                        81          14341820 ns/op         3245091 B/op      40446 allocs/op
PASS
ok      github.com/smtc/benchmarks      4.781s
```