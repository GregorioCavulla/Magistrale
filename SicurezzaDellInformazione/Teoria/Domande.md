Ecco una serie di domande formulate come se provenissero dalla docente del corso, basate sia sui contenuti del tuo `Riassuntone2.0.md` sia sullo stile degli esami precedenti che hai caricato (domande aperte, articolate, con focus su teoria, protocolli e criticità):

---

### 🔐 Sezione 1 – Fondamenti e crittografia

**1.** Descriva le proprietà che deve possedere una funzione hash crittograficamente sicura e illustri l’attacco di estensione della lunghezza (length extension attack), specificando come è possibile prevenirlo.

**2.** Cosa si intende per funzione unidirezionale? Faccia riferimento al ruolo che tali funzioni svolgono nei protocolli di autenticazione, firma e cifratura.

**3.** Illustri le differenze operative e di sicurezza tra PRNG e CSPRNG. Perché nei cifrari a flusso sincroni non è opportuno utilizzare un TRNG?

---

### 🔄 Sezione 2 – Cifrari e modalità operative

**4.** Confronti un cifrario a blocchi in modalità CBC con un cifrario a flusso sincrono, soffermandosi sugli effetti in caso di perdita di sincronismo e di errori di trasmissione.

**5.** Descriva il concetto di malleabilità crittografica e spieghi in che modo si manifesta nei cifrari a flusso.

**6.** Elenchi i vantaggi e svantaggi della modalità ECB rispetto alla modalità CTR, spiegando anche come cambiano le proprietà in caso di attacchi con testo in chiaro noto.

---

### 🔑 Sezione 3 – Gestione chiavi, autenticazione e protocolli

**7.** Confronti il modello di distribuzione delle chiavi basato su KDC e quello basato su infrastrutture PKI. Quali vantaggi e svantaggi presenta ciascuno?

**8.** Descriva un protocollo di autenticazione attiva basato su meccanismo di challenge-response, indicando come può essere vulnerabile a un attacco reflection.

**9.** Supponiamo che un utente A desideri inviare un messaggio riservato, autentico e non ripudiabile a un utente B. Indichi quali trasformazioni applicare al messaggio e con quali chiavi, giustificando le scelte.

---

### 📄 Sezione 4 – Certificati, firme e blockchain

**10.** Descriva il ruolo della CRL e dell’OCSP nella gestione della revoca dei certificati digitali. Quali sono le criticità principali associate all’uso delle CRL?

**11.** Come si costruisce un certificato X.509? Qual è la differenza tra un certificato autofirmato e uno rilasciato da una CA? In quale scenario si possono usare certificati autofirmati?

**12.** Spieghi il funzionamento di un Merkle Tree e il motivo per cui è utilizzato nelle tecnologie blockchain. Come contribuisce alla proprietà di tamper-proof?

---

### 🛡️ Sezione 5 – Analisi, identità e sicurezza applicativa

**13.** Descriva le fasi e i criteri di un’analisi del rischio efficace in ambito aziendale. Quali variabili vanno considerate nella definizione delle contromisure?

**14.** Il sistema di identificazione biometrica rientra tra i metodi basati su possesso, conoscenza o conformità? Quali problemi di sicurezza possono presentarsi in fase di riconoscimento?

**15.** Cosa significa che una blockchain è “tamper-proof”? Quali tecnologie o proprietà crittografiche garantiscono questa caratteristica?

---

**1.** *Spiega il funzionamento del meccanismo di consenso Proof of Stake (PoS) in una blockchain pubblica. Quali vantaggi presenta rispetto a Proof of Work (PoW) in termini di sicurezza e sostenibilità energetica?*

---

**2.** *Confronta le blockchain pubbliche, private e di tipo consortium in termini di accesso, governance, efficienza e sicurezza. In quale contesto aziendale è preferibile adottare Hyperledger Fabric rispetto a Ethereum?*

---

**3.** *Descrivi in dettaglio l’architettura modulare di Hyperledger Fabric, specificando il ruolo dei peer (endorsing e committing), dell’orderer e del chaincode. Come viene raggiunto il consenso e quali sono i vantaggi di un sistema deterministico come Raft?*

---

**4.** *Che differenza c’è tra la blockchain e il WorldState all’interno del ledger di Hyperledger Fabric? Come viene gestito lo stato del sistema e come avviene una lettura/scrittura di un asset tramite chaincode?*

---

**5.** *Cos’è l’identità digitale? Distinguine i livelli (dichiarata, verificata, certificata) e spiega come funzionano i sistemi federati di identità come SAML e OpenID Connect. Quali vantaggi offre il Single Sign-On (SSO) nella gestione aziendale degli accessi?*

---

**6.** *Quali sono i problemi legati all’uso delle password tradizionali? Descrivi il funzionamento dell’autenticazione passwordless con crittografia asimmetrica, e come gli standard FIDO2 (WebAuthn + CTAP) affrontano questi problemi.*

---

**7.** *Nell’ambito dell’autenticazione FIDO2, cosa avviene durante la fase di registrazione e cosa durante la fase di autenticazione? Quali sono i ruoli di Relying Party, Authenticator e WebAuthn API in questo processo?*

---

**8.** *Quali sono i principali limiti delle blockchain in termini di privacy, scalabilità e governance? Fornisci un esempio concreto in cui questi limiti potrebbero compromettere l’adozione della tecnologia.*

---
