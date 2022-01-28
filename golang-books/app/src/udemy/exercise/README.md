# ch01

Q1. 以下の1.11をint型に変換して出力してください。

f := 1.11

Q2. コードを書かずに以下の出力結果を答えてください。

s := []int{1, 2, 5, 6, 2, 3, 1}
fmt.Println(s[2:4])
Q3. 以下のコードを実行した時に

fmt.Printf("%T %v", m, m)

以下のような出力結果となるmを作成してください。

map[string]int map[Mike:20 Nancy:24 Messi:30]

# ch02

Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。

l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

Q2. 以下の果物の価格の合計を出力するコードを書いてください。

m := map[string]int{
    "apple":  200,
    "banana": 300,
    "grapes": 150,
    "orange": 80,
    "papaya": 500,
    "kiwi":   90,
}

# ch03

Q1. 以下の実行結果をコードを書かずに答えてください。エラーがおきますか？正しく表示されるとすると何が表示字されますか？

package main
 
import (
    "fmt"
)
 
func main() {
    var i int = 10
    var p *int
    p = &i
    var j int
    j = *p // Error
    fmt.Println(j)
}
Q2. 以下の実行結果をコードを書かずに答えてください。エラーがおきますか？正しく表示されるとすると何が表示されますか？

package main
 
import (
    "fmt"
)
 
func main() {
    var i int = 100
    var j int = 200
    var p1 *int
    var p2 *int
    p1 = &i
    p2 = &j
    i = *p1 + *p2
    p2 = p1
    j = *p2 + i
    fmt.Println(j)
}

# ch04

Q1. 以下に、7と表示されるメソッドを作成してください。

package main
 
import (
    "fmt"
)
 
type Vertex struct{
    X, Y int
}
 
func main(){
    v := Vertex{3, 4}
    fmt.Println(v.Plus())
}
Q2 X is 3! Y is 4! と表示されるStringerを作成してください。

package main
 
import (
    "fmt"
)
 
type Vertex struct{
    X, Y int
}
 
func main(){
    v := Vertex{3, 4}
    fmt.Println(v)
}

# ch05

Q1. 以下のように文字列を連結させて出力したいコードがありますが、

test1!
test1!test2!
test1!test2!test3!
test1!test2!test3!test4!
以下のコードには間違いがあります。上記の出力になるように修正してください。

package main
 
import "fmt"
 
func goroutine(s []string, c chan int){
    sum := ""
    for _, v := range s{
        sum += v
        c <- sum
    }
}
 
func main(){
    words := []string{"test1!", "test2!", "test3!", "test4!"}
    c := make(chan int)
    go goroutine(words, c)
    for w := range c{
        fmt.Println(w)
    }
}
