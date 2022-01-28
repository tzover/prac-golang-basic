## install

```
https://golang.org/dl/
```

## test

```
go get golang.org/x/tools/cmd/godoc
```

## PATH

```
go env

export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOROOT/bin

go doc fmt
```

## GoLand(IDE)

- 有料

## Language

- Package Document

```
https://golang.org/pkg/

go doc fmt Println
```

- func init()

> main()より先に実行される

- Comment out

> `//` `/* */`

- Format

```
gofmt -w lesson.go
```
