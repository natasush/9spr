package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// Обработка крайнего случая - размер равен 0
	if size <= 0 {
		return []int{}
	}
	// Создаем новый генератор случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Создание слайса заданного размера
	data := make([]int, size)

	// Заполнение слайса случайными числами
	for i := 0; i < size; i++ {
		data[i] = r.Intn(size)
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	n := len(data)
	if n == 0 {
		return 0
	}

	chunks := CHUNKS
	if chunks > n {
		chunks = n
	}

	baseChunkSize := n / chunks
	remainder := n % chunks

	maximumChunks := make([]int, chunks)
	var wg sync.WaitGroup

	startIndex := 0
	for i := 0; i < chunks; i++ {
		size := baseChunkSize
		if remainder > 0 {
			size++
			remainder--
		}

		endIndex := startIndex + size

		wg.Add(1)
		go func(i int, chunk []int) {
			defer wg.Done()

			maximumChunks[i] = maximum(chunk)
		}(i, data[startIndex:endIndex])
		startIndex = endIndex
	}
	wg.Wait()

	// Возвращаем максимальное значение из всех частей
	return maximum(maximumChunks)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numberSequence := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxOneThread := maximum(numberSequence)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxOneThread, elapsed.Milliseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxSeveralThreads := maxChunks(numberSequence)

	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxSeveralThreads, elapsed.Milliseconds())
}
