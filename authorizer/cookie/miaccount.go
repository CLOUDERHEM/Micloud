package cookie

var (
	MiaccountCookieFilepath = "/tmp/micloud/.miaccount_cookie"
)

func GetMiaccountCookie() string {
	return getCookieFromFile(MiaccountCookieFilepath)
}
