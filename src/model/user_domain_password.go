package model

import (
	"crypto/md5"
	"encoding/hex"
)

func (user *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.password))                 //Transformando a senha string em um array de byte
	user.password = hex.EncodeToString(hash.Sum(nil)) //troca a senha que ele vai enviar por uma ecryptada
}
