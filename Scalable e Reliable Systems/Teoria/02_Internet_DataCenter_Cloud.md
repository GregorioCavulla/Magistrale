# Analisi Integrale: Internet Data Center e Cloud Computing

Questo modulo analizza l'intera infrastruttura dei servizi scalabili, dalla fisica dei rack alla logica di business del Cloud.

## Indice
- [Analisi Integrale: Internet Data Center e Cloud Computing](#analisi-integrale-internet-data-center-e-cloud-computing)
  - [Indice](#indice)
  - [1. Architettura Fisica e Ingegneristica degli IDC](#1-architettura-fisica-e-ingegneristica-degli-idc)
    - [Gestione Termica: Il Flusso dell'Aria](#gestione-termica-il-flusso-dellaria)
    - [Efficienza Energetica (PUE)](#efficienza-energetica-pue)
    - [Continuità Elettrica](#continuità-elettrica)
  - [2. Standard Tier (Uptime Institute)](#2-standard-tier-uptime-institute)
  - [3. Strategie di Geo-Distribuzione](#3-strategie-di-geo-distribuzione)
    - [Disaster Recovery (DR) e Business Continuity](#disaster-recovery-dr-e-business-continuity)
    - [Global Server Load Balancing (GSLB)](#global-server-load-balancing-gslb)
  - [4. Cloud Computing: Modelli e Caratteristiche](#4-cloud-computing-modelli-e-caratteristiche)
    - [Modelli di Servizio (La Piramide)](#modelli-di-servizio-la-piramide)
    - [Caratteristiche Essenziali (NIST)](#caratteristiche-essenziali-nist)
  - [5. Modelli di Deployment](#5-modelli-di-deployment)

---

## 1. Architettura Fisica e Ingegneristica degli IDC

Un Internet Data Center (IDC) è un ecosistema complesso dove l'hardware viene protetto da ogni possibile causa di interruzione.

### Gestione Termica: Il Flusso dell'Aria
I server generano una quantità massiccia di calore. Per evitare surriscaldamenti e sprechi energetici, si adotta la configurazione **Hot/Cold Aisle**:
* **Corridoio Freddo (Cold Aisle)**: I rack sono rivolti l'uno verso l'altro. L'aria condizionata viene spinta dal pavimento flottante in questo corridoio.
* **Corridoio Caldo (Hot Aisle)**: Le parti posteriori dei server espellono l'aria calda in questo spazio comune, da dove viene aspirata e raffreddata nuovamente.



### Efficienza Energetica (PUE)
Il **Power Usage Effectiveness (PUE)** è il rapporto tra l'energia totale consumata dal data center e quella effettivamente usata dai server. Più il valore è vicino a 1.0, più il data center è efficiente (ovvero spreca pochissima energia in raffreddamento o illuminazione).

### Continuità Elettrica
Per non fermarsi mai, un IDC usa una catena di sicurezza:
1.  **Rete Elettrica**: Connessioni ridondanti a diverse centrali.
2.  **UPS (Uninterruptible Power Supply)**: Batterie o volani che intervengono istantaneamente se manca la corrente.
3.  **Generatori Diesel**: Motori enormi che si avviano in pochi secondi e possono alimentare il centro per giorni.

---

## 2. Standard Tier (Uptime Institute)

L'affidabilità non è un'opinione, ma una classificazione rigorosa basata sulla ridondanza dei percorsi di distribuzione (elettricità, acqua per raffreddamento, dati).

* **Tier I (Basic)**: Un solo percorso, nessuna ridondanza. Se si rompe un tubo o manca corrente, tutto si ferma.
* **Tier II (Redundant Components)**: Un solo percorso, ma alcuni componenti (es. pompe o condizionatori) sono doppi.
* **Tier III (Concurrently Maintainable)**: Multipli percorsi di distribuzione. Questo è il salto di qualità: puoi smontare e riparare qualsiasi pezzo dell'infrastruttura senza mai spegnere i server.
* **Tier IV (Fault Tolerant)**: Infrastruttura completamente compartimentata. Anche se scoppia un incendio in una sala macchine, l'altra metà del data center continua a lavorare senza che l'utente se ne accorga.

---

## 3. Strategie di Geo-Distribuzione

Un servizio è davvero affidabile solo se sopravvive alla distruzione di un intero data center (es. terremoti o alluvioni).

### Disaster Recovery (DR) e Business Continuity
* **RPO (Recovery Point Objective)**: Indica quanti dati puoi permetterti di perdere (es. "le ultime 2 ore di transazioni").
* **RTO (Recovery Time Objective)**: Indica quanto tempo può stare fermo il servizio prima di tornare online (es. "deve ripartire entro 30 minuti").
* **Replica Geografica**: I siti devono essere distanti tra loro (spesso oltre i 100km) per evitare che lo stesso evento colpisca entrambi.

### Global Server Load Balancing (GSLB)
Il sistema instrada l'utente verso il data center più vicino o meno carico, ottimizzando la latenza e garantendo che, se un sito cade, il traffico venga deviato automaticamente sugli altri.

---

## 4. Cloud Computing: Modelli e Caratteristiche

Il Cloud è l'astrazione totale dell'hardware, offerto come utility (come l'elettricità o l'acqua).

### Modelli di Servizio (La Piramide)
* **IaaS (Infrastructure as a Service)**: Affitti il "ferro" virtuale (CPU, RAM, Disco). Gestisci tu il sistema operativo (es. AWS EC2).
* **PaaS (Platform as a Service)**: Affitti un ambiente dove caricare solo il tuo codice. Il provider gestisce l'OS e il middleware (es. Heroku).
* **SaaS (Software as a Service)**: Usi direttamente l'applicazione via browser. Non gestisci nulla (es. Google Drive).



### Caratteristiche Essenziali (NIST)
* **On-demand Self-service**: Compri e attivi risorse da solo, tramite API o portale.
* **Rapid Elasticity**: Puoi passare da 1 a 1000 server in pochi minuti per gestire un picco di traffico.
* **Resource Pooling**: Molti utenti (tenants) condividono le stesse risorse fisiche in modo sicuro e isolato.
* **Pay-per-use**: Paghi solo i secondi o i minuti in cui le risorse sono state effettivamente attive.

---

## 5. Modelli di Deployment
* **Public Cloud**: Infrastruttura condivisa (multi-tenancy) gestita da giganti come Amazon o Microsoft.
* **Private Cloud**: Infrastruttura dedicata esclusivamente a un'azienda, per esigenze di privacy o controllo estremo.
* **Hybrid Cloud**: Il mix ideale. Usi il private per i dati sensibili e il public per scalare durante i picchi di traffico (Cloud Bursting).