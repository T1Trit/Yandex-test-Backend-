package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TrieNode представляет узел префиксного дерева для цифровых представлений слов.
// Используется фиксированный массив из 10 элементов (индексирование по символу '0'-'9').
type TrieNode struct {
	children [10]*TrieNode
	// Если слово завершено в этом узле, поле word содержит оригинальное слово.
	word string
}

// Insert вставляет в Trie слово по его цифровому представлению.
// originalWord — исходное слово, которое нужно сохранить для восстановления ответа.
func (node *TrieNode) Insert(code string, originalWord string) {
	current := node
	for _, ch := range code {
		index := ch - '0'
		if current.children[index] == nil {
			current.children[index] = &TrieNode{}
		}
		current = current.children[index]
	}
	// Если несколько слов имеют одинаковое числовое представление,
	// переопределяем значение, как и в оригинальном решении.
	current.word = originalWord
}

// buildLetterToNumberMap строит отображение букв в цифровое представление.
func buildLetterToNumberMap() map[rune]string {
	letterToNumber := make(map[rune]string)
	phoneMapping := []string{"", "", "ABC", "DEF", "GHI", "JKL", "MNO", "PQRS", "TUV", "WXYZ"}
	for number, letters := range phoneMapping {
		for i, letter := range letters {
			// Для каждой буквы повторяем номер барабана (число) i+1 раз.
			letterToNumber[letter] = strings.Repeat(string(rune('0'+number)), i+1)
		}
	}
	return letterToNumber
}

// wordToNumber преобразует слово в цифровую строку, используя letterToNumber.
func wordToNumber(word string, letterToNumber map[rune]string) string {
	var result strings.Builder
	for _, char := range word {
		result.WriteString(letterToNumber[char])
	}
	return result.String()
}

// parseInt преобразует строку в целое число (без обработки ошибок).
func parseInt(s string) int {
	var res int
	for _, ch := range s {
		res = res*10 + int(ch-'0')
	}
	return res
}

func main() {
	// Создаём Scanner и увеличиваем размер буфера,
	// чтобы корректно обрабатывать строку s до 128·10³ символов.
	scanner := bufio.NewScanner(os.Stdin)
	// Устанавливаем максимальную длину токена (например, 128КБ + небольшой запас).
	buf := make([]byte, 1024)
	scanner.Buffer(buf, 128*1024+1024)

	// Считываем сообщение (цифровую строку)
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении сообщения")
		return
	}
	message := scanner.Text()

	// Считываем количество слов в словаре
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении размера словаря")
		return
	}
	n := parseInt(scanner.Text())

	// Построение отображения буква -> цифровое представление.
	letterToNumber := buildLetterToNumberMap()

	// Построение Trie на основе цифровых представлений слов словаря.
	trieRoot := &TrieNode{}
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			fmt.Fprintln(os.Stderr, "Ошибка при чтении слова из словаря")
			return
		}
		word := strings.ToUpper(scanner.Text())
		numericCode := wordToNumber(word, letterToNumber)
		trieRoot.Insert(numericCode, word)
	}

	mlen := len(message)
	// dp[i] == true означает, что первые i символов сообщения можно разобрать.
	dp := make([]bool, mlen+1)
	dp[0] = true

	// Для восстановления пути разбивки сохраняем предыдущий индекс и выбранное слово.
	prev := make([]int, mlen+1)
	resultWord := make([]string, mlen+1)
	for i := range prev {
		prev[i] = -1
	}

	// Динамическое программирование с обходом Trie по сообщению.
	for i := 0; i < mlen; i++ {
		if !dp[i] {
			continue
		}
		current := trieRoot
		// Переход по Trie для подстроки, начинающейся с позиции i.
		for j := i; j < mlen; j++ {
			ch := message[j]
			index := ch - '0'
			// Если обнаружен нецифровой символ или ветка отсутствует, прекращаем поиск.
			if index < 0 || index > 9 || current.children[index] == nil {
				break
			}
			current = current.children[index]
			// Если в узле найдено завершение слова, обновляем dp.
			if current.word != "" {
				if !dp[j+1] {
					dp[j+1] = true
					prev[j+1] = i
					resultWord[j+1] = current.word
				}
			}
		}
	}

	// Если разбить сообщение не удалось, выводим "No solution".
	if !dp[mlen] {
		fmt.Println("No solution")
		return
	}

	// Восстанавливаем последовательность слов (обратным проходом).
	var words []string
	pos := mlen
	for pos > 0 {
		if prev[pos] == -1 {
			fmt.Println("No solution")
			return
		}
		words = append(words, resultWord[pos])
		pos = prev[pos]
	}

	// Переворачиваем последовательность для восстановления исходного порядка.
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	fmt.Println(strings.Join(words, " "))
}
