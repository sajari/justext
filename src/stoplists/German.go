package gojustext

var GermanStoplist = map[string]bool {"der": true,
"und": true,
"die": true,
"in": true,
"von": true,
"den": true,
"des": true,
"im": true,
"mit": true,
"dem": true,
"das": true,
"ist": true,
"Die": true,
"er": true,
"zu": true,
"als": true,
"wurde": true,
"eine": true,
"sich": true,
"für": true,
"ein": true,
"auf": true,
"an": true,
"war": true,
"auch": true,
"nach": true,
"bis": true,
"aus": true,
"Der": true,
"bei": true,
"einer": true,
"am": true,
"zum": true,
"sie": true,
"durch": true,
"einem": true,
"nicht": true,
"oder": true,
"es": true,
"sind": true,
"Das": true,
"zur": true,
"wird": true,
"In": true,
"einen": true,
"Im": true,
"über": true,
"um": true,
"dass": true,
"Er": true,
"unter": true,
"wie": true,
"werden": true,
"wurden": true,
"nur": true,
"Jahr": true,
"noch": true,
"seiner": true,
"vom": true,
"aber": true,
"Nach": true,
"seine": true,
"vor": true,
"Jahre": true,
"sowie": true,
"Sie": true,
"eines": true,
"zwischen": true,
"dieser": true,
"zwei": true,
"hat": true,
"jedoch": true,
"so": true,
"Es": true,
"Jahren": true,
"gegen": true,
"werden.": true,
"Zeit": true,
"Stadt": true,
"seit": true,
"man": true,
"ersten": true,
"waren": true,
"Am": true,
"wieder": true,
"wurde.": true,
"etwa": true,
"diese": true,
"sein": true,
"Als": true,
"hatte": true,
"seinen": true,
"seinem": true,
"mehr": true,
"drei": true,
"ihre": true,
"später": true,
"heute": true,
"Bei": true,
"liegt": true,
"erste": true,
"beim": true,
"de": true,
"kann": true,
"dann": true,
"Von": true,
"Mit": true,
"Ende": true,
"of": true,
"wo": true,
"Diese": true,
"anderen": true,
"ihm": true,
"bereits": true,
"ab": true,
"Teil": true,
"da": true,
"ihrer": true,
"Ein": true,
"während": true,
"sehr": true,
"diesem": true,
"gibt": true,
"dort": true,
"Seit": true,
"Eine": true,
"kam": true,
"ihn": true,
"ist.": true,
"beiden": true,
"Mai": true,
"allem": true,
"Gemeinde": true,
"konnte": true,
"Auch": true,
"August": true,
"alle": true,
"Januar": true,
"Oktober": true,
"dessen": true,
"ihr": true,
"März": true,
"keine": true,
"Juni": true,
"April": true,
"September": true,
"damit": true,
"Juli": true,
"Dezember": true,
"ihren": true,
"Jahrhundert": true,
"vier": true,
"deren": true,
"Auf": true,
"Geschichte.": true,
"November": true,
"gab": true,
"erhielt": true,
"war.": true,
"Mitglied": true,
"zunächst": true,
"deutschen": true,
"Namen": true,
"haben": true,
"immer": true,
"an.": true,
"bzw.": true,
"was": true,
"wird.": true,
"Deutschland": true,
"auf.": true,
"Für": true,
"Kirche": true,
"weitere": true,
"schon": true,
"hier": true,
"Universität": true,
"begann": true,
"Februar": true,
"denen": true,
"selbst": true,
"sondern": true,
"ebenfalls": true,
"ohne": true,
"befindet": true,
"dabei": true,
"anderem": true,
"Ort": true,
"können": true,
"neuen": true,
"Während": true,
"neue": true,
"So": true,
"einige": true,
"ist,": true,
"ins": true,
"werden,": true,
"erst": true,
"Sohn": true,
"Seine": true,
"war,": true,
"wenn": true,
"nahm": true,
"dieses": true,
"großen": true,
"Durch": true,
"Leben.": true,
"allerdings": true,
"Familie": true,
"deutscher": true,
"andere": true,
"seines": true,
"neben": true,
"Deutschen": true,
"viele": true,
"Da": true,
"Saison": true,
"wurde,": true,
"zusammen": true,
"diesen": true,
"wegen": true,
"Sein": true,
"trat": true,
"spielte": true,
"m": true,
"erstmals": true,
"Platz": true,
"Zu": true,
"New": true,
"große": true,
"Jahrhunderts": true,
"ging": true,
"Film": true,
"sind.": true,
"ein.": true,
"mehrere": true,
"km": true,
"ihrem": true,
"besteht": true,
"welche": true,
"meist": true,
"letzten": true,
"Ab": true,
"gehört": true,
"B.": true,
"fünf": true,
"wurden.": true,
"gewann": true,
"rund": true,
"zweiten": true,
"Menschen": true,
"nun": true,
"St.": true,
"Leben": true,
"Geschichte": true,
"Prozent": true,
"oft": true,
"gehörte": true,
"aus.": true,
"schließlich": true,
"studierte": true,
"wobei": true,
"Dies": true,
"Anfang": true,
"Art": true,
"ab.": true,
"Titel": true,
"arbeitete": true,
"führte": true,
"Tod": true,
"Beispiel": true,
"Gebiet": true,
"Form": true,
"also": true,
"Friedrich": true,
"deutsche": true,
"verschiedenen": true,
"Nachdem": true,
"Meter": true,
"Band": true,
"Zweiten": true,
"Dieser": true,
"Name": true,
"fast": true,
"fand": true,
"Dabei": true,
"aufgrund": true,
"Bis": true,
"Neben": true,
"an,": true,
"weiteren": true,
"steht": true,
"Berlin": true,
"the": true,
"Johann": true,
"erfolgte": true,
"Mitte": true,
"Bereich": true,
"gilt": true,
"König": true,
"ca.": true,
"weiter": true,
"Außerdem": true,
"soll": true,
"jeweils": true,
"heutigen": true,
"besonders": true,
"Dort": true,
"insbesondere": true,
"Unternehmen": true,
"stark": true,
"Um": true,
"eigenen": true,
"hatte,": true,
"Karl": true,
"hatte.": true,
"sechs": true,
"neu": true,
"Bau": true,
"Gruppe": true,
"innerhalb": true,
"An": true,
"worden": true,
"statt.": true,
"blieb": true,
"findet": true,
"auf,": true,
"meisten": true,
"dies": true,
"entstand": true,
"insgesamt": true,
"Provinz": true,
"damals": true,
"Entwicklung": true,
"Region": true,
"allen": true,
"Weltkrieg": true,
"zahlreiche": true,
"musste": true,
"Aus": true,
"II.": true,
"Beginn": true,
"kommt": true,
"sowohl": true,
"Frau": true,
"Unter": true,
"zurück.": true,
"Millionen": true,
"ehemaligen": true,
"Wilhelm": true,
"übernahm": true,
"sollte": true,
"Mal": true,
"Kilometer": true,
"folgenden": true,
"Liste": true,
"Heinrich": true,
"wechselte": true,
"vielen": true,
"darauf": true,
"Vater": true,
"verschiedene": true,
"and": true,
"Arbeit": true,
"USA": true,
"hatten": true,
"können.": true,
"häufig": true,
"weil": true,
"Amt": true,
"beispielsweise": true,
"doch": true,
"John": true,
"Rahmen": true,
"bekannt": true,
"alten": true,
"Weblinks.": true,
"Bevölkerung": true,
"lange": true,
"bezeichnet": true,
"zog": true,
"daher": true,
"je": true,
"führt": true,
"einigen": true,
"erreichte": true,
"Gebäude": true,
"davon": true,
"Burg": true,
"teilweise": true,
"Norden": true,
"ließ": true,
"Seite": true,
"Folge": true,
"Rolle": true,
"gegenüber": true,
"I.": true,
"zweite": true,
"Zur": true,
"muss": true,
"eigene": true,
"ihnen": true,
"Land": true,
"Album": true,
"französischen": true,
"erneut": true,
"dazu": true,
"Tochter": true,
"Richtung": true,
"Dorf": true,
"hat.": true,
"somit": true,
"Süden": true,
"Krieg": true,
"sein.": true,
"wird,": true,
"erschien": true,
"Sitz": true,
"zwar": true,
"kurz": true,
"geht": true,
"z.": true,
"Begriff": true,
"Peter": true,
"Höhe": true,
"Bedeutung": true,
"aller": true,
"lang": true,
"haben.": true,
"Verein": true,
"welches": true,
"gleichen": true,
"Hier": true,
"Landkreis": true,
"Regierung": true,
"Nähe": true,
"starb": true,
"danach": true,
"Alter": true,
"finden": true,
"kann.": true,
"kleinen": true,
"kein": true,
"Bezeichnung": true,
"stand": true,
"Paul": true,
"sogar": true,
"setzte": true,
"etwas": true,
"Partei": true,
"direkt": true,
"sei": true,
"Werk": true,
"Zum": true,
"v.": true,
"Mannschaft": true,
"Einwohner": true,
"folgte": true,
"größten": true,
"Haus": true,
"zurück": true,
"weniger": true,
"darunter": true,
"Insel": true,
"Schloss": true,
"FC": true,
"Westen": true,
"aus,": true,
"kleine": true,
"Bereits": true,
"Straße": true,
"Dieses": true,
"Professor": true,
"konnten": true,
"gemeinsam": true,
"Danach": true,
"Grund": true,
"gut": true,
"Strecke": true,
"gelang": true,
"gehören": true,
"ganz": true,
"stellte": true,
"nachdem": true,
"Hans": true,
"hinaus": true,
"größte": true,
"waren.": true,
"stellt": true,
"Karriere": true,
"kamen": true,
"Kinder": true,
"zudem": true,
"vor.": true,
"machte": true,
"außerdem": true,
"gründete": true,
"pro": true,
"sind,": true,
"Jahres": true,
"mehreren": true,
"Zwischen": true,
"Den": true,
"letzte": true,
"\"The": true,
"ein,": true,
"Deutsche": true,
"ursprünglich": true,
"Personen": true,
"beträgt": true,
"sieben": true,
"deutlich": true,
"Osten": true,
"Gründung": true,
"Musik": true,
"Franz": true,
"hohen": true,
"Gemeinden": true,
"Schule": true,
"erster": true,
"Wasser": true,
"Damit": true,
"Heute": true,
"besuchte": true,
"Frankreich": true,
"acht": true,
"befinden": true,
"Regel": true,
"Weitere": true,
"Spiel": true,
"Erst": true,
"zehn": true,
"zuvor": true,
"tätig.": true,
"Teile": true,
"Kloster": true,
"Stelle": true,
"Ihre": true,
"Kaiser": true,
"Studium": true,
"Chr.": true,
"gegründet.": true,
"Gegensatz": true,
"Einsatz": true,
"konnte.": true,
"York": true,
"entwickelte": true,
"beide": true,
"Reihe": true,
"Ersten": true,
"Gesellschaft": true,
"ihres": true,
"bekannt.": true,
"Sommer": true,
"Staaten": true,
"Ausbildung": true,
"handelt": true,
"eher": true,
"zeigt": true,
"Berliner": true,
"weit": true,
"hinter": true,
"Ludwig": true,
"Mitglieder": true,
"bezeichnet.": true,
"Frauen": true,
"Lage": true,
"Truppen": true,
"University": true,
"wenig": true,
"The": true,
"Bruder": true,
"Georg": true,
"Erfolg": true,
"Verbindung": true,
"Firma": true,
"schrieb": true,
"Über": true,
"Kreis": true,
"politischen": true,
"\"Die": true,
"einzige": true,
"einmal": true,
"ebenso": true,
"damaligen": true,
"Hilfe": true,
"Serie": true,
"Mutter": true,
"Michael": true,
"Buch": true,
"kehrte": true,
"lässt": true,
"Ziel": true,
"Tag": true,
"Fläche": true,
"umfasst": true,
"Wien": true,
"bestand": true,
"Olympischen": true,
"bisher": true,
"Hälfte": true,
"Präsident": true,
"Nachfolger": true,
"hauptsächlich": true,
"Arten": true,
"stammt": true,
"worden.": true,
"Maria": true,
"Mann": true,
"Welt": true,
"Werke": true,
"Europa": true,
"Spieler": true,
"Aufgrund": true,
"dafür": true,
"Beim": true,
"Robert": true,
"stehen": true,
"München": true,
"frühen": true,
"entstanden": true,
"britischen": true,
"lag": true,
"ob": true,
"Betrieb": true,
"San": true,
"Besitz": true,
"Vereinigten": true,
"spielt": true,
"selben": true,
"wurden,": true,
"Ihr": true,
"Später": true,
"verwendet.": true,
"viel": true,
"hohe": true,
"Linie": true,
"südlich": true,
"Schweizer": true,
"Team": true,
"späteren": true,
"anschließend": true,
"Weg": true,
"Vertrag": true,
"Wiener": true,
"Wie": true,
"Zudem": true,
"gewählt.": true,
"Bahnhof": true,
"veröffentlichte": true,
"County": true,
"Schlacht": true,
"kommen": true,
"Armee": true,
"Thomas": true,
"Herrschaft": true,
"schloss": true,
"Österreich": true,
"weiterhin": true,
"bald": true,
"Spielen": true,
"einzelnen": true,
"nördlich": true,
"liegen": true,
"zahlreichen": true,
"Vor": true,
"nahe": true,
"bevor": true,
"bildet": true,
"Man": true,
"lediglich": true,
"Zahl": true,
"wohl": true,
"ehemalige": true,
"gesamte": true,
"Bad": true,
"betrug": true,
"Dr.": true,
"besitzt": true,
"dritten": true,
"wenige": true,
"politische": true,
"Sprache": true,
"Zentrum": true,
"Fall": true,
"verwendet": true,
"Grenze": true,
"Tage": true,
"erreicht": true,
"gleichzeitig": true,
"internationalen": true,
"Paris": true,
}