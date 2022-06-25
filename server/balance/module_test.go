package balance

import (
	"testing"
)

func TestReadAllData(t *testing.T) {
	salaryPath = "../../data/"
	v, _ := readAllData()
	if len(v) == 0 {
		t.Fail()
	}
}
func TestReadTextFile(t *testing.T) {
	readMonthDir("../../data/202202/")
	//print(s)
	//readTextFile("../../data/202202/salary-cost.txt")
}
func TestReadTextFileToTimeItem(t *testing.T) {
	readTextFileToTimeItem("../../data/202202/salary-time1.txt")
}
