package quic

import (
	"context"
	"crypto/tls"

	quic "github.com/lucas-clemente/quic-go"
	"github.com/wilson818/gsnova/common/channel"
	"github.com/wilson818/gsnova/common/logger"
	"github.com/wilson818/gsnova/common/mux"
)

func servQUIC(lp quic.Listener) {
	for {
		sess, err := lp.Accept(context.Background())
		if nil != err {
			continue
		}
		muxSession := &mux.QUICMuxSession{Connection: sess}
		go channel.ServProxyMuxSession(muxSession, nil, nil)
	}
	//ws.WriteMessage(websocket.CloseMessage, []byte{})
}

func StartQuicProxyServer(addr string, config *tls.Config) error {
	lp, err := quic.ListenAddr(addr, config, nil)
	if nil != err {
		logger.Error("[ERROR]Failed to listen QUIC address:%s with reason:%v", addr, err)
		return err
	}
	logger.Info("Listen on QUIC address:%s", addr)
	servQUIC(lp)
	return nil
}
