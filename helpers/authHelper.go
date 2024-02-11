package helpers
import
	(
	"errors"	
	"github.com/gin-gonic/gin"
)



func MatchUserTypetoUid(c *gin.Context, userId string)(err error){
	userType := c.GetString("user_type")
	uid:= c.GetString("uid")
	err = nil
	if userType =="USER" && uid != userId{
		err = errors.New("UnAuthorized")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

func CheckUserType(c* gin.Context, role string)(err error){
	userType := c.GetString("user_type")
	err = nil
	if userType !=role{
		err = errors.New("UnAuthorized")
		return err
	}
	return err;
}