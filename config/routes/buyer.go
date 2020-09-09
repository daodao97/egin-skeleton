// ****************************
// 该文件为系统生成, 请勿随意更改
// ****************************
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/daodao97/egin"
	"github.com/daodao97/egin/consts"
	"github.com/daodao97/egin/middleware"
	"github.com/daodao97/egin/utils"

	"skeleton/controller"
)

func RegBuyerRouter(r *gin.Engine) {

	r.Handle("GET", "/buyer", func() func(ctx *gin.Context) {
		return func(ctx *gin.Context) {
			var params controller.Params2Validate
			err := ctx.ShouldBind(&params)
			if err != nil {
				errs, _ := utils.TransErr(params, err.(validator.ValidationErrors))
				ctx.JSON(http.StatusOK, gin.H{
					"code":    consts.ErrorParam,
					"message": errs,
				})
				ctx.Abort()
				return
			}

			result, code, err := controller.Buyer{}.Get(ctx, params)
			egin.Response(ctx, result, code, err)
		}
	}(), middleware.IpLimiter())

	r.Handle("POST", "/buyer", func(ctx *gin.Context) {
		result, code, err := controller.Buyer{}.Post(ctx)
		egin.Response(ctx, result, code, err)
	})

	r.Handle("PUT", "/buyer", func(ctx *gin.Context) {
		result, code, err := controller.Buyer{}.Put(ctx)
		egin.Response(ctx, result, code, err)
	})

	r.Handle("DELETE", "/buyer", func(ctx *gin.Context) {
		result, code, err := controller.Buyer{}.Delete(ctx)
		egin.Response(ctx, result, code, err)
	})

}
