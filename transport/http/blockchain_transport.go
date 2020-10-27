package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
				"error":  err.Error(),
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

func (s *Server) handleImportKey() http.Handler {
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
				"error":  err.Error(),
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
			Key:           payload.Key,
		}
		//fmt.Println(s.BlockService)
		a, err := s.BlockService.Import(r.Context(), n)
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

func (s *Server) handleGetAddressByPK() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(createBaasUserRequest)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Key) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetAddressFromPK(r.Context(), payload.Key)
		if err != nil {
			//RespondError(w, err, http.StatusInternalServerError, GetReqID(r))
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status":  1,
			"address": a,
		}

		RespondOK(w, resp)
	})
}

func (s *Server) handleGetPKByKeyStore() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(createBaasUserRequest)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.KeyStore) == 0 || len(payload.Password) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetPKFromKeyStore(r.Context(), payload.KeyStore, payload.Password)
		if err != nil {
			//RespondError(w, err, http.StatusInternalServerError, GetReqID(r))
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status":     1,
			"privatekey": a,
		}

		RespondOK(w, resp)
	})
}

func (s *Server) handleGetKeyStoreByPK() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(createBaasUserRequest)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Key) == 0 || len(payload.Password) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetKeyStoreFromPK(r.Context(), payload.Key, payload.Password)
		if err != nil {
			//RespondError(w, err, http.StatusInternalServerError, GetReqID(r))
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status":   1,
			"keystore": a,
		}

		RespondOK(w, resp)
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
				"error":  err.Error(),
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

		if len(payload.Email) == 0 || len(payload.Password) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
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

		if len(payload.Userid) == 0 || len(payload.Targetid) == 0 || payload.Value == 0.0 {
			RespondOK(w, map[string]interface{}{
				"error":  "Parameter error",
				"status": 0,
			})
			return
		}
		if len(payload.Encode) == 0 {
			payload.Encode = "utf8"
		} else if payload.Encode != "utf8" && payload.Encode != "hex" {
			payload.Encode = "utf8"
		}

		//fmt.Println(s.BlockService)
		a, err := s.BlockService.SendToken(r.Context(), payload.Userid, payload.Password, payload.Targetid, fmt.Sprintf("%f", payload.Value), payload.Message, payload.Encode, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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
		a, err := s.BlockService.WriteMsg(r.Context(), payload.Userid, payload.Password, payload.Targetid, payload.Message, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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
		if len(payload.Encode) == 0 {
			payload.Encode = "utf8"
		} else if payload.Encode != "hex" {
			payload.Encode = "utf8"
		}
		a, t, err := s.BlockService.ReadMsg(r.Context(), payload.Hash, payload.Encode)
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
				"msg":       a,
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

		a, err := s.BlockService.CreateErc721Token(r.Context(), payload.UserId, payload.Password, payload.Contract, payload.TokenId, payload.Meta, payload.Property, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

//设置token property
func (s *Server) handleSet721TokenProperty() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Contract) == 0 || len(payload.TokenId) == 0 || len(payload.Property) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.SetErc721TokenProperty(r.Context(), payload.UserId, payload.Password, payload.Contract, payload.TokenId, payload.Property, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

//设置token property
/*func (s *Server) xxhandleGet721TokenProperty() http.Handler {
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

		a, err := s.BlockService.GetErc721TokenProperty(r.Context(), payload.Contract, payload.TokenId)
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
*/
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
func (s *Server) handleGet721TokenProperty() http.Handler {
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

		a, err := s.BlockService.GetErc721TokenProperty(r.Context(), payload.Contract, payload.TokenId)
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
				"property": fmt.Sprintf("%s", a),
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

		a, err := s.BlockService.SendErc721Token(r.Context(), payload.UserId, payload.Password, payload.Contract, payload.TokenId, payload.Memo, payload.TargetId, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

// 调用erc721 add another token memo address
func (s *Server) handleAdd721TokenMemo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Memo) == 0 || len(payload.TokenId) == 0 || len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.AddErc721TokenMemo(r.Context(), payload.UserId, payload.Password, payload.Contract, payload.TokenId, payload.Memo, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

// 调用erc721 token memolist
func (s *Server) handleGet721TokenMemoList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.TokenId) == 0 || len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetErc721TokenMemoList(r.Context(), payload.Contract, payload.TokenId)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}
		list := make([]map[string]interface{}, 0)
		err = json.Unmarshal([]byte(a), &list)
		//fmt.Println(list)
		resp := &map[string]interface{}{
			"status": 1,
			"data":   list,
			"error":  "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 send a toke to another address
func (s *Server) handleGetSend721TokenMemo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Hash) == 0 { //|| len(payload.TokenId) == 0 || len(payload.Contract) == 0
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetSendErc721TokenMemo(r.Context(), payload.Contract, payload.TokenId, payload.Hash)
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
				"memo": a,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 send a toke to another address
func (s *Server) handleGetUser721Token() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetUserErc721TokenList(r.Context(), payload.UserId, payload.Contract)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}
		list := make([]map[string]interface{}, 0)
		err = json.Unmarshal([]byte(a), &list)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}
		resp := &map[string]interface{}{
			"status": 1,
			"data":   list,
			"error":  "",
		}

		RespondOK(w, resp)
	})
}

