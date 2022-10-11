package error_messages

type ServerErrorMessage = string

const (
	DivisionByZero  ServerErrorMessage = "Division by zero"
	UnknownOperator ServerErrorMessage = "Unknown operator"
)
