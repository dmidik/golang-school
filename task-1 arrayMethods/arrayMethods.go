package arrayMethods

// Добавляем элемент в конец строки
func push[T any](arr []T, value T) []T {
	return append(arr, value)
}
