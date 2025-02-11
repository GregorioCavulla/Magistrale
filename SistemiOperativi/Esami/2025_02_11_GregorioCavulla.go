package main

import ( //stampa, random, sleep
	"fmt"
	"math/rand"
	"time"
)

const MAXBUFF int = 100 // Max buffer per canali

// Funzione per dormire un tempo casuale di millisecondi
func sleepRandomMillis(max int) {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	time.Sleep(time.Duration(rand.Intn(max)) * time.Millisecond)
}

// Funzione per dormire un tempo casuale di millisecondi con tempo minimo
func sleepRandomMillisMin(min int, max int) {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Millisecond)
}

// Funzione when
func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

// Costanti
const (
	MAX_TRATTO = 50 //capienzamassima in peq

	N_AUTO       = 50 //15 peq
	N_BICICLETTE = 50 //1 peq
)

// Canali
var (
	entraAutoSalita  = make(chan int, MAXBUFF)
	entraAutoDiscesa = make(chan int, MAXBUFF)
	esceAutoSalita   = make(chan int, MAXBUFF)
	esceAutoDiscesa  = make(chan int, MAXBUFF)

	ackAuto [N_AUTO]chan bool

	entraBiciclettaSalita  = make(chan int, MAXBUFF)
	entraBiciclettaDiscesa = make(chan int, MAXBUFF)
	esceBiciclettaSalita   = make(chan int, MAXBUFF)
	esceBiciclettaDiscesa  = make(chan int, MAXBUFF)

	ackBicicletta [N_BICICLETTE]chan bool

	addettoBlocca = make(chan int, MAXBUFF)
	addettoChiude = make(chan int, MAXBUFF)
	addettoApre   = make(chan int, MAXBUFF)

	ackAddetto = make(chan bool)

	terminaAddetto = make(chan bool)
	terminaStrada  = make(chan bool)

	done       = make(chan bool)
	doneStrada = make(chan bool)
)

// Routines

func addetto(stato bool) { // se il ponte parte aperto è la fine, si rompe tutto
	// basta cambiare l'ordine dei canali:
	// - ponte chiuso: Apre, Blocca, Chiude
	// - ponte aperto: Blocca, Chiude, Apre
	// posso creare una variabile nel main che determina lo stato di partenza del ponte, passarla alla routine addetto e fare un if che scelga l'ordine giusto

	var res bool

	if stato { // ponte parte aperto (variabile isOpen nella routine main)
		for {

			sleepRandomMillis(600)

			addettoBlocca <- 1
			select {
			case <-terminaAddetto:
				fmt.Printf("[Addetto] terminazione\n")
				done <- true
				return
			case <-ackAddetto:
			}
			// codice che capisce quando la strada si è svuotata
			// probabile sia gestito strada

			for {
				addettoChiude <- 1

				select {
				case <-terminaAddetto:
					fmt.Printf("[Addetto] terminazione\n")
					done <- true
					return
				case res = <-ackAddetto:
				}

				if res {
					break
				}
			}

			sleepRandomMillis(600)

			addettoApre <- 1

			select {
			case <-terminaAddetto:
				fmt.Printf("[Addetto] terminazione\n")
				done <- true
				return
			case <-ackAddetto:
			}

			sleepRandomMillis(600)

		}
	} else { // ponte parte chiuso
		for {

			sleepRandomMillis(600)

			addettoApre <- 1

			select {
			case <-terminaAddetto:
				fmt.Printf("[Addetto] terminazione\n")
				done <- true
				return
			case <-ackAddetto:
			}

			sleepRandomMillis(600)

			addettoBlocca <- 1
			select {
			case <-terminaAddetto:
				fmt.Printf("[Addetto] terminazione\n")
				done <- true
				return
			case <-ackAddetto:
			}
			// codice che capisce quando la strada si è svuotata
			// probabile sia gestito strada

			for {
				addettoChiude <- 1

				select {
				case <-terminaAddetto:
					fmt.Printf("[Addetto] terminazione\n")
					done <- true
					return
				case res = <-ackAddetto:
				}

				if res {
					break
				}
			}

			sleepRandomMillis(600)

		}
	}

}

func auto(id int) { // questo funziona

	sleepRandomMillis(200)

	var percorso int = rand.Intn(2)

	switch percorso {
	case 0: //salita
		entraAutoSalita <- id
		<-ackAuto[id]

		sleepRandomMillisMin(100, 300)

		esceAutoSalita <- id
		<-ackAuto[id]
	case 1: //discesa
		entraAutoDiscesa <- id
		<-ackAuto[id]

		sleepRandomMillisMin(100, 300)

		esceAutoDiscesa <- id
		<-ackAuto[id]
	}

	//fine
	done <- true
}

func bicicletta(id int) { // uguale a prima
	sleepRandomMillis(200)

	var percorso int = rand.Intn(2)

	switch percorso {
	case 0: //salita
		entraBiciclettaSalita <- id
		<-ackBicicletta[id]

		sleepRandomMillisMin(100, 300)

		esceBiciclettaSalita <- id
		<-ackBicicletta[id]
	case 1: //discesa
		entraBiciclettaDiscesa <- id
		<-ackBicicletta[id]

		sleepRandomMillisMin(100, 300)

		esceBiciclettaDiscesa <- id
		<-ackBicicletta[id]
	}

	//fine
	done <- true
}

