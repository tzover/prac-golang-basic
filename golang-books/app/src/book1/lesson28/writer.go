package book1Lesson28

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writen(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func funcProverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writen("Errors are values.")
	sw.writen("Don't just check errors, handle them gracefully.")
	sw.writen("Don't panic.")
	sw.writen("Make the zero value useful.")
	sw.writen("The bigger the interface, the weaker the abstraction.")
	sw.writen("interface{} says nothing.")
	sw.writen("Gofmt's style is no one's favorite," +
		" yet gofmt is everyone's favorite.")
	sw.writen("Documentation is for users.")
	sw.writen("A little copying is better than a little dependency.")
	sw.writen("Clear is better than clever")
	sw.writen("Concurrency is not parallelism")
	sw.writen("Don't communicate by sharing memory" +
		" share memory by communicating.")
	sw.writen("Channels orchestrate; mutexes serialize.")
	sw.writen("")

	return sw.err
}
