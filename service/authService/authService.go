package authService

type Auth struct {
	Username string
	Password string
}

var creds = []Auth{
	{Username: "vishalshinde", Password: "Adm!n@1"},
}

func (a *Auth) Authenticate() (bool, error) {
	for _, c := range creds {
		if c.Username == a.Username {
			if c.Password == a.Password {
				return true, nil
			}
		}
	}

	return false, nil
}
