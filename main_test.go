package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	// размер слайса 0 - крайний случай
	result := generateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("ожидался пустой слайс, получено %d элементов", len(result))
	}

	// размер слайса < 0 - крайний случай
	result = generateRandomElements(-10)
	if len(result) != 0 {
		t.Error("ожидался пустой слайс при отрицательном размере")
	}

	// нормальный случай
	result = generateRandomElements(333)
	if len(result) != 333 {
		t.Errorf("ожидалось 333 элементов, получено %d элементов", len(result))
	}

	// проверка, что все числа положительные
	for _, v := range result {
		if v <= 0 {
			t.Errorf("найдено неположительное число: %d", v)
		}
	}
}

func TestMaximum(t *testing.T) {
	// размер слайса 0 - крайний случай
	if maximum([]int{}) != 0 {
		t.Error("для пустого слайса ожидается 0")
	}

	// размер слайса 1 элемент
	if maximum([]int{421}) != 421 {
		t.Error("для слайса из одного элемента должен вернуться сам элемент")
	}

	// нормальный случай
	data := []int{3, 17, 89, 2, 56, 103, 7, 44, 901, 15, 333, 8, 103}
	if maximum(data) != 901 {
		t.Errorf("ожидался максимум 901, получен %d", maximum(data))
	}

	// все элементы одинаковые
	if maximum([]int{33, 33, 33, 33, 33}) != 33 {
		t.Error("при одинаковых элементах в слайсе в качестве максимума должен вернуться этот элемент")
	}

	// максимальный элемент в начале слайса
	if maximum([]int{192, 23, 46, 95, 11}) != 192 {
		t.Error("неверно обрабатывается максимум в начале слайса")
	}

	// максимальный элемент в конце слайса
	if maximum([]int{23, 46, 95, 11, 192}) != 192 {
		t.Error("неверно обрабатывается максимум в конце слайса")
	}
}

func TestMaxChunks(t *testing.T) {
	// размер слайса 0 - крайний случай
	if maxChunks([]int{}) != 0 {
		t.Error("для пустого слайса ожидается 0")
	}

	// слайс небольшого размера - меньше CHUNKS
	data := []int{5, 2, 8, 5, 4}
	if maxChunks(data) != 8 {
		t.Errorf("ожидался 8, получено %d", maxChunks(data))
	}

	// слайс побольше
	data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 123, 124543, 222, 6666, 3034839}
	if maxChunks(data) != 3034839 {
		t.Errorf("ожидался 3034839, получено %d", maxChunks(data))
	}

	// все элементы одинаковые
	data = make([]int, 124)
	for i := range data {
		data[i] = 124
	}
	if maxChunks(data) != 124 {
		t.Error("при одинаковых элементах в слайсе в качестве максимума должен вернуться этот элемент")
	}
}
