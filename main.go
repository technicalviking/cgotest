package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"tulipindicators"
)

func main() {
	fileName := fmt.Sprintf(".%sindicatorTestData%sbband.csv", string(os.PathSeparator), string(os.PathSeparator))

	var fileData []byte
	var readErr error

	if fileData, readErr = ioutil.ReadFile(fileName); readErr != nil {
		panic(readErr)
		//return
	}

	fileDataBuf := new(bytes.Buffer)
	fileDataBuf.Write(fileData)

	fileReader := csv.NewReader(fileDataBuf)

	var rows [][]string
	var csvErr error

	if rows, csvErr = fileReader.ReadAll(); csvErr != nil {
		panic(csvErr)
		//return
	}

	values := rows[0][len(rows[0])-100:]

	inputs := make([]float64, len(values))

	for i, val := range values {
		bigFloat, _, parseErr := big.ParseFloat(val, 10, 53, big.ToNearestEven)

		if parseErr != nil {
			//never fires.
			panic(parseErr)
		}

		inputs[i], _ = bigFloat.Float64()
	}

	var outputs [][]float64
	var indicatorErr error

	if outputs, indicatorErr = tulipindicators.Indicators["bbands"]([][]float64{inputs}, []float64{18, 2}); indicatorErr != nil {
		panic(indicatorErr)
		return
	}

	fmt.Printf("Outputs: \n %v \n", outputs)
}
