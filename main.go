package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var array = [4][4]int{}

func main() {
	rand.Seed(time.Now().UnixNano())

	randGen(true)
	visualise()

	won := false
	lost := false

	for !won && !lost {
		var input string
		validIn := false

		for !validIn {
			fmt.Scanf("%s", &input)
			if input != "" {
				input = string(input[len(input)-1])
			}

			switch input {
			case "a":
				validIn = true
				arrCopy := array
				left()
				if arrCopy == array {
					validIn = false
				}
				array = arrCopy
			case "d":
				validIn = true
				arrCopy := array
				right()
				if arrCopy == array {
					validIn = false
				}
				array = arrCopy
			case "w":
				validIn = true
				arrCopy := array
				up()
				if arrCopy == array {
					validIn = false
				}
				array = arrCopy
			case "s":
				validIn = true
				arrCopy := array
				down()
				if arrCopy == array {
					validIn = false
				}
				array = arrCopy
			}
		}

		switch input {
		case "a":
			left()
		case "d":
			right()
		case "w":
			up()
		case "s":
			down()
		}

		randGen(false)

		lost = true
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				if array[y][x] == 0 {
					lost = false
				}
				if array[y][x] == 2048 {
					won = true
				}
			}
		}

		if lost {
			arrCopy := array

			left()
			right()
			up()
			down()

			if arrCopy != array {
				lost = false
			}
		}

		visualise()

		if lost {
			fmt.Println()
			fmt.Println("You lost!")
		}

		if won {
			fmt.Println()
			fmt.Println("You won!")
		}
	}
	
	var in string
	fmt.Scanln(&in)
}

func visualise() {
	fmt.Println()
	fmt.Println(" ____ ____ ____ ____ ")

	for y := 0; y < 4; y++ {
		fmt.Print("|")
		for x := 0; x < 4; x++ {
			iS := strconv.Itoa(array[y][x])
			i := len(iS)

			if i < 3 {
				fmt.Print(" ")
			}

			if array[y][x] == 0 {
				for j := 0; j < i; j++ {
					fmt.Print(" ")
				}
			} else {
				fmt.Print(array[y][x])
			}

			for j := i; j < 3; j++ {
				fmt.Print(" ")
			}
			if i == 3 {
				fmt.Print(" ")
			}

			fmt.Print("|")
		}
		fmt.Println()

		if y == 3 {
			fmt.Println("|    |    |    |    |")
		} else {
			fmt.Println("|____|____|____|____|")
		}
	}

	fmt.Println(" ¯¯¯¯ ¯¯¯¯ ¯¯¯¯ ¯¯¯¯ ")
}

func randGen(firstGen bool) {
	var y0 = []int{}
	var x0 = []int{}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if array[y][x] == 0 {
				x0 = append(x0, x)
				y0 = append(y0, y)
			}
		}
	}

	r := 2
	if firstGen {
		r = 2
	} else {
		r = rand.Intn(11)

		if r < 10 {
			r = 2
		}
		if r >= 10 {
			r = 4
		}
	}

	r0 := rand.Intn(len(x0))
	array[y0[r0]][x0[r0]] = r

}

func left() {
	pushL()
	mergeL()
	pushL()
}

func right() {
	pushR()
	mergeR()
	pushR()
}

func up() {
	pushU()
	mergeU()
	pushU()
}

func down() {
	pushD()
	mergeD()
	pushD()
}

func pushL() {
	for i := 0; i < 3; i++ {
		for y := 0; y < 4; y++ {
			for x := 2; x >= 0; x-- {
				if array[y][x] == 0 {
					array[y][x] = array[y][x+1]
					array[y][x+1] = 0
				}
			}
		}
	}
}

func mergeL() {
	for y := 0; y < 4; y++ {
		for x := 0; x < 3; x++ {
			if array[y][x] == array[y][x+1] {
				array[y][x] += array[y][x+1]
				array[y][x+1] = 0
			}
		}
	}
}

func pushR() {
	for i := 0; i < 3; i++ {
		for y := 0; y < 4; y++ {
			for x := 1; x < 4; x++ {
				if array[y][x] == 0 {
					array[y][x] = array[y][x-1]
					array[y][x-1] = 0
				}
			}
		}
	}
}

func mergeR() {
	for y := 0; y < 4; y++ {
		for x := 3; x > 0; x-- {
			if array[y][x] == array[y][x-1] {
				array[y][x] += array[y][x-1]
				array[y][x-1] = 0
			}
		}
	}
}

func pushU() {
	for i := 0; i < 3; i++ {
		for x := 0; x < 4; x++ {
			for y := 2; y >= 0; y-- {
				if array[y][x] == 0 {
					array[y][x] = array[y+1][x]
					array[y+1][x] = 0
				}
			}
		}
	}
}

func mergeU() {
	for x := 0; x < 4; x++ {
		for y := 0; y < 3; y++ {
			if array[y][x] == array[y+1][x] {
				array[y][x] += array[y+1][x]
				array[y+1][x] = 0
			}
		}
	}
}

func pushD() {
	for i := 0; i < 3; i++ {
		for x := 0; x < 4; x++ {
			for y := 1; y < 4; y++ {
				if array[y][x] == 0 {
					array[y][x] = array[y-1][x]
					array[y-1][x] = 0
				}
			}
		}
	}
}

func mergeD() {
	for x := 0; x < 4; x++ {
		for y := 3; y > 0; y-- {
			if array[y][x] == array[y-1][x] {
				array[y][x] += array[y-1][x]
				array[y-1][x] = 0
			}
		}
	}
}
