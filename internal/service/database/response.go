package database

var (
	internalExecutionFailedMessage = "error internal execution failed"
)

// ExecuteResponse - мета информация о выполнении запроса
type ExecuteResponse struct {
	message string
}

// Message -
func (r ExecuteResponse) Message() string {
	return r.message
}
