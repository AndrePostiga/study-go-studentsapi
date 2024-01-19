package response

import studentaggregate "github.com/andrepostiga/api-go-gin/domain/entities/studentAggregate"

type GetAddressResponse struct {
	CEP          string `json:"cep"`
	Number       uint   `json:"number"`
	Complement   string `json:"complement,omitempty"` // omitempty to exclude if empty
	Country      string `json:"country"`
	StudentId    uint   `json:"studentId"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func NewAddressResponse(a *studentaggregate.Address) *GetAddressResponse {
	resp := GetAddressResponse{}
	if a == nil {
		return nil
	}

	resp.Number = a.Number
	resp.Complement = a.Complement
	resp.Country = a.Country
	resp.CEP = a.Cep.Cep
	resp.State = a.State
	resp.City = a.City
	resp.Neighborhood = a.Neighborhood
	resp.Street = a.Street

	return &resp
}
