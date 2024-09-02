package npmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// NPMClient represents an Nginx Proxy Manager client.
type NPMClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewNPMClient creates a new NPMClient instance.
func NewNPMClient(baseURL, apiKey string) *NPMClient {
	return &NPMClient{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// CreateProxy creates a new proxy in NPM.
func (c *NPMClient) CreateProxy(proxy *Proxy) error {
	data, err := json.Marshal(proxy)
	log.Println(string(data))
	if err != nil {
		return fmt.Errorf("failed to marshal proxy: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+"/api/nginx/proxy-hosts", bytes.NewReader(data))
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *NPMClient) PlaceHolder() error {
	return nil
}

// Proxy represents a proxy configuration in NPM.
type Proxy struct {
	DomainNames           []string `json:"domain_names"`
	ForwardScheme         string   `json:"forward_scheme"`
	ForwardHost           string   `json:"forward_host"`
	ForwardPort           int      `json:"forward_port"`
	CachingEnabled        bool     `json:"caching_enabled,omitempty"`
	BlockExploits         bool     `json:"block_exploits,omitempty"`
	AllowWebsocketUpgrade bool     `json:"allow_websocket_upgrade,omitempty"`
	AccessListID          string   `json:"access_list_id,omitempty"`
	CertificateID         int      `json:"certificate_id"`
	SslForced             bool     `json:"ssl_forced,omitempty"`
	HTTP2Support          bool     `json:"http2_support,omitempty"`
	AdvancedConfig        string   `json:"advanced_config"`
}

// Example usage:
// func main() {
// 	npmClient := NewNPMClient("https://pewer.aj.athulcyriac.in", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkiLCJzY29wZSI6WyJ1c2VyIl0sImF0dHJzIjp7ImlkIjoxfSwiZXhwaXJlc0luIjoiMWQiLCJqdGkiOiJVSWlEa2dsSyIsImlhdCI6MTcyNTE5OTkwNywiZXhwIjoxNzI1Mjg2MzA3fQ.P05MRKVgydgr7ZWgiT7hwE4Jo8YkIXeGACRF2hUiv4s2YHgMhu6kWuvNTBFC96RfHgxFgXS7o497ak19x_htcWdsus_PBGF5l5OhSK4J5o7B9iTSrvQfWCDLEzjLPOvSTCevuoblUkLNTmhFGMxHSGa25jEE_y4a8RqrZWTr9D5pTS_moiteMVaYlGOapfxSvphMQmKPGn3Th9_Mz01xFaQNO6vKYxaAL9MKyuwfNj0jiTcuzr4cDJkslHe5a-KF1RxxxBUE2KtPkobWFg8Q-Wi5EgdnhrbH20uLslYrY7VLw4Q4F1Lli6-vGIMULOiXf-4i-Tlabd1aEf9gj0WL5A")
//
// 	proxy := &Proxy{
// 		DomainNames:    []string{"client.aj.athulcyriac.in"},
// 		ForwardScheme:  "http",
// 		ForwardHost:    "127.0.0.1",
// 		ForwardPort:    9090,
// 		CertificateID:  8,
// 		AdvancedConfig: "",
// 	}
//
// 	err := npmClient.CreateProxy(proxy)
// 	if err != nil {
// 		log.Println(err)
// 		fmt.Println("Error creating proxy:", err)
// 	} else {
// 		fmt.Println("Proxy created successfully!")
// 	}
// }
