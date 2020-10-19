package main

//Update struct
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

//Message struct
type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

//Chat struct
type Chat struct {
	ChatID int `json:"id"`
}

//RestResponce struct
type RestResponce struct {
	Result []Update `json:"result"`
}

//BotMessage struct
type BotMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
