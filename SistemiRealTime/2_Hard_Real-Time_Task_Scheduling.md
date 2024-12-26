# Hard Real Time Task Scheduling<div style="text-align: right">[back](./SistemiRealTime.md)</div>
[Link Piter](https://liveunibo-my.sharepoint.com/:o:/r/personal/pietro_focaccia_studio_unibo_it/_layouts/15/Doc.aspx?sourcedoc=%7BD195ED30-F39F-489F-8CD5-2DEA70483705%7D&file=SOM&action=edit&mobileredirect=true&wdorigin=Sharepoint&RootFolder=%2Fpersonal%2Fpietro_focaccia_studio_unibo_it%2FDocuments%2FSOM&d=wd195ed30f39f489f8cd52dea70483705&e=5%3Afa3c89b441c04712b7ed303d1b15acda&sharingv2=true&fromShare=true&at=9&CID=590ccd21-1d7c-4a1a-b106-5cff1daeaf26)

## Indice

- [Introduzioneback](#introduzioneback)
  - [Indice](#indice)
  - [Sistemi di elaborazione in tempo reale](#sistemi-di-elaborazione-in-tempo-reale)
  - [Sistemi in tempo reale: Aspetti principali](#sistemi-in-tempo-reale-aspetti-principali)
    - [Caratteristiche fondamentali](#caratteristiche-fondamentali)
    - [Obiettivi principali nella progettazione del software](#obiettivi-principali-nella-progettazione-del-software)
  - [Metodologia di progetto](#metodologia-di-progetto)
    - [Approccio "Top-Down"](#approccio-top-down)
  - [Aspetti temporali nei sistemi real-time](#aspetti-temporali-nei-sistemi-real-time)
    - [Vincoli temporali principali](#vincoli-temporali-principali)
    - [Tipologie di schedulazione](#tipologie-di-schedulazione)
      - [Classificazione della schedulazione](#classificazione-della-schedulazione)
  - [Tipologie di processi](#tipologie-di-processi)
  - [Parametri temporali](#parametri-temporali)
    - [Processi periodici:](#processi-periodici)
    - [Processi sporadici:](#processi-sporadici)
  - [Funzioni di utilità di un processo](#funzioni-di-utilità-di-un-processo)
  - [Problematiche nei sistemi real-time](#problematiche-nei-sistemi-real-time)
    - [Tempi di blocco](#tempi-di-blocco)
  - [Soluzioni: Protocolli di accesso](#soluzioni-protocolli-di-accesso)
    - [Obiettivi](#obiettivi)
    - [Strategie](#strategie)

## Assunti

$N$ processi $P_i$ con $i = 1, 2, ..., N$ indipendenti
- Senza vincoli di precedenza
- Senza risorse condivise

Ogni processo $P_j$ con $j = 1, 2, ..., N$
- è periodico, con periodo $T_j$ prefissato
- è caratterizzato da un tempo massimo di esecuzione $C_j$ con $C_j < T_j$
- è caratterizzato da una deadline $D_j$ con $D_j = T_j$

L'esecuzione dei processi è affidata a un sistema di elaborazione monoprocessore. Il tempo impiegato dal processore per operare una commutazione di contesto tra processi è trascurabile.

## Teorema sulla schedulabilità

> Condizione necessaria perchè $N$ precessi siano schedulabili
>
> $U = \sum_{j=1}^{N} U_j = \sum_{j=1}^{N} \frac{C_j}{T_j} \leq 1$

$U$ è il **fattore di utilizzazione** del processore

> Il $j$-esimo termine della sommatoria $C_j/T_j = (C_j(H/T_j)) / H$ rappresenta la frazione dell'iperperiodo $H = mcm(T_1, T_2, ..., T_N)$ impiegata dal processo $P_j$

## Schedulazione clock-driven

Schedulazione di tipo:
- offline
- guaranteed
- non preemptive

associate a processi NP-hard

lo schedule viene fatto su un iperperiodo in istanti decisionali predefiniti