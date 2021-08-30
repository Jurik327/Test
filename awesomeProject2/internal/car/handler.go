package car

import (
	"awesomeProject2/internal/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	carsUrl = "/cars"
	carUrl = "/cars/:uuid"
)


type handler struct {

}

func NewHandler() handlers.Handler{
	return &handler{}
}

  func (h *handler) Register( router *httprouter.Router){
	   router.GET( carsUrl, h.GetList)
	   router.GET( carsUrl, h.CreateCar)
	   router.POST( carUrl, h.GetCarByUUID)
	   router.PUT( carUrl, h.UpdateCar)
	   router.PATCH( carUrl, h.PartiallyUpdateCar)
	   router.DELETE(  carUrl, h.DeleteCar)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params){
 w.Write([]byte("This is list of cars"))
}

func (h *handler)CreateCar(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is create cars"))
}

func (h *handler)GetCarByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is cars by uuid"))
}

func (h *handler)UpdateCar(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is update cars"))
}

func (h *handler)PartiallyUpdateCar(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is partially update cars"))
}

func (h *handler)DeleteCar(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is delete cars"))
}
type CarData struct {

		ID        int    `json:"id"`
		Year     int `json:"year"`
		Make string `json:"make"`
		Model  string `json:"model"`
		Engine    string `json:"engine"`

}