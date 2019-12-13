package try

import "fmt"

func Echo(msg interface{}) {
	fmt.Println(msg)
}

type Func func() string

func (f Func) String() string { return f() }

