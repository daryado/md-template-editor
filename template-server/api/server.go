package api

import (
	"github.com/emicklei/go-restful/v3"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func CORSFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.AddHeader(restful.HEADER_AccessControlAllowOrigin, "*")
	chain.ProcessFilter(req, resp)
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 28, //1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	wsContainer := restful.NewContainer()

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"PUT", "POST", "GET", "DELETE"},
		AllowedDomains: []string{"*"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	wsContainer.Filter(wsContainer.OPTIONSFilter)
	wsContainer.Filter(CORSFilter)

	log.Print("start listening on localhost:" + port)
	return s.httpServer.ListenAndServe()
	//return http.ListenAndServe(":"+port, wsContainer)
}
