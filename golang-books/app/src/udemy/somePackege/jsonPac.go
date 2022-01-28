package somePackage

import (
	"encoding/json"
	"fmt"
	"log"
)

type T struct{}

type Person struct {
	Name      string   `json:"name,omitempty"`
	Age       int      `json:"age,omitempty"`
	Nicknames []string `json:"nicknames,omitempty"`
	T         *T       `json:"T,omitempty"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		log.Fatalln(err)
	}
	p.Name = p2.Name + "!"
	return err
}

// 正規表現
func jsonPac() {
	fmt.Println("*** jsonPac() ***")

	b := []byte(`{"name": "mike", "age": 10, "nicknames": ["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// Marshal
	v, _ := json.Marshal(p)
	fmt.Println(string(v))

}
