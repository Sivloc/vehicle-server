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
	if id, err := strconv.ParseInt(id_str, 10, 64); err == nil {
		deleted, err := d.store.Vehicle().Delete(r.Context(), id)
		if deleted && err == nil {
			rw.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(rw, "Unfound vehicle with id "+id_str, http.StatusNotFound)
		}
	} else {
		http.Error(rw, "unvalid id "+id_str, http.StatusBadRequest)
	}
}
