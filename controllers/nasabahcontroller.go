package nasabahcontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/aldiramdan/go-crud-mysql/entities"
	"github.com/aldiramdan/go-crud-mysql/libraries"
	nsbhmodel "github.com/aldiramdan/go-crud-mysql/models"
)

var validation = libraries.NewValidation()

var nasabahModel = nsbhmodel.NewNasabahModel()

func Index(response http.ResponseWriter, request *http.Request) {

	nasabah, _ := nasabahModel.FindAll()

	data := map[string]interface{}{
		"nasabah": nasabah,
	}

	temp, err := template.ParseFiles("views/nasabah/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
	
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/nasabah/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var nasabah entities.Nasabah
		nasabah.NamaLengkap = request.Form.Get("nama_lengkap")
		nasabah.Alamat = request.Form.Get("alamat")
		nasabah.TempatLahir = request.Form.Get("tempat_lahir")
		nasabah.TanggalLahir = request.Form.Get("tanggal_lahir")
		nasabah.RataPenghasilan = request.Form.Get("avg_penghasilan")
		nasabah.JenisNasabah = request.Form.Get("jenis_nasabah")
		nasabah.IdCustomer = request.Form.Get("id_customer")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(nasabah)

		if vErrors != nil {
			data["nasabah"] = nasabah
			data["validation"] = vErrors
		} else {
			data["message"] = "Data berhasil disimpan"
			nasabahModel.Create(nasabah)
		}

		temp, _ := template.ParseFiles("views/nasabah/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var nasabah entities.Nasabah
		nasabahModel.Find(id, &nasabah)

		data := map[string]interface{}{
			"nasabah": nasabah,
		}

		temp, err := template.ParseFiles("views/nasabah/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var nasabah entities.Nasabah
		nasabah.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		nasabah.NamaLengkap = request.Form.Get("nama_lengkap")
		nasabah.Alamat = request.Form.Get("alamat")
		nasabah.TempatLahir = request.Form.Get("tempat_lahir")
		nasabah.TanggalLahir = request.Form.Get("tanggal_lahir")
		nasabah.RataPenghasilan = request.Form.Get("avg_penghasilan")
		nasabah.JenisNasabah = request.Form.Get("jenis_nasabah")
		nasabah.IdCustomer = request.Form.Get("id_customer")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(nasabah)

		if vErrors != nil {
			data["nasabah"] = nasabah
			data["validation"] = vErrors
		} else {
			data["message"] = "Data berhasil diperbarui"
			nasabahModel.Update(nasabah)
		}

		temp, _ := template.ParseFiles("views/nasabah/edit.html")
		temp.Execute(response, data)

	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	nasabahModel.Delete(id)

	http.Redirect(response, request, "/nasabah", http.StatusSeeOther)

}
