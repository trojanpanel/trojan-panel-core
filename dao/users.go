package dao

import (
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/module"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/util"
)

func CountUserByApiPort(apiPort int) (int, error) {
	var total int
	selectFields := []string{"count(1)"}
	where := map[string]interface{}{"api_port": apiPort}
	buildSelect, values, err := builder.BuildSelect(mySQLConfig.UsersTable, where, selectFields)
	if err != nil {
		logrus.Errorln(err.Error())
		return 0, errors.New(constant.SysError)
	}
	if err = db.QueryRow(buildSelect, values...).Scan(&total); err != nil {
		logrus.Errorln(err.Error())
		return 0, errors.New(constant.SysError)
	}
	return total, nil
}

func UpdateUser(apiPort *int, password *string, download *int, upload *int) error {
	where := map[string]interface{}{}
	if apiPort != nil {
		where["api_port"] = apiPort
	}
	if password != nil {
		passwordEncode, err := util.AesEncode(*password)
		if err != nil {
			return err
		}
		where["password"] = passwordEncode
	}
	update := map[string]interface{}{}
	if download != nil {
		update["download"] = *download
	}
	if upload != nil {
		update["upload"] = *upload
	}

	if len(update) > 0 {
		buildUpdate, values, err := builder.BuildUpdate(mySQLConfig.UsersTable, where, update)
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

func InsertUsers(users []module.Users) error {
	var data []map[string]interface{}
	for _, item := range users {
		encodePassword, err := util.AesEncode(*item.Password)
		if err != nil {
			return err
		}
		user := map[string]interface{}{
			"account_id": item.AccountId,
			"api_port":   item.ApiPort,
			"password":   encodePassword,
			"download":   item.Download,
			"upload":     item.Upload,
		}
		data = append(data, user)
	}
	buildInsert, values, err := builder.BuildInsert(mySQLConfig.AccountTable, data)
	if err != nil {
		logrus.Errorln(err.Error())
		return errors.New(constant.SysError)
	}
	if _, err := db.Exec(buildInsert, values...); err != nil {
		logrus.Errorln(err.Error())
		return errors.New(constant.SysError)
	}
	return nil
}

func SelectUsersToApi(isAdd bool) ([]vo.UserApiVo, error) {
	var (
		users       []module.Users
		buildSelect string
		values      []interface{}
		err         error
	)
	data := map[string]interface{}{
		"account_table": mySQLConfig.AccountTable,
		"users_table":   mySQLConfig.UsersTable,
	}
	if isAdd {
		buildSelect, values, err = builder.NamedQuery(`select u.api_port, u.password, u.download, u.upload
from {{ account_table }} a
         left join {{ users_table }} u on a.id = u.account_id
where a.download + a.upload < a.quota
   or a.quota < 0`,
			data)
	} else {
		buildSelect, values, err = builder.NamedQuery(`select u.api_port, u.password, u.download, u.upload
from {{ account_table }} a
         left join {{ users_table }} u on a.id = u.account_id
where a.download + a.upload >= a.quota
  and a.quota >= 0`,
			data)
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
	var apiUserVo = make([]vo.UserApiVo, 0)
	for _, user := range users {
		passwordDecode, err := util.AesDecode(*user.Password)
		if err != nil {
			return nil, err
		}
		apiUserVo = append(apiUserVo, vo.UserApiVo{
			Password: passwordDecode,
			Download: *user.Download,
			Upload:   *user.Upload,
		})
	}
	return apiUserVo, nil
}

func SelectUsersDUByAccountId(accountId int) (int, int, error) {
	var (
		download int
		upload   int
	)
	selectFields := []string{"sum(u.download) download,sum(u.upload) upload"}
	where := map[string]interface{}{"account_id": accountId}
	buildSelect, values, err := builder.BuildSelect(mySQLConfig.UsersTable, where, selectFields)
	if err != nil {
		logrus.Errorln(err.Error())
		return 0, 0, errors.New(constant.SysError)
	}
	if err = db.QueryRow(buildSelect, values...).Scan(&download, &upload); err != nil {
		logrus.Errorln(err.Error())
		return 0, 0, errors.New(constant.SysError)
	}
	return download, upload, nil
}
