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
	tipoParcheggio int
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

func sum(a int, b int) int {
	return a + b
}

// ## Inizio Esercizio ##

// Costanti
const (
	N_STANDARD = 15
	N_MAXI = 5

	N_AUTO = 25
	N_CAMPER = 10

	STANDARD = 0
	MAXI = 1
)

var tipoStr = [2]string{"Standard", "Maxi"} // 0 = Standard, 1 = Maxi

// Canali
var(
	entraAutoSalita = make(chan int, MAXBUFF)
	entraAutoDiscesa = make(chan Richiesta, MAXBUFF)
	esceAutoSalita = make(chan int, MAXBUFF)
	esceAutoDiscesa = make(chan int, MAXBUFF)

	ackAuto[N_AUTO] chan int

	entraCamperSalita = make(chan int, MAXBUFF)
	entraCamperDiscesa = make(chan int, MAXBUFF)
	esceCamperSalita = make(chan int, MAXBUFF)
	esceCamperDiscesa = make(chan int, MAXBUFF)

	ackCamper[N_CAMPER] chan bool

	entraSpazzaneveSalita = make(chan bool, MAXBUFF)
	entraSpazzaneveDiscesa = make(chan bool, MAXBUFF)
	esceSpazzaneveSalita = make(chan bool, MAXBUFF)
	esceSpazzaneveDiscesa = make(chan bool, MAXBUFF)

	ackSpazzaneve = make(chan bool)

	terminaSpazzaneve = make(chan bool)
	terminaStrada = make(chan bool)

	done = make(chan bool)
	doneStrada = make(chan bool)	
)

// Routines

func spazzaneve(){

	for{

		sleepRandomMillis(150)

		entraSpazzaneveDiscesa <- true
		select{
			case <- ackSpazzaneve:
			case <- terminaSpazzaneve:
				fmt.Println("[Spazzaneve] terminazione")
				done <- true
				return
		}

		sleepRandomMillis(100)

		esceSpazzaneveDiscesa <- true
		select{
			case <- ackSpazzaneve:
			case <- terminaSpazzaneve:
				fmt.Println("[Spazzaneve] terminazione")
				done <- true
				return
		}

		sleepRandomMillis(500)

		entraSpazzaneveSalita <- true
		select{
			case <- ackSpazzaneve:
			case <- terminaSpazzaneve:
				fmt.Println("[Spazzaneve] terminazione")
				done <- true
				return
		}

		sleepRandomMillis(70)

		esceSpazzaneveSalita <- true
		select{
			case <- ackSpazzaneve:
			case <- terminaSpazzaneve:
				fmt.Println("[Spazzaneve] terminazione")
				done <- true
				return
		}
	}
}

func auto(id int) {

	var req Richiesta

	var postoAssegnato int

	req.id = id

	sleepRandomMillis(50)

	entraAutoSalita <- id
	postoAssegnato = <- ackAuto[id]

	req.tipoParcheggio = postoAssegnato

	sleepRandomMillis(70)

	esceAutoSalita <- id
	<- ackAuto[id]

	sleepRandomMillis(300)

	entraAutoDiscesa <- req
	<- ackAuto[id]

	sleepRandomMillis(50)

	esceAutoDiscesa <- id
	<- ackAuto[id]

	done <- true
}

func camper(id int) {

	sleepRandomMillis(50)

	entraCamperSalita <- id
	<- ackCamper[id]

	sleepRandomMillis(85)

	esceCamperSalita <- id
	<- ackCamper[id]

	sleepRandomMillis(400)

	entraCamperDiscesa <- id
	<- ackCamper[id]

	sleepRandomMillis(60)

	esceCamperDiscesa <- id
	<- ackCamper[id]

	done <- true
}	

