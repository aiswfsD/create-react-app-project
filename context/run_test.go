package context

import (
	"github.com/baetyl/baetyl-go/v2/http"
	"github.com/baetyl/baetyl-go/v2/log"
	"github.com/baetyl/baetyl-go/v2/mqtt"
	"github.com/baetyl/baetyl-go/v2/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	Run(func(ctx Context) error {
		assert.Equal(t, "etc/baetyl/service.yml", ctx.ConfFile())
		assert.Equal(t, &SystemConfig{
			Certificate: utils.Certificate{CA: "var/lib/baetyl/system/certs/ca.pem", Key: "var/lib/baetyl/system/certs/key.pem", Cert: "var/lib/baetyl/system/certs/crt.pem", Name: "", InsecureSkipVerify: false, ClientAuthType: 0},
			Function:    http.ClientConfig{Address: "https://baetyl-function.baetyl-edge-system", Timeout: 30000000000, KeepAlive: 30000000000, MaxIdleConns: 100, IdleConnTimeout: 90000000000, TLSHandshakeTimeout: 10000000000, ExpectContinueTimeout: 1000000000, Certificate: utils.Certificate{CA: "var/lib/baetyl/system/certs/ca.pem", Key: "var/lib/baetyl/system/certs/key.pem", Cert: "var/lib/baetyl/system/certs/crt.pem", Name: "", InsecureSkipVerify: false, ClientAuthType: 0}},
			Broker:      mqtt.ClientConfig{Address: "ssl://baetyl-broker.baetyl-edge-system:8883", Username: "", Password: "", ClientID: "baetyl-svc-service", CleanSession: false, Timeout: 30000000000, KeepAlive: 30000000000, MaxReconnectInterval: 180000000000, MaxCacheMessages: 10, DisableAutoAck: false, Subscriptions: []mqtt.QOSTopic{}, Certificate: utils.Certificate{CA: "var/lib/baetyl/system/certs/ca.pem", Key: "var/lib/baetyl/system/certs/key.pem", Cert: "var/lib/baetyl/system/certs/crt.pem", Name: "", InsecureSkipVerify: false, ClientAuthType: 0}},
			Logger:      log.Config{Level: "info", Encoding: "json", Filename: "", Compress: false, MaxAge: 15, MaxSize: 50, MaxBackups: 15, EncodeTime: "", EncodeLevel: ""},
		}, ctx.SystemConfig())
		return nil
	})
}
