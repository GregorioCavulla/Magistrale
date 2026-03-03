# Fondamenti di Computer Graphics (A.A. 2025/26)

## Indice
- [Fondamenti di Computer Graphics (A.A. 2025/26)](#fondamenti-di-computer-graphics-aa-202526)
  - [Indice](#indice)
  - [1. Struttura del Corso e Obiettivi](#1-struttura-del-corso-e-obiettivi)
    - [Modulo 1 (Prof. Casciola) - 6 CFU](#modulo-1-prof-casciola---6-cfu)
    - [Modulo 2 (Prof.ssa Morigi) - 2 CFU](#modulo-2-profssa-morigi---2-cfu)
  - [2. Il Sistema Grafico e la Pipeline 3D](#2-il-sistema-grafico-e-la-pipeline-3d)
    - [Definizioni Fondamentali](#definizioni-fondamentali)
    - [Pipeline Grafica 3D](#pipeline-grafica-3d)
  - [3. Argomenti Tecnici di Studio](#3-argomenti-tecnici-di-studio)
    - [Fondamenti 2D e Matematica](#fondamenti-2d-e-matematica)
    - [WebGL e Programmable Pipeline](#webgl-e-programmable-pipeline)
    - [Rendering e Modellazione Avanzata](#rendering-e-modellazione-avanzata)
  - [4. Software e Strumenti di Sviluppo](#4-software-e-strumenti-di-sviluppo)
  - [5. Modalità d'Esame e Progetto](#5-modalità-desame-e-progetto)
  - [6. Risorse Utili](#6-risorse-utili)
---

## 1. Struttura del Corso e Obiettivi

Il corso è diviso in due moduli complementari che coprono la teoria e la pratica della grafica digitale moderna.

### Modulo 1 (Prof. Casciola) - 6 CFU
* **Focus**: Concetti base della Computer Graphics 3D, grafica interattiva, modellazione poligonale e real-time rendering.
* **Carico di lavoro**: Circa 150 ore totali (48h lezione, 52h studio/pratica, 50h preparazione esame).
* **Tecnologie**: HTML5, CSS3, JavaScript, WebGL e GLSL.

### Modulo 2 (Prof.ssa Morigi) - 2 CFU
* **Focus**: Libreria **Three.js** per lo sviluppo di grafica 3D Real-Time su Web.

---

## 2. Il Sistema Grafico e la Pipeline 3D

Il concetto cardine della disciplina è la trasformazione del dato geometrico in informazione visiva.

### Definizioni Fondamentali
* **Computer Graphics**: Processo $3D \rightarrow 2D$ (da modelli matematici a immagini).
* **Image Processing**: Processo $2D \rightarrow 2D$ (manipolazione di immagini esistenti).
* **Computer Vision**: Processo $2D \rightarrow 3D$ (ricostruzione di scene da immagini).

### Pipeline Grafica 3D
La pipeline è il flusso di lavoro che permette di visualizzare un modello su uno schermo:
1.  **3D Models**: Creazione degli oggetti tramite modellazione interattiva, procedurale, scansione 3D o librerie esterne.
2.  **Rendering**: Calcolo della resa visiva (luci, ombre, materiali).
3.  **2D Image Display**: Output finale dei pixel sul dispositivo.

---

## 3. Argomenti Tecnici di Studio

### Fondamenti 2D e Matematica
* **Trasformazioni Geometriche 2D**: Operazioni di traslazione, rotazione e scala nel piano.
* **Algebra Lineare**: Fondamentale per la manipolazione delle coordinate.

### WebGL e Programmable Pipeline
* **GLSL (OpenGL Shading Language)**: Linguaggio per scrivere Vertex e Fragment Shaders che girano direttamente sulla GPU.
* **Versioni**: WebGL 1.0 (basato su OpenGL ES 2.0) e WebGL 2.0 (basato su OpenGL ES 3.0).
* **Vantaggi**: Nessuna compilazione esterna, supporto mobile nativo (iOS/Android) e integrazione completa in HTML5.

### Rendering e Modellazione Avanzata
* **Pipeline di Vista**: Trasformazioni di vista e proiezioni geometriche (prospettiche o ortografiche).
* **Gestione Visibilità**: Algoritmi di Clipping, Rasterizing e utilizzo dello **Z-Buffer** (Depth Buffer) per gestire le sovrapposizioni.
* **Illuminazione**: Implementazione del modello di **Phong** (ambientale, diffusa, speculare).
* **Texture Mapping**: Tecnica per "incollare" immagini 2D su mesh 3D per aumentarne il realismo.

---

## 4. Software e Strumenti di Sviluppo

Il corso adotta strumenti open-source e standard industriali.

* **Editor**: Visual Studio Code (consigliato).
* **Browser**: Google Chrome.
* **Modellazione 3D**: **Blender 5.0**.
    * *Nota*: È previsto un mini-corso specifico su Blender (7 ore) tra aprile e maggio 2026.

---

## 5. Modalità d'Esame e Progetto

L'esame richiede il passaggio dalla conoscenza alla competenza pratica.

1.  **Realizzazione Progetto**: Sviluppo di un'applicazione grafica basata sulle specifiche fornite nell'A.A. di frequenza.
2.  **Prova Orale**: Discussione sul progetto e domande teoriche su tutto il programma del corso.
3.  **Modulo 2**: Discussione di una rielaborazione/modifica di un codice Three.js.

---

## 6. Risorse Utili

* **Sito Web**: [https://www.dm.unibo.it/~casciola/html/CG2526.html](https://www.dm.unibo.it/~casciola/html/CG2526.html).
    * *Credenziali*: `student` / `onward`.
* **Virtuale**: Per iscrizione, esercizi e comunicazioni.
* **Account LABs**: Necessario per accedere alle macchine dei laboratori d'Ingegneria (`infoy.ing.unibo.it`).
* **Bibliografia**:
    * *Fundamentals of Computer Graphics* (Shirley).
    * *3D Computer Graphics: A Mathematical Introduction with OpenGL* (Buss).