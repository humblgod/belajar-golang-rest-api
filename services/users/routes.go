package users

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/humblgod/belajar-golang-rest-api/auth"
	"github.com/humblgod/belajar-golang-rest-api/types"
	"github.com/humblgod/belajar-golang-rest-api/utils"
)

type Handler struct {
	store	types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

// regis api endpoint
func (h *Handler) RegistersRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload

	// parse JSON
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload JSON"))
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors) 
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("payload validation error : %v", errors))
	}

	// check if user 
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterPayload
  
	// parse json
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload JSON"))
		return
	}

	// validate payload
	 if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("payload validation error : %v", errors))
	 }

	 // check if user already exists
	 user, err := h.store.GetUserByEmail(payload.Email)
	 if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return 
	 }
	 if user != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user alrd exist"))
		return
	 }


	 // hash password 
	 hashedPassword, err := auth.CreateHashedPassword(payload.Password)
	 if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("erorr hashing plain password"))
		return 
	 }

	 // register user
	 // ! id is auto increment and the time also
	 err = h.store.CreateUser(types.User{
		Username: payload.Username,
		Email: payload.Email,
		Password: hashedPassword,
	 })
	 if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error creating a new user : %v", err))
		return
	 }

	 utils.WriteJSON(w, http.StatusOK, map[string]bool{
		"register": true,
	 })
}