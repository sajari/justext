package gojustext

var BasqueStoplist = map[string]bool {"eta": true,
"da.": true,
"izan": true,
"bat": true,
"da,": true,
"zen": true,
"zuen": true,
"bere": true,
"ez": true,
"da": true,
"zen.": true,
"edo": true,
"ere": true,
"zen,": true,
"egin": true,
"zuen.": true,
"de": true,
"beste": true,
"a.": true,
"zuten": true,
"K.": true,
"ziren": true,
"den": true,
"baina": true,
"du": true,
"dira": true,
"ere,": true,
"biztanle": true,
"dira.": true,
"zuen,": true,
"bi": true,
"zituen.": true,
"egiten": true,
"zituen": true,
"hau": true,
"udalerria": true,
"dagoen": true,
"Bere": true,
"du.": true,
"baino": true,
"duen": true,
"behar": true,
"baten": true,
"eta,": true,
"erroldaren": true,
"urtarrilaren": true,
"lehen": true,
"arabera": true,
"artean": true,
"oso": true,
"sortu": true,
"zuten.": true,
"hasi": true,
"hartu": true,
"dira,": true,
"izena": true,
"ziren.": true,
"eman": true,
"bizi": true,
"1eko": true,
"bezala": true,
"diren": true,
"urte": true,
"egun": true,
"erabiltzen": true,
"du,": true,
"hainbat": true,
"dago.": true,
"Euskal": true,
"dute": true,
"hil": true,
"duten": true,
"arteko": true,
"ditu.": true,
"batzuk": true,
"batean": true,
"lortu": true,
"eskualdean": true,
"2009ko": true,
"honen": true,
"dagoena.": true,
"dago": true,
"arabera,": true,
"ziren,": true,
"dago,": true,
"izeneko": true,
"hauek": true,
"urtean": true,
"ere.": true,
"dela": true,
"ditu": true,
"San": true,
"ondoren,": true,
"zuten,": true,
"azken": true,
"bertan": true,
"hartzen": true,
"gaur": true,
"izango": true,
"hiru": true,
"dute.": true,
"nahi": true,
"arte": true,
"baita": true,
"batez": true,
"zela": true,
"handia": true,
"ematen": true,
"batek": true,
"osatzen": true,
"euskal": true,
"honek": true,
"Espainiako": true,
"Nafarroako": true,
"Hala": true,
"lan": true,
"asko": true,
"egon": true,
"talde": true,
"joan": true,
"esan": true,
"aurkitu": true,
"zenbait": true,
"alde": true,
"batera": true,
"dituen": true,
"bigarren": true,
"hori": true,
"herri": true,
"nagusia": true,
"bezala,": true,
"honetan": true,
"parte": true,
"utzi": true,
"bihurtu": true,
"zituzten": true,
"ondoren": true,
"agertzen": true,
"Estatu": true,
"izaten": true,
"Frantziako": true,
"gehiago": true,
"Kanpo": true,
"kilometro": true,
"loturak.": true,
"erabili": true,
"zituen,": true,
"zion": true,
"bat,": true,
"Gaur": true,
"2008ko": true,
"ohi": true,
"euren": true,
"nahiz": true,
"gain,": true,
"eraiki": true,
"zehar": true,
"ari": true,
"lehenengo": true,
"guztiak": true,
"hain": true,
"dute,": true,
"zati": true,
"hizkuntza": true,
"hamarkada": true,
"zein": true,
"lau": true,
"Historia.": true,
"2007ko": true,
"zegoen": true,
"arte.": true,
"ditu,": true,
"milioi": true,
"sortzen": true,
"sartu": true,
"non": true,
"gutxi": true,
"musika": true,
"hala": true,
"bidez": true,
"izen": true,
"mendean": true,
"daude.": true,
"del": true,
"jarri": true,
"hiri": true,
"aurka": true,
"kokatzen": true,
"bai": true,
"la": true,
"egindako": true,
"azalera": true,
"Beste": true,
"arren,": true,
"Lehen": true,
"urteko": true,
"jaso": true,
"aurrera": true,
"bera": true,
"egiteko": true,
"ahal": true,
"inguruan": true,
"garai": true,
"dena.": true,
"ezagutzen": true,
"berri": true,
"artean,": true,
"mendearen": true,
"egungo": true,
"probintziako": true,
"izendatu": true,
"iritsi": true,
"sistema": true,
"eragin": true,
"inguruko": true,
"daude": true,
"zituzten.": true,
"bakarra": true,
"Herriko": true,
"mende": true,
"hau,": true,
"dauden": true,
"La": true,
"ugari": true,
"errege": true,
"ezin": true,
"taldea": true,
"Gerra": true,
"metroko": true,
"inguru": true,
"Bigarren": true,
"handiena": true,
"pertsona": true,
"ia": true,
"gero": true,
"hiria": true,
"horretan": true,
"mota": true,
"Ipar": true,
"jaio": true,
"aldiz": true,
"aurkako": true,
"buruzko": true,
"bertako": true,
"horren": true,
"hauen": true,
"berriz,": true,
"irabazi": true,
"Bi": true,
"udalerri": true,
"hitz": true,
"batzuetan": true,
"Europako": true,
"metro": true,
"idatzi": true,
"dituzten": true,
"gisa": true,
"handi": true,
"gabe": true,
"batera,": true,
"dituzte": true,
"antzinako": true,
"ibaiaren": true,
"modu": true,
"bakarrik": true,
"of": true,
"mailako": true,
"sortutako": true,
"hegoaldean": true,
"bezala.": true,
"urriaren": true,
"gertatu": true,
"herriak": true,
"daude,": true,
"oraindik": true,
"munduko": true,
"egun,": true,
"besteak": true,
"Azken": true,
"argitaratu": true,
"egoera": true,
"bat.": true,
"hiriburua": true,
"duena.": true,
"I.a": true,
"gainean": true,
"berriz": true,
"XIX.": true,
"ezarri": true,
"geratu": true,
"Maria": true,
"ondorioz,": true,
"Herri": true,
"Mundu": true,
"Baina": true,
"Jose": true,
"ondorioz": true,
"Juan": true,
"jarraitu": true,
"km²-ko": true,
"erabaki": true,
"mendeko": true,
"herrialde": true,
"edota": true,
"dituzte.": true,
"gune": true,
"horiek": true,
"bilakatu": true,
"zuzen": true,
"Erresuma": true,
"daiteke.": true,
"abenduaren": true,
"batzuen": true,
"irailaren": true,
"agertu": true,
"ekainaren": true,
"Batuetako": true,
"aurreko": true,
"uztailaren": true,
"jo": true,
"aita": true,
"amaitu": true,
"herria": true,
"hasieran": true,
"eraman": true,
"denbora": true,
"zena.": true,
"deitzen": true,
"Bilboko": true,
"zuela": true,
"eskualde": true,
"uste": true,
"itzuli": true,
"zion.": true,
"abuztuaren": true,
"the": true,
"iparraldean": true,
"Gainera,": true,
"aritu": true,
"maiatzaren": true,
"The": true,
"estatu": true,
"XX.": true,
"berak": true,
"batetik": true,
"zeuden": true,
"osoan": true,
"haren": true,
"honetako": true,
"bost": true,
"muga": true,
"gerra": true,
"batera.": true,
"txiki": true,
"maila": true,
"zabaldu": true,
"ezaguna": true,
"dira:": true,
"jasotzen": true,
"gertatzen": true,
"lurralde": true,
"II.a": true,
"aurkitutako": true,
"berean,": true,
"taldeak": true,
"martxoaren": true,
"askotan": true,
"batekin": true,
"mendebaldean": true,
"asteroide": true,
"galdu": true,
"Ez": true,
"ekialdean": true,
"atera": true,
"barnean": true,
"zirela": true,
"adierazten": true,
"kokatuta": true,
"jotzen": true,
"Izan": true,
"Santa": true,
"Urte": true,
"nola": true,
"aldatu": true,
"Erdi": true,
"Bizkaiko": true,
"apirilaren": true,
"aukera": true,
"handiak": true,
"esaten": true,
"luze": true,
"moduan": true,
"Ameriketako": true,
"azaroaren": true,
"gai": true,
"bezalako": true,
"batean,": true,
"arte,": true,
"berria": true,
"berean": true,
"Egiptoko": true,
"Hego": true,
"gehien": true,
"Garai": true,
"gero,": true,
"hirian": true,
"osoa": true,
"aldiz,": true,
"berriro": true,
"onartu": true,
"gehienak": true,
"artean.": true,
"mundu": true,
"ikusi": true,
"bereko": true,
"Txinako": true,
"da:": true,
"beraz,": true,
"taldearen": true,
"zioten": true,
"seme": true,
"burua": true,
"Hau": true,
"daiteke": true,
"bitartean": true,
"otsailaren": true,
"zuena.": true,
"honetan,": true,
"hirugarren": true,
"Biografia.": true,
"probintzia": true,
"barruan": true,
"bati": true,
"lanak": true,
"Luis": true,
"El": true,
"eremu": true,
"egitura": true,
"lana": true,
"beti": true,
"behin": true,
"berriak": true,
"zituzten,": true,
"hiriko": true,
"eginiko": true,
"era": true,
"familia": true,
"urtean,": true,
"aurretik": true,
"nagusi": true,
"New": true,
"horrek": true,
"ezaugarri": true,
"auzoa": true,
"departamenduko": true,
"jatorrizko": true,
"semea": true,
"ur": true,
"zenean,": true,
"banatzen": true,
"gabeko": true,
"hamarkadan": true,
"Eusko": true,
"den,": true,
"garatu": true,
"futbol": true,
"hurrengo": true,
"eraikin": true,
"ezkondu": true,
"handiagoa": true,
"kide": true,
"Herrian": true,
"guztiz": true,
"soilik": true,
"hartan": true,
"eliza": true,
"ikus": true,
"ondoren.": true,
"kanpo": true,
"ordea,": true,
"zion,": true,
"balio": true,
"zeharkatzen": true,
"erabat": true,
"dena": true,
"ibaia": true,
"aldeko": true,
"Bertan": true,
"gora": true,
"jokatu": true,
"guztien": true,
"mendebaldeko": true,
"duela": true,
"honek,": true,
"orain": true,
"iparraldeko": true,
"hegoaldeko": true,
"egoitza": true,
"biltzen": true,
"espezie": true,
"koadroko": true,
"gerrikoan": true,
"izateko": true,
"energia": true,
"asteroidearen": true,
"aldi": true,
"herriko": true,
"ekin": true,
"toki": true,
"baina,": true,
"horien": true,
"idazle": true,
"daitezke.": true,
"buru": true,
"armada": true,
"haien": true,
"direla": true,
"y": true,
"jarraitzen": true,
"Nafarroa": true,
"alderdi": true,
"aldaketa": true,
"Hori": true,
"garaian": true,
"John": true,
"edozein": true,
"osatu": true,
"izatea": true,
"A": true,
"bada": true,
"ikasi": true,
"ekarri": true,
"kaleratu": true,
"bide": true,
"2010eko": true,
"enpresa": true,
"handiko": true,
"izena,": true,
"bateko": true,
"dauka": true,
"gertu": true,
"omen": true,
"leku": true,
"Dena": true,
"sartzen": true,
"buruz": true,
"lortzen": true,
"bitartean,": true,
"nagusiak": true,
"jakin": true,
"nahiko": true,
"heldu": true,
"and": true,
"kultura": true,
"Alderdi": true,
"itxura": true,
"ekialdeko": true,
"urtera": true,
"ospatzen": true,
"Ozeano": true,
"aurkeztu": true,
"azaleran.": true,
"arabera.": true,
"erregea": true,
"disko": true,
"alaba": true,
"elementu": true,
"eraso": true,
"\"The": true,
"hasten": true,
"biak": true,
"Antzinako": true,
"ihes": true,
"Irlandako": true,
"Horrela,": true,
"Horrez": true,
"urtetik": true,
"dituzte,": true,
"urteetan": true,
"behera": true,
"eragina": true,
"burutu": true,
"garrantzitsua": true,
"luzera": true,
"pertsonaia": true,
"bizitza": true,
"bitartean.": true,
"zehar,": true,
"beraien": true,
"ostean,": true,
"lur": true,
"Mendebaldeko": true,
"hiriaren": true,
"da;": true,
"daiteke,": true,
"naturala": true,
"sei": true,
"aurrean": true,
"aurre": true,
"garaiko": true,
"bereziki": true,
"erromatar": true,
"osoko": true,
"esker": true,
"ezberdinak": true,
"batzuek": true,
"geroago": true,
"XVIII.": true,
"bidali": true,
"lotzen": true,
"familiako": true,
"Bretainiako": true,
"argi": true,
"anaia": true,
"erakusten": true,
"jokatzen": true,
"Errepublikako": true,
"Nahiz": true,
"laguntza": true,
"izandako": true,
"esku": true,
"pasa": true,
"Karlos": true,
"Halaber,": true,
"probintzian": true,
"igaro": true,
"jartzen": true,
"mendi": true,
"beren": true,
"Inguru": true,
"Aroko": true,
"batetan": true,
"txikia": true,
"XVI.": true,
"zelarik.": true,
"ikusten": true,
"Errege": true,
"gabe.": true,
"arrakasta": true,
"dio": true,
"forma": true,
"iraun": true,
"euskaraz": true,
"Antonio": true,
"zena,": true,
"lagun": true,
"bakoitzak": true,
"egingo": true,
"udalerriko": true,
"eragiten": true,
"Hirugarren": true,
"kilometrora": true,
"2006ko": true,
"hasiera": true,
"besterik": true,
"kopurua": true,
"saiatu": true,
"jasan": true,
"behin-behineko": true,
"barne": true,
"Arabako": true,
"eskaintzen": true,
"ama": true,
"horrela": true,
"eskatu": true,
"uharte": true,
"delarik.": true,
"Gipuzkoako": true,
"ondorengo": true,
"Alemaniako": true,
"denez,": true,
"egitea": true,
"tenperatura": true,
"azaltzen": true,
"emandako": true,
"eskualdean.": true,
"mantendu": true,
"Iruñeko": true,
"gain": true,
"gainera,": true,
"Errepublika": true,
"Hurrengo": true,
"arren.": true,
"aurrera,": true,
"ezagutu": true,
"izendapen": true,
"bitartez": true,
"bete": true,
"ondoan": true,
"aurkitzen": true,
"aurki": true,
"izan,": true,
"errepidea": true,
"bestelako": true,
"pinyineraz:": true,
"hauek,": true,
"ondo": true,
"(txinera": true,
"(K.": true,
"kasu": true,
"Ekialdeko": true,
"deitu": true,
"Hiru": true,
"Aita": true,
"askoz": true,
"Lehenengo": true,
"departamenduan": true,
"erdialdean": true,
"izan.": true,
"jatorria": true,
"eskualdeko": true,
"diskoa": true,
"2008an": true,
"ibili": true,
"batzuk,": true,
"III.a": true,
"Europa": true,
"indar": true,
"nazioarteko": true,
"zioten.": true,
"eta.": true,
"iristen": true,
"erregearen": true,
"ditu:": true,
"antzeko": true,
"zegoen,": true,
"udal": true,
"gobernuak": true,
"Olinpiar": true,
"Pedro": true,
"II.": true,
"erabilera": true,
"dauzka": true,
"taldeko": true,
"zazpi": true,
"grabatu": true,
"Munduko": true,
"osatutako": true,
"bidea": true,
"lehendabiziko": true,
"Finlandiako": true,
"idatzitako": true,
"historia": true,
"zegoen.": true,
"aktore": true,
"sinplifikatuz:": true,
"eskualdea": true,
"agindu": true,
"guztia": true,
"dena,": true,
"bada,": true,
"horretan,": true,
"honi": true,
"mendia": true,
"erakunde": true,
"banaketa": true,
"izatera": true,
"Donostiako": true,
"askok": true,
"garrantzitsuena": true,
"bota": true,
"gabe,": true,
"Honek": true,
"taldean": true,
"ezagunak": true,
"igarotzen": true,
"hamar": true,
"ostean": true,
"dagokionez,": true,
"politiko": true,
"XVII.": true,
"batean.": true,
"Nazioarteko": true,
"lanean": true,
"eskuratu": true,
"geratzen": true,
"klima": true,
"bildu": true,
"aipatzen": true,
"borroka": true,
"dituena.": true,
"garaitu": true,
"daitezke": true,
"benetako": true,
"abesti": true,
"gainontzeko": true,
"Italiako": true,
"Joan": true,
"batu": true,
"eduki": true,
"desagertu": true,
"euskara": true,
"esker,": true,
"gaineko": true,
"kopuru": true,
"dauka.": true,
"antolatu": true,
"solairu": true,
"prozesu": true,
"boto": true,
"Honen": true,
"doan": true,
"mugimendu": true,
"puntu": true,
"sinatu": true,
"udalerriaren": true,
"kontuan": true,
"Ezaugarriak.": true,
"igo": true,
"inoiz": true,
"B": true,
"borrokatu": true,
"zenean": true,
"udalerrian": true,
"Eta": true,
"lotura": true,
"datu": true,
"botere": true,
"tamaina": true,
"2005ean": true,
"orduan": true,
"Guztira": true,
"eraikuntza": true,
"beharko": true,
"boterea": true,
"herriaren": true,
"greziar": true,
"ongi": true,
"antolatzen": true,
"Erromatar": true,
"beste,": true,
"urtez": true,
"kolore": true,
"I.": true,
"eskola": true,
"liburu": true,
"film": true,
"aldean": true,
"gizarte": true,
"altuera": true,
"genero": true,
"horietako": true,
"kanpoko": true,
"Iparraldeko": true,
"etorri": true,
"2005eko": true,
"eratzen": true,
"irudikatzen": true,
"honako": true,
"guztiek": true,
"a": true,
"ohiko": true,
"zuen:": true,
"izaera": true,
"bertan,": true,
"uhartea": true,
"animalia": true,
"Europan": true,
"mantentzen": true,
"Martin": true,
"dugu": true,
"arazo": true,
"bakar": true,
"2003ko": true,
"etxe": true,
"ibaiak": true,
"Urteko": true,
"batek,": true,
"auzo": true,
"Era": true,
"munduan": true,
"betetzen": true,
"amaieran": true,
"Adibidez,": true,
"bihurtzen": true,
"estatuko": true,
"taula": true,
"ekintza": true,
"izateko.": true,
"bazen": true,
"beste.": true,
"Bestalde,": true,
"III.": true,
"handik": true,
"zientifikoa": true,
"garrantzi": true,
"iparraldean,": true,
"multzo": true,
"hiriburu": true,
"ageri": true,
"bukatu": true,
"independentzia": true,
"erregeak": true,
"doa.": true,
"konkistatu": true,
"Hainbat": true,
"Eguzkiaren": true,
"Gipuzkoa,": true,
"etxebizitza": true,
"hauetan": true,
"Miguel": true,
"lortzeko": true,
"herrian": true,
"gelditu": true,
"helburu": true,
"adibidez,": true,
"tartean": true,
"zenbaki": true,
"izanik.": true,
"han": true,
"gure": true,
"zien": true,
"erresuma": true,
"\"La": true,
"kontrako": true,
"Erromako": true,
"gobernatu": true,
"batetik,": true,
"lagundu": true,
"nafar": true,
"ofiziala": true,
"Egun": true,
"bertan.": true,
"eraikitako": true,
"zuena,": true,
"tresna": true,
"aipatu": true,
"IV.": true,
"teoria": true,
"Ingalaterrako": true,
"beharra": true,
"ura": true,
"bitarteko": true,
"zitzaion": true,
"filma": true,
"Mikel": true,
"txikiak": true,
"auzoak": true,
"nola,": true,
"herrialdeko": true,
"ospatu": true,
"lehenago": true,
"jarduera": true,
"Azkenik,": true,
"zen:": true,
"mailan": true,
"kontra": true,
"izenak": true,
"Fernando": true,
"Horren": true,
"duena": true,
"duena,": true,
"adierazi": true,
"hitza": true,
"uzten": true,
"lotuta": true,
"Erdialdeko": true,
"hasierako": true,
"itsas": true,
"ireki": true,
"ibai": true,
"amaiera": true,
"Gasteizko": true,
"estatua": true,
"gauza": true,
"kokapena.": true,
"Oso": true,
"biztanleria": true,
"gizon": true,
"izar": true,
"Bartzelonako": true,
"harreman": true,
"aukeratu": true,
"baizik": true,
"dauka,": true,
"sarrera": true,
"Los": true,
"batzuetan,": true,
"zena": true,
"Foru": true,
"kidea": true,
"bereganatu": true,
"estilo": true,
"1980ko": true,
"Eliza": true,
"garrantzitsu": true,
"zen;": true,
"zortzi": true,
"garrantzia": true,
"Ondoren,": true,
"giza": true,
"Inperioaren": true,
"eratu": true,
"kendu": true,
"batik": true,
"arduratzen": true,
"jada": true,
"sortzeko": true,
"hartako": true,
"mendean,": true,
"hilabete": true,
"2004ko": true,
"zituela": true,
"gazte": true,
"neurri": true,
"kontatzen": true,
"oinarritzen": true,
"alegia.": true,
"Bertan,": true,
"denean": true,
"falta": true,
"Beraz,": true,
"iturri": true,
"Honela,": true,
"guztietan": true,
"zuzendutako": true,
"irudi": true,
"bada.": true,
"zaio.": true,
"urtero": true,
"oinarrizko": true,
"Gerran": true,
"baita.": true,
"erdi": true,
"horretako": true,
"goi": true,
"zelula": true,
"hedatzen": true,
"Gaztelako": true,
"José": true,
"V.": true,
"osatua": true,
"emakume": true,
"ezik,": true,
"ekialdean,": true,
"kirol": true,
"Manuel": true,
"Jean": true,
"Inperio": true,
"Itsaso": true,
"zehar.": true,
"biztanleriaren": true,
"2003an": true,
"Geografia.": true,
"eraikitzen": true,
"diseinatu": true,
"zutela": true,
"barne.": true,
"politika": true,
"titulua": true,
"arazoak": true,
"hamarkadaren": true,
"bezain": true,
"informazio": true,
"zuzenean": true,
"saria": true,
"eguna": true,
"James": true,
"aldarrikatu": true,
"gehienetan": true,
"banatu": true,
"erreferentzia": true,
"Jaiak": true,
"el": true,
"urtea": true,
"historian": true,
"datuen": true,
"arkitekto": true,
"bultzatu": true,
"bertsio": true,
"gainera.": true,
"horrela,": true,
"bidez,": true,
"idazten": true,
"C": true,
"berezko": true,
"dio.": true,
"Madrilgo": true,
"gorputz": true,
"beraz": true,
"helburua": true,
"AEBetako": true,
"Xabier": true,
"Barrutia": true,
"Hegoaldeko": true,
"emaitza": true,
"km": true,
"daitezke,": true,
"azaldu": true,
"militar": true,
"zinema": true,
"2007an": true,
"baitzen.": true,
"ekonomiko": true,
"aldetik": true,
"abiadura": true,
"zela,": true,
"Amerikako": true,
"horregatik": true,
"lurraldea": true,
"Robert": true,
"garaian,": true,
"eskaini": true,
"berehala": true,
"Afrikako": true,
"erdian": true,
"ateratzen": true,
"ezaguna,": true,
"irakasle": true,
"jainko": true,
"garaiera": true,
"ikerketa": true,
"ekonomia": true,
"ikasketak": true,
"indarra": true,
"zuelarik.": true,
"material": true,
"kokatua.": true,
"errepidearen": true,
"egiteko.": true,
"gobernu": true,
"Antso": true,
"erakutsi": true,
"laugarren": true,
"dinastiako": true,
"har": true,
"kilometroan": true,
"Sierra": true,
"sare": true,
"zuzendari": true,
"mendiaren": true,
"linea": true,
"urtean.": true,
"(edo": true,
"informazioa": true,
"beharreko": true,
"Errepidea": true,
"izenez": true,
"beranduago": true,
"besteko": true,
"bidez.": true,
"gutxiago": true,
"zaio": true,
"hegoaldean,": true,
"aldatzen": true,
"atal": true,
"eraikina": true,
"Herria": true,
"bata": true,
"batzuekin": true,
"Iparraldean": true,
"nabarmentzen": true,
"hitzak": true,
"hasieran,": true,
"askoren": true,
"en": true,
"ona": true,
"Europar": true,
"zuenean,": true,
"eskatzen": true,
"Francisco": true,
"liburua": true,
"Saria": true,
"bereizten": true,
"suntsitu": true,
"zuena": true,
"Gerraren": true,
"zela.": true,
"kokatutako": true,
"politikoa": true,
"kokatu": true,
"Parisko": true,
"alkatea": true,
"ofizialki": true,
"badira": true,
"hedatu": true,
"osatuta": true,
"los": true,
"Baina,": true,
"hura": true,
"Iberiar": true,
"jarrera": true,
"errepidea,": true,
"ideia": true,
"\"El": true,
"denean,": true,
"aritzen": true,
"masa": true,
"eredu": true,
"txapelketa": true,
"mendebaldean,": true,
"legez": true,
"kilometroko": true,
"landare": true,
"hiriak": true,
"in": true,
"Arte": true,
"den.": true,
"datuak": true,
"proposatu": true,
"joaten": true,
"beranduago,": true,
"normalean": true,
"bakoitza": true,
"teknika": true,
"arkitektura": true,
"beheko": true,
"ondorioz.": true,
"erori": true,
"William": true,
"auzoan": true,
"Alde": true,
"parke": true,
"XIII.": true,
"erdiko": true,
"zuzendu": true,
"I": true,
"2001eko": true,
"inguruan.": true,
"itsasoaren": true,
"bizirik": true,
"batzen": true,
"esaterako,": true,
"direnak.": true,
"hazi": true,
"2006an": true,
"Alfontso": true,
"asmatu": true,
"izaki": true,
"sortuz.": true,
"heriotza": true,
"berreskuratu": true,
"Pirinio": true,
"joera": true,
"Hiri": true,
"astronomo": true,
"honela": true,
"berezia": true,
"mendetik": true,
"inguruan,": true,
"Oscar": true,
"XV.": true,
"Hasiera": true,
"Haren": true,
"azkenean": true,
"II.aren": true,
"kokaturik": true,
"ondorio": true,
"milaka": true,
"egongo": true,
"eraikitzeko": true,
"Handia": true,
"gaztelua": true,
"jokalari": true,
"goiko": true,
"zerbitzu": true,
"preso": true,
"bertatik": true,
"industria": true,
"rock": true,
"altuena": true,
"gehitu": true,
"oinarritutako": true,
"Demografia.": true,
"atxilotu": true,
"desberdinak": true,
"ohikoa": true,
"hegazti": true,
"geroago,": true,
"ziren:": true,
"Julio": true,
"Eskoziako": true,
"aztertzen": true,
"honekin": true,
"Lan": true,
"Lau": true,
"hauetako": true,
"kostaldean": true,
"motako": true,
"George": true,
"zegoela": true,
"I.aren": true,
"gisa.": true,
"aldetik,": true,
"abiatu": true,
"sorturiko": true,
"David": true,
"urteen": true,
"probintziaren": true,
"Parke": true,
"saldu": true,
"Aldi": true,
"m)": true,
"aintzira": true,
"hiririk": true,
"uko": true,
"dela,": true,
"Javier": true,
"Burgosko": true,
"Iñaki": true,
"Batasunaren": true,
"ibilbide": true,
"hartzeko": true,
"garapen": true,
"onartzen": true,
"Frantzia": true,
"(bretainieraz": true,
"Errusiako": true,
"Jon": true,
"presidentea": true,
"zerrenda": true,
"hurbil": true,
"nabarmendu": true,
"barneko": true,
"Bilbo": true,
"Real": true,
"Goi": true,
"adibide": true,
"1990eko": true,
"lanetan": true,
"ordu": true,
"emaztea": true,
"garatzen": true,
"ekoizten": true,
"txikiagoa": true,
"handitu": true,
"esker.": true,
"Carlos": true,
"hartan,": true,
"XII.": true,
"ofizial": true,
"hauteskundeetan": true,
"funtzio": true,
"lurrak": true,
"fatxada": true,
"zeuden.": true,
"burutzen": true,
"Handik": true,
"Zenbait": true,
"erabiliz": true,
"aurretik,": true,
"barrutietako": true,
"artikulu": true,
"Tierra": true,
"erosi": true,
"jatorri": true,
"mugak": true,
"euskarazko": true,
"dagokion": true,
"Ikus,": true,
"guztira": true,
"banatuta": true,
"eramaten": true,
"ezartzen": true,
"Charles": true,
"Aroan": true,
"izena.": true,
"zer": true,
"Hauek": true,
"oro": true,
"kokatua": true,
"kontzejua": true,
"1960ko": true,
"espainiar": true,
"idazlea": true,
"daitekeen": true,
"azpian": true,
"berarekin": true,
"telebista": true,
"ahalbidetzen": true,
"izenarekin": true,
"Aragoiko": true,
"diska": true,
"fosil": true,
"mende)": true,
"inguratzen": true,
"horietatik": true,
"zituena.": true,
"hazten": true,
"itxi": true,
"oinarri": true,
"ondoko": true,
"Mari": true,
"dugu.": true,
"erabilitako": true,
"nagusia,": true,
"iraupena": true,
"1970eko": true,
"(euskaraz": true,
"daitezkeen": true,
"buruzagi": true,
"une": true,
"nabarmena": true,
"bien": true,
"berezi": true,
"Hiriburua": true,
"hildako": true,
"aztertu": true,
"programa": true,
"Kataluniako": true,
"menpe": true,
"Britainia": true,
"Talde": true,
"azkenik,": true,
"unitate": true,
"kantu": true,
"gainera": true,
"irten": true,
"ustez,": true,
"gordetzen": true,
"direlarik.": true,
"zatia": true,
"arraza": true,
"Dragoi": true,
"bila": true,
"jende": true,
"ehun": true,
"adibidez": true,
"eguzki": true,
"orduko": true,
"Futbol": true,
"tokia": true,
"ezberdinetan": true,
"Batuko": true,
"zuzendaria": true,
"herrixka": true,
"dutenak.": true,
"dio,": true,
"emateko": true,
"Guztira,": true,
"Aro": true,
"erabil": true,
"botoak": true,
"diru": true,
"politikari": true,
"i": true,
"Valentziako": true,
"ala": true,
"abeslari": true,
"urterekin": true,
"aldaera": true,
"heltzen": true,
"zituen:": true,
"langile": true,
"ezberdin": true,
"Kultura": true,
"hegoaldean.": true,
"inongo": true,
"duelarik.": true,
"presidente": true,
"piztu": true,
"lurraldeak": true,
"hortik": true,
"gehienek": true,
"bitan": true,
"zubia": true,
"Herbeheretako": true,
"laster": true,
"Marko": true,
"izanik,": true,
"ikasle": true,
"nagusia.": true,
"honakoak": true,
"ospetsua": true,
"aldera": true,
"zioten,": true,
"Japoniako": true,
"zuen;": true,
"eginez.": true,
"kokatutakoa.": true,
"Rhône-Alpeak": true,
"Inperioko": true,
"pasatzen": true,
"Ain": true,
"M": true,
"osagai": true,
"planeta": true,
"zenbat": true,
"1ean": true,
"matxinada": true,
"Jesus": true,
"Erkidegoko": true,
"molekula": true,
"arlo": true,
"egoten": true,
"zabalera": true,
"zaila": true,
"urrezko": true,
"bandera": true,
"barruti": true,
"zenean.": true,
"deritzo.": true,
"Bertako": true,
"sekula": true,
"prozesua": true,
"arku": true,
"baitzen": true,
"horretarako": true,
"Zaragozako": true,
"proiektu": true,
"ipar-ekialdean": true,
"ondorengoa": true,
"alemaniarrak": true,
"omenez": true,
"bidezko": true,
"gogor": true,
"beharrezkoa": true,
"bestea": true,
"ikur": true,
"dagoena,": true,
"elkartzen": true,
"lortutako": true,
"objektu": true,
"elkartu": true,
"etxeorratz": true,
"nekazaritza": true,
"eragindako": true,
"Unibertsitatean": true,
"zabaltzen": true,
"Inperioa": true,
"du:": true,
"asmoz.": true,
"ordena": true,
"metal": true,
"(bretoieraz": true,
"arrazoi": true,
"behatokitik.": true,
"guduan": true,
"aitaren": true,
"Sistema": true,
"irlandar": true,
"gudua": true,
"erraz": true,
"behera.": true,
"mitologian,": true,
"Erreferentziak.": true,
"mugitzen": true,
"dagoela": true,
"dauzka.": true,
"bazuen": true,
"Horregatik,": true,
"bakoitzaren": true,
"Autonomia": true,
"Gainera": true,
"merkataritza": true,
"Bestela,": true,
"zabal": true,
"baldintza": true,
"ezaugarriak": true,
"arren": true,
"delako": true,
"Salamancako": true,
"frantziar": true,
"literatura": true,
"horietan": true,
"ibilbidea": true,
"Behin": true,
"hautatu": true,
"ordezkatu": true,
"erlijio": true,
"abestia": true,
"bilduma": true,
"Han": true,
"argia": true,
"musikari": true,
"zuri": true,
"egutegi": true,
"ipar-mendebaldean": true,
"bulego": true,
"Herriaren": true,
"zutabe": true,
"garrantzitsuak": true,
"bakoitzean": true,
"kale": true,
"multzoa": true,
"mitologiako": true,
"ezer": true,
"autonomia": true,
"kontrola": true,
"moduan,": true,
"zeharkatu": true,
"daude:": true,
"L": true,
"O": true,
"Lurraren": true,
"Frantzian": true,
"Karl": true,
"Horrela": true,
"gisa,": true,
"ustez": true,
"dira;": true,
"harri": true,
"Txapelketa": true,
"mendeetan": true,
"dinastia": true,
"Filipe": true,
"zefalopodoen": true,
"Galiziako": true,
"baterako": true,
"geltokia": true,
"desegin": true,
"ammonite": true,
"luzea": true,
"hau.": true,
"mila": true,
"tokiko": true,
"gaixotasun": true,
"Asia": true,
"publikoa": true,
"urak": true,
"egunean": true,
"soldadu": true,
"tenplu": true,
"Sobietar": true,
"María": true,
"von": true,
"uharteak": true,
"orduan,": true,
"euskararen": true,
"partida": true,
"gaztelaniaz": true,
"Ekonomia.": true,
"ipar": true,
"S": true,
"izendatzeko": true,
"atzera": true,
"Atlantikoak": true,
"ate": true,
"ezta": true,
"Indiako": true,
"enpresak": true,
"distantzia": true,
"balioa": true,
"lege": true,
"Paul": true,
"batzuk.": true,
"zientzia": true,
"alegia,": true,
"existitzen": true,
"moduan.": true,
"herrietako": true,
"dator,": true,
"baitzuen.": true,
"harrizko": true,
"bederatzi": true,
"dutela": true,
"testu": true,
"Nolanahi": true,
"ertzean": true,
"zitzaion.": true,
"gorde": true,
"zuhaitz": true,
"XIV.": true,
"jardun": true,
"bukaeran": true,
"Henrike": true,
"babesa": true,
"bidaia": true,
"Portugalgo": true,
"debekatu": true,
"bidean": true,
"Louis": true,
"zuenean": true,
"su": true,
"Behe": true,
"horri": true,
"jasaten": true,
"IV.a": true,
"harremanak": true,
"Udal": true,
"amaitzen": true,
"etxea": true,
"zabala": true,
"landu": true,
"dela.": true,
"zeuden,": true,
"Galesko": true,
"mendilerroan": true,
"baldin": true,
"ezinezkoa": true,
"hamarkada.": true,
"gatazka": true,
"Ama": true,
"zehatz": true,
"arma": true,
"gobernuaren": true,
"etxean": true,
"herrian,": true,
"Mexikoko": true,
"lehena": true,
"estatubatuarra": true,
"barrutian": true,
"okupatu": true,
"egina": true,
"gas": true,
"emanez.": true,
"jauregia": true,
"gero.": true,
"gainean,": true,
"zerrendan": true,
"Bizkaia,": true,
"Son": true,
"diote": true,
"guzti": true,
"Euren": true,
"zuzena": true,
"gertu.": true,
"erabiliz.": true,
"aldaketak": true,
"berdina": true,
"karratuko": true,
"Orduan": true,
"sartzeko": true,
"badaude": true,
"zeukan": true,
"plataforma": true,
"ospea": true,
"edizioa": true,
"baso": true,
"egun.": true,
"amaieran,": true,
"Izen": true,
"presio": true,
"jainkosa": true,
"zuenean.": true,
"azkar": true,
"isurtzen": true,
"hezkuntza": true,
"Klima": true,
"Espainian": true,
"nagusiak.": true,
"Musika": true,
"egutegiaren": true,
"ohitura": true,
"Sant": true,
"urteetan,": true,
"irudia": true,
"Izena": true,
"lehen,": true,
"adar": true,
"gainditu": true,
"zuzen.": true,
"hego-mendebaldean": true,
"hiltzen": true,
"partikula": true,
"elkarrekin": true,
"Andre": true,
"zenez,": true,
"Asiako": true,
"erantzun": true,
"honetatik": true,
"izendatzen": true,
"2002ko": true,
"elkarte": true,
"gainetik": true,
"batetan,": true,
"proiektua": true,
"erdialdeko": true,
"ezagun": true,
"aurka.": true,
"espero": true,
"uhartean": true,
"bana": true,
"etxeko": true,
"bira": true,
"Handiaren": true,
"soinu": true,
"ezagunena": true,
"zurezko": true,
"Prefektura": true,
"enperadore": true,
"ekonomikoa": true,
"non,": true,
"gainerako": true,
"inolako": true,
"ohia": true,
"hartuko": true,
"galtzen": true,
"York": true,
"metodo": true,
"garaipen": true,
"Hasieran": true,
"emazte": true,
"Alemania": true,
"aipatutako": true,
"kasuan": true,
"polizia": true,
"azkenik": true,
"honela,": true,
"mendirik": true,
"VI.": true,
"biztanleak": true,
"Espainia": true,
"2009an": true,
"ikuspegi": true,
"lotutako": true,
"historiako": true,
"Hiria": true,
"G": true,
"Tolkienen": true,
"zaharra": true,
"emango": true,
"gizakiak": true,
"Afrika": true,
"garaia": true,
"geroztik": true,
"igotzen": true,
"mugan": true,
"dator": true,
"egiteko,": true,
"on": true,
"esate": true,
"armadak": true,
"erregina": true,
"km²": true,
"badu": true,
"babestu": true,
"a>": true,
"urriak": true,
"pasatu": true,
"produktu": true,
"gregoriar": true,
"Richard": true,
"erabiltzeko": true,
"dion": true,
"ezberdinen": true,
"eraikia": true,
"lekua": true,
"kanpaina": true,
"2000ko": true,
"kontzertu": true,
"babesten": true,
"Italia)": true,
"Nobel": true,
"zien.": true,
"lotu": true,
"Pierre": true,
"esaterako.": true,
"herrialdea": true,
"Jaurlaritzak": true,
"sortua": true,
"leiho": true,
"Errusiar": true,
"garapena": true,
"Finlandia": true,
"Côtes-d'Armor": true,
"ordaindu": true,
"moduko": true,
"zelarik,": true,
"zizkion": true,
"laguntzen": true,
"arrunta": true,
"Donibane": true,
"beraren": true,
"konderria": true,
"martxan": true,
"(Sardinia,": true,
"taldearekin": true,
"bertara": true,
"propioa": true,
"azido": true,
"Azkenean,": true,
"Hain": true,
"naturala.": true,
"eskubideak": true,
"baitan": true,
"2004an": true,
"gunea": true,
"bizia": true,
"Gero": true,
"dugu,": true,
"alemaniar": true,
"eragile": true,
"Hizkuntza": true,
"berriaren": true,
"muino": true,
"badago": true,
"Santiago": true,
"ardura": true,
"kontutan": true,
"zirelarik.": true,
"asmoa": true,
"lerro": true,
"Nazionalak": true,
"aurrean,": true,
"harrapatu": true,
"bisurteetan.": true,
"antzekoa": true,
"sorrera": true,
"elkarren": true,
"Joanes": true,
"Goku": true,
"betiko": true,
"gehiena": true,
"Ondorioz,": true,
"hurrenez": true,
"ikasten": true,
"kate": true,
"Euskararen": true,
"Mendi": true,
"plaza": true,
"hartzea": true,
"paper": true,
"direnak": true,
"natura": true,
"semea,": true,
"ikurra": true,
"zuzen,": true,
"jabe": true,
"gobernua": true,
"2001ean": true,
"heriotzaren": true,
"probintzia.": true,
"urteurrenak.": true,
"izenaz": true,
"makina": true,
"liburuan": true,
"aktorea": true,
"Joseph": true,
"arloan": true,
"urtearen": true,
"apaiz": true,
"garaiena": true,
"gela": true,
"Kasu": true,
"zerbitzuak": true,
"irudiak": true,
"legeak": true,
"erreka": true,
"dantza": true,
"baitzuen": true,
"hego-ekialdean": true,
"hegoaldera": true,
"Armada": true,
"konposatu": true,
"mugakideak": true,
"posible": true,
"elizaren": true,
"kudeatzen": true,
"geruza": true,
"Parisen": true,
"populatuena": true,
"Le": true,
"fikziozko": true,
"egiten.": true,
"gaitasuna": true,
"pertsonak": true,
"hogeita": true,
"Espainiar": true,
"et": true,
"ospetsu": true,
"hilobi": true,
"mendiak": true,
"izateaz": true,
"(txineraz:": true,
"prefektura": true,
"(gaztelaniaz": true,
"portu": true,
"Michael": true,
"domeinua": true,
"Batuetan": true,
"irauten": true,
"auzoko": true,
"deitua": true,
"Ramon": true,
"kristau": true,
"udalerriak": true,
"barruko": true,
"azalera.": true,
"zaio,": true,
"abiatzen": true,
"herritarren": true,
"generoko": true,
"P": true,
"D": true,
"Urrezko": true,
"argitaratzen": true,
"deskribatzen": true,
"aurretik.": true,
"familiaren": true,
"borrokan": true,
"finkatu": true,
"tren": true,
"armadaren": true,
"erabilia": true,
"sari": true,
"hots,": true,
"zatitan": true,
"Iruñea": true,
"Max": true,
"izanik": true,
"Jatorrizko": true,
"erakundeak": true,
"dago:": true,
"nabarmen": true,
"mailatik": true,
"mugatzen": true,
"Egungo": true,
"hauteskundeak.": true,
"(Gipuzkoa)": true,
"zubi": true,
"enperadorearen": true,
"Kirol": true,
"tokian": true,
"hiritik": true,
"garrantzitsuenak": true,
"ordezkari": true,
"amaitzeko.": true,
"eraikinaren": true,
"Hong": true,
"lurra": true,
"kostaldeko": true,
"estatuaren": true,
"are": true,
"Liga": true,
"Las": true,
"oinordeko": true,
"atzean": true,
"guztian": true,
"txinera": true,
"ostean.": true,
"eztabaida": true,
"Hortik": true,
"zirenak.": true,
"gertu,": true,
"bakoitzeko": true,
"ontzi": true,
"Batuen": true,
"izendatua": true,
"saiakera": true,
"du;": true,
"Batasuneko": true,
"publiko": true,
"ditugu.": true,
"teknologia": true,
"gorri": true,
"hemen": true,
"aipaturiko": true,
"Wilhelm": true,
"Greziar": true,
"ezarritako": true,
"uraren": true,
"hori,": true,
"dute:": true,
"oinordekoa": true,
"jaitsi": true,
"ordez": true,
"zuzeneko": true,
"azterketa": true,
"Orain": true,
"hartuz.": true,
"kimiko": true,
"ezberdina": true,
"Nazio": true,
"kokatua,": true,
"Mediterraneo": true,
"etab.": true,
"Itsas": true,
"zerrenda.": true,
"nahastu": true,
"hiriburutik": true,
"deritzon": true,
"Egipto": true,
"to": true,
"horrekin": true,
"espazio": true,
"dorre": true,
"Saint": true,
"zaharrena": true,
"zuten:": true,
"Bizitza.": true,
"joko": true,
"Lurralde": true,
"urtetan": true,
"uretan": true,
"idazlea.": true,
"jabetza": true,
"erdia": true,
"bigarrena": true,
"kasu.": true,
"joateko": true,
"Jokoetan": true,
"britainiar": true,
"frogatu": true,
"geltoki": true,
"erromatarren": true,
"izenaren": true,
"Herrian,": true,
"estaltzen": true,
"jaten": true,
"Kretazeoan": true,
"lanaren": true,
"eskuetan": true,
"administratiboa.": true,
"Bilbon": true,
"eginez": true,
"lehenbiziko": true,
"labur": true,
"dugun": true,
"zuela.": true,
"berez": true,
"letra": true,
"eremua": true,
"jaiak": true,
"Batuak": true,
"argitaratutako": true,
"Honek,": true,
"Ezker": true,
"(grezieraz": true,
"eskubidea": true,
"mendebaldean.": true,
"kideak": true,
"bai.": true,
"jendea": true,
"guztietako": true,
"15ean": true,
"jarritako": true,
"kalte": true,
"itsasoan": true,
"Nagusia": true,
"herrien": true,
"errepideko": true,
"papera": true,
"Denbora": true,
"atzetik": true,
"haur": true,
"Esan": true,
"web": true,
"bitartez.": true,
"Pablo": true,
"Henry": true,
"korronte": true,
"mugakide": true,
"Orduan,": true,
"Yorkeko": true,
"des": true,
"hogei": true,
"litzateke.": true,
"nagusiki": true,
"Ramses": true,
"taldeen": true,
"har,": true,
"for": true,
"Horretarako,": true,
"Sari": true,
"zirelako.": true,
"zaldi": true,
"ezik": true,
"kontrolatzen": true,
"ezaguna.": true,
"VIII.": true,
"nagusiaren": true,
"Bilbao": true,
"hainbeste": true,
"beltz": true,
"erraldoi": true,
"begira": true,
"militarra": true,
"eskubide": true,
"Herrian.": true,
"funtsezko": true,
"(ingelesez": true,
"banandu": true,
"buruan": true,
"tarte": true,
"Mugakideak.": true,
"neurtzen": true,
"omenezko": true,
"gutxira": true,
"uharteko": true,
"Legea": true,
"agindua": true,
"gerran": true,
"onak": true,
"Horrek": true,
"izatea.": true,
"izenekoa.": true,
"erabateko": true,
"Nazionala": true,
"XI.": true,
"Errioxako": true,
"bestalde,": true,
"gerora": true,
"sarritan": true,
"muturrean": true,
"irakaslea": true,
"erresistentzia": true,
"enperadorea": true,
"hamaika": true,
"artista": true,
"horma": true,
"ondoan.": true,
"errepide": true,
"herritar": true,
"handirik": true,
"inbaditu": true,
"Egoera": true,
"handien": true,
"herrialdearen": true,
"topatu": true,
"gertaera": true,
"elkar": true,
"latinez": true,
"diren.": true,
"(eta": true,
"idazkari": true,
"Baionako": true,
"etengabe": true,
"dagoena": true,
"at": true,
"Joseba": true,
"no": true,
"zelako.": true,
"Kristo": true,
"pisua": true,
"berari": true,
"Britainiar": true,
"agintari": true,
"itsasoa": true,
"(San": true,
"askatu": true,
"jainkoaren": true,
"pentsatu": true,
"baterako,": true,
"lakua": true,
"gutxienez": true,
"eraikiko": true,
"sortzailea": true,
"II.ak": true,
"Toledoko": true,
"Guadalajarako": true,
"obra": true,
"Bizkaia": true,
"III.aren": true,
"aurkezten": true,
"barnean,": true,
"V.a": true,
"geroztik.": true,
"zirenean,": true,
"etorkin": true,
"eusten": true,
"25ean": true,
"honetan.": true,
"ordez,": true,
"egokia": true,
"lanik": true,
"maiz": true,
"(Bizkaia)": true,
"gainean.": true,
"Bola": true,
"bosgarren": true,
"inauguratu": true,
"(jatorrizko": true,
"nazionala": true,
"prestatu": true,
"historikoa": true,
"gehiengoa": true,
"edo,": true,
"R": true,
"orokorra": true,
"las": true,
"arroka": true,
"murriztu": true,
"bereziak": true,
"porrot": true,
"orokorrean": true,
"onena": true,
"gizakien": true,
"karbono": true,
"harremana": true,
"batekin.": true,
"saiatzen": true,
"hamarkadak": true,
"ospe": true,
}