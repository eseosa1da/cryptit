package main

import (
	"fmt"

	"github.com/eseosa1da/cryptit/decrypt"
	"github.com/eseosa1da/cryptit/encrypt"
)

func main() {

	fmt.Println(encrypt.Nimbus("Kodekloud"))

	fmt.Println(decrypt.Nimbus("Nrghnorxg"))

}
