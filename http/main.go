package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type OSResponse struct {
	/*Shards struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`*/
	Hits struct {
		Hits []struct {
			//Index  string      `json:"_index"`
			//Type   string      `json:"_type"`
			//Id     string      `json:"_id"`
			//Score  interface{} `json:"_score"`
			Source struct {
				//CountryCode   string `json:"country_code"`
				//OpenvpnPid    string `json:"openvpn_pid"`
				OpenvpnReason string `json:"openvpn_reason"`
				//Input         struct {
				//	Type string `json:"type"`
				//} `json:"input"`
				//ClientAsn     int    `json:"client_asn"`
				//ClientCountry string `json:"client_country"`
				Timestamp string `json:"timestamp"`
				/*Clientip      string `json:"clientip"`
				Ecs           struct {
					Version string `json:"version"`
				} `json:"ecs"`
				Log struct {
					File struct {
						Path string `json:"path"`
					} `json:"file"`
					Offset int `json:"offset"`
				} `json:"log"`
				Openvpnmessage string `json:"openvpnmessage"`
				Message        string `json:"message"`
				ClientCity     string `json:"client_city"`
				OvpnIdentifier string `json:"ovpn.identifier"`
				SyslogProgram  string `json:"syslog_program"`
				Host           struct {
					Name string `json:"name"`
				} `json:"host"`*/
				SyslogHostname string `json:"syslog_hostname"`
				GeoipClient    struct {
					Ip    string `json:"ip"`
					AsOrg string `json:"as_org"`
				} `json:"geoip_client"`
				/*ReceivedFrom      []string `json:"received_from"`
				ClientGeoLocation struct {
					Lon float64 `json:"lon"`
					Lat float64 `json:"lat"`
				} `json:"client_geo_location"`
				SyslogMessage          string      `json:"syslog_message"`
				City                   string      `json:"city"`
				Timestamp1             time.Time   `json:"@timestamp"`
				Tags                   []string    `json:"tags"`
				Hostname               string      `json:"hostname"`
				OpenvpnCn              string      `json:"openvpn_cn"`
				ClientAutonomousSystem string      `json:"client_autonomous_system"`
				SyslogTimestamp        string      `json:"syslog_timestamp"`
				ReceivedAt             []time.Time `json:"received_at"`
				OpenvpnProgramResult   string      `json:"openvpn_program_result"`
				Agent                  struct {
					Hostname    string `json:"hostname"`
					Name        string `json:"name"`
					Version     string `json:"version"`
					EphemeralId string `json:"ephemeral_id"`
					Type        string `json:"type"`
					Id          string `json:"id"`
				} `json:"agent"`
				Country           string `json:"country"`
				ClientCountryCode string `json:"client_country_code"`
				SyslogPid         string `json:"syslog_pid"`
				Version           string `json:"@version"`
				OpenvpnProgram    string `json:"openvpn_program"`
				Clientport        string `json:"clientport"`
				Next              struct {
					Config struct {
						SecondaryChannel struct {
							Type string `json:"type"`
						} `json:"secondary_channel"`
						PrimaryChannel struct {
							Type string `json:"type"`
						} `json:"primary_channel"`
					} `json:"config"`
					Identifier string `json:"identifier"`
					Facts      struct {
						PrimaryWanIp   string `json:"primary_wan_ip"`
						SecondaryWanIp string `json:"secondary_wan_ip"`
					} `json:"facts"`
					Type string `json:"type"`
				} `json:"next"`*/
			} `json:"_source"`
			//Sort []int64 `json:"sort"`
		} `json:"hits"`
	} `json:"hits"`
}

func main() {

	url := "https://vpc-pls-opensearch-fhg6wf37b4r2juplvqx2mp4erm.eu-central-1.es.amazonaws.com/openvpn-*/_search"
	method := "POST"

	payload := strings.NewReader(`{
    "size": "3",
    "from": "0",
    "sort": [
        {
            "@timestamp": {
                "order": "desc",
                "unmapped_type": "boolean"
            }
        }
    ],
    "query": {
        "bool": {
            "filter": [
                {
                    "bool": {
                        "should": [
                            {
                                "match_phrase": {
                                    "openvpn_cn": "` + "643107D81220" + `"
                                }
                            }
                        ],
                        "minimum_should_match": 1
                    }
                }
            ]
        }
    }
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	auth := base64.StdEncoding.EncodeToString([]byte("next" + ":" + "J4NkbFuviix9jev_"))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data := &OSResponse{}
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	/*body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))*/
	fmt.Println(string(body))
	fmt.Println("-------------------")
	fmt.Println(data)
	fmt.Println("-------------------")
	for i := 0; i < len(data.Hits.Hits); i++ {
		fmt.Println(data.Hits.Hits[i].Source.OpenvpnReason, data.Hits.Hits[i].Source.Timestamp)
	}
}

/*openvpn_reason
timestamp

syslog_hostname
geoip_client*/
