package data

type ProfileFields []ProfileField

// ProfileField is the struct behind a single field in fields of JSON raw message in public.user_profile
type ProfileField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
