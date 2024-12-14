package helpers

import (
	"log"
	"strconv"
)

func ParsingFloat64(value string) float64 {
	result, err := strconv.ParseFloat(value, 64)

	if err != nil {
		log.Default().Println(err)
		return float64(result)
	}

	return result
}

func ParsingInt64(value string) int64 {
	result, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		log.Default().Println(err)
		return int64(result)
	}

	return result
}
