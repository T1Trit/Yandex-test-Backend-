package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readInput читает входные данные: размеры матрицы, минимальное расстояние и саму матрицу.
func readInput() (int, int, int, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(tokens[0])
	m, _ := strconv.Atoi(tokens[1])
	d, _ := strconv.Atoi(tokens[2])

	grid := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		grid[i] = scanner.Text()
	}
	return n, m, d, grid
}

// computeDistances вычисляет для каждой ячейки расстояние до ближайшей ячейки с символом 'x' или 'X'.
// Здесь distances хранится в []int32 для экономии памяти, а очередь реализована без лишнего копирования.
func computeDistances(grid []string, n, m int) []int32 {
	distances := make([]int32, n*m)
	queue := make([]int32, 0, n*m)
	maxDist := int32(n + m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			idx := i*m + j
			if grid[i][j] == 'x' || grid[i][j] == 'X' {
				distances[idx] = 0
				queue = append(queue, int32(idx))
			} else {
				distances[idx] = maxDist
			}
		}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// Обработка очереди без лишнего выделения памяти: идём по срезу, не удаляя первый элемент.
	for i := 0; i < len(queue); i++ {
		currentIdx := int(queue[i])
		x := currentIdx / m
		y := currentIdx % m
		for _, dir := range directions {
			nx := x + dir[0]
			ny := y + dir[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < m {
				idx := nx*m + ny
				if distances[idx] > distances[currentIdx]+1 {
					distances[idx] = distances[currentIdx] + 1
					queue = append(queue, int32(idx))
				}
			}
		}
	}
	return distances
}

// computePrefixSum строит префиксную сумму для "неподходящих" ячеек (где расстояние < d),
// используя одномерный массив типа int32 размером (n+1)×(m+1). Это снижает расход памяти вдвое по сравнению с int.
func computePrefixSum(distances []int32, n, m, d int) []int32 {
	prefix := make([]int32, (n+1)*(m+1))
	// Для каждой ячейки матрицы вычисляем значение: 1, если расстояние меньше d, иначе 0.
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			idx := i*m + j
			var val int32 = 0
			if distances[idx] < int32(d) {
				val = 1
			}
			// Индексы в префиксном массиве:
			// prefix[(i+1)*(m+1)+(j+1)] = prefix[i*(m+1)+(j+1)] + prefix[(i+1)*(m+1)+j] - prefix[i*(m+1)+j] + val
			prefix[(i+1)*(m+1)+j+1] = prefix[i*(m+1)+j+1] + prefix[(i+1)*(m+1)+j] - prefix[i*(m+1)+j] + val
		}
	}
	return prefix
}

// canPlaceSquare проверяет, можно ли разместить квадрат размера size так,
// чтобы все его ячейки были "подходящими" (т.е. расстояние >= d).
// С использованием префиксной суммы вычисление суммы для любого квадратного участка происходит за O(1).
func canPlaceSquare(prefix []int32, n, m, size int) bool {
	// Размер 0 — тривиально подходит.
	if size == 0 {
		return true
	}
	for i := 0; i <= n-size; i++ {
		for j := 0; j <= m-size; j++ {
			topLeft := i*(m+1) + j
			topRight := i*(m+1) + j + size
			bottomLeft := (i+size)*(m+1) + j
			bottomRight := (i+size)*(m+1) + j + size
			sum := prefix[bottomRight] - prefix[topRight] - prefix[bottomLeft] + prefix[topLeft]
			if sum == 0 {
				return true
			}
		}
	}
	return false
}

// min возвращает минимальное из двух чисел.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Чтение входных данных.
	n, m, d, grid := readInput()

	// Вычисляем расстояния до ближайших жилых объектов.
	distances := computeDistances(grid, n, m)

	// Построение одномерной префиксной суммы "неподходящих" ячеек с использованием int32.
	prefix := computePrefixSum(distances, n, m, d)

	// Бинарный поиск по размеру квадрата.
	low, high := 0, min(n, m)
	result := 0
	for low <= high {
		mid := (low + high) / 2
		if canPlaceSquare(prefix, n, m, mid) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// Вывод результата.
	fmt.Println(result)
}
