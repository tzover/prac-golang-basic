package book1Lesson24

type stardater interface {
	YearDay() int
	Hour() int
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}
