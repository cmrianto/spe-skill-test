package constant

type ConstantReason struct {
	Key   string
	Value string
}

var TransporterAppsAssignmentReasons = []*ConstantReason{
	{Key: "price_not_match", Value: "Harga Tidak Sesuai"},
	{Key: "truck_not_available", Value: "Truk Tidak Tersedia"},
	{Key: "driver_not_available", Value: "Driver Sedang Tidak Tersedia"},
	{Key: "not_route_preference", Value: "Bukan Rute Preferensi"},
	{Key: "others", Value: "Alasan Lainnya"},
}

var DriverRejectTripReasons = []*ConstantReason{
	{Key: "driver_already_has_assignment_on_date", Value: "Ada kerjaan lain"},
	{Key: "location_too_far", Value: "Lokasi kejauhan"},
	{Key: "not_enough_cost", Value: "Uang jalan tidak cukup"},
	{Key: "driver_sick", Value: "Sedang sakit"},
	{Key: "me_others", Value: "Alasan lainnya"},
}

var CancelDriverReasons = []*ConstantReason{
	{Key: "driver_not_responding", Value: "Driver terlalu lama merespon"},
	{Key: "item_not_available", Value: "Barang belum tersedia"},
	{Key: "driver_too_far", Value: "Lokasi driver terlalu jauh"},
	{Key: "location_changed", Value: "Ganti lokasi muat bongkar"},
	{Key: "me_others", Value: "Alasan lainnya"},
}

var CancelTruckReasons = []*ConstantReason{
	{Key: "truck_spec_not_match", Value: "Spesifikasi truk tidak sesuai"},
	{Key: "item_not_available", Value: "Barang belum tersedia"},
	{Key: "location_changed", Value: "Ganti lokasi trip"},
	{Key: "me_others", Value: "Alasan lainnya"},
}
