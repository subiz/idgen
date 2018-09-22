package idgen

import (
	"fmt"
	"testing"
	"time"

	"github.com/thanhpk/baseconv"
)

func TestCreated(t *testing.T) {
	nowstr := "14324234234"
	timestr, _ := baseconv.Convert(nowstr, baseconv.DigitsDec, "abcdefghijklmnopqrstuvwxyz")
	println("xxxxxx", timestr)

	created, err := GetCreated("evaaaaaaaaaaaaaaaaaaa", "ev")
	println(time.Unix(0, created).Format(time.RFC3339Nano), err)
	if _, err := GetCreated("evqcxzkihetmvosjgvhctilzx", "ev"); err != nil {
		//"evqalufuxphqubrufrnqbgcgf"
		//"00qalufuxpijiao"
		t.Fatal(err)
	}
}

func TestId(t *testing.T) {
	ids := []struct {
		name, id string
	}{
		{"tag", NewTagID()},
		{"account", NewAccountID()},
		{"group", NewAgentGroupID()},
		{"agent", NewAgentID()},
		{"event", NewEventID()},
		{"user", NewUserID()},
		{"file", NewFileID()},
		{"button", NewButtonID()},
		{"conversation", NewConversationID()},
		{"invitation", NewInvitationID()},
		{"scheduler item", NewScheduleItemID()},
		{"web send", NewWebsendID()},
		{"auth token", NewAuthToken()},
		{"refresh token", NewRefreshToken()},
		{"rule", NewRuleID()},
		{"client", NewClientID()},
		{"error", NewErrorID()},
		{"canned response", NewCannedResponseID()},
		{"ws conn", NewWsConnID(10)},
		{"challenge", NewChallengeID()},
		{"segmentation", NewSegmentationID()},
		{"subiz", NewSubizID()},
		{"user field", NewUserFieldID()},
		{"request", NewRequestID()},
		{"automation", NewAutomationID()},
		{"user session", NewUserSessionID()},
		{"invoice", NewInvoiceID()},
		{"billing", NewBillingID()},
		{"subscription", NewSubscriptionID()},
		{"customer", NewCustomerID()},
		{"payment log", NewPaymentLogID()},
		{"payment method", NewPaymentMethodID()},
		{"attribute", NewAttributeID()},
		{"idempotency key", NewIdempotencyKey()},
		{"comment", NewPaymentCommentID()},
		{"zero", New0()},
	}
	c, err := GetCreated(New0(), "00")
	fmt.Printf("%d, %v\n", c, err)

	for i := range ids {
		fmt.Printf("%15s: %s\n", ids[i].name, ids[i].id)
	}
}

func TestValidateID(t *testing.T) {
	ids := []struct {
		name, id string
		f        func(id string) bool
	}{
		{"user", NewUserID(), IsUserID},
		{"account", NewAccountID(), IsAccountID},
	}

	for _, v := range ids {
		isValid := v.f(v.id)
		fmt.Printf("%s - %s: %t\n", v.name, v.id, isValid)
		if !isValid {
			t.FailNow()
		}
	}
}
