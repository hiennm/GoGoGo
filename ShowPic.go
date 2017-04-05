package main

import "golang.org/x/tour/pic"

/*Pic : Generate picture content*/
func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i := range slice {
		slice[i] = make([]uint8, dx)
		for j := range slice[i] {
			slice[i][j] = uint8(i * j)
		}
	}
	return slice
}

/*main function*/

func main() {
	pic.Show(Pic)
}
