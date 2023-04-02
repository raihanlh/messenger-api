package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	apiPayload "gitlab.com/raihanlh/messenger-api/api/payload"
	http_error "gitlab.com/raihanlh/messenger-api/api/payload/http-error"
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user"
	"gitlab.com/raihanlh/messenger-api/internal/domain/user/payload"
)

type UserHandler struct {
	usecases *dependency.Usecases
}

func New(u *dependency.Usecases) user.Handler {
	return &UserHandler{
		usecases: u,
	}
}

// CreateNewUser godoc
// @Summary Create New User
// @Description create user from request body
// @Tags User
// @Accept application/json
// @Param body body payload.CreateRequest true "Create User"
// @Produce json
// @Success 200 {object} object{status=string,data=payload.CreateResponse}
// @Router /api/v1/user/create [post]
func (h UserHandler) Create(ctx echo.Context) error {
	var body payload.CreateRequest

	if err := ctx.Bind(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(errCustom.HTTPCode, errCustom.HttpResponseError())
	}

	// Validate incoming data
	if err := ctx.Validate(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(http.StatusBadRequest, errCustom)
	}

	// Pass body to usecase
	data, err := h.usecases.User.Create(ctx.Request().Context(), &body)
	if err != nil {
		httpErr, ok := err.(*http_error.Error)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, http_error.InternalServerError(fmt.Sprintf("Failed to create user: %s", httpErr.Error())))
		}
		return ctx.JSON(httpErr.HTTPCode, httpErr.HttpResponseError())
	}

	res := new(apiPayload.BaseResponse)
	res.AddHTTPCode(http.StatusCreated).AddStatus(apiPayload.StatusOK).AddData(data)
	return ctx.JSON(res.HTTPCode, res)
}

// UpdateUser godoc
// @Summary Update User
// @Description update user from request body
// @Tags User
// @Accept application/json
// @Param Authorization header string true "Token"
// @Param id path string true "User ID"
// @Param body body payload.UpdateRequest true "Update User"
// @Produce json
// @Success 200 {object} object{status=string,data=payload.UpdateResponse}
// @Router /api/v1/user/update/{id} [patch]
func (h UserHandler) Update(ctx echo.Context) error {
	var body payload.UpdateRequest

	if err := ctx.Bind(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(errCustom.HTTPCode, errCustom.HttpResponseError())
	}

	// Validate incoming data
	if err := ctx.Validate(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(http.StatusBadRequest, errCustom)
	}

	// Pass body to usecase
	data, err := h.usecases.User.Update(ctx.Request().Context(), &body)
	if err != nil {
		httpErr, ok := err.(*http_error.Error)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, http_error.InternalServerError(fmt.Sprintf("Failed to update user: %s", httpErr.Error())))
		}
		return ctx.JSON(httpErr.HTTPCode, httpErr.HttpResponseError())
	}

	res := new(apiPayload.BaseResponse)
	res.AddHTTPCode(http.StatusOK).AddStatus(apiPayload.StatusOK).AddData(data)
	return ctx.JSON(res.HTTPCode, res)
}

// DeleteUser godoc
// @Summary Delete User
// @Description delete user from request body
// @Tags User
// @Accept application/json
// @Param Authorization header string true "Token"
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} object{status=string,data=payload.DeleteResponse}
// @Router /api/v1/user/delete/{id} [delete]
func (h UserHandler) Delete(ctx echo.Context) error {
	var body payload.DeleteRequest

	if err := ctx.Bind(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(errCustom.HTTPCode, errCustom.HttpResponseError())
	}

	// Validate incoming data
	if err := ctx.Validate(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(http.StatusBadRequest, errCustom)
	}

	// Pass body to usecase
	data, err := h.usecases.User.Delete(ctx.Request().Context(), &body)
	if err != nil {
		httpErr, ok := err.(*http_error.Error)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, http_error.InternalServerError(fmt.Sprintf("Failed to update user: %s", httpErr.Error())))
		}
		return ctx.JSON(httpErr.HTTPCode, httpErr.HttpResponseError())
	}

	res := new(apiPayload.BaseResponse)
	res.AddHTTPCode(http.StatusOK).AddStatus(apiPayload.StatusOK).AddData(data)
	return ctx.JSON(res.HTTPCode, res)
}

