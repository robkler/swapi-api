package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"swapi/db"
	"swapi/routes"
	request_swapi "swapi/swapi"
	"time"
)

type Server struct{}

func (s *Server) New() *gin.Engine {
	planetDb := db.PlanetDb{}
	mapPlanets := request_swapi.MapPlanets{}
	mapPlanets.GetAllPlanets()
	planetRoutes := routes.PlanetRoutes{
		PlanetDb: &planetDb,
		Swapi:    &mapPlanets,
	}
	gin.SetMode(gin.ReleaseMode)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	r := gin.New()
	r.Use(s.loggerMiddleware)
	r.Use(s.logger())
	r.POST("/planet", planetRoutes.InsertPlanet)
	r.GET("/planet", planetRoutes.GetPlanets)
	r.GET("/planet/{user_name}/name", planetRoutes.GetByName)
	r.GET("/planet/{user_uuid}/id", planetRoutes.GetById)
	r.DELETE("/planet/{user_uuid}", planetRoutes.DeletePlanet)
	return r
}

func (s *Server) loggerMiddleware(c *gin.Context) {
	logrus.Info("%s %s", c.Request.Method, c.FullPath())
	startTime := time.Now()
	c.Next()
	logrus.WithField("Status", c.Writer.Status()).WithField("Time Took", time.Since(startTime)).Info("%s %s", c.Request.Method, c.FullPath())

}
func (s *Server) logger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Output: nil,
		Formatter: func(params gin.LogFormatterParams) string {
			fields := logrus.Fields{
				HTTPPath:       params.Path,
				HTTPMethod:     params.Method,
				HTTPStatusCode: params.StatusCode,
				HTTPBodySize:   params.BodySize,
				HTTPClientIP:   params.ClientIP,
				HTTPLatency:    params.Latency.Milliseconds(),
				HTTPUrl:        params.Request.URL.String(),
			}
			for key, val := range params.Keys {
				fields[key] = val
			}
			if params.ErrorMessage != "" {
				fields["error"] = true
				logrus.WithFields(fields).Error(params.ErrorMessage)
			}
			logrus.WithFields(fields).Info("")
			return ""
		},
	})
}
