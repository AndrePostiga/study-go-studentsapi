package externalservices

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	externalservicescontracts "github.com/andrepostiga/api-go-gin/domain/externalServicesContracts"
	valueobjects "github.com/andrepostiga/api-go-gin/domain/valueObjects"
)

const DefaultTimeout = 10 * time.Second
const path = "/cep/v2/"

type Client struct {
	baseURL *url.URL
	cli     *http.Client
}

type Cep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func NewCepClient(baseUrl string) externalservicescontracts.ICepClient {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatalf("Invalid base URL: %s", err)
	}

	httpClient := &http.Client{
		Timeout: DefaultTimeout,
	}

	return &Client{
		cli:     httpClient,
		baseURL: parsedUrl,
	}
}

// Get implements externalservicescontracts.ICepClient.
func (c *Client) Get(cep string) (*valueobjects.Cep, error) {
	resourcePath := c.baseURL.String() + path + cep

	req, err := http.NewRequest("GET", resourcePath, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var decodedResponse Cep
	if err := json.NewDecoder(response.Body).Decode(&decodedResponse); err != nil {
		return nil, err
	}

	return &valueobjects.Cep{
		Cep:          decodedResponse.Cep,
		City:         decodedResponse.City,
		State:        decodedResponse.State,
		Neighborhood: decodedResponse.Neighborhood,
		Street:       decodedResponse.Street,
	}, nil
}
