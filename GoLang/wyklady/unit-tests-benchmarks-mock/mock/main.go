package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	greeter := Greeter{Doer: &HTTPDoer{}}
	//greeter := Greeter{Doer: &DoerMock{}}
	fmt.Printf("Hello %s!\n", greeter.What())
}

type Doer interface {
	Do() string
}

type Greeter struct {
	Doer Doer
}

type DoerMock struct {
	Ret string
}

func (m *DoerMock) Do() string {
	return m.Ret
}

type HTTPDoer struct{}

func (*HTTPDoer) Do() string {
	resp, err := http.Get("http://localhost:9999/hello")
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func (g *Greeter) What() string {
	return g.Doer.Do()
}
