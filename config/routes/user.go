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

	"github.com/daodao97/egin-skeleton/controller"
)

func RegUserRouter(r *gin.Engine) {

	r.Handle("GET", "/user", func() func(ctx *gin.Context) {
		return func(ctx *gin.Context) {
			var params controller.ParamsValidate
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

			result, code, err := controller.User{}.Get(ctx, params)
			egin.Response(ctx, result, code, err)
		}
	}(), middleware.IpLimiter())

	r.Handle("POST", "/user", func(ctx *gin.Context) {
		result, code, err := controller.User{}.Post(ctx)
		egin.Response(ctx, result, code, err)
	})

	r.Handle("PUT", "/user", func(ctx *gin.Context) {
		result, code, err := controller.User{}.Put(ctx)
		egin.Response(ctx, result, code, err)
	})

	r.Handle("DELETE", "/user", func(ctx *gin.Context) {
		result, code, err := controller.User{}.Delete(ctx)
		egin.Response(ctx, result, code, err)
	})

}

func RegUserCustomValidateFunc() {
	utils.RegCustomValidateFuncs([]utils.CustomValidateFunc{
		controller.Bookabledate,
	})
}
