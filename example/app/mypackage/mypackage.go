package stuff

import "strconv"

var Name string = "Derek"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, value := range intArr {
		strArr = append(strArr, strconv.Itoa(value))
	}
	return strArr
}
