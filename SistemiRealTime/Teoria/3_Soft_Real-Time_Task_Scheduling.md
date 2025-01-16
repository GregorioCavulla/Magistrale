# Soft Real Time Task Scheduling

[Link Piter](https://liveunibo-my.sharepoint.com/:o:/r/personal/pietro_focaccia_studio_unibo_it/_layouts/15/Doc.aspx?sourcedoc=%7BD195ED30-F39F-489F-8CD5-2DEA70483705%7D&file=SOM&action=edit&mobileredirect=true&wdorigin=Sharepoint&RootFolder=%2Fpersonal%2Fpietro_focaccia_studio_unibo_it%2FDocuments%2FSOM&d=wd195ed30f39f489f8cd52dea70483705&e=5%3Afa3c89b441c04712b7ed303d1b15acda&sharingv2=true&fromShare=true&at=9&CID=590ccd21-1d7c-4a1a-b106-5cff1daeaf26)

[Return](./SistemiRealTime.md)

---

## Indice

- [Hard Real Time Task Schedulingback](#hard-real-time-task-schedulingback)
  - [Indice](#indice)
  - [Assunti](#assunti)
  - [Teorema sulla schedulabilità](#teorema-sulla-schedulabilità)
  - [Schedulazione clock-driven](#schedulazione-clock-driven)
  - [Timer-driven scheduling](#timer-driven-scheduling)
  - [Ambiente di esecuzione](#ambiente-di-esecuzione)
    - [Sequenziale](#sequenziale)
  - [Cyclic Executive](#cyclic-executive)
    - [Approccio Cyclic Executive](#approccio-cyclic-executive)
    - [Approccio Barker - Shaw](#approccio-barker---shaw)
    - [Costruzione di un feasable schedule](#costruzione-di-un-feasable-schedule)
  - [Schedulazione Priority Driven](#schedulazione-priority-driven)
    - [Algoritmo Rate Monotonic Priority Ordering (RMPO)](#algoritmo-rate-monotonic-priority-ordering-rmpo)
  - [Test di schedulabilità LIU-LAYLAND](#test-di-schedulabilità-liu-layland)
    - [Corollario](#corollario)
  - [Test di Kuo - Mok](#test-di-kuo---mok)
  - [Test di Burchard](#test-di-burchard)
  - [Test di Han](#test-di-han)
  - [Analisi di schedulabilità di Audsley](#analisi-di-schedulabilità-di-audsley)
    - [Alternativa all'algoritmo di Audsley](#alternativa-allalgoritmo-di-audsley)
  - [Processi Sporadici](#processi-sporadici)
  - [Deadline Monotonic Priority Ordering (DMPO)](#deadline-monotonic-priority-ordering-dmpo)
  - [Analisi di schedulabilità attraverso i tempi di risposta](#analisi-di-schedulabilità-attraverso-i-tempi-di-risposta)
    - [Altro test](#altro-test)
  - [Test di Lehoczky](#test-di-lehoczky)
  - [Test di Utilizzazione Efficace dei Processi](#test-di-utilizzazione-efficace-dei-processi)
  - [Algoritmo Earliest Deadline First (EDF)](#algoritmo-earliest-deadline-first-edf)
  - [Algoritmo Least Slack Time First (LST)](#algoritmo-least-slack-time-first-lst)
  - [Metascheduler](#metascheduler)

## Introduzione

Ai task Hard Real Time si aggiungono quelli Soft Real Time, l'obiettivo è velocizzare la loro esecuzione.

## Modalità di getsione

- **Background**
- **Servizio tramite server**:
  Un task HRT va eseguito prima della deadline non per forza subito. Ritardiamo l'esecuzione degli HRT per eseguire un po di SRT.
  Generando un nuovo task HRT (Server) che gestisce gli SRT.
  Il server è un task da dimensionare correttamente per non scavallare gli altri HRT.

- **Priorità statica**
  - Polling server
  - Deferrable server
  - priority exchange server
  - sporadic server
- **Priorità dinamica**
  - constant utilization server
  - total bandwidth server
  
### Servizio in background

![alt text](image-16.png)

Lerichieste aperiodiche sono definite da vincoli temporali soft o hard real time.

La strategia di schedulazione dei processi aperiodici, tipicamente FCFS (First Come First Serve) è indipendente da quella selezionata per gli altri processi.

In questa situazione il tempo di risposta degli SRT non è ottimale.

### Server a priorità statica

I task periodici (HRT) tra cui il server sono schedulati con priorità statica, solitamente RMPO.

![alt text](image-17.png)

Il server ha parametri $T_s$ (periodo) e $C_s$ (Tempo di esecuzione) da dimensionare correttamente per rispettare le deadline degli hrt.

Con questa strategia i tempi di SRT migliorano rispetto a **background**, che vengono schedulati conalgoritmi differenti in base al pito di **Server**.

### Polling Server

Ad ogni periodo $T_s$ viene assegnato un tempo di computazione $C_s$, consumato dall'esecuzione degli SRT ed inizializzato ad ogni ciclo, se non sono presenti task aperiodici viene buttato via.

Essendo anche il serverperiodico possiamo procedere adimensionare $C_s$ e $T_s$ basandoci su regole di schedulabiilità dei task periodici:

$$
U_p + U_s \leq U_{RMPO} (N+1) = (N+1) (2^{\frac{1}{(N+1)}}-1)

U_s \leq (N+1)(2^{\frac{1}{(N+1)}}-1) - U_p =_{N\rarr+\infty} \ln2 -U_p
$$

Valore che rimane conservativo, si può fare di meglio se si considera il server con priorità massima:

$$
U_s\leq\frac{2}{(1+\frac{U_p}{N})^N}-1 =_{N\rarr+\infty} \frac{2}{e^{U_p}}-1
$$

### Deferrable Server

$C_s$ Ripristinato ad ogni $T_s$ (come polling), però $C_s$ viene considerato in assenza di task SRT.

**Problema:** Il contributo del server al fattore di utilizzazione dell'applicazione, in condizioni di carico massimo, è maggiore di quello di un processo periodico, ovvero $U_s>\frac{C_s}{T_s}$.

Il deferrable server risulta più invasivo nei confronti degli HRT, qundo ha priorità massima, conservando $C_s$

Dimensionamento Server con priorità massima:

$$

U_s\leq\frac{2-(1+\frac{U_p}{N})^N}{2(1+\frac{U_p}{N})^N)-1}=_{N\rarr+\infty}\frac{2-e^{U_p}}{2e^{U_p}-1}

$$

Se la coda degli SRT è sempre piena meglio PS, se no DF.

### Priority Exchange Server

$C_s$ può essere accumulata ad ogni livello di priorità inferiore a quella del server.

Ogni $T_s$ viene ripristinata $C_s$ al livello di priorità del server, $C_s$ a priorità inferiore viene mantenuta.

$C_s$ al massimo livello di priorità disponibilie in assenza di SRT e in pr(esenza di HRT a priorità inferiore viene trasferita al livello di HRT in esecuzione.

Se presente, un SRT e il server possiede $C_s$ allora avrà priorità ad ogni livello in cui ha $C_s$.

**Dimensionamento server:**

$$

U_s\leq(N+1)(2^{\frac{1}{(N+1)}-1}-U_p)=_{N\rarr+\infty}\ln2-U_p

$$

Demensionamento server con priorità massima:

$$

U_s\leq\frac{2}{(1+\frac{U_p}{N})^N}-1=_{N\rarr+\infty}\frac{2}{e^{U_p}}-1

$$

### Sporadic Server (rotto)

$C_s$ disponibile conservato se non ci sono richieste aperiodiche.

$C_s$ consumato solo dal servizio di richieste aperiodiche

$C_s$ Reintegrato solo dopo essere consumato e in misura di quanto è stato consumato.

$$

R_T \text{ (istante di reintegro)} = max\{b_1 + T_s, T_D\}  

$$

> In questa forma, lo Sporadic Server può generare problemi agli HRT causati dalla formulazione di $R_T$.

### Sporadi Server (corretto)

Quando il carico di richieste aperiodiche non comporta al totale consumo di $C_s$ del server, le distinte porzioni (chunks) di capacità: quella residua e quella consumata, devono essere gestite separatamente.

![alt text](image-18.png)

- $t_e = t_a$ perchè $C_s$ è massima
- $t_e$ deriva da $R_T$ precedente

### Server a Priorità Dinamica

![alt text](image-19.png)

Il server e gli altri processi sono schedulati con EDF (Earliest Deadline First).

### Constant Utilization Server

Il server rimane in attesa, ogni volta che arriva un task aperiodico vanno calcolati $C_s$ e $d_s$.

$$

C_s - C_{Ra}
\\
d_s = t_{Ra} + \frac{C_s}{U_s}

$$

Il processo poi viene eseguito in accordo con la priorità data da $d_s$ e $C_s$.

