# Cloud device 

Questa libreria ha lo scopo di centralizzare la *creazione* dei device
che vengono gestiti dal cloud. 
Permette di definire un `device` e le sue informazioni e la codifica/decodifica di alto 
livello dei messaggi.

Il `device` quindi è un oggetto che ad alto livello prepara e decodifica i messaggi per essere
utilizzati in pagine web o log. Non ha il compito di serializzare i messaggi per la trasmissione
via mqtt con il broker.

## Device
Un device è un oggetto **unico** per tipologia (id) che sa quali sono i registri da leggere e scrivere
nel dispositivo fisico.

La libreria permette di *creare* un device indicandone:
- `id` Codice univoco che identificata la tipologia di device
- `info` Informazioni generiche come: nome, version, firmware e seriale
- `registry` Registri di lettura e scrittura 

## Dipendenze
La libreria importa `cloud-protocol` solo per avere le interfacce da ritornare.
La serializzazione dei messaggi viene lasciata a `cloud-protocol` e al servizio che ne
ha necessità.