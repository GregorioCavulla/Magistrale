# Dipendenze e architetture superscalari

## instruction level parallelism (ILP)

l'obiettivo è di massimizzare il CPI
- pipline CPI = ideal pipeline CPI + structural stalls + data hazard stalls + control stalls

## Data Dependency

Loop level parallelism
- srotolo il ciclo e faccio più operazioni in parallelo
- uso SIMD (Single Instruction Multiple Data)

Challenges:
- Data dependency:
    - instruction j depends on instruction i if:
        - i produces a result that j uses
        j is data dependent on k and k is data dependent on i

Istruzioni dipendenti non possono essere eseguite in parallelo

