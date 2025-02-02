package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXP = 20 //Max dosi
const MAXM = 20 //Max dosi

const RROSSE = 3 	 	//Numero regioni rosse
const RARANCIONI = 3 	//Numero regioni arancioni
const RGIALLE = 3	 	//Numero regioni gialle

const NL = 10 	//N dosi x lotto
const Q = 3 	//N dosi x regione

const TOTP = 24
const TOTM = 24

var prenotaP = make(chan int)
var prenotaM = make(chan int)

var ackP = make(chan int)
var ackM = make(chan int)

var prelR = make(chan int, 10)
var prelA = make(chan int, 10)
var prelG = make(chan int, 10)

var ack[RROSSE + RARANCIONI + RGIALLE]chan int

var termina = make(chan int)
var done = make(chan int)


func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func pfizer(){
	for{
		timeout := rand.Intn(50)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		prenotaP <-1
		select {
			case <- ackP:
				fmt.Printf("[pfizer] consegna\n")
			case <- termina:
				fmt.Printf("[pfizer] terminazione\n")
				done <- 1
				return 
		}
	}
}

func moderna(){
	for{
		timeout := rand.Intn(50)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		prenotaM <-1
		select {
			case <- ackM:
				fmt.Printf("[moderna] consegna\n")
			case <- termina:
				fmt.Printf("[moderna] terminazione\n")
				done <- 1
				return 
		}
	}
}

func regioneR(id int){

	for {
		timeout := rand.Intn(50)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		prelR <- id
		select{
			case <- termina:
				fmt.Printf("[regioneR][%d] terminazione\n", id)
				done <- 1
			case tipo := <- ack[id]:
				if tipo == 0{
					fmt.Printf("[regioneR][%d] prelevato vacino pfizer\n", id)
				}else {
					fmt.Printf("[regioneR][%d] prelevato vacino moderna\n", id)
				}
						
		}
		fmt.Printf("[regioneR][%d] terminazione\n", id)
	}
	done <- 1
}

func regioneA(id int){

	for {
		timeout := rand.Intn(50)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		prelA <- id
		select{
			case <- termina:
				fmt.Printf("[regioneA][%d] terminazione\n", id)
				done <- 1
			case tipo := <- ack[id]:
				if tipo == 0{
					fmt.Printf("[regioneA][%d] prelevato vacino pfizer\n", id)
				}else {
					fmt.Printf("[regioneA][%d] prelevato vacino moderna\n", id)
				}
						
		}
		fmt.Printf("[regioneA][%d] terminazione\n", id)
	}
	done <- 1
}

func regioneG(id int){

	for {
		timeout := rand.Intn(50)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		prelG <- id
		select{
			case <- termina:
				fmt.Printf("[regioneG][%d] terminazione\n", id)
				done <- 1
			case tipo := <- ack[id]:
				if tipo == 0{
					fmt.Printf("[regioneG][%d] prelevato vacino pfizer\n", id)
				}else {
					fmt.Printf("[regioneG][%d] prelevato vacino moderna\n", id)
				}
						
		}
		fmt.Printf("[regioneG][%d] terminazione\n", id)
	}
	done <- 1
}

func deposito() {

	var postiP int = MAXP
	var postiM int = MAXM

	var totP int = 0
	var totM int = 0

	for {
		if totP > TOTP && totM > TOTM {
			break
		}
		select {
		case <- when((postiP - NL) > 0,prenotaP):	//Consegna PFIZER
			postiP -= NL
			ackP <- 1
		case <- when((postiM - NL) > 0,prenotaM):	//Consegna MODERNA
			postiM -= NL
			ackM <- 1
		case x:= <- when((postiM + Q < MAXM) || (postiP +Q < MAXP) || (MAXM+MAXP-postiM-postiP >= Q), prelR):	//Richiesta ROSSA
			if postiM + Q < MAXM {
				postiM += Q
				totM += Q
				ack[x] <- 1
			} else if postiP +Q < MAXP{
				postiP += Q
				totP += Q
				ack[x] <- 0
			} else {
				totP += Q-(MAXM-postiM)
				totM += MAXM-postiM
				postiP += Q-(MAXM-postiM)
				postiM = MAXM 
			}
		case x:= <- when(((postiM + Q < MAXM) || (postiP +Q < MAXP) || (MAXM+MAXP-postiM-postiP >= Q)) && len(prelR)== 0, prelA):	//Richiesta ARANCIONE
			if postiM + Q < MAXM {
				postiM += Q
				totM += Q
				ack[x] <- 1
			} else if postiP +Q < MAXP{
				postiP += Q
				totP += Q
				ack[x] <- 0
			} else {
				totP += Q-(MAXM-postiM)
				totM += MAXM-postiM
				postiP += Q-(MAXM-postiM)
				postiM = MAXM 
			}	
		case x:= <- when(((postiM + Q < MAXM) || (postiP +Q < MAXP) || (MAXM+MAXP-postiM-postiP >= Q)) && len(prelR)== 0 && len(prelA)== 0, prelG):	//Richiesta GIALLA
			if postiM + Q < MAXM {
				postiM += Q
				totM += Q
				ack[x] <- 1
			} else if postiP +Q < MAXP{
				postiP += Q
				totP += Q
				ack[x] <- 0
			} else {
				totP += Q-(MAXM-postiM)
				totM += MAXM-postiM
				postiP += Q-(MAXM-postiM)
				postiM = MAXM 
			}
		}
	}

	fmt.Printf("[deposito] terminazione\n",)
	done <- 1 
}


func main(){

	go deposito()
	go pfizer()
	go moderna()

	for i := 0; i < RARANCIONI+RGIALLE+RROSSE; i++ {
		ack[i] = make(chan int, 2)
	}


	for i := 0; i<RGIALLE ; i++{
		go regioneG (i)
	}

	for i := RGIALLE; i<RARANCIONI+RGIALLE ; i++{
		go regioneA (i)
	}

	for i := RROSSE; i<RARANCIONI+RGIALLE+RROSSE ; i++{
		go regioneR (i)
	}

	<- done

	for i := 0; i<RARANCIONI+RGIALLE+RROSSE+2 ; i++{
		termina <- 1
	}

	for i := 0; i<RARANCIONI+RGIALLE+RROSSE+2 ; i++{
		<- done
	}


}


