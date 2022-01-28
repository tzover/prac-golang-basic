package book1Lesson26

import "fmt"

func Main() {
	fmt.Println("******* Package book1Lesson26 *******")

	memory()
	pracType()
	home()
	nasa()
	pracStruct()
	superpower()

	rebecca := person{
		name:       "Rebecce",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)

	fmt.Printf("%+v \n", rebecca)

	method()

	day()

	interior()

	pracArray()

	pracSlice()

	mainMartian()

	pracInterface()

}
