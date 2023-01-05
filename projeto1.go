package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	temperaturaK := ebulicaoK
	temperaturaC := temperaturaK - 273.0

	fmt.Printf("A temperatura de ebulição da agua em °K = %g , A temperatura de ebulição da agua em °C = %g .", temperaturaK, temperaturaC)
}
