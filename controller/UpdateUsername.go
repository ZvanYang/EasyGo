package controller

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func UpdateUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {

	log.Println(ps)

}
