package quic

import (
	"crypto/tls"
	"net"
	"net/url"

	quic "github.com/lucas-clemente/quic-go"
	"github.com/wilson818/gsnova/common/channel"
	"github.com/wilson818/gsnova/common/dns"
	"github.com/wilson818/gsnova/common/logger"
	"github.com/wilson818/gsnova/common/mux"
	"github.com/wilson818/gsnova/common/netx"
)

type QUICProxy struct {
	//proxy.BaseProxy
}

func (p *QUICProxy) Features() channel.FeatureSet {
	return channel.FeatureSet{
		AutoExpire: true,
		Pingable:   false,
	}
}

func (tc *QUICProxy) CreateMuxSession(server string, conf *channel.ProxyChannelConfig) (mux.MuxSession, error) {
	rurl, err := url.Parse(server)
	if nil != err {
		return nil, err
	}
	hostport := rurl.Host
	tcpHost, tcpPort, _ := net.SplitHostPort(hostport)
	if net.ParseIP(tcpHost) == nil {
		iphost, err := dns.DnsGetDoaminIP(tcpHost)
		if nil != err {
			return nil, err
		}
		hostport = net.JoinHostPort(iphost, tcpPort)
	}
	var quicSession quic.Connection

	udpAddr, err := net.ResolveUDPAddr("udp", hostport)
	if err != nil {
		return nil, err
	}
	udpConn, err := netx.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 0})
	if err != nil {
		return nil, err
	}
	quicConfig := &quic.Config{
		KeepAlivePeriod: 300000,
	}
	quicSession, err = quic.Dial(udpConn, udpAddr, hostport, &tls.Config{InsecureSkipVerify: true}, quicConfig)

	if err != nil {
		return nil, err
	}
	logger.Debug("Connect %s success.", server)
	return &mux.QUICMuxSession{Connection: quicSession}, nil
}

func init() {
	channel.RegisterLocalChannelType("quic", &QUICProxy{})
}
