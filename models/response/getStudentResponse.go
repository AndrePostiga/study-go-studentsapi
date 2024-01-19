package response

type GetStudentResponse struct {
	Id      string              `json:"id"`
	Name    string              `json:"name"`
	CPF     string              `json:"cpf"`
	Address *GetAddressResponse `json:"address,omitempty"`
}
