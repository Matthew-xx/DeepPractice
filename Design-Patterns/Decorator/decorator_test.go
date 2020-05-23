package Decorator

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPi(t *testing.T)  {
	fmt.Println(Pi(5000))
}

func TestWrapLogger(t *testing.T) {
	f := WrapLogger(Pi,log.New(os.Stdout,"test:",1))
	f(10000)
}
