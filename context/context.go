package context

import (
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"

	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/baetyl/baetyl-go/v2/log"
	"github.com/baetyl/baetyl-go/v2/pki"
	"github.com/baetyl/baetyl-go/v2/utils"
)

// Env keys
const (
	EnvKeyConfFile    = "BAETYL_CONF_FILE"
	EnvKeyNodeName    = "BAETYL_NODE_NAME"
	EnvKeyAppName     = "BAETYL_APP_NAME"
	EnvKeyServiceName = "BAETYL_SERVICE_NAME"
	EnvKeyCodePath    = "BAETYL_CODE_PATH"
	EnvKeyCertPath    = "BAETYL_CERT_PATH"

	SystemCertCA  = "ca.pem"
	SystemCertCrt = "crt.pem"
	SystemCertKey = "key.pem"
	SystemCertOU  = "BAETYL"

	SystemAppInit = "baetyl-init"
)

var (
	ErrSystemCertInvalid = errors.New("failed to verify system certificate")
)

// Context of service
type Context interface {
	// NodeName returns node name.
	NodeName() string
	// AppName returns app name.
	AppName() string
	// ServiceName returns service name.
	ServiceName() string
	// ConfFile returns config file.
	ConfFile() string
	// ServiceConfig returns service config.
	ServiceConfig() ServiceConfig

	// Load returns the value stored in the map for a key, or nil if no value is present.
	// The ok result indicates whether value was found in the map.
	Load(key interface{}) (value interface{}, ok bool)
	// Store sets the value for a key.
	Store(key, value interface{})
	// LoadOrStore returns the existing value for the key if present.
	// Otherwise, it stores and returns the given value.
	// The loaded result is true if the value was loaded, false if stored.
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	// Delete deletes the value for a key.
	Delete(key interface{})

	// Get get system resource object
	GetSystemResource() *SystemResource

	// LoadCustomConfig loads custom config, if path is empty, will load config from default path.
	LoadCustomConfig(cfg interface{}, files ...string) error
	// Log returns logger interface.
	Log() *log.Logger

	// Wait waits until exit, receiving SIGTERM and SIGINT signals.
	Wait()
	// WaitChan returns wait channel.
	WaitChan() <-chan os.Signal
}

type ctx struct {
	sync.Map
	cfg ServiceConfig
	log *log.Logger

	nodeName    string
	appName     string
	serviceName string
	confFile    string
	httpAddress string
	mqttAddress string
	certPath    string
	res         *SystemResource
}

// NewContext creates a new context
func NewContext(confFile string) (Context, error) {
	if confFile == "" {
		confFile = os.Getenv(EnvKeyConfFile)
	}
	c := &ctx{
		confFile:    confFile,
		nodeName:    os.Getenv(EnvKeyNodeName),
		appName:     os.Getenv(EnvKeyAppName),
		serviceName: os.Getenv(EnvKeyServiceName),
		certPath:    os.Getenv(EnvKeyCertPath),
		res:         &SystemResource{},
	}

	var fs []log.Field
	if c.nodeName != "" {
		fs = append(fs, log.Any("node", c.nodeName))
	}
	if c.appName != "" {
		fs = append(fs, log.Any("app", c.appName))
	}
	if c.serviceName != "" {
		fs = append(fs, log.Any("service", c.serviceName))
	}
	c.log = log.With(fs...)

	err := c.LoadCustomConfig(&c.cfg)
	if err != nil {
		c.log.Error("failed to load service config, to use default config", log.Error(err))
		utils.UnmarshalYAML(nil, &c.cfg)
	}

	_log, err := log.Init(c.cfg.Logger, fs...)
	if err != nil {
		c.log.Error("failed to init logger", log.Error(err))
	}
	c.log = _log

	if c.serviceName != SystemAppInit {
		err = c.checkAndSetCert()
		if err != nil {
			c.Log().Error("service has stopped with error", log.Error(err))
			return nil, errors.Trace(err)
		}
	}

	if c.cfg.HTTP.Address == "" {
		if c.cfg.HTTP.Key == "" {
			c.cfg.HTTP.Address = "http://baetyl-function:80"
		} else {
			c.cfg.HTTP.Address = "https://baetyl-function:443"
		}
	}

	if c.cfg.MQTT.Address == "" {
		if c.cfg.MQTT.Key == "" {
			c.cfg.MQTT.Address = "tcp://baetyl-broker:1883"
		} else {
			c.cfg.MQTT.Address = "ssl://baetyl-broker:8883"
		}
	}
	c.log.Debug("context is created", log.Any("file", confFile), log.Any("conf", c.cfg))
	return c, nil
}

func (c *ctx) GetSystemResource() *SystemResource {
	return c.res
}

func (c *ctx) NodeName() string {
	return c.nodeName
}

func (c *ctx) AppName() string {
	return c.appName
}

func (c *ctx) ServiceName() string {
	return c.serviceName
}

func (c *ctx) ConfFile() string {
	return c.confFile
}

func (c *ctx) ServiceConfig() ServiceConfig {
	return c.cfg
}

func (c *ctx) LoadCustomConfig(cfg interface{}, files ...string) error {
	f := c.confFile
	if len(files) > 0 && len(files[0]) > 0 {
		f = files[0]
	}
	if utils.FileExists(f) {
		return errors.Trace(utils.LoadYAML(f, cfg))
	}
	return errors.Trace(utils.UnmarshalYAML(nil, cfg))
}

func (c *ctx) Log() *log.Logger {
	return c.log
}

func (c *ctx) Wait() {
	<-c.WaitChan()
}

func (c *ctx) WaitChan() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	signal.Ignore(syscall.SIGPIPE)
	return sig
}

func (c *ctx) checkAndSetCert() error {
	// get and check ca
	ca, err := ioutil.ReadFile(path.Join(c.certPath, SystemCertCA))
	if err != nil {
		return errors.Trace(err)
	}
	// get crt and key
	crt, err := ioutil.ReadFile(path.Join(c.certPath, SystemCertCrt))
	if err != nil {
		return errors.Trace(err)
	}
	info, err := pki.ParseCertificates(crt)
	if err != nil {
		return errors.Trace(err)
	}
	if len(info) != 1 || len(info[0].Subject.OrganizationalUnit) != 1 ||
		info[0].Subject.OrganizationalUnit[0] != SystemCertOU {
		return errors.Trace(ErrSystemCertInvalid)
	}
	key, err := ioutil.ReadFile(path.Join(c.certPath, SystemCertKey))
	if err != nil {
		return errors.Trace(err)
	}
	// set
	c.res.ca = ca
	c.res.crt = crt
	c.res.key = key
	return nil
}
