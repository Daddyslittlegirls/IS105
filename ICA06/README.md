ICA 06

Text to speech

Her kjørte vi en container gjennom docker på server som var hostet på port 8080 på instansen vår. 
i tillegg kjørte vi en server kode hostet på port 8090 som viste til en html fil hvor vi lot brukeren skrive inn ønsket tekst i en boks,
for deretter å redirecte til 8080 porten med tilsvarende tekst som paramater i url. Dermed fikk vi automatisk opp text-to-speech lydfilen når vi redirectet.

Server-filen er lånt ifra https://github.com/Zwirc/IS-105/tree/master/ICA06/src og brukes kun for demonstrere html-funksjonaliteten.