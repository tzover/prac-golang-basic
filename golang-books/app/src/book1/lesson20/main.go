package book1Lesson20

import (
	"fmt"
	"math/rand"
	"time"
)

// Game rules
// 生きているセルは、生きている隣接セルが2個に満たないと、死ぬ
// 生きているセルは、生きている隣接セルが2個か3個であれば次世代まで生き続けられる
// 生きているセルは、生きている隣接セルが3個を超えると死ぬ
// 死んでいるセルは、生きている隣接セルが3個ならば、生きているセルになる

// よくわからん

const (
	width  = 80
	height = 15
)

type (
	Universe [][]bool // 2次元配列
)

func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

// セルが生きているかどうか判定する
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

// 生きている隣接セルを数える
func (u Universe) Neighbors(x, y int) int {
	n := 0

	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width+1)*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

func (u Universe) Show() {
	fmt.Print("\x0c", u.String())
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func Main() {
	fmt.Println("******* Package book1Lesson20 *******")
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a
	}
}
