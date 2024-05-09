package models

type Mahasiswa struct {
	ID     int64  `gorm:"primaryKey" json:"id"`
	Nama   string `gorm:"type:varchar(300)" json:"nama"`
	Matkul string `gorm:"type:varchar(100)" json:"matkul"`
	Nilai  string `gorm:"type:varchar(10)" json:"nilai"`
}
