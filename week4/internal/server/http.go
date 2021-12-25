package server

import (
	"github.com/cyl615123/geek/week4/internal/biz"
	"github.com/cyl615123/geek/week4/internal/config"
	"github.com/cyl615123/geek/week4/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"net/http"
)

func NewHTTPServer(c *config.Conf, library *service.LibraryService) *http.Server {
	r := gin.Default()
	r.GET("/library/v1/series/book", FindSeriesBool(library))
	return &http.Server{Handler: r}
}

type SeriesReq struct {
	Names []string `form:"names"`
}

type SeriesReply struct {
	Books []*biz.Book `json:"books"`
}

func FindSeriesBool(library *service.LibraryService) gin.HandlerFunc {
	return func(g *gin.Context) {
		req := SeriesReq{}
		if err := g.ShouldBind(&req); err != nil {
			g.JSON(http.StatusBadRequest, nil)
			return
		}
		books, err := library.FindSeriesBool(req.Names)
		if err != nil {
			log.Error("FindSeriesBool err=%+v, err")
			g.JSON(http.StatusInternalServerError, nil)
			return
		}
		g.JSON(http.StatusOK, &SeriesReply{
			Books: books,
		})
	}
}
