package book1Lesson3

import "fmt"

func torch() {
	fmt.Println("----------")
	fmt.Println("torch.go")
	fmt.Println("----------")

	var (
		haveTorch bool = true
		litTorch  bool
	)

	if !haveTorch || !litTorch {
		fmt.Println("I can not see anything")
	}

}
