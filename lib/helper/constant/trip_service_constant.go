package constant

func GetAssignmentReasonLabelByCaetgory(category string) *string {
	reasonMap := map[string]string{
		// cancel vendor
		"me_duplicate":          "Duplikasi DO atau salah input",
		"me_wrong_date":         "Ingin mengubah tanggal pengiriman",
		"vendor_not_coming":     "Vendor tidak datang",
		"vendor_too_far":        "Lokasi driver terlalu jauh",
		"vendor_not_responding": "Vendor tidak merespon",
		// reject vendor
		"price_not_match":      "Harga Tidak Sesuai",
		"truck_not_available":  "Truk Tidak Tersedia",
		"driver_not_available": "Driver Sedang Tidak Tersedia",
		"not_route_preference": "Bukan Rute Trip Preferensi",
		// reject driver
		"not_available":                         "Tanggal itu ada kerjaan lain",
		"location_too_far":                      "Lokasi kejauhan",
		"not_enough_cost":                       "Uang jalan tidak cukup",
		"driver_sick":                           "Sedang Sakit",
		"driver_already_has_assignment_on_date": "Ada kerjaan lain",
		// cancel driver
		"driver_not_responding": "Driver terlalu lama merespon",
		"item_not_available":    "Barang belum tersedia",
		"driver_too_far":        "Lokasi driver terlalu jauh",
		"location_changed":      "Ganti lokasi muat bongkar",
		"vehicle_trouble":       "Truk tiba-tiba bermasalah",
		"cannot_be_fulfilled":   "Permintaan khusus tidak bisa dipenuhi",
		"accidentally_taken":    "Kerjaan tidak sengaja diambil",
		// cancel truck
		"truck_spec_not_match": "Spesifikasi truk tidak sesuai",
	}
	if v, ok := reasonMap[category]; ok {
		return &v
	}
	return nil
}
