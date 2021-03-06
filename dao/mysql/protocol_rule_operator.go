package mysql

import (
	"github.com/LotteWong/giotto-gateway-admin/constants"
	"github.com/LotteWong/giotto-gateway-admin/models/po"
	"github.com/LotteWong/giotto-gateway-admin/utils"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ProtocolRuleOperator struct{}

func NewProtocolRuleOperator() *ProtocolRuleOperator {
	return &ProtocolRuleOperator{}
}

func (o *ProtocolRuleOperator) FindHttpRule(ctx *gin.Context, tx *gorm.DB, req *po.HttpRule) (*po.HttpRule, error) {
	res := &po.HttpRule{}
	err := tx.SetCtx(utils.GetGinTraceContext(ctx)).Where(req).Find(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (o *ProtocolRuleOperator) SaveHttpRule(ctx *gin.Context, tx *gorm.DB, req *po.HttpRule) error {
	if err := tx.SetCtx(utils.GetGinTraceContext(ctx)).Save(req).Error; err != nil {
		return err
	}
	return nil
}

func (o *ProtocolRuleOperator) ListHttpRulesByServiceId(ctx *gin.Context, tx *gorm.DB, serviceId int64) ([]po.HttpRule, int64, error) {
	var httpRules []po.HttpRule
	var count int64
	var err error
	query := tx.SetCtx(utils.GetGinTraceContext(ctx)).Table((&po.HttpRule{}).TableName()).Where("service_id=?", serviceId)

	err = query.Order("id desc").Find(&httpRules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, nil
	}

	err = query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return httpRules, count, nil
}

func (o *ProtocolRuleOperator) GroupByHttpServiceType(ctx *gin.Context, tx *gorm.DB) ([]po.HttpServicePercentage, error) {
	var groups []po.HttpServicePercentage

	query := tx.SetCtx(utils.GetGinTraceContext(ctx)).Table((&po.HttpRule{}).TableName()).Where("is_delete=0")

	// wss protocol
	wssPercent := &po.HttpServicePercentage{
		HttpServiceType:  constants.HttpServiceTypeWss,
		HttpServiceCount: o.countHttpServiceType(query, constants.Enable, constants.Enable),
	}
	groups = append(groups, *wssPercent)

	// ws protocol
	wsPercent := &po.HttpServicePercentage{
		HttpServiceType:  constants.HttpServiceTypeWs,
		HttpServiceCount: o.countHttpServiceType(query, constants.Enable, constants.Disable),
	}
	groups = append(groups, *wsPercent)

	// https protocol
	httpsPercent := &po.HttpServicePercentage{
		HttpServiceType:  constants.HttpServiceTypeHttps,
		HttpServiceCount: o.countHttpServiceType(query, constants.Disable, constants.Enable),
	}
	groups = append(groups, *httpsPercent)

	// http protocol
	httpPercent := &po.HttpServicePercentage{
		HttpServiceType:  constants.HttpServiceTypeHttp,
		HttpServiceCount: o.countHttpServiceType(query, constants.Disable, constants.Disable),
	}
	groups = append(groups, *httpPercent)

	return groups, nil
}

func (o *ProtocolRuleOperator) countHttpServiceType(query *gorm.DB, needWebSocket, needHttps int) int64 {
	var httpServiceTypeCount int64
	httpServiceTypeQuery := query.Where("(need_websocket = ? and need_https = ?)", needWebSocket, needHttps)
	httpServiceTypeQuery.Limit(-1).Offset(-1).Count(&httpServiceTypeCount)
	return httpServiceTypeCount
}

func (o *ProtocolRuleOperator) FindTcpRule(ctx *gin.Context, tx *gorm.DB, req *po.TcpRule) (*po.TcpRule, error) {
	res := &po.TcpRule{}
	err := tx.SetCtx(utils.GetGinTraceContext(ctx)).Where(req).Find(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (o *ProtocolRuleOperator) SaveTcpRule(ctx *gin.Context, tx *gorm.DB, req *po.TcpRule) error {
	if err := tx.SetCtx(utils.GetGinTraceContext(ctx)).Save(req).Error; err != nil {
		return err
	}
	return nil
}

func (o *ProtocolRuleOperator) ListTcpRulesByServiceId(ctx *gin.Context, tx *gorm.DB, serviceId int64) ([]po.TcpRule, int64, error) {
	var tcpRules []po.TcpRule
	var count int64
	var err error
	query := tx.SetCtx(utils.GetGinTraceContext(ctx)).Table((&po.TcpRule{}).TableName()).Where("service_id=?", serviceId)

	err = query.Order("id desc").Find(&tcpRules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	err = query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return tcpRules, count, nil
}

func (o *ProtocolRuleOperator) FindGrpcRule(ctx *gin.Context, tx *gorm.DB, req *po.GrpcRule) (*po.GrpcRule, error) {
	res := &po.GrpcRule{}
	err := tx.SetCtx(utils.GetGinTraceContext(ctx)).Where(req).Find(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (o *ProtocolRuleOperator) SaveGrpcRule(ctx *gin.Context, tx *gorm.DB, req *po.GrpcRule) error {
	if err := tx.SetCtx(utils.GetGinTraceContext(ctx)).Save(req).Error; err != nil {
		return err
	}
	return nil
}

func (o *ProtocolRuleOperator) ListGrpcRulesByServiceId(ctx *gin.Context, tx *gorm.DB, serviceId int64) ([]po.GrpcRule, int64, error) {
	var grpcRules []po.GrpcRule
	var count int64
	var err error
	query := tx.SetCtx(utils.GetGinTraceContext(ctx)).Table((&po.GrpcRule{}).TableName()).Where("service_id=?", serviceId)

	err = query.Order("id desc").Find(&grpcRules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, nil
	}

	err = query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return grpcRules, count, nil
}
