package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agreyfox/baas"
)

func (s *Server) handleCreateKey() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(createBaasUserRequest)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		_, err := s.ApplicationService.Application(r.Context(), payload.ApplicationID)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrApplicationNotFound,
				"status": 0,
			})
			return
		}

		if err := payload.validate(); err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  "Password is not secure",
				"status": 0,
			})
			return
		}

		n := &baas.NewBAASUser{
			Name:          payload.Name,
			Email:         payload.Email,
			CipherText:    payload.CipherText,
			Password:      payload.Password,
			ApplicationID: payload.ApplicationID,
			Description:   payload.Description,
		}
		//fmt.Println(s.BlockService)
		a, err := s.BlockService.Create(r.Context(), n)
		if err != nil {
			//RespondError(w, err, http.StatusInternalServerError, GetReqID(r))
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &baasUserResponse{
			Data:   sanitizeBaasUser(a),
			Status: 1,
		}

		RespondCreated(w, resp)
	})
}

func (s *Server) handleChangePassword() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(updateBaasUserRequest)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		_, err := s.ApplicationService.Application(r.Context(), payload.ApplicationID)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrApplicationNotFound,
				"status": 0,
			})
			return
		}
		if err := payload.validate(); err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		//fmt.Println(s.BlockService)
		err = s.BlockService.ChangePassword(r.Context(), payload.Email, payload.OldPassword, payload.NewPassword)
		if err == baas.ErrBaasInvalidPassword {
			RespondOK(w, map[string]interface{}{
				"error":  err,
				"status": 0,
			})
			return
		} else if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasNoSuchUser,
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
		}

		RespondOK(w, resp)
	})
}
func (s *Server) handleGetKey() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(createBaasUserRequest)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}
		_, err := s.ApplicationService.Application(r.Context(), payload.ApplicationID)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrApplicationNotFound,
				"status": 0,
			})
			return
		}

		if err := payload.validate(); err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		if len(payload.CipherText) < 8 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasCipherTextRequired,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetKey(r.Context(), payload.Email, payload.Password, payload.CipherText)
		if err == baas.ErrBaasInvalidPassword {
			RespondOK(w, map[string]interface{}{
				"error":  err,
				"status": 0,
			})
			return
		} else if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err,
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]string{
				"key": a,
			},
		}

		RespondOK(w, resp)
	})
}

func (s *Server) handleDeleteKey() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(updateBaasUserRequest)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}
		//fmt.Println(s.BlockService)
		_, err := s.ApplicationService.Application(r.Context(), payload.ApplicationID)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrApplicationNotFound,
				"status": 0,
			})
			return
		}
		if payload.ApplicationID == "" && payload.Email == "" && payload.Password == "" {
			RespondOK(w, map[string]interface{}{
				"error":  "Parameter is no valid",
				"status": 0,
			})
			return
		}

		err = s.BlockService.DeleteKey(r.Context(), payload.Email, payload.Password, payload.ApplicationID)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err,
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
		}

		RespondOK(w, resp)
	})
}

func (s *Server) handleGetAddress() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(createBaasUserRequest)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if err := payload.validate(); err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  "Parameter is no valid",
				"status": 0,
			})
			return
		}
		//fmt.Println(s.BlockService)
		a, err := s.BlockService.GetAddress(r.Context(), payload.Email, payload.Password)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasNoSuchUser,
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]string{
				"address": a,
			},
		}

		RespondOK(w, resp)
	})
}

