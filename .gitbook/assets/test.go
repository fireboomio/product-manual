package test

type Admin__AppVersion__CreateOneAppResponseData struct {
	Data Admin__AppVersion__CreateOneAppResponseData_data `json:"data,omitempty"`
}
type Admin__AppVersion__CreateOneAppResponseData_data struct {
	Id string `json:"id"`
}
type Admin__AppVersion__DeleteManyAppInput struct {
	Ids []string `json:"ids"`
}
type Admin__AppVersion__DeleteManyAppInternalInput struct {
	Ids []string `json:"ids"`
}
type Admin__AppVersion__DeleteManyAppResponseData struct {
	Data Admin__AppVersion__DeleteManyAppResponseData_data `json:"data,omitempty"`
}
type Admin__AppVersion__DeleteManyAppResponseData_data struct {
	Count int64 `json:"count"`
}
type Admin__AppVersion__UpdateOneAppInput struct {
	Description string           `json:"description,omitempty"`
	DownloadUrl string           `json:"downloadUrl,omitempty"`
	Id          string           `json:"id"`
	IsForce     bool             `json:"isForce,omitempty"`
	Latest      bool             `json:"latest,omitempty"`
	Type        Freetalk_AppType `json:"type,omitempty"`
	Version     string           `json:"version,omitempty"`
}
type Admin__AppVersion__UpdateOneAppInternalInput struct {
	Version     string           `json:"version,omitempty"`
	Description string           `json:"description,omitempty"`
	DownloadUrl string           `json:"downloadUrl,omitempty"`
	Id          string           `json:"id"`
	IsForce     bool             `json:"isForce,omitempty"`
	Latest      bool             `json:"latest,omitempty"`
	Type        Freetalk_AppType `json:"type,omitempty"`
	UpdateTime  string           `json:"updateTime"`
}
type Admin__AppVersion__UpdateOneAppResponseData struct {
	Data Admin__AppVersion__UpdateOneAppResponseData_data `json:"data,omitempty"`
}
type Admin__AppVersion__UpdateOneAppResponseData_data struct {
	Id string `json:"id"`
}

type Freetalk_BoolFilter struct {
	Equals bool                       `json:"equals,omitempty"`
	Not    *Freetalk_NestedBoolFilter `json:"not,omitempty"`
}

type Freetalk_DateTimeFilter struct {
	Gt     string                         `json:"gt,omitempty"`
	Gte    string                         `json:"gte,omitempty"`
	In     []string                       `json:"in,omitempty"`
	Lt     string                         `json:"lt,omitempty"`
	Lte    string                         `json:"lte,omitempty"`
	Not    *Freetalk_NestedDateTimeFilter `json:"not,omitempty"`
	NotIn  []string                       `json:"notIn,omitempty"`
	Equals string                         `json:"equals,omitempty"`
}
type Freetalk_DateTimeNullableFilter struct {
	Equals string                                 `json:"equals,omitempty"`
	Gt     string                                 `json:"gt,omitempty"`
	Gte    string                                 `json:"gte,omitempty"`
	In     []string                               `json:"in,omitempty"`
	Lt     string                                 `json:"lt,omitempty"`
	Lte    string                                 `json:"lte,omitempty"`
	Not    *Freetalk_NestedDateTimeNullableFilter `json:"not,omitempty"`
	NotIn  []string                               `json:"notIn,omitempty"`
}

type Freetalk_IntNullableFilter struct {
	Equals int64                             `json:"equals,omitempty"`
	Gt     int64                             `json:"gt,omitempty"`
	Gte    int64                             `json:"gte,omitempty"`
	In     []int64                           `json:"in,omitempty"`
	Lt     int64                             `json:"lt,omitempty"`
	Lte    int64                             `json:"lte,omitempty"`
	Not    *Freetalk_NestedIntNullableFilter `json:"not,omitempty"`
	NotIn  []int64                           `json:"notIn,omitempty"`
}

