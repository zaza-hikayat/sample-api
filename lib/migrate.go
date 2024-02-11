package lib

import "gorm.io/gorm"

var CreateTblDosen = `
CREATE TABLE IF NOT EXISTS  "Dosen" (
	"NIK" varchar(20) PRIMARY KEY,
	"Nama" varchar(100) NOT NULL,
	"TanggalLahir" date NOT NULL
)
`
var CreateTblMahasiswa = `
CREATE TABLE IF NOT EXISTS "Mahasiswa" (
	"NoMhs" varchar(20) PRIMARY KEY,
	"Nama" varchar(100) NOT NULL,
	"TanggalLahir" date NOT NULL,
	"TanggalMasuk" date
)
`
var CreateTblMatakuliah = `
CREATE TABLE IF NOT EXISTS "MataKuliah" (
	"Kode" varchar(10) PRIMARY KEY,
	"NIKDosen" varchar(20) NOT NULL,
	"Nama" varchar(100) NOT NULL,
	"SKS" int4 NOT NULL,
 	FOREIGN KEY ("NIKDosen") REFERENCES "public"."Dosen" ("NIK")  ON DELETE CASCADE ON UPDATE CASCADE
)
`
var CreateTblNilai = `
CREATE TABLE IF NOT EXISTS "Nilai" (
	"KodeMK" varchar(10),
	"NoMhs" varchar(20),
	"Nilai" float Default 0,
	PRIMARY KEY ("KodeMK", "NoMhs"),
 	FOREIGN KEY ("NoMhs") REFERENCES "public"."Mahasiswa" ("NoMhs")  ON DELETE CASCADE ON UPDATE CASCADE,
	 FOREIGN KEY ("KodeMK") REFERENCES "public"."MataKuliah" ("Kode")  ON DELETE CASCADE ON UPDATE CASCADE

)
`

func RunMigration(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Exec(CreateTblDosen).Error; err != nil {
			return
		}
		if err = tx.Exec(CreateTblMahasiswa).Error; err != nil {
			return
		}
		if err = tx.Exec(CreateTblMatakuliah).Error; err != nil {
			return
		}
		if err = tx.Exec(CreateTblNilai).Error; err != nil {
			return
		}
		return
	})

	if err != nil {
		return err
	}

	return nil
}
