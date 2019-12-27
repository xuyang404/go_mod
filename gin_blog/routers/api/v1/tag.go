package v1

import (
	"GinHello/models"
	"GinHello/pkg/e"
	"GinHello/pkg/setting"
	"GinHello/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//获取多个标签
func GetTags(ctx *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name := ctx.Query("name"); name != "" {
		maps["name"] = name
	}

	if arg := ctx.Query("state"); arg != "" {
		maps["state"] = com.StrTo(arg).MustInt()
	}

	data["lists"] = models.GetTags(util.GetPage(ctx), setting.PageSize, maps)
	data["count"] = models.GetTagTotal(maps)

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
}

//新增文章标签
func AddTag(ctx *gin.Context) {
	name := ctx.Query("name")
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()
	createdBy := ctx.Query("created_by")

	vaild := validation.Validation{}
	vaild.Required(name, "name").Message("名称不能为空")
	vaild.MaxSize(name, 100, "name").Message("名称最长为100字符")
	vaild.Required(createdBy, "createdBy").Message("创建人不能为空")
	vaild.MaxSize(createdBy, 100, "createdBy").Message("创建人最长为100字符")
	vaild.Range(state, 0, 1, "state").Message("状态只允许为0或1")

	code := e.INVALID_PARAMS
	if !vaild.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg" : e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

//修改文章标签
func EditTag(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()
	modified_by := ctx.Query("modified_by")
	name := ctx.Query("name")

	vaild := validation.Validation{}
	vaild.Required(id, "id").Message("id不能为空")
	vaild.Required(modified_by, "modified_by").Message("修改人不能为空")
	vaild.MaxSize(modified_by, 100, "modified_by").Message("修改人最长为100字符")
	vaild.MaxSize(name, 100, "name").Message("名称最长为100字符")
	vaild.Range(state, 0, 1, "state").Message("状态只允许为0或1")

	code := e.INVALID_PARAMS
	if !vaild.HasErrors() {
		if models.ExistTagById(id) {
			code = e.SUCCESS
			data := make(map[string]interface{})
			data["modified_by"] = modified_by
			data["name"] = name
			data["state"] = state

			models.Edit(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg" : e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

//删除文章标签
func DeleteTag(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	vaild := validation.Validation{}
	vaild.Min(id, 1, "id").Message("id必须大于0")

	code := e.INVALID_PARAMS
	if !vaild.HasErrors() {
		if models.ExistTagById(id) {
			code = e.SUCCESS
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg" : e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
