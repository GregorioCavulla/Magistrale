package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxP = 3 //max pneumatici nel deposito
const maxC = 3 //max cerchi nel deposito
const tipoPA = 0
const tipoPB = 1
const tipoCA = 2
const tipoCB = 3
const RobotA = 0
const RobotB = 1
const TOT = 10 //numero totale di auto da montare

var tipoRobot = [2]string{"Modello A", "Modello B"}
var tipoNastro = [4]string{"pneumatico A", "pneumatico B", "cerchio A", "cerchio B"}

var done = make(chan bool)
var terminaDeposito = make(chan bool)

//canali usati dai robot per prelevare pezzi dal deposito
var prelievoPA = make(chan int, 100)
var prelievoPB = make(chan int, 100)
var prelievoCA = make(chan int, 100)
var prelievoCB = make(chan int, 100)

//canali usati dai nastri per depositare pezzi nel deposito
var consegnaPA = make(chan int, 100)
var consegnaPB = make(chan int, 100)
var consegnaCA = make(chan int, 100)
var consegnaCB = make(chan int, 100)

//canali di ack
var ack_robotA = make(chan int)   // canale di ack per il robot relativo al modello A
var ack_robotB = make(chan int)   // canale di ack per il robot relativo al modello B
var ack_nastroPA = make(chan int) // canale di ack per il nastro relativo a PA
var ack_nastroPB = make(chan int) // canale di ack per il nastro relativo a PB
var ack_nastroCA = make(chan int) // canale di ack per il nastro relativo a CA
var ack_nastroCB = make(chan int) // canale di ack per il nastro relativo a CB

func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func Robot(tipo int) {
	var tt int
	fmt.Printf("[Robot %s]: partenza! \n", tipoRobot[tipo])
	var k, ris int

	for { // per ogni auto..
		for i := 0; i < 4; i++ { // per ognuna delle 4 ruote..
			if tipo == 0 { // Robot per modello A
				prelievoCA <- tipo
				ris = <-ack_robotA
				if ris == -1 {
					fmt.Printf("[Robot %s]: termino !\n", tipoRobot[tipo])
					done <- true
					return
				}
				fmt.Printf("[Robot %s]: prelevato cerchio CA\n", tipoRobot[tipo])
				tt = (rand.Intn(2) + 1)
				time.Sleep(time.Duration(tt) * time.Second) //tempo di montaggio cerchio

				prelievoPA <- tipo
				ris = <-ack_robotA
				if ris == -1 {
					fmt.Printf("[Robot %s]: termino!\n", tipoRobot[tipo])
					done <- true
					return
				}
				fmt.Printf("[Robot %s]: prelevato pneumatico PA\n", tipoRobot[tipo])
				tt = (rand.Intn(2) + 1)
				time.Sleep(time.Duration(tt) * time.Second) //tempo di montaggio pneumatico

			} else { // Robot per modello B
				prelievoCB <- tipo
				ris = <-ack_robotB
				if ris == -1 {
					fmt.Printf("[Robot %s]: termino!\n", tipoRobot[tipo])
					done <- true
					return
				}
				fmt.Printf("[Robot %s]: prelevato cerchio CB\n", tipoRobot[tipo])
				tt = (rand.Intn(2) + 1)
				time.Sleep(time.Duration(tt) * time.Second) //tempo di montaggio cerchio

				prelievoPB <- tipo
				ris = <-ack_robotB
				if ris == -1 {
					fmt.Printf("[Robot %s]: termino!\n", tipoRobot[tipo])
					done <- true
					return
				}
				fmt.Printf("[Robot %s]: prelevato pneumatico PB\n", tipoRobot[tipo])
				tt = (rand.Intn(2) + 1)
				time.Sleep(time.Duration(tt) * time.Second) //tempo di montaggio pneumatico
			}
		}
		k++
		fmt.Printf("[Robot %s]: ho completato l'auto n. %d\n", tipoRobot[tipo], k)
	}
}

func nastro(myType int) {
	var tt int
	var i, ris int
	for {
		tt = rand.Intn(2) + 1
		time.Sleep(time.Duration(tt) * time.Second) //tempo di trasporto del pezzo
		switch {
		case myType == 0: //PA
			consegnaPA <- 1
			ris = <-ack_nastroPA
			if ris == -1 {
				fmt.Printf("[nastro %s]:  termino!\n", tipoNastro[myType])
				done <- true
				return
			}
			fmt.Printf("[nastro %s]: consegnato %s \n", tipoNastro[myType], tipoNastro[myType])
		case myType == 1: //PB
			consegnaPB <- 1
			ris = <-ack_nastroPB
			if ris == -1 {
				fmt.Printf("[nastro %s]:  termino!\n", tipoNastro[myType])
				done <- true
				return
			}
			fmt.Printf("[nastro %s]: consegnato %s \n", tipoNastro[myType], tipoNastro[myType])
		case myType == 2: //CA
			consegnaCA <- 1
			ris = <-ack_nastroCA
			if ris == -1 {
				fmt.Printf("[nastro %s]:  termino!\n", tipoNastro[myType])
				done <- true
				return
			}
			fmt.Printf("[nastro %s]: consegnato %s \n", tipoNastro[myType], tipoNastro[myType])
		case myType == 3: //CB
			consegnaCB <- 1
			ris = <-ack_nastroCB
			if ris == -1 {
				fmt.Printf("[nastro %s]:  termino!\n", tipoNastro[myType])
				done <- true
				return
			}
			fmt.Printf("[nastro %s]: consegnato %s \n", tipoNastro[myType], tipoNastro[myType])
		default:
			fmt.Printf("[nastro %d]: inizializzato erroneamente, termino!\n", myType)
		}
		i++
	}

}

