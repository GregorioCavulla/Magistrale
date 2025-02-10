package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXBUFF int = 100

// Funzione per abilitare/disabilitare un canale in base a una condizione
func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func sleepRandomMillis(max int){
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	time.Sleep(time.Duration(rand.Intn(max)) * time.Millisecond)
}


// Costanti
const(
	MAX_SALA = 60
	MAX_SORVEGLIANTI = 3
	MAX_CORRIDOIO = 30

	DIM_SCOLARESCA = 25

	N_SCOLARESCHE = 3
	N_VISITATORI = 10
	N_SORVEGLIANTI = 5
)

// Canali
var(
	entraInScolaresca = make(chan int, MAXBUFF)
	esceInScolaresca = make(chan int, MAXBUFF)

	entraOutScolaresca = make(chan int, MAXBUFF)
	esceOutScolaresca = make(chan int, MAXBUFF)

	ackScolaresca[N_SCOLARESCHE]chan int

	entraInVisitatore = make(chan int, MAXBUFF)
	esceInVisitatore = make(chan int, MAXBUFF)

	entraOutVisitatore = make(chan int, MAXBUFF)
	esceOutVisitatore = make(chan int, MAXBUFF)

	ackVisitatore[N_VISITATORI]chan int

	entraInSorvegliante = make(chan int, MAXBUFF)
	esceInSorvegliante = make(chan int, MAXBUFF)

 	entraOutSorvegliante = make(chan int, MAXBUFF)
	esceOutSorvegliante = make(chan int, MAXBUFF)

	ackSorvegliante[N_SORVEGLIANTI]chan int

	termina = make(chan int)
	done = make(chan int)
)

// Routines

func sorvegliante(id int) {

	for {
	
		sleepRandomMillis(20)

		entraInSorvegliante <- id

		select {
			case <-ackSorvegliante[id]:
			case <-termina:
				done <- 1
				return
		}

		sleepRandomMillis(20)
		
		esceInSorvegliante <- id

		select {
			case <-ackSorvegliante[id]:
			case <-termina:
				done <- 1
				return
		}

		sleepRandomMillis(20)

		entraOutSorvegliante <- id

		select {
			case <-ackSorvegliante[id]:
			case <-termina:
				done <- 1
				return
		}

		sleepRandomMillis(20)

		esceOutSorvegliante <- id

		select {
			case <-ackSorvegliante[id]:
			case <-termina:
				done <- 1
				return
		}
	}
}

func visitatore(id int) {

	sleepRandomMillis(20)

	entraInVisitatore <- id
	<- ackVisitatore[id]

	sleepRandomMillis(20)

	esceInVisitatore <- id
	<- ackVisitatore[id]

	sleepRandomMillis(20)

	entraOutVisitatore <- id
	<- ackVisitatore[id]

	sleepRandomMillis(20)

	esceOutVisitatore <- id
	<- ackVisitatore[id]

	done <- 1
}

func scolaresca(id int) {

	sleepRandomMillis(20)

	entraInScolaresca <- id
	<- ackScolaresca[id]

	sleepRandomMillis(20)

	esceInScolaresca <- id
	<- ackScolaresca[id]

	sleepRandomMillis(20)

	entraOutScolaresca <- id
	<- ackScolaresca[id]

	sleepRandomMillis(20)

	esceOutScolaresca <- id
	<- ackScolaresca[id]

	done <- 1
}

