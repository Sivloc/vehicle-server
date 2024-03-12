package vehicle

import (
	"net/http"
	"strconv"

	"github.com/Sivloc/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	id_str := r.PathValue("id")
	id, err := strconv.ParseInt(id_str)
	if id, err := strconv.ParseInt(id_str); err == nil {
		// if delete
		//		response 204 "vehicle %d deleted"
		// else
		//		response 404 "unfound vehicle with id: %d"
	} else {
		// response 400 "bad request (unvalid id `%s`)"
	}
	http.Error(rw, "Not Implemented", http.StatusInternalServerError)
}
