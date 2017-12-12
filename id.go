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
	nowstr := strconv.FormatInt(time.Now().UnixNano(), 10)
	timestr, _ := baseconv.Convert(nowstr, baseconv.DigitsDec, "abcdefghijklmnopqrstuvwxyz")
  return sign + timestr + randstr.RandomString(randomfactor, "abcdefghijklmnopqrstuvwxyz")
}

func NewAgentGroupID() string {
	return generateID("gr", 3)
}

func NewAccountID() string {
	return generateID("ac", 5)
}

func NewAgentID() string {
	return generateID("ag", 3)
}

func NewInvitationID() string {
	return generateID("iv", 4)
}

func NewEventID() string {
	return generateID("ev", 10)
}

func NewConversationID() string {
	return generateID("cs", 3)
}

func NewScheduleItemID() string {
	return generateID("si", 20)
}

func NewWebsendID() string {
	return generateID("wd", 20) // websend
}

// NewAuthToken return new oauth2 authorization token
func NewAuthToken() string {
	return generateID("au", 20)
}

func NewRefreshToken() string {
	return generateID("rt", 30)
}

func NewRuleID() string {
	return generateID("ru", 6)
}

func NewClientID() string {
	return generateID("cl", 6)
}

func NewErrorID() string {
	return generateID("er", 6)
}

func NewWebhookID() string {
	return generateID("wh", 6)
}

func NewWsConnID(partition int32) string {
	return strconv.Itoa(int(partition)) + "p" + generateID("ws", 20)
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
	return generateID("us", 6)
}

func NewTagID() string {
	return generateID("tg", 5)
}

func NewCannedResponseID() string {
	return generateID("cn", 5)
}

func NewFileID() string {
	return generateID("fi", 5)
}

func NewButtonID() string {
	return generateID("bt", 8)
}

func NewChallengeID() string {
	return generateID("ch", 40)
}

func NewSegmentationID() string {
	return generateID("sg", 5)
}

func NewSubizID() string {
	return generateID("su", 8)
}
