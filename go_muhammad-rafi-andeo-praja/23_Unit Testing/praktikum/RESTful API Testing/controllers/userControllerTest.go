package controllers

import (
	"net/http"
	"net/http/httptest"
	"project-structure/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
	}
}

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	user := &models.User{
		Name:  "Rafi",
		Email: "rafi@gmail.com",
	}
	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")
	c.SetParamNames("name", "email")
	c.SetParamValues(user.Name, user.Email)

	// Assertions
	if assert.NoError(t, CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
	}
}

func TestUpdateUser(t *testing.T) {
	// Setup
	e := echo.New()
	user := &models.User{
		Name:  "Rafi",
		Email: "rafi@gmail.com",
	}
	req := httptest.NewRequest(http.MethodPut, "/users/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("user", user)

	// Assertions
	if assert.NoError(t, UpdateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
	}
}

func TestDeleteUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, DeleteUser(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
