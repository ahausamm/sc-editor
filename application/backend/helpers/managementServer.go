package management

import (
	"fmt"
)
type conf struct {
	typ string
}
var m *mgmt

type mgmt struct {
	userId string
	clientIp string
	creationTimestamp string
	serviceUrl string
	authorized bool
	config conf
}

func (m *mgmt) IsAuthorized() bool {
    return m.authorized
}
func (m *mgmt) authorize() bool {
    //Request machen
    m.authorized = true;
    return m.IsAuthorized()
}

func createAuthorizationKey() string {

	return "authorizationString"
}

func InitManagement(userId string) bool {
	m := mgmt{
		userId: userId,
		clientIp: GetIp(),
		creationTimestamp: "dd",
		serviceUrl: "http://hand.sg.werk.ch",
		authorized: false,
		config: conf{
			typ: "json",
		},
	}
	fmt.Println(m.clientIp)
	m.authorize()
	return m.IsAuthorized()
}
