package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Costanti
const (
	MAX_PERSONE   = 5
	N_MASCHERINE  = 2
	N_COMMESSI    = 1
	N_CLIENTI     = 5
	N_CLIENTI_VIP = 1
)

type Commesso struct {
	id           int
	isDentro     bool
	clientiCount int
}

type Richiesta struct {
	id        int
	idCommesso int
}

var (
	entraCliente    = make(chan int, N_CLIENTI)
	esciCliente     = make(chan Richiesta, N_CLIENTI)
	ackCliente      = make([]chan int, N_CLIENTI)
	entraClienteVIP = make(chan int, N_CLIENTI_VIP)
	esciClienteVIP  = make(chan Richiesta, N_CLIENTI_VIP)
	ackClienteVIP   = make([]chan int, N_CLIENTI_VIP)
	entraCommesso   = make(chan int, N_COMMESSI)
	esciCommesso    = make(chan int, N_COMMESSI)
	ackCommesso     = make([]chan bool, N_COMMESSI)
	consegnaFornitore = make(chan struct{}, 1)
	ackFornitore    = make(chan struct{})
	terminaNegozio  = make(chan struct{})
	doneClienti     = make(chan struct{})
	doneNegozio     = make(chan struct{})
)

func sleepRandom(maxTime int) {
	time.Sleep(time.Duration(rand.Intn(maxTime)+1) * 10 * time.Millisecond)
}

func fornitore() {
	for {
		sleepRandom(15)
		consegnaFornitore <- struct{}{}
		<-ackFornitore
	}
}

func cliente(id int) {
	sleepRandom(10)
	entraCliente <- id
	<-ackCliente[id]
	sleepRandom(10)
	esciCliente <- Richiesta{id, id}
	<-ackCliente[id]
	doneClienti <- struct{}{}
}

func clienteVIP(id int) {
	sleepRandom(10)
	entraClienteVIP <- id
	<-ackClienteVIP[id]
	sleepRandom(10)
	esciClienteVIP <- Richiesta{id, id}
	<-ackClienteVIP[id]
	doneClienti <- struct{}{}
}

func commesso(id int) {
	for {
		sleepRandom(50)
		entraCommesso <- id
		<-ackCommesso[id]
		sleepRandom(50)
		esciCommesso <- id
		<-ackCommesso[id]
	}
}

func negozio() {
	clientiCount := 0
	mascherineCount := 0
	commessi := make([]Commesso, N_COMMESSI)

	for i := range commessi {
		commessi[i] = Commesso{id: i, isDentro: false, clientiCount: 0}
	}

	for {
		select {
		case <-consegnaFornitore:
			fmt.Println("[Negozio] - Fornitore consegna mascherine")
			mascherineCount += N_MASCHERINE
			ackFornitore <- struct{}{}

		case id := <-entraCommesso:
			fmt.Printf("[Negozio] - Commesso %d entrato\n", id)
			commessi[id].isDentro = true
			ackCommesso[id] <- true

		case id := <-esciCommesso:
			fmt.Printf("[Negozio] - Commesso %d uscito\n", id)
			commessi[id].isDentro = false
			ackCommesso[id] <- true

		case id := <-entraClienteVIP:
			if clientiCount < MAX_PERSONE && mascherineCount > 0 {
				fmt.Printf("[Negozio] - Cliente VIP %d entrato\n", id)
				mascherineCount--
				clientiCount++
				ackClienteVIP[id] <- id
			}

		case id := <-entraCliente:
			if clientiCount < MAX_PERSONE && mascherineCount > 0 {
				fmt.Printf("[Negozio] - Cliente %d entrato\n", id)
				mascherineCount--
				clientiCount++
				ackCliente[id] <- id
			}

		case <-terminaNegozio:
			fmt.Println("[Negozio] termina")
			doneNegozio <- struct{}{}
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := range ackCliente {
		ackCliente[i] = make(chan int)
	}
	for i := range ackClienteVIP {
		ackClienteVIP[i] = make(chan int)
	}
	for i := range ackCommesso {
		ackCommesso[i] = make(chan bool)
	}
	go fornitore()
	go negozio()
	for i := 0; i < N_CLIENTI; i++ {
		go cliente(i)
	}
	for i := 0; i < N_CLIENTI_VIP; i++ {
		go clienteVIP(i)
	}
	for i := 0; i < N_COMMESSI; i++ {
		go commesso(i)
	}
	for i := 0; i < N_CLIENTI+N_CLIENTI_VIP; i++ {
		<-doneClienti
	}
	terminaNegozio <- struct{}{}
	<-doneNegozio
}
