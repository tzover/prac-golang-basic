package book1Lesson26

import "fmt"

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

func day() {
	fmt.Println("----------")
	fmt.Println("day.go")
	fmt.Println("----------")

}
