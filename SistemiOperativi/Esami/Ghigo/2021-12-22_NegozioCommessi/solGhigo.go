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
	id_cliente  int
	id_commesso int
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

	MAX_NEGOZIO = 15
	N_MASCHERINE = 5
	N_COMMESSI = 2
	N_CLIENTI = 5
	N_CLIENTI_VIP = 5
)

type Commesso struct{
	id int
	isDentro bool
	countClienti int
}

// Canali

var(
	fornitoreDeposita = make(chan bool, MAXBUFF)
	ackFornitore = make(chan bool)

	entraCliente = make(chan int, MAXBUFF)
	esceCliente = make(chan Richiesta, MAXBUFF)
	ackCliente[N_CLIENTI]chan int

	entraClienteVip = make(chan int, MAXBUFF)
	esceClienteVip = make(chan Richiesta, MAXBUFF)
	ackClienteVip[N_CLIENTI_VIP]chan int

	entraCommesso = make(chan int, MAXBUFF)
	esceCommesso = make(chan int, MAXBUFF)
	ackCommesso[N_COMMESSI]chan bool

	terminaCommesso = make(chan bool)
	terminaFornitore = make(chan bool)
	terminaNegozio = make(chan bool)

	done = make(chan int)
	doneNegozio = make(chan int)
)

// Routines

func fornitore() {
	for{

		sleepRandomMillis(150)

		fornitoreDeposita <- true

		select{
			case <-ackFornitore:
			case <-terminaFornitore:
				fmt.Println("[Fornitore] terminazione")
				done <- 1
				return
		}

		sleepRandomMillis(150)
	}
}

func cliente(id int) {

	sleepRandomMillis(50)

	entraCliente <- id
	idCommesso := <-ackCliente[id]
	
	sleepRandomMillis(250)

	esceCliente <- Richiesta {id, idCommesso}
	<- ackCliente[id]

	done <- 1
}

func clienteVip(id int) {

	sleepRandomMillis(50)

	entraClienteVip <- id
	idCommesso := <-ackClienteVip[id]
	
	sleepRandomMillis(250)

	esceClienteVip <- Richiesta {id, idCommesso}
	<- ackClienteVip[id]

	done <- 1
}

func commesso(id int) {

	var res bool

	for{

		sleepRandomMillis(20)

		entraCommesso <- id
		select{
			case <- ackCommesso[id]:
			case <- terminaCommesso:
				fmt.Printf("[Commesso][%d] terminazione\n",id)
				done <- 1
				return
		}

		sleepRandomMillis(300)

		for{

			esceCommesso <- id
			select{
				case res = <- ackCommesso[id]:
				case <- terminaCommesso:
					fmt.Printf("[Commesso][%d] terminazione\n",id)
					done <- 1
					return
			}

			if(res){
				break
			}
		}

	}

}

