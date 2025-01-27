**1. Che cos'è un'architettura eterogenea? Perchè l'abbiamo esplorata?**

# MODULO 1

## SIMD
**2. Descrivi in generale le caratteristiche del paradigma SIMD**

**3. Quali sono delle tipiche istruzioni supportate da un estensione SIMD?**

**4. Che cos'è una operazione con saturazione e come mai è utile?**

**5. Parlami del branching in SIMD**

## Dealing with real numbers
**6. Descrivi le tecniche principali di codifica dei reali?**

**7. Che tipi di errori sono possibili con floating-point?**

**8. Parlami degli altri formati per i reali oltre a quelli definiti da IEEE? Come mai si è sentito il bisogno di crearne di altri?**

**9. Che altre tecniche conosci per ridurre il memory footprint di un programma?**

**10. Qual'è il caso d'uso di fixed-point?**

**11. Quando è necessario utilizzare floating point?**

# MODULO 2

## Modello di programmazione CUDA
**11. Che cos'è il modello di programmazione CUDA?**

**12. Che cos'è un thread CUDA**

**13. Qual'è il tipico workflow in CUDA?**

**14. Come vengono organizzati i thread in CUDA?**

**15. Come mai c'è bisogno di una gerarchia di thread?**

**16. Che cos'è un kernel CUDA?**

**17. Ci sono dei limiti per quanto riguarda il dimensionamento di blocchi e griglie?**

**18. Che influenza ha il dimensionamento di blocchi e griglie sulle performance?**

**19. Che influenza ha il mapping dei dati ai thread sulle performance?**

**20. Esistono diversi metodi di mapping dei dati?**

## Modello di esecuzione CUDA
**21. Che cos'è il modello di esecuzione?** 

**22. Che cos'è un SM e da che cosa è composto?**

**23. Come vengono distribuiti i blocchi tra i vari SM?**

**24. Parlami di SIMT e delle sue differenze con il modello SIMD**

**25. Parlami di più dei warp ed in particolare del warp scheduling**

**26. Parlami di come si può ottenere il latency hiding massimo**

**27. Che cos'è Indipendent Thread Scheduling?**

**28. Perchè sono necessarie le operazioni atomiche in CUDA?**

**29. Che cos'è il resource partitioning in CUDA?**

**30. Che cos'è l'occupancy in CUDA?**

**31. Parlami di CUDA Dynamic Parallelism**

## Modello di memoria CUDA
**32. Parlami di kernel compute bound e kernel memory bound, come mai è importante distinguere queste due categorie?** 

**33. Come possiamo capire se un kernel è memory bound o compute bound? Che cos'è il diagramma di roofline?** 

**34. Che tipi di memoria esistono in CUDA?**

**35. Come vengono trasferiti i dati dalla memoria dell'host alla memoria del device? A che cosa bisogna stare attenti in questo processo?**

**36. Che cos'è la memoria zero copy?**

**37. Che cosa sono Unified Virtual Addressing (UVA) e Unified Memory (UM)?**

**38. Che cosa si intende con pattern di accesso alla memoria? Come mai è importante avere un pattern ottimo per le performance di un kernel?**

**39. Puoi farmi qualche esempio di utilizzo della SMEM**

**40. Come si utilizza la SMEM?**

**41. Che cos'è un bank conflict?**
