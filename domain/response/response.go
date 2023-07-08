package response

type Ping struct {
	Message string `json:"message"`
}

type NarcissisticNumberResponse struct {
	IsNarcissisticNumber bool
}

type ParityOutlierResponse struct {
	Outlier string `json:"outlier"`
}

type NeedleInHaystackResponse struct {
	Index int64 `json:"index"`
}

type BlueOceanResponse struct {
	BlueOcean []int64 `json:"blue_ocean"`
}
