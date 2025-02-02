package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Costanti
const (
	MAX_PERSONE   = 15
	NM            = 5
	N_COMMESSI    = 3
	N_CLIENTI     = 10
	N_CLIENTI_VIP = 3
)

// global
var commessi [N_COMMESSI]Commesso

// Funzione per dormire un tempo casuale
func sleepRandom(maxTime int) {
	time.Sleep(time.Duration(rand.Intn(maxTime)+1) * time.Second)
}

// Funzione per abilitare/disabilitare un canale in base a una condizione
func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func sumCommessi(commessi [N_COMMESSI]Commesso) int {
	sum := 0
	for i := 0; i < N_COMMESSI; i++ {
		if(commessi[i].isDentro){
		sum += commessi[i].clientiCount
		}
	}
	return sum
}

// Strutture dati
type Commesso struct {
	id           int
	isDentro	bool
	clientiCount int // Max 3
}

type Richiesta struct {
	id        int
	idCommesso int
}

// Canali
var entraCliente = make(chan int, 10)
var esciCliente = make(chan Richiesta, 10)
var ackCliente = make([]chan int, N_CLIENTI)

var entraClienteVIP = make(chan int, 10)
var esciClienteVIP = make(chan Richiesta, 10)
var ackClienteVIP = make([]chan int, N_CLIENTI_VIP)

var entraCommesso = make(chan int, 10)
var esciCommesso = make(chan int, 10)
var ackCommesso = make([]chan bool, N_COMMESSI)

var consegnaFornitore = make(chan int, 1)
var ackFornitore = make(chan bool)

var termina = make(chan bool)
var terminaNegozio = make(chan bool)

var done = make(chan bool)
var doneNegozio = make(chan bool)
var doneClienti = make(chan int)

// Routine del fornitore
func fornitore() {
	for {
		sleepRandom(15)
		consegnaFornitore <- 1

		select {
		case <-ackFornitore:
		case <-termina:
			fmt.Println("[Fornitore] termino")
			done <- true
			return
		}
	}
}

// Routine cliente
func cliente(id int) {
	fmt.Printf("[Cliente %d] avviato\n", id)
	sleepRandom(5)

	fmt.Printf("[Cliente %d] vorrei entrare...\n", id)
	entraCliente <- id
	idCommesso := <-ackCliente[id]

	sleepRandom(5)

	fmt.Printf("[Cliente %d] vorrei uscire...\n", id)
	esciCliente <- Richiesta{id, idCommesso}
	<-ackCliente[id]

	doneClienti <- id
	return
}

// Routine cliente VIP
func clienteVIP(id int) {
	fmt.Printf("[ClienteVIP %d] avviato\n", id)
	sleepRandom(5)


	fmt.Printf("[ClienteVIP %d] vorrei entrare...\n", id)
	entraClienteVIP <- id
	idCommesso := <-ackClienteVIP[id]

	sleepRandom(5)

	fmt.Printf("[ClienteVIP %d] vorrei uscire...\n", id)
	esciClienteVIP <- Richiesta{id, idCommesso}
	<-ackClienteVIP[id]

	doneClienti <- id
	return
}

// Routine commesso
func commesso(id int) {
	fmt.Printf("[Commesso %d] avviato, Stato: %v\n", id, commessi[id])

	for { 

		sleepRandom(5)

		if(!commessi[id].isDentro){
			fmt.Printf("[Commesso %d] vorrei entrare...\n", id)
			entraCommesso <- id
			 //bloccante, finchè il negozio decide che il commesso può entrare
			
			<- ackCommesso[id]
			sleepRandom(9)
		} else {
			fmt.Printf("[Commesso %d] vorrei uscire...\n", id)
			esciCommesso <- id
			 //bloccante, finchè il negozion decide che il commesso può uscire	
			
			<- ackCommesso[id]
			sleepRandom(9)
		}

		//select di terminazione
		select {
		case <-termina:
			fmt.Printf("[Commesso] termino\n")
			done <- true
			return
		default:
			sleepRandom(2)
		}
	}
}

