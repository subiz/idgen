package ID

import (
	"time"
	"github.com/thanhpk/randstr"
	"github.com/thanhpk/baseconv"
	"strconv"
)

func generateID(sign string, randomfactor int) string {
		timestr, _ := baseconv.Convert(strconv.FormatInt(time.Now().UnixNano(), 10),
		baseconv.DigitsDec,
    baseconv.Digits62)
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

func NewOutEventID() string {
	return generateID("oe", 16)
}

func NewChatEventID() string {
	return generateID("ce", 16)
}

func NewChatID() string {
	return generateID("cs", 16)
}
