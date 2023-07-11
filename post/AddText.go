package post

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddText(c *gin.Context) {
	fmt.Fprintf(c.Writer, "something")
}
