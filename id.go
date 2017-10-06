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
	timestr, _ := baseconv.Convert(nowstr, baseconv.DigitsDec, "abcdefghijklmnopqrstuvwxyz")
  return sign + timestr + randstr.RandomString(randomfactor, "abcdefghijklmnopqrstuvwxyz")
}

func NewAgentGroupID() string {
	return generateID("gr", 16)
}

func NewAccountID() string {
	return generateID("ac", 32)
}

func NewAgentID() string {
	return generateID("ag", 16)
}

func NewInvitationID() string {
	return generateID("iv", 16)
}

func NewEventID() string {
	return generateID("ev", 24)
}

func NewOutEventID() string {
	return generateID("oe", 16)
}

func NewConverstationID() string {
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

func NewRuleID() string {
	return generateID("ru", 10)
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

func NewWsConnID(partition int32) string {
	return strconv.Itoa(int(partition)) + "p" + generateID("ws", 30)
}

func GetPartitionFromWsConnID(id string) int32 {
	for i := 0; i < len(id); i++ {
		if id[i] == "p"[0] {
			par := id[:i]
			partition, err := strconv.Atoi(par)
			if err != nil {
				return -1
			}
			return int32(partition)
		}
	}
	return -1
}

func NewUserID() string {
	return generateID("us", 32)
}

func NewTagID() string {
	return generateID("tg", 2)
}

func NewCannedResponseID() string {
	return generateID("cn", 2)
}
