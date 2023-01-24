package services

import (
	"math/big"
)

type ApiService struct{}

func NewApiService() *ApiService {
	return &ApiService{}
}

func (service *ApiService) Task(slice []int64) ([]bool, error) {
	var result []bool

	for _, s := range slice {
		// Если нужно работать с ОЧЕНЬ большими цифрами -> bigInt из этой же библиотеки
		// Если задача не использовать готовых решений -> можно просто находить корень числа и смотреть на остаток на модулю, если остатка нет - число не подходит
		if big.NewInt(s).ProbablyPrime(0) {
			result = append(result, true)
			continue
		}

		result = append(result, false)
	}

	return result, nil
}
