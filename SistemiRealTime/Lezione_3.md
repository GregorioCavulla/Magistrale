# Lezione 3

## RMPO rate monotonic priority ordering

Ad ogniprocesso viene assegnata una priorità fissa, che è inversamente proporzionale al periodo di esecuzione del processo. In questo modo, i processi con periodo di esecuzione minore avranno priorità maggiore. Questo algoritmo è molto semplice e veloce, ma non è ottimale. Infatti, non tiene conto del tempo di esecuzione del processo, ma solo del periodo. Inoltre, non è possibile utilizzare questo algoritmo se i periodi di esecuzione dei processi non sono multipli l'uno dell'altro.

$p(P_j) = p_j \propto \frac{1}{T_j} = \frac{1}{D_j}$

Algoritmo ottimo per i task periodici, ma non per i task sporadici.

Se un insieme di processi periodichi è schedulabile con un qualche algoritmo che prevede una attribuzione statica di priorità, allora è schedulabile anche con RMPO.

Se un insieme di processi periodici non è schedulabile con RMPO,
allora tale insieme non è schedulabile con alcun altro algoritmo
che preveda un’attribuzione statica di priorità.

