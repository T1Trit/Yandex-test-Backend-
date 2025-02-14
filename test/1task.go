package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

func main() {
    // Чтение входных данных
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := strings.Split(scanner.Text(), " ")
    n, _ := strconv.Atoi(input[0]) // Количество этажей
    m, _ := strconv.Atoi(input[1]) // Количество квартир на этаже
    x, _ := strconv.Atoi(input[2]) // Количество окон в высоту
    y, _ := strconv.Atoi(input[3]) // Количество окон в ширину

    totalWindows := x * y
    minLitWindows := int(math.Ceil(float64(totalWindows) / 2)) // Минимальное количество окон с светом

    // Чтение состояния окон
    windowMatrix := make([]string, n*x)
    for i := 0; i < n*x; i++ {
        scanner.Scan()
        windowMatrix[i] = scanner.Text()
    }

    // Подсчёт бодрствующих квартир
    awakeCount := 0
    for floor := 0; floor < n; floor++ {
        for apartment := 0; apartment < m; apartment++ {
            litCount := 0
            for row := 0; row < x; row++ {
                start := apartment * y
                end := start + y
                litCount += countX(windowMatrix[floor*x+row][start:end])
            }
            if litCount >= minLitWindows {
                awakeCount++
            }
        }
    }

    // Вывод результата
    fmt.Println(awakeCount)
}

// Функция для подсчёта количества 'X' в строке
func countX(s string) int {
    count := 0
    for _, char := range s {
        if char == 'X' {
            count++
        }
    }
    return count
}
