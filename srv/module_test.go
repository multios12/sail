package main

import (
	"testing"
)

func TestReadAllData(t *testing.T) {
	v := readAllData("../../data/")
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
