package e2e

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/airbloc/airbloc-go/shared/blockchain"

	e2eutils "github.com/airbloc/airbloc-go/test/e2e/utils"

	testutils "github.com/airbloc/airbloc-go/test/utils"

	"github.com/airbloc/airbloc-go/provider/api"
	apilib "github.com/airbloc/airbloc-go/shared/service/api"
	serviceMock "github.com/airbloc/airbloc-go/shared/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getAvailablePort() (port int, rerr error) {
	server, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer func() {
		err = server.Close()
		if err != nil {
			rerr = err
		}
	}()

	host := server.Addr().String()

	_, portStr, err := net.SplitHostPort(host)
	if err != nil {
		return 0, err
	}

	// Return the port as an int
	p, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, err
	}

	return p, nil
}

var _ = Describe("Apps", func() {
	var (
		t            = GinkgoT()
		testAppName  = fmt.Sprintf("app-test-%s-%d", generateUniqueName(), rand.Int())
		testClient   *blockchain.Client
		testServer   *httptest.Server
		testEndpoint = func(path string) string {
			return testServer.URL + "/apps/" + path
		}
	)

	{
		BeforeEach(func() {
			ctrl := gomock.NewController(t)

			client, err := e2eutils.ConnectBlockchain()
			Ω(err).ShouldNot(HaveOccurred())
			testClient = client

			backend := serviceMock.NewMockBackend(ctrl)
			backend.EXPECT().Client().Return(testClient)

			_, engine := gin.CreateTestContext(nil)

			api, err := api.NewAppRegistryAPI(backend)
			Ω(err).ShouldNot(HaveOccurred())
			api.AttachToAPI(&apilib.Service{HttpServer: engine})

			testServer = httptest.NewServer(engine)
		})

		AfterEach(func() {
			testServer.Close()
			testClient.Close()
		})
	}

	It("should register app", func() {
		req, err := e2eutils.CreateRequest(
			e2eutils.MethodPost, testEndpoint(""),
			gin.H{"app_name": testAppName}, e2eutils.RequestJSON,
		)
		Ω(err).ShouldNot(HaveOccurred())

		resp, body, err := req.Do()
		Ω(err).ShouldNot(HaveOccurred())

		t.Log("code", resp.StatusCode)
		t.Log("body", string(body))

		Ω(resp.StatusCode).Should(Equal(http.StatusOK))
		Ω(string(body)).Should(Equal(testutils.TestSuccessStr))
	})

	//It("should retrieve app", func() {
	//	req, err := e2eutils.CreateRequest(
	//		e2eutils.MethodGet, testEndpoint(""),
	//		gin.H{"app_name": testAppName}, e2eutils.RequestQuery,
	//	)
	//	Ω(err).ShouldNot(HaveOccurred())
	//
	//	resp, body, err := req.Do()
	//	Ω(err).ShouldNot(HaveOccurred())
	//
	//	t.Log("code", resp.StatusCode)
	//	t.Log("body", string(body))
	//
	//	Ω(resp.StatusCode).Should(Equal(http.StatusOK))
	//	d, err := json.Marshal(types.App{
	//		Name:  testAppName,
	//		Owner: testClient.Account().From,
	//		Addr:  common.BytesToAddress(crypto.Keccak256([]byte(testAppName))),
	//	})
	//	Ω(body).Should(Equal(d))
	//})
	//
	//It("should not retrive app when it does not registered", func() {
	//
	//})

	It("should unregister app", func() {
		req, err := e2eutils.CreateRequest(
			e2eutils.MethodDelete, testEndpoint(""),
			gin.H{"app_name": testAppName}, e2eutils.RequestJSON,
		)
		Ω(err).ShouldNot(HaveOccurred())

		resp, body, err := req.Do()
		Ω(err).ShouldNot(HaveOccurred())

		t.Log("code", resp.StatusCode)
		t.Log("body", string(body))

		Ω(resp.StatusCode).Should(Equal(http.StatusOK))
		Ω(string(body)).Should(Equal(testutils.TestSuccessStr))
	})
})
