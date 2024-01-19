package user

type ServiceUser struct{}

type ServiceUserInterface interface {
	FindAll() error
	Store() error
	FindById(int) error
	Update(int) error
	Delete(int) error
}

func (s *ServiceUser) FindAll() error {
	// var users []User
	// if err := app.App.DB.Find(&users) ; err != nil {
	// 	return err
	// }

	return nil
}

func (s *ServiceUser) Store() error {
	return nil
}

func (s *ServiceUser) FindById(id int) error {
	return nil
}

func (s *ServiceUser) Update(id int) error {
	return nil
}

func (s *ServiceUser) Delete(id int) error {
	return nil
}
