package gojustext

var SwahiliStoplist = map[string]bool {"ya": true,
"na": true,
"wa": true,
"katika": true,
"kwa": true,
"ni": true,
"la": true,
"za": true,
"mwaka": true,
"kama": true,
"cha": true,
"mji": true,
"kuwa": true,
"kutoka": true,
"au": true,
"Kwa": true,
"yake": true,
"wakazi": true,
"kwenye": true,
"hii": true,
"vya": true,
"pia": true,
"jina": true,
"nchini": true,
"nchi": true,
"kata": true,
"wake": true,
"kati": true,
"alikuwa": true,
"Mkoa": true,
"lakini": true,
"pamoja": true,
"zaidi": true,
"hadi": true,
"wakati": true,
"juu": true,
"watu": true,
"kwanza": true,
"sensa": true,
"miaka": true,
"ina": true,
"mujibu": true,
"wapatao": true,
"iliyofanyika": true,
"Wilaya": true,
"mara": true,
"moja": true,
"ambayo": true,
"Katika": true,
"mnamo": true,
"sehemu": true,
"mkuu": true,
"kwamba": true,
"baada": true,
"Tanzania.": true,
"Mji": true,
"ilikuwa": true,
"hasa": true,
"eneo": true,
"nyingi": true,
"sana": true,
"kubwa": true,
"kazi": true,
"sababu": true,
"inahusu": true,
"huko": true,
"Makala": true,
"humo.": true,
"waishio": true,
"chini": true,
"Marekani": true,
"Afrika": true,
"lake": true,
"ambao": true,
"kila": true,
"jimbo": true,
"hiyo": true,
"lugha": true,
"ndani": true,
"Kristo).": true,
"Mwaka": true,
"ili": true,
"vile": true,
"huu": true,
"upande": true,
"mwa": true,
"yao": true,
"mjini": true,
"wengi": true,
"njia": true,
"aina": true,
"wilaya": true,
"hivyo": true,
"nafasi": true,
"hali": true,
"muda": true,
"(Baada": true,
"kutokana": true,
"mbalimbali": true,
"kawaida": true,
"hili": true,
"hata": true,
"kuna": true,
"tu": true,
"tarehe": true,
"una": true,
"karibu": true,
"karne": true,
"Kenya": true,
"Baada": true,
"dhidi": true,
"ajili": true,
"albamu": true,
"ambapo": true,
"yenye": true,
"Eneo": true,
"kuhusu": true,
"maeneo": true,
"muhimu": true,
"maji": true,
"wimbo": true,
"BK": true,
"siku": true,
"maisha": true,
"tangu": true,
"mtu": true,
"takriban": true,
"New": true,
"filamu": true,
"serikali": true,
"wanaoishi": true,
"huu.": true,
"idadi": true,
"huwa": true,
"mkubwa": true,
"tena": true,
"m": true,
"pili": true,
"Mungu": true,
"tofauti": true,
"kutumia": true,
"zake": true,
"huo": true,
"wao": true,
"zote": true,
"mmoja": true,
"nyingine": true,
"Papa": true,
"usawa": true,
"bahari.": true,
"mbili": true,
"upo": true,
"wote": true,
"Kuna": true,
"bila": true,
"Wakati": true,
"hizi": true,
"kadhaa": true,
"walikuwa": true,
"kundi": true,
"ambaye": true,
"kipindi": true,
"kufanya": true,
"milioni": true,
"kabla": true,
"akiwa": true,
"sasa": true,
"muziki": true,
"wenye": true,
"yote": true,
"wengine": true,
"mto": true,
"mkoa": true,
"Uingereza": true,
"kuanzia": true,
"km².": true,
"vita": true,
"kuliko": true,
"mwisho": true,
"of": true,
"tatu": true,
"Idadi": true,
"ambazo": true,
"Hata": true,
"Kanisa": true,
"kuu": true,
"matumizi": true,
"si": true,
"maarufu": true,
"kanda": true,
"baadaye": true,
"kiwango": true,
"Ni": true,
"mpya": true,
"Marekani.": true,
"familia": true,
"baadhi": true,
"kupitia": true,
"kiasi": true,
"yake.": true,
"Uturuki.": true,
"ikiwa": true,
"pekee": true,
"kabisa": true,
"kupata": true,
"Chuo": true,
"Ulaya": true,
"kusini": true,
"mfumo": true,
"bora": true,
"shule": true,
"ndege": true,
"watoto": true,
"habari": true,
"biashara": true,
"nje": true,
"asili": true,
"Jina": true,
"the": true,
"mpaka": true,
"Mashariki": true,
"nyimbo": true,
"Kusini": true,
"dunia": true,
"Kati": true,
"msingi": true,
"uwezo": true,
"maana": true,
"timu": true,
"umri": true,
"kaskazini": true,
"tar.": true,
"Bahari": true,
"mwezi": true,
"yeye": true,
"yaani": true,
"inaweza": true,
"Lakini": true,
"mengine": true,
"hivi": true,
"kampuni": true,
"utawala": true,
"sheria": true,
"kifo": true,
"nguvu": true,
"bado": true,
"hivyo,": true,
"chake": true,
"Pia": true,
"Mnamo": true,
"mwenye": true,
"mengi": true,
"Tangu": true,
"Tuzo": true,
"ulikuwa": true,
"Hii": true,
"rangi": true,
"Kenya.": true,
"jinsi": true,
"Tanzania": true,
"historia": true,
"kutoa": true,
"dawa": true,
"Ujerumani": true,
"haya": true,
"The": true,
"ugonjwa": true,
"duniani": true,
"mwingine": true,
"mfano": true,
"huku": true,
"awali": true,
"hapo": true,
"ingawa": true,
"wanyama": true,
"jamii": true,
"neno": true,
"mambo": true,
"(amezaliwa": true,
"Kikuu": true,
"iliopo": true,
"Ufaransa.": true,
"chama": true,
"zao": true,
"lilikuwa": true,
"hilo": true,
"sawa": true,
"rasmi": true,
"KK.": true,
"mfalme": true,
"kwenda": true,
"ndogo": true,
"damu": true,
"ambalo": true,
"mahali": true,
"Jimbo": true,
"jumla": true,
"Historia.": true,
"makuu": true,
"uchaguzi": true,
"ndiyo": true,
"mshindi": true,
"Kama": true,
"Yesu": true,
"elimu": true,
"Umoja": true,
"miji": true,
"fulani": true,
"zamani": true,
"Kaskazini": true,
"Yeye": true,
"pwani": true,
"Nobel": true,
"Mto": true,
"shirika": true,
"chati": true,
"kisiwa": true,
"namna": true,
"rais": true,
"magharibi": true,
"mabadiliko": true,
"makao": true,
"huduma": true,
"uhuru": true,
"makubwa": true,
"and": true,
"chuo": true,
"\"The": true,
"kumi": true,
"ambako": true,
"mechi": true,
"kidogo": true,
"kitabu": true,
"mbali": true,
"chakula": true,
"hiki": true,
"Rais": true,
"mita": true,
"klabu": true,
"Wimbo": true,
"haki": true,
"binadamu": true,
"Jamhuri": true,
"mtoto": true,
"iko": true,
"yake,": true,
"Hasa": true,
"halafu": true,
"hatari": true,
"vitabu": true,
"asilimia": true,
"mazingira": true,
"Vita": true,
"jeshi": true,
"mashariki": true,
"ndio": true,
"shughuli": true,
"kimataifa": true,
"KK": true,
"wake.": true,
"taifa": true,
"pa": true,
"bali": true,
"Mkuu": true,
"wana": true,
"kusababisha": true,
"picha": true,
"imani": true,
"utafiti": true,
"dini": true,
"mdogo": true,
"matatizo": true,
"Ujerumani.": true,
"Amerika": true,
"nyumba": true,
"kufikia": true,
"Ziwa": true,
"Kiingereza": true,
"ana": true,
"Uingereza.": true,
"ile": true,
"bahari": true,
"hayo": true,
"nne": true,
"maendeleo": true,
}