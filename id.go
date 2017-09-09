package ID

import (
	"time"
	"github.com/thanhpk/randstr"
	"github.com/thanhpk/baseconv"
	"strconv"
)

// New return new random ID
func New() string {
	newid := generateID("", 2)
	return newid
}

func generateID(sign string, randomfactor int) string {
	nowstr := strconv.FormatInt(time.Now().Unix(), 10)
	timestr, _ := baseconv.Convert(nowstr, baseconv.DigitsDec, baseconv.Digits62)
  return timestr + randstr.RandomString(randomfactor) + sign
}

func NewAgentGroupID() string {
	return generateID("gr", 16)
}

func NewAccountID() string {
	return generateID("ac", 16)
}

func NewAgentID() string {
	return generateID("ag", 16)
}

func NewInvitationID() string {
	return generateID("iv", 16)
}

func NewEventID() string {
	return generateID("ev", 16)
}

func NewOutEventID() string {
	return generateID("oe", 16)
}

func NewChatEventID() string {
	return generateID("ce", 16)
}

func NewConverstationID() string {
	return generateID("cs", 16)
}

func NewChatID() string {
	return generateID("cs", 16)
}

func NewScheduleItemID() string {
	return generateID("si", 20)
}

func NewWebsendID() string {
	return generateID("Wd", 20) // websend
}

// NewAuthToken return new oauth2 authorization token
func NewAuthToken() string {
	return generateID("au", 20)
}

func NewRefreshToken() string {
	return generateID("rt", 30)
}

func NewClientID() string {
	return generateID("cl", 16)
}

func NewErrorID() string {
	return generateID("er", 5)
}

func NewWebhookID() string {
	return generateID("wh", 8)
}
