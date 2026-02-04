package solver

import (
	"push-swap/internal/operations"
	"push-swap/internal/stack"
	"testing"
)

func TestNewSolver(t *testing.T) {
	input := []int{3, 2, 1}
	solver := NewSolver(input)
	
	if solver == nil {
		t.Fatal("NewSolver should not return nil")
	}
	
	if solver.stackA.Size() != 3 {
		t.Errorf("Expected stack A size 3, got %d", solver.stackA.Size())
	}
	
	if !solver.stackB.IsEmpty() {
		t.Error("Stack B should be empty initially")
	}
	
	if len(solver.operations) != 0 {
		t.Error("Operations should be empty initially")
	}
}

func TestSolveAlreadySorted(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	solver := NewSolver(input)
	
	ops := solver.Solve()
	
	if len(ops) != 0 {
		t.Errorf("Expected 0 operations for already sorted stack, got %d", len(ops))
	}
}

func TestSolveEmptyStack(t *testing.T) {
	input := []int{}
	solver := NewSolver(input)
	
	ops := solver.Solve()
	
	if len(ops) != 0 {
		t.Errorf("Expected 0 operations for empty stack, got %d", len(ops))
	}
}

func TestSolveSingleElement(t *testing.T) {
	input := []int{42}
	solver := NewSolver(input)
	
	ops := solver.Solve()
	
	if len(ops) != 0 {
		t.Errorf("Expected 0 operations for single element, got %d", len(ops))
	}
}

func TestSolveTwoElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int // expected number of operations
	}{
		{
			name:     "Already sorted",
			input:    []int{1, 2},
			expected: 0,
		},
		{
			name:     "Needs swap",
			input:    []int{2, 1},
			expected: 1,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := NewSolver(tt.input)
			ops := solver.Solve()
			
			if len(ops) != tt.expected {
				t.Errorf("Expected %d operations, got %d", tt.expected, len(ops))
			}
			
			// Verify the result is sorted
			result := validateSolution(tt.input, ops)
			if !result {
				t.Error("Solution should result in sorted stack")
			}
		})
	}
}

func TestSolveThreeElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"Case 3,2,1", []int{3, 2, 1}},
		{"Case 3,1,2", []int{3, 1, 2}},
		{"Case 2,3,1", []int{2, 3, 1}},
		{"Case 2,1,3", []int{2, 1, 3}},
		{"Case 1,3,2", []int{1, 3, 2}},
		{"Case 1,2,3", []int{1, 2, 3}}, // already sorted
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := NewSolver(tt.input)
			ops := solver.Solve()
			
			// Verify the result is sorted
			result := validateSolution(tt.input, ops)
			if !result {
				t.Errorf("Solution for %v should result in sorted stack", tt.input)
			}
			
			// For 3 elements, should not need more than 2 operations (except for already sorted)
			if !isSorted(tt.input) && len(ops) > 2 {
				t.Errorf("Solution for %v uses %d operations, expected <= 2", tt.input, len(ops))
			}
		})
	}
}

func TestSolveFiveElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"Reverse sorted", []int{5, 4, 3, 2, 1}},
		{"Random order", []int{3, 1, 4, 2, 5}},
		{"Mostly sorted", []int{1, 3, 2, 4, 5}},
		{"Four elements", []int{4, 3, 2, 1}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := NewSolver(tt.input)
			ops := solver.Solve()
			
			// Verify the result is sorted
			result := validateSolution(tt.input, ops)
			if !result {
				t.Errorf("Solution for %v should result in sorted stack", tt.input)
			}
			
			// For 5 elements, should be reasonably efficient
			if len(ops) > 12 {
				t.Errorf("Solution for %v uses %d operations, expected <= 12", tt.input, len(ops))
			}
		})
	}
}

func TestSolveSixElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"Case 2,1,3,6,5,8", []int{2, 1, 3, 6, 5, 8}},
		{"Reverse sorted", []int{6, 5, 4, 3, 2, 1}},
		{"Random order", []int{3, 1, 6, 2, 5, 4}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := NewSolver(tt.input)
			ops := solver.Solve()
			
			// Verify the result is sorted
			result := validateSolution(tt.input, ops)
			if !result {
				t.Errorf("Solution for %v should result in sorted stack", tt.input)
			}
			
			// For 6 elements, should use less than 9 operations
			if len(ops) >= 9 {
				t.Errorf("Solution for %v uses %d operations, expected < 9", tt.input, len(ops))
			}
		})
	}
}

