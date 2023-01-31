package routers

import (
	"net/http"

	nsbhctrl "github.com/aldiramdan/go-crud-mysql/controllers"
)

func nasabahRouter() {

	http.HandleFunc("/", nsbhctrl.Index)
	http.HandleFunc("/nasabah", nsbhctrl.Index)
	http.HandleFunc("/nasabah/index", nsbhctrl.Index)
	http.HandleFunc("/nasabah/add", nsbhctrl.Add)
	http.HandleFunc("/nasabah/edit", nsbhctrl.Edit)
	http.HandleFunc("/nasabah/delete", nsbhctrl.Delete)

}
