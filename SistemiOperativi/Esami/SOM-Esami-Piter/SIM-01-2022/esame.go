package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NC = 20
const NF = 15

const LF = 5
const LC = 8
const LM = 3

const NAR = 3

var inizio_fornitore_c = make(chan int,1)
var fine_fornitore_c = make(chan int,1)
var	ack_fornitore_c = make(chan int)

var inizio_fornitore_f = make(chan int,1)
var fine_fornitore_f = make(chan int,1)
var	ack_fornitore_f = make(chan int)

var inizio_ar_c = make(chan int,10)
var fine_ar_c = make(chan int,10)

var inizio_ar_f = make(chan int,10)
var fine_ar_f = make(chan int,10)

var inizio_ar_m = make(chan int,10)
var fine_ar_m = make(chan int,10)

var ack_ar[NAR]chan int

var termina = make(chan int)
var done = make(chan int)

func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func fornitore_c(){
	for{

		timeout := rand.Intn(10)
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		inizio_fornitore_c <-1
		
		select {
		case <- ack_fornitore_c:
		case <- termina:
			fmt.Printf("[Fornitore C] terminazione\n")
			done <- 1
			return 
		}
		
		timeout = rand.Intn(100)
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		fine_fornitore_c <-1
		
		select {
		case <- ack_fornitore_c:
		case <- termina:
			fmt.Printf("[Fornitore C] terminazione\n")
			done <- 1
			return 
		}

	}
}

func fornitore_f(){
	for{

		timeout := rand.Intn(10)
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		inizio_fornitore_f <-1
		
		select {
		case <- ack_fornitore_f:
		case <- termina:
			fmt.Printf("[Fornitore F] terminazione\n")
			done <- 1
			return 
		}
		
		timeout = rand.Intn(100)
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		fine_fornitore_f <-1
		
		select {
		case <- ack_fornitore_f:
		case <- termina:
			fmt.Printf("[Fornitore F] terminazione\n")
			done <- 1
			return 
		}

	}
}

func addettoReparto(id int){

	for i:=0 ; i<5 ; i++{
		timeout := rand.Intn(50)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
	
		k:= rand.Intn(3)

		if (k==0){ //prelievo chirurgiche

			inizio_ar_c <- id
			<- ack_ar[id]
			
			timeout := rand.Intn(20)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
			
			fine_ar_c <- id
			<- ack_ar[id]

		}else if (k==1){ //prelievo ffp2
			
			inizio_ar_f <- id
			<- ack_ar[id]
			
			timeout := rand.Intn(20)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
			
			fine_ar_f <- id
			<- ack_ar[id]

		}else{ //prelievo misto

			inizio_ar_m <- id
			<- ack_ar[id]
			
			timeout := rand.Intn(20)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
			
			fine_ar_m <- id
			<- ack_ar[id]
		}	
	}

	done <- 1

} 


func magazzino(){

	var chirurgiche int = 0
	var ffp2 int = 0

	var arc int = 0	
	var arf int = 0

	var consegna_c bool = false
	var consegna_f bool = false

	for{
		select{
			case <- fine_fornitore_f:
				ffp2 = NF
				consegna_f = false 
				fmt.Println("[Fornitore F] fine consegna")
				ack_fornitore_f <-1
			case <- fine_fornitore_c:
				chirurgiche = NC
				consegna_c = false 
				fmt.Println("[Fornitore C] fine consegna")
				ack_fornitore_c <-1
			case x:= <- fine_ar_m:
				arc --
				arf --
				fmt.Printf("[Addetto Rerparto][%d] fine prelievo misto\n", x)
				ack_ar[x] <- 1
			case x:= <- fine_ar_f:
				arf --
				fmt.Printf("[Addetto Rerparto][%d] fine prelievo ffp2\n", x)
				ack_ar[x] <- 1
			case x:= <- fine_ar_c:
				arc --
				fmt.Printf("[Addetto Rerparto][%d] fine prelievo chirurgiche\n", x)
				ack_ar[x] <- 1
			case <- when(arf==0 && (ffp2 <= chirurgiche),inizio_fornitore_f):
				consegna_f = true;
				fmt.Println("[Fornitore F] inizio consegna")
				ack_fornitore_f <- 1
			case <- when(arc==0 && (ffp2 > chirurgiche || len(inizio_fornitore_f)== 0),inizio_fornitore_c):
				consegna_c = true;
				fmt.Println("[Fornitore C] inizio consegna")
				ack_fornitore_c <- 1	
			case x:= <- when(!consegna_c && !consegna_f && chirurgiche >= LM && ffp2 >= LM,inizio_ar_m):	
				arc ++
				arf ++
				chirurgiche -= LM
				ffp2 -= LM
				fmt.Printf("[Addetto Rerparto][%d] inizio prelievo misto\n",x)
				ack_ar[x] <- 1
			case x:= <- when(len(inizio_ar_m)==0 && !consegna_f && ffp2 >= LF,inizio_ar_f):
				arf ++
				ffp2 -= LF
				fmt.Printf("[Addetto Rerparto][%d] inizio prelievo ffp2\n",x)
				ack_ar[x] <- 1
			case x:= <- when(len(inizio_ar_m)==0 && len(inizio_ar_f)==0 && !consegna_c && chirurgiche >= LC,inizio_ar_c):
				arc ++
				chirurgiche -= LC
				fmt.Printf("[Addetto Rerparto][%d] inizio prelievo chirurgiche\n",x)
				ack_ar[x] <- 1
			case <- termina: 
				fmt.Println("[Magazzino] terminazione")
				done <- 1
				return
		}
	}

}


func main() {

	for i := 0; i < NAR; i++ {
		ack_ar[i] = make(chan int)
	}

	go magazzino()

	for i:=0; i<NAR; i++ {
		go addettoReparto(i)
	}

	go fornitore_c()
	go fornitore_f()

	for i:=0; i<NAR; i++{
		<- done
	}

	termina <- 1
	termina <- 1
	termina <- 1

	<- done
	<- done
	<- done

}