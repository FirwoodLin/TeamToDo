package response

var (
	UnauthorizedError = MakeFailedResponse("Unauthorized")
	InvalidInfoError  = MakeFailedResponse("Invalid information")
	TimeoutError      = MakeFailedResponse("Timeout")
)
