package rad

type RenderService interface {
	ListServices(ownerID string) ([]Service, error)
	DeleteService(serviceID string) error
	ListAuthorizedOwners() ([]Owner, error)
}

type Service struct {
	ID      string
	Name    string
	OwnerID string
}

type Owner struct {
	ID    string
	Name  string
	Email string
	Type  string
}
