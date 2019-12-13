package try

import (
	"fmt"
	"math/rand"
	"time"
)

func Omikuji() {
	t := time.Now().UnixNano()
	rand.Seed(t)

	for i := 1; i <= 100; i = i+1 {
		n := rand.Intn(6)+1
		switch n {
		case 6: fmt.Println("大吉")
		case 5,4: fmt.Println("中吉")
		case 3,2: fmt.Println("吉")
		case 1: fmt.Println("凶")
		}
	}
}
