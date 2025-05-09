# Algorithm Analysis Report

 For more information you can check out the github repository https://github.com/umutciloglu/go-sort-comparison

## Performance Comparison

| Algorithm | Time Complexity | N = 1,000 | N = 10,000 | N = 100,000 | N = 1,000,000 |
|-----------|----------------|-----------|------------|-------------|-------------|
| Merge Sort | O(n log n) | 207.375µs | 2.208292ms | 20.631917ms | 92.091958ms |
| Heap Sort | O(n log n) | 42.666µs | 537.042µs | 7.077166ms | 91.263875ms |
| Selection Sort | O(n²) | 284.166µs | 25.334ms | 2.521276958s | 4m16.645789084s |
| Quick Sort | O(n log n) | 33.542µs | 389.041µs | 4.574917ms | 52.373625ms |
| Binary Search | O(log n) | 84ns | 0s | 0s | 0s |
| Exponential Search | O(log n) | 542ns | 42ns | 0s | 42ns |

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
