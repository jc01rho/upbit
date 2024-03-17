package upbit

import (
	"upbit/model"
	"upbit/model/upbitWebsocket"
)

func _createWSRequestBody(ticketName string, dataType upbitWebsocket.RequestType, markets []string) []interface{} {
	wsReqList := make([]interface{}, 0)
	wsReqList = append(wsReqList, &model.WebSocketTicket{

		Ticket: ticketName,
	})

	wsReqList = append(wsReqList, &model.WebSocketType{

		Type:           string(dataType),
		Codes:          markets,
		IsOnlySnapshot: nil,
		IsOnlyRealtime: nil,
	})

	return wsReqList
}