func TestSolveLargeStack(t *testing.T) {
	// Test with a larger stack
	input := []int{10, 3, 7, 1, 9, 2, 8, 4, 6, 5}
	solver := NewSolver(input)
	ops := solver.Solve()
	
	// Verify the result is sorted
	result := validateSolution(input, ops)
	if !result {
		t.Error("Solution should result in sorted stack")
	}
	
	// Should complete in reasonable number of operations
	if len(ops) > 100 {
		t.Errorf("Solution uses %d operations, might be inefficient", len(ops))
	}
}

func TestFindMinPosition(t *testing.T) {
	solver := NewSolver([]int{3, 1, 4, 2, 5})
	
	minPos := solver.findMinPosition(solver.stackA)
	
	if minPos != 1 {
		t.Errorf("Expected min position 1, got %d", minPos)
	}
	
	// Test with empty stack
	emptyStack := stack.NewEmptyStack()
	solver.stackA = emptyStack
	minPos = solver.findMinPosition(solver.stackA)
	
	if minPos != -1 {
		t.Errorf("Expected -1 for empty stack, got %d", minPos)
	}
}

func TestFindMinValue(t *testing.T) {
	solver := NewSolver([]int{3, 1, 4, 2, 5})
	
	minVal := solver.findMinValue(solver.stackA)
	
	if minVal != 1 {
		t.Errorf("Expected min value 1, got %d", minVal)
	}
	
	// Test with single element
	solver = NewSolver([]int{42})
	minVal = solver.findMinValue(solver.stackA)
	
	if minVal != 42 {
		t.Errorf("Expected min value 42, got %d", minVal)
	}
	
	// Test with negative numbers
	solver = NewSolver([]int{3, -5, 4, -2, 1})
	minVal = solver.findMinValue(solver.stackA)
	
	if minVal != -5 {
		t.Errorf("Expected min value -5, got %d", minVal)
	}
}

func TestGetValueAt(t *testing.T) {
	solver := NewSolver([]int{10, 20, 30, 40, 50})
	
	// Test valid indices
	val := solver.getValueAt(solver.stackA, 0)
	if val != 10 {
		t.Errorf("Expected value at index 0 to be 10, got %d", val)
	}
	
	val = solver.getValueAt(solver.stackA, 4)
	if val != 50 {
		t.Errorf("Expected value at index 4 to be 50, got %d", val)
	}
	
	// Test out of bounds index
	val = solver.getValueAt(solver.stackA, 10)
	if val != 999999 {
		t.Errorf("Expected value at out of bounds index to be 999999, got %d", val)
	}
}

func TestFindMaxPosition(t *testing.T) {
	solver := NewSolver([]int{3, 1, 4, 2, 5})
	
	maxPos := solver.findMaxPosition(solver.stackA)
	
	if maxPos != 4 {
		t.Errorf("Expected max position 4, got %d", maxPos)
	}
	
	// Test with negative numbers
	solver = NewSolver([]int{-3, -1, -4, -2, -5})
	maxPos = solver.findMaxPosition(solver.stackA)
	
	if maxPos != 1 {
		t.Errorf("Expected max position 1 (value -1), got %d", maxPos)
	}
}

func TestMoveToTopOptimized(t *testing.T) {
	solver := NewSolver([]int{1, 2, 3, 4, 5})
	
	// Test moving element at position 2 to top
	solver.moveToTopOptimized(solver.stackA, 2, true)
	
	top, _ := solver.stackA.Top()
	if top != 3 {
		t.Errorf("Expected top element to be 3 after moveToTopOptimized, got %d", top)
	}
	
	// Test with position 0 (should do nothing)
	solver = NewSolver([]int{1, 2, 3, 4, 5})
	originalTop, _ := solver.stackA.Top()
	solver.moveToTopOptimized(solver.stackA, 0, true)
	newTop, _ := solver.stackA.Top()
	
	if originalTop != newTop {
		t.Error("moveToTopOptimized with position 0 should not change the stack")
	}
}

