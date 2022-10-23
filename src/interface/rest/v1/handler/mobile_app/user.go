package handler_mobile_app_v1

import (
	"net/http"
	"time"
	config "github.com/riyanda432/belajar-authentication/config"

	"github.com/golang-jwt/jwt/v4"
	usecases "github.com/riyanda432/belajar-authentication/src/app/v1/use_cases/mobile_app"
	response "github.com/riyanda432/belajar-authentication/src/interface/rest/response"
	request_v1_mobile "github.com/riyanda432/belajar-authentication/src/interface/rest/v1/request/mobile_app"
	transformer_v1_mobile "github.com/riyanda432/belajar-authentication/src/interface/rest/v1/transformer/mobile_app"
)

// IUserHandler ...
type IUserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	response response.IResponseClient
	usecase  usecases.IUserUseCase
}

// NewUserHandler ...
func NewUserHandler(r response.IResponseClient, u usecases.IUserUseCase) IUserHandler {
	return &UserHandler{
		response: r,
		usecase:  u,
	}
}

// Create ...
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := request_v1_mobile.UserCreate{}
	dto, err := req.Validate(r)
	if err != nil {
		h.response.HttpError(w, err)
		return
	}
	if user, err := h.usecase.Create(ctx, dto); err != nil {
		h.response.HttpError(w, err)
		return
	} else {
		if user == nil {
			h.response.HttpError(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		h.response.JSON(w, "Successfully created data", transformer_v1_mobile.TransformCreateUpdate(user), nil)
	}
}

// Login ...
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := request_v1_mobile.UserLogin{}
	dto, err := req.Validate(r)
	if err != nil {
		h.response.HttpError(w, err)
		return
	}
	if user, err := h.usecase.Login(ctx, dto); err != nil {
		h.response.HttpError(w, err)
		return
	} else {
		if user == nil {
			h.response.HttpError(w, err)
			return
		}
		// proses pembuatan token jwt
		expTime := time.Now().Add(time.Minute * 1)
		claims := &config.JWTClaim{
			Username: user.GetUsername(),
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "go-jwt-mux",
				ExpiresAt: jwt.NewNumericDate(expTime),
			},
		}

		// medeklarasikan algoritma yang akan digunakan untuk signing
		tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// signed token
		token, err := tokenAlgo.SignedString(config.JWT_KEY)
		if err != nil {
			return 
		}
		// set token yang ke cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    token,
			HttpOnly: true,
		})
		h.response.JSON(w, "Successfully Login", transformer_v1_mobile.TransformSuccessLogin(user), nil)
	}
}
