package repos



type Repository struct {
	// repo Repository
}

func NewUserRepository() *Repository {
	return &Repository{
		// repo: r,
	}
}



func (r *Repository) GetUser() string {
	

	return "Alisson user"
}
