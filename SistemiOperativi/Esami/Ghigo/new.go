package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ---------------------- COSTANTI ----------------------
// numeri massimi client e risorsa dal testo
const MAX_CLIENTS = 10     // Numero massimo di client concorrenti
const MAX_RESOURCES = 5    // Numero massimo di risorse disponibili

// ---------------------- STRUTTURE DATI ----------------------
// strutture per richiesta
type Request struct {
	id       int       // Identificativo del client
	priority int       // Priorità della richiesta (valore minore = maggiore priorità)
	response chan bool // Canale per ricevere la risposta dal server
}


// ---------------------- CANALI ----------------------
var richiesta = make(chan Request)  // Canale per inviare richieste di risorse
var rilascio = make(chan int)       // Canale per rilasciare risorse
var termina = make(chan bool)       // Canale per terminare il server


// ---------------------- STRUTTURA DEL SERVER ----------------------
type Server struct {
	disponibili int       // Numero di risorse attualmente disponibili
	queue       []Request // Coda delle richieste in attesa
}

// Funzione principale del server: gestisce allocazione e rilascio delle risorse
func (s *Server) run() {
	for {
		select {
		case req := <-richiesta: // Ricezione di una richiesta da un client
			mu.Lock()
			if s.disponibili > 0 { // Se ci sono risorse disponibili
				s.disponibili--           // Decremento le risorse disponibili
				req.response <- true      // Rispondo al client con "OK"
				fmt.Printf("[Server] Risorsa assegnata a client %d\n", req.id)
			} else { // Se non ci sono risorse disponibili
				s.queue = append(s.queue, req) // Aggiungo la richiesta in coda
				fmt.Printf("[Server] Client %d in attesa\n", req.id)
			}
			mu.Unlock()

		case res := <-rilascio: // Un client rilascia una risorsa
			mu.Lock()
			s.disponibili++ // Incremento le risorse disponibili
			fmt.Printf("[Server] Risorsa %d rilasciata\n", res)

			// Se ci sono richieste in coda, serviamo il client con priorità più alta
			if len(s.queue) > 0 {
				highestPriorityIndex := 0 // Trova il client con priorità più alta
				for i := 1; i < len(s.queue); i++ {
					if s.queue[i].priority < s.queue[highestPriorityIndex].priority {
						highestPriorityIndex = i
					}
				}
				// Seleziona il client prioritario e lo serve
				chosen := s.queue[highestPriorityIndex]
				s.queue = append(s.queue[:highestPriorityIndex], s.queue[highestPriorityIndex+1:]...)
				chosen.response <- true // Rispondi al client con "OK"
				s.disponibili--         // Decrementa le risorse disponibili
				fmt.Printf("[Server] Client %d con priorità %d servito\n", chosen.id, chosen.priority)
			}
			mu.Unlock()

		case <-termina: // Ricevuto segnale di terminazione
			fmt.Println("[Server] Terminazione")
			return
		}
	}
}

// ---------------------- FUNZIONE CLIENT ----------------------
func client(id int, priority int) {
	response := make(chan bool) // Canale per ricevere la risposta del server

	// Invio richiesta al server con il mio ID e la mia priorità
	richiesta <- Request{id: id, priority: priority, response: response}

	// Attendo la risposta del server
	approved := <-response
	if approved {
		fmt.Printf("[Client %d] Risorsa ottenuta\n", id)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Simula utilizzo della risorsa
		rilascio <- id                                       // Rilascia la risorsa
		fmt.Printf("[Client %d] Risorsa rilasciata\n", id)
	}
}

// ---------------------- FUNZIONE MAIN ----------------------
func main() {
	rand.Seed(time.Now().Unix()) // Inizializza il generatore di numeri casuali

	// Creazione del server con un numero iniziale di risorse disponibili
	server := Server{disponibili: MAX_RESOURCES}
	go server.run() // Avvio il server come goroutine

	// Creazione di più client concorrenti con priorità casuali
	for i := 0; i < MAX_CLIENTS; i++ {
		go client(i, rand.Intn(3)) // Ogni client ha una priorità casuale tra 0 e 2
	}

	// Attesa per consentire l'esecuzione dei client
	time.Sleep(10 * time.Second)

	// Segnale di terminazione al server
	termina <- true
}