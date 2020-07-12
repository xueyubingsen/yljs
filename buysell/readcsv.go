package buysell

var v1 string = "security_code"

func Readcsv(sec string) map[string]string {
	var sect map[string]string
	sect = make(map[string]string)
	sect[v1] = sec
	return sect
}
