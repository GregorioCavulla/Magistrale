**1. Che cos'è un'architettura eterogenea? Perchè l'abbiamo esplorata?**
Hint: Pensa alla combinazione di GPU e CPU in un sistema e ai vantaggi specifici che ognuna può offrire.

# MODULO 1

## SIMD
**2. Descrivi in generale le caratteristiche del paradigma SIMD**
Hint: Concentrati sull'elaborazione simultanea di più dati usando un'unica istruzione e i benefici in termini di parallelismo a livello di dati.

**3. Quali sono delle tipiche istruzioni supportate da un'estensione SIMD?**
Hint: Elenca operazioni come load/store, operazioni logiche, aritmetiche e di riarrangiamento dati nei registri estesi.

**4. Che cos'è un'operazione con saturazione e come mai è utile?**
Hint: Ricorda l'importanza di evitare overflow e underflow, specialmente in applicazioni come l'elaborazione di immagini.

**5. Parlami del branching in SIMD**
Hint: Rifletti su come le maschere e le operazioni logiche permettono di aggirare le limitazioni del branching in SIMD.

## Dealing with real numbers
**6. Descrivi le tecniche principali di codifica dei reali?**
Hint: Confronta i formati fixed-point e floating-point, con attenzione alle differenze di range dinamico e accuratezza.

**7. Che tipi di errori sono possibili con floating-point?**
Hint: Focalizzati su errori di approssimazione, troncamenti e propagazione dell'errore nelle operazioni.

**8. Parlami degli altri formati per i reali oltre a quelli definiti da IEEE? Come mai si è sentito il bisogno di crearne di altri?**
Hint: Considera applicazioni specifiche come l'AI, dove il range dinamico è prioritario rispetto all'accuratezza.

**9. Che altre tecniche conosci per ridurre il memory footprint di un programma?**
Hint: Pensa al weight sharing e alla quantizzazione per ridurre l'occupazione di memoria.

**10. Qual'è il caso d'uso di fixed-point?**
Hint: Rifletti su come questa rappresentazione consenta calcoli efficienti in termini di costo e potenza, specialmente su microcontrollori.

**11. Quando è necessario utilizzare floating-point?**
Hint: Considera situazioni che richiedono un elevato range dinamico.

# MODULO 2

## Modello di programmazione CUDA
**11. Che cos'è il modello di programmazione CUDA?**
Hint: Spiega la gerarchia di thread, memoria e l'importanza delle API per ottimizzare l'esecuzione parallela.

**12. Che cos'è un thread CUDA**
Hint: Focalizzati sull'unità di elaborazione elementare e sul parallelismo attraverso thread multipli.

**13. Qual'è il tipico workflow in CUDA?**
Hint: Descrivi il processo dall'allocazione della memoria all'elaborazione e recupero dei dati.

**14. Come vengono organizzati i thread in CUDA?**
Hint: Spiega la gerarchia di thread, blocchi e griglie, e come possono essere dimensionati.

**15. Come mai c'è bisogno di una gerarchia di thread?**
Hint: Sottolinea come ciò aiuti nella scalabilità e portabilità del codice CUDA su hardware diverso.

**16. Che cos'è un kernel CUDA?**
Hint: Pensa a una funzione che viene eseguita parallelamente da molti thread.

**17. Ci sono dei limiti per quanto riguarda il dimensionamento di blocchi e griglie?**
Hint: Rifletti sui vincoli dati dalla compute capability della GPU.

**18. Che influenza ha il dimensionamento di blocchi e griglie sulle performance?**
Hint: Considera l'importanza di bilanciare i carichi di lavoro e massimizzare l'occupancy.

**19. Che influenza ha il mapping dei dati ai thread sulle performance?**
Hint: Sottolinea l'importanza di accessi alla memoria coalescenti e del bilanciamento del carico.

**20. Esistono diversi metodi di mapping dei dati?**
Hint: Confronta i metodi lineare e per coordinate, spiegando quando è meglio usare ciascuno.