func TestMoveToTopOptimizedWithReverseRotate(t *testing.T) {
	solver := NewSolver([]int{1, 2, 3, 4, 5})
	
	// Test moving element at position 4 (last) to top - should use reverse rotate
	solver.moveToTopOptimized(solver.stackA, 4, true)
	
	top, _ := solver.stackA.Top()
	if top != 5 {
		t.Errorf("Expected top element to be 5 after moveToTopOptimized, got %d", top)
	}
}

func TestExecuteAndRecord(t *testing.T) {
	solver := NewSolver([]int{2, 1})
	
	initialOpsCount := len(solver.operations)
	solver.executeAndRecord(operations.SA)
	
	if len(solver.operations) != initialOpsCount+1 {
		t.Error("executeAndRecord should add operation to the list")
	}
	
	if solver.operations[0] != operations.SA {
		t.Errorf("Expected first operation to be SA, got %v", solver.operations[0])
	}
	
	// Check that the operation was actually executed
	top, _ := solver.stackA.Top()
	if top != 1 {
		t.Errorf("Expected top element to be 1 after SA, got %d", top)
	}
}

func TestCreateRanks(t *testing.T) {
	solver := NewSolver([]int{5, 1, 4, 2, 3})
	
	ranks := solver.createRanks()
	
	// Verify ranks map
	expected := map[int]int{
		1: 0,
		2: 1,
		3: 2,
		4: 3,
		5: 4,
	}
	
	for val, expectedRank := range expected {
		if ranks[val] != expectedRank {
			t.Errorf("Expected rank of %d to be %d, got %d", val, expectedRank, ranks[val])
		}
	}
	
	// Test with negative numbers
	solver = NewSolver([]int{-5, 10, -3, 0, 7})
	ranks = solver.createRanks()
	
	expectedNeg := map[int]int{
		-5: 0,
		-3: 1,
		0:  2,
		7:  3,
		10: 4,
	}
	
	for val, expectedRank := range expectedNeg {
		if ranks[val] != expectedRank {
			t.Errorf("Expected rank of %d to be %d, got %d", val, expectedRank, ranks[val])
		}
	}
}

func TestApplyRanks(t *testing.T) {
	solver := NewSolver([]int{5, 1, 4, 2, 3})
	
	ranks := map[int]int{
		1: 0,
		2: 1,
		3: 2,
		4: 3,
		5: 4,
	}
	
	solver.applyRanks(ranks)
	
	// Verify that values are replaced with their ranks
	expected := []int{4, 0, 3, 1, 2}
	actual := solver.stackA.ToSlice()
	
	for i, exp := range expected {
		if actual[i] != exp {
			t.Errorf("Expected stackA[%d] to be %d, got %d", i, exp, actual[i])
		}
	}
}

func TestQuickSortAndPartition(t *testing.T) {
	solver := NewSolver([]int{})
	
	// Test quicksort
	arr := []int{5, 2, 8, 1, 9, 3, 7}
	solver.quickSort(arr, 0, len(arr)-1)
	
	expected := []int{1, 2, 3, 5, 7, 8, 9}
	for i, exp := range expected {
		if arr[i] != exp {
			t.Errorf("Expected arr[%d] to be %d after quicksort, got %d", i, exp, arr[i])
		}
	}
	
	// Test with already sorted array
	arr2 := []int{1, 2, 3, 4, 5}
	solver.quickSort(arr2, 0, len(arr2)-1)
	
	for i, exp := range []int{1, 2, 3, 4, 5} {
		if arr2[i] != exp {
			t.Errorf("Expected arr2[%d] to be %d after quicksort, got %d", i, exp, arr2[i])
		}
	}
	
	// Test with reverse sorted array
	arr3 := []int{5, 4, 3, 2, 1}
	solver.quickSort(arr3, 0, len(arr3)-1)
	
	for i, exp := range []int{1, 2, 3, 4, 5} {
		if arr3[i] != exp {
			t.Errorf("Expected arr3[%d] to be %d after quicksort, got %d", i, exp, arr3[i])
		}
	}
}

