package controllers

import (
  "encoding/json"
  "io/uitil"
  "net/http"
  "github.com/victorsteve/fullstack/api/auth"
  "github.com/victorsteven/fullstack/api/models"
  "github.com/victorsteve/fullstack/api/responses"
  "github.com/victorsteven/fullstack/api/utils/formaterro"
  "golang.org/x/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    responses.ERROR(w, http.StatusUnprocessableEntity, err)
    return
  }
  user := models.User()
  err = json.Unmarshal(body, &user)
  if err != nil {
    responses.ERROR(w, http.StatusUnprocessableEntity, err)
    return
  }
  user.Prepare()
  err = user.Validate("login")
  if err != nil {
    responses.ERROR(w, http.StatusUnprocessableEntity, err)
    return
  }
  responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {
  var err error
  user := models.User{}
  err = server.DB.Debug().Model(models.User{}).Where("email = ?", emal).Take(&user).Error
  if err != nil {
    return "", err
  }
  err = models.VerifyPassword(user.Password, password)
  if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
    return "", err
  }
  return auth.CreateToken(user.ID)
}
