package management

import (
	"fmt"
	"time"
    "net/url"
    "net/http"
    "bytes"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

type conf struct {
	Typ string
}

type authorization struct {
	Token string
	Authorized bool
	Expires time.Time
}

type mgmt struct {
	UserId string
	InstanceId string
	ClientIp string
	CreationTimestamp string
	ServiceUrl string
	ApiVersion string
	Authorization authorization
	Config conf
}

var m *mgmt

func (m *mgmt) IsAuthorized() bool {
	if(time.Since(m.Authorization.Expires).Seconds() < 0 && m.Authorization.Authorized) {
    	return true
    } else {
    	return false
    }
}

func createAuthorization(UserId string, InstanceId string, ServiceUrl string, ApiVersion string) authorization {
	apiUrl := ServiceUrl
	h256 := sha256.New()
	var StringToEncrypt bytes.Buffer
	StringToEncrypt.WriteString(InstanceId)
	StringToEncrypt.WriteString("-")
	StringToEncrypt.WriteString(GetLocalIp())
	StringToEncrypt.WriteString("-")
	StringToEncrypt.WriteString(UserId)
	h256.Write([]byte(StringToEncrypt.String()))
	
    resource := "/api/v"+ApiVersion+"/token/create"

    form := url.Values{}
    form.Add("InstanceId", InstanceId)
    form.Add("UserId", UserId)

    urlStr := apiUrl + resource

    client := &http.Client{}
    r, _ := http.NewRequest("POST", urlStr, strings.NewReader(form.Encode()))
    r.Header.Add("Authorization", base64.StdEncoding.EncodeToString(h256.Sum(nil)))
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    resp, _ := client.Do(r)
    fmt.Println(resp.Status)


	// Hier mit ManagementBackend verbinden und ein Token einfordern
	returner := authorization{
		Authorized: true,
		Token: "Jdlsdfj9458lsdfj",
		Expires: time.Now().AddDate(0,0,2),
	}
	return returner
}

func ManagementRequest() {
	fmt.Println("request")
}
func InitManagement(UserId string, InstanceId string, ServiceUrl string, ApiVersion string) bool {
	m := mgmt{
		UserId: UserId,
		InstanceId: InstanceId,
		ClientIp: GetLocalIp(),
		ServiceUrl: ServiceUrl,
		ApiVersion: ApiVersion,
		Authorization: createAuthorization(UserId,InstanceId,ServiceUrl,ApiVersion),
		Config: conf{
			Typ: "json",
		},
	}
	return m.IsAuthorized()
}
