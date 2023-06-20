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

	moduleTugbes "github.com/Fatwaff/be_tugbes/module"
	"github.com/aulianafahrian/be_p1/model"
	"github.com/aulianafahrian/be_p1/module"
	inimodellatihan "github.com/indrariksa/be_presensi/model"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/aulianafahrian/aufa-ulbi",
		"message":     "You are at the root endpoint 😉 sip",
		"success":     true,
	})
}

//Latihan Presensi

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensi [get]
func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
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

// InsertData godoc
// @Summary Insert data presensi.
// @Description Input data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi inimodellatihan.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodullatihan.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// fungsi UPDATE presensi

// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi inimodellatihan.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodullatihan.UpdatePresensi(db, "presensi",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// fungsi DELETE presensi

// DeletePresensiByID godoc
// @Summary Delete data presensi.
// @Description Hapus data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePresensiByID(c *fiber.Ctx) error {
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

	err = inimodullatihan.DeletePresensiByID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

//proyek 1

func GetAllProyek1(c *fiber.Ctx) error {
	py1 := module.GetAllProyek1(config.Ulbimongoconn, "proyek1")
	return c.JSON(py1)
}

func GetAllDataMhs(c *fiber.Ctx) error {
	mhs := module.GetAllDataMahasiswa(config.Ulbimongoconn, "mahasiswa")
	return c.JSON(mhs)
}

func GetAllDataDsn(c *fiber.Ctx) error {
	dsn := module.GetAllDataDosen(config.Ulbimongoconn, "dosen")
	return c.JSON(dsn)
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

// tugbes uas

// GetAllProyek godoc
// @Summary Get All Data Proyek.
// @Description Mengambil semua data proyek.
// @Tags Proyek
// @Accept json
// @Produce json
// @Success 200 {object} Proyek
// @Router /proyek [get]
func GetAllProyek(c *fiber.Ctx) error {
	var data []model.Proyek
	ps := moduleTugbes.GetAllDocs(config.Ulbimongoconn2, "proyek", data)
	return c.JSON(ps)
}

// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
func GetProyekFromID(c *fiber.Ctx) error {
	var data model.Proyek
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
	ps, err := moduleTugbes.GetDocFromID(objID, config.Ulbimongoconn2, "proyek", data)
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

func GetProyekFromID2(c *fiber.Ctx) error {
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
	ps, err := moduleTugbes.GetDocFromID2(objID, config.Ulbimongoconn2, "proyek")
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

// InsertDataProyek godoc
// @Summary Insert data proyek.
// @Description Input data proyek.
// @Tags Proyek
// @Accept json
// @Produce json
// @Param request body Proyek true "Payload Body [RAW]"
// @Success 200 {object} Proyek
// @Failure 400
// @Failure 500
// @Router /proyek [post]
func InsertDataProyek(c *fiber.Ctx) error {
	db := config.Ulbimongoconn2
	var proyek model.Proyek
	if err := c.BodyParser(&proyek); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := moduleTugbes.InsertOneDoc(db, "proyek", proyek)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// fungsi UPDATE presensi

// UpdateDataProyek godoc
// @Summary Update data proyek.
// @Description Ubah data proyek.
// @Tags Proyek
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Proyek true "Payload Body [RAW]"
// @Success 200 {object} Proyek
// @Failure 400
// @Failure 500
// @Router /proyek/{id} [put]
func UpdateDataProyek(c *fiber.Ctx) error {
	db := config.Ulbimongoconn2

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var proyek model.Proyek
	if err := c.BodyParser(&proyek); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = moduleTugbes.UpdateOneDoc(db, "proyek",
		objectID, proyek)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// fungsi DELETE presensi

// DeleteProyekByID godoc
// @Summary Delete data proyek.
// @Description Hapus data proyek.
// @Tags Proyek
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /proyek/{id} [delete]
func DeleteProyekByID(c *fiber.Ctx) error {
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

	err = moduleTugbes.DeleteDocsByID(objID, config.Ulbimongoconn2, "proyek")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
