Here are implementations of these classic programming exercises in Unified:

## Fibonacci

```unified
// Recursive implementation
fn fibonacci_recursive(n: Int) -> Int {
    if n <= 1 {
        return n
    }
    return fibonacci_recursive(n - 1) + fibonacci_recursive(n - 2)
}

// Iterative implementation (more efficient)
fn fibonacci(n: Int) -> Int {
    if n <= 1 {
        return n
    }
    
    let mut a = 0
    let mut b = 1
    let mut result = 0
    
    for _ in 2..=n {
        result = a + b
        a = b
        b = result
    }
    
    return result
}

// Generate Fibonacci sequence
fn fibonacci_sequence(n: Int) -> List<Int> {
    let mut sequence = new List<Int>()
    
    for i in 0..n {
        sequence.push(fibonacci(i))
    }
    
    return sequence
}

fn main() {
    let seq = fibonacci_sequence(10)
    print("First 10 Fibonacci numbers: ${seq.join(", ")}")
    // Output: First 10 Fibonacci numbers: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
}
```

## FizzBuzz

```unified
fn fizzbuzz(n: Int) {
    for i in 1..=n {
        if i % 15 == 0 {
            print("FizzBuzz")
        } else if i % 3 == 0 {
            print("Fizz")
        } else if i % 5 == 0 {
            print("Buzz")
        } else {
            print("${i}")
        }
    }
}

// Alternative implementation using string building
fn fizzbuzz_string(n: Int) -> List<String> {
    let mut results = new List<String>()
    
    for i in 1..=n {
        let mut result = ""
        
        if i % 3 == 0 {
            result += "Fizz"
        }
        
        if i % 5 == 0 {
            result += "Buzz"
        }
        
        if result.isEmpty() {
            result = "${i}"
        }
        
        results.push(result)
    }
    
    return results
}

fn main() {
    fizzbuzz(20)
    
    let results = fizzbuzz_string(100)
    print(results.join("\n"))
}
```

## Towers of Hanoi

```unified
fn towers_of_hanoi(n: Int, source: String, auxiliary: String, target: String) {
    if n == 1 {
        print("Move disk 1 from ${source} to ${target}")
        return
    }
    
    towers_of_hanoi(n - 1, source, target, auxiliary)
    print("Move disk ${n} from ${source} to ${target}")
    towers_of_hanoi(n - 1, auxiliary, source, target)
}

fn main() {
    let num_disks = 3
    print("Solution for ${num_disks} disks:")
    towers_of_hanoi(num_disks, "A", "B", "C")
}
```

## 8-Queens Problem

```unified
struct Board {
    size: Int
    queens: List<(Int, Int)>  // List of (row, col) positions
    
    fn new(size: Int) -> Self {
        return Self{
            size: size,
            queens: new List<(Int, Int)>()
        }
    }
    
    fn is_safe(self, row: Int, col: Int) -> Bool {
        for (queen_row, queen_col) in self.queens {
            // Check if queens attack each other horizontally, vertically, or diagonally
            if queen_row == row || queen_col == col || 
               (queen_row - queen_col) == (row - col) || 
               (queen_row + queen_col) == (row + col) {
                return false
            }
        }
        return true
    }
    
    fn place_queen(self: &mut Self, row: Int, col: Int) {
        self.queens.push((row, col))
    }
    
    fn remove_queen(self: &mut Self) {
        self.queens.pop()
    }
    
    fn display(self) {
        for row in 0..self.size {
            let mut line = ""
            for col in 0..self.size {
                if self.queens.contains(&(row, col)) {
                    line += "Q "
                } else {
                    line += ". "
                }
            }
            print(line)
        }
        print("")
    }
}

fn solve_queens(board: &mut Board, col: Int) -> Bool {
    // Base case: If all queens are placed
    if col >= board.size {
        return true
    }
    
    // Try placing queen in each row of the current column
    for row in 0..board.size {
        if board.is_safe(row, col) {
            // Place queen
            board.place_queen(row, col)
            
            // Recursively place rest of the queens
            if solve_queens(board, col + 1) {
                return true
            }
            
            // If placing queen doesn't lead to a solution, backtrack
            board.remove_queen()
        }
    }
    
    // No solution found with current configuration
    return false
}

fn eight_queens() {
    let mut board = new Board(8)
    
    if solve_queens(&mut board, 0) {
        print("Solution found:")
        board.display()
    } else {
        print("No solution exists.")
    }
}

fn main() {
    eight_queens()
}
```

## Bubble Sort

```unified
fn bubble_sort<T>(arr: &mut List<T>) where T: Ord {
    let n = arr.len()
    let mut swapped = true
    
    while swapped {
        swapped = false
        for i in 0..(n - 1) {
            if arr[i] > arr[i + 1] {
                // Swap elements
                let temp = move arr[i]
                arr[i] = move arr[i + 1]
                arr[i + 1] = move temp
                swapped = true
            }
        }
    }
}

fn main() {
    let mut numbers = [5, 2, 9, 1, 5, 6]
    bubble_sort(&mut numbers)
    print("Sorted array: ${numbers.join(", ")}")
    // Output: Sorted array: 1, 2, 5, 5, 6, 9
}
```

