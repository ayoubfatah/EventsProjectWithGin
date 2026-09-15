package models

import (
	"errors"
	"gin-quickstart/db"
	"gin-quickstart/utils"
)

type User struct {
	Id int64  `json:"id"`
	FirstName string `json:"firstName" binding:"required"` 
	SecondName string `json:"secondName" binding:"required"` 
	UserName string `json:"userName" binding:"required"` 
	Email    string `json:"email" binding:"required"` 
	Password string `json:"password" binding:"required"`  
}


type SignupRequest struct {
	Id		int64
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
 
func (u *User) Save() (error) {
	query := `INSERT INTO users(firstName, secondName, userName, email, password) 
			  VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err 
	}
	defer stmt.Close()
	passwordHashed , err:=   utils.HashPassword(u.Password)
	if err != nil {
		return err
	}


	result, err := stmt.Exec(
		u.FirstName,
		u.SecondName,
		u.UserName,
		u.Email,
		passwordHashed,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.Id = id	
	return nil 
}





func  GetUsers() ([]User, error){
	query := `SELECT id , firstName , secondName  , userName, email, password  FROM users `
	rows , err := db.DB.Query(query)
	 if(err != nil){
		return  nil, err
	 }
	defer rows.Close()


	var users []User
	for rows.Next(){
		var user User 
		// err := rows.Scan(&user.Id, &user.FirstName , &user.SecondName , &user.Email  , &user.UserName )
		err := rows.Scan(&user.Id ,&user.FirstName , &user.SecondName , &user.UserName ,&user.Email , &user.Password)
		 if(err != nil){
		return  nil,err 
		 } 
		users = append(users, user)
	}
	return users , err
}




func (u *SignupRequest) ValidateCredentials() error{
	query := `SELECT  id ,password FROM users WHERE email = ? `
	row :=   db.DB.QueryRow(query, u.Email)	
	

	var retrievedPassword string 


	err := row.Scan(&u.Id , &retrievedPassword)
   if err != nil {
	return  err 
   }
   passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid{
		return  errors.New("credentials invalid")
	}
	return nil 
}


func GetUserByID(userId int64) (User, error) {


    const query = `
        SELECT id, email, firstName, secondName, userName
        FROM users
        WHERE id = ?
    `

    row := db.DB.QueryRow(query, userId)

    var user User

    err := row.Scan(
        &user.Id,
        &user.Email,
        &user.FirstName,
        &user.SecondName,
        &user.UserName,
    )

    if err != nil {

        return User{}, err
    }



    return user, nil
}