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
	MAX_PZFIZER = 200 // frigo a -70
	MAX_MODERNA = 200 // frigo a -8

	N_RIFORNIMENTO = 100 // dosi per rifornimento

	TOT_PZFIZER = 2400 // dosi totali prima di terminare
	TOT_MODERNA = 2400 // dosi totali prima di terminare

	N_DOSI_PRELIEVO = 30 // dosi per prelievo

	N_REGIONI_ROSSE = 3
	N_REGIONI_ARANCIONI = 6
	N_REGIONI_GIALLE = 11
)

// Canali
var (
	depositaPfizer = make(chan int, MAXBUFF)
	ackPfizer = make(chan bool)

	depositaModerna = make(chan int, MAXBUFF)
	ackModerna = make(chan bool)

	regioneRossaPreleva = make(chan int, MAXBUFF)
	ackRegioneRossa [N_REGIONI_ROSSE]chan bool

	regioneArancionePreleva = make(chan int, MAXBUFF)
	ackRegioneArancione [N_REGIONI_ARANCIONI]chan bool

	regioneGiallaPreleva = make(chan int, MAXBUFF)
	ackRegioneGialla [N_REGIONI_GIALLE]chan bool

	terminaPfizer = make(chan bool)
	terminaModerna = make(chan bool)
	terminaRegioneRossa = make(chan bool)
	terminaRegioneArancione = make(chan bool)
	terminaRegioneGialla = make(chan bool)
	terminaDeposito = make(chan bool)

	done = make(chan bool)
	doneDeposito = make(chan bool)
)

func vacciniDisponibili (countPfizer int, countModerna int) bool {
	return countPfizer >= N_DOSI_PRELIEVO || countModerna >= N_DOSI_PRELIEVO || countPfizer + countModerna >= N_DOSI_PRELIEVO
}

// Routines

func pfizer() {

	for{

		sleepRandomMillis(150)

		depositaPfizer <- 1

		select{
			case <- ackPfizer:
			case <- terminaPfizer:
				fmt.Println("Pfizer termina")
				done <- true
				return
		}
	}	
}

func moderna() {

	for{

		sleepRandomMillis(150)

		depositaModerna <- 1

		select{
			case <- ackModerna:
			case <- terminaModerna:
				fmt.Println("Moderna termina")
				done <- true
				return
		}
	}	
}

func regioneRossa(id int) {

	for{

		regioneRossaPreleva <- id

		select{
			case <- ackRegioneRossa[id]:
			case <- terminaRegioneRossa:
				fmt.Println("Regione Rossa termina")
				done <- true
				return
		}
	
		sleepRandomMillis(500)
	}	
}

func regioneArancione(id int) {
	
	for{

		regioneArancionePreleva <- id

		select{
			case <- ackRegioneArancione[id]:
			case <- terminaRegioneArancione:
				fmt.Println("Regione Arancione termina")
				done <- true
				return
		}
	
		sleepRandomMillis(500)
	}	
}

func regioneGialla(id int) {
	
	for{

		regioneGiallaPreleva <- id

		select{
			case <- ackRegioneGialla[id]:
			case <- terminaRegioneGialla:
				fmt.Println("Regione Gialla termina")
				done <- true
				return
		}
	
		sleepRandomMillis(500)
	}	
}

