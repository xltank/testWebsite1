package midware

import (
	"log"
	"websiteGin/db"
	"websiteGin/error"
	"websiteGin/model"
	"websiteGin/res"
	"websiteGin/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, _ := c.Cookie("token")

		// log.Println("Auth, cookie: " + cookie)
		_, claims, err := utils.ParseToken(cookie)

		if err != nil {
			log.Println("get cookie error:", err)
			res.SendParamError(c, error.CodeUserLogin, "Please Login")
			c.Abort()
			return // `return` not works. To return before other handlers, use Abortxxx().
		}
		log.Println("UserId from token: ", claims.ID)
		var user model.User
		r := db.Db.First(&user, claims.ID)
		if r.Error != nil {
			res.SendParamError(c, 0, r.Error.Error())
			return
		}

		user.Pass = ""
		c.Set("user", user)

		c.Next()
	}
}
