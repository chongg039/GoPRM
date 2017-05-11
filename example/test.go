package main

import (
	"log"
	"reflect"
	"time"
)

type test struct {
	name string
	num  time.Time
}

func main() {
	var si test

	log.Println(reflect.Zero(TypeOf(si)))
}
