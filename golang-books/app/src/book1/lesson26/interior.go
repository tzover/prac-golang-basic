package book1Lesson26

import "fmt"

type character struct {
	name  string
	stats stats
}

func interior() {
	fmt.Println("----------")
	fmt.Println("interior.go")
	fmt.Println("----------")

	player := character{name: "Tojima"}
	levelUp(&player.stats)

	fmt.Printf("%+v \n", player)
	fmt.Printf("%+v \n", player.stats)

}