// GetUserById godoc
// @Summary Get User By Id
// @Description get user by id
// @Tags User
// @Accept application/json
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} object{status=string,data=payload.GetByIdResponse}
// @Router /api/v1/user/{id} [get]
func (h UserHandler) GetById(ctx echo.Context) error {
	var body payload.GetByIdRequest

	if err := ctx.Bind(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(errCustom.HTTPCode, errCustom.HttpResponseError())
	}

	// Validate incoming data
	if err := ctx.Validate(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(http.StatusBadRequest, errCustom)
	}

	// Pass body to usecase
	data, err := h.usecases.User.GetById(ctx.Request().Context(), &body)
	if err != nil {
		httpErr, ok := err.(*http_error.Error)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, http_error.InternalServerError(fmt.Sprintf("Failed to get user by id: %s", httpErr.Error())))
		}
		return ctx.JSON(httpErr.HTTPCode, httpErr.HttpResponseError())
	}

	res := new(apiPayload.BaseResponse)
	res.AddHTTPCode(http.StatusOK).AddStatus(apiPayload.StatusOK).AddData(data)
	return ctx.JSON(res.HTTPCode, res)
}

// GetAllUser godoc
// @Summary Get All User
// @Description get all users
// @Tags User
// @Accept application/json
// @Param search query string false "Search by user name"
// @Produce json
// @Success 200 {object} object{status=string,data=payload.GetAllResponse}
// @Router /api/v1/users [get]
func (h UserHandler) GetAll(ctx echo.Context) error {
	var body payload.GetAllRequest
	var data *payload.GetAllResponse

	if err := ctx.Bind(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(errCustom.HTTPCode, errCustom.HttpResponseError())
	}

	// Validate incoming data
	if err := ctx.Validate(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(http.StatusBadRequest, errCustom)
	}

	// Pass body to usecase
	data, err := h.usecases.User.GetAll(ctx.Request().Context(), &body)
	if err != nil {
		httpErr, ok := err.(*http_error.Error)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, http_error.InternalServerError(fmt.Sprintf("Failed to get all user: %s", httpErr.Error())))
		}
		return ctx.JSON(httpErr.HTTPCode, httpErr.HttpResponseError())
	}

	res := new(apiPayload.BaseResponse)
	res.AddHTTPCode(http.StatusOK).AddStatus(apiPayload.StatusOK).AddData(data)
	return ctx.JSON(res.HTTPCode, res)
}

// GetByToken godoc
// @Summary Get By Token
// @Description get user by token
// @Tags User
// @Accept application/json
// @Produce json
// @Success 200 {object} object{status=string,data=payload.GetByTokenResponse}
// @Router /api/v1/user [get]
func (h UserHandler) GetByToken(ctx echo.Context) error {
	var body payload.GetByTokenRequest

	if err := ctx.Bind(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(errCustom.HTTPCode, errCustom.HttpResponseError())
	}

	// Validate incoming data
	if err := ctx.Validate(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(http.StatusBadRequest, errCustom)
	}

	// Pass body to usecase
	body.Token = ctx.Get("token").(string)
	data, err := h.usecases.User.GetByToken(ctx.Request().Context(), &body)
	if err != nil {
		httpErr, ok := err.(*http_error.Error)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, http_error.InternalServerError(fmt.Sprintf("Failed to get all user: %s", httpErr.Error())))
		}
		return ctx.JSON(httpErr.HTTPCode, httpErr.HttpResponseError())
	}

	res := new(apiPayload.BaseResponse)
	res.AddHTTPCode(http.StatusOK).AddStatus(apiPayload.StatusOK).AddData(data)
	return ctx.JSON(res.HTTPCode, res)
}

// Login godoc
// @Summary Login
// @Description user login
// @Tags User
// @Accept application/json
// @Param body body payload.LoginRequest true "Login"
// @Produce json
// @Success 200 {object} object{status=string,data=payload.LoginResponse}
// @Router /api/v1/user/login [post]
func (h UserHandler) Login(ctx echo.Context) error {
	var body payload.LoginRequest

	if err := ctx.Bind(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(errCustom.HTTPCode, errCustom.HttpResponseError())
	}

	// Validate incoming data
	if err := ctx.Validate(&body); err != nil {
		errCustom := http_error.BadRequest(err)
		return ctx.JSON(http.StatusBadRequest, errCustom)
	}

	// Pass body to usecase
	data, err := h.usecases.User.Login(ctx.Request().Context(), &body)
	if err != nil {
		httpErr, ok := err.(*http_error.Error)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, http_error.InternalServerError(fmt.Sprintf("Failed to get all user: %s", httpErr.Error())))
		}
		return ctx.JSON(httpErr.HTTPCode, httpErr.HttpResponseError())
	}

	res := new(apiPayload.BaseResponse)
	res.AddHTTPCode(http.StatusOK).AddStatus(apiPayload.StatusOK).AddData(data)

	ctx.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   data.Token,
		Expires: data.Exp,
		Path:    "/",
	})
	return ctx.JSON(res.HTTPCode, res)
}
