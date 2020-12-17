package domain

type Resource struct {
	Search struct {
		Entry []struct {
			Dn        string `json:"dn"`
			Attribute []struct {
				Name  string   `json:"name"`
				Value []string `json:"value"`
			} `json:"attribute"`
		} `json:"entry"`
		Return struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Count   int    `json:"count"`
		} `json:"return"`
	} `json:"search"`
}
