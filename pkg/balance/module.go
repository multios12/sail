package balance

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var Salaries []Salary // 明細リスト

// 指定されたディレクトリからデータを読み込み、給与明細モデルリストを返す
func readAllData() ([]Salary, error) {
	files, err := ioutil.ReadDir(balancePath)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	salaries := []Salary{}

	for _, file := range files {
		exist := false
		for _, m := range salaries {
			if m.Month == file.Name() {
				exist = true
				break
			}
		}

		if exist || !file.IsDir() || (len(file.Name()) != 6 && len(file.Name()) != 7) {
			continue
		}

		monthDir := filepath.Join(balancePath, file.Name())
		d := readMonthDir(monthDir)
		salaries = append(salaries, d)
		updateBalanceFromSalaries(d.Month[:6], salaries)
	}
	return salaries, nil
}

// 年月ディレクトリを読み込み、明細モデルを返す
func readMonthDir(dirPath string) Salary {
	s := Salary{}
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

	filename = filepath.Join(dirPath, "expense01.txt")
	s.Expense, s.Expenses, s.IsError = readTextFileToExpenseItem(filename)

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Print(err)
		return s
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if file.IsDir() || ext != ".png" {
			continue
		}
		s.Images = append(s.Images, file.Name())
	}
	sort.Slice(s.Images, func(i, j int) bool { return s.Images[i] > s.Images[j] })

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

// テキストファイルの読み込みと解析を行い、ExpenseItemを返す
func readTextFileToExpenseItem(filename string) (expense int, expenses []ExpenseItem, isErr bool) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	regAmount := regexp.MustCompile("^[\\-,0-9 円]+$")

	text := string(bytes)
	lines := strings.Split(text, "\n")
	var targets []string
	for _, v := range lines {
		v = strings.ReplaceAll(v, "\f", "")
		v = strings.ReplaceAll(v, "▲", "-")
		if len(v) == 0 {
			continue
		} else if regAmount.MatchString(v) {
			v = strings.ReplaceAll(v, ",", "")
			v = strings.ReplaceAll(v, " 円", "")
		}

		targets = append(targets, v)
	}

	regAmount = regexp.MustCompile("^[1-9]$")
	r2 := regexp.MustCompile("^[0-9]+$")
	expenseItem := ExpenseItem{}
	for i, v := range targets {
		if i == 0 {
			expense, _ = strconv.Atoi(v)
		} else if regAmount.MatchString(v) {
			if len(expenseItem.Name) > 0 {
				expenses = append(expenses, expenseItem)
			}
			expenseItem = ExpenseItem{}
		} else if expenseItem.Name == "" {
			expenseItem.Name = v
		} else if !r2.MatchString(v) {
			expenseItem.Memo += v
		} else {
			expenseItem.Amount, _ = strconv.Atoi(v)
		}
	}
	if len(expenseItem.Name) > 0 {
		expenses = append(expenses, expenseItem)
	}
	return
}

func updateBalanceFromSalaries(month string, salaries []Salary) {
	b, err := findBalanceByMonth(month)
	b.Month = month
	if err != nil {
		b = Balance{Month: month}
	}
	salary := 0
	paid := 0
	expense := 0
	for _, s := range salaries {
		if s.Month[:6] != month || len(s.Totals) < 3 {
			continue
		}
		salary += s.Totals[0].Value + s.Expense
		paid += s.Totals[2].Value + s.Expense
		expense = s.Expense
		if len(s.Month) == 7 {
			b.Memo = strings.ReplaceAll(b.Memo, "＋賞与", "")
			b.Memo = "＋賞与" + b.Memo
		}
	}

	if b.Salary != salary || b.Paid != paid || b.Expense != expense {
		b.Salary = salary
		b.Paid = paid
		b.Expense = expense
		upsertBalance(b)
	}
}