// Routine negozio
func negozio() {
	// Inizializzazione variabili
	var personeCount, mascherineCount, commessiCount int = 0, 0, N_COMMESSI

	fmt.Println("[Negozio] inizializzazione commessi")
	for i := 0; i < N_COMMESSI; i++ {
		commessi[i] = Commesso{id: i, isDentro: false, clientiCount: 999}
	}

	for {
		select {
		case <-consegnaFornitore:
			mascherineCount += NM
			fmt.Printf("[Fornitore] consegna, numero mascherine: %d, coda in: %d, coda out: %d\n", mascherineCount, len(entraCliente)+len(entraClienteVIP), len(esciCliente)+len(esciClienteVIP))
			ackFornitore <- true

		case id := <-esciCommesso:
			if commessi[id].clientiCount == 0 && commessi[id].isDentro {
				commessiCount--
				personeCount--
				commessi[id].clientiCount = 999
				commessi[id].isDentro = false
				fmt.Printf("[Commesso %d] uscito\n", id)
				ackCommesso[id] <- true
			} else {
				ackCommesso[id] <- false
			}

		case exit := <-esciClienteVIP:
			personeCount--
			commessi[exit.idCommesso].clientiCount--
			fmt.Printf("[Cliente VIP %d] uscito, Commesso %d ha un posto in più\n", exit.id, exit.idCommesso)
			ackClienteVIP[exit.id] <- 42

		case exit := <-esciCliente:
			personeCount--
			commessi[exit.idCommesso].clientiCount--
			fmt.Printf("[Cliente %d] uscito, Commesso %d ha un posto in più\n", exit.id, exit.idCommesso)
			ackCliente[exit.id] <- 42

		case id := <-when(personeCount < MAX_PERSONE, entraCommesso):
			commessiCount++
			personeCount++
			fmt.Printf("[Commesso %d] entrato\n", id)
			commessi[id].clientiCount = 0
			commessi[id].isDentro = true
			ackCommesso[id] <- true

		case id := <-when(personeCount < MAX_PERSONE && mascherineCount > 0 && sumCommessi(commessi) < (commessiCount*3) && len(entraCommesso) == 0, entraClienteVIP):
//			fmt.Printf("[Negozio] Cliente VIP %d tenta di entrare: Persone: %d, Mascherine: %d, Commessi: %d, CodaCommessi: %d\n", id, personeCount, mascherineCount, commessiCount, len(entraCommesso))

			personeCount++
			mascherineCount--

			for i := 0; i < N_COMMESSI; i++ {
				if commessi[i].clientiCount < 3 && commessi[i].isDentro {
					commessi[i].clientiCount++
					fmt.Printf("[Cliente VIP %d] entrato, Commesso %d ha un posto in meno\n", id, i)
					ackClienteVIP[id] <- i
					break
				}
			}

		case id := <-when(personeCount < MAX_PERSONE && mascherineCount > 0 && sumCommessi(commessi) < (commessiCount*3) && len(entraCommesso) == 0 && len(entraClienteVIP) == 0, entraCliente):
//			fmt.Printf("[Negozio] Cliente %d tenta di entrare: Persone: %d, Mascherine: %d, Commessi: %d, CodaCommessi: %d, CodaVIP: %d\n", id, personeCount, mascherineCount, commessiCount, len(entraCommesso), len(entraClienteVIP))

			personeCount++
			mascherineCount--

			for i := 0; i < N_COMMESSI; i++ {
				if commessi[i].clientiCount < 3 && commessi[i].isDentro {
					commessi[i].clientiCount++
					fmt.Printf("[Cliente %d] entrato, Commesso %d ha un posto in meno\n", id, i)
					ackCliente[id] <- i
					break
				}
			}

		case <-terminaNegozio:
			fmt.Println("[Negozio] termino")

			for i := 0; i<N_COMMESSI+1; i++ {
				termina<-true
				<-done
				fmt.Printf("...k_%d\n", i)
			}
			doneNegozio <- true
			return
		}
	}
}

func main() {
	fmt.Println("Inizio")

	fmt.Println("Inizializzazione canali")
	for i := 0; i < N_COMMESSI; i++ {
		ackCommesso[i] = make(chan bool)
	}
	for i := 0; i < N_CLIENTI_VIP; i++ {
		ackClienteVIP[i] = make(chan int)
	}
	for i := 0; i < N_CLIENTI; i++ {
		ackCliente[i] = make(chan int)
	}


	sleepRandom(15)

	fmt.Println("Lancio le routines")
	go negozio()
	go fornitore()
	fmt.Println("Lanciati negozio e fornitore")

	for i := 0; i < N_CLIENTI; i++ {
		go cliente(i)
	}
	for i := 0; i < N_CLIENTI_VIP; i++ {
		go clienteVIP(i)
	}
	for i := 0; i < N_COMMESSI; i++ {
		go commesso(i)
	}

	var idArray [N_CLIENTI+N_CLIENTI_VIP]int

	for i := 0; i < N_CLIENTI_VIP; i++ {
		idArray[i] = i
	}	

	for i := 0; i < N_CLIENTI; i++{
		idArray[i+N_CLIENTI_VIP] = i
	}

	fmt.Println("Aspetto i figli")
	for i := 0; i < N_CLIENTI+N_CLIENTI_VIP; i++ { // aspetto la terminazione  delle routines cliente e clienteVIP
		id := <-doneClienti

		
		for i := 0; i < N_CLIENTI+N_CLIENTI_VIP; i++ {
			if(idArray[i]==id){
				idArray[i]=-1
				break
			}
		}
		fmt.Printf("clienti rimasti: ")
		for i := 0; i < N_CLIENTI+N_CLIENTI_VIP; i++ {
			if!(idArray[i]==-1){
				fmt.Printf("%d, ",idArray[i])
			}
		}
		fmt.Printf("\n")

		//fmt.Printf("\n...c_%d\n", i)
	}

	//terminazione negozio
	fmt.Println("Termino negozio")
	terminaNegozio <- true
	<- doneNegozio

}
