package main

import (
	"fmt"
	"math/rand"
	"time"
)


const NS = 3 		//# posti standard
const NM = 3 		//# posti maxi

const NAUTO = 7 	//# di auto
const NCAMPER = 7	//# di camper

var entra_s_discesa = make(chan int,1)		//richiesta spazzaneve discesa
var entra_s_salita = make(chan int,1)		//richiesta spazzaneve salita
var esci_s = make(chan int)					//ricchiesta usicta spazzaneve

var	ack_s = make(chan int)					//conferma spazzaneve 


var entra_a_discesa = make(chan richiesta,10) //richiesta auto discesa
var esci_a_discesa = make(chan int,10)

var entra_a_salita = make(chan int,10)		//richesta auto salita
var esci_a_salita = make(chan int,10)

var ack_a[NAUTO]chan int					//conferma auto


var entra_c_discesa = make(chan int,10)		//richiesta camper discesa
var esci_c_discesa = make(chan int,10)

var entra_c_salita = make(chan int,10)		//richesta camper salita
var esci_c_salita = make(chan int,10)

var ack_c[NCAMPER]chan int					//conferma camper

var termina = make(chan int)
var done = make(chan int)

type richiesta struct {
	id int
	tipo  int
}

func whenAuto(b bool, c chan richiesta) chan richiesta {
	if !b {
		return nil
	}
	return c
}

func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func spazzaneve(){
	for{
		timeout := rand.Intn(70)
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		entra_s_discesa <-1
		
		select {
		case <- ack_s:
		case <- termina:
			fmt.Printf("[Spazzaneve] terminazione\n")
			done <- 1
			return 
		}
		
		timeout = rand.Intn(20)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		esci_s <- 1

		select {
		case <- ack_s:
		case <- termina:
			fmt.Printf("[Spazzaneve] terminazione\n")
			done <- 1
			return 
		}

		timeout = rand.Intn(200)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		entra_s_salita <-1 

		select {
		case <- ack_s:
		case <- termina:
			fmt.Printf("[Spazzaneve] terminazione\n")
			done <- 1
			return 
		}

		timeout = rand.Intn(20)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		esci_s <- 1

		select {
		case <- ack_s:
		case <- termina:
			fmt.Printf("[Spazzaneve] terminazione\n")
			done <- 1
			return 
		}

	}
}


func auto(id int){
	timeout := rand.Intn(50)
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	entra_a_salita <- id
	tipo := <- ack_a[id] //se 0 posto standard se no maxi
	timeout = rand.Intn(10)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	esci_a_salita <- id
	<- ack_a[id]
	timeout = rand.Intn(30)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	entra_a_discesa <- richiesta {id, tipo}
	<- ack_a[id]
	timeout = rand.Intn(10)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	esci_a_discesa <- id
	<- ack_a[id]
	
	done <- 1
}

func camper(id int){
	timeout := rand.Intn(50)
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	entra_c_salita <- id
	<- ack_c[id]
	timeout = rand.Intn(10)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	esci_c_salita <- id
	<- ack_c[id]
	timeout = rand.Intn(30)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	entra_c_discesa <- id
	<- ack_c[id]
	timeout = rand.Intn(10)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	esci_c_discesa <- id
	<- ack_c[id]
	
	done <- 1
}


func castello(){

	var postiS int = NS; 
	var postiM int = NM;	

	var spazza bool = false; 

	var autoS int = 0;		//Auto in discesa
	var autoD int = 0;		//Auto in salita

	var camperS int = 0;	//Camper in salita
	var camperD int = 0;	//Camper in discesa

	for{
		select{
			case <- esci_s :
				fmt.Printf("[Spazzaneve] esce\n") 
				spazza = false
				ack_s <- 1
			case x:= <- esci_c_discesa:
				fmt.Printf("[Camper][%d] esce discesa\n", x)
				camperD --
				ack_c[x] <- 1	
			case x:= <- esci_a_discesa: 
				fmt.Printf("[Auto][%d] esce discesa\n", x)
				autoD --
				ack_a[x] <- 1
			case x:= <- esci_c_salita:
				fmt.Printf("[Camper][%d] esce salita\n", x)
				camperS --
				ack_c[x] <- 1
			case x:= <- esci_a_salita:
				fmt.Printf("[Auto][%d] esce salita\n", x)
				autoS --
				ack_a[x] <- 1
			case <- when( autoS==0 && autoD==0 && camperD==0 && camperS==0 , entra_s_discesa): 
				fmt.Printf("[Spazzaneve] entra discesa\n")
				spazza = true
				ack_s <- 1
			case x:= <- when( autoS==0 && camperS==0 && !spazza &&  len(entra_s_discesa) ==0 , entra_c_discesa):
				fmt.Printf("[Camper][%d] entra discesa\n", x)
				camperD ++
				postiM ++
				ack_c[x] <- 1
			case x:= <- whenAuto( camperS==0 && !spazza && len(entra_s_discesa) ==0 && len(entra_c_discesa) ==0, entra_a_discesa):
				fmt.Printf("[Auto][%d] entra discesa\n", x.id)
				autoD ++
				if x.tipo == 0 {
					postiS ++
				}else {
					postiM ++
				}
				ack_a[x.id] <- 1
			case x:= <-when( camperD==0 && autoD==0 && postiM>0 && !spazza && len(entra_a_discesa)==0 && len(entra_s_discesa) ==0 && len(entra_c_discesa) ==0, entra_c_salita):
				camperS ++
				postiM --
				fmt.Printf("[Camper][%d] entra salita\n", x)
				ack_c[x] <-1
			case x:= <-when( (postiM>0 || postiS>0) && camperD==0 && !spazza && len(entra_c_salita)==0 && len(entra_a_discesa)==0 && len(entra_s_discesa) ==0 && len(entra_c_discesa) ==0 ,entra_a_salita):
				fmt.Printf("[Auto][%d] entra salita\n", x)
				autoS ++
				if (postiS>0){
					postiS --
					ack_a[x] <- 0
				}else{
					postiM --
					ack_a[x] <- 1
				}
			case <- when( autoS==0 && autoD==0 && camperD==0 && camperS==0 && len(entra_a_salita)==0 && len(entra_c_salita)==0 && len(entra_a_discesa)==0 && len(entra_c_discesa) ==0, entra_s_salita):
				fmt.Printf("[Spazzaneve] entra salita\n")
				spazza=true
				ack_s <- 1
			case <-termina: 
			fmt.Println("FINE !!!!!!")
			done <- 1
			return
		}
	}
}





func main() {

	for i := 0; i < NAUTO; i++ {
		ack_a[i] = make(chan int)
	}

	for i := 0; i < NCAMPER; i++ {
		ack_c[i] = make(chan int)
	}

	go castello()
	go spazzaneve()

	for i:=0; i<NAUTO; i++ {
		go auto(i)
	}

	for i:=0; i<NCAMPER; i++ {
		go camper(i)
	}

	for i:=0; i<NAUTO+NCAMPER; i++{
		<- done
	}

	termina <- 1
	termina <- 1

	<- done
	<- done

}