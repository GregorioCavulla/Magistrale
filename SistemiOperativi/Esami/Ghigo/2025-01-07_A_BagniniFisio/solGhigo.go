// ## Standard Per Esame ##

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXBUFF int = 100

// Richiesta, id per identificare la richiesta, ack per la risposta
type Richiesta struct {
	id  int
	nVolte int
	// altro
}

// Funzione per dormire un tempo casuale di millisecondi
func sleepRandomMillis(max int){
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	time.Sleep(time.Duration(rand.Intn(max)) * time.Millisecond)
}

// Funzione per abilitare/disabilitare un canale in base a una condizione
func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

// Funzione per abilitare/disabilitare un canale in base a una condizione
func whenRichiesta(b bool, c chan Richiesta) chan Richiesta {
	if !b {
		return nil
	}
	return c
}

// ## Inizio Esercizio ##

// Costanti
const (

	MAX_UTENTI = 10
	N_FISIO = 3
	N_BAGNINI = 2
	N_UTENTI = 25

)


// Canali
var(

	entraBagnino = make(chan int, MAXBUFF)
	esceBagnino = make(chan int, MAXBUFF)
	ackBagnino [N_BAGNINI]chan bool	

	entraUtenteFun = make(chan Richiesta, MAXBUFF)
	entraUtenteFisio = make(chan Richiesta, MAXBUFF)
	esceUtenteFun = make(chan int, MAXBUFF)
	esceUtenteFisio = make(chan int, MAXBUFF)
	ackUtente [N_UTENTI]chan bool

	terminaBagnino = make(chan bool)
	terminaServer = make(chan bool)

	done = make(chan bool)
	doneServer = make(chan bool)
)

// Routines

func bagnino(id int) {

	for{
		sleepRandomMillis(150)

		entraBagnino <- id

		select{
			case <- terminaBagnino:
				fmt.Println("[Bagnino", id, "] Terminazione")
				done <- true
				return
			case <- ackBagnino[id]:
				sleepRandomMillis(1000)
		}

		esceBagnino <- id

		select{
		case <- terminaBagnino:
			fmt.Println("[Bagnino", id, "] Terminazione")
			done <- true
			return
		case <- ackBagnino[id]:
			sleepRandomMillis(1)
		}
	}
}

func utente(id int) {

	sleepRandomMillis(50)

	var percorso int = rand.Intn(3)
	var nVolte int = rand.Intn(5) + 1

	for i := 0; i < nVolte; i++ {
		req := Richiesta{id, i}
		switch percorso {
		case 0: // solo fun
			entraUtenteFun <- req
			<- ackUtente[id]
			sleepRandomMillis(100)
			esceUtenteFun <- id
			<- ackUtente[id]
		case 1: // solo fisio
			entraUtenteFisio <- req
			<- ackUtente[id]
			sleepRandomMillis(100)
			esceUtenteFisio <- id
			<- ackUtente[id]
		case 2: // entrambi
			entraUtenteFun <- req
			<- ackUtente[id]
			sleepRandomMillis(100)
			esceUtenteFun <- id
			<- ackUtente[id]
			entraUtenteFisio <- req
			<- ackUtente[id]
			sleepRandomMillis(100)
			esceUtenteFisio <- id
			<- ackUtente[id]
		}
	}
	
	//fine
	done <- true
}

func centro(){

	var countBagnini = 0
	var countUtenti = 0

	var countFisioLiberi = N_FISIO

	for{

		select{
			case req := <- whenRichiesta(countUtenti < MAX_UTENTI && countBagnini > 0, entraUtenteFun):
				fmt.Println("[Centro] Utente", req.id, "entra in fun, per la ", req.nVolte, "volta")
				countUtenti++
				ackUtente[req.id] <- true
			case req := <- whenRichiesta(countUtenti < MAX_UTENTI && countFisioLiberi > 0, entraUtenteFisio):
				fmt.Println("[Centro] Utente", req.id, "entra in fisio, per la ", req.nVolte, "volta")
				countUtenti++
				countFisioLiberi--
				ackUtente[req.id] <- true
			case id := <- esceUtenteFun:
				fmt.Println("[Centro] Utente", id, "esce da fun")
				countUtenti--
				ackUtente[id] <- true
			case id := <- esceUtenteFisio:
				fmt.Println("[Centro] Utente", id, "esce da fisio")
				countUtenti--
				countFisioLiberi++
				ackUtente[id] <- true
			case id := <- entraBagnino:
				fmt.Println("[Centro] Bagnino", id, "entra")
				countBagnini++
				ackBagnino[id] <- true
			case id := <- when(countUtenti == 0 || countBagnini > 1 ,esceBagnino):
				fmt.Println("[Centro] Bagnino", id, "esce")
				countBagnini--
				ackBagnino[id] <- true
			case <- terminaServer:
				//termino tutti i figli che hanno un ciclo infito
				fmt.Println("\n[Centro] Terminazione Bagnini")

				for i := 0; i < N_BAGNINI; i++ {
					terminaBagnino <- true
					<- done
				}

				fmt.Println("[Centro] Terminazione server")
				doneServer <- true
				return
		}
	}
}

// Main

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Inizio esecuzione")

	// Inizializzazioni ack
	for i := 0; i < N_BAGNINI; i++ {
		ackBagnino[i] = make(chan bool)
	}

	for i := 0; i < N_UTENTI; i++ {
		ackUtente[i] = make(chan bool)
	}


	// Avvio routines
	go centro()

	for i := 0; i < N_BAGNINI; i++ {
		go bagnino(i)
	}

	for i := 0; i < N_UTENTI; i++ {
		go utente(i)
	}

	// Attesa terminazione

	for i := 0; i < N_UTENTI; i++ {
		<- done
	}

	// Terminazione server
	terminaServer <- true
	<- doneServer

	fmt.Println("Fine esecuzione")
}