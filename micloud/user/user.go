package usermgr

import "github.com/clouderhem/micloud/micloud/user/user"

func GetLoginUrl() (string, error) {
	return user.GetLoginUrl()
}
