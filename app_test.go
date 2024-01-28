package core_test

import (
	"context"
	"testing"
	"time"

	"gitee.com/keenoho/go-core"
	middleware "gitee.com/keenoho/go-core/middleware"
	"gitee.com/keenoho/go-core/module"
	"gitee.com/keenoho/go-core/protobuf"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DemoGrpcService struct {
	protobuf.UnimplementedBaseServiceServer
}
type User struct {
	Id       int64     `gorm:"column:id;primaryKey" form:"id" json:"id"`
	Account  string    `gorm:"column:account;unique;not null" form:"account" json:"account"`
	Password string    `gorm:"column:password;" form:"password" json:"password"`
	Name     string    `gorm:"column:name" form:"name" json:"name"`
	Avatar   string    `gorm:"column:avatar" form:"avatar" json:"avatar"`
	Email    string    `gorm:"column:email" form:"email" json:"email"`
	Phone    string    `gorm:"column:phone" form:"phone" json:"phone"`
	Gender   int       `gorm:"column:gender" form:"gender" json:"gender"`
	Status   string    `gorm:"column:status;default:Y" form:"status" json:"status"`
	CreateAt time.Time `gorm:"column:create_at" form:"createAt" json:"createAt"`
	UpdateAt time.Time `gorm:"column:update_at" form:"updateAt" json:"updateAt"`
}

func (User) TableName() string {
	return "user"
}

type MyAppService struct {
	module.DbModule
	module.RedisModule
}

func (s *MyAppService) TestDb() ([]User, error) {
	var rows []User
	err := s.Db().Model(&User{}).Find(&rows).Error
	return rows, err
}

func (s *MyAppService) TestRedis() (string, error) {
	s.RedisSet("foo", "bar", 60)
	reply, err := s.RedisGet("foo")
	str, err := redis.String(reply, err)
	return str, err
}

var MyAppServiceInstance = new(MyAppService)

type MyAppController struct {
	core.Controller
}

func (c *MyAppController) URLMapping() {
	c.Mapping("/", "GET", c.Test)
	c.Mapping("/db", "GET", c.TestDb)
	c.Mapping("/redis", "GET", c.TestRedis)
}

func (c *MyAppController) Test(ctx *core.Context) core.ControllerResponse {
	return c.MakeResponse("hi")
}

func (c *MyAppController) TestDb(ctx *core.Context) core.ControllerResponse {
	res, err := MyAppServiceInstance.TestDb()
	return c.MakeResponse(res, err)
}

func (c *MyAppController) TestRedis(ctx *core.Context) core.ControllerResponse {
	res, err := MyAppServiceInstance.TestRedis()
	return c.MakeResponse(res, err)
}

func TestApp(t *testing.T) {
	core.ConfigLoad("development")
	app := core.AppNew()
	middleware.UseDefaultHttpMiddleware(app)
	app.InitModule(new(module.DbModule), new(module.RedisModule))
	app.InitModuleUseCustomCaller(core.RegisterController, new(MyAppController))
	t.Log(app)
	app.Start()
}

func TestGrpcApp(t *testing.T) {
	core.ConfigLoad("development")
	app := core.AppNew(core.AppOption{Type: core.APP_TYPE_GRPC})
	server := new(DemoGrpcService)
	app.RegisterGrpcService(&protobuf.BaseService_ServiceDesc, *server)
	t.Log(app)
	app.Start()
}

func TestGrpcAppClient(t *testing.T) {
	conn, err := grpc.Dial(
		"127.0.0.1:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer conn.Close()
	client := protobuf.NewBaseServiceClient(conn)
	grpcCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Request(grpcCtx, &protobuf.RequestBody{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp, err)
}
