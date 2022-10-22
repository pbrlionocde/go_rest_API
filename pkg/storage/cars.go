package storage

type car struct {
	ID          int64  `json:"id"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Horse_Power int64  `json:"horse_power"`
}

type storageCars struct {
	Cars []car `json:"cars"`
}

var Storage = storageCars{
	Cars: []car{
		{
			ID:          1,
			Brand:       "Mitsubishi",
			Model:       "lancer",
			Horse_Power: 300,
		},
		{
			ID:          2,
			Brand:       "Subaru",
			Model:       "impreza",
			Horse_Power: 300,
		},
	},
}
