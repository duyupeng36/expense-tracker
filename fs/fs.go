// Package fs 从文件反序列化任务列表，或将任务序列化到文件
package fs

import (
	"encoding/csv"
	"encoding/json"
	"expense-tracker/expense"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

var expenseFilePath string

func init() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("Get execution patherror: %v", err)
	}

	expenseFilePath = filepath.Join(filepath.Dir(ex), "expenses.json")
}

func ReadTaskFromFile() (result []*expense.Expense, err error) {

	result = make([]*expense.Expense, 0)

	// 检查文件是否存在
	if _, err := os.Stat(expenseFilePath); os.IsNotExist(err) {
		log.Printf("File %s does not exist\n", expenseFilePath)
		log.Printf("Create task file %s\n", expenseFilePath)
		f, err := os.Create(expenseFilePath)
		if err != nil {
			log.Printf("Create task file %s failed.\n", expenseFilePath)
			return nil, err
		}
		// 程序函数返回时关闭文件
		defer func(closer io.Closer) {
			err := closer.Close()
			if err != nil {
				log.Printf("Close task file %s failed.\n", expenseFilePath)
			}
		}(f)

		if err := json.NewEncoder(f).Encode([]*expense.Expense{}); err != nil {
			log.Printf("Encode task to file %s failed.\n", expenseFilePath)
			return nil, err
		}

		return result, nil
	}

	f, err := os.Open(expenseFilePath)
	if err != nil {
		log.Printf("Open task file %s failed.\n", expenseFilePath)
		return nil, err
	}

	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			log.Printf("Close task file %s failed.\n", expenseFilePath)
		}
	}(f)

	// 读取文件并反序列任务列表
	if err := json.NewDecoder(f).Decode(&result); err != nil {
		log.Printf("Decode task from file %s failed.\n", expenseFilePath)
		return nil, err
	}

	return result, nil
}

func WriteTaskToFile(result []*expense.Expense) error {
	f, err := os.Create(expenseFilePath)
	if err != nil {
		log.Printf("Create task file %s failed.\n", expenseFilePath)
		return err
	}

	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			log.Printf("Close task file %s failed.\n", expenseFilePath)
		}
	}(f)

	if err := json.NewEncoder(f).Encode(result); err != nil {
		log.Printf("Encode task to file %s failed.\n", expenseFilePath)
		return err
	}

	return nil
}

func Export2CSV() error {

	csvFile, err := os.Create(filepath.Join(filepath.Dir(expenseFilePath), "expenses.csv"))
	if err != nil {
		return err
	}

	defer func() {
		err := csvFile.Close()
		if err != nil {
			log.Printf("Close file %s failed.\n", expenseFilePath)
		}
	}()

	csvWriter := csv.NewWriter(csvFile)

	jsonFile, err := os.Open(expenseFilePath)
	if err != nil {
		return err
	}
	defer func() {
		err := jsonFile.Close()
		if err != nil {
			log.Printf("Close file %s failed.\n", expenseFilePath)
		}
	}()

	var expenses []*expense.Expense

	if err := json.NewDecoder(jsonFile).Decode(&expenses); err != nil {
		return err
	}

	var csvHeader []string

	t := reflect.TypeOf(expenses[0]).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		csvHeader = append(csvHeader, field.Tag.Get("json"))
	}

	if err := csvWriter.Write(csvHeader); err != nil {
		return err
	}

	for _, e := range expenses {
		line := []string{strconv.Itoa(e.ID), e.Date.String(), e.Description, strconv.FormatFloat(e.Amount, 'f', -1, 64)}
		if err := csvWriter.Write(line); err != nil {
			return err
		}
	}
	csvWriter.Flush()
	log.Println("export expense to csv file successfully")
	return nil
}
