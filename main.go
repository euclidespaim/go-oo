package main

import (
	"fmt"
	"go-orientado/contas"
)

func main() {
	contaDoElon := contas.ContaCorrente{
		Titular:       "Elon Musk",
		NumeroAgencia: 321,
		NumeroConta:   1234567,
		Saldo:         100.21,
	}

	contaDoJeff := contas.ContaCorrente{
		Titular:       "Jeff Bezos",
		NumeroAgencia: 123,
		NumeroConta:   7654321,
		Saldo:         100.21,
	}

	status := contaDoElon.Transferir(100, &contaDoJeff)

	fmt.Println(status)
	fmt.Println(contaDoJeff)
	fmt.Println(contaDoElon)

}
