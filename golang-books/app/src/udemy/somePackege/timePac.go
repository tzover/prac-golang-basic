package somePackage

import (
	"fmt"
	"time"
)

func timePac() {
	fmt.Println("*** timePac() ***")
	t := time.Now()
	fmt.Println("Default : ", t)
	fmt.Println("RFC3339 : ", t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	fmt.Println(t.Date())
}