func deposito(){

	var disponibiliPfizer int = 0
	var disponibiliModerna int = 0

	var consegnatiPfizer int = 0
	var consegnatiModerna int = 0

	for{
		if( consegnatiPfizer >= TOT_PZFIZER && consegnatiModerna >= TOT_MODERNA){
			break
		}
		select{
			case <- when(disponibiliPfizer + N_RIFORNIMENTO < MAX_PZFIZER, depositaPfizer):
				fmt.Println("[Deposito] Stato: disponibiliPfizer", disponibiliPfizer, "disponibiliModerna", disponibiliModerna, "consegnatiPfizer", consegnatiPfizer, "consegnatiModerna", consegnatiModerna)
				disponibiliPfizer += N_RIFORNIMENTO
				consegnatiPfizer += N_RIFORNIMENTO
				fmt.Println("[Deposito] deposita Pfizer")
				ackPfizer <- true
			case <- when(disponibiliModerna + N_RIFORNIMENTO < MAX_MODERNA, depositaModerna):
				fmt.Println("[Deposito] Stato: disponibiliPfizer", disponibiliPfizer, "disponibiliModerna", disponibiliModerna, "consegnatiPfizer", consegnatiPfizer, "consegnatiModerna", consegnatiModerna)
				disponibiliModerna += N_RIFORNIMENTO
				consegnatiModerna += N_RIFORNIMENTO
				fmt.Println("[Deposito] deposita Moderna")
				ackModerna <- true
			case id := <- when(vacciniDisponibili(disponibiliPfizer, disponibiliModerna), regioneRossaPreleva):
				if disponibiliPfizer >= N_DOSI_PRELIEVO{
					disponibiliPfizer -= N_DOSI_PRELIEVO
					fmt.Printf("[Deposito] Regione Rossa %d preleva Pfizer\n", id)
					ackRegioneRossa[id] <- true
				} else if disponibiliModerna >= N_DOSI_PRELIEVO{
					disponibiliModerna -= N_DOSI_PRELIEVO
					fmt.Printf("[Deposito] Regione Rossa %d preleva Moderna\n", id)
					ackRegioneRossa[id] <- true
				} else {
					var temp int = N_DOSI_PRELIEVO - disponibiliPfizer
					disponibiliPfizer = 0
					disponibiliModerna -= temp
					fmt.Printf("[Deposito] Regione Rossa %d preleva Pfizer e Moderna\n", id)
				}
			case id := <- when(vacciniDisponibili(disponibiliPfizer, disponibiliModerna) && len(regioneRossaPreleva) == 0, regioneArancionePreleva):
				if disponibiliPfizer >= N_DOSI_PRELIEVO{
					disponibiliPfizer -= N_DOSI_PRELIEVO
					fmt.Printf("[Deposito] Regione Arancione %d preleva Pfizer\n", id)
					ackRegioneArancione[id] <- true
				} else if disponibiliModerna >= N_DOSI_PRELIEVO{
					disponibiliModerna -= N_DOSI_PRELIEVO
					fmt.Printf("[Deposito] Regione Arancione %d preleva Moderna\n", id)
					ackRegioneArancione[id] <- true
				} else {
					var temp int = N_DOSI_PRELIEVO - disponibiliPfizer
					disponibiliPfizer = 0
					disponibiliModerna -= temp
					fmt.Printf("[Deposito] Regione Arancione %d preleva Pfizer e Moderna\n", id)
				}
			case id := <- when(vacciniDisponibili(disponibiliPfizer, disponibiliModerna) && len(regioneRossaPreleva) == 0 && len(regioneArancionePreleva) == 0, regioneGiallaPreleva):
				if disponibiliPfizer >= N_DOSI_PRELIEVO{
					disponibiliPfizer -= N_DOSI_PRELIEVO
					fmt.Printf("[Deposito] Regione Gialla %d preleva Pfizer\n", id)
					ackRegioneGialla[id] <- true
				} else if disponibiliModerna >= N_DOSI_PRELIEVO{
					disponibiliModerna -= N_DOSI_PRELIEVO
					fmt.Printf("[Deposito] Regione Gialla %d preleva Moderna\n", id)
					ackRegioneGialla[id] <- true
				} else {
					var temp int = N_DOSI_PRELIEVO - disponibiliPfizer
					disponibiliPfizer = 0
					disponibiliModerna -= temp
					fmt.Printf("[Deposito] Regione Gialla %d preleva Pfizer e Moderna\n", id)
				}
			}
		}
	fmt.Println("Deposito termina")
	doneDeposito <- true
}

// Main

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Inizio esecuzione")

	// Inizializzazioni
	for i := 0; i < N_REGIONI_ROSSE; i++{
		ackRegioneRossa[i] = make(chan bool)
	}

	for i := 0; i < N_REGIONI_ARANCIONI; i++{
		ackRegioneArancione[i] = make(chan bool)
	}

	for i := 0; i < N_REGIONI_GIALLE; i++{
		ackRegioneGialla[i] = make(chan bool)
	}

	// Avvio routines
	go deposito()
	go pfizer()
	go moderna()

	for i := 0; i < N_REGIONI_ROSSE; i ++ {
		go regioneRossa(i)
	}

	for i := 0; i < N_REGIONI_ARANCIONI; i ++ {
		go regioneArancione(i)
	}

	for i := 0; i < N_REGIONI_GIALLE; i ++ {
		go regioneGialla(i)
	}

	// Attesa terminazione deposito
	<- doneDeposito

	fmt.Println("Terminazione routines")

	terminaPfizer <- true
	<- done
	terminaModerna <- true
	<- done

	for i := 0; i < N_REGIONI_ROSSE; i++{
		terminaRegioneRossa <- true
	}

	for i := 0; i < N_REGIONI_ROSSE; i++{
		<- done
	}

	for i := 0; i < N_REGIONI_ARANCIONI; i++{
		terminaRegioneArancione <- true
	}

	for i := 0; i < N_REGIONI_ARANCIONI; i++{
		<- done
	}

	for i := 0; i < N_REGIONI_GIALLE; i++{
		terminaRegioneGialla <- true
	}

	for i := 0; i < N_REGIONI_GIALLE; i++{
		<- done
	}
	
	fmt.Println("Fine esecuzione")
}
