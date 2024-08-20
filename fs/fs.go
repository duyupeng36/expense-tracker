// Package fs 从文件反序列化任务列表，或将任务序列化到文件
package fs

import (
	"encoding/json"
	"expense-tracker/expense"
	"io"
	"log"
	"os"
	"path/filepath"
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
