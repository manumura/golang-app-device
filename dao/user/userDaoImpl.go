package userdao

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/manumura/golang-app-device/config"
	"github.com/manumura/golang-app-device/model/user"
	"github.com/manumura/golang-app-device/security"
)

// UserDaoImpl : implementation for DB operations on user
type UserDaoImpl struct {
}

// GetUser : retrieve user by id from the database
func (ud UserDaoImpl) GetUser(id int) (usermodel.User, error) {

	// user := usermodel.User{}

	if id == 0 {
		return usermodel.User{}, errors.New("id cannot be empty")
	}

	// row := config.Database.QueryRow("SELECT u.id, u.username, u.password, u.last_name, u.first_name FROM app_user u WHERE u.id = $1", id)
	// err := row.Scan(&user.ID, &user.Username, &user.Password, &user.LastName, &user.LastName)

	query := "SELECT u.id, u.username, u.last_name, u.first_name, r.id, r.role_name "
	query += "FROM app_user u "
	query += "LEFT JOIN user_role ur ON ur.user_id = u.id "
	query += "LEFT JOIN role r ON r.id = ur.role_id  "
	query += "WHERE u.id = $1"

	rows, err := config.Database.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var userID, roleID int
	var userName, userLastName, userFirstName, roleName string
	var roles []usermodel.Role
	for rows.Next() {
		err := rows.Scan(&userID, &userName, &userLastName, &userFirstName, &roleID, &roleName)
		if err != nil {
			log.Fatal(err)
		}

		roles = append(roles, usermodel.Role{ID: roleID, Name: roleName})
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return usermodel.User{}, err
	}

	user := usermodel.User{
		ID:        userID,
		Username:  userName,
		LastName:  userLastName,
		FirstName: userFirstName,
		Roles:     roles,
	}

	log.Println(user)
	return user, nil
}

// GetUserByUsername : retrieve user by username from the database
func (ud UserDaoImpl) GetUserByUsername(username string) (usermodel.User, error) {

	user := usermodel.User{}

	if username == "" {
		return user, errors.New("username cannot be empty")
	}

	row := config.Database.QueryRow("SELECT u.id, u.username, u.password, u.last_name, u.first_name FROM app_user u WHERE upper(u.username) = $1", strings.ToUpper(strings.Trim(username, " ")))

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.LastName, &user.LastName)
	if err != nil {
		log.Println(err)
		return user, err
	}

	log.Println(user)
	return user, nil
}

// CheckUsernameUnique : check username is unique
func (ud UserDaoImpl) CheckUsernameUnique(username string) (bool, error) {

	user := usermodel.User{}

	if username == "" {
		return false, errors.New("username cannot be empty")
	}

	row := config.Database.QueryRow("SELECT u.id FROM app_user u WHERE upper(u.username) = $1", strings.ToUpper(strings.Trim(username, " ")))

	err := row.Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		log.Println(err)
		return false, err
	}

	return false, nil
}

// Create : create one user
func (ud UserDaoImpl) Create(u usermodel.User) (usermodel.User, error) {

	log.Println("Create user")

	result := usermodel.User{}

	if u.Username == "" || u.Password == "" {
		return result, errors.New("parameters cannot be empty")
	}

	// Hash password to argon2
	hashedPassword, err := security.GenerateHash(u.Password)
	if err != nil {
		log.Println("Cannot hash password")
	}

	tx, err := config.Database.Begin()
	if err != nil {
		log.Println(err)
		return result, err
	}

	// execute insert on user table
	stmt, err := tx.Prepare("INSERT INTO app_user (username, password, created_date_time, is_active, first_name, last_name) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return result, err
	}
	defer stmt.Close()

	var userid int
	if err = stmt.QueryRow(u.Username, hashedPassword, time.Now(), true, u.FirstName, u.LastName).Scan(&userid); err != nil {
		tx.Rollback()
		log.Println(err)
		return result, err
	}
	log.Println("userid: ", userid)

	// execute insert on user role table
	stmt, err = tx.Prepare("INSERT INTO user_role (user_id, role_id) SELECT $1, id FROM role WHERE role_name = 'STANDARD_USER'")
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return result, err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userid); err != nil {
		tx.Rollback()
		log.Println(err)
		return result, err
	}

	// commit transaction
	if err = tx.Commit(); err != nil {
		log.Println(err)
		return result, errors.New("user cannot be saved")
	}

	result, err = ud.GetUser(userid)

	log.Println(result)

	return result, err
}
