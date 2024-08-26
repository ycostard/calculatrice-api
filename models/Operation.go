package models

type OperationRequest struct {
	Expression string `json:"expression" binding:"required"`
}
