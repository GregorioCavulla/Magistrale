package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX = 15 		
const NM = 5

const NCLIENTI = 5
const NCLIENTIVIP = 5
const NCOMMESSI = 2

var consegna_fornitore = make(chan int,1)
var	ack_fornitore = make(chan int)
var entra_cliente = make(chan int,10)
var esci_cliente = make(chan uscita,10)

var ack_cliente[NCLIENTI]chan int

var entra_cliente_vip = make(chan int,10)
var esci_cliente_vip = make(chan uscita,10)

var ack_clientevip[NCLIENTIVIP]chan int

var entra_commesso = make(chan int,10)
var esci_commesso = make(chan int,10)

var ack_commesso[NCOMMESSI]chan int

var termina = make(chan int)
var done = make(chan int)

type uscita struct {
	id int
	id_commesso  int
}

func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func fornitore(){
	for{

		timeout := rand.Intn(10)
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		consegna_fornitore <-1
		
		select {
		case <- ack_fornitore:
		case <- termina:
			fmt.Printf("[Fornitore] terminazione\n")
			done <- 1
			return 
		}
		
		timeout = rand.Intn(100)
		time.Sleep(time.Duration(timeout) * time.Millisecond)

	}
}

func cliente(id int){
	timeout := rand.Intn(50)
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	entra_cliente <- id
	id_commesso := <- ack_cliente[id] 

	timeout = rand.Intn(50)
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	esci_cliente <- uscita {id, id_commesso}
	<- ack_cliente[id]
	
	done <- 1
} 

func cliente_vip(id int){
	timeout := rand.Intn(50)
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	entra_cliente_vip <- id
	id_commesso := <- ack_clientevip[id] 

	timeout = rand.Intn(50)
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	esci_cliente_vip <- uscita {id, id_commesso}
	<- ack_clientevip[id]
	
	done <- 1
} 

func commesso(id int){
	
	var res int = 0

	for {
		timeout := rand.Intn(50)
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		entra_commesso <- id
		select {
		case <- ack_commesso[id]:
		case <- termina:
			fmt.Printf("[Commesso][%d] terminazione\n",id)
			done <- 1
			return 
		}

		timeout = rand.Intn(100)
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		for {
			esci_commesso <- id

			select {
			case res= <- ack_commesso[id]:
			case <- termina:
				fmt.Printf("[Commesso][%d] terminazione\n",id)
				done <- 1
				return
			}	
			
			if (res == 1) { break }
			
		}

	}
	
}

func negozio(){

	var N int = 0
	var NC int = 0
	var MASCHERINE int = 0

	var posti[NCOMMESSI] int

	for i:=0; i<NCOMMESSI; i++{
		posti[i] = -1
	}

	for{
		select{
			case <- consegna_fornitore :
				MASCHERINE += NM
				fmt.Printf("[Fornitore] consegna\n")
				ack_fornitore <- 1
			case x:= <- esci_commesso:
				if (posti[x] == 3) {
					ack_commesso[x] <- 1
					N --
					NC --
					posti[x] = -1
					fmt.Printf("[Commesso][%d] uscito\n", x)
				}else{
					ack_commesso[x] <- 0
				}
			case x:= <- esci_cliente_vip:
				N--
				posti[x.id_commesso] ++
				fmt.Printf("[Cliente vip][%d] uscito\n", x.id)
				ack_clientevip[x.id] <- 1
			case x:= <- esci_cliente:
				N--
				posti[x.id_commesso] ++
				fmt.Printf("[Cliente][%d] uscito\n", x.id)
				ack_cliente[x.id] <- 1
			case x:= <- when(N < MAX ,entra_commesso):
				N ++
				NC ++
				posti[x] = 3
				fmt.Printf("[Commesso][%d] entrato\n", x)
				ack_commesso[x] <- 1
			case x:= <- when(N < MAX && MASCHERINE>0 && (N-NC)/3 < NC && len(entra_commesso)==0, entra_cliente_vip):
				N ++
				MASCHERINE --
				var id_commesso int
				for i:=0; i<NCOMMESSI; i++{
					if(posti[i]>0){
						id_commesso = i
						break
					}
				}
				posti[id_commesso] -- 
				fmt.Printf("[Cliente vip][%d] entrato\n", x)
				ack_clientevip[x] <-id_commesso
			case x:= <- when(N < MAX && MASCHERINE>0 && (N-NC)/3 < NC && len(entra_cliente_vip)==0 && len(entra_commesso)==0, entra_cliente):
				N ++
				MASCHERINE --
				var id_commesso int
				for i:=0; i<NCOMMESSI; i++{
					if(posti[i]>0){
						id_commesso = i
						break
					}
				} 
				posti[id_commesso] --
				fmt.Printf("[Cliente][%d] entrato\n", x)
				ack_cliente[x] <-id_commesso
			case <- termina: 
				fmt.Println("FINE !!!!!!")
				done <- 1
				return
		}
	}

}


func main() {

	for i := 0; i < NCOMMESSI; i++ {
		ack_commesso[i] = make(chan int)
	}

	for i := 0; i < NCLIENTIVIP; i++ {
		ack_clientevip[i] = make(chan int)
	}

	for i := 0; i < NCLIENTI; i++ {
		ack_cliente[i] = make(chan int)
	}

	go negozio()
	go fornitore()

	for i:=0; i<NCLIENTI; i++ {
		go cliente(i)
	}

	for i:=0; i<NCLIENTIVIP; i++ {
		go cliente_vip(i)
	}

	for i:=0; i<NCOMMESSI; i++ {
		go commesso(i)
	}

	for i:=0; i<NCLIENTI+NCLIENTIVIP; i++{
		<- done
	}

	for i:=0; i<NCOMMESSI+2; i++{
		termina <- 1
	}

	for i:=0; i<NCOMMESSI+2; i++{
		<- done
	}

}