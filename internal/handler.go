package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-calculator/internal/models"
	msg "rest-calculator/internal/models/error_messages"
	"rest-calculator/internal/models/operator"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// ExecuteOperation godoc
// @Description Performs a mathematical operation
// @Accept      json
// @Produce     json
// @Param       input   body     models.Operation true "Enter operation"
// @Success     200     {object} models.ServerSuccessResponse
// @Failure     400,500 {object} models.ServerErrorResponse
// @Router      /execute [post]
func (handler *Handler) ExecuteOperation(context *gin.Context) {
	json := models.Operation{}
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, models.NewServerErrorResponse(err.Error()))
		return
	}

	var result string
	number1 := *json.NumberOne
	number2 := *json.NumberTwo

	switch json.Operator {
	case operator.Add:
		result = fmt.Sprintf("%.5f", number1+number2)
		break
	case operator.Minus:
		result = fmt.Sprintf("%.5f", number1-number2)
		break
	case operator.Divide:
		if number2 == 0 {
			context.AbortWithStatusJSON(http.StatusBadRequest, models.NewServerErrorResponse(msg.DivisionByZero))
			return
		}
		result = fmt.Sprintf("%.5f", number1/number2)
		break
	case operator.Multiply:
		result = fmt.Sprintf("%.5f", number1*number2)
		break
	default:
		context.AbortWithStatusJSON(http.StatusBadRequest, models.NewServerErrorResponse(msg.UnknownOperator))
		return
	}

	context.JSON(http.StatusOK, models.NewServerSuccessResponse(result))
}
