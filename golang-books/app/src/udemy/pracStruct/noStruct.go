package pracStruct

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}

func noStruct() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