// deploy contranct
func (s *Server) handleDeploy721Contract() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(DeployOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Password) == 0 || len(payload.Name) == 0 || len(payload.Symbol) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, hash, err := s.BlockService.CreateERC721Contract(r.Context(), payload.UserId, payload.Password, payload.Name, payload.Symbol, fmt.Sprint(payload.Class))
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + hash,
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"contract": a,
				"txHash":   hash,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

/// Follow function handle query stuff
// 获取一个所有721的交易列表
func (s *Server) handleGet721TxList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
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
		var size = "25"
		var page = fmt.Sprint(payload.Page)
		//fmt.Println(size, page)

		a, err := s.BlockService.GetErc721TxList(r.Context(), payload.Contract, page, size)
		if err == baas.ErrBaasQueryNoResult {
			RespondOK(w, map[string]interface{}{
				"error":  "",
				"status": 0,
				"Data":   []interface{}{},
			})
			return
		} else if err != nil {
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
				"error":  err.Error(),
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

// 获取一个所有721 通证的交易列表
func (s *Server) GetErc721TokenTxList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
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
		var size = "25"
		var page = fmt.Sprint(payload.Page)
		fmt.Println(size, page)

		a, err := s.BlockService.GetErc721TokenTxList(r.Context(), payload.Contract, payload.TokenId, page, size)
		if err == baas.ErrBaasQueryNoResult {
			RespondOK(w, map[string]interface{}{
				"error":  "",
				"status": 0,
				"Data":   []interface{}{},
			})
			return
		} else if err != nil {
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

// 获取一个所有用户关于721的交易列表
func (s *Server) GetErc721TxListByUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 || len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		var size = "25"
		var page = fmt.Sprint(payload.Page)
		//fmt.Println(size, page)

		a, err := s.BlockService.GetErc721TxListByUser(r.Context(), payload.Userid, payload.Contract, page, size)
		if err == baas.ErrBaasQueryNoResult {
			RespondOK(w, map[string]interface{}{
				"error":  "",
				"status": 0,
				"Data":   []interface{}{},
			})
			return
		} else if err != nil {
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

// 获取一个所有721 通证对于某用户的交易列表
func (s *Server) GetErc721TokenTxListByUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 || len(payload.Contract) == 0 || len(payload.TokenId) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		var size = "25"
		var page = fmt.Sprint(payload.Page)
		//fmt.Println(size, page)

		a, err := s.BlockService.GetErc721TokenTxListByUser(r.Context(), payload.Userid, payload.Contract, payload.TokenId, page, size)
		if err == baas.ErrBaasQueryNoResult {
			RespondOK(w, map[string]interface{}{
				"error":  "",
				"status": 0,
				"Data":   []interface{}{},
			})
			return
		} else if err != nil {
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

/// Following is erc20 process portion =====================================================

// deploy contranct
func (s *Server) handleDeploy20Token() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(DeployOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Password) == 0 || len(payload.Name) == 0 || len(payload.Symbol) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		if payload.TotalSupply == 0.0 || payload.Decimal == 0 || (payload.Class == 4 && payload.Capacity == 0) {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasContractInValid,
				"status": 0,
			})
			return
		}
		a, hash, err := s.BlockService.DeployERC20Contract(r.Context(), payload.UserId, payload.Password, payload.Name, payload.Symbol, fmt.Sprint(payload.Class), payload.TotalSupply, payload.Decimal, payload.Capacity)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + hash,
				"status": 0,
			})
			return
		}

		resp := &map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"contract": a,
				"txHash":   hash,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc20 获得info
