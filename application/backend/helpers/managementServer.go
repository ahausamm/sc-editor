package management

import (
	"fmt"
	"time"
)

type conf struct {
	typ string
}

type authorization struct {
	token string
	authorized bool
	expires time.Time
}

type mgmt struct {
	userId string
	instanceId string
	clientIp string
	creationTimestamp string
	serviceUrl string
	authorization authorization
	config conf
}

var m *mgmt

func (m *mgmt) IsAuthorized() bool {
	if(time.Since(m.authorization.expires).Seconds() < 0 && m.authorization.authorized) {
    	return true
    } else {
    	return false
    }
}

func (m *mgmt) createAuthorization() authorization {
	// Hier mit ManagementBackend verbinden und ein Token einfordern
	returner := authorization{
		authorized: true,
		token: "Jdlsdfj9458lsdfj",
		expires: time.Now().AddDate(0,0,2),
	}
	return returner
}

func ManagementRequest() {
	fmt.Println("request")
}
func InitManagement(userId string, instanceId string) bool {
	m := mgmt{
		userId: userId,
		instanceId: instanceId,
		clientIp: GetLocalIp(),
		serviceUrl: "http://hand.sg.werk.ch",
		authorization: m.createAuthorization(),
		config: conf{
			typ: "json",
		},
	}
	return m.IsAuthorized()
}
