// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/aulianafahrian",
            "email": "1214049@std.ulbi.ac.id"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delete/{id}": {
            "delete": {
                "description": "Hapus data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Delete data presensi.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/ins": {
            "post": {
                "description": "Input data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Insert data presensi.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/presensi": {
            "get": {
                "description": "Mengambil semua data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get All Data Presensi.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                }
            }
        },
        "/presensi/{id}": {
            "get": {
                "description": "Ambil per ID data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get By ID Data Presensi.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/proyek": {
            "get": {
                "description": "Mengambil semua data proyek.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Get All Data Proyek.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Proyek"
                        }
                    }
                }
            },
            "post": {
                "description": "Input data proyek.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Insert data proyek.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Proyek"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Proyek"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/proyek/{id}": {
            "get": {
                "description": "Ambil per ID data proyek.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Get By ID Data Proyek.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Proyek"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "Ubah data proyek.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Update data proyek.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Proyek"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Proyek"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Hapus data proyek.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Delete data proyek.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/proyekdosen": {
            "get": {
                "description": "Mengambil semua data dosen.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Get All Data Dosen.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Dosen"
                        }
                    }
                }
            },
            "post": {
                "description": "Input data dosen.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Insert data dosen.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Dosen"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Dosen"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/proyekdosen/{id}": {
            "get": {
                "description": "Ambil per ID data dosen.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Get By ID Data Dosen.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Dosen"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "Ubah data dosen.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Update data dosen.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Dosen"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Dosen"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Hapus data dosen.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Delete data dosen.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/proyekmahasiswa": {
            "get": {
                "description": "Mengambil semua data mahasiswa.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Get All Data Mahasiswa.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Mahasiswa"
                        }
                    }
                }
            },
            "post": {
                "description": "Input data mahasiswa.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Insert data mahasiswa.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Mahasiswa"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Mahasiswa"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/proyekmahasiswa/{id}": {
            "get": {
                "description": "Ambil per ID data mahasiswa.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Get By ID Data Mahasiswa.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Mahasiswa"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "Ubah data mahasiswa.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Update data mahasiswa.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Mahasiswa"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Mahasiswa"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Hapus data mahasiswa.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Proyek"
                ],
                "summary": "Delete data mahasiswa.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/upd/{id}": {
            "put": {
                "description": "Ubah data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Update data presensi.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Dosen": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "nama": {
                    "type": "string"
                },
                "nid": {
                    "type": "string"
                },
                "prodi": {
                    "$ref": "#/definitions/controller.Prodi"
                }
            }
        },
        "controller.JamKerja": {
            "type": "object",
            "properties": {
                "durasi": {
                    "type": "integer",
                    "example": 8
                },
                "gmt": {
                    "type": "integer",
                    "example": 7
                },
                "hari": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Senin",
                        "Selasa",
                        "Rabu",
                        "Kamis",
                        "Jumat",
                        "Sabtu",
                        "Minggu"
                    ]
                },
                "jam_keluar": {
                    "type": "string",
                    "example": "16:00"
                },
                "jam_masuk": {
                    "type": "string",
                    "example": "08:00"
                },
                "piket_tim": {
                    "type": "string",
                    "example": "Piket Z"
                },
                "shift": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "controller.Jurusan": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "kode_jurusan": {
                    "type": "string"
                },
                "nama": {
                    "type": "string"
                }
            }
        },
        "controller.Karyawan": {
            "type": "object",
            "properties": {
                "hari_kerja": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Senin",
                        "Selasa",
                        "Rabu",
                        "Kamis",
                        "Jumat",
                        "Sabtu",
                        "Minggu"
                    ]
                },
                "jabatan": {
                    "type": "string",
                    "example": "Anonymous"
                },
                "jam_kerja": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.JamKerja"
                    }
                },
                "nama": {
                    "type": "string",
                    "example": "Tes Swagger"
                },
                "phone_number": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        },
        "controller.Mahasiswa": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "jurusan": {
                    "$ref": "#/definitions/controller.Jurusan"
                },
                "kelas": {
                    "type": "string"
                },
                "nama": {
                    "type": "string"
                },
                "npm": {
                    "type": "string"
                },
                "prodi": {
                    "$ref": "#/definitions/controller.Prodi"
                }
            }
        },
        "controller.Presensi": {
            "type": "object",
            "properties": {
                "biodata": {
                    "$ref": "#/definitions/controller.Karyawan"
                },
                "checkin": {
                    "description": "Datetime     primitive.DateTime ` + "`" + `bson:\"datetime,omitempty\" json:\"datetime,omitempty\"` + "`" + `",
                    "type": "string",
                    "example": "MASUK"
                },
                "latitude": {
                    "type": "number",
                    "example": 123.11
                },
                "location": {
                    "type": "string",
                    "example": "Bandung"
                },
                "longitude": {
                    "type": "number",
                    "example": 123.11
                },
                "phone_number": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        },
        "controller.Prodi": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "kode_prodi": {
                    "type": "string"
                },
                "nama": {
                    "type": "string"
                }
            }
        },
        "controller.Proyek": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "biodata_mahasiswa": {
                    "$ref": "#/definitions/controller.Mahasiswa"
                },
                "dosen_pembimbing": {
                    "$ref": "#/definitions/controller.Dosen"
                },
                "dosen_penguji": {
                    "$ref": "#/definitions/controller.Dosen"
                },
                "judul": {
                    "type": "string"
                },
                "partner_mahasiswa": {
                    "$ref": "#/definitions/controller.Mahasiswa"
                },
                "ruang_sidang": {
                    "type": "string"
                },
                "tanggal_sidang": {
                    "type": "string"
                },
                "tipe_proyek": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "aufa-ulbi.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Anganaufea TES SWAG",
	Description:      "this is a simple server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
