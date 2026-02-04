package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"push-swap/internal/operations"
	"push-swap/internal/parser"
	"push-swap/internal/stack"
	"strconv"
	"strings"
	"time"
)

// Modified Solver that captures operations step-by-step
type VisualizerSolver struct {
	stackA     *stack.Stack
	stackB     *stack.Stack
	operations []operations.Operation
}

func NewVisualizerSolver(input []int) *VisualizerSolver {
	return &VisualizerSolver{
		stackA:     stack.NewStack(input),
		stackB:     stack.NewEmptyStack(),
		operations: make([]operations.Operation, 0),
	}
}

func (s *VisualizerSolver) executeAndRecord(op operations.Operation) {
	operations.ExecuteOperation(s.stackA, s.stackB, op)
	s.operations = append(s.operations, op)
}

func (s *VisualizerSolver) Solve() []operations.Operation {
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
		s.solveSix()
	default:
		s.solveLarge()
	}
	
	return s.operations
}

func (s *VisualizerSolver) solveTwo() {
	first, _ := s.stackA.At(0)
	second, _ := s.stackA.At(1)
	if first > second {
		s.executeAndRecord(operations.SA)
	}
}

func (s *VisualizerSolver) solveThree() {
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

func (s *VisualizerSolver) solveFive() {
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

func (s *VisualizerSolver) solveSix() {
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
	
	minPos := s.findMinPosition(s.stackA)
	s.moveToTopOptimized(s.stackA, minPos, true)
	s.executeAndRecord(operations.PB)
	s.solveThree()
	s.executeAndRecord(operations.PA)
	
	s.executeAndRecord(operations.PA)
	s.executeAndRecord(operations.PA)
}

func (s *VisualizerSolver) solveLarge() {
	size := s.stackA.Size()
	
	ranks := s.createRanks()
	s.applyRanks(ranks)
	
	var chunkSize int
	if size <= 100 {
		chunkSize = 20
	} else if size <= 500 {
		chunkSize = 30
	} else {
		chunkSize = 45
	}
	
	numChunks := (size + chunkSize - 1) / chunkSize
	
	for chunk := 0; chunk < numChunks; chunk++ {
		minRange := chunk * chunkSize
		maxRange := (chunk + 1) * chunkSize
		if maxRange > size {
			maxRange = size
		}
		
		elementsInChunk := maxRange - minRange
		rotations := 0
		
		for elementsInChunk > 0 && rotations < size*2 {
			if s.stackA.IsEmpty() {
				break
			}
			
			top, _ := s.stackA.Top()
			
			if top >= minRange && top < maxRange {
				s.executeAndRecord(operations.PB)
				elementsInChunk--
				
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
	
	for s.stackB.Size() > 0 {
		maxPos := s.findMaxPosition(s.stackB)
		s.moveToTopOptimized(s.stackB, maxPos, false)
		s.executeAndRecord(operations.PA)
	}
}

func (s *VisualizerSolver) findMinPosition(st *stack.Stack) int {
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

func (s *VisualizerSolver) findMinValue(st *stack.Stack) int {
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

func (s *VisualizerSolver) getValueAt(st *stack.Stack, index int) int {
	if index >= st.Size() {
		return 999999
	}
	val, _ := st.At(index)
	return val
}

func (s *VisualizerSolver) findMaxPosition(st *stack.Stack) int {
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

func (s *VisualizerSolver) moveToTopOptimized(st *stack.Stack, position int, isStackA bool) {
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

func (s *VisualizerSolver) createRanks() map[int]int {
	data := s.stackA.ToSlice()
	n := len(data)
	
	sorted := make([]int, n)
	copy(sorted, data)
	
	if n > 50 {
		s.quickSort(sorted, 0, n-1)
	} else {
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

func (s *VisualizerSolver) quickSort(arr []int, low, high int) {
	if low < high {
		pi := s.partition(arr, low, high)
		s.quickSort(arr, low, pi-1)
		s.quickSort(arr, pi+1, high)
	}
}

func (s *VisualizerSolver) partition(arr []int, low, high int) int {
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

func (s *VisualizerSolver) applyRanks(ranks map[int]int) {
	data := s.stackA.ToSlice()
	for i, val := range data {
		data[i] = ranks[val]
	}
	*s.stackA = *stack.NewStack(data)
}

type VisualizerState struct {
	StackA    []int  `json:"stackA"`
	StackB    []int  `json:"stackB"`
	Operation string `json:"operation"`
	OpCount   int    `json:"opCount"`
}

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Push-Swap Visualizer</title>
    <style>
        :root {
            --primary: #667eea;
            --secondary: #764ba2;
            --danger: #ef4444;
            --bg-dark: #0f0f1e;
            --bg-card: rgba(255, 255, 255, 0.05);
            --border: rgba(255, 255, 255, 0.1);
            --text-main: #e0e0e0;
            --text-dim: #a0a0b0;
        }

        * { margin: 0; padding: 0; box-sizing: border-box; }
        
        body {
            font-family: 'JetBrains Mono', monospace;
            background: linear-gradient(135deg, #0f0f1e 0%, #1a1a2e 100%);
            color: var(--text-main);
            min-height: 100vh;
            padding: 2rem;
            display: flex;
            flex-direction: column;
            align-items: center;
        }

        .wrapper { max-width: 1000px; width: 100%; }
        header { text-align: center; margin-bottom: 2rem; }
        
        h1 {
            font-size: 2.5rem;
            background: linear-gradient(135deg, var(--primary) 0%, var(--secondary) 100%);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .card {
            background: var(--bg-card);
            backdrop-filter: blur(10px);
            border-radius: 12px;
            padding: 1.5rem;
            margin-bottom: 1.5rem;
            border: 1px solid var(--border);
        }

        .controls { display: flex; flex-direction: column; gap: 1rem; }
        .input-row { display: flex; gap: 1rem; align-items: center; }
        
        input[type="text"] {
            flex: 1;
            padding: 0.75rem;
            background: rgba(0,0,0,0.3);
            border: 1px solid var(--border);
            border-radius: 6px;
            color: #fff;
            font-family: inherit;
        }

        button {
            padding: 0.75rem 1.25rem;
            border-radius: 6px;
            border: none;
            cursor: pointer;
            font-weight: 600;
            transition: 0.2s;
        }

        .btn-primary { background: var(--primary); color: white; }
        .btn-danger { background: var(--danger); color: white; }
        .btn-outline { background: transparent; border: 1px solid var(--primary); color: var(--primary); }
        button:hover:not(:disabled) { opacity: 0.8; transform: translateY(-1px); }
        button:disabled { opacity: 0.4; cursor: not-allowed; }

        .hidden { display: none; }

        .stats { display: flex; justify-content: space-around; margin-bottom: 1.5rem; }
        .stat-item { text-align: center; }
        .stat-val { font-size: 1.5rem; font-weight: bold; color: var(--primary); }

        .stacks-container {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 1.5rem;
            margin-bottom: 1.5rem;
        }

        .stack-box {
            background: rgba(255,255,255,0.02);
            border: 1px solid var(--border);
            border-radius: 8px;
            padding: 1rem;
            min-height: 400px;
        }

        .stack-label { text-align: center; margin-bottom: 1rem; color: var(--text-dim); font-size: 0.8rem; text-transform: uppercase; }

        .item-bar {
            margin-bottom: 2px;
            border-radius: 2px;
            color: white;
            font-size: 0.7rem;
            padding: 2px 8px;
            text-align: right;
        }

        #historyLog {
            display: flex;
            flex-wrap: wrap;
            gap: 0.5rem;
            max-height: 200px;
            overflow-y: auto;
            padding: 10px;
            background: rgba(0,0,0,0.2);
            border-radius: 8px;
        }

        .op-chip {
            padding: 4px 10px;
            background: rgba(102, 126, 234, 0.1);
            border: 1px solid var(--primary);
            color: var(--primary);
            border-radius: 4px;
            font-size: 0.8rem;
            font-weight: bold;
        }
    </style>
</head>
<body>
    <div class="wrapper">
        <header>
            <h1>Push-Swap</h1>
            <p style="color: var(--text-dim)">Visualizing sorting efficiency (Max: 100)</p>
        </header>

        <div class="card">
            <div class="controls">
                <div class="input-row">
                    <input type="text" id="numsInput" placeholder="Enter numbers or generate...">
                    <button class="btn-outline" id="r100" onclick="generateRandom(100)">Rand 100</button>
                    <button class="btn-outline" id="r50" onclick="generateRandom(50)">Rand 50</button>
                    <button class="btn-primary" id="runBtn" onclick="startSort()">Run</button>
                    <button class="btn-danger hidden" id="stopBtn" onclick="stopSort()">Stop</button>
                </div>
                <div class="input-row">
                    <label style="font-size: 0.8rem; color: var(--text-dim)">Speed (ms):</label>
                    <input type="range" id="speedInput" min="1" max="500" value="50">
                    <span id="speedDisplay">50</span>
                </div>
            </div>
        </div>

        <div class="stats">
            <div class="stat-item">
                <div class="stat-val" id="moveCount">0</div>
                <div style="font-size: 0.7rem; color: var(--text-dim)">OPERATIONS</div>
            </div>
            <div class="stat-item">
                <div class="stat-val" id="lastMove">-</div>
                <div style="font-size: 0.7rem; color: var(--text-dim)">LAST OP</div>
            </div>
        </div>

        <div class="stacks-container">
            <div class="stack-box">
                <div class="stack-label">Stack A</div>
                <div id="stackA"></div>
            </div>
            <div class="stack-box">
                <div class="stack-label">Stack B</div>
                <div id="stackB"></div>
            </div>
        </div>

        <div class="card">
            <div class="stack-label" style="text-align: left;">Command History</div>
            <div id="historyLog"></div>
        </div>
    </div>

    <script>
        let eventSource = null;

        document.getElementById('speedInput').oninput = function() {
            document.getElementById('speedDisplay').innerText = this.value;
        };

        function generateRandom(count) {
            const limit = Math.min(count, 100);
            const arr = Array.from({length: limit}, (_, i) => i + 1);
            for (let i = arr.length - 1; i > 0; i--) {
                const j = Math.floor(Math.random() * (i + 1));
                [arr[i], arr[j]] = [arr[j], arr[i]];
            }
            document.getElementById('numsInput').value = arr.join(' ');
        }

        function getColor(val, min, max) {
            const ratio = (val - min) / (max - min || 1);
            return "hsl(" + (220 + ratio * 140) + ", 70%, 60%)";
        }

        function render(state) {
            const all = [...state.stackA, ...state.stackB];
            const min = Math.min(...all), max = Math.max(...all);
            const draw = (id, items) => {
                const container = document.getElementById(id);
                container.innerHTML = items.map(n => {
                    const width = 20 + ((n - min) / (max - min || 1) * 80);
                    return '<div class="item-bar" style="width:' + width + '%; background:' + getColor(n, min, max) + '">' + n + '</div>';
                }).join('');
            };
            draw('stackA', state.stackA);
            draw('stackB', state.stackB);
            document.getElementById('moveCount').innerText = state.opCount;
            document.getElementById('lastMove').innerText = state.operation || '-';
            if (state.operation) {
                const log = document.getElementById('historyLog');
                const chip = document.createElement('div');
                chip.className = 'op-chip';
                chip.innerText = state.operation;
                log.appendChild(chip);
                log.scrollTop = log.scrollHeight;
            }
        }

        function toggleUI(isRunning) {
            document.getElementById('runBtn').classList.toggle('hidden', isRunning);
            document.getElementById('stopBtn').classList.toggle('hidden', !isRunning);
            document.getElementById('r100').disabled = isRunning;
            document.getElementById('r50').disabled = isRunning;
            document.getElementById('numsInput').disabled = isRunning;
        }

        function stopSort() {
            if (eventSource) {
                eventSource.close();
                eventSource = null;
            }
            toggleUI(false);
        }

        function startSort() {
            const input = document.getElementById('numsInput').value;
            const speed = document.getElementById('speedInput').value;
            const log = document.getElementById('historyLog');

            if (input.trim().split(/\s+/).length > 100) {
                alert("Limit is 100 numbers.");
                return;
            }

            log.innerHTML = '';
            toggleUI(true);

            eventSource = new EventSource('/visualize?numbers=' + encodeURIComponent(input) + '&speed=' + speed);
            eventSource.onmessage = (e) => render(JSON.parse(e.data));
            eventSource.addEventListener('complete', () => stopSort());
            eventSource.onerror = () => stopSort();
        }
    </script>
</body>
</html>`

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/visualize", handleVisualize)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("index").Parse(htmlTemplate)
	tmpl.Execute(w, nil)
}

func handleVisualize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Get context to detect when client stops/disconnects
	ctx := r.Context()

	numbersStr := r.URL.Query().Get("numbers")
	speedStr := r.URL.Query().Get("speed")
	speed := 150
	if s, err := strconv.Atoi(speedStr); err == nil {
		speed = s
	}

	numbers, _ := parser.ParseArguments(strings.Fields(numbersStr))
	if len(numbers) == 0 { return }

	flusher := w.(http.Flusher)
	solver := NewVisualizerSolver(numbers)
	ops := solver.Solve()

	stackA := stack.NewStack(numbers)
	stackB := stack.NewEmptyStack()

	sendState(w, flusher, stackA, stackB, "", 0)

	for i, op := range ops {
		// CHECK IF CLIENT DISCONNECTED
		select {
		case <-ctx.Done():
			// User clicked stop or closed tab, stop processing
			return
		default:
			time.Sleep(time.Duration(speed) * time.Millisecond)
			operations.ExecuteOperation(stackA, stackB, op)
			sendState(w, flusher, stackA, stackB, string(op), i+1)
		}
	}

	completeData, _ := json.Marshal(map[string]interface{}{"operations": len(ops)})
	fmt.Fprintf(w, "event: complete\ndata: %s\n\n", completeData)
	flusher.Flush()
}

func sendState(w http.ResponseWriter, flusher http.Flusher, stackA, stackB *stack.Stack, op string, opCount int) {
	state := VisualizerState{
		StackA:    stackA.ToSlice(),
		StackB:    stackB.ToSlice(),
		Operation: op,
		OpCount:   opCount,
	}
	data, _ := json.Marshal(state)
	fmt.Fprintf(w, "data: %s\n\n", data)
	flusher.Flush()
}