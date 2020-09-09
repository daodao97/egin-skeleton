package controller

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/daodao97/egin/cache"
	"github.com/daodao97/egin/consts"
	"github.com/daodao97/egin/db"

	"skeleton/model"
)

// @Controller BUYER管理 这里是简介
type Buyer struct {
	BaseApi
}

type Params2Validate struct {
	CheckIn  time.Time `form:"check_in" time_format:"2006-01-02" json:"check_in" binding:"required" label:"输入时间" in:"query"`
	CheckOut time.Time `form:"check_out" json:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02" label:"输出时间" in:"query"`
}

// @GetApi /buyer 若无默认为全小写的方法名
// @Summary BUYER列表接口
// @Desc 接口简介, 若无则为空 维护者: 刀刀
// @Params Params2Validate 接口参数所对应的结构体
// @Response
// @Middleware IpLimiter
func (u Buyer) Get(c *gin.Context, params Params2Validate) (interface{}, consts.ErrCode, error) {

	var result []model.UserEntity
	user := model.User
	err := user.Get(db.Filter{
		"id": map[string]int{
			">": 20,
		},
	}, db.Attr{
		Select:  []string{"realname", "id", "username", "password"},
		OrderBy: "id desc",
	}, &result)

	redis := cache.Redis{Connection: "default"}
	setV, _ := json.Marshal([]int{1, 2, 4})
	err = redis.Set("egin:test", setV, 0)
	_cache, err := redis.Get("egin:test")

	return []interface{}{result, params, _cache}, 0, err
}

// @PostApi /buyer 若无默认为全小写的方法名
// @Summary 创建BUYER
// @Desc 这个接口支持你创建一个BUYER 维护者: 刀刀
// @Response
func (u Buyer) Post(c *gin.Context) (interface{}, consts.ErrCode, error) {
	user := model.User
	result, _, err := user.Insert(db.Record{
		"username": "test33333",
		"realname": "你好",
		"password": "cool",
	})
	var code consts.ErrCode
	if err != nil {
		code = consts.ErrorSystem
	}
	return []interface{}{result}, code, err
}

// @PutApi /buyer
// @Summary 更新BUYER信息
// @Desc 维护者: 刀刀
// @Response
func (u Buyer) Put(c *gin.Context) (interface{}, consts.ErrCode, error) {
	user := model.User
	_, affected, err := user.Update(
		db.Filter{
			"id": 13,
		},
		db.Record{
			"username": "test12",
		})
	var code consts.ErrCode
	if err != nil {
		code = consts.ErrorSystem
	}
	return affected, code, err
}

// @DeleteApi /buyer
// @Summary 删除BUYER
// @Desc 维护者: 刀刀
// @Response
func (u Buyer) Delete(c *gin.Context) (interface{}, consts.ErrCode, error) {
	user := model.User
	_, affected, err := user.Delete(db.Filter{
		"id": 22,
	})
	return affected, 0, err
}
