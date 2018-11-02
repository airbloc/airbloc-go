package main

import (
	"log"

	"github.com/pkg/errors"
)

//go:generate ./build.sh

func init() {

}

func a() error {
	return errors.New("hello world!")
}

func b() error {
	return errors.Wrap(a(), "rehello!")
}

func main() {
	log.Println((b()))
	log.Println(len([]byte("4957744b3ac54434b8270f2c854cc1040228c82ea4e72d66d2887a4d3e30b317")))
}
