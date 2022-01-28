package panicAndRecover

import "fmt"

func thirdPartyConnectDB() {
	// panic はあまり使わない　エラーハンドリングをしっかりやる
	panic("Unable to connect database")
}

func save() {

	// Error が起きるが defer recover で復帰できる
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func Main() {
	fmt.Println("********** Packege panicAndRecover -> panicAndRecover.go **********")

	save()
	fmt.Println("ok?")
}
