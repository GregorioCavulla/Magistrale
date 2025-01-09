# Hard Real Time Process Scheduling

## Primo modello per lo scheduling di processi periodici

N processi $P_1, P_2, ..., P_N$ indipendenti

- senza vincoli di precedenza
- senza risorse condivise

Ogni processo $P_j$:

- ha un periodo $T_j$
- ha un tempo di esecuzione $C_j$ con $C_j < T_j$
- caratterizzato da una Deadline $D_j = T_j$

L'esecuzione è affidata a un sistema di elaborazione monoprocessore.

Il tempo impiegato dal processore per commutare di contesto tra due processi è trascurabile.

## Teorema sulla schedulabilità

## I requisiti temporali sono coerenti e consistenti?

**Condizione necessaria (ma in generale non sufficiente) affinché un insieme di N processi periodici sia schedulabile è che il risultante fattore di utilizzazione del processore sia non superiore a 1**

$$U = \sum_{j=1}^{N} U_j = \sum_{j=1}^{N} \frac{C_j}{T_j} \leq 1$$

Il $j^{mo}$ termine della sommatoria $Cj / Tj = (C_j (H / T_j)) / H$
rappresenta la frazione dell’iperperiodo $H = mcm (T_1, T_2, ..., T_N)$
richiesta per l’esecuzione di $P_j$.

## Schedulazione CLOCK-DRIVEN

Schedulazione di tipo:

- off-line
- guaranteed
- non pre-emptive

semplice da realizzare e efficiente (l'overhead è trascurabile), non è idonea in contesti che implicano dinamicità e flessibilità.

I parametri temporali dei processi sono noti a priori e non soggetti a variazioni a run-time.

Tutti i vincoli vengono soddisfatti a priori. **Problema NP-hard**.

Allo scopo possono essere usati algoritmi anche complessi, senza incorrere in una penalizzazione delle prestazioni conseguibili a run-time.

Lo schedule viene costruito con riferimento al generico iperperiodo, assumendo
che qualunque decisione riguardante la schedulazione dei processi venga presa in
corrispondenza di predefiniti “istanti decisionali”. Tali istanti possono essere o
meno equidistanziati.

Lo schedule risultante è di norma esplicitato in termini di una tabella, la cui generica k-esima entry è del tipo:

- $(\Delta t_k, P(tk))$ nel caso di istanti decisionali non equidistanziati, con $\Delta t_k =t_k+1-t_k$
- $(P(t_k))$ nel caso di istanti decisionali equidistanziati, dove $P(t_k)$ e $P(t_k)$ identificano rispettivamente il processo o l’insieme di processi da schedulare all’istante $t_k$.

Lo scheduler a run-time, avvalendosi di un timer hardware che, opportunamente programmato, genera un interrupt in corrispondenza di ogni istante decisionale, ciclicamente si limita a interpretare il contenuto di tale tabella, procedendo a riprogrammare il timer se necessario (ovvero nel caso di istanti decisionali non equidistanziati), a selezionare il processo o i processi da schedulare, a porsi quindi in attesa del successivo interrupt.

## Scheduler TIME-DRIVEN

...

## Cyclic Executive

...

## Ambiente di esecuzione

...

## Ambienti di esecuzione sequenziale

...

## Approccio cyclic executive

periodo T multiplo del periodo del task precedente = relazione armonica tra i task

## Identificazione del frame size

## Pianificazione

## Costruzione di un feasible schedule
