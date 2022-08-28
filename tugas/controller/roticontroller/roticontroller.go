package roticontroller

import (
	"html/template"
	"net/http"

	"github.com/jeypc/TUGAS/entities"
	"github.com/jeypc/TUGAS/models"
)

// var validate = libraries.NewValidation()
var rotiModel = models.NewRotiModel()

func Index(response http.ResponseWriter, request *http.Request) {

	roti, _ := rotiModel.FindAll()
	data := map[string]interface{}{
		"roti": roti,
	}

	temp, err := template.ParseFiles("views/roti/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/roti/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)

	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var roti entities.Roti
		roti.Nama_roti = request.Form.Get("nama_roti")
		roti.Tittle = request.Form.Get("tittle")
		roti.Description = request.Form.Get("description")
		roti.Rating = request.Form.Get("rating")
		roti.Image = request.Form.Get("image")

		// var data = make(map[string]interface{})
		// vError := validator.Struct(roti)

		// if vError != nil {
		// 	data["validation"] = vErrors
		// } else {
		// 	data["pesan"] = "data berhasil di simpan"
		// 	rotiModel.Create(roti)
		// }

		rotiModel.Create(roti)
		data := map[string]interface{}{
			"pesan": "data berhasil disimpan",
		}

		temp, _ := template.ParseFiles("views/roti/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {

}
