package projects

import "golang.org/x/crypto/bcrypt"


func Hasher (password []byte) string {
	bytes, _ := bcrypt.GenerateFromPassword(password, 10)
	hashedPassword := string(bytes)
	return hashedPassword
}

func CompareHash(password, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}