package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(acc)

	if err := store.CreateAccount(acc); err != nil {
		println("conta no Create Account => ", acc.BankNumber)
		println("encrypted pass Create Account => ", acc.EncryptedPassword)

		log.Fatal("Erro ao criar a conta => ", err)
	}

	fmt.Println("Conta criada => ", acc.BankNumber)

	return acc
}

func seedAccounts(s Storage) {

	seedAccount(s, "victor", "reis", "hunter9999")
}

func main() {
	seed := flag.Bool("seed", false, "seed do DB")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Fazendo seed do banco")

		// coisas para fazer o seed
		seedAccounts(store)
	}

	server := NewApiServer(":3000", store)
	server.Run()
}
