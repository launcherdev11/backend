package schemas

type RequestUUID_by_name struct {
	Username string `json:"username" validate:"required,min=3,max=128"`
	Time     int64  `json:"at"` //timestamp
}

type ResponseUUID_by_name struct {
	UUID   string `json:"id" validate:"required" pattern:"^[0-9a-f]{32}$"`
	Name   string `json:"name" validate:"required,min=3,max=128"`
	Lagacy bool   `json:"legacy"` //exist only if true
	Demo   bool   `json:"demo"`   //exist only if true
}

type Response400 struct {
	Error        string `json:"error" validate:"required"`
	ErrorMessage string `json:"errorMessage" validate:"required"`
}

type UUIDrequest struct {
	UUID string `json:"id" validate:"required" pattern:"^[0-9a-f]{32}$"`
}

type ResponseName struct {
	Name        string `json:"name" validate:"required,min=3,max=128"`
	ChangedToAt int64  `json:"changedToAt"` //ms timestamp
}

type ResponseNames []ResponseName

type RequestProfile struct {
	UUID     string `json:"id" validate:"required" pattern:"^[0-9a-f]{32}$"`
	Unsigned bool   `query:"unsigned"`
}

type TextureProperty struct {
	Name      string `json:"name" validate:"required,eq=textures"`
	Value     string `json:"value" validate:"required,base64"`
	Signature string `json:"signature,omitempty"` // only if unsigned=false
}
type ResponseProfile struct {
	UUID       string            `json:"id" validate:"required"`
	Name       string            `json:"name" validate:"required"`
	Properties []TextureProperty `json:"properties" validate:"required"`
}
