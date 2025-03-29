package handlers

import (
	"TicTacToe/assert"
	"TicTacToe/gameServer/internal/event"
	"TicTacToe/gameServer/message"
	"errors"

	"github.com/google/uuid"
)

type EventExit struct {
	RoomUUID uuid.UUID
	Player *Player
	ConnectionId uuid.UUID
	OpponentConnId uuid.UUID
	RoomExist bool
}

type EventMove struct {
	X int
	Y int
	Player *Player
}

type EventSendMessage struct {
	ConnectionId uuid.UUID
	Msg message.Message
}

func (eType EventExit) GetType() event.EventType {
	return event.EventTypeExit;
}
func (eType EventMove) GetType() event.EventType {
	return event.EventTypeMove;
}
func (eType EventSendMessage) GetType() event.EventType {
	return event.EventTypeSendMessage;
}

func EventFromMessage(msg message.Message) (event.Event, error) {
	assert.NotNil(msg, "message was nil")

	switch message.ClientMsg(msg.Type) {
	case message.Move:
		data, err := message.ParseMessage[message.MoveMessage](msg)

		if err != nil {
			return nil, err
		}

		return EventMove{
			X: data.X,
			Y: data.Y,
		}, nil

	default:
		return nil, errors.New("this message has no corresponding event")
	}
}