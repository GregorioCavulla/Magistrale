package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 60
const MAXS = 3
const NC = 30

const NSCOLARESCHE = 3
const NVISITATORI = 10
const NADDETTI = 5

var entra_in_scolaresca = make(chan int,10)
var esci_in_scolaresca = make(chan int,10)

var entra_out_scolaresca = make(chan int,10)
var esci_out_scolaresca = make(chan int,10)

var ack_s[NSCOLARESCHE]chan int


var entra_in_visitatore = make(chan int,10)
var esci_in_visitatore = make(chan int,10)

var entra_out_visitatore = make(chan int,10)
var esci_out_visitatore = make(chan int,10)

var ack_v[NVISITATORI]chan int


var entra_in_addetto = make(chan int,10)
var esci_in_addetto = make(chan int,10)

var entra_out_addetto = make(chan int,10)
var esci_out_addetto = make(chan int,10)

var ack_a[NVISITATORI]chan int

var termina = make(chan int)
var done = make(chan int)

func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func addetto(id int){
	for{
		timeout := rand.Intn(70)
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		entra_in_addetto <- id
		
		select {
		case <- ack_a[id]:
		case <- termina:
			fmt.Printf("[Addetto][%d] terminazione\n", id)
			done <- 1
			return 
		}
		
		timeout = rand.Intn(20)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		esci_in_addetto <- id

		select {
		case <- ack_a[id]:
		case <- termina:
			fmt.Printf("[Addetto][%d] terminazione\n", id)
			done <- 1
			return 
		}

		timeout = rand.Intn(200)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		entra_out_addetto <- id 

		select {
		case <- ack_a[id]:
		case <- termina:
			fmt.Printf("[Addetto][%d] terminazione\n", id)
			done <- 1
			return 
		}

		timeout = rand.Intn(20)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		esci_out_addetto <- id

		select {
		case <- ack_a[id]:
		case <- termina:
			fmt.Printf("[Addetto][%d] terminazione\n", id)
			done <- 1
			return 
		}

	}
}


func visitatore(id int){
	timeout := rand.Intn(50)
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	entra_in_visitatore <- id
	<- ack_v[id]
	timeout = rand.Intn(10)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	esci_in_visitatore <- id
	<- ack_v[id]
	timeout = rand.Intn(30)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	entra_out_visitatore <- id
	<- ack_v[id]
	timeout = rand.Intn(10)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	esci_out_visitatore <- id
	<- ack_v[id]
	
	done <- 1	
}

func scolaresca(id int){
	timeout := rand.Intn(50)
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	entra_in_scolaresca <- id
	<- ack_s[id]
	timeout = rand.Intn(10)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	esci_in_scolaresca <- id
	<- ack_s[id]
	timeout = rand.Intn(30)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	entra_out_scolaresca <- id
	<- ack_s[id]
	timeout = rand.Intn(10)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	esci_out_scolaresca <- id
	<- ack_s[id]
	
	done <- 1	
}


func mostra(){

	var persone int = 0; 
	var nA int = 0;

	var inTransito int = 0
	var outTransito int = 0

	var scolarescaIN int = 0
	var scolarescaOUT int = 0

	for {
		select{
			case x:= <- esci_out_scolaresca :
				scolarescaOUT --
				outTransito -= 25
				fmt.Printf("[Scolaresca][%d] esce OUT\n",x)
				ack_s[x] <- 1
			case x:= <- esci_out_visitatore:
				outTransito --
				fmt.Printf("[Visitatore][%d] esce OUT\n",x)
				ack_v[x] <- 1
			case x:= <- esci_out_addetto:
				outTransito --
				fmt.Printf("[Addetto][%d] esce OUT\n",x)
				ack_a[x] <- 1
			case x:= <- esci_in_addetto:
				inTransito --
				fmt.Printf("[Addetto][%d] esce IN\n",x)
				ack_a[x] <- 1	
			case x:= <- esci_in_visitatore:
				inTransito --
				fmt.Printf("[Visitatore][%d] esce IN\n",x)
				ack_v[x] <- 1
			case x:= <- esci_in_scolaresca :
				scolarescaIN --
				inTransito -= 25
				fmt.Printf("[Scolaresca][%d] esce IN\n",x)
				ack_s[x] <- 1	
			case x:= <- when(outTransito+25 <= NC && inTransito==0 ,entra_out_scolaresca):
				scolarescaOUT ++;
				outTransito += 25;
				persone -= 25;
				fmt.Printf("[Scolaresca][%d] entra OUT\n",x)
				ack_s[x] <- 1 
			case x:= <- when(scolarescaIN==0 && inTransito+outTransito<NC && len(entra_out_scolaresca)==0,entra_out_visitatore):
				outTransito ++
				persone --
				fmt.Printf("[Visitatore][%d] entra OUT\n",x)
				ack_v[x] <- 1
			case x:= <- when(scolarescaIN==0 && inTransito+outTransito<NC && len(entra_out_scolaresca)==0 && len(entra_out_visitatore)==0 && persone-nA==0 ,entra_out_addetto):
				outTransito ++
				persone --
				nA --
				fmt.Printf("[Addetto][%d] esce OUT\n",x)
				ack_a[x] <- 1
			case x:= <- when(inTransito+outTransito<NC && scolarescaOUT==0 && persone<N && nA<MAXS && len(entra_out_addetto)==0 && len(entra_out_scolaresca)==0 && len(entra_out_visitatore)==0 ,entra_in_addetto):
				inTransito ++
				persone ++
				nA ++
				fmt.Printf("[Addetto][%d] entra IN\n",x)
				ack_a[x] <-	1	
			case x:= <- when( inTransito+outTransito<NC && scolarescaOUT==0 && persone<N && nA>0 && len(entra_in_addetto)==0 && len(entra_out_addetto)==0 && len(entra_out_scolaresca)==0 && len(entra_out_visitatore)==0,entra_in_visitatore):
				inTransito ++
				persone ++
				fmt.Printf("[Visitatore][%d] entra IN\n",x)
				ack_v[x] <- 1
			case x:= <- when( persone+25<N && inTransito+25<NC && outTransito==0 && nA>0 && len(entra_in_visitatore)==0 && len(entra_in_addetto)==0 && len(entra_out_addetto)==0 && len(entra_out_scolaresca)==0 && len(entra_out_visitatore)==0, entra_in_scolaresca):		
				inTransito += 25
				scolarescaIN ++
				persone +=25
				fmt.Printf("[Scolaresca][%d] entra IN\n",x)
				ack_s[x] <- 1
			case <-termina: 
				fmt.Println("FINE !!!!!!")
				done <- 1
				return
		}
	}

}


func main() {

	for i := 0; i < NADDETTI; i++ {
		ack_a[i] = make(chan int)
	}

	for i := 0; i < NSCOLARESCHE; i++ {
		ack_s[i] = make(chan int)
	}

	for i := 0; i < NVISITATORI; i++ {
		ack_v[i] = make(chan int)
	}

	go mostra()

	for i:=0; i<NADDETTI; i++ {
		go addetto(i)
	}

	for i:=0; i<NSCOLARESCHE; i++ {
		go scolaresca(i)
	}

	for i:=0; i<NVISITATORI; i++ {
		go visitatore(i)
	}

	for i:=0; i<NSCOLARESCHE+NVISITATORI; i++{
		<- done
	}

	for i:=0; i<NADDETTI+1; i++{
		termina <- 1
	}

	for i:=0; i<NADDETTI+1; i++{
		<- done 
	}

}