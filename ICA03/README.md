# ICA03

# Oppgave 1

Go filene har egne main funksjoner og bruker main package bare for å gjøre det enklere å teste hver enkelt fil lokalt og på instansen.


## A
Ikke alle tegn vil vises på enhver terminal. I default terminal på windows vises ikke noe symbol for bytes fra 0x00 til 0x0F ettersom det er ascii control characters.
Det samme resultatet får vi også på skyinstansen, men istedet for et "erstatningsikon" som vi så lokalt på windows, vises det bare empty space.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/iterateOverASCIIStringLiteral%20-%20lokalt.png)
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/iterateOverASCIIStringLiteral%20-%20instans.png)

## B
Som vist i resultatet på bildene så printes stringen helt likt på instansen, som lokalt. 

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/greetingASCII%20-%20lokalt.png)
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/greetingASCII%20-%20instans.png)

## C
Ettersom den høyeste decimal verdien i standard ascii table er 127, kan vi iterere gjennom stringen og sjekke at ikke den verdien blir overskredet.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/TestGreetingASCII%20-%20resultat.png)

# Oppgave 2

## A
Lokalt i cmd ser vi at alle ascii tegnene som kan vises er til stede. Forskjellen på instansen er at vi i tillegg ser en "Â" på siden av tegn. I tillegg vises ikke tegnene fra C0 til DF på instansen, men de er synlige lokalt.
%c i koden lar oss vise ikonet som slicen representerer, istenfor %s som heller ville vist byteverdien uten å "oversette" den.
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/IterateOverExtendedASCIIStringLiteral%20-%20lokalt.png)
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/IterateOverExtendedASCIIStringLiteral%20-%20instans.png)

## B
Vi ser at default print metoden viser tegnene helt fint på både cmd og git bash. Men på instansen er det enkelte tegn som ikke vises. 
Her kan det være bedre å se på andre verb for print som foreksempel %c, eller verb som har escape code så man kan se forskjellen på de.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/GreetingExtendedASCII%20-%20instans.png)
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/GreetingExtendedASCII%20-%20git%20bash.png)
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/GreetingExtendedASCII%20-%20cmd.png)
## C
Her sjekker koden hvert tegn for å se om verdien tilsvarer et av de vanlige ASCII tegnene, istedet for et av de utvidede. Ettersom det var vanlige tegn i byteslicen \x20\ for eksempel, så resulterte testen i fail.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/TestExtendedGreetingASCII%20-%20resultat.png)

# Oppgave 3

## A

Her eksperimenterte vi litt med de forskjellene konverteringene, og representeringene av symboler. Ettersom eksperimentering er kommentert gjennom koden, legger vi med et bilde av den her. 

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/Oppgave%203a%20-%20Kode.png)

Og resultet av å kjøre koden ser slik ut 
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/Oppgave%203a%20-%20Resultat.png)

## B

Utskriften fra hver av filene ser slik ut. 

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/lang01%20-%20default.png)
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/lang02%20-%20default.png)
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/lang03%20-%20default.png)

Vi lagde en funksjon som iterer over hele byteslicen og erstatter "Ø" i dette tilfellet med "O" for å demonstrere.

![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/editLetter.png)

Da ble resultet slik 
![](https://github.com/Daddyslittlegirls/IS105/blob/master/ICA03/Bilder/lang03%20-%20erstattet.png)