func (s *Server) handleGetBalance() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  "Parameter error",
				"status": 0,
			})
			return
		}
		//fmt.Println(s.BlockService)
		a, err := s.BlockService.GetBalance(r.Context(), payload.Userid)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]string{
				"balance": a,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

func (s *Server) handleSendToken() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 || len(payload.Targetid) == 0 || len(payload.Value) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  "Parameter error",
				"status": 0,
			})
			return
		}
		//fmt.Println(s.BlockService)
		a, err := s.BlockService.SendToken(r.Context(), payload.Userid, payload.Password, payload.Targetid, payload.Value, payload.Message)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]string{
				"txhash": a,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

func (s *Server) handleWriteMsg() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 || len(payload.Targetid) == 0 || len(payload.Message) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  "Parameter error",
				"status": 0,
			})
			return
		}
		//fmt.Println(s.BlockService)
		a, err := s.BlockService.WriteMsg(r.Context(), payload.Userid, payload.Password, payload.Targetid, payload.Message)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]string{
				"txhash": a,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

func (s *Server) handleReadMsg() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Hash) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, t, err := s.BlockService.ReadMsg(r.Context(), payload.Hash)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]string{
				"txhash":    a,
				"timestamp": t,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

func (s *Server) handleGetTxByHash() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Hash) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetTxByHash(r.Context(), payload.Hash)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}
		um := map[string]interface{}{}
		err = json.Unmarshal([]byte(a), &um)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}
		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"Tx": um,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 获取一个用户的交易列表
func (s *Server) handleGetUserTxList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		var size = "25"
		var page = fmt.Sprint(payload.Page)
		/* if payload.Page > 0 {
			page = fmt.Sprint(payload.Page)
		} else if payload.Page == 0 {
			page = "1"
		} else if payload.Page == -99 {
			page = "-99"
		} */
		a, err := s.BlockService.GetTxByUserAddress(r.Context(), payload.Userid, page, size)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}
		dataArray := []map[string]interface{}{} // empty

		err = json.Unmarshal([]byte(a), &dataArray)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  "Internal Data Error",
				"status": 0,
			})
			return
		}
		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"Tx": dataArray,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 获取一个用户对特定用户的交易列表
func (s *Server) handleGetUserToUserTxList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 || len(payload.Targetid) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		var size = "25"
		var page = fmt.Sprint(payload.Page)
		/* if payload.Page > 0 {
			page = fmt.Sprint(payload.Page)
		} else if payload.Page == 0 {
			page = "1"
		} else if payload.Page == -99 {
			page = "-99"
		} */
		a, err := s.BlockService.GetPeerToPeerTxByUserAddress(r.Context(), payload.Userid, payload.Targetid, page, size)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}
		dataArray := []map[string]interface{}{} // empty

		err = json.Unmarshal([]byte(a), &dataArray)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  "Internal Data Service Error",
				"status": 0,
			})
			return
		}
		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"Tx": dataArray,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 合约mint一个token
func (s *Server) handleCreate721Token() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.TokenId) == 0 || len(payload.Meta) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.CreateErc721Token(r.Context(), payload.UserId, payload.Password, payload.Contract, payload.TokenId, payload.Meta)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"txHash": a,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 获得info
func (s *Server) handleGet721Info() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetErc721Info(r.Context(), payload.Contract)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data":   a,
			"error":  "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 获得info
