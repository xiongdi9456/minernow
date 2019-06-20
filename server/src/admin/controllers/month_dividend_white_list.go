package controllers

import (
	"admin/controllers/errcode"
	"common"
	"encoding/json"
	"fmt"
	"utils/admin/dao"
	agentdao "utils/agent/dao"
)

type MonthDividendWhiteListController struct {
	BaseController
}

func (c *MonthDividendWhiteListController) GetMonthDividendWhiteList() {
	//_, errCode := c.CheckPermission()
	//if errCode != controllers.ERROR_CODE_SUCCESS {
	//	c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteList, errCode, string(c.Ctx.Input.RequestBody))
	//	return
	//}

	var input string
	id, err := c.GetInt(KEY_ID, -1)
	if err != nil {
		c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteList, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
		return
	}
	if id > 0 {
		input = fmt.Sprintf("{\"id\":%v}", id)
		data, err := dao.MonthDividendWhiteListDaoEntity.QueryById(uint32(id))
		if err != nil {
			c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteList, controllers.ERROR_CODE_DB, input)
			return
		}
		c.SuccessResponseAndLog(OPActionReadMonthDividendWhiteList, input, *data)
	} else {
		page, err := c.GetInt("page", 1)
		if err != nil {
			c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteList, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
			return
		}
		perPage, err := c.GetInt("per_page", DEFAULT_PER_PAGE)
		if err != nil {
			c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteList, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
			return
		}

		//分页查询
		res := map[string]interface{}{}
		input = fmt.Sprintf("{\"page\":%v,\"per_page\":%v}", page, perPage)
		count, data, err := dao.MonthDividendWhiteListDaoEntity.QueryByPage(page, perPage)
		if err != nil {
			c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteList, controllers.ERROR_CODE_DB, input)
			return
		}
		meta := dao.PageInfo{
			Limit: perPage,
			Total: int(count),
			Page:  page,
		}
		res["meta"] = meta
		res["list"] = data
		c.SuccessResponseAndLog(OPActionReadMonthDividendWhiteList, input, res)
	}
}

