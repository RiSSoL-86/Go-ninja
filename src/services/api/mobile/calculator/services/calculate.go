package services

import (
	"app/src/services/api/mobile/calculator/models"
	"errors"
)

type ICalculateRepository interface {
	SaveCalculation(req models.CalculateRequest, res models.CalculateResponse) error
}

type CalculateService struct {
	repository ICalculateRepository
}

func NewCalculateService(repository ICalculateRepository) *CalculateService {
	return &CalculateService{repository: repository}
}

func (s *CalculateService) Execute(req models.CalculateRequest) (models.CalculateResponse, error) {
	if len(req.Args) == 0 {
		return models.CalculateResponse{}, errors.New("no arguments provided")
	}

	var result float64
	switch req.Operator {
	case "+":
		for _, arg := range req.Args {
			result += arg
		}
	case "-":
		result = req.Args[0]
		for i := 1; i < len(req.Args); i++ {
			result -= req.Args[i]
		}
	case "*":
		result = 1
		for _, arg := range req.Args {
			result *= arg
		}
	case "/":
		result = req.Args[0]
		for i := 1; i < len(req.Args); i++ {
			if req.Args[i] == 0 {
				return models.CalculateResponse{}, errors.New("division by zero")
			}
			result /= req.Args[i]
		}
	default:
		return models.CalculateResponse{}, errors.New("unsupported operator")
	}

	res := models.CalculateResponse{Result: result}

	_ = s.repository.SaveCalculation(req, res)

	return res, nil
}
