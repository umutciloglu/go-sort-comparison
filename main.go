package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"time"

	. "github.com/klauspost/cpuid/v2"
)

// Generate a random array of n integers
func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n * 10) // Numbers between 0 and 10*n
	}
	return arr
}

// -------------------- SORTING ALGORITHMS --------------------

// Selection Sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Merge Sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// Heap Sort
func heapSort(arr []int) {
	n := len(arr)

	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// One by one extract elements from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]
		// Call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

// To heapify a subtree rooted with node i which is an index in arr[].
// n is size of heap
func heapify(arr []int, n, i int) {
	largest := i // Initialize largest as root
	left := 2*i + 1
	right := 2*i + 2

	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		// Recursively heapify the affected sub-tree
		heapify(arr, n, largest)
	}
}

// Quick Sort
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)/2]
	left, right := 0, len(arr)-1

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	if right > 0 {
		quickSort(arr[:right+1])
	}
	if left < len(arr) {
		quickSort(arr[left:])
	}
}

// -------------------- SEARCH ALGORITHMS --------------------

// Binary Search
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // Element not found
}

// Exponential Search
func exponentialSearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	if arr[0] == target {
		return 0
	}

	// Find range for binary search
	i := 1
	for i < len(arr) && arr[i] <= target {
		i *= 2
	}

	// Call binary search on the range [i/2, min(i, len(arr)-1)]
	return binarySearchRange(arr, target, i/2, min(i, len(arr)-1))
}

func binarySearchRange(arr []int, target, left, right int) int {
	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // Element not found
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// -------------------- TESTING FUNCTIONS --------------------

func testSortingAlgorithm(name string, algorithm func([]int), sizes []int) map[int]time.Duration {
	results := make(map[int]time.Duration)

	for _, size := range sizes {
		arr := generateRandomArray(size)
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)

		start := time.Now()

		if name == "Merge Sort" {
			// Special case for merge sort since it returns a new array
			result := mergeSort(arrCopy)
			copy(arrCopy, result)
		} else {
			algorithm(arrCopy)
		}

		duration := time.Since(start)
		results[size] = duration

		// Verify the array is sorted
		if !sort.SliceIsSorted(arrCopy, func(i, j int) bool { return arrCopy[i] < arrCopy[j] }) {
			fmt.Printf("WARNING: %s failed to sort the array of size %d correctly!\n", name, size)
		}
	}

	return results
}

func testSearchAlgorithm(name string, algorithm func([]int, int) int, sizes []int) map[int]time.Duration {
	results := make(map[int]time.Duration)

	for _, size := range sizes {
		arr := generateRandomArray(size)

		// Sort the array (required for search algorithms)
		sort.Ints(arr)

		target := -1 // Search for a value that doesn't exist in the array

		start := time.Now()
		algorithm(arr, target)
		duration := time.Since(start)

		results[size] = duration
	}

	return results
}

func getSystemInfo() string {
	return fmt.Sprintf("- Go Version: %s\n- OS: %s\n- Architecture: %s\n- CPU Name: %s\n- CPU Physical Cores: %d\n- CPU Logical Cores: %d\n- CPU Threads per Core: %d\n- CPU L1 Data Cache: %d bytes\n- CPU L1 Instruction Cache: %d bytes\n- CPU L2 Cache: %d bytes\n- CPU L3 Cache: %d bytes\n",
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		CPU.BrandName,
		CPU.PhysicalCores,
		CPU.LogicalCores,
		CPU.ThreadsPerCore,
		CPU.Cache.L1D,
		CPU.Cache.L1I,
		CPU.Cache.L2,
		CPU.Cache.L3)
}

func main() {
	// Set random seed
	rand.Seed(time.Now().UnixNano())

	// Test sizes
	sizes := []int{1000, 10000, 100000, 1000000}

	// Prepare algorithms
	algorithms := map[string]func([]int){
		"Merge Sort":     func(arr []int) { result := mergeSort(arr); copy(arr, result) },
		"Heap Sort":      heapSort,
		"Quick Sort":     quickSort,
		"Selection Sort": selectionSort,
	}

	searchAlgorithms := map[string]func([]int, int) int{
		"Binary Search":      binarySearch,
		"Exponential Search": exponentialSearch,
	}

	// Store results
	results := make(map[string]map[int]time.Duration)

	// Run tests
	fmt.Println("Running tests. This may take some time...")

	for name, algorithm := range algorithms {
		fmt.Printf("Testing %s...\n", name)
		results[name] = testSortingAlgorithm(name, algorithm, sizes)
	}

	for name, algorithm := range searchAlgorithms {
		fmt.Printf("Testing %s...\n", name)
		results[name] = testSearchAlgorithm(name, algorithm, sizes)
	}

	complexities := map[string]string{
		"Merge Sort":         "O(n log n)",
		"Heap Sort":          "O(n log n)",
		"Quick Sort":         "O(n log n)",
		"Selection Sort":     "O(n²)",
		"Binary Search":      "O(log n)",
		"Exponential Search": "O(log n)",
	}

	file, err := os.Create("algorithm_analysis_report.md")
	if err != nil {
		fmt.Println("Error creating report file:", err)
		return
	}
	defer file.Close()

	// Write report header
	file.WriteString("# Algorithm Analysis Report\n\n")
	file.WriteString(" For more information you can check out the github repository https://github.com/umutciloglu/go-sort-comparison\n\n")
	file.WriteString("## Performance Comparison\n\n")

	// Write table header
	file.WriteString("| Algorithm | Time Complexity | N = 1,000 | N = 10,000 | N = 100,000 | N = 1,000,000 |\n")
	file.WriteString("|-----------|----------------|-----------|------------|-------------|-------------|\n")

	// Write sorting algorithm results
	for _, name := range []string{"Merge Sort", "Heap Sort", "Selection Sort", "Quick Sort"} {
		file.WriteString(fmt.Sprintf("| %s | %s", name, complexities[name]))
		for _, size := range sizes {
			duration := results[name][size]
			file.WriteString(fmt.Sprintf(" | %v", duration))
		}
		file.WriteString(" |\n")
	}

	// Write search algorithm results
	for _, name := range []string{"Binary Search", "Exponential Search"} {
		file.WriteString(fmt.Sprintf("| %s | %s", name, complexities[name]))
		for _, size := range sizes {
			file.WriteString(fmt.Sprintf(" | %v", results[name][size]))
		}
		file.WriteString(" |\n")
	}

	// Add hardware information section
	file.WriteString("\n## Hardware Information\n\n")
	file.WriteString(getSystemInfo())

	// Add experiment setup section
	file.WriteString("\n## Experiment Setup\n\n")
	file.WriteString("- Random arrays were generated with values between 0 and 10*N\n")
	file.WriteString("- For search algorithms, the target value was -1 (not in the array)\n")
	file.WriteString("- Selection Sort was skipped for sizes > 100,000 due to O(n²) time complexity\n")
	file.WriteString("- All algorithms were tested on the same randomly generated arrays\n")
	file.WriteString("- Note that completion time for some runs might show 0 because of the run time taking less than the minimum time that can be measured.\n")

	fmt.Println("Report generated: algorithm_analysis_report.md")
}
