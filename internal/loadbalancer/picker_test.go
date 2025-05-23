package loadbalancer_test

// import (
// 	"testing"

// 	"google.golang.org/grpc/attributes"
// 	"google.golang.org/grpc/balancer"
// 	"google.golang.org/grpc/balancer/base"
// 	"google.golang.org/grpc/connectivity"
// 	"google.golang.org/grpc/resolver"

// 	"github.com/stretchr/testify/require"

// 	"github.com/YashTike/proglog/internal/loadbalancer"
// )

// func TestPickerNoSubConnAvailable(t *testing.T) {
// 	picker := &loadbalancer.Picker{}
// 	for _, method := range []string{
// 		"/log.vX.Log/Produce",
// 		"/log.vX.Log/Consume",
// 	} {
// 		info := balancer.PickInfo{
// 			FullMethodName: method,
// 		}
// 		result, err := picker.Pick(info)
// 		require.Equal(t, balancer.ErrNoSubConnAvailable, err)
// 		require.Nil(t, result.SubConn)
// 	}
// }

// func TestPickerProducesToLeader(t *testing.T) {
// 	picker, subConns := setupTest()
// 	info := balancer.PickInfo{
// 		FullMethodName: "/log.vX.Log/Produce",
// 	}
// 	for i := 0; i < 5; i++ {
// 		gotPick, err := picker.Pick(info)
// 		require.NoError(t, err)
// 		require.Equal(t, subConns[0], gotPick.SubConn)
// 	}
// }

// func TestPickerConsumesFromFollowers(t *testing.T) {
// 	picker, subConns := setupTest()
// 	info := balancer.PickInfo{
// 		FullMethodName: "/log.vX.Log/Consume",
// 	}
// 	for i := 0; i < 5; i++ {
// 		pick, err := picker.Pick(info)
// 		require.NoError(t, err)
// 		require.Equal(t, subConns[i%2+1], pick.SubConn)
// 	}
// }

// func setupTest() (*loadbalancer.Picker, []*subConn) {
// 	var subConns []*subConn
// 	buildInfo := base.PickerBuildInfo{
// 		ReadySCs: make(map[balancer.SubConn]base.SubConnInfo),
// 	}
// 	for i := 0; i < 3; i++ {
// 		sc := &subConn{}
// 		addr := resolver.Address{
// 			Attributes: attributes.New("is_leader", i == 0),
// 		}
// 		// 0th sub conn is the leader
// 		sc.UpdateAddresses([]resolver.Address{addr})
// 		buildInfo.ReadySCs[sc] = base.SubConnInfo{Address: addr}
// 		subConns = append(subConns, sc)
// 	}
// 	picker := &loadbalancer.Picker{}
// 	picker.Build(buildInfo)
// 	return picker, subConns
// }

// // subConn implements balancer.SubConn.
// type subConn struct {
// 	addrs []resolver.Address
// }

// func (s *subConn) UpdateAddresses(addrs []resolver.Address) {
// 	s.addrs = addrs
// }

// func (s *subConn) Connect() {}

// func (s *subConn) GetOrBuildProducer(balancer.ProducerBuilder) (balancer.Producer, func()) {
// 	return nil, nil
// }

// func (s *subConn) RegisterHealthListener(func(balancer.SubConnState)) {}

// func (s *subConn) Shutdown() {}

// type SubConnState struct {
// 	ConnectivityState connectivity.State
// 	ConnectionError   error
// }

// var _ balancer.SubConn = (*subConn)(nil)
