package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

type ServiceDept struct {
}

func (s *ServiceDept) GetDeptList(requestDeptList system.RequestDeptList) (err error, role interface{}, total int64) {
	pageSize := requestDeptList.PageSize
	offset := requestDeptList.PageSize * (requestDeptList.Page - 1)
	db := global.GqaDb.Model(&system.SysDept{})
	var deptList []system.SysDept
	//配置搜索
	if requestDeptList.DeptCode != ""{
		db = db.Where("dept_code like ?", "%" + requestDeptList.DeptCode + "%")
	}
	if requestDeptList.DeptName != ""{
		db = db.Where("dept_name like ?", "%" + requestDeptList.DeptName + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(requestDeptList.SortBy, requestDeptList.Desc)).Find(&deptList).Error
	return err, deptList, total
}

func (s *ServiceDept) EditDept(toEditDept system.SysDept) (err error) {
	var sysDept system.SysDept
	if err = global.GqaDb.Where("id = ?", toEditDept.Id).First(&sysDept).Error; err != nil {
		return err
	}
	if sysDept.Stable == "yes" {
		return errors.New("系统内置不允许编辑：" + toEditDept.DeptCode)
	}
	err = global.GqaDb.Updates(&toEditDept).Error
	return err
}

func (s *ServiceDept) AddDept(toAddDept system.SysDept) (err error) {
	var dept system.SysDept
	if !errors.Is(global.GqaDb.Where("dept_code = ?", toAddDept.DeptCode).First(&dept).Error, gorm.ErrRecordNotFound) {
		return errors.New("此部门已存在：" + toAddDept.DeptCode)
	}
	err = global.GqaDb.Create(&toAddDept).Error
	return err
}

func (s *ServiceDept) DeleteDept(id uint) (err error) {
	var sysDept system.SysDept
	if err = global.GqaDb.Where("id = ?", id).First(&sysDept).Error; err != nil {
		return err
	}
	if sysDept.Stable == "yes" {
		return errors.New("系统内置不允许删除：" + sysDept.DeptCode)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysDept).Error
	return err
}

func (s *ServiceDept) QueryDeptById(id uint) (err error, deptInfo system.SysDept) {
	var dept system.SysDept
	err = global.GqaDb.Preload("Owner").Preload("CreatedByUser").Preload("UpdatedByUser").First(&dept, "id = ?", id).Error
	return err, dept
}
