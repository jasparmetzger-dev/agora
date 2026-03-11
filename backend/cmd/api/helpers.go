package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func ValidateUUID(c *gin.Context) (pgtype.UUID, error) {
	userIdString := c.MustGet("UserId").(string)
	var id pgtype.UUID
	err := id.Scan(userIdString)
	return id, err
}
