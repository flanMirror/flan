package sdk

type Host string

func (h Host) IsLocal() bool {
	return h == ""
}
