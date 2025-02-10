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
	id_personalTrainer int
	nVolte int
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
	MAX_PALESTRA = 18
	MAX_PESI = 10

	N_UTENTI = 30

	N_PERSONAL_TRAINER = 3

)


// Canali
var(

	entraPersonalTrainer = make(chan int, MAXBUFF)
	escePersonalTrainer = make(chan int, MAXBUFF)
	ackPersonalTrainer [N_PERSONAL_TRAINER]chan bool

	entraUtentePesi = make(chan Richiesta, MAXBUFF)
	esceUtentePesi = make(chan int, MAXBUFF)
	entraUtenteCorsi = make(chan Richiesta, MAXBUFF)
	esceUtenteCorsi = make(chan Richiesta, MAXBUFF)
	ackUtente [N_UTENTI]chan Richiesta

	terminaPersonalTrainer = make(chan bool)
	terminaPalestra = make(chan bool)

	done = make(chan bool)
	donePalestra = make(chan bool)
)

// Routines

func personalTrainer(id int) {

	var res bool

	for{

		sleepRandomMillis(150)

		entraPersonalTrainer <- id

		select{
			case <- terminaPersonalTrainer:
				fmt.Println("[Personal Trainer", id, "] Terminazione")
				done <- true
				return
			case <- ackPersonalTrainer[id]:
				sleepRandomMillis(1000)
		}


		for{
			escePersonalTrainer <- id

			select{
			case <- terminaPersonalTrainer:
				fmt.Println("[Personal Trainer", id, "] Terminazione")
				done <- true
				return
			case res = <- ackPersonalTrainer[id]:
				sleepRandomMillis(1)
			}

			if res {
				break
			}
		}
	}
}

func utente(id int) {

	sleepRandomMillis(50)

	var percorso int = rand.Intn(3)
	var nVolte int = rand.Intn(5) + 1

	for i := 0; i < nVolte; i++ {
		req := Richiesta{id, -1, i}
		switch percorso {
		case 0: // solo pes
			entraUtentePesi <- req
			<- ackUtente[id]
			sleepRandomMillis(100)
			esceUtentePesi <- id
			<- ackUtente[id]
		case 1: // solo corsi
			entraUtenteCorsi <- req
			req := <- ackUtente[id]
			sleepRandomMillis(100)
			esceUtenteCorsi <- req
			<- ackUtente[id]
		case 2: // entrambi
			entraUtentePesi <- req
			<- ackUtente[id]
			sleepRandomMillis(100)
			esceUtentePesi <- id
			<- ackUtente[id]
			entraUtenteCorsi <- req
			req := <- ackUtente[id]
			sleepRandomMillis(100)
			esceUtenteCorsi <- req
			<- ackUtente[id]
		}
	}
	
	//fine
	done <- true
}

func personalTrainerLiberi(countPersonalTrainer int, countUtentiCorsi int) bool {
	// se il numero di personal trainer è maggiore del numero di utenti corsi
	// allora ci sono personal trainer liberi
	return countPersonalTrainer > countUtentiCorsi
}

type PersonalTrainer struct {
	id int
	isDentro bool
	id_utente int
}

func palestra(){

	var countPersonalTrainer = 0
	var countUtenti = 0
	var countUtentiPesi = 0
	var countUtentiCorsi = 0

	var personalTrainers [N_PERSONAL_TRAINER]PersonalTrainer

	for i := 0; i < N_PERSONAL_TRAINER; i++ {
		personalTrainers[i] = PersonalTrainer{i, false, -1}
	}

	var id_personalTrainer = -1

	for{

		select{
			case req := <- whenRichiesta(countUtenti < MAX_PALESTRA && countUtentiPesi < MAX_PESI && len(entraUtenteCorsi) == 0, entraUtentePesi):
				fmt.Printf("\n[Palestra] Utente %d entra in zona pesi per la %d volta", req.id, req.nVolte)
				countUtenti++
				countUtentiPesi++
				ackUtente[req.id] <- req
			case req := <- whenRichiesta(countUtenti < MAX_PALESTRA && personalTrainerLiberi(countPersonalTrainer, countUtentiCorsi) && len(entraPersonalTrainer) == 0, entraUtenteCorsi):
				fmt.Printf("\n[Palestra] Utente %d entra in zona corsi per la %d volta", req.id, req.nVolte)
				countUtenti++
				countUtentiCorsi++
				for i := 0; i < N_PERSONAL_TRAINER; i++ {
					if personalTrainers[i].id_utente == -1 && personalTrainers[i].isDentro {
						personalTrainers[i].id_utente = req.id
						id_personalTrainer = i
						break
					}
				}
				ackUtente[req.id] <- Richiesta{req.id, id_personalTrainer, req.nVolte}
			case id := <- esceUtentePesi:
				fmt.Printf("\n[Palestra] Utente %d esce dalla zona pesi", id)
				countUtenti--
				countUtentiPesi--
				ackUtente[id] <- Richiesta{id, -1, 0}
			case req := <- esceUtenteCorsi:
				fmt.Printf("\n[Palestra] Utente %d esce dalla zona corsi", req.id)
				countUtenti--
				countUtentiCorsi--
				id_personalTrainer = req.id_personalTrainer
				personalTrainers[id_personalTrainer].id_utente = -1
				ackUtente[req.id] <- Richiesta{req.id, -1, 0}
			case id := <- entraPersonalTrainer:
				fmt.Printf("\n[Palestra] Personal Trainer %d entra", id)
				countPersonalTrainer++
				personalTrainers[id].isDentro = true
				ackPersonalTrainer[id] <- true
			case id := <- escePersonalTrainer:
				fmt.Printf("\n[Palestra] Personal Trainer %d vuole uscire", id)
				if(personalTrainers[id].id_utente == -1){
					fmt.Printf("\n[Palestra] Personal Trainer %d esce", id)
					countPersonalTrainer--
					personalTrainers[id].isDentro = false
					ackPersonalTrainer[id] <- true
				}else{
					fmt.Printf("\n[Palestra] Personal Trainer %d non può uscire", id)
					ackPersonalTrainer[id] <- false
				}
			case <- terminaPalestra:
				//termino tutti i figli che hanno un ciclo infito
				fmt.Printf("\n\nTerminazione figli")

				for i := 0; i < N_PERSONAL_TRAINER; i++ {
					terminaPersonalTrainer <- true
					<- done
				}

				fmt.Printf("\nTerminazione palestra")
				donePalestra <- true
				return
		}
	}	
}

// Main

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("\nInizio esecuzione")

	// Inizializzazioni ack
	for i := 0; i < N_PERSONAL_TRAINER; i++ {
		ackPersonalTrainer[i] = make(chan bool)
	}

	for i := 0; i < N_UTENTI; i++ {
		ackUtente[i] = make(chan Richiesta)
	}

	// Avvio routines
	go palestra()

	for i := 0; i < N_PERSONAL_TRAINER; i++ {
		go personalTrainer(i)
	}

	for i := 0; i < N_UTENTI; i++ {
		go utente(i)
	}

	// Attesa terminazione

	for i := 0; i < N_UTENTI; i++ {
		<- done
	}

	// Terminazione Palestra
	terminaPalestra <- true
	<- donePalestra

	fmt.Printf("\nFine esecuzione")
}