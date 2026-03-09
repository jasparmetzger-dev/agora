package auth

func register(username, password string) (string, error) {

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return "", err
	}
	//
	return "", nil
}
