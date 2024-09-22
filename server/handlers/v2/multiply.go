package v2Handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Limerio/calculator-api/server/utils"
	"github.com/Limerio/calculator-api/server/utils/constants"
)

func Multiply(w http.ResponseWriter, r *http.Request) {
	values, _ := r.Context().Value(constants.BODY_JSON).([]float64)

	result := utils.Float64Reducer(values, func(a, b float64) float64 { return a * b })

	fmt.Fprint(w, strconv.FormatFloat(result, 'f', 5, 64))
}
