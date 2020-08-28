package http

import "net/http"

func (s *Server) routes() {
	//////////////////////////////
	// HEALTH //
	/////////////////////////////
	const healthCheckPath = "/health"
	s.Handle(healthCheckPath,
		s.handleHealth()).Methods("GET")

	//////////////////////////////
	// APPLICATIONS //
	/////////////////////////////
	const createApplicationPath = "/applications"
	s.Handle(createApplicationPath,
		s.Authenticate(
			s.handleCreateApplication())).Methods("POST")

	const getApplicationPath = "/applications/{id}"
	s.Handle(getApplicationPath,
		s.Authenticate(
			s.handleGetApplication())).Methods("GET")

	const listApplicationsPath = "/applications"
	s.Handle(listApplicationsPath,
		s.Authenticate(
			s.handleListApplications())).Methods("GET")

	const updateApplication = "/applications/{id}"
	s.Handle(updateApplication,
		s.Authenticate(
			s.handleUpdateApplication())).Methods("PUT")

	const deleteApplicationPath = "/applications/{id}"
	s.Handle(deleteApplicationPath,
		s.Authenticate(
			s.handleDeleteApplication())).Methods("DELETE")

	//////////////////////////////
	// UPLOAD //
	/////////////////////////////
	const uploadFilePath = "/upload"
	s.Handle(uploadFilePath,
		s.Authenticate(
			s.handleUpload())).Methods("POST")

	const uploadChunkFilePath = "/chunk-upload"
	s.Handle(uploadChunkFilePath,
		s.Authenticate(
			s.handleChunkUpload())).Methods("POST")

	const completedChunksPath = "/chunks-completed"
	s.Handle(completedChunksPath,
		s.Authenticate(
			s.handleCompleteChunkUpload())).Methods("POST")

	//////////////////////////////
	// USAGES //
	/////////////////////////////
	const listUsages = "/usages"
	s.Handle(listUsages,
		s.Authenticate(
			s.handleListUsages())).Methods("GET")

	const listApplicationUsages = "/applications/{application_id}/usages"
	s.Handle(listApplicationUsages,
		s.Authenticate(
			s.handleListApplicationUsages())).Methods("GET")

	//////////////////////////////
	// SERVE FILE //
	/////////////////////////////
	s.PathPrefix("/{app_name}/{date}/{file_name}").Handler(
		s.handleServeFile()).Methods("GET")

	//////////////////////////////
	// BLOCK //
	/////////////////////////////
	const baasCreateKey = "/api/baas/createKey"
	s.Handle(baasCreateKey,
		s.Authenticate(
			s.handleCreateKey())).Methods("POST")
	const baasChangePassword = "/api/baas/changePassword"
	s.Handle(baasChangePassword,
		s.Authenticate(
			s.handleChangePassword())).Methods("POST")
	const baasDeleteKey = "/api/baas/deleteKey"
	s.Handle(baasDeleteKey,
		s.Authenticate(
			s.handleDeleteKey())).Methods("POST")

	const baasGetKey = "/api/baas/getKey"
	s.Handle(baasGetKey,
		s.Authenticate(
			s.handleGetKey())).Methods("POST")
	const baasGetAddress = "/api/baas/getAddress"
	s.Handle(baasGetAddress,
		s.Authenticate(
			s.handleGetAddress())).Methods("POST")
	const baasGetBalance = "/api/baas/getBalance"
	s.Handle(baasGetBalance,
		s.Authenticate(
			s.handleGetBalance())).Methods("POST")

	const baasSendToken = "/api/baas/sendToken"
	s.Handle(baasSendToken,
		s.Authenticate(
			s.handleSendToken())).Methods("POST")
	const baasWriteMsg = "/api/baas/writeMsg"
	s.Handle(baasWriteMsg,
		s.Authenticate(
			s.handleWriteMsg())).Methods("POST")
	const baasReadMsg = "/api/baas/readMsg"
	s.Handle(baasReadMsg,
		s.Authenticate(
			s.handleReadMsg())).Methods("POST")

	const baasGetTx = "/api/baas/getTxByHash"
	s.Handle(baasGetTx,
		s.Authenticate(
			s.handleGetTxByHash())).Methods("POST")

	//////////////////////////////
	// PPROF //
	/////////////////////////////
	s.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

}
