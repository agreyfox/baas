package http

import (
	"encoding/json"
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

		if err := payload.validate(); err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  "Password is not secure",
				"status": 0,
			})
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
			Data: sanitizeBaasUser(a),
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

		if err := payload.validate(); err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
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

		if err := payload.validate(); err != nil {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasParameterNotFound,
				"status": 0,
			})
			return
		}
		if len(payload.CipherText) == 0 {
			RespondOK(w, map[string]interface{}{
				"error":  baas.ErrBaasCipherTextRequired,
				"status": 0,
			})
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
		//fmt.Println(s.BlockService)
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

		if payload.ApplicationID == "" && payload.Email == "" && payload.Password == "" {
			RespondOK(w, map[string]interface{}{
				"error":  "Parameter is no valid",
				"status": 0,
			})
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
		err = s.BlockService.DeleteKey(r.Context(), payload.Email, payload.Password)
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
		a, err := s.BlockService.SendToken(r.Context(), payload.Userid, payload.Password, payload.Targetid, payload.Value)
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
