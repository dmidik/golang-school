package task_1

// Добавляем элемент в конец строки
func push[T any](arr []T, value T) []T {
	return append(arr, value)
}

func delete[T any](arr []T, index int) []T {
	copy(arr[index:], arr[index+1:])
	arr = arr[:len(arr)-1] // Обрезаем срез
	return arr
}
