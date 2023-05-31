package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// STRUCT PROYEK1
type Proyek1 struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Biodata_mahasiswa Mahasiswa          `bson:"biodata_mahasiswa,omitempty" json:"biodata_mahasiswa,omitempty"`
	Dosen_pembimbing  Dosen              `bson:"dosen_pembimbing,omitempty" json:"dosen_pembimbing,omitempty"`
	Dosen_penguji     Dosen              `bson:"dosen_penguji,omitempty" json:"dosen_penguji,omitempty"`
	Judul             string             `bson:"judul,omitempty" json:"judul,omitempty"`
	Tanggal_sidang    string             `bson:"tanggal_sidang,omitempty" json:"tanggal_sidang,omitempty"`
}

type Mahasiswa struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NPM     string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama    string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Kelas   string             `bson:"kelas,omitempty" json:"kelas,omitempty"`
	Jurusan Jurusan            `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	Prodi   Prodi              `bson:"prodi,omitempty" json:"prodi,omitempty"`
}

type Dosen struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NID   string             `bson:"nid,omitempty" json:"nid,omitempty"`
	Nama  string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Prodi Prodi              `bson:"prodi,omitempty" json:"prodi,omitempty"`
}

type Prodi struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_Prodi string             `bson:"kode_prodi,omitempty" json:"kode_prodi,omitempty"`
	Nama       string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

type Jurusan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_jurusan string             `bson:"kode_jurusan,omitempty" json:"kode_jurusan,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

// LATIHAN STRUCT PRESENSI

type Karyawan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan     string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Jam_kerja   []JamKerja         `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja  []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	//Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin string   `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Biodata Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

type Lokasi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}
