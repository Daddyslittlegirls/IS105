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

<!-- kilde: http://stackoverflow.com/questions/26028700/write-to-client-udp-socket-in-go -->

### i

#### 1
I menylinjen til wireshark velger man statistics deretter protocoll hierarchy, et vindu vil åpne seg med en oversikt over valgt element. Her ser vi oversikten over den lokale hendelsen mellom 1234 altså klienten vår, og 52735 som var serverprogrammet. Tabellen viser at totalt hundre prosent av data i denne overføring var protokoll-data. Pakker inneholder vanligvis flere protokoller, det vil si at mer enn én protokoll vil bli regnet med for hver pakke. Protokoll-lag kan bestå av pakker som ikke inneholder noen høyere lag, summen av de øverste lagene pakker trenger ikke summere opp til protokoll-pakkene.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/protocol-chart.png)


#### 2
Teoretisk sett kan en UDP pakke være 65,535 bytes (8 byte for "tittelen" og 65,527 bytes med data), men på grunn av IPv4 protokolen som trenger 20 bytes for IP adressen, blir det da plass til 65,507 bytes med data. De resterende 28 bytesene brukes da til UDP tittelen og IP adressen.

### ii
#### 1

Over for eksempel WiFi er prosentandelen protokoll data hundre prosent totalt. Som nevnt tidligere kan pakker inneholde flere protokoller, i tillegg trenger ikke sub-protokollene summere opp til de øverste protokoll-lagene. For eksempel ser vi TCP ligger på 97,9% mens sub-protokollene som XMPP, SSL og HTTP kun summerer opp til 22,1%. Årsaken kan være TCP overheading, data som ikke bidrar til innhold av meldingen.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/wireshark-wifi-protocol.png)

#### 2
Det finnes to måter å filtere data på i wireshark; display filter og capture filter. 
Ved å sette et display filter velger man hvilke capture files en vil se, for at display filteret skal fungere må wireshark ha fanget opp trafikk på capture filteret. Et capture filter brukes til å velge hvilke pakker som skal bli lagret på disken mens trafikk fanges opp. Capture filter bruker en BPF syntax, en modul som kjører i kjernen og kan derfor opprettholde høyt capture rate fordi pakkene ikke trenger å bevege seg fra rommet i kjernen til rommet til brukeren under filtering. Hva en kan filtere er definert og begrenset i forhold til et display filter.

For å effektivt finne frem til relevante meldinger kan man bruke filter som for eksempel http, tcp eller udp. Da vil kun de valgte meldingene komme opp, wireshark har også ferdig definerte filter man kan velge mellom. Om man vil være mer spesifikk, la oss si du kun vil følge med på http GET info, kan man skrive inn filteret: http.request.method == "GET". På samme måte som i programmeringsspråk kan man bruke tegnene || til å velge og/eller. Filteret: http.request.method == "GET" || (http.request.method == "POST") vil vise all http GET og POST trafikk, om det finnes noen.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/wireshark-filter.png)

#### 3
Tidligere i oppgaven analyserte vi loopback capture, data sendt innenfor vår lokale node. Her så vi de samme kildene og destinasjonene gå igjen, som 127.0.0.1 som er localhost og 192.168.1 som er ruteren. Det er lettere å finne ut hvor data kommer fra og hvor det ender opp hen på det lokale nettverket, det er som regel en aktivitet som er kjent for brukeren.
Når vi fanger opp kommunikasjon fra nettverkskortet ser vi et stort mangfold av IP adresser og protokoller, dette kan være trafikk fra egen node eller andre brukere av det samme nettverket. Vi ser tydelig at det er flere måter å kommunisere på via NIC, et større pool av protokoller. På bildet under ser man en encrypted handshake message, dette er en kobling mellom server og klient ved TLS v1.0 protokoll. Fremgangsmåten består av flere sekvenser; client hello, server hello, certificate, server hello done, client key exchange og change Cipher Spec. Det er altså flere ledd i komunikasjonen, det er viktig at informasjonen som overføres skal havne hos riktig mottaker og forstås på korrekt måte. 

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/wireshark-wifi.png)

# Oppgave 2

## A
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/TCPserver.png)

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/TCPclient.png)

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/TCPservclient.png)

#### Kilde: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

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

# Oppgave 4

## A

For å studere trafikken på nrk direkte tv må vi først trykke play, derretter gå inn på statistikk og velge conversations på wireshark. Vi huker av for TCP kommunikasjon og sorterer etter flest pakker overført, da vil vi få opp dette resultatet: 

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/nrklivetv.png)

Her streamer jeg serien SKAM på nrk nett tv. Legg merke til at resultatet er ganske likt, dette være seg at kvaliteten på streamen er lik for direkte tv og andre filer på bibloteket.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/nrknett.png)

Til slutt, etter å ha lukket begge fanene sitter jeg igjen med github og fronter åpent. Her ser kommunikasjonen litt annerledens ut.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/ingenstream.png)

## B

Når vi analyserer en stream fra youtube er det lett å legge merke til kvaliteten som blir spilt av. Her spilte jeg av Apples superbowl reklame fra 1984 i 240p. Legg merke til byte forskjellen fra NRKs stream.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/applereklame.png)

For å vise hvor stor forkjell kvaliteten på streamen kan utgjøre valgte jeg å streame en 4k video for å sammenligne med superbowl reklamen i 240p.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/4k.png)

NRK lar deg ikke endre streamingkvalitet selv men heller kompanserer etter hvor rask forbinnelsen din er. For moroskyld prøve jeg meg frem på forskjellige youtubevideor med ulik kvalitet for å se hvilke som liknet mest på nrks.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA07/Vedlegg/1080p.png)

Svaret er at NRK mest sannsynlig streamer i 1080p.
