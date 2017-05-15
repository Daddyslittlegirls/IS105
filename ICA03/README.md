# is105/ICA03

Det ligger bilder vedlagt til oppgavene i Bilder mappen som viser resultater lokalt og på instans.

1. 

a) Ikke alle tegn vil vises på enhver terminal. I default terminal på windows vises ikke noe symbol for bytes fra 0x00 til 0x0F ettersom det er ascii control characters.
Det samme resultatet får vi også på skyinstansen, men istedet for et "erstatningsikon" som vi så lokalt på windows, vises det bare empty space.

b) som vist i resultatet på bildene så printes stringen helt likt på instansen, som lokalt. 

c) Ettersom den høyeste decimal verdien i standard ascii table er 127, kan vi iterere gjennom stringen og sjekke at ikke den verdien blir overskredet.

2.

a) 
Lokalt i cmd ser vi at alle ascii tegnene som kan vises er til stede. Forskjellen på instansen er at vi i tillegg ser en "Â" på siden av tegn. I tillegg vises ikke tegnene fra C0 til DF på instansen, men de er synlige lokalt.
%c i koden lar oss vise ikonet som slicen representerer, istenfor %s som heller ville vist byteverdien uten å "oversette" den.
