package main

import (
	"github.com/sirupsen/logrus"
	"net"
)

var portlist = map[string][]string{
	"53": {"udp", "tcp"},
	"67": {"udp"},
	"69": {"udp"},
	"80": {"tcp"},
	"443": {"tcp"},
	"546": {"udp"},
	"6443": {"tcp"},
	"8080": {"tcp"},
	"9000": {"tcp"},
	"9090": {"tcp"},
	"22623": {"tcp"},
}

// preflight error counter
var preflightErrorCount int = 0

func portCheck() int {
	logrus.Info("Starting Port Checks")
	// set the error count to 0
	porterrorcount := 0

	for port, protocolArray := range portlist {
		for _, protocol := range protocolArray {
			//check if you can listen on this port on TCP
			if protocol == "tcp" {
			    logrus.Infof("Testing port %s on protocol %s", port, protocol)
	            if t, err := net.Listen(protocol, ":" + port); err == nil {
                    logrus.Infof("Error TCP %s", err)
                } else if err != nil {
                    logrus.Infof("Error TCP %s", err)
					logrus.Warnf("Port check  %s/%s is in use", port, protocol)
					porterrorcount += 1
	                defer t.Close()
                }

			} else if protocol == "udp" {
			    logrus.Infof("Testing port %s on protocol %s", port, protocol)
				if u, err := net.ListenPacket(protocol, ":"+port); err == nil {
                    logrus.Infof("Error UDP %s", err)
					// If this returns an error, then something else is listening on this port
				} else if err != nil {
                        logrus.Infof("Error UDP %s", err)
					    logrus.Warnf("Port check  %s/%s is in use", port, protocol)
						porterrorcount += 1
				        defer u.Close()
                }
			}
		}
	}

	// Display that no errors were found
	if porterrorcount > 0 {
		preflightErrorCount += 1
	}
	logrus.WithFields(logrus.Fields{"Port Issues": porterrorcount}).Info("Preflight checks for Ports")
	return porterrorcount
}

func main () {
    portCheck()
}
