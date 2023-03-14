package exporter

type ExtensionListResp struct {
	List []Extension `json:"list"`
}

type Extension struct {
	Id           int
	IsRegistered bool
	Number       string
	FirstName    string
	LastName     string
}

func (api *API) ExtensionList() ([]Extension, error) {
	var resp ExtensionListResp
	err := api.getResponse("ExtensionList", &resp)
	if err != nil {
		return nil, err
	}
	return resp.List, nil
}
