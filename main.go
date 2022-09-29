package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoElon := ContaCorrente{
		titular: "Elon Musk", 
		numeroAgencia: 321, 
		numeroConta: 1234567, 
		saldo: 100.21,
	}
	fmt.Println(contaDoElon)

	contaDoBezos := ContaCorrente{"Jeff Bezos", 123, 1234567, 200.00}
	fmt.Println(contaDoBezos)

	var contaDoStark *ContaCorrente
	contaDoStark = new(ContaCorrente)
	contaDoStark.titular = "Tony Stark"
	contaDoStark.numeroAgencia = 123
	fmt.Println(contaDoStark)
}
