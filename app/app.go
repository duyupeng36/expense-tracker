package app

import (
	"errors"
	"expense-tracker/expense"
	"expense-tracker/fs"
	"fmt"
	"log"
	"sort"
	"time"
)

var App = New()

type ExpenseTracker interface {
	Add(description string, amount float64) error
	Delete(id int) error
	Summary(month time.Month)
	List()
}

type application struct {
	expenses []*expense.Expense
}

func (a *application) Add(description string, amount float64) error {
	if a == nil {
		return errors.New("application must not be nil")
	}
	id, err := a.generateID()
	if err != nil {
		return err
	}

	a.expenses = append(a.expenses, expense.New(id, description, amount))

	if err := a.writeFile(); err != nil {
		return err
	}

	log.Printf("Expense added successfully (ID: %d)\n", id)

	return nil
}

func (a *application) Delete(id int) error {

	for i, e := range a.expenses {
		if e.ID == id {
			a.expenses = append(a.expenses[:i], a.expenses[i+1:]...)
			log.Println("Expense delete successfully")
			return a.writeFile()
		}
	}

	return errors.New("expense not found")
}

func (a *application) Summary(month time.Month) {
	total := 0.0
	mothTotal := 0.0
	for _, e := range a.expenses {
		switch {
		case month == 0:
			total += e.Amount
		case e.Date.Month() == month:
			mothTotal += e.Amount
		}
	}
	if month == 0 {
		fmt.Printf("Total expense: $%v\n", total)
	} else {
		fmt.Printf("Total expense  for %v: $%v\n", time.Month(month), total)
	}
}

func (a *application) List() {
	fmt.Printf("%-5v %-13v %-14v %-5v\n", "ID", "Date", "Description", "Amount")
	for _, e := range a.expenses {
		fmt.Printf("%-5v %-13v %-14v %-5v\n", e.ID, e.Date, e.Description, e.Amount)
	}
}

func (a *application) generateID() (int, error) {
	if a == nil {
		return 0, errors.New("application must not be nil")
	}

	if len(a.expenses) == 0 {
		return 1, nil
	}

	if err := a.sortTasks(); err != nil {
		return 0, err
	}

	minID := a.expenses[0].ID

	if minID != 1 {
		return 1, nil
	}

	for i, j := 0, 1; j < len(a.expenses); i, j = i+1, j+1 {
		if a.expenses[j].ID-a.expenses[i].ID != 1 {
			return a.expenses[i].ID + 1, nil
		}
	}

	return a.expenses[len(a.expenses)-1].ID + 1, nil
}

func (a *application) sortTasks() error {
	if a == nil {
		return errors.New("application must not be nil")
	}
	sort.Slice(a.expenses, func(i, j int) bool {
		return a.expenses[i].ID < a.expenses[j].ID
	})
	return nil
}

func (a *application) readFile() (err error) {
	if a == nil {
		return errors.New("application must not be nil")
	}

	a.expenses, err = fs.ReadTaskFromFile()
	if err != nil {
		return err
	}
	return nil
}

func (a *application) writeFile() (err error) {
	if a == nil {
		return errors.New("application must not be nil")
	}

	return fs.WriteTaskToFile(a.expenses)
}

func New() ExpenseTracker {
	app := &application{}

	if err := app.readFile(); err != nil {
		log.Fatal(err)
	}

	return app

}
