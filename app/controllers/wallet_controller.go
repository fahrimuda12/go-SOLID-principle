package controllers

import (
	"go-simple-MVC/app/models"
	"go-simple-MVC/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletDTO struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Balance         float64 `json:"balance"`
	Key_Phrase      string  `json:"key_phrase"`
	User_ID         uint    `json:"user_id"`
	Virtual_Account string  `json:"virtual_account"`
	Tag_Name        string  `json:"tag_name"`
}

type WalletController interface {
	GetAllWallet(c *gin.Context)
	GetDetailWallet(c *gin.Context)
	CreateWallet(c *gin.Context)
	UpdateWallet(c *gin.Context)
	DeleteWallet(c *gin.Context)
  }

type WalletControllerImpl struct {
	svc services.WalletService
}

func (w WalletControllerImpl) GetAllWallet(c *gin.Context) {
	w.svc.GetAllWallet(c)
}

func (w WalletControllerImpl) GetDetailWallet(c *gin.Context) {
	w.svc.GetDetailWallet(c)
}

func (idb *InDB) CreateWallet(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	person.First_Name = first_name
	person.Last_Name = last_name
	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateWallet(c *gin.Context) {
		id := c.Query("id")
		first_name := c.PostForm("first_name")
		last_name := c.PostForm("last_name")

		var (
			person    models.Person
			newPerson models.Person
			result    gin.H
		)

		err := idb.DB.First(&person, id).Error
		if err != nil {
			result = gin.H{
				"result": "Data not found",
			}
		}

		newPerson.First_Name = first_name
		newPerson.Last_Name = last_name
		err = idb.DB.Model(&person).Updates(newPerson).Error
		if err != nil {
			result = gin.H{
				"result": "Update failed",
			}
		} else {
			result = gin.H{
				"result": "Data updated successfully",
			}
		}
		c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteWallet(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	err = idb.DB.Delete(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}

func WalletControllerInit(userService services.WalletService) *WalletControllerImpl {
	return &WalletControllerImpl{
		svc: userService,
	}
}
