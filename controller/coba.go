package controller

import (
	"errors"
	"fmt"

	cek "github.com/aiteung/presensi"
	"github.com/aulianafahrian/aufa-ulbi/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"net/http"

	"github.com/aulianafahrian/be_p1/model"
	"github.com/aulianafahrian/be_p1/module"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/aulianafahrian/aufa-ulbi",
		"message":     "You are at the root endpoint 😉 sip",
		"success":     true,
	})
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func GetProyek1(c *fiber.Ctx) error {
	proyek1 := module.GetProyek1FromNPM("1214049", config.Ulbimongoconn, "proyek1")
	return c.JSON(proyek1)
}
func GetDataMahasiswa(c *fiber.Ctx) error {
	mahasiswa := module.GetDataMahasiswaFromNPM("1214049", config.Ulbimongoconn, "mahasiswa")
	return c.JSON(mahasiswa)
}
func GetDataDosen(c *fiber.Ctx) error {
	dosen := module.GetDataDosenFromNID("123456", config.Ulbimongoconn, "dosen")
	return c.JSON(dosen)
}
func GetDataProdi(c *fiber.Ctx) error {
	prodi := module.GetDataProdiFromKodeProdi("11111", config.Ulbimongoconn, "prodi")
	return c.JSON(prodi)
}
func GetDataJurusan(c *fiber.Ctx) error {
	jurusan := module.GetDataJurusanFromKodeJurusan("22222", config.Ulbimongoconn, "jurusan")
	return c.JSON(jurusan)
}

func InsertProyek1(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var proyek1 model.Proyek1
	if err := c.BodyParser(&proyek1); err != nil {
		return err
	}
	insertedID := module.InsertProyek1(db, "proyek1",
		proyek1.Biodata_mahasiswa,
		proyek1.Dosen_pembimbing,
		proyek1.Dosen_penguji,
		proyek1.Judul,
		proyek1.Tanggal_sidang)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data proyek 1 berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
func InsertDataMahasiswa(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var mahasiswa model.Mahasiswa
	if err := c.BodyParser(&mahasiswa); err != nil {
		return err
	}
	insertedID := module.InsertDataMahasiswa(db, "mahasiswa",
		mahasiswa.NPM,
		mahasiswa.Nama,
		mahasiswa.Kelas,
		mahasiswa.Jurusan,
		mahasiswa.Prodi,
	)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data mahasiswa berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
func InsertDataDosen(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var dosen model.Dosen
	if err := c.BodyParser(&dosen); err != nil {
		return err
	}
	insertedID := module.InsertDataDosen(db, "dosen",
		dosen.NID,
		dosen.Nama,
		dosen.Prodi,
	)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data dosen berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataProdi(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var prodi model.Prodi
	if err := c.BodyParser(&prodi); err != nil {
		return err
	}
	insertedID := module.InsertDataProdi(db, "prodi",
		prodi.Kode_Prodi,
		prodi.Nama,
	)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data prodi berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataJurusan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var jurusan model.Jurusan
	if err := c.BodyParser(&jurusan); err != nil {
		return err
	}
	insertedID := module.InsertDataJurusan(db, "jurusan",
		jurusan.Kode_jurusan,
		jurusan.Nama,
	)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data jurusan berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
