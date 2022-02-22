package models

import (
	"mallxx_server/product/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

var (
	productCacheUsernamePrefix = "mallxxadmin:product:%s"
)

type Product struct {
	engine *xorm.Engine
	table  string
}

func NewProduct(engine *xorm.Engine) *Product {
	return &Product{
		table:  "product",
		engine: engine,
	}
}

func (m *Product) Insert(in *pb.Product) bool {
	session := m.engine.NewSession()
	defer session.Close()

	err := session.Commit()
	if err != nil {
		logx.Error(err)
		return false
	}

	if in.PromotionStartTime == "" {
		in.PromotionStartTime = time.Now().Format("2006-01-02 15:04:05")
	}

	if in.PromotionEndTime == "" {
		in.PromotionEndTime = time.Now().Format("2006-01-02 15:04:05")
	}

	//insert product
	_, err = session.InsertOne(in)

	if err != nil {
		logx.Error(err)
		session.Rollback()
		return false
	}

	lastId := in.Id
	//insert sku stock
	skuOk := NewSkuStock(m.engine).InsertListBySession(session, lastId, in.SkuStock)
	if !skuOk {
		session.Rollback()
		return false
	}

	//insert member price
	memberPriceIsOk := NewMemberPrice(m.engine).InsertListBySession(session, lastId, in.MemberPrice)
	if !memberPriceIsOk {
		session.Rollback()
		return false
	}

	//insert product_attribute_value
	attrValueIsOk := NewProductAttributeValue(m.engine).InsertListBySession(session, lastId, in.ProductAttrValue)
	if !attrValueIsOk {
		session.Rollback()
		return false
	}

	//insert Product Ladder
	ladderIsok := NewProductLadder(m.engine).InsertListBySession(session, lastId, in.ProductLadder)
	if !ladderIsok {
		session.Rollback()
		return false
	}

	//insert ProductFullReduction
	reductionIsOk := NewProductFullReduction(m.engine).InsertListBySession(session, lastId, in.ProductFullReduction)
	if !reductionIsOk {
		session.Rollback()
		return false
	}

	if err = session.Commit(); err != nil {
		logx.Error(err)
		return false
	}

	in.Id = lastId
	return true
}

func (m *Product) Update(in *pb.Product) bool {
	session := m.engine.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		logx.Error(err)
		return false
	}

	_, err = session.Where("id = ?", in.Id).Update(in)
	if err != nil {
		session.Rollback()
		logx.Error(err)
		return false
	}

	//delete sku stock
	skuModel := NewSkuStock(m.engine)
	skuIsDel := skuModel.DeleteByProductIdAndSession(session, in.Id)
	if !skuIsDel {
		session.Rollback()
		return false
	}
	//inser sku stock
	skuIsInserted := skuModel.InsertListBySession(session, in.Id, in.SkuStock)
	if !skuIsInserted {
		session.Rollback()
		return false
	}

	//delete member price
	memberPriceModel := NewMemberPrice(m.engine)
	memIsDel := memberPriceModel.DeleteBySession(session, in.Id)
	if !memIsDel {
		session.Rollback()
		return false
	}
	// insert member price
	memIsInserted := memberPriceModel.InsertListBySession(session, in.Id, in.MemberPrice)
	if !memIsInserted {
		session.Rollback()
		return false
	}

	//delete product_attribute_value
	productAttrValueModel := NewProductAttributeValue(m.engine)
	pavIsDel := productAttrValueModel.DeleteBySession(session, in.Id)
	if !pavIsDel {
		session.Rollback()
		return false
	}
	// insert product_attribute_value
	pavIsInserted := productAttrValueModel.InsertListBySession(session, in.Id, in.ProductAttrValue)
	if !pavIsInserted {
		session.Rollback()
		return false
	}

	//delete Product Ladder
	plModel := NewProductLadder(m.engine)
	plIsDel := plModel.DeleteByProductIdAndSession(session, in.Id)
	if !plIsDel {
		session.Rollback()
		return false
	}
	//insert Product Ladder
	plIsInserted := plModel.InsertListBySession(session, in.Id, in.ProductLadder)
	if !plIsInserted {
		session.Rollback()
		return false
	}

	pfrModel := NewProductFullReduction(m.engine)
	//delete ProductFullReduction
	pfrIsDel := pfrModel.DeleteByProductIdAndSession(session, in.Id)
	if !pfrIsDel {
		session.Rollback()
		return false
	}
	//insert ProductFullReduction
	reductionIsInserted := pfrModel.InsertListBySession(session, in.Id, in.ProductFullReduction)
	if !reductionIsInserted {
		session.Rollback()
		return false
	}

	if err = session.Commit(); err != nil {
		logx.Error(err)
		return false
	}

	return true
}

