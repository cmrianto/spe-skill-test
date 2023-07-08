package request

type NarcissisticNumberRequest struct {
	Number string `json:"number"`
}

type ParityOutlierRequest struct {
	Numbers []string `json:"numbers"`
}

type NeedleInHaystackRequest struct {
	Haystack []string `json:"haystack"`
	Needle   string   `json:"needle"`
}

type BlueOceanRequest struct {
	BlueOcean []int64 `json:"blue_ocean"`
	Remove    []int64 `json:"remove"`
}
