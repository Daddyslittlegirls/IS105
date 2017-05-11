# IS105

Gruppe 6: Daddy's little girls

1.2.1

Vi brukte dele metoden som ble vist i undervisningen:

1) 1 = 1. 1 bits.
	2) 2 = 10. 2 bits.
	3) 5 = 101. 3 bits.
	4) 8 = 1000. 4 bits.
	5) 16 = 10 000. 5 bits.
	6) 256 = 100 000 000. 9 bits.

For å løse den andre delen av oppgaven brukte vi denne metoden:
 
	1) 100 = 4
	2) 1001 = 9
	3) 1100110011 = 411



1.2.2

Flere personer prøver å gjette et tresifret (3-bit) binært tall. 
(1) Lise har fått vite / lærer at tallet er et oddetall. 
(2) Per har fått vite at tallet er IKKE et multiplum av 3 (dvs. ikke 0, 3, 6). 
(3) Oskar har fått vite at tallet inneholder nøyaktig 2 enere. 
(4) Louise har fått vite alt det Lise, Per og Oskar vet. 
Hvor mye informasjon (i bits) har hver spiller fått?

8 alternativer siden det er et tresifret binært tall

Log^2(1/(M/N)) = log^2(N/M) bits

(1) Log^2(1/(4/8)) = log^2(8/4) = log^2(2)= 1 bit

(2) Log^2(1/(5/8)) = log^2(8/5)= log^2(1,6)= 0.6780719051126377 bit

(3) Log^2(1/(3/8)) = log^2(8/3)) = log^2(2.67=3) = 1,6 eller 2 bit

(4) Log^2(1/(1/8) = log^2(8/1)) = log^2(8) = 3 b

1.2.3: 
Link til våre Githubs:
https://github.com/Chopflox/is105-uke04v2
https://github.com/vegardst/is105-uke04
https://github.com/mosvold/is105-uke4
https://github.com/andreaweert/is105-uke04
https://github.com/JonasOmdal/Is105-uk4-ny
https://github.com/Sovre/IS105-uke04-egen
https://github.com/ChristofferO/is105-uke04-ny
https://github.com/bjornarguttormsen

Link til vår Gruppegithub: 
https://github.com/Daddyslittlegirls/IS105















1.2.4
Hvilken fordeler og ulemper har en git-flow-modell med en hovedrepository?
	Fordeler:
Alle medlemmene kan jobbe med sin egen branch. 
Man slipper merge konflikter

	Ulemper:
Man må ha oversikt over hva alle andre jobber med.
Hvis man jobber samtidig med noe så kan det bli overskrevet dersom man ikke er oppmerksom.

Finn ut hva heter objektfiler for de mest brukte platformer (Unix/Linux, MS Windows, Mac OS x)! Hvorfor, etter deres mening, har disse plattformene så forskjellig objektfil-formater?
COFF (Unix/Linux) / Executable and Linkable Format (ELF)
none, .o, .obj
none, .axf, .bin, .elf,.o, .prx, .puff and .so
Mach-O (Mac OSx)
none, .o, .dylib,.bundle
Portable Executable (PE) (Windows)
.acm, .ax, .cpl, .dll, .drv, .efi,.exe, .mui, .ocx, .scr, .sys, .tsp	
De har forskjellige objektfil-formater på grunn av gamle standarder/gamle systemer.

Hvilke forskjeller ser dere i forhold til programmeringsspråket Java?
Det er ikke objektorientert programmeringsspråk.
Compile tiden generelt mye kjappere.
cpu loaden er generelt mindre.

Hvilke viktige poeng illusterer denne øvelsen når det gjelder bruk av et programmeringsmiljø på en plattform?
Et programmeringsmiljø inkluderer automatisk fil og mappestruktur og gjør det lettere å debugge kode. I tillegg kan et programmeringsmiljø være svært effektivt ettersom man ikke trenger å gjøre alt manuelt. Man har gjerne autofyll og forslag på funksjoner.

Er det hensiktsmessig å legge inn denne filen i git repository? Begrunn svaret!https://github.com/Daddyslittlegirls/IS105/tree/master/log	
I dette tilfellet vil det være hensiktsmessig å legge filen i et repository, for å teste om koden vi har skrevet funker. I lengden vil det være unødvendig mye arbeid å oppdatere den kjørbare filen konstant, ettersom endringer er hyppige.


Hvordan skiller pakken log, som dere har implementert, seg fra andre pakker i go, for eksempel, fmt?
https://github.com/Daddyslittlegirls/IS105/tree/master/log
	
	Log er en pakke inne i “math” mens for eksempel fmt er bare en pakke. 
	Log er for å logge mens fmt er for å formatere. 

