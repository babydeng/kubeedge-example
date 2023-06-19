package main

import (
	"errors"
	"flag"
	"os"
	"sync"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/glog"
)

const (
	modelName                 = "BUZZER-MQTT"
	powerStatus               = "power-status"
	pinNumberConfig           = "TOPIC-MQTT"
	DeviceETPrefix            = "$hw/events/device/"
	DeviceETStateUpdateSuffix = "/state/update"
	TwinETUpdateSuffix        = "/twin/update"
	TwinETCloudSyncSuffix     = "/twin/cloud_updated"
	TwinETGetResultSuffix     = "/twin/get/result"
	TwinETGetSuffix           = "/twin/get"
)

var cli MQTT.Client
var Token_client Token
var wg sync.WaitGroup
var deviceTwinResult DeviceTwinUpdate
var deviceID string
var topic float64

// init for getting command line arguments for glog and initiating the MQTT connection
func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
	err := configFile.ReadFromConfigFile()
	if err != nil {
		glog.Error(errors.New("Error while reading from config file " + err.Error()))
		os.Exit(1)
	}
	glog.Info("MQTT URL: ", configFile.MQTTURL)
	// connect to mqtt broker
	connectToMqtt(configFile.MQTTURL)

	err := LoadConfigMap()
	if err != nil {
		glog.Error(errors.New("Error while reading from config map " + err.Error()))
		os.Exit(1)
	}
}
func main() {

}
