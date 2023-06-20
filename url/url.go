package url

import (
	"github.com/aulianafahrian/aufa-ulbi/controller"
	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)                                        //ujicoba panggil package musik
	page.Get("/proyek1", controller.GetProyek1)
	page.Get("/allproyek1", controller.GetAllProyek1)
	page.Get("/allmahasiswa", controller.GetAllDataMhs)
	page.Get("/alldosen", controller.GetAllDataDsn)
	page.Get("/mahasiswa", controller.GetDataMahasiswa)
	page.Get("/dosen", controller.GetDataDosen)
	page.Get("/prodi", controller.GetDataProdi)
	page.Get("/jurusan", controller.GetDataJurusan)
	page.Post("/insertproyek1", controller.InsertProyek1)
	page.Post("/insertmahasiswa", controller.InsertDataMahasiswa)
	page.Post("/insertdosen", controller.InsertDataDosen)
	page.Post("/insertprodi", controller.InsertDataProdi)
	page.Post("/insertjurusan", controller.InsertDataJurusan)
	//presensi
	page.Get("/presensi", controller.GetAllPresensi)    //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	page.Get("/presensi", controller.GetPresensi)
	page.Post("/ins", controller.InsertData)
	page.Put("/upd/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeletePresensiByID)
	page.Get("/docs/*", swagger.HandlerDefault)
	// tugbes uas
	page.Get("/proyek", controller.GetAllProyek)          //menampilkan seluruh data proyek
	page.Get("/proyek/:id", controller.GetProyekFromID)   //menampilkan data proyek berdasarkan id
	page.Get("/proyek2/:id", controller.GetProyekFromID2) //menampilkan data proyek berdasarkan id
	page.Post("/proyek", controller.InsertDataProyek)
	page.Put("/proyek/:id", controller.UpdateDataProyek)
	page.Delete("/proyek/:id", controller.DeleteProyekByID)
}
