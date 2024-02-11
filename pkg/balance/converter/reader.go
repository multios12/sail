package converter

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// テキストファイルの読み込みと解析を行い、DetailItemを返す
func readTextFileToDetailItem(text string) (items []DetailItem, isErr bool) {
	r := regexp.MustCompile(`^[\-,0-9]+$`)

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
			var err error
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
func readTextFileToTimeItem(text string) (items []TimeItem, isErr bool) {
	r := regexp.MustCompile("時間")
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
			var err error
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
func readTextFileToExpenseItem(text string) (item ExpenseDetail, err error) {
	regAmount := regexp.MustCompile(`^[\-,0-9 円]+$`)
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
	amountItem := AmountItem{}
	for i, v := range targets {
		if i == 0 {
			item.Expense, _ = strconv.Atoi(v)
		} else if regAmount.MatchString(v) {
			if len(amountItem.Name) > 0 {
				item.Amounts = append(item.Amounts, amountItem)
			}
			amountItem = AmountItem{}
		} else if amountItem.Name == "" {
			amountItem.Name = v
		} else if !r2.MatchString(v) {
			amountItem.Memo += v
		} else {
			amountItem.Amount, _ = strconv.Atoi(v)
		}
	}
	if len(amountItem.Name) > 0 {
		item.Amounts = append(item.Amounts, amountItem)
	}
	return
}

// ----------------------------------------------------------------------------

// pdftotextコマンドを実行し、テキストデータを出力する
func pdftostring(src string, opt string, pdfPassword string) string {
	opt = src + " - -opw " + pdfPassword + " " + opt
	args := strings.Split(opt, " ")
	bytes, _ := exec.Command("pdftotext", args...).Output()
	text := string(bytes)
	text = strings.ReplaceAll(text, "\f", "")
	text = strings.ReplaceAll(text, "\r\n", "\n")
	return text
}

// pdfinfoコマンドを実行し、ページ数を返す
func pdfinfo(filename string, pdfPassword string) (string, error) {
	command := "pdfinfo"

	b, err := exec.Command(command, filename, "-opw", pdfPassword).Output()
	if err != nil {
		return "1", err
	}

	text := string(b)
	lines := strings.Split(text, "\n")
	for _, v := range lines {
		if strings.Contains(v, "Pages:") {
			v = strings.ReplaceAll(v, "Pages:", "")
			v = strings.TrimSpace(v)
			return v, nil
		}
	}
	return "1", nil
}
