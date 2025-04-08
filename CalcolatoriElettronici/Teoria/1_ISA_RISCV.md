
# ISA RISCV <div style="text-align: right"> [back](./CalcolatoriElettronici.md) </div>

# Indice

- [ISA RISCV  back ](#isa-riscv--back-)
- [Indice](#indice)
  - [RISC-V ISA nomenclatura](#risc-v-isa-nomenclatura)
  - [ISA RISC-V](#isa-risc-v)
  - [Ripasso Concetti Base - Pipeline RISC-V](#ripasso-concetti-base---pipeline-risc-v)
  - [Pipeline RISC-V a 5 stadi](#pipeline-risc-v-a-5-stadi)
  - [Concetti Principali](#concetti-principali)
  - [Soluzioni alle Alee](#soluzioni-alle-alee)
  - [Prestazioni Pipeline](#prestazioni-pipeline)
  - [\[\[2.2\_Modello\_di\_Esecuzione\_Cuda\_GC\]\]](#22_modello_di_esecuzione_cuda_gc)
  - [Ripasso concetti base](#ripasso-concetti-base)
  - [Data Hazards in ALU Instructions](#data-hazards-in-alu-instructions)
  - [Rilevazione della necessità di Forwarding](#rilevazione-della-necessità-di-forwarding)
  - [Condizioni di Forwarding](#condizioni-di-forwarding)
  - [Double Data Hazard](#double-data-hazard)
  - [Dynamic Branch Prediction](#dynamic-branch-prediction)
  - [1-Bit Predictor: Problema](#1-bit-predictor-problema)
  - [2-Bit Predictor](#2-bit-predictor)
  - [Calcolo del Branch Target](#calcolo-del-branch-target)
- [\[\[Lezione\_3\]\]](#lezione_3)
  - [Intructions Level Parallelism (ILP)](#intructions-level-parallelism-ilp)
  - [Data Level Parallelism (DLP)](#data-level-parallelism-dlp)
  - [Thread Level Parallelism (TLP)](#thread-level-parallelism-tlp)
  - [Data Dependance](#data-dependance)
  - [Alee di Dato](#alee-di-dato)
 

## RISC-V ISA nomenclatura

- **Caratteristiche RISC-V:**
  - Architettura semplice e estensibile, adatta a vari ambiti (IoT, HPC).
  - Formato: `RV + word-width + estensioni` (es. RV32IMC: 32 bit, integer, moltiplicazione, compressione).
  - **Specifiche:**
    - **User-level:** solo l'estensione I è obbligatoria.
    - **Privileged-level:** in sviluppo, include funzioni del sistema operativo (interrupt, gestione virtuale degli indirizzi).
  - **Estensioni standard (fissate):**
    - I (Integer), E (Registri ridotti), M (Moltiplicazione/Divisione), A (Istruzioni Atomiche), F (Floating Point Singola Precisione), D (Doppia Precisione), C (Compressed).
  - **Estensioni future:**
    - Q (Floating Point Quadrupla Precisione), L (Decimal Floating Point), B (Manipolazione di Bit), V (Operazioni Vettoriali).

---

## ISA RISC-V

- **Concetti principali:**
  - **Instruction Count e CPI:** determinano le prestazioni della CPU, influenzati dall'ISA, dal programma e dal compilatore.
  - **RISC-V vs CISC:**
    - **RISC-V:** architettura con ≈200 istruzioni, operandi nei registri, formato regolare (4 formati base).
    - **CISC (x86):** oltre 1000 istruzioni con lunghezze variabili e modalità di indirizzamento complesse.
  - **Principi di progettazione RISC-V:**
    - Semplicità, regolarità, focus sulla velocità dei casi comuni e supporto a compromessi progettuali.
  - **Istruzioni RISC-V:**
    - Formati principali: R-type (aritmetiche), I-type (immediate), S-type (memoria), SB-type (branch), UJ-type (jump).
  - **Procedure Call Convention:** gestione di parametri, trasferimento di controllo e ritorno da procedure.
  - **Estensioni future:** V (Operazioni vettoriali), B (Manipolazione di bit), L (Decimal Floating Point).

## Ripasso Concetti Base - Pipeline RISC-V

- **Instruction Execution**
  - PC → Instruction memory, fetch instruction
  - Register numbers → Register file, read registers
  - Depending on instruction:
    - ALU calculates:
      - Arithmetic result
      - Memory address (load/store)
      - Branch comparison
  - Access data memory (load/store)
  - PC ← Target address or PC + 4

## Pipeline RISC-V a 5 stadi

1. **IF** (Instruction Fetch)
   - Fetch dell'istruzione da memoria.
   - Incremento PC per puntare alla prossima istruzione.
2. **ID** (Instruction Decode)
   - Decodifica dell'istruzione e lettura dei registri.
   - Test condizionale per branch.
3. **EX** (Execution)
   - Esecuzione dell'operazione ALU o calcolo indirizzo.
4. **MEM** (Memory Access)
   - Accesso alla memoria per load/store.
   - Aggiornamento del Program Counter (PC).
5. **WB** (Write Back)
   - Scrittura del risultato nei registri.

## Concetti Principali

- **Pipeline Speedup**
  - Maggiore throughput (aumento delle istruzioni completate per ciclo).
  - La latenza non diminuisce per singola istruzione.
  - Se tutti gli stadi sono bilanciati:
    - `Tempo istruzione pipelined = Tempo istruzione non pipelined / Numero stadi`.

- **Hazards (Alee)**
  - **Structural Hazards**: risorse richieste da più istruzioni contemporaneamente.
  - **Data Hazards**: dipendenze tra dati di istruzioni consecutive.
  - **Control Hazards**: decisioni sui salti condizionali non ancora determinate.

## Soluzioni alle Alee

- **Structural Hazards**: pipeline con memorie separate per istruzioni e dati.

- **Data Hazards**:
  - **Pipeline Interlock**: stall della pipeline finché i dati non sono pronti.
  - **Forwarding (Bypassing)**: uso dei risultati appena calcolati, senza attendere la scrittura nei registri.
- **Control Hazards**:
  - **Branch Prediction Static**: predizione del salto come "taken" o "untaken".
  - **Branch Prediction Dinamica**: uso di hardware per monitorare la storia dei salti e prevedere il comportamento futuro.

## Prestazioni Pipeline

- **CPI (Cycles per Instruction)** dipende da:
  - Stalli per risolvere le alee.
  - Progettazione della microarchitettura (forwarding, branch resolution).

- **Formula del Tempo di CPU**:
  - `CPU Time = Instruction Count × CPI × Clock Cycle Time`.

## [[2.2_Modello_di_Esecuzione_Cuda_GC]]

- La pipeline migliora il throughput eseguendo istruzioni in parallelo.
- Affetta da alee strutturali, di dati e di controllo.
- La progettazione dell'ISA (Instruction Set Architecture) influisce sulla complessità dell'implementazione della pipeline.

## Ripasso concetti base

- **Pipeline**
- **Forwarding unit**
- **Branch prediction**

## Data Hazards in ALU Instructions

- Esempio di sequenza:
  - `sub  x2, x1,x3`
  - `and  x12,x2,x5`
  - `or   x13,x6,x2`
  - `add  x14,x2,x2`
  - `sd   x15,100(x2)`
- Risoluzione degli hazard con **forwarding**

## Rilevazione della necessità di Forwarding

- Passaggio dei numeri dei registri lungo la pipeline
- Numeri dei registri degli operandi ALU nella fase EX sono:
  - `ID/EX.RegisterRs1`, `ID/EX.RegisterRs2`
- Hazards di dati:
  - 1a. `EX/MEM.RegisterRd == ID/EX.RegisterRs1`
  - 1b. `EX/MEM.RegisterRd == ID/EX.RegisterRs2`
  - 2a. `MEM/WB.RegisterRd == ID/EX.RegisterRs1`
  - 2b. `MEM/WB.RegisterRd == ID/EX.RegisterRs2`
- Forwarding da:
  - Registro pipeline `EX/MEM`
  - Registro pipeline `MEM/WB`

## Condizioni di Forwarding

- Solo se l'istruzione di forwarding scrive su un registro:
  - `EX/MEM.RegWrite`, `MEM/WB.RegWrite`
- Solo se `Rd ≠ x0`:
  - `EX/MEM.RegisterRd ≠ 0`
  - `MEM/WB.RegisterRd ≠ 0`

## Double Data Hazard

- Sequenza di esempio:
  - `add x1,x1,x2`
  - `add x1,x1,x3`
  - `add x1,x1,x4`
- Entrambi gli hazard si verificano
- Utilizzare il risultato più recente:
  - Revisione della condizione di hazard MEM:
    - Forward solo se la condizione di hazard EX non è vera.

## Dynamic Branch Prediction

- Più profonda la pipeline, più significativa la penalità del branch
- **Branch prediction buffer** (branch history table)
  - Indicizzato dagli indirizzi delle istruzioni branch recenti
  - Memorizza il risultato (taken/not taken)
- Fasi di esecuzione di un branch:
  - Controllo della tabella, aspettandosi lo stesso risultato
  - Avvio del fetch dal fall-through o target
  - Se errato, flush della pipeline e inversione della predizione

## 1-Bit Predictor: Problema

- Predizione errata due volte nei branch interni a un loop:
  - Mispredizione come taken all'ultima iterazione del loop interno
  - Mispredizione come not taken alla prima iterazione successiva

## 2-Bit Predictor

- Cambio della predizione solo dopo due mispredizioni successive.

## Calcolo del Branch Target

- Anche con un predittore, è necessario calcolare l'indirizzo target
  - Penalità di 1 ciclo per un branch taken se il calcolo dell'indirizzo non è completato in ID.
- **Branch target buffer**:
  - Cache degli indirizzi target
  - Indicizzato dal PC quando l'istruzione viene fetchata
  - Se c'è un hit e l'istruzione è predetta taken, è possibile fetchare immediatamente il target.

# [[Lezione_3]]

## Intructions Level Parallelism (ILP)

- **Definizione:** esecuzione di più istruzioni in parallelo all'interno di un singolo flusso di istruzioni.

Le tecniche per sfruttare l'ILPsono **piplining**, **out-of-order execution** e **esecuzione speculativa**.

## Data Level Parallelism (DLP)

- **Definizione:** esecuzione di più operazioni su più dati indipendenti in parallelo per migliorare le prestazioni del sistema. Sfuttato molto nelle architetture SIMD e vettoriali.

## Thread Level Parallelism (TLP)

- **Definizione:** esecuzione di più flussi di controllo in parallelo per migliorare le prestazioni del sistema sfruttando a pieno le risorse HW. Sfuttato nelle architetture multicore e multiprocessore.

**Due istruzioni sono INDIPENDENTI se si puo invertire l'ordine delle istruzioni senza modificare il risultato**

## Data Dependance

**Loop-Level parallelism**: parallelismo tra iterazioni di un ciclo.

- si srotola il ciclo staticamente o dinamicamente
- si utilizza SIMD (Single Instruction Multiple Data)

- Istruzione j dipende da i se:
  - i produce un risultato che può essere usato da j
  - se j dipende da k e k dipende da i, allora j dipende da i

**Due istruzioni dipendenti non possono essere eseguite in parallelo**

- Le dipendenze sono proprietarie del codice, non dell'HW
- L'organizzazione della pipeline determina se le ci sono dipendenze e se esse causano stalli

Una dipendenza di dato porta a:

- La possibilità di causare stalli
- Ad un ordine obbligatori in cui le istruzioni devono essere eseguite
- Ad un limite superiore al parallelismo

Le dipendenze che passano attraverso la memoria sono più difficili da gestire.

Se due istruzioni usano lo stesso nome ma non c'è dipendenza di dati, non c'è bisogno di stallare la pipeline.

**Anti-dependence (Write-after-Read)**: una istruzione **i** scrive un registro che un'altra istruzione **j** legge.

**Output-dependence (Write-after-Write)**: due istruzioni scrivono lo stesso registro.

**Entrambe sono false dipendenze**

Per risolvere questi problemi si usa **Register Renaming**, ovvero possiamo riordinare o eseguire parallelamente le istruzioni se il nome viene cambiato per evitare il conflitto.

## Alee di Dato

- **RAW (Read After Write)**: **j** prova a leggere un registro prima che **i** lo scriva leggendo così un valore sbagliato (non aggiornato).

-**WAW (Write After Write)**: **j** scrive un registro prima che **i** lo faccia, sovrascrivendo il valore.
Possibile sono in pipline che permettono la scrittura in piu stadi o che permettono l'esecuzione fuori ordine.
Non da problimi perchè il **WB** della prima istruzione avviene prima del **WB** della seconda istruzione.

- **WAR (Write After Read)**: **j** scrive un registro prima che **i** lo legga, **i** legge un valore sbagliato.
Non puo succedere se la pipeline legge fli operandi in ID, quindi prima del fetch della seconda istruzione. Puo succedere se alcune istruzioni scrivono nei primi stadi della pipeline o se le istruzioni eseugono fuori ordine.