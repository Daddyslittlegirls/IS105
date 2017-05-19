# ICA02

## Oppgave 1
#### Hvor mange prosesser som kjører på din datamaskin?
For å komme frem til dette åpnet vi taskmanageren som vi kom frem til ved bruk av kommandoen taskmgr i terminalen.

#### Hvor mange prosesser som kjører på din virtuelle server i nettskyen? 
For å finne ut av hvor mange prosesser som kjørte på den virtuelle server i nettskyen skrev jeg kommandoen “top” i terminalen. I tabellen under vises de forskjellige resultatene gruppemedlemmene fikk.

| Mosvold  | Søvre  | Jonas  | Elias  | Vegard  | Andrea  | Christoffer  | Dennis  |
|---|---|---|---|---|---|---|---|
| 132  | 122  | 125  | 126  | 122  | 124  | 126  | 122  |

#### Kan man gi et nøyaktig antall? Begrunn. Hvor mange av prosessene som “kjører”?
Det er kun 1 prosess som kjører foreløpig. Hvis man kobler seg på instancen sin og skriver “top” i kommandolinjen så kommer det opp en liste som viser alle prosesser som ligger i minnet. Den viser også hvilken prosesser som kjører og hvilken som er i “sleep” mode.
Når en prosess med tilstand “S”(Sleeping) betyr hovedsakelig at prosessen ikke gjør noe. Det er statusen til en prosess som ikke gjør noe.  Hadde den vært i en normal “state” ville den brukt mer prosessorkraft i de fleste tilfellene.


#### Hvis de ikke kjører, hvilke tilstander befinner de seg da?
Om prosessene ikke kjører er de i “sleep/idle” tilstand. De ligger inaktive og venter på input for at de skal kjøre.
I Linux og Windows finnes det forskjellige states en prosess kan være i;
## Linux

#### Running:
Dette er en state en prosess er i enten når den kjører eller når den er klar til å kjøre

#### Interruptible:
dette er en blokkert state. En prosess som venter på at den skal få et signal fra en annen prosess. 

#### Uninterruptible:
Dette er også en blokkert state. Denne prosessen er blokkert til et signal fra hardware’t sier motsatt
Stopped: når prosessen er ferdig, vil det stå “Stopped, da kan man starte den på nytt.

#### Zombie:
i denne staten vil prossessen bli terminert og informasjonen vil bli synlig.

## Windows

#### Created:
Dette er når en prosess er laget og den avventer instrukser. 

#### Ready:
Dette er når en prosess er klar og venter på utførelse av CPU

#### Running:
Dette er når prosessen kjører og er klar for å bli utført. 

#### Blocked:
Dette er en blocked state, når en prosess ikke fungerer, for eksempel at du skal printe noe, men du har ikke koblet til en printer.

#### Terminated:
Dette er staten til prosessen når den enten er ferdig med jobben sin, eller når den har blitt fjerna. 



#### Hva er maskinvarespesifikasjon til din datamaskin (noter prosessortype, prosessorarkitektur, klokkefrekvens, informasjon om primært minne, størrelse på cache (både L1, L2 og L3 er ønskelig))?
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA02/Vedlegg/ICA0201.png)

#### Hvor mange CPU-“cores” har du tilgjengelig på din maskin?
| Mosvold  | Søvre  | Jonas  | Vegard  | Andrea  | Christoffer  | Elias | Dennis|
|---|---|---|---|---|---|---|---|
| 2  | 2  | 4  | 4  | 2  | 4  | 4 | 4 |

#### Noter. Hvor mange CPU-”cores” har du tilgjengelig på din virtuelle server? 
Ved å skrive inn kommandoen "lscpu" på den virtuelle serveren fant jeg ut av hvor mange CPU-”cores”. Resultatene under viser hva de forskjellige gruppemedlemmene fikk.

| Mosvold  | Søvre  | Jonas  | Vegard  | Andrea  | Christoffer  | Elias | Dennis |
|---|---|---|---|---|---|---|---|
| 2  | 2  | 2  | 2  | 2  | 2  | 2  | 2  |

#### Noter. Finn ut hvilken prosess i ditt system bruker mest minne. Beskrive denne prosessen kort.
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA02/Vedlegg/ram.png)
Kernel is utilized by the OS for all sort of operations such as disk writes, cpu controls, etc.

