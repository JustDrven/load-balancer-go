package data

type Config struct {
	Address     string `json:address`
	Username    string `json:username`
	Password    string `json:password`
	DbName      string `json:db`
	ServiceType string `json:serviceType`
}
