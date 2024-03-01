//question link:    https://exercism.org/tracks/go/exercises/card-tricks

package cards

//Favori Kartlar, bu sırayla 2, 6 ve 9 numaralı kartlarla bir dilim döndürür.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem, belirli bir konumdaki bir dilimden bir öğeyi alır.
// Dizin aralık dışındaysa, -1 değerini döndürmesini istiyoruz.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

func SetItem(slice []int, index, newCard int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, newCard)
		return slice
	} else {
		slice[index] = newCard
	}
	return slice
}

func PrependItems(slice []int, values ...int) []int {
	slice = append(values, slice...)
	return slice
}

func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	} else {
		slice = append(slice[:index], slice[index+1:]...)
		return slice
	}
}
