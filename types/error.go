package types

type ErrorResponse struct {
	Param   string `json:"param,omitempty"`
	Message string `json:"message"`
}

type ErrorType string

const (
	RecipeError ErrorType = "RecipeError"
	PowerError  ErrorType = "PowerError"
	DataError   ErrorType = "DataError"
)

type AppError interface {
	error
	Type() ErrorType
	Param() string
}

type appErrorImpl struct {
	msg   string
	typ   ErrorType
	param string
}

func (e *appErrorImpl) Error() string   { return e.msg }
func (e *appErrorImpl) Type() ErrorType { return e.typ }
func (e *appErrorImpl) Param() string   { return e.param }

func ThrowMsg(msg string) AppError {
	return &appErrorImpl{msg: msg, typ: RecipeError, param: ""}
}

func ThrowRecipe(msg string, param string) AppError {
	return &appErrorImpl{msg: msg, typ: RecipeError, param: param}
}

func ThrowPower(msg string) AppError {
	return &appErrorImpl{msg: msg, typ: PowerError, param: ""}
}

func ThrowData(msg string) AppError {
	return &appErrorImpl{msg: msg, typ: DataError, param: ""}
}
