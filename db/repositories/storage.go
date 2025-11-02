package db


type Storage struct{     // facilitates dependency injection for repo
	UserRepository UserRepository	
}

func NewStorage() *Storage{
	return &Storage{
		UserRepository: &UserRepositoryImpl{},
	}
}