package pracStruct

import "fmt"

type Person2 struct {
	Name string
	Age  int
}

func (p Person2) String() string {
	return fmt.Sprintf("My Name is %v \n", p.Name)
}

func stringPrint() {
	mike := Person2{Name: "Mike", Age: 22}
	fmt.Println(mike)

}