func TestSolveLargeOptimized(t *testing.T) {
	// Test with 100 elements
	input := make([]int, 100)
	for i := 0; i < 100; i++ {
		input[i] = 100 - i // Reverse sorted
	}
	
	solver := NewSolver(input)
	ops := solver.Solve()
	
	// Verify the result is sorted
	result := validateSolution(input, ops)
	if !result {
		t.Error("Solution should result in sorted stack")
	}
	
	// Should use less than 700 operations
	if len(ops) >= 700 {
		t.Errorf("Solution uses %d operations, expected < 700", len(ops))
	}
}

func TestSpecificCase(t *testing.T) {
	// Test the specific case mentioned: "2 1 3 6 5 8"
	input := []int{2, 1, 3, 6, 5, 8}
	solver := NewSolver(input)
	ops := solver.Solve()
	
	// Should be less than 9 operations for this case
	if len(ops) >= 9 {
		t.Errorf("Expected less than 9 operations for input %v, got %d operations", input, len(ops))
		t.Logf("Operations: %v", ops)
	}
	
	// Verify the result is sorted
	result := validateSolution(input, ops)
	if !result {
		t.Errorf("Solution for %v should result in sorted stack", input)
	}
}

func TestVariousRandomInputs(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		maxOps   int
		describe string
	}{
		{
			name:     "10 elements",
			input:    []int{7, 3, 9, 1, 5, 2, 8, 4, 10, 6},
			maxOps:   100,
			describe: "Medium size stack",
		},
		{
			name:     "15 elements",
			input:    []int{15, 7, 3, 11, 1, 13, 5, 9, 2, 14, 6, 10, 4, 12, 8},
			maxOps:   150,
			describe: "Larger medium stack",
		},
		{
			name:     "Already sorted 10",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			maxOps:   0,
			describe: "Already sorted should use 0 operations",
		},
		{
			name:     "Reverse sorted 10",
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			maxOps:   100,
			describe: "Worst case ordering",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := NewSolver(tt.input)
			ops := solver.Solve()
			
			// Verify the result is sorted
			result := validateSolution(tt.input, ops)
			if !result {
				t.Errorf("%s: Solution for %v should result in sorted stack", tt.describe, tt.input)
			}
			
			// Check operation count
			if len(ops) > tt.maxOps {
				t.Errorf("%s: Solution uses %d operations, expected <= %d", tt.describe, len(ops), tt.maxOps)
			}
		})
	}
}

// Helper function to validate that a solution actually sorts the input
func validateSolution(input []int, ops []operations.Operation) bool {
	if len(input) == 0 {
		return true
	}
	
	stackA := stack.NewStack(input)
	stackB := stack.NewEmptyStack()
	
	// Execute all operations
	for _, op := range ops {
		err := operations.ExecuteOperation(stackA, stackB, op)
		if err != nil {
			return false
		}
	}
	
	// Check if stack A is sorted and stack B is empty
	return stackA.IsSorted() && stackB.IsEmpty()
}

// Helper function to check if a slice is sorted
func isSorted(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}

// Benchmark tests for performance evaluation
func BenchmarkSolveSmall(b *testing.B) {
	input := []int{3, 1, 4, 2, 5}
	
	for i := 0; i < b.N; i++ {
		solver := NewSolver(input)
		solver.Solve()
	}
}

func BenchmarkSolveMedium(b *testing.B) {
	input := []int{10, 3, 7, 1, 9, 2, 8, 4, 6, 5, 15, 12, 11, 13, 14}
	
	for i := 0; i < b.N; i++ {
		solver := NewSolver(input)
		solver.Solve()
	}
}

func BenchmarkSolveLarge(b *testing.B) {
	input := make([]int, 100)
	for i := 0; i < 100; i++ {
		input[i] = 100 - i // Reverse sorted
	}
	
	for i := 0; i < b.N; i++ {
		solver := NewSolver(input)
		solver.Solve()
	}
}

func BenchmarkSolveVeryLarge(b *testing.B) {
	input := make([]int, 500)
	for i := 0; i < 500; i++ {
		input[i] = 500 - i // Reverse sorted
	}
	
	for i := 0; i < b.N; i++ {
		solver := NewSolver(input)
		solver.Solve()
	}
}