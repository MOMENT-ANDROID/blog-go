package admin

import (
	"blog/controllers"
	"blog/service/admin"
	"fmt"
	"strconv"
	"blog/model"
	"blog/fox/url"
	"blog/service/conf"
)

type AdminRole struct {
	Base
}

func (c *AdminRole) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
}
//列表
// @router /admin_role [get]
func (c *AdminRole)List() {
	ser := admin.NewTypeService()
	data, err := ser.Query(conf.ADMIN_ROLE)
	fmt.Println(err)
	c.Data["data"] = data
	c.Data["title"] = "角色-列表"
	c.Data["HtmlHead"] = controllers.ExecuteTemplateHtml("admin/type/head.html", c.Data)
	c.TplName = "admin/admin_role/list.html"
}
//添加
// @router /admin_role/add [get]
func (c *AdminRole)Add() {
	mod := model.NewType()
	mod.TypeId = conf.ADMIN_ROLE
	mod.IsDefault = 0
	mod.IsDel = 0
	mod.IsSystem = 0
	mod.IsShow = 1
	mod.IsChild = 0
	c.Data["type_id"] = conf.ADMIN_ROLE
	c.Data["type_id_name"] = "无"
	c.Data["parent_id_name"] = "无"
	c.Data["info"] = mod
	c.Data["_method"] = "post"
	c.Data["title"] = "角色-添加"
	c.TplName = "admin/admin_role/get.html"
}
//保存
// @router /admin_role [post]
func (c *AdminRole)Post() {
	mod := model.NewType()
	//参数传递
	if err := url.ParseForm(c.Input(), mod); err != nil {
		c.Error(err.Error())
		return
	}
	mod.TypeId = conf.ADMIN_ROLE
	//创建
	ser := admin.NewTypeService()
	id, err := ser.Create(mod)
	if err != nil {
		c.Error(err.Error())
		return
	} else {
		fmt.Println("创建成功！:", id)
		c.Success("操作成功")
	}
}
//编辑
// @router /admin_role/:id [get]
func (c *AdminRole)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	ser := admin.NewTypeService()
	data, err := ser.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Data["info"] = data["info"]
		c.Data["type_id_name"] = data["type_id_name"]
		c.Data["parent_id_name"] = data["parent_id_name"]
		c.Data["_method"] = "put"
		c.Data["is_put"] = true
		c.Data["title"] = "角色-修改"
		c.TplName = "admin/admin_role/get.html"
	}
}
//更新
// @router /admin_role/:id [put]
func (c *AdminRole)Put() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	mod := model.NewType()
	if err := url.ParseForm(c.Input(), mod); err != nil {
		c.Error(err.Error())
		return
	}
	mod.TypeId = conf.ADMIN_ROLE
	//更新
	ser := admin.NewTypeService()
	_, err := ser.Update(int_id, mod)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
//检测名称重复
// @router /admin_role/check_name [post]
func (c *AdminRole)CheckName() {
	//ID 获取 格式化
	id, _ := c.GetInt("id")
	name := c.GetString("name")
	//创建
	ser := admin.NewTypeService()
	ok, err := ser.CheckNameTypeId(conf.ADMIN_ROLE, name, id)
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		c.Success("操作成功")
	}
}
//删除
// @router /admin_role/:id [delete]
func (c *AdminRole)Delete() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//更新
	ser := admin.NewTypeService()
	_, err := ser.DeleteAndTypeId(int_id,conf.ADMIN_ROLE)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
