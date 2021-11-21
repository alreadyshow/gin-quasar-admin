package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysApi = new(sysApi)

type sysApi struct{}

var sysApiData = []system.SysApi{
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "获取用户列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "编辑用户信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, Remark: "新增用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, Remark: "删除用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5, Remark: "根据ID查找用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 6, Remark: "获取用户的菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-menu", ApiMethod: "GET",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 7, Remark: "获取用户的角色列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-role", ApiMethod: "GET",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 8, Remark: "获取角色列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 9, Remark: "编辑角色信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 10, Remark: "新增角色", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 11, Remark: "删除角色", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 12, Remark: "根据ID查找角色", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 13, Remark: "获取角色菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-menu", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 14, Remark: "编辑角色菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-menu-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 15, Remark: "获取角色API", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-api", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 16, Remark: "编辑角色Api", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-api-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 16, Remark: "根据角色查找用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-user", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 16, Remark: "从角色中移除某个用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-user-remove", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 16, Remark: "添加用户到某个角色", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-user-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 17, Remark: "获取菜单列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "menu", ApiPath: "/menu/menu-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 18, Remark: "编辑菜单信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "menu", ApiPath: "/menu/menu-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 19, Remark: "新增菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "menu", ApiPath: "/menu/menu-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 20, Remark: "删除菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "menu", ApiPath: "/menu/menu-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 21, Remark: "根据ID查找菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "menu", ApiPath: "/menu/menu-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 22, Remark: "获取部门列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 23, Remark: "编辑部门信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 24, Remark: "新增部门", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 25, Remark: "删除部门", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 26, Remark: "根据ID查找部门", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 27, Remark: "获取根字典列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dict", ApiPath: "/dict/dict-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 28, Remark: "编辑字典信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dict", ApiPath: "/dict/dict-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 29, Remark: "新增字典", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dict", ApiPath: "/dict/dict-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 30, Remark: "删除字典", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dict", ApiPath: "/dict/dict-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 31, Remark: "根据ID查找字典", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dict", ApiPath: "/dict/dict-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 32, Remark: "获取api列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "api", ApiPath: "/api/api-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 33, Remark: "编辑api信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "api", ApiPath: "/api/api-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 34, Remark: "新增api", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "api", ApiPath: "/api/api-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 35, Remark: "删除api", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "api", ApiPath: "/api/api-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 36, Remark: "上传头像", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "upload", ApiPath: "/upload/avatar", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 37, Remark: "上传文件", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "upload", ApiPath: "/upload/file", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 38, Remark: "上传网站Logo", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "upload", ApiPath: "/upload/web-logo", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 39, Remark: "获取后台配置列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-backend", ApiPath: "/config-backend/config-backend-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 40, Remark: "编辑后台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-backend", ApiPath: "/config-backend/config-backend-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 41, Remark: "新增后台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-backend", ApiPath: "/config-backend/config-backend-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 42, Remark: "删除后台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-backend", ApiPath: "/config-backend/config-backend-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 43, Remark: "获取前台配置列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-frontend", ApiPath: "/config-frontend/config-frontend-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 44, Remark: "编辑前台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-frontend", ApiPath: "/config-frontend/config-frontend-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 45, Remark: "新增前台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-frontend", ApiPath: "/config-frontend/config-frontend-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 46, Remark: "删除前台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-frontend", ApiPath: "/config-frontend/config-frontend-delete", ApiMethod: "DELETE",
	},
}

func (s *sysApi) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysApi{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_api 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[Gin-Quasar-Admin] --> sys_api 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_api 表初始数据成功！")
		global.GqaLog.Error("[Gin-Quasar-Admin] --> sys_api 表初始数据成功！")
		return nil
	})
}
