package idgen

import (
	"fmt"
	"testing"
)

func TestCreated(t *testing.T) {
	if c, err := GetCreated("evqcxzkihetmvosjgvhctilzx", "ev"); err != nil {
		t.Fatal(err)
	} else {
		if c != 1537588731096863554 {
			t.Fatalf("should be %d, got %d", 1537588731096863554, c)
		}
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
		{"template", NewTemplateID()},
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
		{"user note", NewUserNoteID()},
		{"ticket", NewTicketID()},
		{"pipeline", NewPipelineID()},
		{"stage", NewStageID()},
		{"currency", NewCurrencyID()},
		{"exchange rate", NewExchangeRateID()},
	}
	c, err := GetCreated(New0(), "00")
	fmt.Printf("%d, %v\n", c, err)

	for i := range ids {
		fmt.Printf("%15s: %s\n", ids[i].name, ids[i].id)
	}
}

func BenchmarkGenID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewEventID()
	}
}

func TestValidateID(t *testing.T) {
	ids := []struct {
		name, id string
		f        func(id string) bool
	}{
		{"user", NewUserID(), IsUserID},
		{"user", "usqdbezwcnrbtmbiocsip", IsUserID},
		{"account", NewAccountID(), IsAccountID},
		{"account", "acqdbeyvxzinwalityyn", IsAccountID},
		{"tag", NewTagID(), IsTagID},
		{"tag", "tgqdbeyvxzimrxrqqjyh", IsTagID},
		{"group", NewAgentGroupID(), IsAgentGroupID},
		{"group", "grqdbezhjyzvdrhhul", IsAgentGroupID},
		{"agent", NewAgentID(), IsAgentID},
		{"agent", "agqdbezhjyzvvzjszi", IsAgentID},
		{"event", NewEventID(), IsEventID},
		{"event", "evqdbezhjyzwmrglcrzzrnadt", IsEventID},
		{"file", NewFileID(), IsFileID},
		{"file", "fiqdbezwcnrckjnhokhf", IsFileID},
		{"button", NewButtonID(), IsButtonID},
		{"button", "btqdbezwcnrdbdieiktpgwz", IsButtonID},
		{"conversation", NewConversationID(), IsConversationID},
		{"conversation", "csqdbezwcnrdsclxgr", IsConversationID},
		{"invitation", NewInvitationID(), IsInvitationID},
		{"invitation", "ivqdbezwcnreisdbghf", IsInvitationID},
		{"scheduler item", NewScheduleItemID(), IsScheduleItemID},
		{"scheduler item", "siqdbezwcnrezjwfjtwuaqvjluileucrxmt", IsScheduleItemID},
		{"web send", NewWebsendID(), IsWebsendID},
		{"web send", "wdqdbezwcnrfsiczskacihukneqfmokwubm", IsWebsendID},
		{"auth token", NewAuthToken(), IsAuthToken},
		{"auth token", "auqdbezwcnrgjuaknybgomzggkomxigkfch", IsAuthToken},
		{"refresh token", NewRefreshToken(), IsRefreshToken},
		{"refresh token", "rtqdbezwcnrhbbfcmnahxowjwwpsydsgxjvbmedqipdfy", IsRefreshToken},
		{"rule", NewRuleID(), IsRuleID},
		{"rule", "ruqdbezwcnrhszlsbpbjp", IsRuleID},
		{"client", NewClientID(), IsClientID},
		{"client", "clqdbezwcnrijqvkwigot", IsClientID},
		{"error", NewErrorID(), IsErrorID},
		{"error", "erqdbezwcnrjaklhawqau", IsErrorID},
		{"canned response", NewCannedResponseID(), IsCannedResponseID},
		{"canned response", "cnqdbezwcnrjrbeceozl", IsCannedResponseID},
		{"template", NewTemplateID(), IsTemplateID},
		{"template", "tpqdbezwcnrjrbeceozl", IsTemplateID},
		{"ws conn", NewWsConnID(10), IsWsConnID},
		{"ws conn", "10pwsqdbezwcnrkhuewsnoicahfgmsjsiixcsd", IsWsConnID},
		{"challenge", NewChallengeID(), IsChallengeID},
		{"challenge", "chqdbezwcnrkzjmmqwksdurdtmnijfcojwdeetnkrwcqmnhwtgzqyuf", IsChallengeID},
		{"segmentation", NewSegmentationID(), IsSegmentationID},
		{"segmentation", "sgqdbezwcnrlryjdunoo", IsSegmentationID},
		{"subiz", NewSubizID(), IsSubizID},
		{"subiz", "suqdbezwcnrmiskvudpjlfi", IsSubizID},
		{"user field", NewUserFieldID(), IsUserFieldID},
		{"user field", "ufqdbezwcnrmzoskmu", IsUserFieldID},
		{"request", NewRequestID(), IsRequestID},
		{"request", "rqqdbezwcnrnqclbwgztgmlab", IsRequestID},
		{"automation", NewAutomationID(), IsAutomationID},
		{"automation", "atqdbezwcnrohdwmqloj", IsAutomationID},
		{"user session", NewUserSessionID(), IsUserSessionID},
		{"user session", "ssqdbezwcnrozhlymatud", IsUserSessionID},
		{"invoice", NewInvoiceID(), IsInvoiceID},
		{"invoice", "icqdbezwcnrprixgmowtl", IsInvoiceID},
		{"billing", NewBillingID(), IsBillingID},
		{"billing", "biqdbezwcnrqiddoodced", IsBillingID},
		{"subscription", NewSubscriptionID(), IsSubscriptionID},
		{"subscription", "scqdbezwcnrqyuohnychx", IsSubscriptionID},
		{"customer", NewCustomerID(), IsCustomerID},
		{"customer", "crqdbezwcnrrpjzcdiuzb", IsCustomerID},
		{"payment log", NewPaymentLogID(), IsPaymentLogID},
		{"payment log", "plqdbezwcnrsfzuxwugar", IsPaymentLogID},
		{"payment method", NewPaymentMethodID(), IsPaymentMethodID},
		{"payment method", "pmqdbezwcnrswpyiqm", IsPaymentMethodID},
		{"attribute", NewAttributeID(), IsAttributeID},
		{"attribute", "abqdbezwcnrtoocyzs", IsAttributeID},
		{"idempotency key", NewIdempotencyKey(), IsIdempotencyKey},
		{"idempotency key", "ikqdbezwcnruezpqgzeaugyohxmjfdbjbfz", IsIdempotencyKey},
		{"comment", NewPaymentCommentID(), IsPaymentCommentID},
		{"comment", "cmqdbezwcnruwifmecaxrzqrc", IsPaymentCommentID},
		{"user note", NewUserNoteID(), IsUserNoteID},
		{"user note", "ntqdfmydsqqpeevtfeuhntbwd", IsUserNoteID},
		{"ticket", NewTicketID(), IsTicketID},
		{"ticket", "tkqdgcajtiakmkflxaneneknj", IsTicketID},
		{"pipeline", NewPipelineID(), IsPipelineID},
		{"pipeline", "plqeeknatxuphbskrhlr", IsPipelineID},
		{"stage", NewStageID(), IsStageID},
		{"stage", "stqeeknatxupvwilqs", IsStageID},
		{"currency", "crqefeeclybkbzsjsjza", IsCurrencyID},
		{"exchange rate", "exqefeeclybkquiinjkv", IsExchangeRateID},
	}

	for _, v := range ids {
		isValid := v.f(v.id)
		if !isValid {
			fmt.Printf("%s - %s: %t\n", v.name, v.id, isValid)
			t.FailNow()
		}
	}
}