func (s *Server) handleGet721Balance() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Contract) == 0 || len(payload.UserId) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetErc721Balance(r.Context(), payload.UserId, payload.Contract)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"Balance": a,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 获得 metadata
func (s *Server) handleGet721MetaData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Contract) == 0 || len(payload.TokenId) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetErc721MetaData(r.Context(), payload.Contract, payload.TokenId)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"metaData": fmt.Sprintf("%s", a),
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 send a toke to another address
func (s *Server) handleSend721Token() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.TokenId) == 0 || len(payload.Contract) == 0 || len(payload.TargetId) == 0 || len(payload.Password) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.SendErc721Token(r.Context(), payload.UserId, payload.Password, payload.Contract, payload.TokenId, payload.TargetId)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"txHash": a,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 获取一个所有721的交易列表
func (s *Server) handleGet721TxList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		var size = "25"
		var page = fmt.Sprint(payload.Page)
		fmt.Println(size, page)
		/*
			a, err := s.BlockService.GetErc721TxList(r.Context(), payload.Userid, page, size)
			if err != nil {
				RespondOK(w, map[string]interface{}{
					"error":  err.Error(),
					"status": 0,
				})
				return
			}
			dataArray := []map[string]interface{}{} // empty

			err = json.Unmarshal([]byte(a), &dataArray)
			if err != nil {
				RespondOK(w, map[string]interface{}{
					"error":  "Internal Data Error",
					"status": 0,
				})
				return
			} */
		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"Tx": "UnderConstruction",
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 获取一个所有721 通证的交易列表
func (s *Server) GetErc721TokenTxList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		var size = "25"
		var page = fmt.Sprint(payload.Page)
		fmt.Println(size, page)
		/*
			a, err := s.BlockService.GetErc721TxList(r.Context(), payload.Userid, page, size)
			if err != nil {
				RespondOK(w, map[string]interface{}{
					"error":  err.Error(),
					"status": 0,
				})
				return
			}
			dataArray := []map[string]interface{}{} // empty

			err = json.Unmarshal([]byte(a), &dataArray)
			if err != nil {
				RespondOK(w, map[string]interface{}{
					"error":  "Internal Data Error",
					"status": 0,
				})
				return
			} */
		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"Tx": "UnderConstruction",
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 获取一个所有用户关于721的交易列表
func (s *Server) GetErc721TxListByUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		var size = "25"
		var page = fmt.Sprint(payload.Page)
		fmt.Println(size, page)
		/*
			a, err := s.BlockService.GetErc721TxList(r.Context(), payload.Userid, page, size)
			if err != nil {
				RespondOK(w, map[string]interface{}{
					"error":  err.Error(),
					"status": 0,
				})
				return
			}
			dataArray := []map[string]interface{}{} // empty

			err = json.Unmarshal([]byte(a), &dataArray)
			if err != nil {
				RespondOK(w, map[string]interface{}{
					"error":  "Internal Data Error",
					"status": 0,
				})
				return
			} */
		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"Tx": "UnderConstruction",
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 获取一个所有721 通证对于某用户的交易列表
func (s *Server) GetErc721TokenTxListByUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		var size = "25"
		var page = fmt.Sprint(payload.Page)
		fmt.Println(size, page)
		/*
			a, err := s.BlockService.GetErc721TxList(r.Context(), payload.Userid, page, size)
			if err != nil {
				RespondOK(w, map[string]interface{}{
					"error":  err.Error(),
					"status": 0,
				})
				return
			}
			dataArray := []map[string]interface{}{} // empty

			err = json.Unmarshal([]byte(a), &dataArray)
			if err != nil {
				RespondOK(w, map[string]interface{}{
					"error":  "Internal Data Error",
					"status": 0,
				})
				return
			} */
		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"Tx": "UnderConstruction",
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 获得info
func (s *Server) handleGetBlockNumber() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		a, err := s.BlockService.GetBlockNumber(r.Context())
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"count": a,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

func (s *Server) handleUsages() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//applicationID := mux.Vars(r)["application_id"]
		payload := new(BlockChainUsageQuery)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		startDate, err := parseTime(payload.StartDate)
		if err != nil {
			RespondError(w, baas.ErrInvalidDate, http.StatusBadRequest, GetReqID(r))
			return
		}
		endDate, err := parseTime(payload.EndDate)
		if err != nil {
			RespondError(w, baas.ErrInvalidDate, http.StatusBadRequest, GetReqID(r))
			return
		}
		applicationID := payload.ApplicationID
		usages, err := s.UsageService.ApplicationUsages(r.Context(), applicationID, startDate, endDate)
		if err != nil {
			RespondError(w, err, http.StatusInternalServerError, GetReqID(r))
			return
		}
		data := sanitizeUsages(usages)
		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"txcounts":   data.Totals.TxCount,
				"usercounts": data.Totals.UserCount,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}
