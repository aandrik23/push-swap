package solver

import (
	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

type Solver struct {
	stackA     *stack.Stack
	stackB     *stack.Stack
	operations []operations.Operation
}

func NewSolver(input []int) *Solver {
	return &Solver{
		stackA:     stack.NewStack(input),
		stackB:     stack.NewEmptyStack(),
		operations: make([]operations.Operation, 0),
	}
}

func (s *Solver) Solve() []operations.Operation {
	if s.stackA.IsSorted() {
		return s.operations
	}
	
	size := s.stackA.Size()
	
	switch {
	case size <= 1:
		return s.operations
	case size == 2:
		s.solveTwo()
	case size == 3:
		s.solveThree()
	case size <= 5:
		s.solveFive()
	case size == 6:
		s.solveSixImproved()
	default:
		s.solveLargeOptimized()
	}
	
	return s.operations
}

func (s *Solver) executeAndRecord(op operations.Operation) {
	operations.ExecuteOperation(s.stackA, s.stackB, op)
	s.operations = append(s.operations, op)
}

func (s *Solver) solveTwo() {
	first, _ := s.stackA.At(0)
	second, _ := s.stackA.At(1)
	if first > second {
		s.executeAndRecord(operations.SA)
	}
}

func (s *Solver) solveThree() {
	for !s.stackA.IsSorted() {
		first, _ := s.stackA.At(0)
		second, _ := s.stackA.At(1)
		third, _ := s.stackA.At(2)
		
		if first > second && second < third && first < third {
			s.executeAndRecord(operations.SA)
		} else if first > second && second > third && first > third {
			s.executeAndRecord(operations.SA)
			s.executeAndRecord(operations.RRA)
		} else if first > second && second < third && first > third {
			s.executeAndRecord(operations.RA)
		} else if first < second && second > third && first < third {
			s.executeAndRecord(operations.SA)
			s.executeAndRecord(operations.RA)
		} else if first < second && second > third && first > third {
			s.executeAndRecord(operations.RRA)
		}
	}
}

func (s *Solver) solveFive() {
	size := s.stackA.Size()
	elementsToMove := size - 3
	
	for i := 0; i < elementsToMove; i++ {
		minPos := s.findMinPosition(s.stackA)
		s.moveToTopOptimized(s.stackA, minPos, true)
		s.executeAndRecord(operations.PB)
	}
	
	s.solveThree()
	
	for !s.stackB.IsEmpty() {
		s.executeAndRecord(operations.PA)
	}
}

func (s *Solver) solveSixImproved() {
	first, _ := s.stackA.At(0)
	second, _ := s.stackA.At(1)
	minVal := s.findMinValue(s.stackA)
	
	if second == minVal || (first > second && second < s.getValueAt(s.stackA, 2)) {
		s.executeAndRecord(operations.SA)
	}
	
	for i := 0; i < 2; i++ {
		minPos := s.findMinPosition(s.stackA)
		s.moveToTopOptimized(s.stackA, minPos, true)
		s.executeAndRecord(operations.PB)
	}
	
	s.solveFourInPlace()
	
	s.executeAndRecord(operations.PA)
	s.executeAndRecord(operations.PA)
}

func (s *Solver) solveFourInPlace() {
	minPos := s.findMinPosition(s.stackA)
	s.moveToTopOptimized(s.stackA, minPos, true)
	s.executeAndRecord(operations.PB)
	s.solveThree()
	s.executeAndRecord(operations.PA)
}

// solveLargeOptimized uses a simple but reliable push-all then insert-back approach
func (s *Solver) solveLargeOptimized() {
	size := s.stackA.Size()
	
	// Convert to ranks for easier handling
	ranks := s.createRanks()
	s.applyRanks(ranks)
	
	// Calculate chunk size
	var chunkSize int
	if size <= 100 {
		chunkSize = 20
	} else if size <= 500 {
		chunkSize = 30
	} else {
		chunkSize = 45
	}
	
	numChunks := (size + chunkSize - 1) / chunkSize
	
	// Push all elements to B in chunks
	for chunk := 0; chunk < numChunks; chunk++ {
		minRange := chunk * chunkSize
		maxRange := (chunk + 1) * chunkSize
		if maxRange > size {
			maxRange = size
		}
		
		// Push all elements in this range
		elementsInChunk := maxRange - minRange
		rotations := 0
		
		for elementsInChunk > 0 && rotations < size*2 { // Safety limit
			if s.stackA.IsEmpty() {
				break
			}
			
			top, _ := s.stackA.Top()
			
			if top >= minRange && top < maxRange {
				s.executeAndRecord(operations.PB)
				elementsInChunk--
				
				// Smart rotation in B
				if s.stackB.Size() > 1 {
					midRange := (minRange + maxRange) / 2
					if top < midRange {
						s.executeAndRecord(operations.RB)
					}
				}
			} else {
				s.executeAndRecord(operations.RA)
			}
			
			rotations++
		}
	}
	
	// Push everything back from B to A (largest first)
	for s.stackB.Size() > 0 {
		maxPos := s.findMaxPosition(s.stackB)
		s.moveToTopOptimized(s.stackB, maxPos, false)
		s.executeAndRecord(operations.PA)
	}
}

func (s *Solver) findMinPosition(st *stack.Stack) int {
	if st.IsEmpty() {
		return -1
	}
	
	minVal, _ := st.At(0)
	minPos := 0
	
	for i := 1; i < st.Size(); i++ {
		val, _ := st.At(i)
		if val < minVal {
			minVal = val
			minPos = i
		}
	}
	
	return minPos
}

func (s *Solver) findMinValue(st *stack.Stack) int {
	if st.IsEmpty() {
		return 0
	}
	
	minVal, _ := st.At(0)
	
	for i := 1; i < st.Size(); i++ {
		val, _ := st.At(i)
		if val < minVal {
			minVal = val
		}
	}
	
	return minVal
}

func (s *Solver) getValueAt(st *stack.Stack, index int) int {
	if index >= st.Size() {
		return 999999
	}
	val, _ := st.At(index)
	return val
}

func (s *Solver) findMaxPosition(st *stack.Stack) int {
	if st.IsEmpty() {
		return -1
	}
	
	maxVal, _ := st.At(0)
	maxPos := 0
	
	for i := 1; i < st.Size(); i++ {
		val, _ := st.At(i)
		if val > maxVal {
			maxVal = val
			maxPos = i
		}
	}
	
	return maxPos
}

func (s *Solver) moveToTopOptimized(st *stack.Stack, position int, isStackA bool) {
	size := st.Size()
	if position == 0 || size <= 1 {
		return
	}
	
	if position <= size/2 {
		for i := 0; i < position; i++ {
			if isStackA {
				s.executeAndRecord(operations.RA)
			} else {
				s.executeAndRecord(operations.RB)
			}
		}
	} else {
		rotations := size - position
		for i := 0; i < rotations; i++ {
			if isStackA {
				s.executeAndRecord(operations.RRA)
			} else {
				s.executeAndRecord(operations.RRB)
			}
		}
	}
}

func (s *Solver) createRanks() map[int]int {
	data := s.stackA.ToSlice()
	n := len(data)
	
	// Create a copy for sorting
	sorted := make([]int, n)
	copy(sorted, data)
	
	// Use a more efficient sort for larger arrays
	if n > 50 {
		s.quickSort(sorted, 0, n-1)
	} else {
		// Bubble sort for small arrays
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if sorted[j] > sorted[j+1] {
					sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				}
			}
		}
	}
	
	ranks := make(map[int]int)
	for i, val := range sorted {
		ranks[val] = i
	}
	
	return ranks
}

func (s *Solver) quickSort(arr []int, low, high int) {
	if low < high {
		pi := s.partition(arr, low, high)
		s.quickSort(arr, low, pi-1)
		s.quickSort(arr, pi+1, high)
	}
}

func (s *Solver) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func (s *Solver) applyRanks(ranks map[int]int) {
	data := s.stackA.ToSlice()
	for i, val := range data {
		data[i] = ranks[val]
	}
	*s.stackA = *stack.NewStack(data)
}