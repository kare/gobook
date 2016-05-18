package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Soundex</title>
<body><h3>Soundex</h3>
<p>Compute Soundex codes a list of names</p>`
	form = `<form action="/" method="POST">
<label for="names">Names (comma or space-separated):</label><br />
<input type="text" name="names" size="30"><br />
<input type="submit" value="Compute">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">An error occurred with the input</p>`
	tableStart = `<table><tr><td>Name</td><td>Soundex</td></tr>`
	row        = `<tr><td>%s</td><td>%s</td></tr>`
	tableEnd   = `</table>`
)

type Word struct {
	Value   string
	Soundex string
}

// Generates the Soundex value corresponding to the Word's `Value`
func (w *Word) calculateSoundex() error {

	// Initialize character codes
	codes := [][]rune{}
	zero := []rune{'a', 'e', 'i', 'o', 'u', 'h', 'w', 'y'}
	one := []rune{'b', 'f', 'p', 'v'}
	two := []rune{'c', 'g', 'j', 'k', 'q', 's', 'x', 'z'}
	three := []rune{'d', 't'}
	four := []rune{'l'}
	five := []rune{'m', 'n'}
	six := []rune{'r'}

	codes = append(codes, zero)
	codes = append(codes, one)
	codes = append(codes, two)
	codes = append(codes, three)
	codes = append(codes, four)
	codes = append(codes, five)
	codes = append(codes, six)

	codeMap := map[rune]int{}
	for code, characters := range codes {
		for _, character := range characters {
			codeMap[character] = code
		}
	}

	// Encode letters - fill in final encoding rune by rune
	final := []rune{rune(w.Value[0])}
	var finalIndex int
	for _, runeValue := range w.Value {
		encoded := rune(strconv.Itoa(codeMap[runeValue])[0])
		if encoded != '0' && (finalIndex == 0 || encoded != final[finalIndex]) {
			finalIndex++
			final = append(final, encoded)
			if finalIndex == 3 {
				break
			}
		}
	}

	// Pad when necessary
	if finalIndex < 3 {
		for i := finalIndex; i < 3; i++ {
			final = append(final, '0')
		}
	}

	w.Soundex = string(final)

	return nil
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Must be called before writing response
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if words, ok := processRequest(request); ok {
			writeOutput(writer, words)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func writeOutput(writer http.ResponseWriter, words []Word) {
	fmt.Fprint(writer, tableStart)
	for _, word := range words {
		fmt.Fprintf(writer, row, word.Value, word.Soundex)
	}
	fmt.Fprint(writer, tableEnd)
}

func processRequest(request *http.Request) ([]Word, bool) {
	var words []Word
	if slice, found := request.Form["names"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, name := range strings.Fields(text) {
			word := Word{Value: name}
			word.calculateSoundex()
			words = append(words, word)
		}
	} else {
		return words, false // no data first time form is shown
	}
	return words, true
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8088", nil)
}
