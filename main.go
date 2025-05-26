package main

import (
	"CleanArchitecture/handlers"
	"CleanArchitecture/repo"
	"CleanArchitecture/service"
	"log"
	"net/http"
)

func main() {
	// real
	//realRepo := repo.NewUserRepository()
	//userService := service.NewUserService(realRepo)

	// fake
	fakeRepo := repo.NewFakeUserRepository()
	userService := service.NewUserService(fakeRepo)

	handler := handlers.NewUserHandler(userService)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.CreateUserHandler(w, r)
		} else if r.Method == http.MethodGet {
			handler.ListUsersHandler(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
