package balance

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/multios12/sail/pkg/balance/converter"
	"github.com/multios12/sail/pkg/balance/models"
	"gorm.io/gorm"
)

func ensureDataDirs() error {
	for _, dir := range []string{
		balancePath,
		path.Join(balancePath, "_pdf"),
		path.Join(balancePath, "db"),
	} {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

var (
	salaryPDFName  = regexp.MustCompile(`^([0-9]{4})年([0-9]{1,2})月(給与|.*賞与|経費)_.+\.pdf$`)
	salaryImportMu sync.Mutex
)

const (
	imported          = "imported"
	replaced          = "replaced"
	alreadyRegistered = "already_registered"
	invalidFilename   = "invalid_filename"
	conversionFailed  = "conversion_failed"
	saveFailed        = "save_failed"
)

type salaryImportResult struct {
	Filename string `json:"filename"`
	Status   string `json:"status"`
}

func salaryPDFKey(filename string) (string, models.BalanceType, bool) {
	match := salaryPDFName.FindStringSubmatch(filename)
	if match == nil {
		return "", 0, false
	}
	month, _ := strconv.Atoi(match[2])
	if month < 1 || month > 12 {
		return "", 0, false
	}
	kind := models.DetailTypeSalary
	switch {
	case match[3] == "経費":
		kind = models.DetailTypeExpense
	case strings.HasSuffix(match[3], "賞与"):
		kind = models.DetailTypeBonus
	}
	return fmt.Sprintf("%s%02d", match[1], month), kind, true
}

// 最新のPDFを先に確認し、DBにない明細を取り込む。
func readSalaries(balancePath string, pdfPassword string) error {
	pdfPath := filepath.Join(balancePath, "_pdf")
	entries, err := os.ReadDir(pdfPath)
	if err != nil {
		return err
	}

	type pdfFile struct {
		name     string
		modified time.Time
	}
	var files []pdfFile
	for _, entry := range entries {
		info, err := entry.Info()
		if err == nil && info.Mode().IsRegular() {
			files = append(files, pdfFile{entry.Name(), info.ModTime()})
		}
	}
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].modified.After(files[j].modified)
	})
	for _, file := range files {
		importSalaryPDF(file.name, filepath.Join(pdfPath, file.name), pdfPassword, false, "")
	}
	return nil
}

func importSalaryPDF(name, source, password string, replace bool, destination string) salaryImportResult {
	result := salaryImportResult{Filename: name, Status: invalidFilename}
	month, kind, valid := salaryPDFKey(name)
	if !valid {
		return result
	}

	result.Status = saveFailed
	detail, err := findBalanceDetailByMonthType(month, kind)
	if err != nil {
		log.Println("給与明細の照合:", err)
		return result
	}
	exists := len(detail.Json) > 0
	if exists && !replace {
		result.Status = alreadyRegistered
		return result
	}

	if kind == models.DetailTypeExpense {
		detail.Json, detail.Image, err = converter.ConvertExpenseDetail(source, password)
	} else {
		detail.Json, detail.Image, err = converter.ConvertSalaryDetail(source, password)
	}
	if err != nil {
		log.Println("給与明細の変換:", err)
		result.Status = conversionFailed
		return result
	}

	// 同名の原本がある場合、DBのコミットに失敗しても元に戻せるようにする。
	backup := filepath.Join(filepath.Dir(source), ".original")
	hadOriginal := false
	if destination != "" {
		err = os.Link(destination, backup)
		hadOriginal = err == nil
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			log.Println("給与明細の保存準備:", err)
			return result
		}
	}
	saved := false
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := upsertBalanceDetail(tx, detail); err != nil {
			return err
		}
		if destination != "" {
			if err := os.Rename(source, destination); err != nil {
				return err
			}
			saved = true
		}
		return nil
	})
	if err != nil {
		log.Println("給与明細の保存:", err)
		if saved {
			if hadOriginal {
				err = os.Rename(backup, destination)
			} else {
				err = os.Remove(destination)
			}
			if err != nil {
				log.Println("給与明細の原本の復元:", err)
			}
		}
		return result
	}

	result.Status = imported
	if exists {
		result.Status = replaced
	}
	return result
}

func uploadSalaryPDF(name string, input io.Reader, replace bool) salaryImportResult {
	if _, _, valid := salaryPDFKey(name); !valid {
		return salaryImportResult{Filename: name, Status: invalidFilename}
	}
	result := salaryImportResult{Filename: name, Status: saveFailed}
	pdfPath := filepath.Join(balancePath, "_pdf")
	stage, err := os.MkdirTemp(pdfPath, ".upload-")
	if err != nil {
		log.Println("給与明細の一時保存:", err)
		return result
	}
	defer os.RemoveAll(stage)
	source := filepath.Join(stage, name)
	output, err := os.OpenFile(source, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		log.Println("給与明細の一時保存:", err)
		return result
	}
	_, copyErr := io.Copy(output, input)
	closeErr := output.Close()
	if err := errors.Join(copyErr, closeErr); err != nil {
		log.Println("給与明細の一時保存:", err)
		return result
	}

	salaryImportMu.Lock()
	defer salaryImportMu.Unlock()
	return importSalaryPDF(name, source, pdfPassword, replace, filepath.Join(pdfPath, name))
}