func (c *MonthDividendWhiteListController) CreateMonthDividendWhiteList() {
	//_, errCode := c.CheckPermission()
	//if errCode != controllers.ERROR_CODE_SUCCESS {
	//	c.ErrorResponseAndLog(OPActionAddMonthDividendWhiteList, errCode, string(c.Ctx.Input.RequestBody))
	//	return
	//}

	type msg struct {
		Name          string `json:"name"`
		DividendRatio int32  `json:"dividend_ratio"`
	}
	req := &msg{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		common.LogFuncDebug("json decode: %s \nfailed : %v", string(c.Ctx.Input.RequestBody), err)
		c.ErrorResponseAndLog(OPActionAddMonthDividendWhiteList, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
		return
	}

	app, err := dao.MonthDividendWhiteListDaoEntity.Create(req.DividendRatio, req.Name)
	if err != nil {
		c.ErrorResponseAndLog(OPActionAddMonthDividendWhiteList, controllers.ERROR_CODE_DB, string(c.Ctx.Input.RequestBody))
		return
	}

	//返回查询结果
	c.SuccessResponseAndLog(OPActionAddMonthDividendWhiteList, string(c.Ctx.Input.RequestBody), app)
}

func (c *MonthDividendWhiteListController) UpdateMonthDividendWhiteList() {
	//_, errCode := c.CheckPermission()
	//if errCode != controllers.ERROR_CODE_SUCCESS {
	//	c.ErrorResponseAndLog(OPActionEditMonthDividendWhiteList, errCode, string(c.Ctx.Input.RequestBody))
	//	return
	//}

	type msg struct {
		Id            uint32 `json:"id"`
		Name          string `json:"name"`
		DividendRatio int32  `json:"dividend_ratio"`
	}
	req := &msg{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		common.LogFuncDebug("json decode: %s \nfailed : %v", string(c.Ctx.Input.RequestBody), err)
		c.ErrorResponseAndLog(OPActionEditMonthDividendWhiteList, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
		return
	}

	app, err := dao.MonthDividendWhiteListDaoEntity.UpdateById(req.Id, req.DividendRatio, req.Name)
	if err != nil {
		c.ErrorResponseAndLog(OPActionEditMonthDividendWhiteList, controllers.ERROR_CODE_DB, string(c.Ctx.Input.RequestBody))
		return
	}

	//返回查询结果
	c.SuccessResponseAndLog(OPActionEditMonthDividendWhiteList, string(c.Ctx.Input.RequestBody), app)
}

func (c *MonthDividendWhiteListController) DelMonthDividendWhiteList() {
	//_, errCode := c.CheckPermission()
	//if errCode != controllers.ERROR_CODE_SUCCESS {
	//	c.ErrorResponseAndLog(OPActionDelMonthDividendWhiteList, errCode, string(c.Ctx.Input.RequestBody))
	//	return
	//}

	id, err := c.GetUint32(KEY_ID, 0)
	if err != nil {
		c.ErrorResponseAndLog(OPActionDelMonthDividendWhiteList, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
		return
	}

	input := fmt.Sprintf("{\"id\":%v}", id)
	err = dao.MonthDividendWhiteListDaoEntity.DelById(id)
	if err != nil {
		c.ErrorResponseAndLog(OPActionDelMonthDividendWhiteList, controllers.ERROR_CODE_DB, input)
		return
	}
	err = agentdao.AgentPathDaoEntity.DelDividendWhiteList(id)
	if err != nil {
		common.LogFuncDebug("AgentPathDaoEntity.DelWhiteList fail error:%v", err)
	}

	//返回查询结果
	ack := map[string]interface{}{
		"id": id,
	}
	c.SuccessResponseAndLog(OPActionDelMonthDividendWhiteList, input, ack)
}

func (c *MonthDividendWhiteListController) GetAgent() {
	//_, errCode := c.CheckPermission()
	//if errCode != controllers.ERROR_CODE_SUCCESS {
	//	c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteListAgent, errCode, string(c.Ctx.Input.RequestBody))
	//	return
	//}
	errCode := controllers.ERROR_CODE_SUCCESS
	pUid, err := c.GetUint64("puid", 0)
	if err != nil {
		c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteListAgent, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
		return
	}
	wId, err := c.GetInt("month_dividend_whitelist_id", 0)
	if err != nil {
		c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteListAgent, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
		return
	}
	inviteCode := c.GetString("invite_code")
	mobile := c.GetString(KEY_MOBILE)

	if pUid == 0 && len(mobile) > 0 {
		pUid, errCode = GetOtcUidByMobile(mobile)
		if errCode != controllers.ERROR_CODE_SUCCESS {
			c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteListAgent, errCode, string(c.Ctx.Input.RequestBody))
			return
		}
	}

	page, err := c.GetInt(KEY_PAGE, 1)
	if err != nil {
		c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteListAgent, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
		return
	}
	perPage, err := c.GetInt(KEY_LIMIT, DEFAULT_PER_PAGE)
	if err != nil {
		c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteListAgent, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
		return
	}
	input := fmt.Sprintf("{\"mobile\":\"%v\",\"whitelist_id\":%v,\"inviteCode\":\"%v\",\"page\":%v,\"per_page\":%v}", mobile, wId, inviteCode, page, perPage)

	count, data, err := agentdao.AgentPathDaoEntity.GetAgentPathByPage(page, perPage, wId, pUid, inviteCode)
	if err != nil {
		c.ErrorResponseAndLog(OPActionReadMonthDividendWhiteListAgent, controllers.ERROR_CODE_DB, input)
		return
	}

	meta := dao.PageInfo{
		Limit: perPage,
		Total: int(count),
		Page:  page,
	}

	res := map[string]interface{}{}
	res["meta"] = meta
	res["list"] = ClientAgentInfos(data)
	c.SuccessResponseAndLog(OPActionReadMonthDividendWhiteListAgent, input, res)
}

//easyjson:json
type AddDividendAgentMsg struct {
	Mobile string `json:"mobile"`
	WId    uint32 `json:"whitelist_id"`
}

func (c *MonthDividendWhiteListController) AddAgent() {
	//_, errCode := c.CheckPermission()
	//if errCode != controllers.ERROR_CODE_SUCCESS {
	//	c.ErrorResponseAndLog(OPActionEditMonthDividendWhiteListAgent, errCode, string(c.Ctx.Input.RequestBody))
	//	return
	//}

	req := &AddDividendAgentMsg{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		common.LogFuncDebug("json decode: %s \nfailed : %v", string(c.Ctx.Input.RequestBody), err)
		c.ErrorResponseAndLog(OPActionEditMonthDividendWhiteListAgent, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
		return
	}

	otcUid, errCode := GetOtcUidByMobile(req.Mobile)
	if errCode != controllers.ERROR_CODE_SUCCESS {
		c.ErrorResponseAndLog(OPActionEditMonthDividendWhiteListAgent, errCode, string(c.Ctx.Input.RequestBody))
		return
	}

	err = agentdao.AgentPathDaoEntity.SetAgentDividendWhiteList(otcUid, req.WId)
	if err != nil {
		if dao.ErrParam == err {
			c.ErrorResponseAndLog(OPActionEditMonthDividendWhiteListAgent, controllers.ERROR_CODE_EXIST_WHITE_LIST_AGENT, string(c.Ctx.Input.RequestBody))
		} else {
			c.ErrorResponseAndLog(OPActionEditMonthDividendWhiteListAgent, controllers.ERROR_CODE_DB, string(c.Ctx.Input.RequestBody))
		}
		return
	}

	//返回查询结果
	data := map[string]interface{}{
		"mobile":       req.Mobile,
		"whitelist_id": req.WId,
	}
	c.SuccessResponseAndLog(OPActionEditMonthDividendWhiteListAgent, string(c.Ctx.Input.RequestBody), data)
}

func (c *MonthDividendWhiteListController) DelAgent() {
	//_, errCode := c.CheckPermission()
	//if errCode != controllers.ERROR_CODE_SUCCESS {
	//	c.ErrorResponseAndLog(OPActionDelMonthDividendWhiteListAgent, errCode, string(c.Ctx.Input.RequestBody))
	//	return
	//}

	otcUid, err := c.GetUint64(KEY_UID, 0)
	if err != nil || otcUid <= 0 {
		c.ErrorResponseAndLog(OPActionDelMonthDividendWhiteListAgent, controllers.ERROR_CODE_PARAMS_ERROR, string(c.Ctx.Input.RequestBody))
		return
	}
	input := fmt.Sprintf("{\"uid\":%v}", otcUid)

	err = agentdao.AgentPathDaoEntity.SetAgentDividendWhiteList(otcUid, 0)
	if err != nil {
		c.ErrorResponseAndLog(OPActionDelMonthDividendWhiteListAgent, controllers.ERROR_CODE_DB, input)
		return
	}

	//返回查询结果
	data := map[string]interface{}{
		"uid": fmt.Sprintf("%v", otcUid),
	}
	c.SuccessResponseAndLog(OPActionDelMonthDividendWhiteListAgent, input, data)
}
