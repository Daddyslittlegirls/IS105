# IS105

Main kjøres med: go run main.go log.go

locli kjøres med: <filnavn> Arg1(float)

logbcli kjøres med: <filnavn> Arg1(float) Arg2(int)
____________________________________________________

# ICA01

**1.2.1**

Vi brukte denne metoden som ble vist i undervisningen:
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA01/Bilder%20ICA01/Oppgaver%20Ica1.2.jpg)

a)  1 = 00000001. 8 bits.

b)  2 = 00000010. 8 bits.

c)  5 = 00000101. 8 bits.

d)  8 = 00001000. 8 bits.

e)  16 = 0010000. 8 bits.

f)  256 = 0000000100000000. 16 bits.

Vi brukte denne metoden for å løse oppgaven:
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA01/Bilder%20ICA01/Oppgaver%20ica1.jpg)
a)  100 = 4

b)  1001 = 9

c)  1100110011 = 411

For å konvertere de binære tallene til titallsystemet brukte vi potenser av 2. Vi begynte med den laveste potensen (2^0) og startet fra høyre og leste til venstre. For hvert binærtall øker vi opphøyning med 1. (2^1*1, 2^2*1). Så må 2-tallspotensen bare multipliseres med det binære tallet i rekken.


**1.2.2**

Flere personer prøver å gjette et tresifret (3-bit) binært tall.
(1) Lise har fått vite / lærer at tallet er et oddetall.
(2) Per har fått vite at tallet er IKKE et multiplum av 3 (dvs. ikke 0, 3, 6).
(3) Oskar har fått vite at tallet inneholder nøyaktig 2 enere.
(4) Louise har fått vite alt det Lise, Per og Oskar vet.
Hvor mye informasjon (i bits) har hver spiller fått?

8 alternativer siden det er et tresifret binært tall
 
I denne oppgaven er det tilsammen 8 muligheter med 3 sifret binære tall. Dvs. at N=8. M, begrenser valgmulighetene. Siden de forskjellige personene har ulik informasjon om dette tallet varierer M fra person til person. I dette tilfellet ser vi at svaret ikke lander på hele bits, men for eksempel ~0.7 bits som er mulig dersom det er flere mulige utfall, som også er forskjellig vektlagt.

Log^2(1/(M/N)) = log^2(N/M) bits

(1) Log^2(1/(4/8)) = log^2(8/4) = log^2(2)= 1 bit

(2) Log^2(1/(5/8)) = log^2(8/5)= log^2(1,6)= 0.6780719051126377 bit

(3) Log^2(1/(3/8)) = log^2(8/3)) = log^2(2.67=3) = 1,6 eller 2 bit

(4) Log^2(1/(1/8) = log^2(8/1)) = log^2(8) = 3 bit


**1.2.3:**

Link til våre Githubs:
https://github.com/Chopflox/is105-uke04v2
https://github.com/vegardst/is105-uke04
https://github.com/mosvold/is105-uke4
https://github.com/andreaweert/is105-uke04
https://github.com/JonasOmdal/Is105-uk4-ny
https://github.com/Sovre/IS105-uke04-egen
https://github.com/ChristofferO/is105-uke04-ny
 
Link til vår Gruppegithub:
https://github.com/Daddyslittlegirls/IS105


**1.2.4**

1)	Hvilken fordeler og ulemper har en git-flow-modell med en hovedrepository?

Fordeler:
- Alle medlemmer kan jobbe med sin egen branch.
- Man slipper merge konflikter
- Man har oversikt over hva andre jobber med, og kan enkelt hente og dele fra branches.
 
Ulemper:
- Man må kontinuerlig oppdatere når endringer gjøres, slik at man ikke havner for langt bak master branchen.
- Hvis man jobber samtidig med noe så kan det bli overskrevet dersom man ikke er oppmerksom.

2)  Finn ut hva heter objektfiler for de mest brukte platformer (Unix/Linux, MS Windows, Mac OS x)! Hvorfor, etter deres mening, har disse plattformene så forskjellig objektfil-formater?

●   COFF (Unix/Linux) / Executable and Linkable Format (ELF)
○     none, .o, .obj
○     none, .axf, .bin, .elf,.o, .prx, .puff and .so
●   Mach-O (Mac OSx)
○     none, .o, .dylib,.bundle
●   Portable Executable (PE) (Windows)
○     .acm, .ax, .cpl, .dll, .drv, .efi,.exe, .mui, .ocx, .scr, .sys, .tsp	
●   De har forskjellige objektfil-formater på grunn av gamle standarder/gamle systemer.


**3)	Hvilke forskjeller ser dere i forhold til programmeringsspråket Java?**

- Det er ikke objektorientert programmeringsspråk.
- Compile tiden generelt mye kjappere.
- Mindre kode
- Kompakt. Flere unødvendig tegn har blitt redusert. I mange andre språk er det tegn som er meningsløse, men som vi har blitt vandt til     ettersom de forskjellige språkene har eksistert lang tid.
- Siktet mot HTTP
- Et slags funksjonelt språk
- Skrive funksjoner inne i funksjoner, og kan returnere funksjoner akkurat som i et funksjonelt språk og lokale variabler rundt det         tjener som variabler i en closure.
- Funksjoner returner ikke bare en enkel verdi, men flere verdier.


**4)	Hvilke viktige poeng illusterer denne øvelsen når det gjelder bruk av et programmeringsmiljø på en plattform?**

Et programmeringsmiljø lar deg ta i bruk pakker fra standardbiblioteket ved hjelp av environment variabler i systemet. Du kan sette workpath til prosjekter du holder på med slik at systemet vet hvor det kan finne pakkene du referer til.


**5)	Er det hensiktsmessig å legge inn denne filen i git repository? Begrunn svaret!**  	

Github er ment for å dele kildekode og i lengden vil det være unødvendig mye arbeid å oppdatere den kjørbare filen konstant, ettersom endringer er hyppige. I tillegg er kjørbare filer som regel spesifikke til enkelte operativsystem, og derfor ikke gunstige når de skal deles blant flere folk. 


**6)	Hvordan skiller pakken log, som dere har implementert, seg fra andre pakker i go, for eksempel, fmt?**

Pakken log stammer fra lokalt, den er selvlaget og skal beregne logaritme med base 2 og 10. Denne pakken importeres lokalt, mens feks fmt blir importert fra standardbiblioteket




