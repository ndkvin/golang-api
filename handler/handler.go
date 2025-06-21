package handler

type Handler struct {
	AuthHandler AuthHandlerInterface
	LinkHandler LinkHandlerInterface
}

func NewHandler(authHandler AuthHandlerInterface, linkHandler LinkHandlerInterface) *Handler {
	return &Handler{
		AuthHandler: authHandler,
		LinkHandler: linkHandler,
	}
}