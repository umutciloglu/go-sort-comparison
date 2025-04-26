# Algorithm Analysis Report

## Performance Comparison

| Algorithm | Time Complexity | N = 1,000 | N = 10,000 | N = 100,000 | N = 1,000,000 |
|-----------|----------------|-----------|------------|-------------|-------------|
| Merge Sort | O(n log n) | 57.875µs | 821.25µs | 8.054292ms | 80.162625ms |
| Heap Sort | O(n log n) | 44.833µs | 576.584µs | 7.578625ms | 92.210125ms |
| Selection Sort | O(n²) | 281.833µs | 25.480625ms | 2.494347917s | 4m17.541171375s |
| Quick Sort | O(n log n) | 254.375µs | 1.337417ms | 10.274667ms | 70.553708ms |
| Binary Search | O(log n) | 41ns | 0s | 0s | 83ns |
| Exponential Search | O(log n) | 0s | 0s | 0s | 41ns |

## Hardware Information

- Go Version: go1.23.2
- OS: darwin
- Architecture: arm64
- CPUs: 11

## Experiment Setup

- Random arrays were generated with values between 0 and 10*N
- For search algorithms, the target value was -1 (not in the array)
- Selection Sort was skipped for sizes > 100,000 due to O(n²) time complexity
- All algorithms were tested on the same randomly generated arrays
- Note that completion time for some runs might show 0 because of the run time taking less than the minimum time that can be measured.
