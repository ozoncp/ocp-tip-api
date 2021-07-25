package models

import "fmt"

// Tip - информация о совете по решению сложности
type Tip struct {
	Id        uint64
	UserId    uint64
	ProblemId uint64
	Text      string
}

// String возвращает строковое представление структуры Tip
func (t Tip) String() string {
	return fmt.Sprintf("User: %v, Text: %v, Problem: %v", t.UserId, t.Text, t.ProblemId)
}
