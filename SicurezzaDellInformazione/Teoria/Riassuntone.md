# Riassuntone

# Indice
- [Riassuntone](#riassuntone)
- [Indice](#indice)
- [Cap.1 - Fondamenti](#cap1---fondamenti)
  - [Funzioni Hash Crittograficamente sicure](#funzioni-hash-crittograficamente-sicure)
    - [Proprietà](#proprietà)
    - [Attacco con estensione (length extension attack)](#attacco-con-estensione-length-extension-attack)
  - [TRNG, PRNG e PRNG sicuri](#trng-prng-e-prng-sicuri)
    - [TRNG (True Random Number Generator)](#trng-true-random-number-generator)
    - [PRNG (Pseudo-Random Number Generator)](#prng-pseudo-random-number-generator)
    - [PRNG Sicuri (CSPRNG - Cryptographically Secure Pseudo-Random Number Generator)](#prng-sicuri-csprng---cryptographically-secure-pseudo-random-number-generator)
    - [Utilizzo nei cifrari a flusso](#utilizzo-nei-cifrari-a-flusso)
- [Cap.2 - Cifrari a flusso e a blocchi](#cap2---cifrari-a-flusso-e-a-blocchi)
  - [Cifrari a blocchi](#cifrari-a-blocchi)
    - [Funzionamento](#funzionamento)
  - [Cifrari a flusso](#cifrari-a-flusso)
    - [Funzionamento](#funzionamento-1)
  - [Differenza tra Cifrari a blocchi e a flusso](#differenza-tra-cifrari-a-blocchi-e-a-flusso)
  - [Modalità di cifratura a blocchi](#modalità-di-cifratura-a-blocchi)
    - [ECB (Electronic Codebook)](#ecb-electronic-codebook)
    - [CBC (Cipher Block Chaining)](#cbc-cipher-block-chaining)
    - [CFB (Chiper Feedback)](#cfb-chiper-feedback)
    - [CTR (Counter](#ctr-counter)
  - [Sincronismo nei cifrari a flusso](#sincronismo-nei-cifrari-a-flusso)
    - [Cifrari a flusso sincroni](#cifrari-a-flusso-sincroni)
    - [Cifrari a flusso autosincronizzanti](#cifrari-a-flusso-autosincronizzanti)
  - [Malleabilità nei cifrari a flusso](#malleabilità-nei-cifrari-a-flusso)
  - [Perdita di sincronismo nei cifrari a flusso](#perdita-di-sincronismo-nei-cifrari-a-flusso)
- [Cap. 3 - Distribuzione chiavi e protocolli](#cap-3---distribuzione-chiavi-e-protocolli)
  - [Distribuzione chiavi](#distribuzione-chiavi)
    - [Modelli con accordo preliminare:](#modelli-con-accordo-preliminare)
      - [Modello a master key](#modello-a-master-key)
      - [Modello KDC (Key Distribution Center)](#modello-kdc-key-distribution-center)
    - [Modelli senza accordo preliminare:](#modelli-senza-accordo-preliminare)
      - [Diffie-Hellman](#diffie-hellman)
      - [Cifratura asimmetrica (es. RSA)](#cifratura-asimmetrica-es-rsa)
- [Cap. 4 - Autenticazione e Identificazione](#cap-4---autenticazione-e-identificazione)
  - [Tipi di autenticazione](#tipi-di-autenticazione)
    - [Autenticazione passiva](#autenticazione-passiva)
    - [Autenticazione attiva](#autenticazione-attiva)
    - [Challenge response](#challenge-response)
  - [Attacchi ai protocolli di autenticazione](#attacchi-ai-protocolli-di-autenticazione)
    - [Reply attack](#reply-attack)
    - [Interleaving attack](#interleaving-attack)
    - [Reflection attack](#reflection-attack)
- [Cap. 5 - Integrità, autenticazione e firme digitali](#cap-5---integrità-autenticazione-e-firme-digitali)
  - [Integrità e autenticazione di messaggi](#integrità-e-autenticazione-di-messaggi)
    - [MAC (Message Authentication Code)](#mac-message-authentication-code)
    - [HMAC (Hash-based Message Authentication Code)](#hmac-hash-based-message-authentication-code)
  - [Modelli di cifratura autenticata (Authenticated Encryption)](#modelli-di-cifratura-autenticata-authenticated-encryption)
    - [tre modelli principali:](#tre-modelli-principali)
  - [Firme digitali](#firme-digitali)
    - [Funzionamento](#funzionamento-2)
  - [Confronto](#confronto)
- [Cap. 6 - Infrastrutture a chiave pubblica (PKI) e gestione dei certificati](#cap-6---infrastrutture-a-chiave-pubblica-pki-e-gestione-dei-certificati)
  - [come fidarsi della chiave pubblica di qualcuno?](#come-fidarsi-della-chiave-pubblica-di-qualcuno)
  - [PKI (Public Key Infrastructure)](#pki-public-key-infrastructure)
  - [Autorità di Certificazione (CA)](#autorità-di-certificazione-ca)
  - [Revoca dei certificati e gestione CRL](#revoca-dei-certificati-e-gestione-crl)
    - [CRL (Certificate Revocation List)](#crl-certificate-revocation-list)
    - [CRL partizionate](#crl-partizionate)
    - [OCSP (Online Certificate Status Protocol)](#ocsp-online-certificate-status-protocol)
  - [Firma e verifica di un certificato](#firma-e-verifica-di-un-certificato)
  - [Certificati autofirmati](#certificati-autofirmati)
- [Cap. 7 - Blockchain](#cap-7---blockchain)
  - [Proprietà](#proprietà-1)
  - [Tipi di blockchain](#tipi-di-blockchain)
  - [Meccanismi di consenso](#meccanismi-di-consenso)
  - [Limiti](#limiti)
- [Cap. 8 - Hyperledger Fabric](#cap-8---hyperledger-fabric)
  - [Architettura](#architettura)
  - [Transazioni e consenso](#transazioni-e-consenso)
  - [Sviluppo e utilizzo](#sviluppo-e-utilizzo)
- [Cap. 9 - Identità digitale](#cap-9---identità-digitale)
  - [Livelli di identità digitale](#livelli-di-identità-digitale)
  - [Gestione dell'identità digitale](#gestione-dellidentità-digitale)
  - [Tecnologie e standard](#tecnologie-e-standard)
- [Cap. 10 - Passwordless Authentication e FIDO2](#cap-10---passwordless-authentication-e-fido2)
  - [Problemi dei metodi basati su password](#problemi-dei-metodi-basati-su-password)
  - [Autenticazione senza password (Passwordless Authentication)](#autenticazione-senza-password-passwordless-authentication)
  - [FIDO2](#fido2)
  - [Funzionamento](#funzionamento-3)
  - [Vantaggi](#vantaggi)
  - [Svantaggi](#svantaggi)


# Cap.1 - Fondamenti

## Funzioni Hash Crittograficamente sicure

Una funzione hash crittografica ```H(x)``` è una funzione che, a partire da un imput id lunghezza arbitraria, genera un output di lunghezza fissa **digest**. È uno degli strumenti più importanti della crittografia moderna.

### Proprietà

1. **Efficienza**:
   - Il calcolo di ```H(x)``` deve essere computazionalmente facile per ogni ```x```.C
2. **Unidirezionalità**:
   - Dato ```H(x)```, è computazionalmente difficile trovare ```x```.
3. **Resistenza alle collisioni**:
   - **debole**: Dato un ```x``` è difficile trovare un ```y ≠ x``` tale che ```H(x) = H(y)```.
   - **forte**: È difficile trovare una coppia ```(x,y)``` tale che ```H(x) = H(y)```.

Una lunghezza di output di 160bit è considerata sicura per quanto riguarda la resistenza debole, 320bit per la resistenza forte. Lo standard moderno è SHA-512 (Secure Hash Algorithm 512 bit).

### Attacco con estensione (length extension attack)

Questo tipo di attacco sfrutta la modalità con cui sono implementate la maggior parte delle funzioni hash, ovvero un meccanismo di **compressione iterata**.

**Esempio**: Se si calcola ```H(s||m)``` per un messaggio, un attaccante può intercettare il digest, costruire un ```H(s||m||m')``` senza conoscere ```s``` e ```m```, utilizzando il digest originale come lo stato di partenza della funzione hash.

**Contromisure**:
- Usare ```H(m||s)``` invece di ```H(s||m)```.
- Applicare un padding contenente la lunghezza del messaggio originale prima di calcolare l'hash.
- Utilizzare **HMAC** (Hash-based Message Authentication Code), che combina una funzione hash con una chiave segreta per garantire l'integrità e l'autenticità del messaggio, progettato per resistere agli attacchi di estensione.

## TRNG, PRNG e PRNG sicuri

### TRNG (True Random Number Generator)

- Si basa su fenomeni fisici imprevedibili e analogici.
- **Pro**: Genera numeri casuali veramente imprevedibili, non deterministici.
- **Contro**: Lentezza nella generazione, difficoltà di implementazione, necessità di hardware specifico.

### PRNG (Pseudo-Random Number Generator)
- **Deterministico**: Genera sequenze di numeri che sembrano casuali ma sono determinate da un seme iniziale (**seed**).
- **Pro**: Velocità, facilità di implementazione, portabilità.
- **Contro**: Non sono veramente casuali, la sequenza può essere riprodotta se si conosce il seme.
  
### PRNG Sicuri (CSPRNG - Cryptographically Secure Pseudo-Random Number Generator)
È sicuro se:
- Non è possibile prevedere i numeri futuri della sequenza anche se si conoscono (**test next bit**).
- È computazionalmente difficile risalire al seme anche conoscendo una parte della sequenza.
- Utilizza almeno una funzione unidirezionale nella transizione di stato o output.

* Un periodo desiderabile per un RNG è di almeno ```10^50``` per garantire sicurezza a lungo termine.
* Il seed dovrebbe essere generato da un TRNG o da una fonte di entropia sufficientemente casuale.

### Utilizzo nei cifrari a flusso

I cifrari a flusso (stream ciphers) generano un flusso di chiave bit a bit da combinare con li testo in chiaro tramite operazioni bitwise (XOR).

**Non si utilizzano i TRNG** per i cifrari a flusso perché:
- I TRNG sono lenti e non adatti per generare grandi quantità di dati in
- Nei cifrari a flusso è necessario che mittente e destinatario generino lo stesso flusso di chiavi, impossibile con un TRNG.

Si utilizzano PRNG sicuri per generare il flusso di chiavi garantendo che:
- Non sia possibile prevedere i bit futuri della chiave.
- Non sia possibile risalire al seme della chiave.

# Cap.2 - Cifrari a flusso e a blocchi

## Cifrari a blocchi

I cifrari a blocchi operano su blocchi di dati di dimensione fissa (es. 64 o 128 bit) e utilizzano una chiave segreta per cifrare e decifrare i dati.

### Funzionamento

Un cifrario a blocchi prende in input un blocco di dati e una chiave, e produce un blocco cifrato. Il processo può essere descritto come segue:
1. **Input**: Un blocco di dati di dimensione fissa (es.
2. **Chiave**: Una chiave segreta di lunghezza fissa.
3. **Cifratura**: Il blocco di dati viene trasformato in un bloc
4. **Output**: Un blocco cifrato di dimensione fissa.
5. **Decifratura**: Il blocco cifrato può essere trasformato nuovamente nel blocco di dati originale utilizzando la stessa chiave.


## Cifrari a flusso

I cifrari a flusso operano su dati in modo continuo, cifrando bit o byte singoli alla volta. Utilizzano un flusso di chiavi generato da un PRNG sicuro.

I cifrari a flusso sono particolarmente adatti per applicazioni in tempo reale come la trasmissione di audio e video, dove è necessario cifrare i dati in modo rapido e continuo.

### Funzionamento

Un cifrario a flusso funziona come segue:
1. **Input**: Dati in chiaro di dimensione variabile.
2. **Chiave**: Una chiave segreta di lunghezza fissa.
3. **Flusso di chiavi**: Un PRNG sicuro genera un flusso di chiavi bit a bit.
4. **Cifratura**: I dati in chiaro vengono combinati con il flusso di chiavi tramite operazioni bitwise (tipicamente XOR).
5. **Output**: Dati cifrati di dimensione variabile.
6. **Decifratura**: I dati cifrati possono essere decifrati combinando nuovamente con lo stesso flusso di chiavi.


## Differenza tra Cifrari a blocchi e a flusso

| Aspetto                 | Cifratura a blocchi                             | Cifratura a flusso                     |
| ----------------------- | ----------------------------------------------- | -------------------------------------- |
| **Unità di operazione** | Blocchi fissi (es. 64 o 128 bit)                | Bit o byte singoli                     |
| **Modello teorico**     | Teoria di Shannon: sostituzione + trasposizione | One Time Pad (OTP)                     |
| **Performance**         | Maggiore overhead, elaborazione a blocchi       | Più veloce, utile in tempo reale       |
| **Error handling**      | Possibili propagazioni                          | Errori spesso localizzati              |
| **Uso tipico**          | File, email, dati statici (store & forward)     | Streaming audio/video, reti, telefonia |


## Modalità di cifratura a blocchi

### ECB (Electronic Codebook)
- Ogni blocco di testo in chiaro è cifrato sepratamente con la stessa chiave
- **Pro**: supporta cifratura parallela, semplice da implementare.
- **Contro**: è **deterministico**: lo stesso blocco di testo in chiaro produce sempre lo stesso blocco cifrato, rendendo possibile l'analisi statistica.

### CBC (Cipher Block Chaining)
- Ogni blocco in chiaro è messo a XOR con il cifrato precedente poi cifrato.
- Richiede un vettore di inizializzazione (IV) casuale e unico per il primo blocco.
- **Pro**: risolve la ripetizione di blocchi.
- **contro**: non supporta cifratura parallela, c'è propagazione dell'errore (se un blocco è corrotto, anche il successivo lo sarà).

### CFB (Chiper Feedback)
- Struttura simile ai cifrari a flusso autosincronizzanti.
- Modifica/cancellazione di bit, errore limitato al transitorio dello shift register.

### CTR (Counter
- Il flusso di chiave è generato cifrando un contatore incrementale.
- **Pro**: completamente parallelizzabile, non ha propagazione dell'errore, non dipende dal messaggio.

## Sincronismo nei cifrari a flusso

### Cifrari a flusso sincroni

- Il flusso di chiave è indipendente dal testo cifrato.
- Un errore di cancellazione//inserzione porta a una perdita di sincronismo **permanente**.

### Cifrari a flusso autosincronizzanti
- Il flusso di chiave dipende dal testo cifrato precedente (es. tramite shift register).
- Perdita di sincronismo **temporanea**, si risolve dopo ```n``` bit (```n``` lunghezza del registro di shift).

## Malleabilità nei cifrari a flusso

Il problema della **malleabilità** si manifesta quando un intrusore può modificare il cifrato in modo da ottenere un chiaro con significato controllato.

**Esempio**:
- sia ```c = m ⊕ k```.
- un attaccante, senza conoscere ```k```, può fare: ```c' = c ⊕ p``` → ```m' = m ⊕ p```. Quindi trasformare ```m``` in ```m'``` scegliendo opportunamente ```p```.
- possibile solo se l'attaccante conosce il formato del messaggio in chiaro.

## Perdita di sincronismo nei cifrari a flusso

Un cifrario a flusso perde il sincronismo se:
- Viene iniettato o cancellato un bit nel canale di trasmissione.
- L'intera sequenza di chiavi generata non è allineata.

# Cap. 3 - Distribuzione chiavi e protocolli

## Distribuzione chiavi

In crittografia simmetrica, due entità che vogliono comunicare devono condvidere una chiave segreta. Il problema è come distribuire questa chiave in modo sicuro quando:
- le entità non si sono mai scambiate chiavi prima.
- il numero delle entità nella rete è elevato.

### Modelli con accordo preliminare:

#### Modello a master key

- ogni coppia di utenti condivide una chiave segreta personale (```K_AB```).
- Le chiavi di sessione vengono poi scambiate cifrandole con la master key.
- Servono ```n(n-1)/2``` chiavi per ```n``` utenti.

#### Modello KDC (Key Distribution Center)

- Ogni utente condivide una sola chiave segreta con il KDC.
- Quando A vuole comunicare con B:
  - A chiede al KDC una chiave di sessione ```k```
  - il KDC genera ```k```, la invia cifrata sia a A che a B utilizzando ```K_A``` e ```K_B```.
- Si hanno solo ```n``` chiavi segrete, ma il KDC diventa un punto critico di fallimento.

### Modelli senza accordo preliminare:

#### Diffie-Hellman

- Protocollo per lo sscambio di chiavi sicuro senza necessità di un canale sicuro iniziale.
- Si basa su operazioni matematiche difficili della teoria dei numeri, come il logaritmo discreto.
- Permettono a due entità di calcolare una chiave comune scambiandosi pubblicamente dei valori.
- Nonn serve che le due parti si siano mai incontrate prima.

#### Cifratura asimmetrica (es. RSA)
- Una parte invia una chiave di sessione cifrandola con la chiave pubblica dell'altra parte.
- La parte ricevente decifra la chiave di sessione con la propria chiave privata.
- Serve una **Public Key Infrastructure (PKI)** per gestire le chiavi pubbliche e private e garantire che siano autentiche.

| Modello              | Scalabilità | Sicurezza                               | Esempi di uso                      |
| -------------------- | ----------- | --------------------------------------- | ---------------------------------- |
| Master Key           | Bassa       | Dipende da ogni coppia                  | Reti piccole, dispositivi dedicati |
| KDC                  | Media       | Dipende dal server centrale             | Reti aziendali, ambienti chiusi    |
| Diffie-Hellman / RSA | Alta        | Sicurezza basata su problemi matematici | Reti aperte, HTTPS, e-mail         |

# Cap. 4 - Autenticazione e Identificazione

**Identificazione**: Dichiarazione dell'identità di un soggetto.
**Autenticazione**: Verifica dell'identità dichiarata da un soggetto.

## Tipi di autenticazione

### Autenticazione passiva

- Per verificare l'identità del mittente di un messaggio.
- Esempi: MAC (Message Authentication Code), HMAC (Hash-based Message Authentication Code), firme digitali.
- Garantisce integrità e autenticità del messaggio.
- Non impedisce attacchi di replay.

### Autenticazione attiva

- Coinvolge l'interazione dinamica tra le parti, ad esempio tramite scambio di nonce o challenge-response.
- Può prevenire attacchi di replay e impersonificazione.

### Challenge response

Struttura generale:
1. Il verificatore (server) invia una sfida (challenge) al soggetto da autenticare (client), ad esempio un nonce casuale.
2. L'utente risponde firmando o cifrando la sfida a seconda del metodo.
3. Il verificatore controlla la risposta.

| Metodo              | Operazione        | Rischi                                      |          |                                                     |
| ------------------- | ----------------- | ------------------------------------------- | -------- | --------------------------------------------------- |
| Hash                | `H(nonce\|\|segreto)`  | se la funzione hash non è resistente all’estensione |
| Cifrario simmetrico | `E_k(nonce)`     | attacco reflection                          |          |                                                     |
| Firme digitali      | `Sign_k−1(nonce)` | sicuro, se la chiave privata è ben protetta |          |                                                     |

Devono garantire freschezza (uso di nonce o timestamp) per prevenire attacchi di replay.

## Attacchi ai protocolli di autenticazione

### Reply attack

- L'attaccante intercetta e memorizza un messaggio di autenticazione valido e lo ripete in un momento successivo per impersonare l'utente.

- **Contromisure**:
  - Utilizzare nonce o timestamp per garantire freschezza.
  - Implementare challenge-response per verificare l'autenticità della richiesta.
  - Utilizzare sessioni temporanee che scadono dopo un certo periodo.
  - Applicare HMAC o firme digitali per garantire l'integrità e l'autenticità del messaggio.

### Interleaving attack

- Due sessioni attive vengono fatte comunicare tra loro da un atttaccante che media e rimescola i messaggi.

### Reflection attack

- Valido nei protocolli simmetrici dove la sfida viene semplicemente riflessa
- Se A invia una sfida a B, B può rispedirla indietro cifrata con la stessa chiave.

# Cap. 5 - Integrità, autenticazione e firme digitali

## Integrità e autenticazione di messaggi

Garantire l'integrità di un messaggio significa assicurarsi che non sia stato alterato durante la trasmissione. Quando a questo si aggiunge la prova dell'identità del mittente, si parla di autenticazione del messaggio.

Due strumenti principali:
- **MAC (Message Authentication Code)**: Un codice generato da una funzione hash combinata con una chiave segreta. È utilizzato per verificare l'integrità e l'autenticità del messaggio.
- **HMAC (Hash-based Message Authentication Code)**: Una variante del MAC che utilizza una funzione hash crittografica. È più sicuro rispetto al MAC tradizionale.

### MAC (Message Authentication Code)
Funzione che prende in input un messaggio ```m``` e una chiave ```k```, produce un codice ```t=MAC_k(m)```

Serve per verificare che il messaggio:
- sia integro
- provenga da chi lo ha firmato (conosce la chiave segreta ```k```).

Non fornisce non ripudiabilità, chiunque conosca la chiave può generare un MAC valido.

### HMAC (Hash-based Message Authentication Code)
Viariante sicura costruita su una funzione hash crittografica. Combina la chiave segreta con il messaggio e applica la funzione hash.

Struttura:
```HMAC_K(m) = H((k ⊕ opad) || H((k ⊕ ipad) || m))``` con ```opad``` e ```ipad``` costanti di padding.

È sicuro anche se la funzione hash non è resistente all'estensione, poiché la chiave segreta ```k``` viene utilizzata in modo tale da impedire attacchi di estensione.
Ampiamente utilizzato in protocolli come TLS, IPsec e SSH.

Resiste agli attacchi di estensione e non rivela il contenuto della chiave.

## Modelli di cifratura autenticata (Authenticated Encryption)

La cifratura autenticata fornisce **confidenzialità e integrità** in una unica operazione.

### tre modelli principali:

| Modello              | Ordine delle operazioni                                        | Vantaggi                | Svantaggi                                        |
| -------------------- | -------------------------------------------------------------- | ----------------------- | ------------------------------------------------ |
| **MAC-then-Encrypt** | Prima si genera il MAC, poi si cifra tutto (messaggio + MAC)   | usato in SSL/TLS vecchi | vulnerabile a attacchi tipo BEAST                |
| **Encrypt-then-MAC** | Prima si cifra il messaggio, poi si calcola il MAC sul cifrato | ✅ più sicuro            | oggi preferito in IPsec, GCM                     |
| **Encrypt-and-MAC**  | Si cifra il messaggio e si calcola il MAC separatamente        | facile da implementare  | rischi di inconsistenza se gestiti separatamente |

## Firme digitali

Le firme digitali sono l'equivalente crittografico della firma autografa.
- Garantisce identità del mittente
- Protegge integrità del contenuto
- Grarantisce non ripudiabilità

### Funzionamento

1. Il mittente genera l'has del messaggio da firmare ```H(m)```.
2. Firma l'hash con la propria chiave privata ```s = Sign_k−1(H(m))```.
3. Invia ```(m, s)``` al destinatario.
4. Il destinatorio verifica la firma con la chiave pubblica del mittente.

La firma [ univoca, solo chi possiede la chiave privata può generarla.]

## Confronto

| Caratteristica    | MAC/HMAC                             | Firma digitale                           |
| ----------------- | ------------------------------------ | ---------------------------------------- |
| Chiave condivisa  | Sì (segreta)                         | No (pubblica/privata)                    |
| Integrità         | ✅                                    | ✅                                        |
| Autenticazione    | ✅                                    | ✅                                        |
| Non ripudiabilità | ❌                                    | ✅                                        |
| Uso tipico        | Comunicazioni simmetriche (TLS, API) | Comunicazioni aperte (e-mail, contratti) |

# Cap. 6 - Infrastrutture a chiave pubblica (PKI) e gestione dei certificati

## come fidarsi della chiave pubblica di qualcuno?

In un sistema assimetrico ogni utente possiede una chiave pubblica, nota a tutti e una chiave privata segreta.

Chi garantiesce che quella chiave pubblica appartenga effettivamente a quella persona?

La risposta è l'**Infrastruttura a Chiave Pubblica (PKI)**.

## PKI (Public Key Infrastructure)

La PKI è un insieme di entità, regole  e strumenti per creare, distribuire, verificare, revocare e gestire i certificati digitali.

Il certificato digitale associa una chiave pubblica a una identità.

**un certificato contiene**:
- Nome del soggetto
- Chiave pubblica
- Dati di validità temporale
- Nome dell'autorità certificatrice (CA)
- Firma digitale della CA sul tutto

## Autorità di Certificazione (CA)

Una Certification Authority (CA) è un'entità fidata che emette certificati digitali. La CA verifica l'identità del richiedente e firma il certificato con la propria chiave privata.

- La CA ha una coppia di chiavi: una chiave pubblica e una chiave privata.
- Frima i certificati con la chiave privata, permettendo a chiunque di verificare la firma utilizzando la chiave pubblica della CA.
- I clienti conoscono la chiave pubblica della CA come fidata e preinstallata
  

Gerarchia della fiducia:
- Certificato del sisto firmato da CA intermedia firmato da CA root.
- Il client verifica la catena di certificati fino a una root conosciuta e fidata.

## Revoca dei certificati e gestione CRL

Un certificato può diventare non valido prima della scadenza. Per questo ci sono 2 sistemi per gestire la revoca dei certificati:

### CRL (Certificate Revocation List)
- Lista di certificati revocati, firmata dalla CA.
- il client deve scaricarla e vrificarla periodicamente.
- Può diventare obsoleta se non aggiornata e molto lunga.
### CRL partizionate
- Le CRL vengono suddivise per intervalli
- Permette aggiornamenti più rapidi ed efficienti.
### OCSP (Online Certificate Status Protocol)
- Protocollo che consente a un cliente di chiedere in tempo reale lo stato di un certificato
- Più leggero e aggiornato rispetto alle CRL.
- Introduce una dipendenza online dalla CA.

## Firma e verifica di un certificato
Per verificare un certificato il client esegue:
1. Controllo validità temporale (non scaduto).
2. Controllo firma della CA (catena fino a root).
3. Controllo che il nome nel certificato corrisponda all'identità.
4. Controllo su CRL o OCSP.

## Certificati autofirmati
- Un certificato può essere autofirmato, (per uso interno o test).
- In questo caso non esiste una CA fidata, quindi i client lo rifiutano.

# Cap. 7 - Blockchain

Una blockchain è una struttura dati distrubiuta, immutabile e verificabile, composta da una catena di blocchi collegati tra loro, che memorizzano transazioni validate in ordine cronologiico. Ogni blocco contiene:
- Un hash del blocco precedente
- Un insieme di transazioni
- Un timestamp
- Un nonce (nei sistemi basati su proof-of-work)

## Proprietà

- **decentralizzazione**: nessuna autorità centrale, ogni nodo ha una copia del ledger (registro).
- **immutabilità**: una volta validata, l'informazione non può essere modificata senza il consenso della rete.
- **trasparenza**: tutti i nodi hanno accesso alla cronologia completa
- **sicurezza**: garantita tramite crittografia (hash, firme digitali, meccanismi di consenso).

## Tipi di blockchain

- **pubbliche**: chiunque può partecipare, come Bitcoin ed Ethereum.
- **private**: accesso ristretto a un gruppo selezionato, come Hyperledger.
- **consortium**: un gruppo di organizzazioni controlla la rete, come R3 Corda.

## Meccanismi di consenso
Servono a far si che tutti i nodi concordino su una unica versione della catena.
- **Proof of Work (PoW)**: i nodi competono per risolvere un problema computazionale, il primo che lo risolve aggiunge il blocco. (Bitcoin).
- **Proof of Stake (PoS)**: i nodi vengono selezionati per aggiungere blocchi in base alla quantità di criptovaluta che possiedono e sono disposti a "bloccare" come garanzia. (Ethereum 2.0).
- **PBFT (Practical Byzantine Fault Tolerance)**: i nodi raggiungono un consenso attraverso un processo di voto, tollerando fino a un terzo di nodi malevoli. (Hyperledger Fabric).
- **Raft e Kafka**: utilizzati in sistemi distribuiti per garantire coerenza e consenso tra nodi.

## Limiti

- **scalabilità**: le blockchain pubbliche possono avere problemi di scalabilità a causa del numero elevato di transazioni e della necessità di consenso.
- **consumo energetico**: i meccanismi come PoW richiedono un elevato consumo energetico per la risoluzione dei problemi computazionali.
- **privacy**: le blockchain pubbliche sono trasparenti, il che può essere un problema per la privacy degli utenti.
- **governance**: la mancanza di un'autorità centrale può rendere difficile prendere decisioni rapide e risolvere conflitti.

# Cap. 8 - Hyperledger Fabric

Hyperledger Fabric è una piattaforma open-source per la creazione di blockchain permissioned, sviluppata da IBM.
Si distingue per:
- architettura modulare
- Supporto per contratti intelligenti (smart contracts)
- Bassa latenza
- Controllo granulare degli accessi (channels)
- Modello di consenso flessibile e pluggable (es. Raft, Kafka).
  
## Architettura

I componenti principali di Hyperledger Fabric sono:
- Client: applicazione che invia richieste alla blockchain.
- Peer: nodo che mantiene una copia del ledger e gestisce le transazioni.
- Orderer: ordina le transazioni, le raggruppa in blocchi e le distribuisce ai peer.
- Certificate Authority (CA): gestisce l'identità degli utenti e dei nodi, emettendo certificati digitali.
- Channel: sottorete privata per transazioni riservate tra un gruppo di nodi.
- Ledger: composto da blockchain (cronologia immutabile) e WorldState (stato coorente)
- Chaincode: implementazione dei contratti intelligenti

Esistono due tipi di peer:
- **Endorsing Peer**: simula le transazioni e fornisce apprivazioni (endorsement) per le transazioni.
- **Committing Peer**: mantiene il ledger e applica le transazioni approvate.

## Transazioni e consenso
1. Il client invia una proposta di transazione agli endorsing peer.
2. Gli endorsing peer simulano la transazione e restituiscono l'output.
3. Il client raccoglie gli endorsment richiesti e li invia all'orderer.
4. L'orderer ordina le transazioni, crea un blocco e lo invia ai peer.
5. I peer verificano e registrano il blocco nel ledger.

La politica di endorment stabilisce quali peer devono approvare una transazione. Il consenso è deterministico: ogni blocco validato è finale. Il protocollo **Raft** è il servizio di ordinamento consigliato.

## Sviluppo e utilizzo

- La rete viene avviata con ```./network.sh up createChannel -ca```
- Chaincode esemplare: **Asset Transfer** (creazione, lettura, aggiornamento, eliminazione, trasferimento).
- Le funzioni ```submitTransaction``` e ```evaluateTransacrion``` permettono di scrivere e leggere dallo stato
- WorldState accessibile tramite ```ctx.stub.getState()``` e ```putState()```.

# Cap. 9 - Identità digitale

L'identità digitale è l'insieme di **dati e attributi digitali** che rappresentano univocamente una persona, organizzazione, dispositivo all'interno di un sistema informatico o rete.

Può includere:
- Nome, data di nascita, codice fiscale
- Email e numero di telefono
- Certificati digitali, credenziali biometriche, chiavi crittografiche

## Livelli di identità digitale

- **Identità dichiarata**: fornita dall'utente.
- **Identità verificata**: validata da un ente terzo.
- **Identità certificata**: rilascata da una trusted authority con valore legale.
  
## Gestione dell'identità digitale

- **IAM (Identity and Access Management)**: sistemi per gestire identità digitali, accessi e autorizzazioni.
- **IdP (Identity Provider)**: forniscono servizi di autenticazione e autorizzazione.
- **Federated Identity**: consente l'uso di identità digitali su più sistemi senza dover creare account separati.
- **Single Sign-On (SSO)**: permette di accedere a più applicazioni con un'unica autenticazione.

## Tecnologie e standard

- **SAML (Security Assertion Markup Language)**: protocollo per lo scambio di dati di autenticazione e autorizzazione tra IdP e SP (Service Provider).
- **OAuth 2.0 e OpenID Connect**: standard per l'autorizzazione e l'autenticazione basati su token e identità su web.
- **DID (Decentralized Identifier)**: standard emergente per identificatori decentralizzati che non dipendono da un'autorità centrale.

# Cap. 10 - Passwordless Authentication e FIDO2

## Problemi dei metodi basati su password

- L'utente medio ha 25 account ma usa password simili o uguali.
- Password deboli e riutilizzate espongono a: 
  - attacchi brute force
  - credential stuffing (riutilizzo di credenziali trafugate) 
  - Phishing (furto di credenziali tramite inganno).
- Le aziende devono proteggere le password secondo il GDPR.

## Autenticazione senza password (Passwordless Authentication)

Metodo di autenticazione che elimina completamente l'uso delle password e dei segreti condivisi. Utilizza:
- Dispositivi registrati (es. smartphone, token hardware)
- Biometria (PIN, impronta digitale, riconoscimento facciale)
- Crittografia asimmetrica per generazione e verifica credenzialo

**Esempio:**
1. L'utente inserisce il proprio identificatore pubblico
2. il sistema invia una sfida crittografica
3. il dispositivo risponde firmando la sfida con una chiave privata locale
4. il server verifica con la chiave pubblica registrata

## FIDO2

FIDO Alliance (2012) ha sviluppato:
- **FIDO UAF, FIDO U2F**: standard per l'autenticazione senza password.
- **FIDO2/WebAuthn**: standard per l'autenticazione basata su chiavi pubbliche, supportato dai principali browser e piattaforme.

**Componenti principali:**
- **Relying Party (RP)**: il servizio che richiede l'autenticazione.
- **Authenticator**: il dispositivo che genera e verifica le credenziali.
- **WebAuthn API**: interfaccia per l'integrazione nei browser e nelle applicazioni web.
- **CTAP (Client To Authenticator Protocol)**: protocollo di comunicazione tra il browser e l'autenticatore.

## Funzionamento

**Registrazione (Registration Ceremony):**
1. Il client invia una sfida, l'identificatore RP, e preferenze
2. L'autenticatore genera una coppia di chiavi pubblica/privata
3. Il server registra la chiave pubblica e l'ID dell'autenticatore (AAGUID) associati all'utente.

**Autenticazione (Authentication Ceremony):**
1. Il client invia una nuova sfida.
2. L'autenticatore usa la chiave privata per firmarla
3. Il server verifica la firma usando la chiave pubblica salvata.

## Vantaggi

- **Sicurezza**: elimina il rischio di furto di password, phishing e credential stuffing.
- **Reistenza agli attacchi**: le chiavi private non lasciano il dispositivo, riducendo il rischio di compromissione.
- **Usabilità**: l'utente non deve ricordare password complesse, può
- **Conformita GDPR**.

## Svantaggi

- Costo di implementazione iniziale.
- Formazione
- Dipendenza da dispositivi specifici (smartphone, token hardware).