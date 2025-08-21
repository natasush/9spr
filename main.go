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

	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создание слайса заданного размера
	data := make([]int, size)

	// Заполнение слайса случайными числами
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
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
	var wg sync.WaitGroup
	sizeChunk := SIZE / CHUNKS           //размер одной части (среза)
	maximumChunks := make([]int, CHUNKS) //слайс для максимумов частей
	for i := 0; i < CHUNKS; i++ {
		startIndex := i * sizeChunk        //начальный индекс среза
		endIndex := startIndex + sizeChunk //конечный индекс среза
		// Проверяем, чтобы endIndex не выходил за пределы длины data
		if endIndex > len(data) {
			endIndex = len(data)
		}
		wg.Add(1)

		go func(i, start, end int) {
			defer wg.Done()
			maxChunk := maximum(data[start:end]) //максимум части(среза)
			maximumChunks[i] = maxChunk
		}(i, startIndex, endIndex) //добавляем найденный максимум части в слайс
	}

	wg.Wait()

	// Возвращаем максимальное значение из всех частей
	maxOverall := maximum(maximumChunks)
	return maxOverall
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
