package hasher

import "golang.org/x/crypto/bcrypt"

func HashText(txt string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(txt), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func CompareHashText(hashStr string, cmpStr string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashStr), []byte(cmpStr))
}
