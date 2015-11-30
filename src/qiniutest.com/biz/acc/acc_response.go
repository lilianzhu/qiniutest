package acc

type AccTokenResult struct {
	AccToken string `json:"access_token"`
	Expires  int    `json:"expires_in"`
	RefToken string `json:"refresh_token"`
}
