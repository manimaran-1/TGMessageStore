// (c) Jisin0
//
// config/buttons.go contains basic commands buttons.

package config

import "github.com/PaulSonOfLars/gotgbot/v2"

var Buttons map[string][][]gotgbot.InlineKeyboardButton = map[string][][]gotgbot.InlineKeyboardButton

// Single buttons used to build composite markups.
var (
	aboutButton = gotgbot.InlineKeyboardButton{Text: "About ℹ️", CallbackData: "cmd_ABOUT"}
	helpButton  = gotgbot.InlineKeyboardButton{Text: "Help ❓", CallbackData: "cmd_HELP"}
	homeButton  = gotgbot.InlineKeyboardButton{Text: "Home 🏠", CallbackData: "cmd_START"}
)
