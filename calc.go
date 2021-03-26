package main

import "strconv"

func DoCalc(x string, y string) string {
	x1, _ := strconv.Atoi(x)
	y1, _ := strconv.Atoi(y)
	return strconv.Itoa(x1 / y1)
}
