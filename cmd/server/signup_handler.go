package main

import (
	"errors"
	"net/http"

	"github.com/aifaniyi/sample/pkg/domain"
	"github.com/go-chi/render"
)

type signupRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmPassword"`
}

// add request validations here
func (b *signupRequest) Bind(r *http.Request) error {
	if b.Password != b.ConfirmedPassword {
		return errors.New("passwords are not equal")
	}
	return nil
}

type signupResponse struct {
	Status bool `json:"status"`
}

func (s *server) handleSignup() http.HandlerFunc {
	// one time service initialization
	domain := domain.NewDomain(s.repo)

	return func(w http.ResponseWriter, r *http.Request) {
		data := &signupRequest{}
		var err error
		if err = render.Bind(r, data); err != nil {
			httpError(w, err, http.StatusBadRequest)
			return
		}

		if err = domain.Signup(r.Context(), data.Email, data.Password); err != nil {
			httpError(w, err, http.StatusInternalServerError)
			return
		}

		httpResponse(w, signupResponse{
			Status: true,
		}, http.StatusOK)
	}
}
