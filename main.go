package main

import (
	"fmt"
	"go-orientado/clientes"
	"go-orientado/contas"
)

func main() {
	contaDoElon := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Elon Musk",
			CPF:       "123.456.789-10",
			Profissao: "Programador",
		},
		NumeroAgencia: 321,
		NumeroConta:   1234567,
		Saldo:         100.21,
	}

	contaDoJeff := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Jeff Bezos",
			CPF:       "123.456.789-11",
			Profissao: "Front-end",
		},
		NumeroAgencia: 123,
		NumeroConta:   7654321,
		Saldo:         100.21,
	}

	status := contaDoElon.Transferir(100, &contaDoJeff)

	fmt.Println(status)
	fmt.Println(contaDoJeff)
	fmt.Println(contaDoElon)

}
