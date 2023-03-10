package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/postgre"
	"dewetour/repositories"

	"github.com/gorilla/mux"
)

func User(r *mux.Router) {
	UserRepository := repositories.MakeRepository(postgre.DB)
	h := handlers.HandlerUser(UserRepository)

	// menghandle request dengan method GET pada endpoint /users
	r.HandleFunc("/users", middleware.AdminAuth(h.GetAllUsers)).Methods("GET")

	// menghandle request dengan method DELETE pada endpoint /user
	r.HandleFunc("/user/{id_user}", middleware.AdminAuth(h.DeleteUser)).Methods("DELETE")

	// menghandle request dengan method GET pada endpoint /user <improve>
	r.HandleFunc("/user", middleware.UserAuth(h.GetDetailUser)).Methods("GET")

	// menghandle request dengan method PATCH pada endpoint /user <improve>
	r.HandleFunc("/user", middleware.UserAuth(middleware.UpdateUserImage(h.UpdateImageUser))).Methods("PATCH")
}
