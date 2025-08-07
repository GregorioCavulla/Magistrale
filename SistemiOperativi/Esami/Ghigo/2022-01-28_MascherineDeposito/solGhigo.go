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
	tipoLotto int // 0 chirurgica, 1 FFP2, 2 misto
}

// Funzione per dormire un tempo casuale di millisecondi
func sleepRandomMillis(max int){
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	time.Sleep(time.Duration(rand.Intn(max)) * time.Millisecond)
}

func sleepRandomMillisMin(min int,max int){
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	time.Sleep(time.Duration(rand.Intn(max - min) + min) * time.Millisecond)
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

	LOTTO_CHIRURGICHE = 3
	LOTTO_FFP2 = 3
	LOTTO_MISTO_CHIRURGICHE = 2
	LOTTO_MISTO_FFP2 = 1

	MAX_CHIRURGICHE = 25
	MAX_FFP2 = 20

	N_ADDETTI = 3
)

// Canali
var(
	
	iniziaRifChirurgiche = make(chan int)
	iniziaRifFFP2 = make(chan int)
	ackRifChirurgiche = make(chan bool)

	finisceRifChirurgiche = make(chan int)
	finisceRifFFP2 = make(chan int)
	ackRifFFP2 = make(chan bool)

	iniziaAddetto = make(chan int)
	finisceAddettoChirurgiche = make(chan Richiesta)
	finisceAddettoFFP2 = make(chan Richiesta)
	finisceAddettoMisto = make(chan Richiesta)

	ackAddetto[N_ADDETTI] chan int

	termina = make(chan bool)
	terminaDeposito = make(chan bool)

	done = make(chan bool)
	doneDeposito = make(chan bool)
)

// Routines

func rifornimentoChirurgiche() {

	for{

		sleepRandomMillis(150)

		iniziaRifChirurgiche <- 1

		select{
			case <- termina:
				fmt.Println("Terminazione rifornimento chirurgiche")
				done <- true
				return
			case <- ackRifChirurgiche:
		}

		sleepRandomMillisMin(200, 300)

		finisceRifChirurgiche <- 1

		select{
			case <- termina:
				fmt.Println("Terminazione rifornimento chirurgiche")
				done <- true
				return
			case <- ackRifChirurgiche:
				sleepRandomMillis(1)
		}

	}
}

func rifornimentoFFP2() {

	for{

		sleepRandomMillis(150)

		iniziaRifFFP2 <- 1

		select{
			case <- termina:
				fmt.Println("Terminazione rifornimento FFP2")
				done <- true
				return
			case <- ackRifFFP2:
				sleepRandomMillis(1)
		}

		sleepRandomMillisMin(200, 300)

		finisceRifFFP2 <- 1

		select{
		case <- termina:
			fmt.Println("Terminazione rifornimento FFP2")
			done <- true
			return
		case <- ackRifFFP2:
			sleepRandomMillis(1)
	}

	}
}

func addettoReparto(id int) {

	var tipoLotto int


		sleepRandomMillis(150)

		iniziaAddetto <- id

		select{
			case <- termina:
				fmt.Println("Terminazione addetto")
				done <- true
				return
			case tipoLotto = <- ackAddetto[id]:
				sleepRandomMillis(1)
		}

		sleepRandomMillisMin(200, 300)

		switch tipoLotto{
			case 0:
				finisceAddettoChirurgiche <- Richiesta{id, tipoLotto}
			case 1:
				finisceAddettoFFP2 <- Richiesta{id, tipoLotto}
			case 2:
				finisceAddettoMisto <- Richiesta{id, tipoLotto}
		}

		select{
			case <- termina:
				fmt.Println("Terminazione addetto")
				done <- true
				return
			case <- ackAddetto[id]:
				sleepRandomMillis(1)
		}

	done <- true
}

