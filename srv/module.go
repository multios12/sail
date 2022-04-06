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

		if exist || !file.IsDir() || (len(file.Name()) != 6 && len(file.Name()) != 7) {
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

	s.Title = s.Month[:4] + "年" + s.Month[4:6] + "月"
	if s.Month[6:] == "S" {
		s.Title += "賞与"
	} else {
		s.Title += "給与"
	}

	filename := filepath.Join(dirPath, "salary01.txt")
	s.Counts, s.IsError = readTextFileToDetailItem(filename)

	filename = filepath.Join(dirPath, "salary02.txt")
	s.Times, s.IsError = readTextFileToTimeItem(filename)

	filename = filepath.Join(dirPath, "salary10.txt")
	s.Salarys, s.IsError = readTextFileToDetailItem(filename)

	filename = filepath.Join(dirPath, "salary20.txt")
	s.Costs, s.IsError = readTextFileToDetailItem(filename)

	filename = filepath.Join(dirPath, "salary30.txt")
	s.Totals, s.IsError = readTextFileToDetailItem(filename)
	return s
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
			if item.Value > 0 {
				items = append(items, item)
			}
			item = DetailItem{Name: "", Value: 0}
		} else if item.Name == "" && r.MatchString(v) {
			item = DetailItem{Name: "", Value: 0}
		} else {
			if item.Value > 0 {
				items = append(items, item)
			}
			if !r.MatchString(v) {
				item = DetailItem{Name: v, Value: 0}
			} else {
				item = DetailItem{Name: "", Value: 0}
				item.Value, _ = strconv.Atoi(v)
			}
		}
	}
	if item.Name != "" && item.Value > 0 {
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
