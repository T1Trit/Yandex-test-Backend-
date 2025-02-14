package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // Чтение входных данных
    scanner := bufio.NewScanner(os.Stdin)

    // Размер векторов
    scanner.Scan()
    N, _ := strconv.Atoi(scanner.Text())

    // Вектор Q
    scanner.Scan()
    QStr := strings.Split(scanner.Text(), " ")
    Q := make([]int64, N)
    for i := 0; i < N; i++ {
        Q[i], _ = strconv.ParseInt(QStr[i], 10, 64)
    }

    // Вектор C
    scanner.Scan()
    CStr := strings.Split(scanner.Text(), " ")
    C := make([]int, N)
    for i := 0; i < N; i++ {
        C[i], _ = strconv.Atoi(CStr[i])
    }

    // Значения A и B
    scanner.Scan()
    ABStr := strings.Split(scanner.Text(), " ")
    A, _ := strconv.ParseInt(ABStr[0], 10, 64)
    B, _ := strconv.ParseInt(ABStr[1], 10, 64)

    // Вычисление скалярного произведения
    var dotProduct int64 = 0
    if A == B {
        // Если все элементы D одинаковы, то D_i = A для всех i
        for i := 0; i < N; i++ {
            dotProduct += Q[i] * A
        }
    } else {
        // Вычисляем множитель заранее
        scale := float64(B-A) / 255.0
        for i := 0; i < N; i++ {
            // Математическое округление: добавляем 0.5 перед преобразованием
            Di := int64(float64(A) + float64(C[i])*scale + 0.5)
            dotProduct += Q[i] * Di
        }
    }

    // Вывод результата
    fmt.Println(dotProduct)
}
