# Lezione 16/09/24

## Intodruzione

## CI SIAMO

## sistemi di elaborazione in tempo reale

Un sistema di elaborazione opera in tempo reale sosltanto se fornisce i risultati attesi entro prestabiliti limiti temporali (che dipendonon dal contesto applicativo)

### obiettivi di progetto?

1. correttezza
2. efficienza
   ...
   n. predicibilità

### contesto applicativo

Packaging valley (macchine automatiche)

Tanti Task con diversi vincoli temporali (tendenzialmente quanche centinia di milli/micro secondi)

## metodologia di progetto Top-Down

### livelli di descrizione

La descrizione del comportamento di un sisitema complesso in ogni ambito ingegneristico è inevitabimente articolata su più livelli.

Ogni livello individua entità astratte o concrete cooperanti.

Esplorando i livelli della gerarchia dall'alto verso il basso aumenta il numero di componenti, diminuisce la complessità

## Aspetti temporali

### prefissate interazioni

1. vincoli
2. riso

### prefissate specifiche temporali

1. freq
2. tempo max
3. deadline

### architettura

1. singlecore
2. multicore

Individuare una opportuna strategia di schedulazione dei porcessi in modo da rispettare tutti i vincoli imposti dalla applicazione.

## Tipologie di scheduling

### off-line

pianificate a priori

### on-line

stabilita a tempo di esecuzione in base a parametri:

1. static
2. dynamic

### guaranteed

se rispetta i vincoli temporali di tutti i processi

### best effort

se tende ad ottimizzare le prestazioni medie dell'insieme di processi

### preemptive

se l'esecuzione di un processo può essere sospesa

### non preemptive

al contrario

## tipologie di processo

### real time

#### hard

se i vincoli temporali devono essere rispettati

1. periodico con freq costante
2. sporadico al contrrario

#### soft

se i vincoli possono essere disattesi in condizione di temporaneo sovraccarico

1. periodico
2. aperiodico

### non real time

## parametri temporali di un processo

- parameti variabili con il tempo (lettere minuscole)

  1. $a_i (r_i)$ arrival (release) time
  2. $d_i$ deadline
  3. $s_i$ start time
  4. $f_i$ finishing time

- parametri indipendenti dal tempo, derivati

  1. $C_i = f_i - s_i$ computation
  2. $D_i = d_i - a_i$ relative deadline
  3. $R_i = f_i - a_i$ response
  4. $L_i = f_i - d_i$ lateness
  5. $E_i = max(0, L_i)$ tardiness
  6. $X_i = D_i - C_i$ slack time (laxity)

### Processi periodici

1. a\_(i+1) - a_i = T period
2. D_i = T
3. a_1 = \phi (phase)

### processi sporadici

1. a\_(i+1) - a_i >= MIT minimum interarrival time
2.

## Funzione di utilita di un precesso

finchè non è hard real time ha utilià anche dopo la deadline, quando diventa hard real time dopo la deadline, l'utilità è nulla, o addirittura meglio non fare niente che farlo in ritardo.

## Paradosso di Graham

A applicazione costituita da 9 processi P_1 ... P_9
di priorità decrescente p_1 > ... > p_9
zio pera
P_1 precede ... P_9

con tempi di esecuzione

...

## preemptive scheduling

...

## Scheduling di processi con risorse condivise

un macello

...

### Inversione di priorità incontrollata

...

$P_M$ che blocca $P_H$ utilizzando una risorsa condivisa eredita la priorità di $P_H$

### problema della concatenazione dei blocchi

...

### ultima slide

...
