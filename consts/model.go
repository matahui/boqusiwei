package consts

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

// 自定义的时间格式为 "2006-01-02 15:04:05"
const customTimeFormat = "2006-01-02 15:04:05"

// 实现自定义的 JSON 序列化方法
func (c CustomTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", c.Format(customTimeFormat))
	return []byte(formatted), nil
}

// 实现 Scanner 接口
func (c *CustomTime) Scan(value interface{}) error {
	if v, ok := value.(time.Time); ok {
		*c = CustomTime{Time: v}
		return nil
	}
	return fmt.Errorf("cannot scan value %v into CustomTime", value)
}

// 实现 Valuer 接口
func (c CustomTime) Value() (driver.Value, error) {
	return c.Time, nil
}