/*
JWT üç bölümden oluşur:

Header: belirtecin türü ve kullanılan imzalama algoritması. Belirteç türü "JWT" ​​olabilirken İmzalama Algoritması HMAC veya SHA256 olabilir. Ben HS256 kullandım.

Payload: Belirtecin, talepleri içeren ikinci kısmı. Bu iddialar, uygulamaya özel verileri (ör. kullanıcı kimliği, kullanıcı adı), belirteç sona erme süresi (exp) vb. içerir.

İmza: imza oluşturmak için bir header, bir payload ve sağladığınız bir anahtar kullanılır.
*/



package main

import (
	"crypto/rand"
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
	vars := mux.Vars(r)
	id := vars["name"]
	bytes := make([]byte, 256) //AES-256 için rastgele bir 32 bayt anahtar oluşturun.
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	fmt.Printf("\n")

	clamis := &Claims{
		Name:           id,
		StandardClaims: jwt.StandardClaims{},
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis) //token şifreleme
	tokenString, _ := Token.SignedString(bytes)                //şifrelenen tokeni anahtarımıza göndererek imzalı tokeni elde etme
	fmt.Fprint(w, tokenString)                                 // imzalı tokeni endpointe response etme (encoded_header + “.” + encoded_payload, “server_secret”)
	fmt.Println(tokenString)
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/token{name}", getToken).Methods("GET")
	http.ListenAndServe(":8080", mux)
}

