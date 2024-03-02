// https://exercism.org/tracks/go/exercises/bird-watcher

package bird

func TotalBirdCount(birds []int) int {
	var sum int
	for i := 0; i < len(birds); i++ {
		sum += birds[i]
	}
	return sum
}

func BirdsInWeek(birds []int, week int) int {

	var sum int

	for i := (week - 1) * 7; i < week*7; i++ {
		sum += birds[i]
	}
	return sum
}

func FixBirdCountLog(birds []int) []int {

	for i := 0; i < len(birds); i++ {
		if i%2 == 0 {
			birds[i]++
		}
	}
	return birds
}
