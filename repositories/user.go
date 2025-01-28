// This file is for Database operations related to the user
package repositories

import (
	"database/sql"
	"fmt"
	"strconv"

	// "log"
	"encoding/json"

	"github.com/awaisamjad/whisp/backend/db"
	"github.com/awaisamjad/whisp/backend/internal"
	log "github.com/sirupsen/logrus"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connet to dB")
	}
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(SignUpInfo internal.SignUpRequest) error {
	// ? Check if username already exists in dB
	var user_exists bool
	var email_exists bool

	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", SignUpInfo.Username).Scan(&user_exists)
	if err != nil {
		log.Error("Failed to check if username already exists in dB")
		return fmt.Errorf("Internal Server Error")
	}
	if user_exists {
		return fmt.Errorf("User already exists")
	}

	// ? Check if email already exists in dB
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", SignUpInfo.Email).Scan(&email_exists)
	if err != nil {
		log.Error("Failed to check if email already exists in dB")
		return fmt.Errorf("Internal Server Errir")
	}
	if email_exists {
		return fmt.Errorf("Email is already in use")
	}

	// ?  Create the user
	query := `INSERT INTO users (username, first_name, last_name, email, password) VALUES (?, ?, ?, ?, ?)`
	_, err = r.db.Exec(query, SignUpInfo.Username, SignUpInfo.FirstName, SignUpInfo.LastName, SignUpInfo.Email, SignUpInfo.Password)
	if err != nil {
		log.Error("Failed to create user")
		return fmt.Errorf("Either username or email is already in use")
	}
	return nil
}

func (r *UserRepository) LogInUser(logInInfo internal.LogInRequest) (internal.LogInReturn, error) {

	var storedPassword, username, id string
	//? Check if email exists and get the stored password and username
	err := r.db.QueryRow("SELECT password, username, id FROM users WHERE email = ?", logInInfo.Email).Scan(&storedPassword, &username, &id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("Email is not registered")
			return internal.LogInReturn{Username: "", User_Id: ""}, fmt.Errorf("Email is not registered")
		}
		log.Error("Internal Server Error")
		return internal.LogInReturn{Username: "", User_Id: ""}, fmt.Errorf("Internal server error")
	}

	//? Check if the password matches
	if !internal.CheckPasswordAgainstPasswordHash(logInInfo.Password, storedPassword) {
		log.Error("Password does not match hashed password")
		return internal.LogInReturn{Username: "", User_Id: ""}, fmt.Errorf("Incorrect Password")
	}

	logInReturn := internal.LogInReturn{
		User_Id:       id,
		Username: username,
	}
	return logInReturn, nil
}

func (r *UserRepository) CreatePost(createPostInfo internal.CreatePostRequest) error {
	query := `INSERT INTO posts (content, user_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, createPostInfo.Content, createPostInfo.User_Id)
	if err != nil {
		log.Error("Failed to insert user post in posts table")
		return fmt.Errorf("Failed to create post")
	}

	return nil
}

func (r *UserRepository) GetPosts() ([]internal.Post, error) {
	var posts []internal.Post
	query := `SELECT * FROM posts;`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Error("Failed to get posts from the database")
		return nil, fmt.Errorf("Failed to get posts")
	}
	defer rows.Close()

	for rows.Next() {
		var post internal.Post
		// tags and image_content are stored as JSON so handled differently
		var image_content string
		var tags string

		err = rows.Scan(
			&post.Id,
			&post.User_Id,
			&post.Username,
			&post.Avatar,
			&post.Text_Content,
			&image_content,
			&tags,
			&post.Comment_Count,
			&post.Retweet_Count,
			&post.Like_Count,
			&post.Created_At,
			&post.Updated_At)

		if err != nil {
			log.Error("Failed to scan post row : ", err)
			return nil, fmt.Errorf("Failed to get posts")
		}

		if tags != "" {
			err := json.Unmarshal([]byte(tags), &post.Tags)
			if err != nil {
				log.Error(err)
			}
		} else {
			log.Warn("Tags are empty")
		}

		if image_content != "" {
			err := json.Unmarshal([]byte(image_content), &post.Image_Content)
			if err != nil {
				log.Error(err)
			}
		} else {
			log.Warn("Image Content is empty")
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		log.Error("Error iterating over rows")
		return nil, fmt.Errorf("Failed to get posts")
	}
	return posts, nil
}

func (r *UserRepository) GetPostByID(id string) (internal.Post, error) {
	var post internal.Post
	query := `SELECT * FROM posts WHERE id=?;`
	int_id, err := strconv.Atoi(id) 
	if err != nil {
		log.Error("Failed to convert id type string to int")
	}
	row := r.db.QueryRow(query, int_id)
	// tags and image_content are stored as JSON so handled differently
	var image_content string
	var tags string

	err = row.Scan(
		&post.Id,
		&post.User_Id,
		&post.Username,
		&post.Avatar,
		&post.Text_Content,
		&image_content,
		&tags,
		&post.Comment_Count,
		&post.Retweet_Count,
		&post.Like_Count,
		&post.Created_At,
		&post.Updated_At)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("No post found with the given ID")
			return internal.Post{}, fmt.Errorf("No post found with the given ID")
		}
		log.Error("Failed to scan post row : ", err)
		return internal.Post{}, fmt.Errorf("Failed to get post")
	}

	if tags != "" {
		err := json.Unmarshal([]byte(tags), &post.Tags)
		if err != nil {
			log.Error(err)
		}
	} else {
		log.Warn("Tags are empty")
	}

	if image_content != "" {
		err := json.Unmarshal([]byte(image_content), &post.Image_Content)
		if err != nil {
			log.Error(err)
		}
	} else {
		log.Warn("Image Content is empty")
	}

	return post, nil
}

