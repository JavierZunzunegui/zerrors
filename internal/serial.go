package internal

var (
	basic  func(error) string
	detail func(error) string
)

func SetBasic(f func(error) string) {
	if f == nil {
		return
	}
	basic = f
}

func Basic(err error) string {
	return basic(err)
}

func SetDetail(f func(error) string) {
	if f == nil {
		return
	}
	detail = f
}

func Detail(err error) string {
	return detail(err)
}
