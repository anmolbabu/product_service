package api

// chandan@kubric.io
import (
	"encoding/json"
	"net/http"

	"github.com/anmolbabu/product_service/pkg/backend"
)

func init() {
	RegisterAPI("/products", http.MethodPost, "AddProduct")
}

func (ctx ApplicationContext) AddProduct(w http.ResponseWriter, r *http.Request) {
	product := backend.Product{}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&product)

	err := product.Validate()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = product.Add(ctx.client)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfully added the product"))
	return
}
