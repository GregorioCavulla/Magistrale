# Introduzione e Applicazioni della Computer Graphics

La Computer Graphics (CG) è la disciplina informatica che si occupa della generazione e manipolazione di immagini e oggetti attraverso il calcolatore. Il suo scopo principale è trasformare modelli matematici e geometrici in rappresentazioni visive comprensibili.

## Indice
- [Introduzione e Applicazioni della Computer Graphics](#introduzione-e-applicazioni-della-computer-graphics)
  - [Indice](#indice)
  - [1. Il Panorama dell'Immagine Digitale](#1-il-panorama-dellimmagine-digitale)
  - [2. Architettura di un Sistema Grafico](#2-architettura-di-un-sistema-grafico)
  - [3. La Pipeline Grafica: Dal Modello al Pixel](#3-la-pipeline-grafica-dal-modello-al-pixel)
  - [4. Ambiti Applicativi](#4-ambiti-applicativi)
    - [Progettazione Industriale (CAD/CAM)](#progettazione-industriale-cadcam)
    - [Medicina e Visualizzazione Scientifica](#medicina-e-visualizzazione-scientifica)
    - [Cartografia e GIS](#cartografia-e-gis)
    - [Intrattenimento e Arte](#intrattenimento-e-arte)
  - [5. Standard e Portabilità](#5-standard-e-portabilità)
---

## 1. Il Panorama dell'Immagine Digitale

Per comprendere la CG, è fondamentale distinguerla dalle discipline correlate che operano nel dominio delle immagini:

* **Computer Graphics**: Parte da una descrizione matematica (modello 3D) per sintetizzare un'immagine 2D. È un processo di "sintesi".
* **Image Processing**: Prende come input un'immagine e restituisce un'immagine modificata. Si occupa di miglioramento, restauro e compressione.
* **Computer Vision**: Il processo inverso della CG; cerca di estrarre una descrizione 3D partendo da immagini 2D reali.
* **Interazione Uomo-Macchina (HCI)**: Studia come la grafica (GUI) possa facilitare la comunicazione tra l'utente e la potenza di calcolo del sistema.

---

## 2. Architettura di un Sistema Grafico

Un sistema moderno non è composto solo dal software, ma da una stretta sinergia hardware:

* **Processore Centrale (CPU)**: Gestisce la logica dell'applicazione e prepara i dati geometrici.
* **Processore Grafico (GPU)**: Un'unità di calcolo massivamente parallela dedicata esclusivamente alla pipeline grafica.
* **Frame Buffer**: Una memoria speciale che contiene il valore cromatico di ogni pixel destinato al monitor. La qualità dell'immagine dipende dalla risoluzione e dalla profondità di bit di questa memoria.

L'innovazione fondamentale è stata il passaggio dalla pipeline a funzioni fisse alla **Programmable Pipeline**, che permette allo sviluppatore di scrivere i propri shader (Vertex e Fragment) in linguaggio **GLSL**.

---

## 3. La Pipeline Grafica: Dal Modello al Pixel

La pipeline è la "catena di montaggio" digitale che trasforma i dati in visione.



1.  **Modellazione**: Gli oggetti vengono definiti nello spazio tramite vertici e primitive (punti, linee, triangoli).
2.  **Trasformazioni**: Gli oggetti vengono posizionati, orientati e proiettati secondo le leggi della prospettiva per simulare una telecamera.
3.  **Rasterizzazione**: I triangoli geometrici vengono convertiti in "frammenti", ovvero i pixel potenziali che comporranno l'immagine.
4.  **Shading**: Viene calcolato il colore finale di ogni pixel applicando modelli di illuminazione (come il modello di Phong) e texture.

---

## 4. Ambiti Applicativi

La Computer Graphics è ormai pervasiva in quasi ogni settore della società moderna:

### Progettazione Industriale (CAD/CAM)
In ingegneria, il **Computer Aided Design (CAD)** permette di progettare componenti complessi e testarli in ambienti virtuali (aerodinamica, stress meccanico) prima della produzione fisica. Il **CAM** collega questi modelli direttamente alle macchine utensili.

### Medicina e Visualizzazione Scientifica
La CG permette di rendere visibili dati complessi. In medicina, i dati di TAC e Risonanze vengono ricostruiti in 3D per diagnosi non invasive e pianificazione chirurgica. Nella scienza, aiuta a visualizzare flussi di fluidi, fenomeni meteorologici o strutture molecolari.

### Cartografia e GIS
I sistemi **GIS (Geographic Information Systems)** utilizzano la grafica per gestire mappe interattive che integrano dati geologici, demografici e ambientali, fondamentali per la gestione del territorio.

### Intrattenimento e Arte
* **Videogiochi**: Spingono l'innovazione nel **Real-Time Rendering**, dove l'immagine deve essere generata in millisecondi in risposta all'utente.
* **Cinema**: Utilizza il **Pre-rendering** (off-line) per ottenere un fotorealismo estremo, dedicando anche ore al calcolo di un singolo fotogramma.



---

## 5. Standard e Portabilità
L'adozione di standard come **WebGL** permette di eseguire queste pipeline complesse direttamente all'interno dei browser web, garantendo che le applicazioni grafiche siano accessibili universalmente su PC, tablet e smartphone senza installazioni aggiuntive.