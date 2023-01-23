package services

type ApiService struct{}

func NewApiService() *ApiService {
	return &ApiService{}
}

func (service *ApiService) Task(slice []int) ([]bool, error) {
	result := []bool{
		true,
		false,
		true,
	}

	return result, nil
}
