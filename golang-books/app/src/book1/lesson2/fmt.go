package book1Lesson2

import "fmt"

func pracFmt() {
	fmt.Println("----------")
	fmt.Println("fmt.go")
	fmt.Println("----------")

	fmt.Printf("火星の表面で、私の体重は、%v ポンド \n", 176.0*0.3783)
	fmt.Printf("年齢は、%v 歳になるでしょう \n", 28*365/687)
	fmt.Printf("私の体重は、%v の表面で %v ポンドです。\n", "Earth", 176)

	fmt.Printf("%-15v $%4v \n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v \n", "SpaceA", 100)
}
