package rad

type Service struct {
	ID    string
	Name  string
	Owner string
}

type Redis struct {
	ID    string
	Name  string
	Owner Owner
}

type Postgres struct {
	ID    string
	Name  string
	Owner Owner
}

type Owner struct {
	ID    string
	Name  string
	Email string
	Type  string
}
