package feed

type cdata_t struct {
	Text string `xml:",cdata"`
}

func cdata(str *string) *cdata_t {
	if str == nil {
		return nil
	}

	return &cdata_t{Text: *str}
}

func copy_str_ptr(str *string) *string {
	s := *str
	return &s
}
