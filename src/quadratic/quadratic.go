package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Quadratic Equation Solver</title>
<body><h3>Quadratic Equation Solver</h3>
<p>Solves equations of the form ax² + bx + c</p>`
	form = `<form action="/" method="POST">
<input type="text" name="a" size="3">
<label for="x²">x²</label> +
<input type="text" name="b" size="3">
<label for="x">x</label> +
<input type="text" name="c" size="3"> →
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

type equation struct {
	a        float64
	b        float64
	c        float64
	positive complex128
	negative complex128
}

var eqn equation

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Must be called before writing response
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if inputs, message, ok := processRequest(request); ok {
			eqn := solve(inputs)
			fmt.Fprint(writer, formatSolutions(eqn))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) (map[string]float64, string, bool) {
	inputs := make(map[string]float64)
	inputFields := []string{"a", "b", "c"}
	for _, field := range inputFields {
		if value, found := request.Form[field]; found && len(value) > 0 {
			text := value[0]
			if x, err := strconv.ParseFloat(text, 64); err != nil {
				return inputs, "'" + text + "' is invalid", false
			} else {
				inputs[field] = x
			}
		}
	}

	if len(inputs) == 0 {
		return inputs, "", false // no data first time form is shown
	}
	return inputs, "", true
}

func solve(inputs map[string]float64) equation {
	eqn.a = inputs["a"]
	eqn.b = inputs["b"]
	eqn.c = inputs["c"]
	discriminant := math.Pow(eqn.b, 2) - (4 * eqn.a * eqn.c)
	denominator := 2 * eqn.a
	eqn.positive = (complex(-1*eqn.b, 0) + cmplx.Sqrt(complex(discriminant, 0))) / complex(denominator, 0)
	eqn.negative = (complex(-1*eqn.b, 0) - cmplx.Sqrt(complex(discriminant, 0))) / complex(denominator, 0)
	return eqn
}

func formatSolutions(eqn equation) string {
	return fmt.Sprintf(
		`<label for="x²">%fx² + %fx + %f → %f or %f </label>`,
		eqn.a, eqn.b, eqn.c, eqn.positive, eqn.negative)
}