## Quick Sort

```unified
fn quick_sort<T>(arr: &mut List<T>) where T: Ord {
    quick_sort_range(arr, 0, arr.len() - 1)
}

fn quick_sort_range<T>(arr: &mut List<T>, low: Int, high: Int) where T: Ord {
    if low < high {
        let pivot_index = partition(arr, low, high)
        
        // Recursively sort elements before and after partition
        quick_sort_range(arr, low, pivot_index - 1)
        quick_sort_range(arr, pivot_index + 1, high)
    }
}

fn partition<T>(arr: &mut List<T>, low: Int, high: Int) -> Int where T: Ord {
    // Use last element as pivot
    let pivot = &arr[high]
    
    // Index of smaller element
    let mut i = low - 1
    
    for j in low..high {
        // If current element is smaller than or equal to pivot
        if arr[j] <= *pivot {
            i += 1
            
            // Swap elements
            let temp = move arr[i]
            arr[i] = move arr[j]
            arr[j] = move temp
        }
    }
    
    // Swap pivot element with element at i+1
    let temp = move arr[i + 1]
    arr[i + 1] = move arr[high]
    arr[high] = move temp
    
    // Return position of pivot
    return i + 1
}

fn main() {
    let mut numbers = [10, 7, 8, 9, 1, 5]
    quick_sort(&mut numbers)
    print("Sorted array: ${numbers.join(", ")}")
    // Output: Sorted array: 1, 5, 7, 8, 9, 10
}
```

## Conway's Game of Life

```unified
struct GameOfLife {
    width: Int
    height: Int
    grid: List<List<Bool>>
    
    fn new(width: Int, height: Int) -> Self {
        // Initialize empty grid
        let mut grid = new List<List<Bool>>()
        
        for _ in 0..height {
            let mut row = new List<Bool>()
            for _ in 0..width {
                row.push(false)
            }
            grid.push(row)
        }
        
        return Self{
            width: width,
            height: height,
            grid: grid
        }
    }
    
    fn set_cell(self: &mut Self, row: Int, col: Int, alive: Bool) {
        if row >= 0 && row < self.height && col >= 0 && col < self.width {
            self.grid[row][col] = alive
        }
    }
    
    fn random_seed(self: &mut Self, probability: Float) {
        for row in 0..self.height {
            for col in 0..self.width {
                // Random value between 0.0 and 1.0
                let random = std.random.float()
                self.grid[row][col] = random < probability
            }
        }
    }
    
    fn count_neighbors(self, row: Int, col: Int) -> Int {
        let mut count = 0
        
        for dx in -1..=1 {
            for dy in -1..=1 {
                if dx == 0 && dy == 0 {
                    continue  // Skip the cell itself
                }
                
                let neighbor_row = row + dy
                let neighbor_col = col + dx
                
                // Check bounds and count live neighbors
                if neighbor_row >= 0 && neighbor_row < self.height && 
                   neighbor_col >= 0 && neighbor_col < self.width &&
                   self.grid[neighbor_row][neighbor_col] {
                    count += 1
                }
            }
        }
        
        return count
    }
    
    fn next_generation(self: &mut Self) {
        // Create a new grid for the next generation
        let mut new_grid = new List<List<Bool>>()
        
        for row in 0..self.height {
            let mut new_row = new List<Bool>()
            
            for col in 0..self.width {
                let alive = self.grid[row][col]
                let neighbors = self.count_neighbors(row, col)
                
                let will_be_alive = if alive {
                    // Live cell survives if it has 2 or 3 neighbors
                    neighbors == 2 || neighbors == 3
                } else {
                    // Dead cell becomes alive if it has exactly 3 neighbors
                    neighbors == 3
                }
                
                new_row.push(will_be_alive)
            }
            
            new_grid.push(new_row)
        }
        
        // Update grid to new generation
        self.grid = new_grid
    }
    
    fn display(self) {
        for row in 0..self.height {
            let mut line = ""
            for col in 0..self.width {
                line += if self.grid[row][col] { "■ " } else { "□ " }
            }
            print(line)
        }
        print("")
    }
    
    fn run_simulation(self: &mut Self, generations: Int, delay_ms: Int) {
        print("Initial state:")
        self.display()
        
        for gen in 1..=generations {
            // Delay between generations
            std.time.sleep(delay_ms.milliseconds)
            
            self.next_generation()
            print("Generation ${gen}:")
            self.display()
        }
    }
}

fn main() {
    // Create a new game with a 20x10 grid
    let mut game = new GameOfLife(20, 10)
    
    // Set up a "glider" pattern
    game.set_cell(1, 2, true)
    game.set_cell(2, 3, true)
    game.set_cell(3, 1, true)
    game.set_cell(3, 2, true)
    game.set_cell(3, 3, true)
    
    // Run simulation for 20 generations
    game.run_simulation(20, 500)
}
```

These examples demonstrate key features of Unified including:
- Type safety and inference
- Pattern matching
- Mutable and immutable variables
- Explicit ownership management with `move`
- Generic programming
- Functional and imperative programming styles
- Interfaces for type constraints
- Strong standard library functionality

Each implementation follows Unified's idioms and best practices while solving classic programming problems.

