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

func GetCreated(id, prefix string) (int64, error) {
	timestr, err := baseconv.Convert(id[len(prefix):13 + len(prefix)], "abcdefghijklmnopqrstuvwxyz", baseconv.DigitsDec)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(timestr, 10, 0)
}

func New0() string {
	return generateID("00", 0)
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

func NewUserFieldID() string {
	return generateID("uf", 3)
}

func NewRequestID() string {
	return generateID("rq", 10)
}

func NewAutomationID() string {
	return generateID("at", 5)
}

func NewUserSessionID() string {
	return generateID("ss", 6)
}

func NewBillingID() string {
	return generateID("bi", 6)
}

func NewInvoiceID() string {
	return generateID("ic", 6)
}

func NewSubscriptionID() string {
	return generateID("sc", 6)
}

func NewCustomerID() string {
	return generateID("cr", 6)
}

func NewPaymentLogID() string {
	return generateID("pl", 6)
}

func NewPaymentMethodID() string {
	return generateID("pm", 3)
}
