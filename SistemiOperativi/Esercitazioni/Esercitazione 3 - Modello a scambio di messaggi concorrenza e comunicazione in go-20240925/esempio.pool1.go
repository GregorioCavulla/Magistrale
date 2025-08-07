// pool di risorse equivalenti senza guardie logiche
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXPROC = 100 //massimo numero di processi
const MAXRES = 5    //massimo numero di risorse nel pool

var richiesta = make(chan int)
var rilascio = make(chan int)
var risorsa [MAXPROC]chan int
var done = make(chan int)
var termina = make(chan int)

func client(i int) {	// funzione client
	richiesta <- i	// invio richiesta
	r := <-risorsa[i]	// attesa risorsa
	fmt.Printf("\n [client %d] uso della risorsa %d\n", i, r)	// output
	timeout := rand.Intn(3)	// generazione random del timeout
	time.Sleep(time.Duration(timeout) * time.Second)	// attesa
	rilascio <- r	// rilascio risorsa
	done <- i //comunico al main la terminazione
}

func server(nris int, nproc int) {	// funzione server

	var disponibili int = nris	// numero di risorse disponibili
	var res, p, i int	// indice risorsa, processo, indice
	var libera [MAXRES]bool	// vettore di risorse libere
	var sospesi [MAXPROC]bool	// vettore di processi sospesi
	var nsosp int = 0	// numero di processi sospesi

	for i := 0; i < nris; i++ {	// all'avvio del server libero tutte le risorse
		libera[i] = true
	}
	for i := 0; i < nproc; i++ { // all'avvio del server libero tutti i processi
		sospesi[i] = false
	}

	for { // ciclo infinito del server
		time.Sleep(time.Second * 1) // sleep 1s
		fmt.Println("nuovo ciclo server")
		select { // selezione tra i canali
		case res = <-rilascio: // rilascio della risorsa
			if nsosp == 0 { // se non ci sono processi sospesi
				disponibili++ // aggiungo 1 processo disponibile
				libera[res] = true // la risorsa è libera
				fmt.Printf("[server]  restituita risorsa: %d  \n", res)
			} else { // se ci sono processi sospesi
				for i = 0; i < nproc && !sospesi[i]; i++ {  // usa il for per trovare il primo processo sospeso (trova la i corrispondente al primo processo sospeso)
				}
				sospesi[i] = false 
				nsosp--
				risorsa[i] <- res
			}
		case p = <-richiesta:
			if disponibili > 0 { //allocazione della risorsa
				for i = 0; i < nris && !libera[i]; i++ { // usa il for per trovare la prima risorsa libera (trova la i corrispondente alla prima risorsa libera)
				}
				libera[i] = false
				disponibili--
				risorsa[p] <- i
				fmt.Printf("[server]  allocata risorsa %d a cliente %d \n", i, p)
			} else { // attesa
				nsosp++
				sospesi[p] = true
				fmt.Printf("[server]  il cliente %d attende..\n", i)
			}
		case <-termina: // quando tutti i processi clienti hanno finito
			fmt.Println("FINE !!!!!!")
			done <- 1
			return

		}
	}
}

func main() {
	var cli, res int	// dichiarazioni variabili,

	rand.Seed(time.Now().Unix())	// seeding del random
	fmt.Printf("\n quanti clienti (max %d)? ", MAXPROC)	// richiesta input
	fmt.Scanf("%d", &cli)	// lettura input
	fmt.Println("clienti:", cli)	// output
	fmt.Printf("\n quante risorse (max %d)? ", MAXRES)	// richiesta input
	fmt.Scanf("%d", &res)	// lettura input
	fmt.Println("risorse da gestire:", res)	// output

	// inizializzazione canali
	for i := 0; i < cli; i++ {	// per ogni client
		risorsa[i] = make(chan int)	// crea un canale di comunicazione
	}

	for i := 0; i < cli; i++ {	// per ogni client
		go client(i)	// avvia il client
	}
	go server(res, cli)	// avvia il server

	//attesa della terminazione dei clienti:
	for i := 0; i < cli; i++ { // per ogni client
		<-done	// attesa terminazione
	}
	termina <- 1 //terminazione server
	<-done // attesa terminazione server
}
