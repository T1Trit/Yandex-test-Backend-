package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Taxi struct {
    id        int
    position  int
    timestamp int
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // Чтение первой строки
    scanner.Scan()
    firstLine := strings.Split(scanner.Text(), " ")
    N, _ := strconv.Atoi(firstLine[0]) // Число событий
    L, _ := strconv.Atoi(firstLine[1]) // Длина круга
    S, _ := strconv.Atoi(firstLine[2]) // Максимальная скорость

    // Хранилище для таксистов
    taxis := make(map[int]Taxi)

    // Обработка событий
    var results []string
    orderID := 0
    for i := 0; i < N; i++ {
        scanner.Scan()
        event := strings.Fields(scanner.Text())
        timestamp, _ := strconv.Atoi(event[1])

        switch event[0] {
        case "TAXI":
            taxiID, _ := strconv.Atoi(event[2])
            taxiPosition, _ := strconv.Atoi(event[3])
            taxis[taxiID] = Taxi{id: taxiID, position: taxiPosition, timestamp: timestamp}

        case "ORDER":
            orderPosition, _ := strconv.Atoi(event[3])
            orderTime, _ := strconv.Atoi(event[4])

            // Список подходящих таксистов
            var suitableTaxis []int
            for _, taxi := range taxis {
                minPos := (taxi.position + S*(timestamp-taxi.timestamp)) % L
                maxPos := (taxi.position + S*(timestamp-taxi.timestamp)) % L

                if minPos > maxPos {
                    // Диапазон пересекает начало круга
                    if (orderPosition-minPos+L)%L <= orderTime*S || (maxPos-orderPosition+L)%L <= orderTime*S {
                        suitableTaxis = append(suitableTaxis, taxi.id)
                    }
                } else {
                    // Диапазон не пересекает начало круга
                    if (orderPosition-minPos+L)%L <= orderTime*S && (maxPos-orderPosition+L)%L <= orderTime*S {
                        suitableTaxis = append(suitableTaxis, taxi.id)
                    }
                }
            }

            // Формирование ответа
            if len(suitableTaxis) == 0 {
                results = append(results, "-1")
            } else if len(suitableTaxis) <= 5 {
                result := ""
                for _, id := range suitableTaxis {
                    result += strconv.Itoa(id) + " "
                }
                results = append(results, strings.TrimSpace(result))
            } else {
                result := ""
                for i := 0; i < 5; i++ {
                    result += strconv.Itoa(suitableTaxis[i]) + " "
                }
                results = append(results, strings.TrimSpace(result))
            }

            orderID++
        }
    }

    // Вывод результатов
    for _, res := range results {
        fmt.Println(res)
    }
}
