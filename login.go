package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var users = make(map[string]User)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	if _, exists := users[user.Username]; exists {
		http.Error(w, "El usuario ya existe", http.StatusConflict)
		return
	}

	users[user.Username] = user
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Usuario %s registrado con éxito", user.Username)
}

func main() {
	http.HandleFunc("/register", registerHandler)
	fmt.Println("Servidor escuchando en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
