# HTML5 Canvas e Geometria delle Trasformazioni 2D

Questo blocco analizza come il browser gestisce la grafica programmabile e come la matematica delle matrici permetta di manipolare oggetti nello spazio bidimensionale.

## Indice
- [HTML5 Canvas e Geometria delle Trasformazioni 2D](#html5-canvas-e-geometria-delle-trasformazioni-2d)
  - [Indice](#indice)
  - [1. Architettura dell'elemento Canvas HTML5](#1-architettura-dellelemento-canvas-html5)
  - [2. Matematica delle Trasformazioni Geometriche](#2-matematica-delle-trasformazioni-geometriche)
    - [Coordinate Omogenee](#coordinate-omogenee)
    - [Trasformazioni Fondamentali](#trasformazioni-fondamentali)
  - [3. Lo Stack di Stato: Save e Restore](#3-lo-stack-di-stato-save-e-restore)
  - [4. Animazione e Interattività](#4-animazione-e-interattività)
  - [5. Composizione delle Trasformazioni (L'ordine conta)](#5-composizione-delle-trasformazioni-lordine-conta)

---

## 1. Architettura dell'elemento Canvas HTML5
L'elemento `<canvas>` è una delle innovazioni principali di HTML5, fornendo una superficie di disegno bitmap controllabile interamente tramite JavaScript.

* **Il Rendering Context**: Il canvas è solo un contenitore vuoto. Per operare, si deve richiamare il "contesto" (solitamente `2d`). Questo oggetto contiene lo stato del disegno (colori, spessori, trasformazioni) e i metodi per tracciare primitive.
* **Immediate Mode Rendering**: Il sistema non mantiene un grafo della scena. Quando disegni una linea, il computer cambia il colore dei pixel interessati nel buffer e "dimentica" l'entità linea. Per muovere un oggetto, devi cancellare l'area e ridisegnarlo.
* **Sistema di Coordinate**: L'origine (0,0) si trova nell'angolo in alto a sinistra. L'asse X cresce verso destra, mentre l'asse Y cresce verso il basso. Questa convenzione è opposta al piano cartesiano standard e influenza la direzione delle rotazioni.

---

## 2. Matematica delle Trasformazioni Geometriche
Per manipolare la geometria (muovere, ruotare, scalare), applichiamo trasformazioni ai vertici degli oggetti. Per rendere questi calcoli efficienti, si utilizza l'algebra matriciale.

### Coordinate Omogenee
Un punto nel piano $(x, y)$ viene rappresentato come una terna $(x, y, 1)$. Questa estensione a una terza dimensione fittizia è fondamentale perché permette di rappresentare la **Traslazione** (che sarebbe una somma) come una moltiplicazione tra matrici, proprio come la rotazione e la scala.



### Trasformazioni Fondamentali
1.  **Traslazione**: Sposta i punti di un offset $(dx, dy)$. In una matrice 3x3, questi valori occupano l'ultima colonna.
2.  **Scalatura**: Moltiplica le coordinate per fattori di scala. Se i fattori sono diversi da 1, l'oggetto cambia dimensione. Se sono negativi, l'oggetto viene specchiato rispetto all'asse.
3.  **Rotazione**: Ruota i punti di un angolo $\theta$ attorno all'origine. Nel canvas, un angolo positivo genera una rotazione oraria a causa dell'orientamento dell'asse Y.

---

## 3. Lo Stack di Stato: Save e Restore
Le trasformazioni nel Canvas sono cumulative. Se esegui due rotazioni da 10 gradi, l'oggetto successivo sarà ruotato di 20. Per gestire scene complesse senza fare calcoli inversi infiniti, si usa lo **State Stack**:

* `ctx.save()`: Salva la matrice di trasformazione corrente, il colore di riempimento e lo stile del tratto in una pila.
* `ctx.restore()`: Recupera l'ultimo stato salvato. È una pratica standard "salvare" prima di trasformare un singolo oggetto e "ripristinare" subito dopo averlo disegnato, per non influenzare il resto della scena.

---

## 4. Animazione e Interattività
Un'animazione è una successione di frame statici visualizzati rapidamente (tipicamente 60 volte al secondo).



1.  **Cancellazione**: Si pulisce il frame precedente con `clearRect()`.
2.  **Aggiornamento**: Si calcolano le nuove posizioni (es. la nuova coordinata X di un proiettile).
3.  **Disegno**: Si applicano le trasformazioni e si renderizzano le primitive.
4.  **Loop**: Si usa `requestAnimationFrame()`, una funzione ottimizzata che sincronizza il ridisegno con il refresh rate del monitor, evitando "tearing" visivi.

---

## 5. Composizione delle Trasformazioni (L'ordine conta)
Le trasformazioni non sono commutative. Moltiplicare la matrice di rotazione per quella di traslazione dà un risultato diverso rispetto all'ordine inverso.
* **Rotazione attorno a un punto**: Per ruotare un quadrato attorno al suo centro (e non all'angolo del canvas), la sequenza corretta è: Traslare l'origine sul centro dell'oggetto $\rightarrow$ Ruotare $\rightarrow$ Traslare l'oggetto indietro della metà della sua dimensione.