package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
)

func Encrypt(password string) string {
	/*
		Encrypt the password that we receive from the parameters

		Encriptamos la contraseña que recivimos desde los parametros
	*/
	h := sha256.New()
	h.Write([]byte(password))

	/*
		Convert the hexadecimal created in string with this I
		get the encrypted string

		Convierto el Hexadecimal creado en string con esto
		obtengo la cadena encriptada
	*/
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateToken(user string, password string) string {
	/*
		We generate the token with which you can make
		changes in the internal administration application
		this token must be saved in the client's memory to
		make a request send it so that the api can validate
		if you have privileges to consult or edit the
		information . To generate the token we take the user
		and the password we concatenate them and encrypt it
		under Sha-256

		Generamos el token con el cual se podran realizar
		cambios en la aplicacion de administracion interna
		este token se tiene que guardar en la memoria del
		cliente para al momento de realizar una peticion
		enviarla para que la api pueda validar si tiene
		privilegios para consultar o edital la informacion.
		Para generar el token tomamos el usuario y la
		contraseña los concatenamos y lo encriptamos bajo
		Sha-256
	*/
	var concatOne bytes.Buffer
	concatOne.WriteString(user)
	concatOne.WriteString(password)
	cryptP1 := sha256.New()
	cryptP1.Write([]byte(concatOne.String()))
	return hex.EncodeToString(cryptP1.Sum(nil))
}
