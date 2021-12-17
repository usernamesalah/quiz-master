package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"text/template"
	"time"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/usernamesalah/quiz-master/internal/config"
	"github.com/usernamesalah/quiz-master/internal/constants"
	"github.com/usernamesalah/quiz-master/internal/datatransfers"
	"golang.org/x/crypto/bcrypt"
)

const (
	encodingBase        = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	orderedEncodingBase = "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
	FormatDate1         = "2006-01-02 15:04:05"
	FormatDate2         = "Monday, 02 January 2006 15:04"
	FormatDate3         = "02 January 2006"
	FormatDate4         = "2006-01-02T15:04:05-0700"
	FormatDate5         = "01/02/2006"
	IssuerJWT           = "usernamesalah"
)

var letters = []rune(encodingBase)

func ComparePasswords(hashedPwd string, plainPwd []byte) (err error) {
	byteHash := []byte(hashedPwd)
	err = bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		err = errors.New("400: Your password is incorrect")
		return
	}
	return
}

func HashAndSalt(pwd []byte) (hashPwd string, err error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return
	}

	hashPwd = string(hash)
	return
}

func GetDateByUnixTime(format string, unixtime int) string {
	if unixtime <= (7 * 3600) {
		return ""
	}
	result := time.Unix(int64(unixtime), 0).Format(format)
	return result
}

func GenerateSha256Hash(value string) string {
	hasher := sha256.New()
	hasher.Write([]byte(value))
	signature := hex.EncodeToString(hasher.Sum(nil))
	return signature
}

func CalculateOffset(limit, currentPage int) (result int) {
	result = (currentPage - 1) * limit
	return
}

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Now() int {
	return int(time.Now().Unix())
}

func BuildTemplateEmail(value string, data interface{}) (content string, err error) {
	t, err := template.New("email").Parse(value)
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// func SendEmail(content, subject, destination string) (err error) {
// 	// dont send if email as tester
// 	if destination == "" || strings.Contains(destination, constants.PostFixEmailTest) {
// 		return
// 	}

// 	m := gomail.NewMessage()

// 	m.SetBody("text/html", content)
// 	m.SetHeaders(map[string][]string{
// 		"From":    {m.FormatAddress(conf.AppConfig.SMTPSenderEmail, conf.AppConfig.SMTPSenderName)},
// 		"To":      {destination},
// 		"Subject": {subject},
// 	})

// 	d := gomail.NewPlainDialer(
// 		conf.AppConfig.SMTPHost,
// 		conf.AppConfig.SMTPPort,
// 		conf.AppConfig.Secrets.SMTPUser,
// 		conf.AppConfig.Secrets.SMTPPassword)

// 	if err = d.DialAndSend(m); err != nil {
// 		return
// 	}

// 	return
// }

func SetPaginator(limit, page int, cnt int64, c echo.Context) (pageData *datatransfers.PageData) {
	p := pagination.NewPaginator(c.Request(), limit, cnt)
	// because this beego pagination package use 'p' instad of 'page' for page param
	// the p.Pages() is invalid, so we use our own hasNext and currentPage
	hasNext := page < p.PageNums()
	pageData = &datatransfers.PageData{
		HasNext:     hasNext,
		TotalData:   p.Nums(),
		TotalPages:  p.PageNums(),
		CurrentPage: page,
		Limit:       p.PerPageNums,
	}
	return
}

func ReturnInvalidResponse(httpcode int, message string) error {

	errData := datatransfers.ErrorData{
		Code:    httpcode,
		Status:  httpcode,
		Message: message,
	}

	responseBody := datatransfers.Response{
		Success: false,
		Error:   &errData,
	}
	return echo.NewHTTPError(httpcode, responseBody)
}

func GetUserID(c echo.Context) (userID int64, err error) {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	userID, err = strconv.ParseInt(claims["uid"].(string), 10, 64)
	return
}

func CleanPhoneNumber(phoneNumber string) (result string, err error) {
	reg := regexp.MustCompile("[^Z0-9]+")

	result = "+" + reg.ReplaceAllString(phoneNumber, "")
	if len(result) < len(constants.PrefixPhoneNumber)+2 {
		// the min len of phone number to be useful (eg: to get its provider)
		err = errors.New("400: phone number is too short")
		return "", err
	}

	if result[:len(constants.PrefixPhoneNumber)] == constants.PrefixPhoneNumber {
		return
	}

	// just in case if fe still fetch when the user input the wrong number
	if result[:len(constants.PrefixPhoneNumber)] != "+620" {
		return
	}

	// example of phone number from cotter +6208XXXXXXXXXX
	result = constants.PrefixPhoneNumber + result[len(constants.PrefixPhoneNumber)+1:]
	return
}

func GenerateToken(userData *datatransfers.UserData) (token string, err error) {
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":     userData.UID,
		"isAdmin": userData.IsAdmin,
		"iss":     IssuerJWT,
		"sub":     IssuerJWT,
		"iat":     time.Now().Unix(),
	})

	token, err = rawToken.SignedString([]byte(config.AppConfig.JWTSecret))
	return
}
