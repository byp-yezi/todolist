package conf

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Dictinary 字典
var Dictinary map[string]interface{}

// LoadLocales 读取国际化文件
func LoadLocales(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	m := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return err
	}
	Dictinary = m
	return nil
}

// T 翻译
func T(key string) string {
	dic := Dictinary
	keys := strings.Split(key, ".")
	for index, path := range keys {
		// 如果到达了最后一层，寻找目标翻译
		if len(keys) == (index + 1) {
			// 查找对应的键
			if value, ok := dic[path].(string); ok {
				return value
			}
			return fmt.Sprintf("Translation not found for '%s'", path) // 返回翻译未找到
		}
		// 如果还有下一层，继续寻找
		if nextDic, ok := dic[path].(map[string]interface{}); ok {
			dic = nextDic // 继续在子字典中查找
		} else {
			return fmt.Sprintf("Invalid structure for key '%s'", key) // 返回结构错误
		}
	}
	return ""
}
