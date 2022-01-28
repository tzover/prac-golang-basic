package book1Lesson26

import "fmt"

func nasa() {
	fmt.Println("----------")
	fmt.Println("nasa.go")
	fmt.Println("----------")

	var administrator *string
	scolese := "Scolese"

	administrator = &scolese

	fmt.Println(*administrator) // Scolese

	bolden := "Bolden"
	administrator = &bolden

	fmt.Println(*administrator) // Bolden

	bolden = "test"
	fmt.Println(*administrator) // test

	// こっちからも変えられる
	*administrator = "test administrator"
	fmt.Println(*administrator) // test administrator
	fmt.Println(bolden)         // test administrator

	major := administrator
	*major = "Major test"

	fmt.Println("Bolden :", bolden) // Major test

	fmt.Println(administrator == major) // true

	lightfoot := "Robot M"
	administrator = &lightfoot
	fmt.Println(administrator == major) // false

	charles := *major
	*major = "bolden charles"
	fmt.Println(charles) // Major test

	fmt.Println(bolden) // bolden charles

	charles = "bolden charles"
	fmt.Println(charles == bolden) // true

	fmt.Println(&charles == &bolden) // false
	fmt.Println(&bolden)             //
	fmt.Println(&charles)            //

}
