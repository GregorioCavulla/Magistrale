package main

import (
	"fmt"
	"math/rand"
	"time"
)

const LITRI = 30	//N litri
const M1 = 5 		//Max monete da 10 cent
const M2 = 5 		//Max monete da 20 cent

const NCLITTADINIPICCOLI = 2	//# di cittadini che devono riempire bottiglie piccola
const NCITTADINIGRANDI = 2		//# di cittadini che devono riempire bottiglie grande

var manutenzione = make(chan int, 1)	//richiesta accesso addetto	
var ackA = make(chan int)			//conferma addetto

var piccoli = make(chan int, 10)		//richiesta riempimento bottiglie piccole
var grandi = make(chan int, 10)			//richiesta riempimento bottiglie grandi

var ack[NCITTADINIGRANDI + NCLITTADINIPICCOLI]chan int	//conferma erogazione

var termina = make(chan int)
var done = make(chan int)



func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func addetto(){
	for{
		timeout := rand.Intn(200)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		fmt.Printf("[Addetto] richiesta manutenzione\n")
		manutenzione <-1
		select {
			case <- ackA:
				fmt.Printf("[Addetto] fine manutenzione\n")
			case <- termina:
				fmt.Printf("[Addetto] terminazione\n")
				done <- 1
				return 
		}
	}
}

func cittadiniP(id int){
	timeout := rand.Intn(50)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	piccoli <- id
	<- ack[id]
	fmt.Printf("[Cittadino Piccolo][%d] terminazione\n", id)
	done <- 1
}

func cittadiniG(id int){
	timeout := rand.Intn(50)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	grandi <- id
	<- ack[id]
	fmt.Printf("[Cittadino Grande][%d] terminazione\n", id)
	done <- 1
}

func distributore(){

	var acqua float32 = LITRI	//# di litri disponibili
	var c1 int = M1				//# di posti da 10 cent disponibili
	var c2 int = M2				//# di posti da 20 cent disponibili

	for{
		select{
			case <- when((acqua==0 || c1 == 0 || c2 ==0) || (len(piccoli)==0 && len(grandi)==0),manutenzione):
				fmt.Printf("[Distributore] inizio manutenzione\n")
				timeout := rand.Intn(50)
				time.Sleep(time.Duration(timeout) * time.Millisecond)
				acqua = LITRI
				c1 = M1
				c2 = M2
				ackA <- 1
			case x:= <- when(((acqua-0.5)>0 && c1>0) && (len(manutenzione)==0 || c2>0),piccoli):
				fmt.Printf("[Distributore] erogo acqua piccola a %d\n", x)
				timeout := rand.Intn(50)
				time.Sleep(time.Duration(timeout) * time.Millisecond)
				acqua -= 0.5
				c1 --
				ack[x] <-1
			case x:= <- when(((acqua-1.5)>0 && c2>0 && len(piccoli)==0) && (len(manutenzione)==0 || c1>0),grandi):
				fmt.Printf("[Distributore] erogo acqua grande a %d\n", x)
				timeout := rand.Intn(50)
				time.Sleep(time.Duration(timeout) * time.Millisecond)
				acqua -= 1.5
				c2 --
				ack[x] <-1
			case <-termina: 
			fmt.Println("FINE !!!!!!")
			done <- 1
			return
		}
	}
}

func main() {

	//inizializzazione canali
	for i := 0; i < NCITTADINIGRANDI + NCLITTADINIPICCOLI; i++ {
		ack[i] = make(chan int, 1)
	}

	go distributore()
	go addetto()

	for i:= 0 ; i<NCITTADINIGRANDI ; i++{
		go cittadiniG(i)
	}

	for i:= NCITTADINIGRANDI ; i<NCLITTADINIPICCOLI+NCITTADINIGRANDI ; i++{
		go cittadiniP(i)
	}

	for i:= 0 ; i<NCLITTADINIPICCOLI+NCITTADINIGRANDI ; i++{
		<- done
	}

	termina <- 1
	termina <- 1

	<- done
	<- done
}