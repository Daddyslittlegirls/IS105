# ICA04

# Oppgave 1

## A
I Windows har filene samme størrelse, men på GitHub har de forskjellig størrelse. De har like “byte slices”. Grunnlag for at den ene er større enn den andre er at det brukes forskjellige metoder for linjeskift på Linux, Mac og Windows.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.32.58.png)

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.33.10.png)

Slik ser det ut på en Mac

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.33.18.png)

## B
“Text1.txt” bruker \r\n (CR+LF) for linjeskifte, som er vanlig i Windows.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.33.26.png)

# Oppgave 2

## A
Koden vi brukte:

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.33.37.png)

Resultat fra instansen når vi sjekker filen "pg100.txt":

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.33.44.png)

## B
Resultat for "dev/stdin":

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.33.49.png)

Resultat for "dev/ram0":

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.33.54.png)

## C
Resultat på Mac:

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.34.00.png)

Resultat på Windows:

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/oppg2c.png)

Mac og Windows får nesten like resultatet når vi sjekker “pg100.txt”, det eneste som skiller Mac og Windows er rettighetene til filen og hvordan størrelsen blir representert. Windows har større rettigheter enn på mac. På instansen får vi “is not a regular file” og “Is a device file”, men det er omvendt på mac og windows. 

# Oppgave 3

## A
Os-pakken inneholder flere metoder og types for behandling av filer. 
Fileinfo typen lar oss returnere informasjon om filen slik som navn, størrelse i bytes, mode, mod time, sjekker directory (bool), og data source. 
Ellers har vi metoder som lar oss direkte behandle filen slik som rename, delete, copy osv. 

Vi har også en pakke som heter bufio som lar oss jobbe med filer i en buffer. Det vil si at vi kan gjøre midlertidige forandringer før vi skriver de til disken. I bufio pakken finner vi også scanner typen som kan opprettes for å lese data av for eksempel en fil, linjevis. 

Zip & compress/gzip pakkene lar oss skrive direkte til .zip filer, og compresse før vi skriver informasjon videre til den underliggende filskriveren.

path/filepath pakken gir oss muligheten til å endre på filstien, i tråd med det designerte operativsystemet.

Ioutil pakken inneholder flere funksjoner som for eksempel lar oss åpne en fil, skrive en byteslice til den og lukke filen igjen. 

net/http lar oss laste ned filer gjennom http

## B
Finner antall linjeskift

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.34.11.png)

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.34.17.png)

## C

# Oppgave 4

## A
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.34.22.png)

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.34.27.png)

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.34.33.png)

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.34.39.png)

## B
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA04/Bilder/Skjermbilde%202017-05-19%20kl.%2014.34.43.png)
