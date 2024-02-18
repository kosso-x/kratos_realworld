package pwdoptions

import (
	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func EncryptPassword(password string) (ciphertext string, err error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hp), nil
}

// 对比加密后的密码
func ComparePass(passHash, passwd string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(passwd))

	return
}
