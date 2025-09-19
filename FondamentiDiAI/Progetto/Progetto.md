# Attività progettuale Fondamenti di AI

## Indice
- [Attività progettuale Fondamenti di AI](#attività-progettuale-fondamenti-di-ai)
  - [Indice](#indice)
  - [1. Introduzione](#1-introduzione)
    - [Obiettivi del progetto](#obiettivi-del-progetto)
    - [Regole del gioco](#regole-del-gioco)
  - [2. Modellizzazione del problema](#2-modellizzazione-del-problema)
    - [Griglia e regioni](#griglia-e-regioni)
    - [Rappresentazione della soluzione](#rappresentazione-della-soluzione)
    - [Vincoli del problema](#vincoli-del-problema)
    - [Unicità della soluzione](#unicità-della-soluzione)
  - [3. Generazione delle regioni](#3-generazione-delle-regioni)
    - [Algoritmo Seed-Growth](#algoritmo-seed-growth)
    - [Garanzia di connettività e bilanciamento](#garanzia-di-connettività-e-bilanciamento)
  - [4. Generazione del livello](#4-generazione-del-livello)
    - [Controllo risolvibilità con Prolog](#controllo-risolvibilità-con-prolog)
    - [Controllo unicità della soluzione](#controllo-unicità-della-soluzione)
    - [Output in JSON](#output-in-json)
  - [5. Interfaccia utente](#5-interfaccia-utente)
    - [Visualizzazione griglia](#visualizzazione-griglia)
    - [Interazione utente](#interazione-utente)
    - [Condizioni di vittoria](#condizioni-di-vittoria)
  - [6. Conclusioni e sviluppi futuri](#6-conclusioni-e-sviluppi-futuri)
    - [Stato attuale](#stato-attuale)
    - [Possibili estensioni](#possibili-estensioni)

## 1. Introduzione

### Obiettivi del progetto
Il progetto ha come obiettivo la realizzazione di un **generatore di livelli** per un gioco logico ispirato al classico problema delle *N-Queens*, ma arricchito da un vincolo aggiuntivo basato sulla suddivisione della griglia in regioni.

### Regole del gioco
Il gioco si svolge su una matrice di dimensione $N \times N$, suddivisa in **N regioni connesse** ciascuna caratterizzata da un colore o da un identificativo numerico.
Il giocatore deve posizionare **N regine** rispettando i seguenti vincoli:

* una regina per ogni riga,
* una regina per ogni colonna,
* una regina per ogni regione.

Non sono previsti vincoli sulle diagonali, a differenza della formulazione classica delle *N-Queens*.
L’obiettivo finale è ottenere un puzzle con **soluzione unica**, così da garantire una sfida ben definita al giocatore.

Il generatore è implementato in **Prolog**, sfruttando tecniche di *constraint programming* per la verifica della risolvibilità e dell’unicità. L’output viene esportato in formato **JSON**, così da poter essere caricato in un’interfaccia utente che gestisce la parte di gioco.

---

## 2. Modellizzazione del problema

### Griglia e regioni
La griglia è una matrice $N \times N$. Ogni cella è identificata da una coppia di coordinate $(r, c)$, con $r$ = indice di riga e $c$ = indice di colonna.
Ogni cella appartiene a una **regione**, indicata da un ID numerico compreso tra 1 e $N$. Una regione è definita come un insieme di celle **connesse ortogonalmente** (adiacenza su/giù/sinistra/destra). L’insieme delle regioni deve coprire l’intera griglia senza sovrapposizioni.

### Rappresentazione della soluzione
La posizione delle regine è modellata tramite un **vettore di variabili**:

$$
Cols = [C_1, C_2, \dots, C_N]
$$

dove $C_r$ indica la colonna in cui è posizionata la regina della riga $r$.
Questa rappresentazione garantisce automaticamente che vi sia **esattamente una regina per riga**.

### Vincoli del problema
I vincoli logici si esprimono come segue:

* **Colonne**: l’elenco dei valori in `Cols` deve essere *all-different*, cioè non possono esserci due righe con regina nella stessa colonna.
* **Regioni**: per ogni riga $r$, dalla coppia $(r, C_r)$ si ricava l’ID della regione. Anche in questo caso, l’elenco degli ID deve essere *all-different*, cioè ciascuna regione deve contenere esattamente una regina.

In sintesi, una configurazione è valida se e solo se:

* ogni riga ospita una regina,
* ogni colonna ospita una sola regina,
* ogni regione ospita una sola regina.

Questa formalizzazione consente di ridurre il problema a un classico caso di *constraint satisfaction problem* (CSP), risolvibile in Prolog con la libreria CLP(FD).

### Unicità della soluzione
Nel progetto distinguiamo due forme di unicità della soluzione:

**Unicità della disposizione:** esiste una sola configurazione finale di $N$ regine che soddisfa i vincoli. Questa è la nozione adottata nel progetto, poiché garantisce che il livello non sia ambiguo e che la sfida abbia una sola risposta corretta.

**Unicità del percorso:** esiste un unico ordine di mosse “forzato” per arrivare alla soluzione, senza alternative equivalenti. Questa nozione riguarda la difficoltà del puzzle ed è lasciata come possibile sviluppo futuro.

Per questa prima versione, il generatore e il solver sono progettati per garantire soltanto l’unicità della disposizione.

## 3. Generazione delle regioni
### Algoritmo Seed-Growth
### Garanzia di connettività e bilanciamento

## 4. Generazione del livello
### Controllo risolvibilità con Prolog
### Controllo unicità della soluzione
### Output in JSON

## 5. Interfaccia utente
### Visualizzazione griglia
### Interazione utente
### Condizioni di vittoria

## 6. Conclusioni e sviluppi futuri
### Stato attuale
### Possibili estensioni