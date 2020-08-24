package device

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-benchmark/config"
	"github.com/go-benchmark/service"
	"github.com/gobench-io/gobench/clients/http"
	"github.com/gobench-io/gobench/clients/mqtt"
	"github.com/gobench-io/gobench/logger"
)

var (
	//ErrServiceNotFound Error service not found
	ErrServiceNotFound = errors.New("service not found")
)

type mqtter interface {
	Connect(context.Context) error
	Publish(context.Context, string, byte, []byte) error
	Subscribe(context.Context, string, byte, paho.MessageHandler) error
}

// Device represent to device
type Device struct {
	ID         string     `json:"id"`
	Credential credential `json:"mqtt"`

	services map[string]*service.Service

	vu              int // vu index
	mc              mqtt.MqttClient
	host            string
	mqttPrefix      string
	runServiceDelay int
	historyInterval float64
	dc              defaultConfig
	log             *logger.Log
}
type defaultConfig struct {
	realtimeInterval float64
	realtimeLength   int
}

type credential struct {
	Endpoint          string `json:"endpoint"`
	Port              int    `json:"port"`
	Protocol          string `json:"protocol"`
	PrivateKey        string `json:"privateKey"`
	Certificate       string `json:"certificate"`
	RootCaCertificate string `json:"rootCaCertificate"`
	ClientID          string `json:"clientId"`
}

// DeviceReq represent to device
type DeviceReq struct {
	ID           string `json:"id" faker:"uuid_digit"`
	Type         string `json:"type"`
	Name         string `json:"name" faker:"string"`
	MacAddress   string `json:"macAddress" faker:"mac_address"`
	WifiSsid     string `json:"wifiSsid" faker:"uuid_hyphenated"`
	WifiPassword string `json:"wifiPassword"  faker:"password"`
}

// NewDevice create a new device
func NewDevice(opts *config.Options, vu int,log *logger.Log) (*Device, error) {
	d := &Device{
		ID:       faker.UUIDDigit(),
		services: make(map[string]*service.Service),
		log :     log,
	}

	d.setOpts(opts, vu)

	return d, nil
}

func (d *Device) setOpts(opts *config.Options, vu int) *Device {
	d.vu = vu
	d.host = opts.Host
	d.mqttPrefix = opts.MQTTPrefix
	d.historyInterval = opts.DC.HistoryInterval
	d.dc = defaultConfig{
		realtimeInterval: opts.DC.RealtimeInterval,
		realtimeLength:   opts.SC.RealtimeLength,
	}

	return d
}

// GetMqttAuth get mqtt auth
func (d *Device) GetMqttAuth(ctx context.Context) (err error) {
	path := "/api/fws"
	url := "https://" + d.host + path

	req := map[string]string{
		"deviceId": d.ID,
	}
	reqB, _ := json.Marshal(req)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client, err := http.NewHttpClient(ctx, path)
	if err != nil {
		return
	}

	buf, err := client.Post(ctx, url, reqB, headers)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &d)

	return err
}

// ConnectMqtt connects the device with credential to broker.
// The real device has last will setup; however on this simulation we ignore
func (d *Device) ConnectMqtt(ctx context.Context) error {
	tlsConfig := new(tls.Config)
	tlsConfig.ClientCAs = x509.NewCertPool()
	tlsConfig.RootCAs = x509.NewCertPool()

	certificate := []byte(d.Credential.Certificate)
	key := []byte(d.Credential.PrivateKey)
	rootCA := []byte(d.Credential.RootCaCertificate)
	cert, err := tls.X509KeyPair(certificate, key)
	if err != nil {
		return fmt.Errorf("X509KeyPair: %v", err)
	}
	tlsConfig.Certificates = []tls.Certificate{cert}
	if !tlsConfig.ClientCAs.AppendCertsFromPEM(certificate) {
		return fmt.Errorf("append certs from pem: %v", err)
	}
	if !tlsConfig.RootCAs.AppendCertsFromPEM(rootCA) {
		return fmt.Errorf("append certs from pem: %v", err)
	}

	host := fmt.Sprintf("tls://%s:8883", d.Credential.Endpoint)

	opts := mqtt.NewClientOptions()
	opts.
		SetMaxReconnectInterval(time.Second * 55).
		SetAutoReconnect(true).
		SetClientID(d.ID).
		SetTLSConfig(tlsConfig).
		AddBroker(host)
		// SetDefaultPublishHandler(handlers.HandleMqttMsg(ctx))

	if d.mc, err = mqtt.NewMqttClient(ctx, opts); err != nil {
		return err
	}

	err = d.mc.Connect(ctx)

	return err
}

// Run setups device properties:
// 1. subscribe to $prefix/:deviceId/config
// 2. publish network config and
// 3. publish active bots
func (d *Device) Run(ctx context.Context) (err error) {
	// subscribe
	if err = d.sub(ctx); err != nil {
		return
	}

	// send networkconfig
	nc, err := d.networkconfig(ctx)
	if err != nil {
		return
	}

	// send active bot
	if _, err = d.activeBots(ctx, nc.Network.Bots); err != nil {
		return
	}

	return
}

// Finish disconnect mqtt broker
func (d *Device) Finish(ctx context.Context) error {
	return d.mc.Disconnect(ctx)
}

func (d *Device) opStateTopic() string {
	return fmt.Sprintf("%s/%s/opState", d.mqttPrefix, d.ID)
}
