package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

//AUTH

// Секретный ключ для подписи токенов
var jwtKey = []byte("svo")

// Структура для хранения полезной информации токена
type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

// Функция генерации подписанного токена с полезной информацией, срок жизни 24 часа
func GenerateJWT(login string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Login: login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTAuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cookie, err := r.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				log.Printf("Пользователь \"%s\" неверная подпись токена\n", claims.Login)
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			log.Printf("Пользователь \"%s\" токен невалиден\n", claims.Login)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next(w, r, ps)
	}
}

func getLoginFromCookie(r *http.Request) (string, error) {
	// Извлекаем токен из куки
	cookie, err := r.Cookie("token")
	if err != nil {
		return "", errors.New("токен не найден в куках")
	}

	tokenString := cookie.Value

	// Парсинг токена и извлечение информации
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Вставьте ваш секретный ключ
		return []byte("svo"), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("не удалось извлечь claims")
	}

	login, ok := claims["login"].(string)
	if !ok {
		return "", errors.New("логин не найден в claims")
	}

	return login, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validateCreds(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = AuthenticateUser(creds.Login, creds.Password)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			http.Error(w, ErrInvalidCredentials.Error(), http.StatusUnauthorized)
			return
		}
		log.Printf("loginHandler | ошибка аутентификации: %s", err)
		http.Error(w, "Непредвиденная ошибка аутентификации", http.StatusInternalServerError)
		return
	}

	tokenString, err := GenerateJWT(creds.Login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
	w.WriteHeader(http.StatusOK) // Успешный вход
}

func registerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validateCreds(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = RegisterUser(creds.Login, creds.Password)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			http.Error(w, "пользователь уже существует", http.StatusUnauthorized)
			return
		}
		log.Printf("registerHandler | ошибка аутентификации: %s", err)
		http.Error(w, "Непредвиденная ошибка аутентификации", http.StatusInternalServerError)
		return
	}

	log.Printf("Пользователь \"%s\" успешно зарегистрирован", creds.Login)

	tokenString, err := GenerateJWT(creds.Login)
	if err != nil {
		log.Println("registerHandler | ошибка генерации токена:", err)
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	w.WriteHeader(http.StatusCreated) // Возвращаем статус 201 Created
}

func meHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	log.Printf("Пользователь \"%s\" отправил токен: %s\n", claims.Login, cookie.Value)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Printf("Пользователь \"%s\" неверная подпись токена\n", claims.Login)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		log.Printf("Пользователь \"%s\" токен невалиден\n", claims.Login)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var data UserData
	data.Login = claims.Login
	json.NewEncoder(w).Encode(data)
}

func validateCreds(creds Credentials) error {
	if creds.Login == "" {
		return ErrEmptyLogin
	}
	if creds.Password == "" {
		return ErrEmptyPassword
	}
	return nil
}

func logoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0), // Ставим истекшее время
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	// Вернуть успешный ответ
	w.WriteHeader(http.StatusOK)
}
