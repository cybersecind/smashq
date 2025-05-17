package mqtt

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "sync"

    MQTT "github.com/eclipse/paho.mqtt.golang"
)

var Client MQTT.Client
var ResponseStore sync.Map // map[agentID]string

func Init() {
    caCert, _ := ioutil.ReadFile("certs/ca.crt")
    caPool := x509.NewCertPool()
    caPool.AppendCertsFromPEM(caCert)

    clientCert, err := tls.LoadX509KeyPair("../broker/certs/client.crt", "../broker/certs/client.key")
    if err != nil {
        panic(err)
    }

    tlsConfig := &tls.Config{
        RootCAs:      caPool,
        ServerName:   "localhost",
        Certificates: []tls.Certificate{clientCert},
        MinVersion:   tls.VersionTLS12,
    }

    opts := MQTT.NewClientOptions().
        AddBroker("tls://localhost:8883").
        SetTLSConfig(tlsConfig).
        SetClientID("c2-server")

    Client = MQTT.NewClient(opts)
    if token := Client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
}
