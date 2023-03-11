package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type pythagoras struct {
	a, b, c int
	m       sync.Mutex
}

func (p *pythagoras) calculate() {
	p.m.Lock()
	p.c = p.a*p.a + p.b*p.b
	p.c = int(math.Sqrt(float64(p.c)))
	p.m.Unlock()
	fmt.Println(p.c)
}

func main() {
	p := pythagoras{a: 3, b: 4}
	go p.calculate()
	time.Sleep(3 * time.Second)
}