func (s *Server) handleGet20Info() http.Handler {
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

		a, err := s.BlockService.GetErc20Info(r.Context(), payload.Contract)
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

func (s *Server) handleGet20Balance() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 || len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.GetErc20Balance(r.Context(), payload.Userid, payload.Contract)
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

// 调用erc721 send a toke to another address
func (s *Server) handleSend20Token() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Contract) == 0 || len(payload.TargetId) == 0 || len(payload.Password) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		if payload.Quantity == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  "wrong quantity",
				"status": 0,
			})
			return
		}
		vv := payload.Quantity
		/* err := strconv.ParseFloat(payload.Quantity, 64)
		if err != nil {
			s.Log.Err(err).Msg("Quantity  format is not correct")
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasInvalidQuantity,
				"status": 0,
			})
			return
		} */

		a, err := s.BlockService.SendErc20Token(r.Context(), payload.UserId, payload.Password, payload.TargetId, payload.Contract, payload.Memo, vv, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

// 调用erc721 add another token memo address
func (s *Server) handleApprove() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Password) == 0 || len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		if payload.Quantity == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  "wrong quantity",
				"status": 0,
			})
			return
		}
		a, err := s.BlockService.ApproveErc20(r.Context(), payload.UserId, payload.Password, payload.TargetId, payload.Contract, payload.Quantity, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

// 调用erc721 add another token memo address
func (s *Server) handleAllowance() http.Handler {
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

		a, err := s.BlockService.AllowanceErc20(r.Context(), payload.UserId, payload.TargetId, payload.Contract)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error(),
				"status": 0,
			})
			return
		}
		//nn, err := strconv.Atoi(a)
		nn, err := strconv.ParseFloat(a, 10)
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
				"allowance": nn,
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 add another token memo address
func (s *Server) handleIncreaseAllowance() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Password) == 0 || len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		if payload.Quantity == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  "wrong quantity",
				"status": 0,
			})
			return
		}
		a, err := s.BlockService.IncreaseAllowanceErc20(r.Context(), payload.UserId, payload.Password, payload.TargetId, payload.Contract, payload.Quantity, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

// 调用erc721 add another token memo address
func (s *Server) handleDecreaseAllowance() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Password) == 0 || len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		if payload.Quantity == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  "wrong quantity",
				"status": 0,
			})
			return
		}
		a, err := s.BlockService.DecresaseAllowanceErc20(r.Context(), payload.UserId, payload.Password, payload.TargetId, payload.Contract, payload.Quantity, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

// 调用erc721 add another token memo address
func (s *Server) handleTransferFrom() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Contract) == 0 || len(payload.TargetId) == 0 || len(payload.Password) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		if len(payload.SendUserId) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  "no sender id",
				"status": 0,
			})
			return
		}
		if payload.Quantity == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  "wrong quantity",
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.TransferFromErc20(r.Context(), payload.UserId, payload.Password, payload.SendUserId, payload.TargetId, payload.Contract, payload.Memo, payload.Quantity, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

// 调用erc721 获得 metadata
func (s *Server) handleGet20TxMemo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
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

		a, err := s.BlockService.GetErc20TxMemo(r.Context(), payload.Hash)
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
				"memo": fmt.Sprintf("%s", a),
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 add another token memo address
func (s *Server) handleBurn20Token() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Password) == 0 || len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		if payload.Quantity == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasInvalidQuantity,
				"status": 0,
			})
			return
		}

		a, err := s.BlockService.BurnErc20(r.Context(), payload.UserId, payload.Password, payload.Contract, payload.Quantity, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

// 调用erc721 add another token memo address
func (s *Server) handlePause20Token() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Contract) == 0 || len(payload.UserId) == 0 || len(payload.Password) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		if len(payload.Action) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBassContractCmdNotFound,
				"status": 0,
			})
			return
		}
		var action bool
		if payload.Action == "pause" {
			action = true
		}
		if payload.Action == "start" {
			action = false
		}
		a, err := s.BlockService.PauseErc20(r.Context(), payload.UserId, payload.Password, payload.Contract, action, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

// 调用erc721 获得 metadata
func (s *Server) handleGet20PauseStatus() http.Handler {
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

		a, err := s.BlockService.GetErc20PauseStatus(r.Context(), payload.Contract)
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
				"status": fmt.Sprintf("%s", a),
			},
			"error": "",
		}

		RespondOK(w, resp)
	})
}

// 调用erc721 add another token memo address
func (s *Server) handleMint20Token() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(SmartContractOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.UserId) == 0 || len(payload.Password) == 0 || len(payload.Contract) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		if payload.Quantity == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasInvalidQuantity,
				"status": 0,
			})
			return
		}
		a, err := s.BlockService.MintErc20(r.Context(), payload.UserId, payload.Password, payload.Contract, payload.Quantity, payload.GasLimit)
		if err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  err.Error() + ":" + a,
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

// 获取一个用户的交易列表
func (s *Server) handleGet20TxList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
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
		var size = "25"
		var page = fmt.Sprint(payload.Page)

		a, err := s.BlockService.GetERC20TxList(r.Context(), payload.Contract, page, size)
		if err == baas.ErrBaasQueryNoResult {
			RespondOK(w, map[string]interface{}{
				"error":  "",
				"status": 0,
				"Data":   []interface{}{},
			})
			return
		} else if err != nil {
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
				"error":  err.Error(),
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

// 获取一个用户的交易列表
func (s *Server) handleGet20UserTxList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := new(BlockOperation)
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			RespondError(w, err, http.StatusBadRequest, GetReqID(r))
			return
		}

		if len(payload.Userid) == 0 || len(payload.Contract) == 0 {
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
		a, err := s.BlockService.GetERC20TxByUserAddress(r.Context(), payload.Userid, payload.Contract, page, size)
		if err == baas.ErrBaasQueryNoResult {
			RespondOK(w, map[string]interface{}{
				"error":  "",
				"status": 0,
				"Data":   []interface{}{},
			})
			return
		} else if err != nil {
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
