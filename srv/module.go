package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func readAllData(dataPath string) []DetailModel {
	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Print(err)
		return nil
	}

	details := []DetailModel{}

	for _, file := range files {
		exist := false
		for _, m := range details {
			if m.Month == file.Name() {
				exist = true
				break
			}
		}

		if exist || !file.IsDir() || len(file.Name()) != 6 {
			continue
		}

		monthDir := filepath.Join(dataPath, file.Name())
		d := readMonthDir(monthDir)
		details = append(details, d)
	}
	return details
}

// 年月ディレクトリを読み込み、明細モデルを返す
func readMonthDir(dirPath string) DetailModel {
	s := DetailModel{}
	s.Month = filepath.Base(dirPath)

	filename := filepath.Join(dirPath, "salary-title.txt")
	s.Title = readTextFileToString(filename)

	filename = filepath.Join(dirPath, "salary-count1.txt")
	s.Counts, s.IsError = readTextFileToDetailItem(filename)
	filename = filepath.Join(dirPath, "salary-count2.txt")
	counts2, isErr := readTextFileToDetailItem(filename)
	s.Counts = append(s.Counts, counts2...)
	s.IsError = s.IsError || isErr
	filename = filepath.Join(dirPath, "salary-count3.txt")
	Counts3, isErr := readTextFileToDetailItem(filename)
	s.Counts = append(s.Counts, Counts3...)
	s.IsError = s.IsError || isErr
	filename = filepath.Join(dirPath, "salary-count5.txt")
	Counts5, isErr := readTextFileToDetailItem(filename)
	s.Counts = append(s.Counts, Counts5...)
	s.IsError = s.IsError || isErr

	filename = filepath.Join(dirPath, "salary-time1.txt")
	s.Times, s.IsError = readTextFileToTimeItem(filename)
	filename = filepath.Join(dirPath, "salary-time2.txt")
	time2, isErr := readTextFileToTimeItem(filename)
	s.Times = append(s.Times, time2...)
	s.IsError = s.IsError || isErr

	filename = filepath.Join(dirPath, "salary1.txt")
	s.Salarys, s.IsError = readTextFileToDetailItem(filename)
	filename = filepath.Join(dirPath, "salary2.txt")
	Salarys2, isErr := readTextFileToDetailItem(filename)
	s.Salarys = append(s.Salarys, Salarys2...)
	s.IsError = s.IsError || isErr

	filename = filepath.Join(dirPath, "salary-cost1.txt")
	s.Costs, s.IsError = readTextFileToDetailItem(filename)
	filename = filepath.Join(dirPath, "salary-cost2.txt")
	costs2, isErr := readTextFileToDetailItem(filename)
	s.Costs = append(s.Costs, costs2...)
	s.IsError = s.IsError || isErr

	filename = filepath.Join(dirPath, "salary-total.txt")
	s.Totals, s.IsError = readTextFileToDetailItem(filename)
	return s
}

// テキストファイルの読み込み、文字列を返す
func readTextFileToString(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}

	text := string(bytes)
	text = strings.ReplaceAll(text, "\f", "")
	text = strings.ReplaceAll(text, "\n", "")
	return text
}

// テキストファイルの読み込みと解析を行い、DetailItemを返す
func readTextFileToDetailItem(filename string) (items []DetailItem, isErr bool) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	r := regexp.MustCompile("^[\\-,0-9]+$")

	text := string(bytes)
	lines := strings.Split(text, "\n")
	var targets []string
	for _, v := range lines {
		v = strings.ReplaceAll(v, "\f", "")
		v = strings.ReplaceAll(v, "▲", "-")
		if len(v) == 0 || v == "\f" {
			continue
		} else if r.MatchString(v) {
			v = strings.ReplaceAll(v, ",", "")
		}

		targets = append(targets, v)
	}

	item := DetailItem{Name: "", Value: 0}
	for _, v := range targets {
		if item.Name == "" && !r.MatchString(v) {
			item.Name = v
		} else if item.Name != "" && r.MatchString(v) {
			item.Value, err = strconv.Atoi(v)
			if err != nil {
				isErr = true
			}
			items = append(items, item)
			item = DetailItem{Name: "", Value: 0}
		} else {
			items = append(items, item)
			isErr = true
			if !r.MatchString(v) {
				item = DetailItem{Name: v, Value: 0}
			} else {
				item = DetailItem{Name: "", Value: 0}
				item.Value, _ = strconv.Atoi(v)
			}
		}
	}
	if item.Name != "" {
		items = append(items, item)
	}

	return
}

// テキストファイルの読み込みと解析を行い、TimeItemを返す
func readTextFileToTimeItem(filename string) (items []TimeItem, isErr bool) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	r := regexp.MustCompile("時間")

	text := string(bytes)
	lines := strings.Split(text, "\n")
	var targets []string
	for _, v := range lines {
		if len(v) == 0 || v == "\f" {
			continue
		}

		targets = append(targets, v)
	}

	item := TimeItem{Name: "", Value: ""}
	for _, v := range targets {
		if item.Name == "" && r.MatchString(v) {
			item.Name = v
		} else if item.Name != "" && !r.MatchString(v) {
			item.Value = v
			if err != nil {
				isErr = true
			}
			items = append(items, item)
			item = TimeItem{Name: "", Value: ""}
		} else {
			items = append(items, item)
			isErr = true
			if !r.MatchString(v) {
				item = TimeItem{Name: v, Value: ""}
			} else {
				item = TimeItem{Name: "", Value: ""}
				item.Name = v
			}
		}
	}
	if item.Name != "" {
		items = append(items, item)
	}

	return
}
