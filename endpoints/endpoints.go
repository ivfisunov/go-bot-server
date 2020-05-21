package endpoints

import (
	"bot-backend/mongodb"
	"bot-backend/types"
	"bot-backend/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"
)

var ldapConfig utils.ConfigLdapJSON

func init() {
	config, err := utils.LoadJSONConfig("ldapconf-secret.json")
	if err != nil {
		log.Fatal("Error reading ldap config file...\n", err)
	}
	ldapConfig = *config
}

type errorResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// RootHandler - dummy route
func RootHandler(ctx *gin.Context) {
	_ = mongodb.GetConnection()
	ctx.JSON(http.StatusOK, gin.H{"Hi": "there!"})
}

// SearchUserByMobile serach user by mobile
func SearchUserByMobile(ctx *gin.Context) {
	var userMobileJSON types.MobileJSONRequest
	if err := ctx.ShouldBindJSON(&userMobileJSON); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "User not found",
			Status:  "Error",
		})
		return
	}
	user := mongodb.SearchUserByMobile(userMobileJSON)
	ctx.JSON(http.StatusOK, user)
}

// SearchUsersByLastName serach user by last name
func SearchUsersByLastName(ctx *gin.Context) {
	var userLastNameJSON types.UserLastNameJSONRequest
	if err := ctx.ShouldBindJSON(&userLastNameJSON); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "Users not found.",
			Status:  "Error",
		})
		return
	}
	users := mongodb.SearchUserByLastName(userLastNameJSON)
	ctx.JSON(http.StatusOK, users)
}

// GetAccountStatus gets and parse account status from AD
func GetAccountStatus(ctx *gin.Context) {
	var mailJSON types.MailJSONRequest
	if err := ctx.ShouldBindJSON(&mailJSON); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			Status:  "Error",
		})
	}

	ldapClient, err := ldap.Dial("tcp", ldapConfig.URL)
	if err != nil {
		fmt.Println("!!!!", err)
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			Status:  "Error",
		})
	}
	defer ldapClient.Close()
	ldapClient.Bind(ldapConfig.BindDN, ldapConfig.Password)
	searchRequest := ldap.NewSearchRequest(
		ldapConfig.SearchDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(mail="+mailJSON.Mail+")",
		[]string{}, // A list attributes to retrieve
		nil,
	)

	searchResult, err := ldapClient.Search(searchRequest)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			Status:  "Error",
		})
	}

	lockoutTime := searchResult.Entries[0].GetAttributeValue("lockoutTime")
	userAccountControl := searchResult.Entries[0].GetAttributeValue("userAccountControl")
	pwdLastSet := searchResult.Entries[0].GetAttributeValue("pwdLastSet")

	userStatusAD := utils.ParseADStatus(lockoutTime, userAccountControl, pwdLastSet)
	ctx.JSON(http.StatusOK, struct {
		Code int `json:"code"`
	}{Code: userStatusAD})
}
