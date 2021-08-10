package wshandlers

import (
	"gopkg.in/olahol/melody.v1"
)

var WsService = melody.New()

func InitWsServise(m *melody.Melody) {
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		println(string(msg))
		m.Broadcast(msg)
	})
}
