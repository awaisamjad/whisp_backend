// This file contains objects such as structs that are used throughout the project
package internal

var StandardRoutes = []string{
	"explore",
	"search",
}

type Post struct {
	Id            int    `json:"id"`
	User_Id       int    `json:"user_id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Text_Content  string `json:"text_content"`
	Image_Content string `json:"image_content,omitempty"`  // Could be extended to an array for multiple images
	Tags          string `json:"tags,omitempty"` // Tags for the post
	Comment_Count int    `json:"comment_count"`  // Count of comments
	Retweet_Count int    `json:"retweet_count"`  // Count of retweets
	Like_Count    int    `json:"like_num"`
	Created_At    string `json:"created_at"`
	Updated_At    string `json:"updated_at"`
}

// ? Following and Followers are a list of user id's (ints) of those they are following/followed by
type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Avatar     string `json:"avatar"`
	Posts      []Post `json:"posts"`
	Feed       []Post `json:"feed"`
	Following  []int  `json:"following"`
	Followers  []int  `json:"followers"`
	Bio        string `json:"bio"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}

type SignUpRequest struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// What the dB gives back when the user is logged in
type LogInReturn struct {
	Username string `json:"username"`
	User_Id       string `json:"user_id"`
}

type CreatePostRequest struct {
	Content string `json:"content"`
	User_Id string `json:"user_id"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

type SuccessResponse struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}
