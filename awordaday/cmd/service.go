package main

import (
	"awordaday/channels"
	"awordaday/database"
	"awordaday/models"
	"fmt"

	"awordaday/handler"
	"awordaday/helper"

	"bytes"

	"io/ioutil"
	"time"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	nats "github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

const (
	port              = ":50051"
	connectionRetries = 5
)

var (
	provider              string
	dbConnectionStr       string
	dbName                string
	natsConnection        string
	err                   error
	mysupersecretpassword = "TheAimIsToUseThisJWT"

	nc *nats.Conn
)

func init() {
	viper.SetConfigName("app")     // no need to include file extension
	viper.AddConfigPath("config/") // set the path of your config file

	err := viper.ReadInConfig()

	if err != nil {
		glog.Fatal(err)
	} else {
		/*if os.Getenv("RUNTIME")=="CLOUD"{
		dbConnectionStr = viper.GetString("connections.cloud.DBConnection")
		dbName = viper.GetString("connections.cloud.DBName")
		natsConnection = viper.GetString("connections.cloud.NATSURL")
		}*/
		provider = viper.GetString("connections.Provider")

		dbConnectionStr = viper.GetString("connections.DBConnection")
		dbName = viper.GetString("connections.DBName")
		natsConnection = viper.GetString("connections.NATSURL")
	}
}

func main() {

	//Connect to the database
	glog.Info(provider, dbConnectionStr)
	session, err := database.New(provider, dbConnectionStr)
	glog.Info(dbConnectionStr)
	defer glog.Flush()
	if err != nil {
		if session != nil {
			session.Client.Close()
		}
		glog.Info("no connection to the database", err.Error())
	}
	if session != nil {
		defer session.Client.Close()
	}

	// Connect to Casbin Authorization
	e, err := casbin.NewEnforcerSafe("auth/auth_model.conf", "auth/policy.csv")

	if err != nil {
		glog.Error("Authorization engine not started")
	}

	func() {
		retries := 0
	try:
		retries++
		nc, err = nats.Connect(natsConnection)
		if err != nil {
			if retries < connectionRetries {
				time.Sleep(5000)
				glog.Info("Trying to connect ---", retries)
				goto try
			}
			glog.Info("Not connected to Nats.. hence the application cannot be started.")
		}
		glog.Info(natsConnection)
	}()

	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	//router.Use(cors.Default())
	router.Use(CORSMiddleware())
	// create user handler instance
	handler.Init(nc)
	channels.InitAudit(session, dbName)

	router.Use(AuditMiddleware())

	//Authorization Middleware
	router.Use(Authorization(e))

	wordGroup := router.Group("/v1/word")
	{
		//wordGroup.Use(jwt.Auth(mysupersecretpassword))
		wordGroup.GET("/getMagicWord", handler.GetMagicWord(session))
		wordGroup.GET("/getAll", handler.GetAllWords(session))
		wordGroup.GET("/get/:skip/:limit", handler.GetWords(session))
		wordGroup.POST("/insert", handler.InsertWord(session))
		wordGroup.POST("/sentence/insert", handler.InsertSentence(session))
		wordGroup.POST("/request", handler.InsertRequestedWord(session))
		wordGroup.DELETE("/:word", handler.DeleteWord(session))
		wordGroup.PUT("/update/:id", handler.UpdateWord(session))

	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(port)
}

//CORSMiddleware a simple middle ware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// Authorization is a middleware
func Authorization(e *casbin.Enforcer) func(*gin.Context) {
	return func(c *gin.Context) {
		role := c.Request.Header.Get("role")
		if role == "" {
			role = "anonymous"
			//respondWithError(c,401,"failed","authorization subject has not provided")
			//return
		}
		if e == nil {
			respondWithError(c, 401, "failed", "authorization has not been set-1")
			return
		}

		// casbin enforce
		res, err := e.EnforceSafe(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			respondWithError(c, 401, "failed", "authorization has not been set-2")
			return
		}
		if res {
			c.Next()
		} else {
			respondWithError(c, 401, "failed", "authorization has not been set-3")
			return
		}
		// Have to use Casbin here..
		// Pass on to the next-in-chain
		c.Next()
	}
}

// AuditMiddleware to audit and find details of hits
func AuditMiddleware() func(*gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "GET" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" || c.Request.Method == "PATCH" {
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			//var data interface{}

			//json.Unmarshal(bodyBytes, &bodyBytes)
			var headers string
			for key, val := range c.Request.Header {
				// Convert each key/value pair in m to a string
				headers = fmt.Sprintf("%s=\"%s\"", key, val)
				// Do whatever you want to do with the string;
				// in this example I just print out each of them.
				//fmt.Println(s)
			}
			ip, _ := helper.GetClientIPHelper(c.Request)
			channels.ChanAudit <- models.Audit{Data: string(bodyBytes), Headers: headers, URLPath: c.Request.Host + helper.GetPath(c), IP: ip, Device: c.Request.Header.Get("User-Agent"), DateTime: time.Now().UTC()}
		}
		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, status string, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"message": message,
		"status":  status,
	})
}
