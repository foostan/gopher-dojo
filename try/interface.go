package try

import "fmt"

type I int

func (i I) String() string {
	return string(i)
}

type S string

func (s S) String() string {
	return string(s)
}

type B bool

func (b B) String() string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func PrintType(s Stringer) {
	switch s.(type) {
	case I: fmt.Println("I " + s.String())
	case S: fmt.Println("S " + s.String())
	case B: fmt.Println("B " + s.String())
	}
}
