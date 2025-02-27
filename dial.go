package morndial

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/metadata"
)

var (
	conn map[string]*grpc.ClientConn
	ma   map[string]*sync.Mutex
)

func (service *Morndial) NewConnection() error {
	contexto, cancel := context.WithTimeout(context.Background(), service.Timeout)
	defer cancel()

	service.uid = uuid.New()

	inter := grpc.WithUnaryInterceptor(func(ctx context.Context, metodo string, requisicao, resposta interface{}, cc *grpc.ClientConn, invocador grpc.UnaryInvoker, opcoes ...grpc.CallOption) error {
		valor := ctx.Value("RID").(string)
		ctx = metadata.AppendToOutgoingContext(ctx, "RID", valor)

		if _, ok := ctx.Deadline(); !ok {
			ctx2, cancelar := context.WithTimeout(ctx, service.Timeout)
			defer cancelar()
			ctx = ctx2
		}

		return invocador(ctx, metodo, requisicao, resposta, cc, opcoes...)
	})

	service.Interceptors = append(service.Interceptors, inter)

	var erro error
	if conn[service.Name], erro = grpc.DialContext(contexto, service.Url, service.Interceptors...); erro != nil {
		return erro
	}

	return nil
}

func Get(uid uuid.UUID) grpc.ClientConnInterface {
	in := services[uid]

	ma[in.Name].Lock()
	defer ma[in.Name].Unlock()

	if state := conn[in.Name].GetState(); state != connectivity.Ready {
		_ = in.NewConnectionWithContext(context.Background())
	}

	return conn[in.Name]
}
