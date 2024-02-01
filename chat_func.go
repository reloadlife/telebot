package telebot

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func (b *bot) Ban(chatID Recipient, userID int64, untilDate *int64, revokeMessages *bool) error {
	params := banChatMemberRequest{
		ChatID: chatID,
		UserID: userID,
	}

	if untilDate != nil {
		params.UntilDate = untilDate
	}

	if revokeMessages != nil {
		params.RevokeMessages = revokeMessages
	}

	_, err := b.sendMethodRequest(methodBanChatMember, params)
	return err
}

func (b *bot) Unban(chatID Recipient, userID int64, onlyIfBanned *bool) error {
	params := unbanChatMemberRequest{
		ChatID:       chatID,
		UserID:       userID,
		OnlyIfBanned: true, // to prevent users getting kicked when UnBan function is called on user that is not already banned.
	}

	if onlyIfBanned != nil {
		params.OnlyIfBanned = *onlyIfBanned
	}

	_, err := b.sendMethodRequest(methodUnbanChatMember, params)
	return err
}

func (b *bot) Restrict(chatID Recipient, userID int64, permissions ChatPermissions, useIndependentChatPermissions *bool, untilDate *time.Duration) error {
	params := restrictChatMemberRequest{
		ChatID:                        chatID,
		UserID:                        userID,
		Permissions:                   permissions,
		UseIndependentChatPermissions: false,
		UntilDate:                     0,
	}

	if useIndependentChatPermissions != nil {
		params.UseIndependentChatPermissions = *useIndependentChatPermissions
	}

	if untilDate != nil {
		params.UntilDate = time.Now().Add(*untilDate).Unix()
	}

	_, err := b.sendMethodRequest(methodRestrictChatMember, params)
	return err
}

func (b *bot) Promote(chatID Recipient, userID int64, roles ...ChatMemberPermission) error {
	params := promoteChatMemberRequest{
		ChatID: chatID,
		UserID: userID,
	}
	paramsValue := reflect.ValueOf(&params).Elem()
	for _, role := range roles {
		fieldValue := paramsValue.FieldByName(role.String())

		if fieldValue.IsValid() && fieldValue.Kind() == reflect.Bool {
			fieldValue.SetBool(true)
		} else {
			return fmt.Errorf("unsupported role: %s", role.String())
		}
	}

	_, err := b.sendMethodRequest(methodPromoteChatMember, params)
	return err
}

func (b *bot) SetChatAdministratorCustomTitle(chatID Recipient, userID int64, customTitle string) error {
	_, err := b.sendMethodRequest(methodSetChatAdministratorCustomTitle, setChatAdministratorCustomTitleRequest{
		ChatID:      chatID,
		UserID:      userID,
		CustomTitle: customTitle,
	})
	return err
}

func (b *bot) BanChatSenderChat(chatID Recipient, userID int64) error {
	_, err := b.sendMethodRequest(methodBanChatMember, banChatSenderChatRequest{
		ChatID:       chatID,
		SenderChatID: userID,
	})
	return err
}

func (b *bot) UnbanChatSenderChat(chatID Recipient, userID int64) error {
	_, err := b.sendMethodRequest(methodUnbanChatMember, unbanChatSenderChatRequest{
		ChatID:       chatID,
		SenderChatID: userID,
	})
	return err
}

func (b *bot) SetChatPermissions(chatID Recipient, permissions ChatPermissions, useIndependentChatPermissions *bool) error {
	_, err := b.sendMethodRequest(methodSetChatPermissions, setChatPermissionsRequest{
		ChatID:                        chatID,
		Permissions:                   permissions,
		UseIndependentChatPermissions: useIndependentChatPermissions,
	})
	return err
}

