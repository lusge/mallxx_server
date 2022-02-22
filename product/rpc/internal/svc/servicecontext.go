package svc

import (
	"mallxx_server/common/mysqlx"
	"mallxx_server/product/models"
	"mallxx_server/product/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config                                config.Config
	RedisClient                           *redis.Redis
	BrandModel                            *models.Brand
	CategoryModel                         *models.Category
	CategoryAttributeModel                *models.CategoryAttribute
	MemberPriceModel                      *models.MemberPrice
	ProductAttributeModel                 *models.ProductAttribute
	ProductModel                          *models.Product
	ProductAttributeValueModel            *models.ProductAttributeValue
	ProductCategoryAttributeRelationModel *models.ProductCategoryAttributeRelation
	ProductFullReductionModel             *models.ProductFullReduction
	ProductLadderModel                    *models.ProductLadder
	ProductRecommendModel                 *models.ProductRecommend
	SkuStockModel                         *models.SkuStock
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := mysqlx.InitMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                                c,
		RedisClient:                           redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass),
		BrandModel:                            models.NewBrand(engine),
		CategoryModel:                         models.NewCategory(engine),
		CategoryAttributeModel:                models.NewCategoryAttribute(engine),
		MemberPriceModel:                      models.NewMemberPrice(engine),
		ProductAttributeModel:                 models.NewProductAttribute(engine),
		ProductModel:                          models.NewProduct(engine),
		ProductAttributeValueModel:            models.NewProductAttributeValue(engine),
		ProductCategoryAttributeRelationModel: models.NewProductCategoryAttributeRelation(engine),
		ProductFullReductionModel:             models.NewProductFullReduction(engine),
		ProductLadderModel:                    models.NewProductLadder(engine),
		ProductRecommendModel:                 models.NewProductRecommend(engine),
		SkuStockModel:                         models.NewSkuStock(engine),
	}
}
