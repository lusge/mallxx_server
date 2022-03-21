// Code generated by goctl. DO NOT EDIT.
package types

type Member struct {
	Id          int64  `json:"id,optional"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Phone       string `json:"phone,optional"`
	Email       string `json:"email,optional"`
	Avatar      string `json:"avatar,optional"`
	Status      int32  `json:"status,optional"`
	AddTime     int64  `json:"add_time,optional"`
	MemberLevel int64  `json:"member_level,optional"`
	Gender      int32  `json:"gender,optional"`
	Birthday    int32  `json:"birthday,optional"`
	Token       string `json:"token,optional" xorm:"-"`
	Integral    int32  `json:"integral,optional"`
	Coupon      int32  `json:"coupon,optional"`
	Followers   int32  `json:"followers,optional"`
	Nickname    string `json:"nickname,optional"`
}

type MemberLevel struct {
	Id                    int64   `json:"id,optional" db:"id"`
	GrowthPoint           int32   `json:"growth_point" db:"growth_point"`
	DefaultStatus         int32   `json:"default_status" db:"default_status"`
	FreeFreightPoint      float32 `json:"free_freight_point" db:"free_freight_point"`
	CommentGrowthPoint    int32   `json:"comment_growth_point" db:"comment_growth_point"`
	PriviledgeFreeFreight int32   `json:"priviledge_free_freight" db:"priviledge_free_freight"`
	PriviledgeSignIn      int32   `json:"priviledge_sign_in" db:"priviledge_sign_in"`
	PriviledgeComment     int32   `json:"priviledge_comment" db:"priviledge_comment"`
	PriviledgePromotion   int32   `json:"priviledge_promotion" db:"priviledge_promotion"`
	PriviledgeMemberPrice int32   `json:"priviledge_member_price" db:"priviledge_member_price"`
	PriviledgeBirthday    int32   `json:"priviledge_birthday" db:"priviledge_birthday"`
	Note                  string  `json:"note" db:"note"`
	Name                  string  `json:"name" db:"name"`
}

type MemberData struct {
	Member      *Member      `json:"member"`
	MemberLevel *MemberLevel `json:"member_level"`
}

type MemberResponse struct {
	Code   int32       `json:"code"`
	Data   *MemberData `json:"data"`
	Detail string      `json:"detail"`
}

type MemberLoginRequest struct {
	Username string `json:"username,optional"`
	Password string `json:"password,optional"`
	Phone    string `json:"phone,optional"`
	Email    string `json:"email,optional"`
}

type Response struct {
	Code   int32  `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
}
