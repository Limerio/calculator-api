package v1Handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Limerio/calculator-api/server/utils"
)

func Add(w http.ResponseWriter, r *http.Request) {
	var values []float64

	err := json.NewDecoder(r.Body).Decode(&values)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := utils.Float64Reducer(values, func(a, b float64) float64 { return a + b })

	fmt.Fprint(w, strconv.FormatFloat(result, 'f', 5, 64))
}
