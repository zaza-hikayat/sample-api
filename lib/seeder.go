package lib

import "gorm.io/gorm"

var MahasiswaSeed = `
INSERT INTO "public"."Mahasiswa" ("NoMhs", "Nama", "TanggalLahir", "TanggalMasuk") VALUES
('AA20110200', 'Rezza', '1992-08-01', '2020-01-01'),
('AA20110201', 'Ghozi', '1992-06-11', '2020-01-01'),
('AA20110202', 'Raka', '1990-12-19', '2020-01-01'),
('AA20110203', 'Haikal', '1990-10-25', '2020-01-01'),
('AA20110204', 'Abdul', '1990-02-18', '2020-01-01')  ON CONFLICT DO NOTHING;;
`
var DosenSeed = `
INSERT INTO "public"."Dosen" ("NIK", "Nama", "TanggalLahir") VALUES
('AB00123', 'Rikeu Nur', '1990-01-01'),
('AB00124', 'Aziz Abdul', '1989-08-01'),
('AB00125', 'Wawan S', '1989-11-03'),
('AB00126', 'Andri Setiawan', '1988-03-13')  ON CONFLICT DO NOTHING;
`

var MataKuliahSeed = `
INSERT INTO "public"."MataKuliah" ("Kode", "NIKDosen", "Nama", "SKS") VALUES
('KOM210', 'AB00125', 'Algoritma 1', 1),
('KOM211', 'AB00124', 'B. Inggris', 2),
('KOM212', 'AB00126', 'Struktur Data', 3),
('KOM213', 'AB00123', 'Statistik', 2)  ON CONFLICT DO NOTHING;
`
var NilaiSeed = `
INSERT INTO "public"."Nilai" ("KodeMK", "NoMhs", "Nilai") VALUES
('KOM210', 'AA20110200', 81),
('KOM211', 'AA20110201', 82),
('KOM212', 'AA20110202', 83),
('KOM213', 'AA20110203', 84)  ON CONFLICT DO NOTHING;;
`

func RunSeeder(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Exec(MahasiswaSeed).Error; err != nil {
			return
		}
		if err = tx.Exec(DosenSeed).Error; err != nil {
			return
		}
		if err = tx.Exec(MataKuliahSeed).Error; err != nil {
			return
		}
		if err = tx.Exec(NilaiSeed).Error; err != nil {
			return
		}
		return
	})

	if err != nil {
		return err
	}

	return nil
}
