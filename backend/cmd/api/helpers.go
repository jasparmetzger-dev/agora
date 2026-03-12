package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/jasparmetzger-dev/agora/cmd/database"
)

//USER HELPERS

func ValidateUUID(c *gin.Context) (pgtype.UUID, error) {
	userIdString := c.MustGet("UserId").(string)
	var id pgtype.UUID
	err := id.Scan(userIdString)
	return id, err
}

func MakeUserFromHeader(q *db.Queries, c *gin.Context) (db.User, error, int) { //also return httpStatus
	id, err := ValidateUUID(c)
	if err != nil {
		return db.User{}, err, 401
	}
	user, err := q.GetUserById(c, id)
	if err != nil {
		return db.User{}, err, 500
	}
	return user, nil, 200
}

func UserUpdateHelper(q *db.Queries, c *gin.Context, u db.User) (db.User, error) {
	params := db.UpdateUserByIdParams{}
	params.ID = u.ID
	params.Username = u.Username
	params.Email = u.Email
	params.PasswordHash = u.PasswordHash
	return q.UpdateUserById(c, params)
}

// POST HELPERS
func ValidatePostUUID(c *gin.Context) (pgtype.UUID, error) {
	var idString string = c.Param("id")
	var id pgtype.UUID
	err := id.Scan(idString)
	return id, err
}

func MakePostFromPath(q *db.Queries, c *gin.Context) (db.Post, error, int) { //also return httpStatus
	id, err := ValidatePostUUID(c)
	if err != nil {
		return db.Post{}, err, 401
	}
	post, err := q.GetPostById(c, id)
	if err != nil {
		return db.Post{}, err, 500
	}
	return post, nil, 200
}
