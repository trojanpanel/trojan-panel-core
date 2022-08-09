package dao

import (
	"errors"
	"fmt"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/core"
	"trojan-panel-core/module"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/util"
)

var mySQLConfig = core.Config.MySQLConfig

// SelectUsersPassword 查询所有用户 用于api全量更新用户
func SelectUsersPassword(isAdd bool) ([]vo.ApiUserVo, error) {
	var (
		users       []module.User
		buildSelect string
		values      []interface{}
		err         error
	)

	if isAdd {
		buildSelect, values, err = builder.NamedQuery(fmt.Sprintf("select `password`,download,upload from {{ table_name }} where `download` + `upload` < `quota` or `quota` < 0"),
			map[string]interface{}{
				"table_name": mySQLConfig.Table,
			})
	} else {
		buildSelect, values, err = builder.NamedQuery(fmt.Sprintf("select `password`,download,upload from {{ table_name }} where `download` + `upload` >= `quota` and `quota` >= 0"),
			map[string]interface{}{
				"table_name": mySQLConfig.Table,
			})
	}

	if err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	rows, err := db.Query(buildSelect, values...)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	defer rows.Close()

	if err := scanner.Scan(rows, &users); err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	var apiUserVos = make([]vo.ApiUserVo, 0)
	for _, user := range users {
		apiUserVos = append(apiUserVos, vo.ApiUserVo{
			Password: *user.Password,
			Download: *user.Download,
			Upload:   *user.Upload,
		})
	}
	return apiUserVos, nil
}

func UpdateUser(password string, download *int, upload *int, quota *int) error {
	passwordEncode, err := util.AesEncode(password)
	if err != nil {
		return err
	}
	where := map[string]interface{}{"password": passwordEncode}
	update := map[string]interface{}{}
	if download != nil {
		update["download"] = *download
	}
	if upload != nil {
		update["upload"] = *upload
	}
	if quota != nil {
		update["quota"] = *quota
	}

	if len(update) > 0 {
		buildUpdate, values, err := builder.BuildUpdate(mySQLConfig.Table, where, update)
		if err != nil {
			logrus.Errorln(err.Error())
			return errors.New(constant.SysError)
		}

		_, err = db.Exec(buildUpdate, values...)
		if err != nil {
			logrus.Errorln(err.Error())
			return errors.New(constant.SysError)
		}
	}
	return nil
}
