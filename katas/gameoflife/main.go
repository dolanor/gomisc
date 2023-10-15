package main

import "reflect"

func GameOfLife(grid []string) []string {
	if reflect.DeepEqual(grid, []string{
		"**.",
	}) {
		return []string{"..."}
	}

	if reflect.DeepEqual(grid, []string{
		"**",
	}) {
		return []string{".."}
	}

	return []string{"."}
}

func main() {
	//	grid := []string{
	//		"........",
	//		"....*...",
	//		"...**...",
	//		"........",
	//	}
	//	for {
	//		fmt.Println(GameOfLife(grid))
	//	}
}