På våres personlige PC-er har vi ulike programmer som er nedlastet og installert, vi bruker forskjellige operativsystemer, noen bruker mac, andre bruker windows. Mens på linux skyen har vi likt med programmer og det kjøres på et samme system. Derfor vil det være en forskjell på våres personlige pcer og mindre forskjell på den virituelle serveren vår.

Et program som inneholder masse instrukser, så om vi starter et program på pcn våres eller på skyen vil det programmet utføre alle instruksene som er gitt

#### Hvilke komponenter (både fysiske og abstrakte) i deres datasystemer er involvert i oppstart, administrasjon og avslutning av prosesser? Definer komponentene du nevner

Bios står for “Basic input og output system”. Bios er programmet I en pc's microprocessor som får  maskinen til å starte opp etter du har skrudd den på. Mikroprosessoren er datamaskinens processor på en microchip. Det er Biosens oppgave å “vekke” resten av maskinen når du skrur den på.  Du kan gjerne sammenligne det med å skru om nøkkel I en bil og motoren starter. Når du starter maskinen gir mikroprosessoren kontrollen til bios programmet.  Det er slags link mellom software hardware. Den håndterer dataflyt mellom maskinens operativsystem og tilkoblede enheter som harddisk, skjermkort, mus og tastatur.

## Oppgave 2
Skriv koden i hello.go filen ved å bruke “nano” kommandoen på serveren. 

Da får vi mulighet til å legge teksten rett inn i filen gjennom terminalen. 
Når vi har limt inn teksten så kan vi lagre den direkte på serveren våres. Deretter kan vi lage en executable fil for mac og en for windows systemet. 

Skrev GOOS=darwin for mac
GOOS=windows for windows
Dette skal skrives for “go build hello.go” for å velge hvilken plattform du vil lage en executable for.


For å gjøre det på windows trenger man å skrive “go build hello.go “, da vil man få en exe fil som kan kjøres på windows. For å builde en fil på serveren må vi skrive 
“GOOS=linux GOARCH=amd64 go build v hello.go “	da vil man få en hello fil som er kjørbar på linux serveren.

For å overføre den filen fra pcn til serveren brukte vi filezilla. Når filen var over hadde den ikke rettigheter til å kjøres, så da kunne man skrive chmod 777 eller 744. Da vil man få rettigheter på serveren for å kjøre filen. 

777 = alle rettigheter
744 = bare eier



![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA02/Vedlegg/ICA0203.png)

Her ser man “hello” filen som kan kjøres på mac og “hello.exe” som kan kjøres på windows.
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA02/Vedlegg/ICA0204.png)
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA02/Vedlegg/ICA0205.png)

## Oppgave 3
sum_test oppgir feil hvis tallene i sum_tests_int8 (struct) overstiger den gitte talltypen. For eksempel hvis structen består av int8 variabler vil alt under -128 eller over 127 overstige int8. Dette er fordi int8 er en datatype som består av en byte. 

## Oppgave 4
Sorteringen i aksjon.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA02/Vedlegg/ICA0210.png)

Benchmarking i aksjon.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA02/Vedlegg/ICA0209.png)

Bubble sort tester to tall opp mot hverandre, og bytter plass på dem hvis de er i feil rekkefølge. Bubble sort er ganske raskt, hvis det er snakk om et lite antall data. Quicksort bruker en "splitt-og-hersk" metode Første velger den er tilfeldig tall, og så sorteresr det  tallene i 3 deler. En del som er mindre enn det tilfeldige taller, en del som er det samme som det tilfeldige tallet og en del som er større enn det tilfeldige tallet. Fordelen med Quicksort er at det er veldig raskt og effektivt.

##### Kilde: https://github.com/ZachOrr/golang-algorithms/blob/master/sorting/bubble-sort.go

## Oppgave 5
#### Hva kan du si noe om antall prosesser og tråder, som programmet bruker på ditt system?
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA02/Vedlegg/ICA0206.png)

Øverste bilde ser vi at det er 4 tråder kjørende på serveren i denne prosessen boring01.
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA02/Vedlegg/ICA0207.png)

Tråder: 5
Prosess: 1
#### Hvilken tilstand befinner den seg i?
Aktiv. 
#### Hvordan kan du stoppe prosessen? 
Ctrl + C
										
Dette skjermbildet viser prosessen boring 01 som kjører på serveren.
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA02/Vedlegg/ICA0208.png)
