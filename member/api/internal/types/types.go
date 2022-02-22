// Code generated by goctl. DO NOT EDIT.
package types

type ListRequest struct {
	PageNum  int32  `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty" form:"pageNum"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" form:"pageSize"`
	Keyword  string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty" form:"keyword"`
	Pid      int64  `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty" form:"pid"`
}

type Response struct {
	Code   int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

type EmptyRequest struct {
}

type FollowerResponse struct {
	Code   int32     `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Data   []*Member `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
	Detail string    `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail"`
	Total  int32     `protobuf:"varint,4,opt,name=total,proto3" json:"total"`
}

type MemberLevel struct {
	Id                    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id" db:"id"`
	GrowthPoint           int32   `protobuf:"varint,2,opt,name=growth_point,json=growthPoint,proto3" json:"growth_point" db:"growth_point"`
	DefaultStatus         int32   `protobuf:"varint,3,opt,name=default_status,json=defaultStatus,proto3" json:"default_status" db:"default_status"`
	FreeFreightPoint      float32 `protobuf:"fixed32,4,opt,name=free_freight_point,json=freeFreightPoint,proto3" json:"free_freight_point" db:"free_freight_point"`
	CommentGrowthPoint    int32   `protobuf:"varint,5,opt,name=comment_growth_point,json=commentGrowthPoint,proto3" json:"comment_growth_point" db:"comment_growth_point"`
	PriviledgeFreeFreight int32   `protobuf:"varint,6,opt,name=priviledge_free_freight,json=priviledgeFreeFreight,proto3" json:"priviledge_free_freight" db:"priviledge_free_freight"`
	PriviledgeSignIn      int32   `protobuf:"varint,7,opt,name=priviledge_sign_in,json=priviledgeSignIn,proto3" json:"priviledge_sign_in" db:"priviledge_sign_in"`
	PriviledgeComment     int32   `protobuf:"varint,8,opt,name=priviledge_comment,json=priviledgeComment,proto3" json:"priviledge_comment" db:"priviledge_comment"`
	PriviledgePromotion   int32   `protobuf:"varint,9,opt,name=priviledge_promotion,json=priviledgePromotion,proto3" json:"priviledge_promotion" db:"priviledge_promotion"`
	PriviledgeMemberPrice int32   `protobuf:"varint,10,opt,name=priviledge_member_price,json=priviledgeMemberPrice,proto3" json:"priviledge_member_price" db:"priviledge_member_price"`
	PriviledgeBirthday    int32   `protobuf:"varint,11,opt,name=priviledge_birthday,json=priviledgeBirthday,proto3" json:"priviledge_birthday" db:"priviledge_birthday"`
	Note                  string  `protobuf:"bytes,12,opt,name=note,proto3" json:"note" db:"note"`
	Name                  string  `protobuf:"bytes,13,opt,name=name,proto3" json:"name" db:"name"`
}

type MemverLevelInfoRequest struct {
}

type MemverLeveListResponse struct {
	Code   int32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data   []*MemberLevel `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Detail string         `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
}

type Member struct {
	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,optional"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username"`
	Password    string `protobuf:"bytes,3,opt,name=password,proto3" json:"password"`
	Phone       string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone"`
	Email       string `protobuf:"bytes,5,opt,name=email,proto3" json:"email"`
	Avatar      string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar"`
	Status      int32  `protobuf:"varint,7,opt,name=status,proto3" json:"status"`
	AddTime     int64  `protobuf:"varint,8,opt,name=add_time,json=addTime,proto3" json:"add_time"`
	MemberLevel int64  `protobuf:"varint,9,opt,name=member_level,json=memberLevel,proto3" json:"member_level"`
	Gender      int32  `protobuf:"varint,10,opt,name=gender,proto3" json:"gender"`
	Birthday    int32  `protobuf:"varint,11,opt,name=birthday,proto3" json:"birthday"`
	Token       string `protobuf:"bytes,12,opt,name=token,proto3" json:"token" xorm:"-"`
	Integral    int32  `protobuf:"varint,13,opt,name=integral,proto3" json:"integral"`
	Coupon      int32  `protobuf:"varint,14,opt,name=coupon,proto3" json:"coupon"`
	Followers   int32  `protobuf:"varint,15,opt,name=followers,proto3" json:"followers"`
	Nickname    string `protobuf:"bytes,16,opt,name=nickname,proto3" json:"nickname"`
}

type MemberData struct {
	Member      *Member      `protobuf:"bytes,1,opt,name=member,proto3" json:"member"`
	MemberLevel *MemberLevel `protobuf:"bytes,2,opt,name=member_level,json=memberLevel,proto3" json:"member_level"`
}

type MemberResponse struct {
	Code   int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Data   *MemberData `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
	Detail string      `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail"`
}

type MemberListResponse struct {
	Code   int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Data   []*MemberData `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
	Detail string        `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail"`
	Total  int32         `protobuf:"varint,4,opt,name=total,proto3" json:"total"`
}

type MemberRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid" form:"uid"`
}

type MemberLoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email"`
}

type ReceiveAddress struct {
	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	MemberId      int64  `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	Phone         string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone"`
	DefaultStatus int32  `protobuf:"varint,4,opt,name=default_status,json=defaultStatus,proto3" json:"default_status"`
	PostCode      string `protobuf:"bytes,5,opt,name=post_code,json=postCode,proto3" json:"post_code"`
	Province      string `protobuf:"bytes,6,opt,name=province,proto3" json:"province"`
	City          string `protobuf:"bytes,7,opt,name=city,proto3" json:"city"`
	Region        string `protobuf:"bytes,8,opt,name=region,proto3" json:"region"`
	Detail        string `protobuf:"bytes,9,opt,name=detail,proto3" json:"detail"`
}

type ReceiveAddressResponse struct {
	Code   int32           `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Data   *ReceiveAddress `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
	Detail string          `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail"`
}

type ReceiveAddressListResponse struct {
	Code   int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Data   []*ReceiveAddress `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
	Detail string            `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail"`
	Total  int32             `protobuf:"varint,4,opt,name=total,proto3" json:"total"`
}

type ReceiveAddressRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"member_id"`
}
