package v1

import (
	"backend/api/v1/info"
	"log"
	"os"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("SbAPI")

type SbAPI struct {
}

func NewSbAPI() *SbAPI {
	logger.SetLogLevel(loggo.INFO)

	api := &SbAPI{}

	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.SetLogger(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds))

	return api
}

func (api *SbAPI) Register(wsContainer *restful.Container, insecure bool) error {
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept", "Access-Control-Allow-Headers"},
		AllowedMethods: []string{"PUT", "POST", "GET", "DELETE", "OPTIONS"},
		AllowedDomains: []string{"*"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	wsContainer.Filter(enableCORS)
	wsContainer.Filter(measureFilter)
	wsContainer.Filter(logFilter)
	//wsContainer.Filter(authFilter)

	info.NewResource().RegisterGet(wsContainer)
	info.NewResource().RegisterPost(wsContainer)
	info.NewResource().RegisterDelete(wsContainer)

	return nil
}

func enableCORS(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.AddHeader(restful.HEADER_AccessControlAllowOrigin, "*")

	chain.ProcessFilter(req, resp)
}

func authFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {

	chain.ProcessFilter(req, resp)
}

func logFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	logger.Infof("HTTP %s %s\n", req.Request.Method, req.Request.URL)
	logger.Infof("HTTP Req Header %s\n", req.Request.Header)
	logger.Infof("HTTP Req Body %s\n", req.Request.Body)

	chain.ProcessFilter(req, resp)
}

func measureFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	start := time.Now()
	chain.ProcessFilter(req, resp)
	logger.Infof("requester  %s %s completed for %v\n", req.Request.Method, req.Request.URL, time.Since(start))
}
