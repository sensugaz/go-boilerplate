package handler_test

import (
	"net/http"
	"net/http/httptest"

	"gitlab.bitkub.io/bo/gokub-boilerplate/cmd/domain/user"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gitlab.bitkub.io/bo/gokub-boilerplate/cmd/handler"

	handlerMock "gitlab.bitkub.io/bo/gokub-boilerplate/cmd/handler/mock"
)

var _ = Describe("CreateUsers", Label("library"), func() {
	var (
		mockCtrl    *gomock.Controller
		mockService *handlerMock.MockUserService
		mockAppPort = "1234"
		mockServer  *echo.Echo

		h *handler.Handler
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockServer = echo.New()
		mockServer.Validator = &handler.MockValidator{
			Validator: validator.New(),
		}
		mockService = handlerMock.NewMockUserService(mockCtrl)

		h = &handler.Handler{
			Server:      mockServer,
			Port:        mockAppPort,
			UserService: mockService,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Given a valid request payload", func() {
		When("user call service to create a new user", func() {
			var (
				c echo.Context
			)

			BeforeEach(func() {
				mockUser := &user.User{
					Name:  "unit-test",
					Email: "unit-test@unittest.com",
				}

				mockService.EXPECT().Create(gomock.Any()).Return(mockUser, nil).Times(1)

				req := httptest.NewRequest(http.MethodPost, "/v1/users", nil)
				rec := httptest.NewRecorder()
				c = mockServer.NewContext(req, rec)
			})

			It("should create a new user successful", func() {
				err := h.CreateUserHandler(c)
				Expect(err).Should(BeNil())
			})
		})
	})
})
