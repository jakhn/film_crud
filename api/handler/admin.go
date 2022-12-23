package handler

import (
	"context"
	"github.com/jakhn/film_crud/config"
	"github.com/jakhn/film_crud/models"
	"github.com/jakhn/film_crud/pkg/helper"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginAdmin godoc
// @ID loginAdmin
// @Router /loginadmin [POST]
// @Summary Create LoginAdmin
// @Description Create LoginAdmin
// @Tags LoginAdmin
// @Accept json
// @Produce json
// @Param Login body models.Login true "LoginAdminRequestBody"
// @Success 201 {object} models.LoginResponse "GetLoginAdminBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) LoginAdmin(c *gin.Context) {
	var login models.Login

	err := c.ShouldBindJSON(&login)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.storage.User().GetByPKey(
		context.Background(),
		&models.UserPrimarKey{Login: login.Login},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	if login.Password != resp.Password {
		c.JSON(http.StatusInternalServerError, errors.New("error password is not correct").Error())
		return
	}

	data := map[string]interface{}{
		"user_id": resp.Id,
	}

	token, err := helper.GenerateJWT(data, config.SuperTimeExpiredAt, h.cfg.AuthSecretKey, h.cfg.SuperAdmin)
	if err != nil {
		log.Printf("error whiling GenerateJWT: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GenerateJWT").Error())
		return
	}

	c.JSON(http.StatusCreated, models.LoginResponse{AccessToken: token})
}