type Freetalk_NestedBoolFilter struct {
	Equals bool                       `json:"equals,omitempty"`
	Not    *Freetalk_NestedBoolFilter `json:"not,omitempty"`
}
type Freetalk_NestedDateTimeFilter struct {
	Not    *Freetalk_NestedDateTimeFilter `json:"not,omitempty"`
	NotIn  []string                       `json:"notIn,omitempty"`
	Equals string                         `json:"equals,omitempty"`
	Gt     string                         `json:"gt,omitempty"`
	Gte    string                         `json:"gte,omitempty"`
	In     []string                       `json:"in,omitempty"`
	Lt     string                         `json:"lt,omitempty"`
	Lte    string                         `json:"lte,omitempty"`
}
type Freetalk_NestedDateTimeNullableFilter struct {
	Lt     string                                 `json:"lt,omitempty"`
	Lte    string                                 `json:"lte,omitempty"`
	Not    *Freetalk_NestedDateTimeNullableFilter `json:"not,omitempty"`
	NotIn  []string                               `json:"notIn,omitempty"`
	Equals string                                 `json:"equals,omitempty"`
	Gt     string                                 `json:"gt,omitempty"`
	Gte    string                                 `json:"gte,omitempty"`
	In     []string                               `json:"in,omitempty"`
}
type Freetalk_NestedIntNullableFilter struct {
	Equals int64                             `json:"equals,omitempty"`
	Gt     int64                             `json:"gt,omitempty"`
	Gte    int64                             `json:"gte,omitempty"`
	In     []int64                           `json:"in,omitempty"`
	Lt     int64                             `json:"lt,omitempty"`
	Lte    int64                             `json:"lte,omitempty"`
	Not    *Freetalk_NestedIntNullableFilter `json:"not,omitempty"`
	NotIn  []int64                           `json:"notIn,omitempty"`
}
type Freetalk_NestedStringFilter struct {
	StartsWith string                       `json:"startsWith,omitempty"`
	Equals     string                       `json:"equals,omitempty"`
	Gt         string                       `json:"gt,omitempty"`
	Lt         string                       `json:"lt,omitempty"`
	Lte        string                       `json:"lte,omitempty"`
	NotIn      []string                     `json:"notIn,omitempty"`
	Contains   string                       `json:"contains,omitempty"`
	EndsWith   string                       `json:"endsWith,omitempty"`
	Gte        string                       `json:"gte,omitempty"`
	In         []string                     `json:"in,omitempty"`
	Not        *Freetalk_NestedStringFilter `json:"not,omitempty"`
}
type Freetalk_NestedStringNullableFilter struct {
	Not        *Freetalk_NestedStringNullableFilter `json:"not,omitempty"`
	Contains   string                               `json:"contains,omitempty"`
	Gte        string                               `json:"gte,omitempty"`
	Gt         string                               `json:"gt,omitempty"`
	In         []string                             `json:"in,omitempty"`
	Lt         string                               `json:"lt,omitempty"`
	Lte        string                               `json:"lte,omitempty"`
	NotIn      []string                             `json:"notIn,omitempty"`
	StartsWith string                               `json:"startsWith,omitempty"`
	EndsWith   string                               `json:"endsWith,omitempty"`
	Equals     string                               `json:"equals,omitempty"`
}
type Freetalk_NestedUuidFilter struct {
	Not    *Freetalk_NestedUuidFilter `json:"not,omitempty"`
	NotIn  []string                   `json:"notIn,omitempty"`
	Equals string                     `json:"equals,omitempty"`
	Gt     string                     `json:"gt,omitempty"`
	Gte    string                     `json:"gte,omitempty"`
	In     []string                   `json:"in,omitempty"`
	Lt     string                     `json:"lt,omitempty"`
	Lte    string                     `json:"lte,omitempty"`
}

type Freetalk_StringFilter struct {
	EndsWith   string                       `json:"endsWith,omitempty"`
	Mode       Freetalk_QueryMode           `json:"mode,omitempty"`
	Not        *Freetalk_NestedStringFilter `json:"not,omitempty"`
	StartsWith string                       `json:"startsWith,omitempty"`
	Contains   string                       `json:"contains,omitempty"`
	Equals     string                       `json:"equals,omitempty"`
	Gt         string                       `json:"gt,omitempty"`
	Gte        string                       `json:"gte,omitempty"`
	In         []string                     `json:"in,omitempty"`
	Lt         string                       `json:"lt,omitempty"`
	Lte        string                       `json:"lte,omitempty"`
	NotIn      []string                     `json:"notIn,omitempty"`
}
type Freetalk_StringNullableFilter struct {
	Contains   string                               `json:"contains,omitempty"`
	Equals     string                               `json:"equals,omitempty"`
	Gte        string                               `json:"gte,omitempty"`
	Lt         string                               `json:"lt,omitempty"`
	EndsWith   string                               `json:"endsWith,omitempty"`
	Gt         string                               `json:"gt,omitempty"`
	In         []string                             `json:"in,omitempty"`
	Lte        string                               `json:"lte,omitempty"`
	Mode       Freetalk_QueryMode                   `json:"mode,omitempty"`
	Not        *Freetalk_NestedStringNullableFilter `json:"not,omitempty"`
	NotIn      []string                             `json:"notIn,omitempty"`
	StartsWith string                               `json:"startsWith,omitempty"`
}

type Freetalk_UuidFilter struct {
	Equals string                     `json:"equals,omitempty"`
	Gt     string                     `json:"gt,omitempty"`
	In     []string                   `json:"in,omitempty"`
	Not    *Freetalk_NestedUuidFilter `json:"not,omitempty"`
	Gte    string                     `json:"gte,omitempty"`
	Lt     string                     `json:"lt,omitempty"`
	Lte    string                     `json:"lte,omitempty"`
	Mode   Freetalk_QueryMode         `json:"mode,omitempty"`
	NotIn  []string                   `json:"notIn,omitempty"`
}

type Freetalk_AppType string

const (
	Freetalk_AppType_Android Freetalk_AppType = "Android"
	Freetalk_AppType_IOS     Freetalk_AppType = "IOS"
)

type Freetalk_QueryMode string

const (
	Freetalk_QueryMode_default     Freetalk_QueryMode = "default"
	Freetalk_QueryMode_insensitive Freetalk_QueryMode = "insensitive"
)
