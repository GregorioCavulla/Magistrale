# Hardware per la Computer Graphics: Evoluzione e Architetture

Il supporto hardware è l'elemento che ha permesso alla Computer Graphics di passare da semplici linee vettoriali a mondi 3D fotorealistici. Questa sezione analizza come i dispositivi si sono evoluti per gestire carichi di calcolo sempre più massicci.

## Indice
- [Hardware per la Computer Graphics: Evoluzione e Architetture](#hardware-per-la-computer-graphics-evoluzione-e-architetture)
  - [Indice](#indice)
  - [1. Evoluzione dei Dispositivi di Visualizzazione](#1-evoluzione-dei-dispositivi-di-visualizzazione)
  - [2. Il Sistema Grafico Moderno](#2-il-sistema-grafico-moderno)
  - [3. La Rivoluzione della GPU (Graphics Processing Unit)](#3-la-rivoluzione-della-gpu-graphics-processing-unit)
  - [4. Dispositivi di Input e Interazione](#4-dispositivi-di-input-e-interazione)
  - [5. Standard e Astrazione Hardware](#5-standard-e-astrazione-hardware)
---

## 1. Evoluzione dei Dispositivi di Visualizzazione

Storicamente, la visualizzazione è passata attraverso diverse tecnologie chiave:

* **Sistemi Vettoriali (Random Scan)**: Caratteristici degli anni '60 e '70, questi monitor disegnavano le immagini come una serie di segmenti di linea. Non esisteva il concetto di pixel; il raggio catodico si muoveva seguendo direttamente le coordinate dei vettori.
* **Sistemi Raster (Raster Scan)**: È la tecnologia alla base dei moderni monitor (CRT, LCD, Plasma, LED). L'immagine è una matrice di punti chiamati **pixel**. Il raggio (o il segnale digitale) scansiona lo schermo riga per riga (scanline).

---

## 2. Il Sistema Grafico Moderno

Un'architettura hardware dedicata alla grafica si compone di elementi specifici che lavorano in parallelo:

* **Video Controller**: Si occupa di leggere costantemente il contenuto della memoria grafica e inviare i segnali al monitor per il refresh dell'immagine.
* **Frame Buffer (Video Memory)**: Una porzione di RAM dedicata a contenere il valore cromatico di ogni pixel. 
    * La **Risoluzione** è determinata dal numero di pixel memorizzabili (es. 1920x1080).
    * La **Profondità di bit** determina quanti colori diversi può assumere un pixel (es. 24 bit per il True Color, ovvero 8 bit per canale R, G, B).
* **Display Processor (Graphics Accelerator)**: Un'unità che scarica la CPU dal compito di convertire le primitive geometriche (punti, linee) in pixel nel frame buffer.



---

## 3. La Rivoluzione della GPU (Graphics Processing Unit)

La GPU rappresenta il salto evolutivo definitivo. A differenza della CPU, progettata per la logica sequenziale complessa, la GPU è ottimizzata per il **calcolo parallelo**.

* **Parallelismo Massivo**: Una GPU moderna contiene migliaia di core semplificati che eseguono la stessa operazione su dati diversi (SIMD - Single Instruction, Multiple Data). Questo è ideale per la grafica, dove bisogna calcolare il colore di milioni di pixel contemporaneamente.
* **Pipeline Programmabile**: Mentre i primi acceleratori avevano funzioni fisse, le GPU moderne permettono di programmare fasi specifiche (Vertex e Fragment Processing) tramite gli **Shader**.

---

## 4. Dispositivi di Input e Interazione

L'hardware grafico non è solo output. Per la grafica interattiva sono fondamentali:

* **Dispositivi di puntamento**: Mouse, tavolette grafiche e touch screen.
* **Scanner e Digitalizzatori**: Per trasformare oggetti fisici in modelli matematici 3D (scansione laser, fotogrammetria).
* **Dispositivi VR/AR**: Visori e sensori di tracciamento del movimento che richiedono latenze bassissime per evitare il "motion sickness".

---

## 5. Standard e Astrazione Hardware

Per evitare che i programmatori debbano scrivere codice diverso per ogni scheda video (NVIDIA, AMD, Intel), si utilizzano le **API Grafiche**:

* **Livello di Astrazione**: Librerie come **OpenGL** e **WebGL** agiscono come intermediari. Il programmatore invia comandi standard e il driver della scheda video li traduce in istruzioni specifiche per l'hardware.
* **Vantaggi**: Questo garantisce la portabilità del software su diverse piattaforme (Windows, Linux, Mobile) e permette di sfruttare l'accelerazione hardware direttamente dal browser.