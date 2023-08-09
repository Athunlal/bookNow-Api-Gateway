package utils

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"erro
	rs,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func ExtractError(err error) (string, error) {
	// Check if the error is a gRPC error
	if errStatus, ok := status.FromError(err); ok {
		// Extract the error message from the gRPC error
		errorMessage := errStatus.Message()
		return errorMessage, nil
	} else {
		// Handle non-gRPC errors here
		return "", errors.New("Not a grpc error")
	}
}

func JsonInputValidation(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"Success": false,
		"Message": "client-side input validation failed",
		"Error":   "Error in Binding the JSON Data",
	})
}

func FailureJson(ctx *gin.Context, statusCode int, booleanValue bool, message string, err string) {
	ctx.JSON(statusCode, gin.H{
		"Success": booleanValue,
		"Message": message,
		"Error":   err,
	})
}

func SeatValidation(data string) error {
	pattern := regexp.MustCompile(`A[1-9]|A1[0-9]|A20|B[1-9]|B1[0-9]|B20|S[1-9]|S1[0-9]|S20`)
	if !pattern.MatchString(data) {
		return fmt.Errorf("Compartment name should be A1 to A20 , B1 t0 B2 and S1 to S20")
	}
	return nil
}
