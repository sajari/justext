package gojustext

var NeapolitanStoplist = map[string]bool {"'e": true,
"è": true,
"nu": true,
"pruvincia": true,
"d''a": true,
"comune": true,
"e": true,
"'o": true,
"a": true,
"'a": true,
"abitante": true,
"crestiane": true,
"ca": true,
"d'": true,
"di": true,
"e'": true,
"se": true,
"cu": true,
"cchiù": true,
"San": true,
"na": true,
"'O": true,
"in": true,
"d''o": true,
"d\"a": true,
"pe": true,
"o": true,
"pure": true,
"'A": true,
"o'": true,
"a'": true,
"fuje": true,
"da": true,
"ô": true,
"comme": true,
"nun": true,
"â": true,
"si": true,
"’e": true,
"de": true,
"del": true,
"la": true,
"d''e": true,
"'E": true,
"che": true,
"pe'": true,
"ma": true,
"Napule": true,
"ll'anno": true,
"fine": true,
"ha": true,
"parte": true,
"ca'": true,
"juorno": true,
"’a": true,
"primma": true,
"l'anne": true,
"p''a": true,
"ppe": true,
"('o": true,
"secunno": true,
"calannario": true,
"journe": true,
"so'": true,
"Napule,": true,
"bisestile).": true,
"Ammancano": true,
"canzone": true,
"fa": true,
"'nu": true,
"era": true,
"ra": true,
"stato": true,
"'na": true,
"tutte": true,
"greguriano": true,
"Torino.": true,
"ll'anno.": true,
"vene": true,
"storia": true,
"cu'": true,
"tutto": true,
"il": true,
"ra'": true,
"nu'": true,
"chi": true,
"dei": true,
"A": true,
"int''e": true,
"Napule.": true,
"ancora": true,
"una": true,
"granne": true,
"sta": true,
"ro'": true,
"belli": true,
"ed": true,
"nomme": true,
"della": true,
"Cuneo.": true,
"Bergamo.": true,
"assaje": true,
"de'": true,
"canzona": true,
"ce": true,
"tra": true,
"songo": true,
"cità": true,
"fatto": true,
"Trento.": true,
"'n": true,
"tene": true,
"addò": true,
"i": true,
"into": true,
"na'": true,
"’o": true,
"Santa": true,
"nce": true,
"quanno": true,
"con": true,
"do": true,
"dinte": true,
"Brescia.": true,
"al": true,
"un": true,
"'int'ô": true,
"dint'": true,
"d": true,
"Milano.": true,
"the": true,
"p'": true,
"ad": true,
"doppo": true,
"l'ha": true,
"Pavia.": true,
"Alessandria.": true,
"ê": true,
"vota": true,
"'int'": true,
"d\"o": true,
"napulitane,": true,
"accussì": true,
"periodo": true,
"stanno": true,
"e,": true,
"È": true,
"Chesta": true,
"‘e": true,
"d'o": true,
"d'a": true,
"stata": true,
"meglio": true,
"fanno": true,
"Salierno.": true,
"scritto": true,
"chillu": true,
"Como.": true,
"La": true,
"per": true,
"De": true,
"lengua": true,
"sul": true,
"sò": true,
"Reggio": true,
"sua": true,
"Cosenza.": true,
"ccose": true,
"'int'â": true,
"cumune": true,
"chella": true,
"Se": true,
"munno,": true,
"Fuje": true,
"primmo": true,
"dal": true,
"gruppo": true,
"uno": true,
"seculo": true,
"Varese.": true,
"comm": true,
"quase": true,
"le": true,
"of": true,
"Udine.": true,
"fino": true,
"fu": true,
"trova": true,
"nfra": true,
"canzone,": true,
"fujeno": true,
"città": true,
"tre": true,
"chesti": true,
"senza": true,
"Romma.": true,
"anne": true,
"isso": true,
"facette": true,
"Valle": true,
"sempe": true,
"maggio": true,
"forma": true,
"storia;": true,
"nel": true,
"mpurtante": true,
"canzona.": true,
"belle": true,
"(Napule,": true,
"Vicenza.": true,
"‘a": true,
"Festival": true,
"chesta": true,
"chille": true,
"ll'": true,
"sujo": true,
"po'": true,
"hanno": true,
"Maria": true,
"cchiú": true,
"ne": true,
"Asti.": true,
"suo": true,
"l'": true,
"cantante": true,
"pecché": true,
"settembre": true,
"Monte": true,
"duje": true,
"Cremona.": true,
"int'": true,
"naturale": true,
"pparte": true,
"Bolzano.": true,
"fà": true,
"paese": true,
"Villa": true,
"jennaro": true,
"munno.": true,
"Messina.": true,
"avette": true,
"soja": true,
"parole": true,
"int'a": true,
"lo": true,
"Il": true,
"Giovanni": true,
"Berlusconi": true,
"cui": true,
"d‘": true,
"li": true,
"napulitana.": true,
"nascette": true,
"puete": true,
"u": true,
"Pruvincia": true,
"cantate": true,
"ogge,": true,
"l'Aquila.": true,
"Padova.": true,
"I": true,
"Lecce.": true,
"Chiete.": true,
"paricchie": true,
"famiglia": true,
"cultura": true,
"state": true,
"straniere": true,
"Castel": true,
"Verona.": true,
"aroppo": true,
"proprio": true,
"cchiu'": true,
"Putenza.": true,
"contro": true,
"abbrile": true,
"Pietro": true,
"Giorgio": true,
"int'ô": true,
"non": true,
"dicembre": true,
"’nu": true,
"Treviso.": true,
"'int'a": true,
"particulare": true,
"lu": true,
"delle": true,
"nuvembre": true,
"zona": true,
"fora": true,
"nummere": true,
"E": true,
"nord": true,
"'int'o": true,
"d\"e": true,
"essere": true,
"Lecco.": true,
"poco": true,
"Calabbria.": true,
"juorne": true,
}