func server() {

	//inizializzazione

	var countPersone int = 0
	var countSorveglianti int = 0

	var countTransitoIn int = 0
	var countTransitoOut int = 0

	var scolarescaIN int = 0
	var scolarescaOUT int = 0

	for {
		select{
		// esce dal corridoio
		// direzione fuori
			// esceOutScolaresca
		case id := <- esceOutScolaresca:
			scolarescaOUT -= 1
			countTransitoOut -= DIM_SCOLARESCA
			fmt.Printf("[Scolaresca][%d] esce OUT\n", id)
			ackScolaresca[id] <- 1
			// esceOutVisitatore
		case id := <- esceOutVisitatore:
			countTransitoOut -= 1
			fmt.Printf("[Visitatore][%d] esce OUT\n", id)
			ackVisitatore[id] <- 1
			// esceOutSorvegliante
		case id := <- esceOutSorvegliante:
			countTransitoOut -= 1
			fmt.Printf("[Sorvegliante][%d] esce OUT\n", id)
			ackSorvegliante[id] <- 1

		// direzione dentro
			// esceInSorvegliante
		case id := <- esceInSorvegliante:
			countTransitoIn -= 1
			fmt.Printf("[Sorvegliante][%d] esce IN\n", id)
			ackSorvegliante[id] <- 1
			// esceInVisitatore
		case id := <- esceInVisitatore:
			countTransitoIn -= 1
			fmt.Printf("[Visitatore][%d] esce IN\n", id)
			ackVisitatore[id] <- 1
			// esceInScolaresca
		case id := <- esceInScolaresca:
			scolarescaIN -= 1
			countTransitoIn -= DIM_SCOLARESCA
			fmt.Printf("[Scolaresca][%d] esce IN\n", id)
			ackScolaresca[id] <- 1

		// entra nel corridoio
		// direzione fuori
			// entraOutScolaresca
		case id := <- when(countTransitoOut+25 <= MAX_SALA && countTransitoIn == 0, entraOutScolaresca):
			scolarescaOUT += 1
			countTransitoOut += DIM_SCOLARESCA
			countPersone -= DIM_SCOLARESCA
			fmt.Printf("[Scolaresca][%d] entra OUT\n", id)
			ackScolaresca[id] <- 1			
			// entraOutVisitatore
		case id := <- when(scolarescaIN == 0 && countTransitoIn+countTransitoOut < MAX_CORRIDOIO && len(entraOutScolaresca) == 0, entraOutVisitatore):
			countTransitoOut += 1
			countPersone -= 1
			fmt.Printf("[Visitatore][%d] entra OUT\n", id)
			ackVisitatore[id] <- 1
			// entraOutSorvegliante
		case id := <- when(scolarescaIN == 0 && countTransitoIn+countTransitoOut < MAX_CORRIDOIO && len(entraOutScolaresca) == 0 && len(entraOutVisitatore) == 0 && countPersone-countSorveglianti == 0, entraOutSorvegliante):
			countTransitoOut += 1
			countPersone -= 1
			countSorveglianti -= 1
			fmt.Printf("[Sorvegliante][%d] entra OUT\n", id)
			ackSorvegliante[id] <- 1

		// direzione dentro
			// entraInSorvegliante
		case id := <- when(countTransitoIn + countTransitoOut < MAX_CORRIDOIO && scolarescaOUT == 0 && countPersone < MAX_SALA && countSorveglianti < MAX_SORVEGLIANTI && len(entraOutScolaresca) == 0 && len(entraOutSorvegliante) == 0 && len(entraOutVisitatore) == 0, entraInSorvegliante):
			countTransitoIn += 1
			countSorveglianti += 1
			countPersone += 1
			fmt.Printf("[Sorvegliante][%d] entra IN\n", id)
			ackSorvegliante[id] <- 1
			// entraInVisitatore
		case id := <- when(countTransitoIn + countTransitoOut < MAX_CORRIDOIO && scolarescaOUT == 0 && countPersone < MAX_SALA && countSorveglianti > 0 && len(entraInSorvegliante) == 0 && len(entraOutScolaresca) == 0 && len(entraOutSorvegliante) == 0 && len(entraOutVisitatore) == 0, entraInVisitatore):
			countTransitoIn += 1
			countPersone += 1
			fmt.Printf("[Visitatore][%d] entra IN\n", id)
			ackVisitatore[id] <- 1			
			// entraInScolaresca
		case id := <- when(countTransitoIn + countTransitoOut < MAX_CORRIDOIO && scolarescaOUT == 0 && countPersone < MAX_SALA && countSorveglianti > 0 && len(entraInSorvegliante) == 0 && len(entraInVisitatore) == 0 && len(entraOutScolaresca) == 0 && len(entraOutSorvegliante) == 0 && len(entraOutVisitatore) == 0, entraInScolaresca):
			scolarescaIN += 1
			countTransitoIn += DIM_SCOLARESCA
			countPersone += DIM_SCOLARESCA
			fmt.Printf("[Scolaresca][%d] entra IN\n", id)
			ackScolaresca[id] <- 1

		// termina
		case <- termina:
			done <- 1
			return
				
		}
	}

}

// Main

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Inizio esecuzione")

	// Inizializzazioni

	// ack
	for i := 0; i < N_SCOLARESCHE; i++ {
		ackScolaresca[i] = make(chan int)
	}

	for i := 0; i < N_VISITATORI; i++ {
		ackVisitatore[i] = make(chan int)
	}

	for i := 0; i < N_SORVEGLIANTI; i++ {
		ackSorvegliante[i] = make(chan int)
	}

	// Avvio routines

	go server()

	for i := 0; i < N_SORVEGLIANTI; i++ {
		go sorvegliante(i)
	}

	for i := 0; i < N_VISITATORI; i++ {
		go visitatore(i)
	}

	for i := 0; i < N_SCOLARESCHE; i++ {
		go scolaresca(i)
	}

	// Attesa terminazione

	for i := 0; i < N_SCOLARESCHE+N_VISITATORI; i++ {
		<-done
	}

	for i := 0; i < N_SORVEGLIANTI+1; i++ {
		termina <- 1
	}

	for i := 0; i < N_SORVEGLIANTI+1; i++ {
		<-done
	}

	fmt.Println("Fine esecuzione")
}