func strada() {

	var spazzaneveIn bool = false
	var countPostiStandard int = N_STANDARD
	var countPostiMaxi int = N_MAXI

	var camperInSalita bool = false
	var camperInDiscesa bool = false

	var postoAssegnato int = -1

	for{

		select{
			//spazzaneve
			case <- entraSpazzaneveDiscesa:
				fmt.Println("[Strada] entra spazzaneve discesa")
				spazzaneveIn = true
				ackSpazzaneve <- true
			case <- esceSpazzaneveDiscesa:
				fmt.Println("[Strada] esce spazzaneve discesa")
				spazzaneveIn = false
				ackSpazzaneve <- true
			case <- entraSpazzaneveSalita:
				fmt.Println("[Strada] entra spazzaneve salita")
				spazzaneveIn = true
				ackSpazzaneve <- true
			case <- esceSpazzaneveSalita:
				fmt.Println("[Strada] esce spazzaneve salita")
				spazzaneveIn = false
				ackSpazzaneve <- true
			//auto
			case id := <- when(sum(countPostiStandard, countPostiMaxi)<=sum(N_STANDARD, N_MAXI) && !camperInDiscesa && len(entraCamperSalita) == 0 && sum(len(entraCamperDiscesa),len(entraAutoDiscesa)) == 0 && !spazzaneveIn, entraAutoSalita): //se ci sono posti standard o maxi, non ci sono camper in discesa, i camper in salita sono 0 e non c'è lo spazzaneve 
				fmt.Printf("[Strada] entra auto %d salita\n", id)
				if countPostiStandard > 0 {
					countPostiStandard--
					postoAssegnato = STANDARD
				} else {
					countPostiMaxi--
					postoAssegnato = MAXI
				}
				ackAuto[id] <- postoAssegnato
			case req := <- whenRichiesta(!camperInSalita && len(entraCamperDiscesa) == 0 && !spazzaneveIn, entraAutoDiscesa): //se ci sono posti standard o maxi, non ci sono camper in salita, i camper in discesa sono 0 e non c'è lo spazzaneve
				fmt.Printf("[Strada] entra auto %d discesa\n", req.id)
				postoLiberato := req.tipoParcheggio
				if postoLiberato == STANDARD {
					countPostiStandard++
				} else {
					countPostiMaxi++
				}
				ackAuto[req.id] <- 42
			case id := <- esceAutoSalita: 
				fmt.Printf("[Strada] esce auto %d salita\n", id)
				ackAuto[id] <- 42
			case id := <- esceAutoDiscesa:
				fmt.Printf("[Strada] esce auto %d discesa\n", id)
				ackAuto[id] <- 42
			//camper
			case id := <- when(countPostiMaxi <= N_MAXI && !camperInDiscesa && sum(len(entraAutoDiscesa),len(entraCamperDiscesa)) == 0 && !spazzaneveIn, entraCamperSalita): // se ci sono posti maxi, non c'è lo spazzaneve e non ci sono camper in discesa,  
				fmt.Printf("[Strada] entra camper %d salita\n", id)
				countPostiMaxi--
				camperInSalita = true
				ackCamper[id] <- true
			case id := <- when(!camperInSalita && !spazzaneveIn, entraCamperDiscesa): // se ci sono posti maxi, non c'è lo spazzaneve e non ci sono camper in salita,
				fmt.Printf("[Strada] entra camper %d discesa\n", id)
				countPostiMaxi++
				camperInDiscesa = true
				ackCamper[id] <- true
			case id := <- esceCamperSalita:
				fmt.Printf("[Strada] esce camper %d salita\n", id)
				camperInSalita = false
				ackCamper[id] <- true
			case id := <- esceCamperDiscesa:
				fmt.Printf("[Strada] esce camper %d discesa\n", id)
				camperInDiscesa = false
				ackCamper[id] <- true
			case <-terminaStrada:

				fmt.Println("\nKillo spazzaneve\n")
				
				terminaSpazzaneve <- true
				<- done

				fmt.Println("\nKillo strada\n")

				doneStrada <- true
				return
		}
	}

}

// Main

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Inizio esecuzione")

	// Inizializzazioni
	for i := 0; i < N_AUTO; i++{
		ackAuto[i] = make(chan int)
	}

	for i := 0; i < N_CAMPER; i++{
		ackCamper[i] = make(chan bool)
	}

	// Avvio routines
	
	go strada()
	go spazzaneve()

	for i := 0; i < N_AUTO; i ++ {
		go auto(i)
	}

	for i := 0; i < N_CAMPER; i ++ {
		go camper(i)
	}

	// Attesa terminazione

	for i := 0; i < N_AUTO + N_CAMPER; i++ {
		<- done
	}

	terminaStrada <- true
	<- doneStrada

	fmt.Println("Fine esecuzione")
}