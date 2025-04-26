# Algorithm Analysis Report

 To generate a report install dependencies by running `go mod tidy` then run `go run main.go`.

 If you wish to compile a binary executable you can do so by running the `go build main.go` command.

 Doing so would generate a comprarion report for Merge, Heap, Selection, Quick sort algorithms and a comparison for Binary and Exponential seach algorithms. Generated algorithm can be viewed in [Report](/algorithm_analysis_report.md) in markdown format.


## Performance Comparison

| Algorithm | Time Complexity | N = 1,000 | N = 10,000 | N = 100,000 | N = 1,000,000 |
|-----------|----------------|-----------|------------|-------------|-------------|
| Merge Sort | O(n log n) | 57.542µs | 704.041µs | 8.377875ms | 83.2615ms |
| Heap Sort | O(n log n) | 45.042µs | 591.916µs | 7.116291ms | 91.607708ms |
| Selection Sort | O(n²) | 298.417µs | 25.904791ms | 2.531950459s | 4m21.222088375s |
| Quick Sort | O(n log n) | 162.5µs | 1.604959ms | 9.650334ms | 68.892167ms |
| Binary Search | O(log n) | 83ns | 42ns | 42ns | 1.75µs |
| Exponential Search | O(log n) | 41ns | 0s | 333ns | 42ns |

## Hardware Information

- Go Version: go1.23.2
- OS: darwin
- Architecture: arm64
- CPU Name: Apple M3 Pro
- CPU Physical Cores: 11
- CPU Logical Cores: 11
- CPU Threads per Core: 1
- CPU L1 Data Cache: 65536 bytes
- CPU L1 Instruction Cache: 131072 bytes
- CPU L2 Cache: 4194304 bytes
- CPU L3 Cache: -1 bytes

## Experiment Setup

- Random arrays were generated with values between 0 and 10*N
- For search algorithms, the target value was -1 (not in the array)
- Selection Sort was skipped for sizes > 100,000 due to O(n²) time complexity
- All algorithms were tested on the same randomly generated arrays
- Note that completion time for some runs might show 0 because of the run time taking less than the minimum time that can be measured.
