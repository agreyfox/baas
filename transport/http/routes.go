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

	const baasGetUserTxList = "/api/baas/getUserTxList"
	s.Handle(baasGetUserTxList,
		s.Authenticate(
			s.handleGetUserTxList())).Methods("POST")
	const baasGetUserToUserTxList = "/api/baas/getUserToUserTxList"
	s.Handle(baasGetUserToUserTxList,
		s.Authenticate(
			s.handleGetUserToUserTxList())).Methods("POST")

	//////////////////////////ERC721//////////////////////////////
	const baasCreate721Token = "/api/baas/erc721/createToken"
	s.Handle(baasCreate721Token,
		s.Authenticate(
			s.handleCreate721Token())).Methods("POST")

	const baasSet721TokenProperty = "/api/baas/erc721/setTokenProperty"
	s.Handle(baasSet721TokenProperty,
		s.Authenticate(
			s.handleSet721TokenProperty())).Methods("POST")
	const baasGet721TokenProperty = "/api/baas/erc721/getTokenProperty"
	s.Handle(baasGet721TokenProperty,
		s.Authenticate(
			s.handleGet721TokenProperty())).Methods("POST")

	const baasGet721Info = "/api/baas/erc721/getContractInfo"
	s.Handle(baasGet721Info,
		s.Authenticate(
			s.handleGet721Info())).Methods("POST")

	const baasGet721Balance = "/api/baas/erc721/getBalance"
	s.Handle(baasGet721Balance,
		s.Authenticate(
			s.handleGet721Balance())).Methods("POST")

	const baasGet721MetaData = "/api/baas/erc721/getTokenMetadata"
	s.Handle(baasGet721MetaData,
		s.Authenticate(
			s.handleGet721MetaData())).Methods("POST")

	const baasSend721 = "/api/baas/erc721/sendToken"
	s.Handle(baasSend721,
		s.Authenticate(
			s.handleSend721Token())).Methods("POST")
	const baasGetSend721Memo = "/api/baas/erc721/getSendTokenMemo"
	s.Handle(baasGetSend721Memo,
		s.Authenticate(
			s.handleGetSend721TokenMemo())).Methods("POST")
	const baasAddErc721TokenMemo = "/api/baas/erc721/addTokenMemo"
	s.Handle(baasAddErc721TokenMemo,
		s.Authenticate(
			s.handleAdd721TokenMemo())).Methods("POST")

	const baasGetErc721TokenMemoList = "/api/baas/erc721/getTokenMemoList"
	s.Handle(baasGetErc721TokenMemoList,
		s.Authenticate(
			s.handleGet721TokenMemoList())).Methods("POST")
	const baasDeploy721 = "/api/baas/erc721/createContract"
	s.Handle(baasDeploy721,
		s.Authenticate(
			s.handleDeploy721Contract())).Methods("POST")

	//////////////////////////ERC20//////////////////////////////
	const baasDeploy20 = "/api/baas/erc20/createToken"
	s.Handle(baasDeploy20,
		s.Authenticate(
			s.handleDeploy20Token())).Methods("POST")

	const baasGet20Info = "/api/baas/erc20/getContractInfo"
	s.Handle(baasGet20Info,
		s.Authenticate(
			s.handleGet20Info())).Methods("POST")

	const baasGet20Balance = "/api/baas/erc20/getBalance"
	s.Handle(baasGet20Balance,
		s.Authenticate(
			s.handleGet20Balance())).Methods("POST")

	const baasSend20 = "/api/baas/erc20/transfer"
	s.Handle(baasSend20,
		s.Authenticate(
			s.handleSend20Token())).Methods("POST")

	const baasApprove20 = "/api/baas/erc20/approve"
	s.Handle(baasApprove20,
		s.Authenticate(
			s.handleApprove())).Methods("POST")

	const baasAllowance = "/api/baas/erc20/allowance"
	s.Handle(baasAllowance,
		s.Authenticate(
			s.handleAllowance())).Methods("POST")

	const baasIncreaseAllowance = "/api/baas/erc20/increaseAllowance"
	s.Handle(baasIncreaseAllowance,
		s.Authenticate(
			s.handleIncreaseAllowance())).Methods("POST")

	const baasDecreaseAllowance = "/api/baas/erc20/decreaseAllowance"
	s.Handle(baasDecreaseAllowance,
		s.Authenticate(
			s.handleDecreaseAllowance())).Methods("POST")

	const baasTransferFrom = "/api/baas/erc20/transferFrom"
	s.Handle(baasTransferFrom,
		s.Authenticate(
			s.handleTransferFrom())).Methods("POST")

	const baasGetSend20Memo = "/api/baas/erc20/getTxMemo"
	s.Handle(baasGetSend20Memo,
		s.Authenticate(
			s.handleGet20TxMemo())).Methods("POST")

	const baasBurn20Token = "/api/baas/erc20/burnToken"
	s.Handle(baasBurn20Token,
		s.Authenticate(
			s.handleBurn20Token())).Methods("POST")

	const baasPause20Token = "/api/baas/erc20/pauseToken"
	s.Handle(baasPause20Token,
		s.Authenticate(
			s.handlePause20Token())).Methods("POST")

	const baasGet20Status = "/api/baas/erc20/getPauseStatus"
	s.Handle(baasGet20Status,
		s.Authenticate(
			s.handleGet20PauseStatus())).Methods("POST")
	const baasMint20Token = "/api/baas/erc20/mintToken"
	s.Handle(baasMint20Token,
		s.Authenticate(
			s.handleMint20Token())).Methods("POST")

	//////////////////////////////////////
	//////get erc721   logs             //////
	/////////////////////////////////////
	const baasGetUser721Token = "/api/baas/erc721/getTokenList"
	s.Handle(baasGetUser721Token,
		s.Authenticate(
			s.handleGetUser721Token())).Methods("POST")

	const baasGet721List = "/api/baas/erc721/getTxList"
	s.Handle(baasGet721List,
		s.Authenticate(
			s.handleGet721TxList())).Methods("POST")

	const baasGet721TokenList = "/api/baas/erc721/getTokenTxList"
	s.Handle(baasGet721TokenList,
		s.Authenticate(
			s.GetErc721TokenTxList())).Methods("POST")

	const baasGet721UserTxList = "/api/baas/erc721/getUserTxList"
	s.Handle(baasGet721UserTxList,
		s.Authenticate(
			s.GetErc721TxListByUser())).Methods("POST")
	const baasGet721TokenUserTxList = "/api/baas/erc721/getTokenUserTxList"
	s.Handle(baasGet721TokenUserTxList,
		s.Authenticate(
			s.GetErc721TokenTxListByUser())).Methods("POST")
	//////////////////////////////////////
	//////get erc20   logs             //////
	/////////////////////////////////////
	const baasGet20TxLog = "/api/baas/erc20/getTxList"
	s.Handle(baasGet20TxLog,
		s.Authenticate(
			s.handleGet20TxList())).Methods("POST")

	const baasGet20UserList = "/api/baas/erc20/getUserTxList"
	s.Handle(baasGet20UserList,
		s.Authenticate(
			s.handleGet20UserTxList())).Methods("POST")
	/////////////////////////////////
	//other api methods
	////////////////////////////////
	const baasGetBlockCount = "/api/baas/getBlockCount"
	s.Handle(baasGetBlockCount,
		s.Authenticate(
			s.handleGetBlockNumber())).Methods("POST")
	const baasGetTotalTxCounts = "/api/baas/getTotalTxCounts"
	s.Handle(baasGetTotalTxCounts,
		s.Authenticate(
			s.handleUsages())).Methods("POST")
	const baasGetTotalUserCount = "/api/baas/getTotalUserCounts"
	s.Handle(baasGetTotalUserCount,
		s.Authenticate(
			s.handleUsages())).Methods("POST")

	//////////////////////////////
	// PPROF //
	/////////////////////////////
	s.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

}
