# Dispensa di Sicurezza Informatica – Ingegneria e Sistemi Informativi

## Indice
- [Dispensa di Sicurezza Informatica – Ingegneria e Sistemi Informativi](#dispensa-di-sicurezza-informatica--ingegneria-e-sistemi-informativi)
  - [Indice](#indice)
  - [1. Principi Fondamentali della Sicurezza](#1-principi-fondamentali-della-sicurezza)
- [Dispensa di Sicurezza Informatica – Ingegneria e Sistemi Informativi](#dispensa-di-sicurezza-informatica--ingegneria-e-sistemi-informativi-1)
  - [Indice](#indice-1)
  - [1. Principi Fondamentali della Sicurezza](#1-principi-fondamentali-della-sicurezza-1)
  - [2. Analisi del Rischio e Contromisure](#2-analisi-del-rischio-e-contromisure)
    - [ALE (Annual Loss Expectancy)](#ale-annual-loss-expectancy)
    - [Tipologie di contromisure](#tipologie-di-contromisure)
  - [3. Crittografia](#3-crittografia)
    - [3.1 Concetti fondamentali](#31-concetti-fondamentali)
    - [3.2 Funzioni Hash](#32-funzioni-hash)
    - [3.3 Cifrari a blocchi e a flusso](#33-cifrari-a-blocchi-e-a-flusso)
  - [4. Protocollo Kerberos](#4-protocollo-kerberos)
  - [5. Infrastrutture a Chiave Pubblica (PKI)](#5-infrastrutture-a-chiave-pubblica-pki)
  - [6. Firme digitali e autenticazione](#6-firme-digitali-e-autenticazione)
  - [7. IPSec e protocolli di rete sicuri](#7-ipsec-e-protocolli-di-rete-sicuri)
  - [8. Blockchain e Hyperledger](#8-blockchain-e-hyperledger)
  - [9. Identità Digitale e FIDO2](#9-identità-digitale-e-fido2)
  - [10. Argomenti aggiuntivi](#10-argomenti-aggiuntivi)

---

## 1. Principi Fondamentali della Sicurezza
# Dispensa di Sicurezza Informatica – Ingegneria e Sistemi Informativi

## Indice

- [Dispensa di Sicurezza Informatica – Ingegneria e Sistemi Informativi](#dispensa-di-sicurezza-informatica--ingegneria-e-sistemi-informativi)
  - [Indice](#indice)
  - [1. Principi Fondamentali della Sicurezza](#1-principi-fondamentali-della-sicurezza)
- [Dispensa di Sicurezza Informatica – Ingegneria e Sistemi Informativi](#dispensa-di-sicurezza-informatica--ingegneria-e-sistemi-informativi-1)
  - [Indice](#indice-1)
  - [1. Principi Fondamentali della Sicurezza](#1-principi-fondamentali-della-sicurezza-1)
  - [2. Analisi del Rischio e Contromisure](#2-analisi-del-rischio-e-contromisure)
    - [ALE (Annual Loss Expectancy)](#ale-annual-loss-expectancy)
    - [Tipologie di contromisure](#tipologie-di-contromisure)
  - [3. Crittografia](#3-crittografia)
    - [3.1 Concetti fondamentali](#31-concetti-fondamentali)
    - [3.2 Funzioni Hash](#32-funzioni-hash)
    - [3.3 Cifrari a blocchi e a flusso](#33-cifrari-a-blocchi-e-a-flusso)
  - [4. Protocollo Kerberos](#4-protocollo-kerberos)
  - [5. Infrastrutture a Chiave Pubblica (PKI)](#5-infrastrutture-a-chiave-pubblica-pki)
  - [6. Firme digitali e autenticazione](#6-firme-digitali-e-autenticazione)
  - [7. IPSec e protocolli di rete sicuri](#7-ipsec-e-protocolli-di-rete-sicuri)
  - [8. Blockchain e Hyperledger](#8-blockchain-e-hyperledger)
  - [9. Identità Digitale e FIDO2](#9-identità-digitale-e-fido2)
  - [10. Argomenti aggiuntivi](#10-argomenti-aggiuntivi)

---

## 1. Principi Fondamentali della Sicurezza

La sicurezza informatica si basa su un insieme di proprietà fondamentali che, nel loro insieme, garantiscono la protezione dei dati e dei sistemi informatici. Questi principi sono:

* **Riservatezza (Confidentiality)**: consiste nell’assicurare che solo gli utenti autorizzati possano accedere alle informazioni. È spesso realizzata tramite tecniche di cifratura e controlli di accesso. Ad esempio, quando inviamo un messaggio crittografato, stiamo garantendo la riservatezza del contenuto.

* **Integrità (Integrity)**: garantisce che le informazioni non vengano modificate accidentalmente o intenzionalmente da soggetti non autorizzati. Tecniche comuni includono l’uso di checksum, firme digitali e codici MAC.

* **Disponibilità (Availability)**: assicura che i sistemi e le informazioni siano accessibili agli utenti autorizzati quando necessario. Un attacco DoS (Denial of Service), ad esempio, compromette questa proprietà.

* **Autenticazione (Authentication)**: permette di verificare l’identità di un utente o di un sistema. Serve a evitare che attori malevoli si spaccino per entità legittime.

* **Non Ripudio (Non-repudiation)**: garantisce che un soggetto non possa negare di aver effettuato una certa azione, come l’invio di un messaggio. Le firme digitali offrono questo tipo di garanzia.

Questi principi, noti anche con l'acronimo CIA (Confidentiality, Integrity, Availability), sono il fulcro della progettazione dei sistemi sicuri.

## 2. Analisi del Rischio e Contromisure

L’analisi del rischio è una metodologia fondamentale per la gestione della sicurezza. Consiste nella valutazione sistematica delle minacce, vulnerabilità e potenziali impatti sugli asset aziendali.

### ALE (Annual Loss Expectancy)

La formula `ALE = Asset × Probability × Exposure × Impact` consente di calcolare l’aspettativa annuale di perdita economica dovuta a un certo rischio.

Ad esempio:

* Se un attacco ha una probabilità del 60% di verificarsi due volte all’anno, con un impatto stimato di 125€ per evento su un asset del valore di 300€, l’ALE sarà: `300 × 0,6 × 2 × 125 = 45.000€`.
* Se implementiamo contromisure che riducono la probabilità residua (ad es. al 40%), l’ALE scende a `9.000€`, consentendoci di valutare se vale la pena di investirvi.

### Tipologie di contromisure

* **Prevenzione**: agisce per impedire l’attacco. Include firewall, crittografia, policy di sicurezza.
* **Rilevazione**: consente di identificare un attacco in corso o appena avvenuto. Esempi: IDS (Intrusion Detection System), log e audit.
* **Reazione**: azioni che mitigano o bloccano l’attacco, come spegnere un sistema o isolare una rete.
* **Deviazione**: si usano honeypot o altri sistemi falsi per attirare l’attaccante e studiarne il comportamento.

L’obiettivo è ridurre il rischio a un livello accettabile, bilanciando i costi delle contromisure con il valore degli asset protetti.

## 3. Crittografia

### 3.1 Concetti fondamentali
La crittografia è la disciplina che studia i metodi per proteggere l'informazione, rendendola inaccessibile a soggetti non autorizzati. È uno dei pilastri della sicurezza informatica e si divide in due grandi categorie:

- **Crittografia simmetrica**: lo stesso segreto (chiave) è usato sia per cifrare che per decifrare. Esempi: AES, DES.
- **Crittografia asimmetrica**: si usano due chiavi differenti ma correlate: una pubblica per cifrare, una privata per decifrare. Esempi: RSA, ECC.

Obiettivi principali della crittografia:
- Riservatezza
- Integrità
- Autenticazione
- Non ripudio

### 3.2 Funzioni Hash
Le funzioni hash crittografiche trasformano un messaggio di lunghezza arbitraria in una stringa di lunghezza fissa, detta digest. Proprietà fondamentali:
- Efficienza
- Unidirezionalità
- Resistenza alle collisioni (debole e forte)

Sono utilizzate per verificare l'integrità dei messaggi, nelle firme digitali, nei protocolli di autenticazione.

### 3.3 Cifrari a blocchi e a flusso
- **Cifrari a blocchi**: operano su blocchi di dati di lunghezza fissa. Modalità di funzionamento: ECB, CBC, CFB, OFB, CTR.
- **Cifrari a flusso**: operano su bit o byte singoli. Possono essere sincroni o autosincronizzanti. Utilizzano un generatore pseudocasuale (PRNG sicuro).

---

## 4. Protocollo Kerberos
Kerberos è un protocollo di autenticazione di rete basato su chiavi simmetriche e su un'autorità centrale fidata (KDC). Permette l’autenticazione mutua tra client e server e supporta la Single Sign-On (SSO). Utilizza ticket cifrati (Ticket Granting Ticket e Service Ticket) per autorizzare accessi a servizi.

---

## 5. Infrastrutture a Chiave Pubblica (PKI)
Una PKI (Public Key Infrastructure) gestisce la generazione, distribuzione e revoca dei certificati digitali, i quali associano chiavi pubbliche a identità verificate.

Componenti:
- Certificate Authority (CA)
- Registration Authority (RA)
- Repository di certificati e revoche (CRL/OCSP)

Ogni certificato è firmato dalla CA con la sua chiave privata e può essere verificato da chiunque abbia la chiave pubblica della CA.

---

## 6. Firme digitali e autenticazione
- **Firme digitali**: garantiscono autenticità, integrità e non ripudio. Basate sulla crittografia asimmetrica. Esempio: `Sign(H(m))`.
- **MAC e HMAC**: codici di autenticazione basati su funzioni hash e chiavi segrete. Garantiscono autenticità e integrità ma non non ripudio.

Modelli di cifratura autenticata:
- Encrypt-then-MAC (raccomandato)
- MAC-then-Encrypt
- Encrypt-and-MAC

---

## 7. IPSec e protocolli di rete sicuri
IPSec è un insieme di protocolli che opera a livello di rete (IP) e fornisce confidenzialità, autenticazione e integrità dei pacchetti IP.

Componenti:
- AH (Authentication Header)
- ESP (Encapsulating Security Payload)
- IKE (Internet Key Exchange)

Modalità operative:
- Transport Mode (host-to-host)
- Tunnel Mode (gateway-to-gateway, VPN)

---

## 8. Blockchain e Hyperledger
La blockchain è un registro distribuito, immutabile e condiviso. Ogni blocco contiene:
- un hash del blocco precedente
- transazioni
- timestamp
- nonce

Meccanismi di consenso: PoW, PoS, PBFT, Raft.

**Hyperledger Fabric**: piattaforma enterprise permissioned con canali privati, smart contract (chaincode) e consenso modulare.

---

## 9. Identità Digitale e FIDO2
- **Identità Digitale**: insieme di attributi che identificano univocamente un utente. Può essere dichiarata, verificata o certificata.
- **FIDO2 e Passwordless**: standard di autenticazione basati su chiavi pubbliche, che eliminano la necessità di password.

Componenti:
- WebAuthn (browser)
- CTAP (protocollo)
- Autenticatore (token, smartphone)

---

## 10. Argomenti aggiuntivi
- Analisi degli attacchi (spoofing, sniffing, DoS, buffer overflow)
- Challenge-Response e nonce
- Replay attack, reflection attack
- Sicurezza nei sistemi distribuiti
- Tecniche di difesa attiva e passiva
- Honeypot e deception technologies
