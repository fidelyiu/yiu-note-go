package dao

import (
	"errors"
	yiuStr "github.com/fidelyiu/yiu-go-tool/string"
	yiuStrList "github.com/fidelyiu/yiu-go-tool/string_list"
	yiuLang "github.com/fidelyiu/yiu-go-tool/yiu_lang"
	langCore "github.com/fidelyiu/yiu-note/core/lang"
	"go.etcd.io/bbolt"
)

// FindAllByTableName 根据表明查找所有数据
func FindAllByTableName(db *bbolt.DB, tableName string) ([]string, error) {
	result := make([]string, 0)
	err := db.View(func(tx *bbolt.Tx) error {
		table := GetTableByName(tx, tableName)
		if table == nil {
			return nil
		}
		err := table.ForEach(func(k, v []byte) error {
			result = append(result, string(v))
			return nil
		})
		return err
	})
	return result, err
}

// CountAllByTableName 根据表明统计所有数据
func CountAllByTableName(db *bbolt.DB, tableName string) (int, error) {
	var result int
	err := db.View(func(tx *bbolt.Tx) error {
		table := GetTableByName(tx, tableName)
		err := table.ForEach(func(_, _ []byte) error {
			result++
			return nil
		})
		return err
	})
	return result, err
}

// FindByTableNameAndKey 根据表名&Key查找一个数据
func FindByTableNameAndKey(y yiuLang.YiuLang, db *bbolt.DB, tableName string, key string, entityName string) (string, error) {
	if key == "" {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "查询" + entityName + "报错，key不能为空",
			yiuLang.EnUS: "Query " + entityName + " error, key cannot be empty",
		}
		return "", errors.New(langCore.GetLangByKey(y, lMap))
	}
	var result string
	err := db.View(func(tx *bbolt.Tx) error {
		table := GetTableByName(tx, tableName)
		result = string(table.Get([]byte(key)))
		return nil
	})
	return result, err
}

// IsEffectiveByTableNameAndKey 判断该ID是否有效
func IsEffectiveByTableNameAndKey(y yiuLang.YiuLang, db *bbolt.DB, tableName string, key string, entityName string) bool {
	v, err := FindByTableNameAndKey(y, db, tableName, key, entityName)
	if err != nil || len(v) == 0 {
		return false
	}
	return true
}

// SaveByTableNameAndKey 根据表明保存一条数据，key&数据不能为空
func SaveByTableNameAndKey(y yiuLang.YiuLang, db *bbolt.DB, tableName string, key string, entityByte []byte, entityName string) error {
	if key == "" {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "保存" + entityName + "报错，key不能为空",
			yiuLang.EnUS: "Save " + entityName + " error, key cannot be empty",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	if len(entityByte) == 0 {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "保存" + entityName + "报错，数据不能为空",
			yiuLang.EnUS: "Save " + entityName + " error, Data cannot be empty",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	err := db.Update(func(tx *bbolt.Tx) error {
		table := GetTableByName(tx, tableName)
		err := table.Put([]byte(key), entityByte)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateByTableNameAndKey 根据表名、Key、名称修改一个数据，如果数据不存在报错
func UpdateByTableNameAndKey(y yiuLang.YiuLang, db *bbolt.DB, tableName string, key string, entityByte []byte, entityName string) error {
	dbEntity, err := FindByTableNameAndKey(y, db, tableName, key, entityName)
	if err != nil {
		return err
	}
	if dbEntity == "" {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "修改" + entityName + "报错，数据不存在",
			yiuLang.EnUS: "Update " + entityName + " error, Data does not exist",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	if key == "" {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "修改" + entityName + "报错，key不能为空",
			yiuLang.EnUS: "Update " + entityName + " error, key cannot be empty",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	if len(entityByte) == 0 {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "修改" + entityName + "报错，数据不能为空",
			yiuLang.EnUS: "Update " + entityName + " error, Data cannot be empty",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	return SaveByTableNameAndKey(y, db, tableName, key, entityByte, entityName)
}

// DeleteByTableNameAndKey 根据表名、Key删除一条数据
func DeleteByTableNameAndKey(y yiuLang.YiuLang, db *bbolt.DB, tableName string, key string, entityName string) error {
	if key == "" {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "删除" + entityName + "报错，key不能为空",
			yiuLang.EnUS: "Delete " + entityName + " error, key cannot be empty",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	err := db.Update(func(tx *bbolt.Tx) error {
		table := GetTableByName(tx, tableName)
		err := table.Delete([]byte(key))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetTableByName 根据字符串获取数据库中的表
func GetTableByName(tx *bbolt.Tx, tableName string) *bbolt.Bucket {
	if yiuStr.IsBlank(tableName) {
		return nil
	}
	tableNameList := yiuStr.ToStrListBySep(tableName, ".")
	baseBucket := tx.Bucket([]byte(tableNameList[0]))
	if len(tableNameList) <= 1 {
		return baseBucket
	}
	yiuStrList.OpDeleteByIndex(&tableNameList, 0)
	for _, v := range tableNameList {
		baseBucket = baseBucket.Bucket([]byte(v))
	}
	return baseBucket
}

func GetCustomizeWork(db *bbolt.DB, opFunc func(tx *bbolt.Tx) error) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer func(tx *bbolt.Tx) {
		_ = tx.Rollback()
	}(tx)
	err = opFunc(tx)
	if err != nil {
		return err
	}
	// Commit the transaction and check for error.
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