func deposito() {
	var numCAMontati int = 0
	var numCBMontati int = 0
	var numPAMontati int = 0
	var numPBMontati int = 0

	var numAMontati int = 0
	var numBMontati int = 0

	var numCA int = 0
	var numCB int = 0
	var numPA int = 0
	var numPB int = 0

	var totP int = 0
	var totC int = 0
	var fine bool = false // diventa true quando sono stati completati i montaggi di TOT auto
	for {

		select {
		case <-when(fine == false && (totC < maxC && numCA < maxC-1) && (numAMontati < numBMontati ||
			(numAMontati >= numBMontati && len(consegnaCB) == 0)), consegnaCA):
			numCA++
			totC++
			ack_nastroCA <- 1
			fmt.Printf("[deposito]: aggiunto cerchio A, ci sono %d CA e %d CB -> tot cerchi nel deposito = %d\n", numCA, numCB, totC)
		case <-when(fine == false && (totC < maxC && numCB < maxC-1) && (numAMontati >= numBMontati ||
			(numAMontati < numBMontati && len(consegnaCA) == 0)), consegnaCB):
			numCB++
			totC++
			ack_nastroCB <- 1
			fmt.Printf("[deposito]: aggiunto cerchio B, ci sono %d CA e %d CB -> tot cerchi nel deposito = %d\n", numCA, numCB, totC)
		case <-when(fine == false && (totP < maxP && numPA < maxP-1) && (numAMontati < numBMontati ||
			(numAMontati >= numBMontati && len(consegnaPB) == 0)), consegnaPA):
			numPA++
			totP++
			ack_nastroPA <- 1
			fmt.Printf("[deposito]: aggiunto pneumatico A, ci sono %d PA e %d PB -> tot pneumatici nel deposito = %d\n", numPA, numPB, totP)
		case <-when(fine == false && (totP < maxP && numPB < maxP-1) && (numAMontati >= numBMontati ||
			(numAMontati < numBMontati && len(consegnaPA) == 0)), consegnaPB):
			numPB++
			totP++
			ack_nastroPB <- 1
			fmt.Printf("[deposito]: aggiunto pneumatico B, ci sono %d PA e %d PB -> tot pneumatici nel deposito = %d\n", numPA, numPB, totP)
		case <-when(fine == false && numCA > 0 && (numAMontati < numBMontati ||
			(numAMontati >= numBMontati && len(prelievoCB) == 0)), prelievoCA):
			numCA--
			totC--
			numCAMontati++
			ack_robotA <- 1
			fmt.Printf("[deposito]: prelevato cerchio A, tot cerchi nel deposito = %d\n", totC)
		case <-when(fine == false && numCB > 0 && (numAMontati >= numBMontati ||
			(numAMontati < numBMontati && len(prelievoCA) == 0)), prelievoCB):
			numCB--
			totC--
			numCBMontati++
			ack_robotB <- 1
			fmt.Printf("[deposito]: prelevato cerchio B, tot cerchi nel deposito = %d\n", totC)
		case <-when(fine == false && numPA > 0 && (numAMontati < numBMontati ||
			(numAMontati >= numBMontati && len(prelievoPB) == 0)), prelievoPA):
			numPA--
			totP--
			numPAMontati++
			ack_robotA <- 1
			fmt.Printf("[deposito]: prelevato pneumatico A, tot pneumatici nel deposito = %d\n", totC)
		//prelievo pneumatico B
		case <-when(fine == false && numPB > 0 && (numAMontati >= numBMontati ||
			(numAMontati < numBMontati && len(prelievoPA) == 0)), prelievoPB):
			numPB--
			totP--
			numPBMontati++
			ack_robotB <- 1
			fmt.Printf("[deposito]: prelevato pneumatico B, tot pneumatici nel deposito = %d\n", totC)

		//terminazione
		case <-when(fine == true, consegnaCA):
			ack_nastroCA <- -1
		case <-when(fine == true, consegnaCB):
			ack_nastroCB <- -1
		case <-when(fine == true, consegnaPA):
			ack_nastroPA <- -1
		case <-when(fine == true, consegnaPB):
			ack_nastroPB <- -1
		case <-when(fine == true, prelievoCA):
			ack_robotA <- -1
		case <-when(fine == true, prelievoCB):
			ack_robotB <- -1
		case <-when(fine == true, prelievoPA):
			ack_robotA <- -1
		case <-when(fine == true, prelievoPB):
			ack_robotB <- -1

		case <-terminaDeposito:
			fmt.Printf("[deposito]: termino\n")
			done <- true
			return
		}
		//conteggio auto montate modello A e B e verifica per terminazione
		if numCAMontati == 4 && numPAMontati == 4 {
			numAMontati++
			numCAMontati = 0
			numPAMontati = 0
		}
		if numCBMontati == 4 && numPBMontati == 4 {
			numBMontati++
			numCBMontati = 0
			numPBMontati = 0
		}
		fmt.Printf("[deposito] montate A = %d. montate B = %d\n", numAMontati, numBMontati)
		if numAMontati+numBMontati == TOT {
			fine = true
		}

	}
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Printf("[main] attivo 4 nastri trasportatori e 2 robot\n")

	go deposito()

	for i := 0; i < 4; i++ {
		go nastro(i)
	}

	for i := 0; i < 2; i++ {
		go Robot(i)
	}

	for i := 0; i < 4; i++ { //terminazione nastri
		<-done
	}

	for i := 0; i < 2; i++ { //terminazione robot
		<-done
	}

	terminaDeposito <- true
	<-done
	fmt.Printf("[main] APPLICAZIONE TERMINATA \n")
}
