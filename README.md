# Recursion

## Greatest common divisor

The ***greatest common divisor*** (GCD) of two nonzero integers $a$ and $b$ is the
greatest positive integer $d$ such that $d$ is a divisor of both $a$ and $b$; that is,
there are integers $e$ and $f$ such that $a = de$ and $b = df$, and $d$ is the largest
such integer. The GCD of $a$ and $b$ is generally denoted $gcd(a, b)$.

### Example

$$gcd(12, 30) = 6$$

## Euclid's algorithm

The method introduced by Euclid for computing greatest common divisors is based
on the fact that, given two positive integers $a$ and $b$ such that $a > b$, the
common divisors of $a$ and $b$ are the same as the common divisors of $a – b$
and $b$.

$$gcd(a, b) = gcd(a-b, b)$$

### Example

$$gcd(52, 18) = gcd(52-18, 18)$$
$$gcd(34, 18) = gcd(34-18, 18)$$
$$gcd(16, 18) = gcd(18, 16) = gcd(18-16, 16)$$
$$gcd(16, 2) = 2$$

### Program

```golang
func gcd1(a, b uint64) uint64 {
    if a == b { return a }
    if a > b { return gcd1(a-b, b) }
    return gcd1(a, b-a)
}
```

## Euclidian algorithm

$$gcd(a, b) = gcd(a\ \text{mod}\ b, b)$$

### Example

$$gcd(52, 18) = gcd(52\ \text{mod} 18, 18)$$
$$gcd(16, 18) = gcd(18, 16) = gcd(18\ \text{mod}\ 16, 16)$$
$$gcd(2, 16) = gcd(16, 2) = 2$$

### Program

```golang
func gcd2(a, b uint64) uint64 {
    if b == 0 { return a }
    return gcd2(b, a%b)
}
```

## Exercises

### Exercise 1

https://codeforces.com/contest/527/problem/A

### Exercise 2

Write a recursive function that returns a slice of all file names (not directories)
in a given directory and its subdirectories. Use
[os.ReadDir](https://pkg.go.dev/os@go1.19.1#ReadDir) function to get the
directory content.

The output must contain file paths relative to the directory provided.

Testing hint: create testdata directory, which contains directories and files that
you can use for testing.

In your solution is isn't allowed to use library functions that implement similar
functionality e.g. `filepath.Walk`.

```golang
func ListFiles(dir string) []string
```

#### Example

```go
/*
Assuming the following directory tree structure:

/
+---- usr
|      +----- bin
|      |          cat
|      |          gcc
|      |          ls
|      +----- local
|             +------ bin
|                        go1.19
|                        mytool
+---- home
      +----- yarcat
                README.md
*/
fmt.Println(ListFiles("/usr")) 
Output: [bin/cat bin/gcc bin/ls local/bin/go1.19 local/bin/mytool]
```

### Exercise 3

Given a binary matrix $A$ of size $N \times M$. Cells which contain 1 are called filled
cell and cell that contain 0 are called empty cell. Two cells are said to be
connected if they are adjacent to each other horizontally, vertically, or
diagonally. If one or more filled cells are also connected, they form a region.
Find the area of the largest region.

```
func Area(m [][]int) int
```

#### Example

```
Input:
[][]int{
    []int{0, 0, 1, 1, 0},
    []int{1, 0, 1, 1, 0},
    []int{0, 1, 0, 0, 0},
    []int{0, 0, 0, 0, 0},
}
Output:  6
```

```
Input:
[][]int{
    []int{1, 1, 1},
    []int{0, 0, 1},
}
Output:  4
```
