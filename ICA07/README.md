# ICA07

# Oppgave 1

## A
Serverprogrammet lytter til port 1234. Om et klientprogram lykkes i å kommunisere med server vil det printes ut en hemmelig melding fra klient. Server vil gi beskjed til klient om at melding er mottat.
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/udpserver.png)

Klientprogrammet lager en datagram-socket på port 1234. En melding skal overføres fra klient til server. En if settning ser til om meldingen overføres uten error, om error==nil, det vil si at det ikke skjedde noe feil,  vil programmet hoppe til neste linje etter medlingen er sent. En error print vil derimot komme frem om overføringen ikke lyktes. 
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/udpclient.png)

## B
For å kjøre de to programmene åpner man to vinduer i terminalen, hver av dem kjører hver sin go fil. udpserver.go startes først, deretter udpklient.go, da vil meldignen automatisk sendes over fra klient til server og server vil gi en tilbakemelding.
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/terminal-run.png)

## C
I wireshark kan man studere all nettverkstrafikk inn og ut fra en node.  Øverst i søkefeltet kan man legge til filter, et udp filter vil kun vise udp trafikk. Localhost henviser alltid til loopback IP-adressen 127.0.0.1 i IPv4, som er vårt tilfelle, eller ::1 om man bruker IPv6

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/wireshark-inspect.png)

### i

#### 1
I menylinjen til wireshark velger man statistics deretter protocoll hierarchy, et vindu vil åpne seg med en oversikt over valgt element. Her ser vi oversikten over den lokale hendelsen mellom 1234 altså klienten vår, og 52735 som var serverprogrammet. Tabellen viser at totalt hundre prosent av data i denne overføring var protokoll-data. Pakker inneholder vanligvis flere protokoller, det vil si at mer enn én protokoll vil bli regnet med for hver pakke. Protokoll-lag kan bestå av pakker som ikke inneholder noen høyere lag, summen av de øverste lagene pakker trenger ikke summere opp til protokoll-pakkene.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/protocol-chart.png)


#### 2
Teoretisk sett kan en UDP pakke være 65,535 bytes (8 byte for "tittelen" og 65,527 bytes med data), men på grunn av IPv4 protokolen som trenger 20 bytes for IP adressen, blir det da plass til 65,507 bytes med data. De resterende 28 bytesene brukes da til UDP tittelen og IP adressen.

### ii
#### 1

#### 2

#### 3


# Oppgave 2

## A
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/TCPserver.png)
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/TCPclient.png)
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/TCPservclient.png)
## B
### i 
TCP garanterer at alle pakker som blir sent over et nettverk havner i riktig rekkefølge hos mottaker. Bekreftelsespakker blir sendt tilbake til sender for å forsikre om at mottaker har fått den riktige informasjonen. Dette er en mindre effektiv overføring enn UDP som ikke tar hensyn til omstokking.

### ii
Teoretisk sett kan en IP pakke være 65535 bytes inkludert tittel. 
I praksis sendes TCP pakker i størrelse av minste MTU, eller maximum transmission unit, denne maksimalen er forskjellig fra hviken teknologi som brukes, for eksempel er ethernet sin MTU 1500 bytes. Større pakker enn MTU kan sendes ved fragmentering.

### iii
Fragmentering er en prosess som deler datagrammer i mindre biter, slik at de kan bli overført uten å overskride MTU grensen.

### iv
Man bruker UDP der fart er viktigst. Det er mest brukt steder hvor et par tapte pakker ikke ødelegger det som blir sendt. Streaming av videoer, VoIP og spill er eksempler på tjenester som utnytter UDP.

TCP brukes på steder hvor man ikke er så opptatt er fart, men mer opptatt av at man får pakkene slik de var når de ble sendt. Noen gode eksempler er web, e-mail (SMTP/IMAP/POP) og overføring av filer (FTP).

# Oppgave 3

## A

## B

## C
