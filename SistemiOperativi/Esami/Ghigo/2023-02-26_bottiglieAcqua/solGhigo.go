// ## Standard Per Esame ##

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXBUFF int = 100

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

func condizioneAddetto(countLitri float64, countDieci int, countVenti int, lenPiccola int, lenGrande int) bool {
	if countLitri == 0 || countDieci == MAX_DIECI || countVenti == MAX_VENTI { // se il serbatoio è vuoto o una cassetta è piena ho la precedenza
		return true
	} else if lenPiccola == 0 && lenGrande == 0 { // se non ci sono richieste in coda ho la precedenza
		return true
	} else { // altrimenti non ho la precedenza
		return false
	}
}

// Costanti
const (

	MAX_DIECI = 90
	MAX_VENTI = 80
	LITRI = 80

	N_PICCOLA = 100
	N_GRANDE = 100
)

// Struttura
type Struttura struct {

}

// Canali
var(
	addettoInizia = make(chan int, MAXBUFF)
	addettoFinisce = make(chan int, MAXBUFF)
	ackAddetto = make(chan bool)

	piccolaRiempie = make(chan int, MAXBUFF)
	piccolaFinisce = make(chan int, MAXBUFF)
	ackPiccola[N_PICCOLA] chan bool

	grandeRiempie = make(chan int, MAXBUFF)
	grandeFinisce = make(chan int, MAXBUFF)
	ackGrande[N_GRANDE] chan bool

	terminaAddetto = make(chan bool)
	terminaServer = make(chan bool)

	done = make(chan bool)
	doneServer = make(chan bool)
)

// Routines

func addetto() {

	for{

		sleepRandomMillis(150)

		addettoInizia <- 42

		select{
			case <- ackAddetto:
			case <- terminaAddetto:
				fmt.Println("Addetto terminazione")
				done <- true
				return
		}

		sleepRandomMillisMin(200, 300)

		addettoFinisce <- 42

		select{
			case <- ackAddetto:
			case <- terminaAddetto:
				fmt.Println("Addetto terminazione")
				done <- true
				return
		}

	}
}

func grande(id int) {

	sleepRandomMillis(50)

	grandeRiempie <- id
	<- ackGrande[id]

	sleepRandomMillisMin(150, 250)

	grandeFinisce <- id
	<- ackGrande[id]

	//fine
	done <- true
}

func piccola(id int) {

	sleepRandomMillis(50)

	piccolaRiempie <- id
	<- ackPiccola[id]

	sleepRandomMillisMin(150, 250)

	piccolaFinisce <- id
	<- ackPiccola[id]

	//fine
	done <- true
}

func server(){

	var countLitri float64 = 0 // litri rimasti, start at: 250
	var countDieci int = 0 // cassetta da 10, max: 150
	var countVenti int = 0 // cassetta da 20, max: 200

	var isOccupato bool = false

	for{

		select{
			case <- when(condizioneAddetto(countLitri,countDieci,countVenti,len(piccolaRiempie),len(grandeRiempie)) && !isOccupato ,addettoInizia): // condizione da verificare: se il serbatoio è vuoto o una cassetta è piena ho la precedenza, altrimenti no
				fmt.Printf("Addetto inizia\n")
				isOccupato = true
				ackAddetto <- true
			case <- addettoFinisce:
				fmt.Printf("Addetto finisce\n")
				isOccupato = false
				countLitri = LITRI // riempio il serbatoio
				countDieci = 0 // svuoto la cassetta da 10
				countVenti = 0 // svuoto la cassetta da 20
				ackAddetto <- true			
			case id := <- when(countLitri >= 0.5 && countDieci < MAX_DIECI && !isOccupato, piccolaRiempie):
				fmt.Printf("Piccola %d riempie\n", id)
				isOccupato = true
				ackPiccola[id] <- true
			case id := <- piccolaFinisce:
				fmt.Printf("Piccola %d finisce\n", id)
				countLitri -= 0.5 // tolgo 0.5 litri
				countDieci++ // aggiungo 1 moneta da 10
				isOccupato = false
				ackPiccola[id] <- true
			case id := <- when(countLitri >= 1.5 && countVenti < MAX_VENTI && len(piccolaRiempie) == 0 && !isOccupato, grandeRiempie):
				fmt.Printf("Grande %d riempie\n", id)
				isOccupato = true
				ackGrande[id] <- true
			case id := <- grandeFinisce:
				fmt.Printf("Grande %d finisce\n", id)
				countLitri -= 1.5 // tolgo 1.5 litri
				countVenti++ // aggiungo 1 moneta da 20
				isOccupato = false
				ackGrande[id] <- true
			case <- terminaServer:
				//termino tutti i figli che hanno un ciclo infito
				fmt.Println("\nTerminazione figli")

				terminaAddetto <- true
				<- done
			
				fmt.Println("Terminazione server")
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
	for i := 0; i < N_PICCOLA; i++{
		ackPiccola[i] = make(chan bool)
	}

	for i := 0; i < N_GRANDE; i++{
		ackGrande[i] = make(chan bool)
	}

	// Avvio routines

	go server()
	go addetto()

	for i := 0; i < N_GRANDE; i ++ {
		go grande(i)
	}

	sleepRandomMillisMin(1000, 2000)

	for i := 0; i < N_PICCOLA; i ++ {
		go piccola(i)
	}

	// Attesa terminazione

	for i := 0; i < N_PICCOLA + N_GRANDE; i++ {
		<- done
	}

	// Terminazione server
	terminaServer <- true
	<- doneServer

	fmt.Println("Fine esecuzione")
}