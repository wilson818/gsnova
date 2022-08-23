package common

import (
	_ "github.com/wilson818/gsnova/common/channel/direct"
	_ "github.com/wilson818/gsnova/common/channel/http"
	_ "github.com/wilson818/gsnova/common/channel/http2"
	_ "github.com/wilson818/gsnova/common/channel/kcp"
	_ "github.com/wilson818/gsnova/common/channel/quic"
	_ "github.com/wilson818/gsnova/common/channel/ssh"
	_ "github.com/wilson818/gsnova/common/channel/tcp"
	_ "github.com/wilson818/gsnova/common/channel/websocket"
)