func precedenzaChirurgiche(countChirurgiche int, countFFP2 int) bool{
	if countChirurgiche == countFFP2{
		return false
	} else if countChirurgiche > countFFP2{
		return false
	} else if countChirurgiche < countFFP2{
		return true
	}else{
		return false
	}
}

func deposito(){

	var countChirurgiche int = 0
	var countFFP2 int = 0

	var freeChirurhiche bool = false
	var freeFFP2 bool = false

	var randTipo int

	for{

		select{
			case <- when(freeChirurhiche && precedenzaChirurgiche(countChirurgiche, countFFP2),iniziaRifChirurgiche):
				fmt.Println("Inizio rifornimento chirurgiche")
				freeChirurhiche = false
				ackRifChirurgiche <- true
			case <- finisceRifChirurgiche:
				fmt.Println("Fine rifornimento chirurgiche")
				freeChirurhiche = true
				countChirurgiche = MAX_CHIRURGICHE
				ackRifChirurgiche <- true
			case <- when(freeFFP2 && !precedenzaChirurgiche(countChirurgiche, countFFP2) ,iniziaRifFFP2):
				fmt.Println("Inizio rifornimento FFP2")
				freeFFP2 = false
				ackRifFFP2 <- true
			case <- finisceRifFFP2:
				fmt.Println("Fine rifornimento FFP2")
				freeFFP2 = true
				countFFP2 = MAX_FFP2
				ackRifFFP2 <- true
			case id := <- when(true, iniziaAddetto):
				fmt.Println("Inizio richiesta addetto", id)
				randTipo = rand.Intn(3)
				switch randTipo{
				case 0:
					freeChirurhiche = false
				case 1:
					freeFFP2 = false
				case 2:
					freeChirurhiche = false
					freeFFP2 = false
				}
				ackAddetto[id] <- randTipo
			case req := <- whenRichiesta(len(finisceAddettoMisto) == 0 && len(finisceAddettoFFP2) == 0 && freeChirurhiche && countChirurgiche > LOTTO_CHIRURGICHE, finisceAddettoChirurgiche):
				fmt.Println("Fine richiesta addetto chirurgiche", req.id)
				freeChirurhiche = true
				countChirurgiche -= LOTTO_CHIRURGICHE
				ackAddetto[req.id] <- 42
			case req := <- whenRichiesta(len(finisceAddettoMisto) == 0 && freeFFP2 && countFFP2 > LOTTO_FFP2, finisceAddettoFFP2):
				fmt.Println("Fine richiesta addetto FFP2", req.id)
				freeFFP2 = true
				countFFP2 -= LOTTO_FFP2
				ackAddetto[req.id] <- 42
			case req := <- whenRichiesta(freeChirurhiche && freeFFP2 && countChirurgiche > LOTTO_MISTO_CHIRURGICHE && countFFP2 > LOTTO_MISTO_FFP2, finisceAddettoMisto):
				fmt.Println("Fine richiesta addetto misto", req.id)
				freeChirurhiche = true
				freeFFP2 = true				
				countChirurgiche -= LOTTO_MISTO_CHIRURGICHE
				countFFP2 -= LOTTO_MISTO_FFP2
				ackAddetto[req.id] <- 42
			case <- terminaDeposito:
				fmt.Println("Terminazione deposito")
				doneDeposito <- true
				return
		}
	}
}

// Main

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Inizio esecuzione")

	// Inizializzazioni ack
	for i := 0; i < N_ADDETTI; i++{
		ackAddetto[i] = make(chan int)
	}

	// Avvio routines

	go rifornimentoChirurgiche()
	go rifornimentoFFP2()
	go deposito()

	for i := 0; i < N_ADDETTI; i ++ {
		go addettoReparto(i)
	}

	fmt.Println("\nTerminazione figli")
	for i := 0; i < N_ADDETTI; i++ {
		<- done
	}

	termina <- true
	<- done
	termina <- true
	<- done


	terminaDeposito <- true
	<- doneDeposito
	fmt.Println("Fine esecuzione")
}