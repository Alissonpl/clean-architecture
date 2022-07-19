package book


// type Repository interface {}

//Service book usecase
type Service struct {
	// repo Repository
}

//NewService create new service
func NewService() *Service {
	return &Service{
		// repo: r,
	}
}



//GetBook get a book
func (s *Service) GetBook() string {
	

	return "Alisson"
}