func strada(stato bool) { // routine server

	var countPeq int = 0 //peq presenti sul tratto di strada

	var isOpen bool = stato //ponte aperto o chiuso

	// booleans per evitare frontali
	var autoInSalita bool = false
	var autoInDiscesa bool = false

	var biciclettaInSalita bool = false
	var biciclettaInDiscesa bool = false

	for {

		select {
		//addetto blocca
		case <-when(isOpen, addettoBlocca):
			fmt.Printf("[Strada] addetto blocca\n")
			// codice
			isOpen = false
			ackAddetto <- true
		//addetto chiude
		case <-addettoChiude:
			// codice
			if countPeq == 0 {
				fmt.Printf("[Strada] addetto chiude\n")
				isOpen = false
				ackAddetto <- true
			} else {
				ackAddetto <- false
			}
		//addetto apre
		case <-when(!isOpen, addettoApre):
			fmt.Printf("[Strada] addetto apre\n")
			// codice
			isOpen = true
			ackAddetto <- true
		//auto entra in salita
		case id := <-when(isOpen && countPeq+15 <= MAX_TRATTO && !autoInDiscesa && !biciclettaInDiscesa, entraAutoSalita): //non sicuro della condizione sulla bicicletta in direzione opposta, ho chiesto alla prof
			fmt.Printf("[Strada] auto %d entra in salita\n", id)
			// codice
			countPeq += 15
			autoInSalita = true
			ackAuto[id] <- true
		//auto esce in salita
		case id := <-esceAutoSalita:
			fmt.Printf("[Strada] auto %d esce in salita \n", id)
			// codice
			countPeq -= 15
			autoInSalita = false
			ackAuto[id] <- true
		//auto entra in discesa
		case id := <-when(isOpen && countPeq+15 <= MAX_TRATTO && !autoInSalita && !biciclettaInSalita && len(entraAutoSalita)+len(entraBiciclettaSalita) == 0 && len(entraBiciclettaDiscesa) == 0, entraAutoDiscesa): //non sicuro della condizione sulla bicicletta in direzione opposta, ho chiesto alla prof
			fmt.Printf("[Strada]  auto %d entra in discesa\n", id)
			// codice
			countPeq += 15
			autoInDiscesa = true
			ackAuto[id] <- true
		//auto esce in discesa
		case id := <-esceAutoDiscesa:
			fmt.Printf("[Strada]  auto %d esce in discesa\n", id)
			// codice
			countPeq -= 15
			autoInDiscesa = false
			ackAuto[id] <- true
		//bicicletta entra in salita
		case id := <-when(isOpen && countPeq < MAX_TRATTO && !autoInDiscesa && len(entraAutoSalita) == 0, entraBiciclettaSalita):
			fmt.Printf("[Strada]  bicicletta %d entra in salita\n", id)
			// codice
			countPeq += 1
			biciclettaInSalita = true
			ackBicicletta[id] <- true
		//bicicletta esce in salita
		case id := <-esceBiciclettaSalita:
			fmt.Printf("[Strada] bicicletta %d esce in salita\n", id)
			// codice
			countPeq -= 1
			biciclettaInSalita = false
			ackBicicletta[id] <- true
		//bicicletta entra in discesa
		case id := <-when(isOpen && countPeq < MAX_TRATTO && !autoInSalita && len(entraAutoSalita)+len(entraBiciclettaSalita) == 0, entraBiciclettaDiscesa):
			fmt.Printf("[Strada] bicicletta %d entra in discesa\n", id)
			// codice
			countPeq += 1
			biciclettaInDiscesa = true
			ackBicicletta[id] <- true
		//bicicletta esce in discesa
		case id := <-esceBiciclettaDiscesa:
			fmt.Printf("[Strada] bicicletta %d esce in discesa\n", id)
			// codice
			countPeq -= 1
			biciclettaInDiscesa = false
			ackBicicletta[id] <- true
		//termina
		case <-terminaStrada:
			//termino tutti i figli che hanno un ciclo infito

			terminaAddetto <- true
			<-done

			fmt.Printf("[Strada] terminazione\n")
			doneStrada <- true
			return
		}
	}
}

// Main

func main() { // Standard con for dei done
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Inizio esecuzione\n\n")

	var isOpen bool = true

	// Inizializzazioni ack
	for i := 0; i < N_AUTO; i++ {
		ackAuto[i] = make(chan bool)
	}

	for i := 0; i < N_BICICLETTE; i++ {
		ackBicicletta[i] = make(chan bool)
	}

	// Avvio routines

	go strada(isOpen)
	go addetto(isOpen)

	for i := 0; i < N_AUTO; i++ {
		go auto(i)
	}

	for i := 0; i < N_BICICLETTE; i++ {
		go bicicletta(i)
	}

	// Attesa terminazione

	for i := 0; i < N_AUTO+N_BICICLETTE; i++ {
		<-done
	}

	// Terminazione strada
	fmt.Printf("\nTerminazione\n")
	terminaStrada <- true
	<-doneStrada

	fmt.Printf("\nFine esecuzione\n")
}
