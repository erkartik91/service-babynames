package controllers

import (
	"errors"
	"strings"
	"unicode"

	"github.com/erkartik91/service-babynames/models"
	"github.com/erkartik91/service-babynames/orm"
	"github.com/erkartik91/service-babynames/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	funk "github.com/thoas/go-funk"
)

type List struct {
	ORM *orm.ORM
}

func (l *List) GetListIDHandlerFunc(params operations.GetListIDParams) middleware.Responder {
	list, err := l.ORM.List.Read(params.ID)
	if err != nil {
		errorMessage := err.Error()
		errorPayload := models.Error{
			Code:    404,
			Message: &errorMessage,
		}
		return operations.NewGetListIDDefault(404).WithPayload(&errorPayload)
	}

	successPayload := models.List{
		ID:        &list.ID,
		Name:      &list.Name,
		BabyNames: list.BabyNames,
	}

	return operations.NewGetListIDOK().WithPayload(&successPayload)
}

func (l *List) PostListHandlerFunc(params operations.PostListParams) middleware.Responder {
	paramName := strings.TrimSpace(params.Name)
	err := validateName(paramName)
	if err != nil {
		errorMessage := err.Error()
		errorPayload := models.Error{
			Code:    403,
			Message: &errorMessage,
		}
		return operations.NewPostListDefault(403).WithPayload(&errorPayload)
	}
	list, err := l.ORM.List.Create(paramName)
	if err != nil {
		errorMessage := err.Error()
		errorPayload := models.Error{
			Code:    500,
			Message: &errorMessage,
		}
		return operations.NewPostListDefault(500).WithPayload(&errorPayload)
	}

	successPayload := models.List{
		ID:        &list.ID,
		Name:      &list.Name,
		BabyNames: list.BabyNames,
	}

	return operations.NewPostListOK().WithPayload(&successPayload)

}

func (l *List) PutListIDAddBabyNameHandlerFunc(params operations.PutListIDAddBabyNameParams) middleware.Responder {

	paramName := strings.TrimSpace(params.BabyName)

	err := validateName(paramName)
	if err != nil {
		errorMessage := err.Error()
		errorPayload := models.Error{
			Code:    403,
			Message: &errorMessage,
		}
		return operations.NewPostListDefault(403).WithPayload(&errorPayload)
	}

	list, err := l.ORM.List.Read(params.ID)
	if err != nil {
		errorMessage := err.Error()
		errorPayload := models.Error{
			Code:    404,
			Message: &errorMessage,
		}
		return operations.NewPutListIDAddBabyNameDefault(404).WithPayload(&errorPayload)
	}
	if funk.Contains(list.BabyNames, paramName) {
		errorMessage := "Baby name already added to the list"
		errorPayload := models.Error{
			Code:    403,
			Message: &errorMessage,
		}
		return operations.NewPutListIDAddBabyNameDefault(403).WithPayload(&errorPayload)
	}

	list.BabyNames = append(list.BabyNames, paramName)
	l.ORM.List.Update(list)

	successPayload := models.List{
		ID:        &list.ID,
		Name:      &list.Name,
		BabyNames: list.BabyNames,
	}

	return operations.NewPutListIDAddBabyNameOK().WithPayload(&successPayload)
}

func (l *List) PutListIDRemoveBabyNameHandlerFunc(params operations.PutListIDRemoveBabyNameParams) middleware.Responder {
	list, err := l.ORM.List.Read(params.ID)
	if err != nil {
		errorMessage := err.Error()
		errorPayload := models.Error{
			Code:    404,
			Message: &errorMessage,
		}
		return operations.NewPutListIDRemoveBabyNameDefault(404).WithPayload(&errorPayload)
	}

	paramName := strings.TrimSpace(params.BabyName)
	if !funk.Contains(list.BabyNames, paramName) {
		errorMessage := "Baby name not present in list"
		errorPayload := models.Error{
			Code:    403,
			Message: &errorMessage,
		}
		return operations.NewPutListIDRemoveBabyNameDefault(403).WithPayload(&errorPayload)
	}

	// remove name from string
	babyNames := []string{}
	for _, babyName := range list.BabyNames {
		if babyName != paramName {
			babyNames = append(babyNames, babyName)
		}
	}

	list.BabyNames = babyNames
	l.ORM.List.Update(list)

	successPayload := models.List{
		ID:        &list.ID,
		Name:      &list.Name,
		BabyNames: list.BabyNames,
	}

	return operations.NewPutListIDRemoveBabyNameOK().WithPayload(&successPayload)
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func validateName(name string) error {
	if !IsLetter(name) {
		return errors.New("Name contains not letter characters")
	}

	if strings.Count(name, " ") > 1 {
		return errors.New("Name contains more than one space")
	}

	return nil
}
