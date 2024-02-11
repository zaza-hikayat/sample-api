# Test - Sample API
This is technical test api using db postgres and programming language golang
## How to install
### Without docker
- Please copy `.env-example` to `.env`
- Setup `.env` value according your environment
- Before run application please make sure you have been installed go, if you not installed yet, you can download and installed at [Golang Web Official](https://go.dev/doc/install)
- Finally you can run sample api with command at your terminal `go run main.go`

### With docker
- Please copy `.env-example` to `.env`
- Setup `.env` value according your environment
- run your terminal `docker compose up`

### Dokumentasi
- Buatlah pseudocode program
 ```
 Masukkan nilai : 5
  12345
  54321
  12345
  54321
  12345

  -------
  Jawab:
  Deklarasi
  var endNumber: integer;
  INPUT endNumber
  FOR i = 1 To endNumber
    IF i % 2 == 0
      OUTPUT printReverseSequence(endNumber)
    ELSE
      OUTPUT printSequence(endNumber)
    END IF
  END FOR

  SUBROUTINE printSequence(endNumber)
    var s: string
    FOR i = 1 To endNumber
      s += <- string(i)
    END FOR
    RETURN s
  END SUBROUTINE

  SUBROUTINE printReverseSequence(endNumber)
    var s: string
    FOR i = endNumber To 0 downto 1
      s += <- string(i)
    END FOR
    return s
  END SUBROUTINE
  ```
- Pseucode Menampilkan data JSON
```
  Deklarasi
  var name: string;
  var hitPoints, strength, defense, intelligence, class: integer
  var person: object
  var collectionPerson: []person
  INPUT name
  INPUT hitPoints, strength, defense, intelligence, class
  person <- name, hitPoints, strength, defense, intelligence, class
  collectionPerson <- person
  OUTPUT collectionPerson with format json
```

- Soal query
```
-- 1.a Mendapatkan seluruh isi table Mata Kuliah
select * FROM "MataKuliah";

-- 1.b Mendapatkan kolom No Mhs dan kolom  Nama dari table Mahasiswa
select "NoMhs", "Nama" FROM "Mahasiswa";

-- 1.c Mendapatkan seluruh isi table Mata Kuliah yang SKSnya = 1
select * FROM "MataKuliah" WHERE "SKS" = 1;

-- 1.d Mendapatkan dosen yang namanya diawali huruf “A”
select * FROM "Dosen" WHERE "Nama" like 'A%';

-- 1.e Mendapatkan dosen yang mengajar matakuliah dengan kode “KOM210” (diperlukan semua kolom dosen, tidak lebih tidak kurang)
SELECT d.* FROM "Dosen" d JOIN "MataKuliah" m ON d."NIK" = m."NIKDosen" WHERE m."Kode" = 'KOM210';

-- 1.f Mendapatkan nama mahasiswa yang diajar oleh dosen dengan NIK “AB00123” (diperlukan hanya kolom Nama mahasiswa)
SELECT m."Nama" FROM "Mahasiswa" m JOIN "Nilai" n ON n."NoMhs" = m."NoMhs" JOIN "MataKuliah" mk ON mk."Kode" = n."KodeMK" JOIN "Dosen" d ON d."NIK" = mk."NIKDosen" WHERE d."NIK"='AB00123'

-- 1.g Mendapatkan nilai paling tinggi dari mata kuliah dengan kode “KOM210” (hanya 1 kolom yg diperlukan : “NilaiTertinggi”)
select max(n."Nilai") as "NilaiTertinggi" FROM "Nilai" n where "KodeMK" = 'KOM210';

-- 1.h Mendapatkan Nilai rata –rata dari semua mata kuliah (kolom yang diperlukan : “NamaMataKuliah”, “NilaiRataRata”)
select mk."Nama" as "NamaMataKuliah", AVG(n."Nilai") OVER() as "NilaiRataRata" FROM "Nilai" n JOIN "MataKuliah" mk on n."KodeMK" = mk."Kode" ;

-- 2.a.Mengubah nilai mahasiswa dengan NoMhs “AA20110200” danKodeMata Kuliah “KOM210” menjadi 90
UPDATE "Nilai"  set "Nilai" = 90 WHERE "NoMhs" = 'AA20110200' AND "KodeMK" = 'KOM210';

-- 2.b.Mengubah semua nilai mahasiswa yang dosennya bernama “SUDIRMAN” menjadi 0
UPDATE "Nilai"  set "Nilai" = 0 FROM "MataKuliah" mk
JOIN "Dosen" d ON d."NIK" = mk."NIKDosen"
 WHERE "Nilai"."KodeMK" = mk."Kode"
 AND d."Nama" = 'SUDIRMAN';

 -- 3.a.Mengosongkan table nilai
 DELETE FROM "Nilai";

 -- 3.b.enghapus mahasiswa dengan NoMhs “AA20110201”
DELETE FROM "Mahasiswa" WHERE "NoMhs" = 'AA20110201';

-- 4.a Memasukkan data mahasiswa berikut ini : Nama : “Budi”, NoMhs : “AA2011012”, TanggalLahir: 1 januari 1990
INSERT INTO "public"."Mahasiswa" ("NoMhs", "Nama", "TanggalLahir") VALUES
('AA2011012', 'Budi', '1990-01-01');
``
## Endpoints
1. GET      /api/v1/sequence          Get data sequence
2. GET      /api/v1/fizzbuzz          Get data print fizzbuzz
3. GET      /api/v1/contact-chip      Show Contact chip
4. GET      /api/v1/mahasiswa         Get data mahasiswa
5. GET      /api/v1/dosen             Get data dosen
6. GET      /api/v1/matakuliah        Get data matakuliah
7. GET      /api/v1/nilai             Get data nilai
8. POST     /api/v1/show-json         Get JSON and Insert data