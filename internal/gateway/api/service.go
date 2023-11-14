package api

import "github.com/KentaOKada/line-go-sample/internal/domain/usecase/line"

func NewLineServiceServer(
	interactor line.Interactor,
) *LineServiceServer {
	return &LineServiceServer{
		interactor: interactor,
	}
}

type LineServiceServer struct {
	interactor line.Interactor
}
