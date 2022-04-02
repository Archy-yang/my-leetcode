package cn

func addBinary(a string, b string) string {
	s := ""
	al := len(a) - 1
	bl := len(b) - 1
	f := 0
	for {
		ai := 0
		if al >= 0 {
			ai = int(a[al]) - '0'
		}
		bi := 0
		if bl >= 0 {
			bi = int(b[bl]) - '0'
		}

		if bi+ai+f == 0 {
			s = "0" + s
			f = 0
		} else if bi+ai+f == 1 {
			s = "1" + s
			f = 0
		} else if bi+ai+f == 2 {
			s = "0" + s
			f = 1
		} else if bi+ai+f == 3 {
			s = "1" + s
			f = 1
		}

		if al <= 0 && bl <= 0 {
			break
		}
		al--
		bl--
	}
	if f > 0 {
		s = "1" + s
	}
	return s
}