func negozio() {

	// inizializzazioni

	var countMascherine = 0	// numero di mascherine nel negozio
	var countPersone = 0	// numero di persone nel negozio, max: MAX_NEGOZIO
	var countCommessi = 0 	// numero di commessi dentro al negozio, max: N_COMMESSI
	var commessi [N_COMMESSI]Commesso	//array di commessi

	// inizializzo i commessi
	for i := 0; i < N_COMMESSI; i++{
		commessi[i] = Commesso{i, false, 0}
	}

	for{

		select{
			//fornitore
			case <- fornitoreDeposita:
				fmt.Printf("[Negozio] Fornitore deposita mascherine\n")
				countMascherine += N_MASCHERINE
				ackFornitore <- true
			// entra
				//commesso
			case id := <- when(countPersone < MAX_NEGOZIO, entraCommesso):
				fmt.Printf("[Negozio] Entra commesso %d\n", id)
				countPersone += 1
				countCommessi += 1
				commessi[id].isDentro = true
				commessi[id].countClienti = 0
				ackCommesso[id] <- true
				//cliente vip
			case id := <- when(countPersone < MAX_NEGOZIO && (countPersone-countCommessi)/3 < countCommessi && countMascherine > 0 && len(entraCommesso) == 0, entraClienteVip):
				fmt.Printf("[Negozio] Entra cliente vip %d\n", id)
				countPersone += 1
				countMascherine -= 1

				var idCommesso int
				// associa commesso
				for i := 0; i < N_COMMESSI; i++{
					if(commessi[i].isDentro && commessi[i].countClienti < 3){
						idCommesso = i
						break
					}
				}
				commessi[idCommesso].countClienti += 1
				ackClienteVip[id] <- idCommesso
				//cliente
			case id := <- when(countPersone < MAX_NEGOZIO && (countPersone-countCommessi)/3 < countCommessi && countMascherine > 0 && len(entraCommesso) == 0 && len(entraClienteVip) == 0, entraCliente):
				fmt.Printf("[Negozio] Entra cliente %d\n", id)
				countPersone += 1
				countMascherine -= 1
	
				var idCommesso int
				// associa commesso
				for i := 0; i < N_COMMESSI; i++{
					if(commessi[i].isDentro && commessi[i].countClienti < 3){
						idCommesso = i
						break
					}
				}
				commessi[idCommesso].countClienti += 1
				ackCliente[id] <- idCommesso
			// esce
				//commesso
			case id := <- esceCommesso:
				fmt.Printf("[Negozio] Commesso %d vorrebbe uscire...\n", id)
				if(commessi[id].countClienti == 0){
					fmt.Printf("[Negozio] Commesso %d esce\n", id)
					ackCommesso[id] <- true
					countPersone -= 1
					countCommessi -= 1
					commessi[id].isDentro = false
				}else{
					fmt.Printf("[Negozio] Commesso %d non può uscire\n", id)
					ackCommesso[id] <- false
				}
				//cliente vip
			case req := <- esceClienteVip:
				fmt.Printf("[Negozio] Cliente vip %d esce\n", req.id_cliente)
				commessi[req.id_commesso].countClienti -= 1
				ackClienteVip[req.id_cliente] <- 1
				countPersone -= 1
				//cliente
			case req := <- esceCliente:
				fmt.Printf("[Negozio] Cliente %d esce\n", req.id_cliente)
				commessi[req.id_commesso].countClienti -= 1
				ackCliente[req.id_cliente] <- 1
				countPersone -= 1
			// termina
			case <- terminaNegozio:

				fmt.Printf("\nKillo Commessi\n")
				for i := 0; i < N_COMMESSI; i++ {
					terminaCommesso <- true
				}
			
				for i := 0; i < N_COMMESSI; i++ {
					<- done
				}
			
				fmt.Println("\nKillo Fornitore\n")
				terminaFornitore <- true
				<- done
			
				fmt.Println("\nKillo Negozio\n")

				doneNegozio <- 1
				return
		}
	}
}

// Main

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Inizio esecuzione")

	// Inizializzazioni
	for i := 0; i < N_CLIENTI; i++{
		ackCliente[i] = make(chan int)
	}

	for i := 0; i < N_CLIENTI_VIP; i++{
		ackClienteVip[i] = make(chan int)
	}

	for i := 0; i < N_COMMESSI; i++{
		ackCommesso[i] = make(chan bool)
	}

	// Avvio routines

	go negozio()
	go fornitore()

	for i := 0; i < N_COMMESSI; i ++ {
		go commesso(i)
	}

	for i := 0; i < N_CLIENTI; i ++ {
		go cliente(i)
	}

	for i := 0; i < N_CLIENTI_VIP; i ++ {
		go clienteVip(i)
	}

	// Attesa terminazione

	for i := 0; i < N_CLIENTI + N_CLIENTI_VIP; i++ {
		<- done
	}

	// killo commessi, fornitore, negozio

	// termina negozio
	terminaNegozio <- true
	<- doneNegozio
	
	fmt.Println("Fine esecuzione")
}