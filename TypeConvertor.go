package TypeConvertor

import (
	"fmt"
	"strconv"
)

func StringToInt(str string) int {
	integer, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("unable to convert string to int", err)
	}
	return integer
}

func StringToInt64(str string) int64 {
	integer, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("unable to convert string to int64", err)
	}
	return integer
}

func ByteArrayToString(arrayOfBytes []byte) string {
	str, err := strconv.ParseInt(string(arrayOfBytes), 10, 64)
	if err != nil {
		fmt.Println("unable to convert byte array to string", err)
	}
	return strconv.FormatInt(str, 10)
}

func FloatToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

func StringToFloat(str string) float64 {
	floatValue, err := strconv.ParseFloat(str, 8)
	if err != nil {
		fmt.Println("unable to convert string to float64", err)
	}
	return floatValue
}

func StringToBool(str string) bool {
	parsedBool, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println("unable to convert string to bool", err)
	}
	return parsedBool
}
