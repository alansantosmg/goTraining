package arrays

import "fmt"

func main() {
	var planets [3]string
	planets[0] = "mercury"
	planets[1] = "venus"
	planets[2] = "earth"

	// literal array

	planets2 := [3]string{"mercury", "venus", "earth"}

	// comopositing literal array

	planets3 := [...]string{
		"mercury",
		"venus",
		"earth", // <---needs trailing comma
	}

	// compositing literal array

	dwarfs := [...]string{
		"ceres",
		"pluto",
		"mars",
	}

	// interacting with index
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

	// listing
	// 0 ceres
	// 1 pluto
	// 2 mars

}
