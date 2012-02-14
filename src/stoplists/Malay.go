package gojustext

var MalayStoplist = map[string]bool {"yang": true,
"dan": true,
"di": true,
"pada": true,
"dengan": true,
"dalam": true,
"untuk": true,
"ini": true,
"dari": true,
"oleh": true,
"merupakan": true,
"adalah": true,
"telah": true,
"sebagai": true,
"sebuah": true,
"tahun": true,
"tidak": true,
"terletak": true,
"kepada": true,
"juga": true,
"ke": true,
"atau": true,
"daripada": true,
"ialah": true,
"the": true,
"Pada": true,
"Lapangan": true,
"Terbang": true,
"satu": true,
"itu": true,
"menjadi": true,
"bagi": true,
"mereka": true,
"lebih": true,
"mempunyai": true,
"orang": true,
"akan": true,
"of": true,
"seperti": true,
"boleh": true,
"seorang": true,
"beliau": true,
"bandar": true,
"Malaysia": true,
"kerana": true,
"luar.": true,
"Pautan": true,
"Ia": true,
"bahawa": true,
"and": true,
"ia": true,
"secara": true,
"iaitu": true,
"bahasa": true,
"tetapi": true,
"digunakan": true,
"antara": true,
"Sekolah": true,
"negara": true,
"kawasan": true,
"dua": true,
"serta": true,
"nama": true,
"pertama": true,
"menggunakan": true,
"terbang": true,
"beberapa": true,
"Dalam": true,
"banyak": true,
"Beliau": true,
"kedudukan": true,
"memiliki": true,
"tersebut": true,
"lain": true,
"Amerika": true,
"dikenali": true,
"selepas": true,
"meter": true,
"in": true,
"dia": true,
"lapangan": true,
"hari": true,
"masa": true,
"to": true,
"Malaysia.": true,
"atas": true,
"itu,": true,
"terdapat": true,
"salah": true,
"dapat": true,
"Indonesia.": true,
"ketinggian": true,
"hanya": true,
"aras": true,
"Di": true,
"sebelum": true,
"ketika": true,
"Negeri": true,
"sehingga": true,
"mana": true,
"apabila": true,
"Melayu": true,
"kerajaan": true,
"a": true,
"melalui": true,
"sistem": true,
"lagi": true,
"Islam": true,
"bawah": true,
"pelajar": true,
"besar": true,
"termasuk": true,
"hingga": true,
"tentera": true,
"utama": true,
"sama": true,
"kumpulan": true,
"ada": true,
"masih": true,
"ahli": true,
"pernah": true,
"kapal": true,
"tempat": true,
"laut..": true,
"The": true,
"suatu": true,
"Kebangsaan": true,
"terhadap": true,
"baru": true,
"ini,": true,
"bahagian": true,
"ini.": true,
"setiap": true,
"penduduk": true,
"paling": true,
"negeri": true,
"semasa": true,
"Menteri": true,
"bermula": true,
"bentuk": true,
"tiga": true,
"kemudian": true,
"Kuala": true,
"dunia": true,
"filem": true,
"pasukan": true,
"wilayah": true,
"daerah": true,
"buah": true,
"pula": true,
"anak": true,
"jumlah": true,
"lelaki": true,
"Sungai": true,
"para": true,
"semua": true,
"Sultan": true,
"bersama": true,
"Burung": true,
"turut": true,
"mendapat": true,
"pesawat": true,
"Dia": true,
"kini": true,
"semula": true,
"kecamatan": true,
"Syarikat": true,
"Tentera": true,
"pelbagai": true,
"kuasa": true,
"Pulau": true,
"raya": true,
"seramai": true,
"pihak": true,
"kali": true,
"berjaya": true,
"air": true,
"kecil": true,
"bulan": true,
"awal": true,
"diterbitkan": true,
"Perang": true,
"Bahasa": true,
"membuat": true,
"mula": true,
"nombor": true,
"biasanya": true,
"sekolah": true,
"Jepun": true,
"terdiri": true,
"Ini": true,
"is": true,
"Raja": true,
"sekitar": true,
"lagu": true,
"latitude": true,
"tanah": true,
"longtitude": true,
"tanpa": true,
"diri": true,
"Datuk": true,
"amat": true,
"Abdul": true,
"orang.": true,
"biasa": true,
"menyebabkan": true,
"bukan": true,
"sejak": true,
"jenis": true,
"membawa": true,
"itu.": true,
"lain.": true,
"menjadikan": true,
"abad": true,
"Mereka": true,
"syarikat": true,
"mengambil": true,
"tinggi": true,
"selama": true,
"kedua": true,
"luar": true,
"Kabupaten": true,
"mungkin": true,
"sering": true,
"rumah": true,
"kurang": true,
"Selepas": true,
"sebahagian": true,
"Kumpulan": true,
"sebanyak": true,
"seluruh": true,
"zaman": true,
"haiwan": true,
"siri": true,
"kebanyakan": true,
"empat": true,
"Latitud": true,
"Longitud": true,
"tersebut.": true,
"Kampung": true,
"Bandar": true,
"Timur": true,
"China": true,
"kereta": true,
"menerima": true,
"Dasar": true,
"mengenai": true,
"berada": true,
"setelah": true,
"British": true,
"politik": true,
"memberi": true,
"tentang": true,
"pusat": true,
"bidang": true,
"Dunia": true,
"Taman": true,
"keseluruhan": true,
"berasal": true,
"Gunung": true,
"berlaku": true,
"mampu": true,
"Indonesia": true,
"for": true,
"musim": true,
"bahan": true,
"menghasilkan": true,
"dibina": true,
"terkenal": true,
"Presiden": true,
"kembali": true,
"bin": true,
"agama": true,
"memberikan": true,
"lama": true,
"as": true,
"Perdana": true,
"ibu": true,
"kemudiannya": true,
"masyarakat": true,
"Kerajaan": true,
"laut": true,
"Nama": true,
"was": true,
"berbanding": true,
"segi": true,
"dianggap": true,
"mereka.": true,
"walaupun": true,
"Seri": true,
"penting": true,
"Bukit": true,
"cara": true,
"sendiri": true,
"sejenis": true,
"Selain": true,
"Jalan": true,
"sangat": true,
"Asia": true,
"Cina": true,
"akhir": true,
"New": true,
"Barat": true,
"menunjukkan": true,
"Januari": true,
"Walaupun": true,
"manakala": true,
"Universiti": true,
"terbesar": true,
"Dengan": true,
"Arab": true,
"keadaan": true,
"manusia": true,
"Masjid": true,
"keluarga": true,
"Jerman": true,
"jalan": true,
"waktu": true,
"Disember": true,
"Daerah": true,
"disebabkan": true,
"Terdapat": true,
"album": true,
"kod": true,
"ramai": true,
"juta": true,
"sepanjang": true,
"with": true,
"perlu": true,
"Oleh": true,
"muzik": true,
"mengikut": true,
"that": true,
"menyatakan": true,
"hampir": true,
"jika": true,
"Mac": true,
"serangan": true,
"hidup": true,
"Dewan": true,
"Ketua": true,
"pilihan": true,
"murid": true,
"India": true,
"Menengah": true,
"bekas": true,
"September": true,
"kalendar": true,
"Laut": true,
"Majlis": true,
"Wilayah": true,
"Antarabangsa": true,
"sahaja": true,
"Mei": true,
"namun": true,
"kedua-dua": true,
"April": true,
"asli": true,
"Kota": true,
"terus": true,
"merujuk": true,
"maka": true,
"Eropah": true,
"mencapai": true,
"Ogos": true,
"perang": true,
"hasil": true,
"barat": true,
"baik": true,
"on": true,
"sejarah": true,
"pemain": true,
"by": true,
"Jun": true,
"sekali": true,
"\"The": true,
"Negara": true,
"Apabila": true,
"sudah": true,
"peringkat": true,
"dilakukan": true,
"Julai": true,
"sementara": true,
"kota": true,
"wanita": true,
"Sri": true,
"sesuatu": true,
"Yang": true,
"kaum": true,
"lalu": true,
"pokok": true,
"buku": true,
"selatan": true,
"menyertai": true,
"mengandungi": true,
"Parti": true,
"ekonomi": true,
"Republik": true,
"anggota": true,
"Oktober": true,
"udara": true,
"Timur,": true,
"penyanyi": true,
"penggunaan": true,
"lima": true,
"diadakan": true,
"diberikan": true,
"perempuan,": true,
"November": true,
"Selatan": true,
"Al": true,
"rakyat": true,
"perkhidmatan": true,
"membentuk": true,
"dilantik": true,
"Raya": true,
"kampung": true,
"sedang": true,
"parti": true,
"Inggeris": true,
"pemerintahan": true,
"Sebagai": true,
"Tanah": true,
"undang-undang": true,
"tempoh": true,
"berdasarkan": true,
"tinggal": true,
"sesetengah": true,
"rasmi": true,
"hutan": true,
"membantu": true,
"dipanggil": true,
"api": true,
"lain,": true,
"spesies": true,
"bangunan": true,
"akhirnya": true,
"Muhammad": true,
"utara": true,
"Pokok": true,
"Filem": true,
"muncul": true,
"Besar": true,
"Batu": true,
"melakukan": true,
"dikatakan": true,
"badan": true,
"kata": true,
"bernama": true,
"senjata": true,
"harga": true,
"Jakarta": true,
"Ketika": true,
"Ahmad": true,
"pulau": true,
"jawatan": true,
"perkataan": true,
"permainan": true,
"versi": true,
"popular": true,
"Malaysia,": true,
"Setelah": true,
"pentadbiran": true,
"mudah": true,
"batu": true,
"dunia.": true,
"agar": true,
"United": true,
"semakin": true,
"Singapura": true,
"watak": true,
"belakang": true,
"Sulawesi": true,
"golongan": true,
"Perancis": true,
"makanan": true,
"Pilihan": true,
"ditubuhkan": true,
"Februari": true,
"tahun.": true,
"proses": true,
"sebarang": true,
"masalah": true,
"Tun": true,
"Gregory.": true,
"akibat": true,
"Haji": true,
"dibuat": true,
"Nabi": true,
"rancangan": true,
"Namun": true,
"Bagaimanapun,": true,
"Sejarah.": true,
"bola": true,
"stesen": true,
"diberi": true,
"sebelah": true,
"ditulis": true,
"are": true,
"tangan": true,
"timur": true,
"Antara": true,
"jam": true,
"Utara": true,
"ketua": true,
"tiada": true,
"laut.": true,
"mata": true,
"berbeza": true,
"seni": true,
"supaya": true,
"desa": true,
"Untuk": true,
"Buku": true,
"Rakyat": true,
"masuk": true,
"orang-orang": true,
"tempatan": true,
"Dari": true,
"ataupun": true,
"ringkasnya": true,
"istilah": true,
"negara-negara": true,
"menentang": true,
"KRI": true,
"sumber": true,
"kira-kira": true,
"Allah": true,
"terakhir": true,
"de": true,
"kelas": true,
"Belanda": true,
"Orang": true,
"Afrika": true,
"keluar": true,
"asal": true,
"Jabatan": true,
"kalangan": true,
"Melayu.": true,
"Perak": true,
"Satu": true,
"menarik": true,
"Kebanyakan": true,
"komputer": true,
"diambil": true,
"km": true,
"hubungan": true,
"mengalami": true,
"enam": true,
"panjang": true,
"terutamanya": true,
"apa": true,
"memerlukan": true,
"didapati": true,
"dilihat": true,
"dikeluarkan": true,
"memenangi": true,
"from": true,
"puluh": true,
"cuba": true,
"enjin": true,
"Syarikat.": true,
"Johor": true,
"televisyen": true,
"berhampiran": true,
"kerja": true,
"Pasukan": true,
"Sistem": true,
"Mohd": true,
"arah": true,
"dilahirkan": true,
"peluru": true,
"Shah": true,
"asas": true,
"dinamakan": true,
"pelakon": true,
"tenaga": true,
"dihasilkan": true,
"sedikit": true,
"Semasa": true,
"pendidikan": true,
"alam": true,
"Mukim": true,
"agak": true,
"melihat": true,
"wang": true,
"unit": true,
"Piala": true,
"umum": true,
"Selatan,": true,
"Empayar": true,
"Selatan.": true,
"meninggalkan": true,
"bermaksud": true,
"masjid": true,
"tahap": true,
"hak": true,
"membenarkan": true,
"pembangunan": true,
"II": true,
"seseorang": true,
"wujud": true,
"membolehkan": true,
"Abu": true,
"Pusat": true,
"Utara,": true,
"(nombor).": true,
"mendapatkan": true,
"disebut": true,
"Tengah,": true,
"Menurut": true,
"peranan": true,
"begitu": true,
"Melaka": true,
"jauh": true,
"penuh": true,
"Udara": true,
"keputusan": true,
"Parlimen": true,
"I": true,
"mencari": true,
"Islam.": true,
"bekerja": true,
"meninggal": true,
"dipilih": true,
"minyak": true,
"dijadikan": true,
"memasuki": true,
"tersebut,": true,
"or": true,
"rendah": true,
"Teluk": true,
"terpaksa": true,
"Israel": true,
"datang": true,
"laluan": true,
"memulakan": true,
"Hari": true,
"ratus": true,
"bantuan": true,
"be": true,
"jarak": true,
"nilai": true,
"kaki": true,
"kesan": true,
"awam": true,
"Australia": true,
"Dr": true,
"Aceh": true,
"video": true,
"bilangan": true,
"pegawai": true,
"Yahudi": true,
"Tan": true,
"pantai": true,
"tujuan": true,
"(bahasa": true,
"Korea": true,
"aktif": true,
"polis": true,
"gunung": true,
"termasuklah": true,
"ketiga": true,
"budaya": true,
"Harry": true,
"baginda": true,
"membina": true,
"warna": true,
"industri": true,
"milik": true,
"kes": true,
"Kapal": true,
"harus": true,
"khas": true,
"Kawasan": true,
"tetap": true,
"Persekutuan": true,
"pemimpin": true,
"Jepun.": true,
"Lumpur": true,
"moden": true,
"bersama-sama": true,
"membunuh": true,
"(lahir": true,
"program": true,
"sungai": true,
"Swasta": true,
"melibatkan": true,
"sayap": true,
"Baru": true,
"Ibrahim": true,
"Abdullah": true,
"raja": true,
"tengah": true,
"kematian": true,
"sendiri.": true,
"Ahli": true,
"menggantikan": true,
"Sebelum": true,
"perubahan": true,
"mengeluarkan": true,
"menulis": true,
"Pagi": true,
"operasi": true,
"Anugerah": true,
"penyakit": true,
"sebenarnya": true,
"Ibu": true,
"Sdn": true,
"antarabangsa": true,
"dirinya": true,
"at": true,
"bebas": true,
"mereka,": true,
"sebab": true,
"huruf": true,
"mengatakan": true,
"Umum": true,
"kesemua": true,
"Barat,": true,
"terlibat": true,
"Pertama": true,
"pembinaan": true,
"malam": true,
"Selangor": true,
"sewaktu": true,
"Sumatera": true,
"UMNO": true,
"buat": true,
"pengguna": true,
"Tidak": true,
"Sejak": true,
"mati": true,
"John": true,
"bermain": true,
"model": true,
"saluran": true,
"sekarang": true,
"which": true,
"tahun,": true,
"Ali": true,
"karya": true,
"berkenaan": true,
"berat": true,
"maklumat": true,
"Sejarah": true,
"negara.": true,
"menyerang": true,
"Bahagian": true,
"berasaskan": true,
"teknologi": true,
"Bagi": true,
"bergantung": true,
"NPSN": true,
"pun": true,
"memegang": true,
"kita": true,
"TV": true,
"sepak": true,
"kekal": true,
"tumbuhan": true,
"gelaran": true,
"Air": true,
"asalnya": true,
"usaha": true,
"In": true,
"menurut": true,
"Timbalan": true,
"an": true,
"menerusi": true,
"aktiviti": true,
"tindakan": true,
"kuat": true,
"bergerak": true,
"Sukan": true,
"berkaitan": true,
"Brunei.": true,
"sesuai": true,
"langsung": true,
"selain": true,
"mana-mana": true,
"perdagangan": true,
"pertempuran": true,
"sampai": true,
"cerita": true,
"Muslim": true,
"A": true,
"belum": true,
"Lagu": true,
"berbentuk": true,
"perempuan": true,
"cukup": true,
"kadar": true,
"berkuasa": true,
"kebangsaan": true,
"putih": true,
"Jawa": true,
"Barisan": true,
"kehidupan": true,
"Kementerian": true,
"keluasan": true,
"hadapan": true,
"Zaman": true,
"bertindak": true,
"berwarna": true,
"mewakili": true,
"were": true,
"khususnya": true,
"Jika": true,
"suku": true,
"permukaan": true,
"Kampong": true,
"Kesatuan": true,
"ikan": true,
"merangkumi": true,
"Semenanjung": true,
"tujuh": true,
"Rom": true,
"rangkaian": true,
"projek": true,
"Perak.": true,
"bunyi": true,
"tambahan": true,
"Nasional": true,
"Kanada.": true,
"kaedah": true,
"Baginda": true,
"burung": true,
"ruang": true,
"Bhd": true,
"bangsa": true,
"Persatuan": true,
"dibuka": true,
"menyokong": true,
"alat": true,
"sokongan": true,
"bumi": true,
"hitam": true,
"dasar": true,
"menyediakan": true,
"Polis": true,
"Tetapi": true,
"guru": true,
"diterima": true,
"Akta": true,
"surat": true,
"darah": true,
"berkembang": true,
"tak": true,
"kelihatan": true,
"suara": true,
"tertinggi": true,
"it": true,
"Diraja": true,
"mencipta": true,
"pengaruh": true,
"ingin": true,
"Mesir": true,
"bagaimanapun,": true,
"majalah": true,
"ilmu": true,
"dibawa": true,
"tapak": true,
"Tengah": true,
"kawalan": true,
"kegunaan": true,
"Mahkamah": true,
"diketahui": true,
"kulit": true,
"berkhidmat": true,
"sahaja.": true,
"sempadan": true,
"dipercayai": true,
"memilih": true,
"Secara": true,
"penerbangan": true,
"latihan": true,
}