/*
JWT üç bölümden oluşur:

Header: belirtecin türü ve kullanılan imzalama algoritması. Belirteç türü "JWT" ​​olabilirken İmzalama Algoritması HMAC veya SHA256 olabilir. Ben HS256 kullandım.

Payload: Belirtecin, talepleri içeren ikinci kısmı. Bu iddialar, uygulamaya özel verileri (ör. kullanıcı kimliği, kullanıcı adı), belirteç sona erme süresi (exp) vb. içerir.

İmza: imza oluşturmak için bir header, bir payload ve sağladığınız bir anahtar kullanılır.
*/



package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Claims struct { //Payload
	Name string `json:"name"`
	jwt.StandardClaims
}

func getToken(w http.ResponseWriter, r *http.Request) {
	var anahtar = []byte("furkan") //Token oluşturmak için anahtar

	clamis := &Claims{
		StandardClaims: jwt.StandardClaims{},
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis) //token şifreleme
	tokenString, _ := Token.SignedString(anahtar)              //şifrelenen tokeni anahtarımıza göndererek imzalı tokeni elde etme
	fmt.Fprint(w, tokenString)                                 // imzalı tokeni endpointe response etme (encoded_header + “.” + encoded_payload, “server_secret”)
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", getToken)
	http.ListenAndServe(":8080", mux)
}
