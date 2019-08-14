betserver je mala app napravljena da slusa na portu 1234 te trenutno ima implementirano:
- povlacenje popisa liga s https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json
- povlacenje popisa ponude s https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json
- GET method API call koji vraca popis liga s popisom mogucih ponuda s putanjom: GetOffersPerLeagues
- GET methodAPI call koji vraca podatke za pojedinu ponudu s putanjom: GetOfferById a parametar koji cita je offer_id
- POST method API call kojim se moze dodati nova ponuda s putanjom: AddOffer koji u bodyu mora dobiti JSON identican onom koji se povlaci s zadanog URLa (moguce je dodati jednu ili vise ponuda istim requestom ovisno o izgledu JSONa)
- POST method API call koji simulira uplatu listica na putanji: NewTicket. Format JSONa je zadan u zadatku

Iskreno, nisam sigaran da li je rijesenje koje sam odabrao pravo, no sve krucijalne podatke drzim u globalnim varijablama. 

Nadalje, aplikacija moze, a ne mora korisiti MySQL bazu podataka koja se mora nalaziti na localhost 3306 port te mora biti rucno kreirana. Naravno podaci za kreiranje baze se nalaze u folderu. Koristenje baze se konfigurira globalnim varijablama: useDB, truncateUsers, truncateTickets

Od problema/stvari koje nisam stigao rijesiti ostalo je:
- povlacenje konfiguracije iz filea
- Pri unosu u bazu puca mi error za hrvatska specificna slova. Nisam stigao to rijesiti
- Nisam rijesio pri validaciji ticketa, provjeru da li je vrijeme za ticket vec proslo (iskreno, nije ovo problem, no parovi koje pullam svi imaju datume u debeloj proslosit
- Kreiranje baze podataka iz aplikacije
- I vjerujem, jos more provjera podataka

Dodatno:
U git folderu se takoder nalazi limitirani set test caseova koje sam koristio tijekom rada na razvoju ove aplikacije. Ovo je standardni export iz Postmana

EDIT: Internal test 20190813 2x
EDIT: Internal test 20190814
