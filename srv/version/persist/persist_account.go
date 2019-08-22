package persist

import (
	"github.com/gopperin/sme-ray/srv/version/bean"
)

// GetVersion GetVersion Persist
func (maria *Mariadb) GetVersion(_type string) (bean.VersionBase, error) {
	var _obj bean.VersionBase
	err := maria.db.Table("versions").Where("type = ?", _type).Last(&_obj).Error
	return _obj, err
}
