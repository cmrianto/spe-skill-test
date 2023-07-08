package helper

import "os"

func GetWaresixCompanyId() string {
	return os.Getenv("WARESIX_COMPANY_ID") // get waresix company id value in config map
}