func (m *Product) FindOne(id int64) *pb.Product {
	product := new(pb.Product)
	// m.engine.ShowSQL(true)
	has, err := m.engine.Where("id = ?", id).Get(product)
	if err != nil {
		logx.Error(err)
		return nil
	}

	if has == false {
		return nil
	}

	return product
}

func (m *Product) FindAll(in *pb.ProductListRequest, start, size int32) []*pb.Product {
	list := make([]*pb.Product, 0)
	db := m.engine.Where("delete_status = ?", 0)

	if len(in.Name) > 0 {
		db.And("name LIKE  ? ", "%"+in.Name+"%")
	}

	if len(in.Sn) > 0 {
		db.And("product_sn = ?", in.Sn)
	}

	if in.CategoryId > 0 {
		db.And("product_category_id = ?", in.CategoryId)
	}

	if in.BrandId > 0 {
		db.And("brand_id = ?", in.BrandId)
	}

	if in.PublishStatus > 0 {
		db.And("publish_status = ?", in.PublishStatus)
	}

	if in.VerifyStatus > 0 {
		db.And("verify_status = ?", in.VerifyStatus)
	}

	if in.RecommendStatus > 0 {
		db.And("recommend_status = ?", in.RecommendStatus)

	}

	if in.NewStatus > 0 {
		db.And("new_status = ?", in.NewStatus)
	}

	err := db.Limit(int(size), int(start)).Find(&list)

	if err != nil {
		logx.Error(err)
		return nil
	}

	return list
}

func (m *Product) Count(in *pb.ProductListRequest) int64 {
	db := m.engine.Where("delete_status = ?", 0)

	if len(in.Name) > 0 {
		db.And("name LIKE  ? ", "%"+in.Name+"%")
	}

	if len(in.Sn) > 0 {
		db.And("product_sn = ?", in.Sn)
	}

	if in.CategoryId > 0 {
		db.And("product_category_id = ?", in.CategoryId)
	}

	if in.BrandId > 0 {
		db.And("brand_id = ?", in.BrandId)
	}

	if in.PublishStatus > 0 {
		db.And("publish_status = ?", in.PublishStatus)
	}

	if in.VerifyStatus > 0 {
		db.And("verify_status = ?", in.VerifyStatus)
	}

	if in.RecommendStatus > 0 {
		db.And("recommend_status = ?", in.RecommendStatus)

	}

	if in.NewStatus > 0 {
		db.And("new_status = ?", in.NewStatus)
	}

	total, err := db.Table(m.table).Count()
	if err != nil {
		logx.Error(err)
		return 0
	}

	return total
}

func (m *Product) ChangeStatus(id int64, params map[string]interface{}) bool {
	affected, err := m.engine.Table(m.table).Where("id = ?", id).Update(params)
	if err != nil {
		logx.Error(err)
		return false
	}

	if affected <= 0 {
		logx.Error("update error , affected row 0")
		return false
	}

	return true
}

func (m *Product) FinaApiPage(in *pb.ProductApiRequest, size, start int32) []*pb.Product {
	list := make([]*pb.Product, 0)
	db := m.engine.Where("delete_status = ?", 0)
	db.And("verify_status = ?", 1)

	if in.Sort == "def" {
		db.Desc("sort")
	} else if in.Sort == "sale" {
		db.Desc("sale")
	} else if in.Sort == "price+" {
		db.Desc("price")
	} else if in.Sort == "price-" {
		db.Asc("price")
	}

	if len(in.CategoryIds) > 0 {
		db.In("product_category_id", in.CategoryIds)
	}

	if len(in.BrandIds) > 0 {
		db.In("brand_id", in.BrandIds)
	}

	if in.MinPrice > 0 {
		db.And("price >= ?", in.MinPrice)
	}

	if in.MixPrice > 0 {
		db.And("price <= ?", in.MixPrice)
	}

	err := db.Limit(int(size), int(start)).Find(&list)
	if err != nil {
		logx.Error(err)
		return nil
	}

	return list
}

func (m *Product) FindByIds(ids []int64) []*pb.Product {
	list := make([]*pb.Product, 0)
	if len(ids) <= 0 {
		return nil
	}

	err := m.engine.In("id", ids).Find(&list)
	if err != nil {
		logx.Error(err)
		return nil
	}
	return list
}