## Modello di esecuzione CUDA
**21. Che cos'è il modello di esecuzione?**
Hint: Spiega come il modello di programmazione si traduce in esecuzione effettiva sulla GPU.

**22. Che cos'è un SM e da che cosa è composto?**
Hint: Descrivi le unità di calcolo, registri e memoria condivisa all'interno di uno SM.

**23. Come vengono distribuiti i blocchi tra i vari SM?**
Hint: Spiega il ruolo del GigaThread Engine e come avviene il load balancing.

**24. Parlami di SIMT e delle sue differenze con il modello SIMD**
Hint: Confronta l'abilità di supportare divergenza tra thread con la rigidità del SIMD.

**25. Parlami di più dei warp ed in particolare del warp scheduling**
Hint: Spiega come il warp scheduler massimizza il parallelismo e nasconde la latenza.

**26. Parlami di come si può ottenere il latency hiding massimo**
Hint: Rifletti sull'importanza di avere warp pronti e sull'uso della Legge di Little.

**27. Che cos'è Indipendent Thread Scheduling?**
Hint: Spiega come ogni thread ottiene il proprio stato di esecuzione per ridurre i deadlock e migliorare il throughput.

**28. Perché sono necessarie le operazioni atomiche in CUDA?**
Hint: Considera il problema delle corse critiche e l'importanza della coerenza nei risultati.

**29. Che cos'è il resource partitioning in CUDA?**
Hint: Descrivi come registri, memoria condivisa e thread siano suddivisi tra blocchi per massimizzare l'occupancy.

**30. Che cos'è l'occupancy in CUDA?**
Hint: Spiega come una maggiore occupancy aiuti nel latency hiding.

**31. Parlami di CUDA Dynamic Parallelism**
Hint: Pensa ai vantaggi di creare kernel dinamicamente dalla GPU senza intervento della CPU.

## Modello di memoria CUDA
**32. Parlami di kernel compute bound e kernel memory bound, come mai è importante distinguere queste due categorie?**
Hint: Rifletti su come i colli di bottiglia nella memoria o nei calcoli influenzano le strategie di ottimizzazione.

**33. Come possiamo capire se un kernel è memory bound o compute bound? Che cos'è il diagramma di roofline?**
Hint: Utilizza l'intensità aritmetica e la soglia per identificare il tipo di kernel e spiega come il diagramma visualizzi queste informazioni.

**34. Che tipi di memoria esistono in CUDA?**
Hint: Fai una panoramica dei vari tipi di memoria e delle loro caratteristiche principali (es. registri, memoria globale, memoria condivisa).

**35. Come vengono trasferiti i dati dalla memoria dell'host alla memoria del device? A che cosa bisogna stare attenti in questo processo?**
Hint: Descrivi il ruolo del bus PCIe e l'importanza di minimizzare i trasferimenti multipli.

**36. Che cos'è la memoria zero copy?**
Hint: Spiega come la memoria pinned consente accessi diretti tra host e device senza copia esplicita.

**37. Che cosa sono Unified Virtual Addressing (UVA) e Unified Memory (UM)?**
Hint: Confronta UVA e UM, evidenziando come UM automatizzi i trasferimenti di memoria.

**38. Che cosa si intende con pattern di accesso alla memoria? Come mai è importante avere un pattern ottimo per le performance di un kernel?**
Hint: Rifletti sull'importanza di accessi coalescenti e allineati per minimizzare le transazioni di memoria.

**39. Puoi farmi qualche esempio di utilizzo della SMEM**
Hint: Pensa a come la memoria condivisa possa migliorare i pattern di accesso e ridurre le sincronizzazioni.

**40. Come si utilizza la SMEM?**
Hint: Descrivi il processo dal caricamento dei dati alla scrittura dei risultati nella memoria globale.

**41. Che cos'è un bank conflict?**
Hint: Spiega come accessi a banchi di memoria condivisa possano essere sequenzializzati se non gestiti correttamente.

