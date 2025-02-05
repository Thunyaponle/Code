package utils

func NumberToThai(num int) string {
	thaiNum := []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
	if num > 9 || num < 0 {
		return ""
	} else {
		return thaiNum[num]
	}
}
