package telebot

type method int

const (
	methodGetUpdates method = iota
	methodSetWebhook
	methodDeleteWebhook
	methodGetWebhookInfo
	methodGetMe
	methodLogOut
	methodClose
	methodSendMessage
	methodEditMessageText
	methodEditMessageCaption
	methodEditMessageMedia
	methodEditMessageLiveLocation
	methodStopMessageLiveLocation
	methodEditMessageReplyMarkup
	methodStopPoll
	methodDeleteMessage
	methodDeleteMessages
	methodForwardMessage
	methodForwardMessages
	methodCopyMessage
	methodCopyMessages
	methodSendPhoto
	methodSendAudio
	methodSendDocument
	methodSendVideo
	methodSendAnimation
	methodSendVoice
	methodSendVideoNote
	methodSendMediaGroup
	methodSendLocation
	methodSendVenue
	methodSendContact
	methodSendPoll
	methodSendDice
	methodSendChatAction
	methodSetMessageReaction
	methodGetUserProfilePhotos
	methodGetFile
	methodBanChatMember
	methodUnbanChatMember
	methodRestrictChatMember
	methodPromoteChatMember
	methodSetChatAdministratorCustomTitle
	methodBanChatSenderChat
	methodUnbanChatSenderChat
	methodSetChatPermissions
	methodExportChatInviteLink
	methodCreateChatInviteLink
	methodEditChatInviteLink
	methodRevokeChatInviteLink
	methodApproveChatJoinRequest
	methodDeclineChatJoinRequest
	methodSetChatPhoto
	methodDeleteChatPhoto
	methodSetChatTitle
	methodSetChatDescription
	methodPinChatMessage
	methodUnpinChatMessage
	methodUnpinAllChatMessages
	methodLeaveChat
	methodGetChat
	methodGetChatAdministrators
	methodGetChatMemberCount
	methodGetChatMember
	methodSetChatStickerSet
	methodDeleteChatStickerSet
	methodGetForumTopicIconStickers
	methodCreateForumTopic
	methodEditForumTopic
	methodCloseForumTopic
	methodReopenForumTopic
	methodDeleteForumTopic
	methodUnpinAllForumTopicMessages
	methodEditGeneralForumTopic
	methodCloseGeneralForumTopic
	methodReopenGeneralForumTopic
	methodHideGeneralForumTopic
	methodUnhideGeneralForumTopic
	methodUnpinAllGeneralForumTopicMessages
	methodAnswerCallbackQuery
	methodGetUserChatBoosts
	methodSetMyCommands
	methodDeleteMyCommands
	methodGetMyCommands
	methodSetMyName
	methodGetMyName
	methodSetMyDescription
	methodGetMyDescription
	methodSetMyShortDescription
	methodGetMyShortDescription
	methodSetChatMenuButton
	methodGetChatMenuButton
	methodSetMyDefaultAdministratorRights
	methodGetMyDefaultAdministratorRights
	methodSendSticker
	methodGetStickerSet
	methodGetCustomEmojiStickers
	methodUploadStickerFile
	methodCreateNewStickerSet
	methodAddStickerToSet
	methodSetStickerPositionInSet
	methodDeleteStickerFromSet
	methodSetStickerEmojiList
	methodSetStickerKeywords
	methodSetStickerMaskPosition
	methodSetStickerSetTitle
	methodSetStickerSetThumbnail
	methodSetCustomEmojiStickerSetThumbnail
	methodDeleteStickerSet
	methodAnswerInlineQuery
	methodAnswerWebAppQuery
	methodSentWebAppMessage
	methodSendInvoice
	methodCreateInvoiceLink
	methodAnswerShippingQuery
	methodAnswerPreCheckoutQuery
	methodSetPassportDataErrors
	methodSendGame
	methodSetGameScore
	methodGetGameHighScores
)