func (b *bot) ExportChatInviteLink(chatID Recipient) (*string, error) {
	params := exportChatInviteLinkRequest{
		ChatID: chatID,
	}

	data, err := b.sendMethodRequest(methodExportChatInviteLink, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Result string `json:"result"`
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result.Result, nil
}

func (b *bot) CreateChatInviteLink(chatID Recipient, name string, expireDate int64, memberLimit int, createsJoinRequest bool) (*ChatInviteLink, error) {
	params := createChatInviteLinkRequest{
		ChatID:             chatID,
		Name:               name,
		ExpireDate:         expireDate,
		MemberLimit:        memberLimit,
		CreatesJoinRequest: createsJoinRequest,
	}
	var result struct {
		Result ChatInviteLink `json:"result"`
	}
	data, err := b.sendMethodRequest(methodCreateChatInviteLink, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Result, err
}

func (b *bot) EditChatInviteLink(chatID Recipient, inviteLink, name string, expireDate int64, memberLimit int, createsJoinRequest bool) (*ChatInviteLink, error) {
	params := editChatInviteLinkRequest{
		ChatID:             chatID,
		InviteLink:         inviteLink,
		Name:               name,
		ExpireDate:         expireDate,
		MemberLimit:        memberLimit,
		CreatesJoinRequest: createsJoinRequest,
	}
	var result struct {
		Result ChatInviteLink `json:"result"`
	}
	data, err := b.sendMethodRequest(methodEditChatInviteLink, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Result, err
}

func (b *bot) RevokeChatInviteLink(chatID Recipient, inviteLink string) (*ChatInviteLink, error) {
	params := revokeChatInviteLinkRequest{
		ChatID:     chatID,
		InviteLink: inviteLink,
	}
	var result struct {
		Result ChatInviteLink `json:"result"`
	}
	data, err := b.sendMethodRequest(methodRevokeChatInviteLink, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result.Result, err
}

func (b *bot) ApproveChatJoinRequest(chatID Recipient, userID int64) error {
	params := approveChatJoinRequestParams{
		ChatID: chatID,
		UserID: userID,
	}
	_, err := b.sendMethodRequest(methodApproveChatJoinRequest, params)
	return err
}

func (b *bot) DeclineChatJoinRequest(chatID Recipient, userID int64) error {
	params := declineChatJoinRequestParams{
		ChatID: chatID,
		UserID: userID,
	}
	_, err := b.sendMethodRequest(methodDeclineChatJoinRequest, params)
	return err
}

func (b *bot) SetChatPhoto(chatID Recipient, photo InputFile) error {
	params := setChatPhotoParams{
		ChatID: chatID,
		Photo:  photo,
	}
	_, err := b.sendMethodRequest(methodSetChatPhoto, params)
	return err
}

func (b *bot) DeleteChatPhoto(chatID Recipient) error {
	params := deleteChatPhotoParams{
		ChatID: chatID,
	}
	_, err := b.sendMethodRequest(methodDeleteChatPhoto, params)
	return err
}

func (b *bot) SetChatTitle(chatID Recipient, title string) error {
	params := setChatTitleParams{
		ChatID: chatID,
		Title:  title,
	}
	_, err := b.sendMethodRequest(methodSetChatTitle, params)
	return err
}

func (b *bot) SetChatDescription(chatID Recipient, description string) error {
	params := setChatDescriptionParams{
		ChatID:      chatID,
		Description: description,
	}
	_, err := b.sendMethodRequest(methodSetChatDescription, params)
	return err
}

func (b *bot) PinChatMessage(chatID Recipient, messageID int, disableNotification bool) error {
	params := pinChatMessageRequest{
		ChatID:              chatID,
		MessageID:           messageID,
		DisableNotification: disableNotification,
	}
	_, err := b.sendMethodRequest(methodPinChatMessage, params)
	return err
}

func (b *bot) UnpinChatMessage(chatID Recipient, messageID int) error {
	params := unpinChatMessageRequest{
		ChatID:    chatID,
		MessageID: messageID,
	}
	_, err := b.sendMethodRequest(methodUnpinChatMessage, params)
	return err
}

func (b *bot) UnpinAllChatMessages(chatID Recipient) error {
	params := unpinAllChatMessagesRequest{
		ChatID: chatID,
	}
	_, err := b.sendMethodRequest(methodUnpinAllChatMessages, params)
	return err
}

func (b *bot) LeaveChat(chatID Recipient) error {
	params := leaveChatRequest{
		ChatID: chatID,
	}
	_, err := b.sendMethodRequest(methodLeaveChat, params)
	return err
}

func (b *bot) GetChat(chatID Recipient) (*Chat, error) {
	params := getChatRequest{
		ChatID: chatID,
	}
	var result struct {
		Result Chat `json:"result"`
	}
	data, err := b.sendMethodRequest(methodGetChat, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return &result.Result, err
}

func (b *bot) GetChatAdministrators(chatID Recipient) ([]ChatMember, error) {
	params := getChatAdministratorsRequest{
		ChatID: chatID,
	}
	var result struct {
		Result []ChatMember `json:"result"`
	}
	data, err := b.sendMethodRequest(methodGetChatAdministrators, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return result.Result, err
}

func (b *bot) GetChatMemberCount(chatID Recipient) (*int, error) {
	params := getChatMemberCountRequest{
		ChatID: chatID,
	}
	var result struct {
		Result int `json:"result"`
	}
	data, err := b.sendMethodRequest(methodGetChatMemberCount, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return &result.Result, err
}

func (b *bot) GetChatMember(chatID Recipient, userID int64) (*ChatMember, error) {
	params := getChatMemberRequest{
		ChatID: chatID,
		UserID: userID,
	}
	var result struct {
		Result ChatMember `json:"result"`
	}
	data, err := b.sendMethodRequest(methodGetChatMember, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return &result.Result, err
}
