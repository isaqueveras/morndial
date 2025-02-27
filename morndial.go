package morndial

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var servicesM *sync.Mutex
var services map[uuid.UUID]*Morndial

type Morndial struct {
	UID         uuid.UUID     `json:"uid"`
	Name        string        `json:"name"`
	Url         string        `json:"url"`
	Insecure    bool          `json:"insecure"`
	Timeout     time.Duration `json:"timeout"`
	Certificate Certificate   `json:"certificate"`

	Interceptors []grpc.DialOption
}

type Certificate struct {
	Crt string `json:"crt"`
	Key string `json:"key"`
	Ca  string `json:"ca"`
}

// NewService ...
func NewService(name, url string, insecure bool, timeout time.Duration, interceptors ...grpc.DialOption) {
	servicesM.Lock()
	defer servicesM.Unlock()

	services[uuid.New()] = &Morndial{
		UID:          uuid.New(),
		Name:         name,
		Url:          url,
		Insecure:     insecure,
		Timeout:      timeout,
		Interceptors: interceptors,
	}
}