func (m method) String() string {
	switch m {
	case methodGetUpdates:
		return "getUpdates"
	case methodSetWebhook:
		return "setWebhook"
	case methodDeleteWebhook:
		return "deleteWebhook"
	case methodGetWebhookInfo:
		return "getWebhookInfo"
	case methodGetMe:
		return "getMe"
	case methodLogOut:
		return "logOut"
	case methodClose:
		return "close"
	case methodSendMessage:
		return "sendMessage"
	case methodEditMessageText:
		return "editMessageText"
	case methodEditMessageCaption:
		return "editMessageCaption"
	case methodEditMessageMedia:
		return "editMessageMedia"
	case methodEditMessageLiveLocation:
		return "editMessageLiveLocation"
	case methodStopMessageLiveLocation:
		return "stopMessageLiveLocation"
	case methodEditMessageReplyMarkup:
		return "editMessageReplyMarkup"
	case methodStopPoll:
		return "stopPoll"
	case methodDeleteMessages:
		return "deleteMessages"
	case methodForwardMessages:
		return "forwardMessages"
	case methodCopyMessages:
		return "copyMessages"
	case methodSendPhoto:
		return "sendPhoto"
	case methodSendAudio:
		return "sendAudio"
	case methodSendDocument:
		return "sendDocument"
	case methodSendVideo:
		return "sendVideo"
	case methodSendAnimation:
		return "sendAnimation"
	case methodSendVoice:
		return "sendVoice"
	case methodSendVideoNote:
		return "sendVideoNote"
	case methodSendMediaGroup:
		return "sendMediaGroup"
	case methodSendLocation:
		return "sendLocation"
	case methodSendVenue:
		return "sendVenue"
	case methodSendContact:
		return "sendContact"
	case methodSendPoll:
		return "sendPoll"
	case methodSendDice:
		return "sendDice"
	case methodSendChatAction:
		return "sendChatAction"
	case methodSetMessageReaction:
		return "setMessageReaction"
	case methodGetUserProfilePhotos:
		return "getUserProfilePhotos"
	case methodGetFile:
		return "getFile"
	case methodBanChatMember:
		return "banChatMember"
	case methodUnbanChatMember:
		return "unbanChatMember"
	case methodRestrictChatMember:
		return "restrictChatMember"
	case methodPromoteChatMember:
		return "promoteChatMember"
	case methodSetChatAdministratorCustomTitle:
		return "setChatAdministratorCustomTitle"
	case methodBanChatSenderChat:
		return "banChatSenderChat"
	case methodUnbanChatSenderChat:
		return "unbanChatSenderChat"
	case methodSetChatPermissions:
		return "setChatPermissions"
	case methodExportChatInviteLink:
		return "exportChatInviteLink"
	case methodCreateChatInviteLink:
		return "createChatInviteLink"
	case methodEditChatInviteLink:
		return "editChatInviteLink"
	case methodRevokeChatInviteLink:
		return "revokeChatInviteLink"
	case methodApproveChatJoinRequest:
		return "approveChatJoinRequest"
	case methodDeclineChatJoinRequest:
		return "declineChatJoinRequest"
	case methodSetChatPhoto:
		return "setChatPhoto"
	case methodDeleteChatPhoto:
		return "deleteChatPhoto"
	case methodSetChatTitle:
		return "setChatTitle"
	case methodSetChatDescription:
		return "setChatDescription"
	case methodPinChatMessage:
		return "pinChatMessage"
	case methodUnpinChatMessage:
		return "unpinChatMessage"
	case methodUnpinAllChatMessages:
		return "unpinAllChatMessages"
	case methodLeaveChat:
		return "leaveChat"
	case methodGetChat:
		return "getChat"
	case methodGetChatAdministrators:
		return "getChatAdministrators"
	case methodGetChatMemberCount:
		return "getChatMemberCount"
	case methodGetChatMember:
		return "getChatMember"
	case methodSetChatStickerSet:
		return "setChatStickerSet"
	case methodDeleteChatStickerSet:
		return "deleteChatStickerSet"
	case methodGetForumTopicIconStickers:
		return "getForumTopicIconStickers"
	case methodCreateForumTopic:
		return "createForumTopic"
	case methodEditForumTopic:
		return "editForumTopic"
	case methodCloseForumTopic:
		return "closeForumTopic"
	case methodReopenForumTopic:
		return "reopenForumTopic"
	case methodDeleteForumTopic:
		return "deleteForumTopic"
	case methodUnpinAllForumTopicMessages:
		return "unpinAllForumTopicMessages"
	case methodEditGeneralForumTopic:
		return "editGeneralForumTopic"
	case methodCloseGeneralForumTopic:
		return "closeGeneralForumTopic"
	case methodReopenGeneralForumTopic:
		return "reopenGeneralForumTopic"
	case methodHideGeneralForumTopic:
		return "hideGeneralForumTopic"
	case methodUnhideGeneralForumTopic:
		return "unhideGeneralForumTopic"
	case methodUnpinAllGeneralForumTopicMessages:
		return "unpinAllGeneralForumTopicMessages"
	case methodAnswerCallbackQuery:
		return "answerCallbackQuery"
	case methodGetUserChatBoosts:
		return "getUserChatBoosts"
	case methodSetMyCommands:
		return "setMyCommands"
	case methodDeleteMyCommands:
		return "deleteMyCommands"
	case methodGetMyCommands:
		return "getMyCommands"
	case methodSetMyName:
		return "setMyName"
	case methodGetMyName:
		return "getMyName"
	case methodSetMyDescription:
		return "setMyDescription"
	case methodGetMyDescription:
		return "getMyDescription"
	case methodSetMyShortDescription:
		return "setMyShortDescription"
	case methodGetMyShortDescription:
		return "getMyShortDescription"
	case methodSetChatMenuButton:
		return "setChatMenuButton"
	case methodGetChatMenuButton:
		return "getChatMenuButton"
	case methodSetMyDefaultAdministratorRights:
		return "setMyDefaultAdministratorRights"
	case methodGetMyDefaultAdministratorRights:
		return "getMyDefaultAdministratorRights"
	case methodSendSticker:
		return "sendSticker"
	case methodGetStickerSet:
		return "getStickerSet"
	case methodGetCustomEmojiStickers:
		return "getCustomEmojiStickers"
	case methodUploadStickerFile:
		return "uploadStickerFile"
	case methodCreateNewStickerSet:
		return "createNewStickerSet"
	case methodAddStickerToSet:
		return "addStickerToSet"
	case methodSetStickerPositionInSet:
		return "setStickerPositionInSet"
	case methodDeleteStickerFromSet:
		return "deleteStickerFromSet"
	case methodSetStickerEmojiList:
		return "setStickerEmojiList"
	case methodSetStickerKeywords:
		return "setStickerKeywords"
	case methodSetStickerSetTitle:
		return "setStickerSetTitle"
	case methodSetStickerSetThumbnail:
		return "setStickerSetThumbnail"
	case methodDeleteStickerSet:
		return "deleteStickerSet"
	case methodAnswerInlineQuery:
		return "answerInlineQuery"
	case methodAnswerWebAppQuery:
		return "answerWebAppQuery"
	case methodSentWebAppMessage:
		return "sentWebAppMessage"
	case methodSendInvoice:
		return "sendInvoice"
	case methodCreateInvoiceLink:
		return "createInvoiceLink"
	case methodAnswerShippingQuery:
		return "answerShippingQuery"
	case methodAnswerPreCheckoutQuery:
		return "answerPreCheckoutQuery"
	case methodSetPassportDataErrors:
		return "setPassportDataErrors"
	case methodSendGame:
		return "sendGame"
	case methodSetGameScore:
		return "setGameScore"
	case methodGetGameHighScores:
		return "getGameHighScores"
	case methodSetStickerMaskPosition:
		return "setStickerMaskPosition"
	case methodSetCustomEmojiStickerSetThumbnail:
		return "setCustomEmojiStickerSetThumbnail"
	default:
		panic("telebot: unknown method")
	}
}
