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

)

// Struttura
type Struttura struct {

}

// Canali
var(
	

	terminaRoutineCiclica = make(chan bool)
	terminaServer = make(chan bool)

	done = make(chan bool)
	doneServer = make(chan bool)
)

// Routines

func routineCiclica() {

	for{

		sleepRandomMillis(150)

		select{
			case <- terminaRoutine:
				done <- true
				return
			case default:
				sleepRandomMillis(1)
		}

	}
}

func routine() {

	sleepRandomMillis(50)

	//fine
	done <- true
}

func server(){

	for{

		select{
			case <- terminaServer:
				//termino tutti i figli che hanno un ciclo infito
				fmt.Println("\nTerminazione figli")

				for i := 0; i < N; i++ {
					terminaRoutineCiclica <- true
					<- done
				}

				fmt.Println("Terminazione server")
				doneServer <- true
				return
	}
}

// Main

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Inizio esecuzione")

	// Inizializzazioni ack

	// Avvio routines

	// Attesa terminazione

	// Terminazione server
	terminaServer <- true
	<- doneServer

	fmt.Println("Fine esecuzione")
}