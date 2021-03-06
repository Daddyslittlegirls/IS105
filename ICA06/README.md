# ICA06

# Eksperiment 1
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA06/Bilder/Skjermbilde%202017-05-19%20kl.%2016.21.14.png)
Vi har spilt inn lyd fra hver person i gruppen og sammelignet. Bildene viser ordet : "hjelp" På de øverste bildene kan vi se at det er mye bakgrunnsstøy, men vi kan se at det er forskjell på tonefall og dialekt. Hvis vi skulle implementert dette i et program, er det viktig at programmet kan forstå forskjellige tonefall og dialekter. Programmet bør dele opp det personen sier i delte organiserte deler i en akustisk database med lagret språk. Ideelt vil vi ha et speech to text-program som oversetter hver setning istedet for vært ord. Dette er fordi et ord kan høres vidt forskjellig ut etter hvilken kontekst det er i, for at dette skal bli mulig må programmet på forhånd vite hvordan et ord høres ut i et satt miljø. Et slikt program må ta hensyn til srpåk, dialekt og slang.

# Eksperiment 2
Tekst til tale:
Her kjørte vi en container gjennom docker på server som var hostet på port 8080 på instansen vår. 
i tillegg kjørte vi en server kode hostet på port 8090 som viste til en html fil hvor vi lot brukeren skrive inn ønsket tekst i en boks,
for deretter å redirecte til 8080 porten med tilsvarende tekst som paramater i url. Dermed fikk vi automatisk opp text-to-speech lydfilen når vi redirectet.

![Docker container status](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA06/Bilder/Eksperiment%202%20-%20Docker%20container.png)
![Html-siden](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA06/Bilder/Eksperiment%202%20-%20html-siden.png)
![Resultatet av redirect](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA06/Bilder/Eksperiment%202%20-%20Resultat.png)

Server-filen er lånt ifra https://github.com/Zwirc/IS-105/tree/master/ICA06/src og brukes kun for demonstrere html-funksjonaliteten.

# Eksperiment 3
Tale til tekst:
Vi brukte eksempelet:https://github.com/amsehili/gspeech-rec. 
Vi lastet ned HomeBrew for å installere flac og wget. HomeBrew er en package-manager for Mac. Samme som at APT er for Linux. Dette er tilleggsprogrammer scriptet trenger for å kjøre. Dette kalles Dependencies.
For å kjøre shell-filen, bruker man kommandoen ./speech-rec.sh. Koden ligger under filen speech-rec.sh.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA06/Bilder/Tale%20til%20tekst.png)
