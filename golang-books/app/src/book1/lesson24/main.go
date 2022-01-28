package book1Lesson24

import (
	"fmt"
	"time"
)

func Main() {
	fmt.Println("******* Package book1Lesson24 *******")

	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(5))

	s := startship{laser(5)}
	fmt.Println(s.talk())

	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)

	fmt.Printf("%.1f Curiosity has landed \n", stardate(day))

	so := sol(1422)

	fmt.Printf("%.1f Happy birthday \n", stardate(so))

	stringer()
}
