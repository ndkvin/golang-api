package handler

type Handler struct {
	AuthHandler AuthHandlerInterface
}

func NewHandler(authHandler AuthHandlerInterface) *Handler {
	return &Handler{
		AuthHandler: authHandler,
	}
}