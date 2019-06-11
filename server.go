package main

import (
	"net/http"

	//"github.com/byuoitav/adcp-control-microservice/handlers"
	"./handlers"
	"github.com/byuoitav/common"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/v2/auth"
	"github.com/byuoitav/hateoas"
	"github.com/labstack/echo"
)

func main() {
	port := ":8099"
	router := common.NewRouter()
	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))

	write := router.Group("", auth.AuthorizeRequest("write-state", "room", auth.LookupResourceFromAddress))
	read := router.Group("", auth.AuthorizeRequest("read-state", "room", auth.LookupResourceFromAddress))

	write.GET("/:address/volume/set/:level", handlers.SetVolume)
	write.GET("/:address/power/on", handlers.PowerOn)
	write.GET("/:address/power/standby", handlers.PowerStandby)
	write.GET("/:address/volume/mute", handlers.Mute)
	write.GET("/:address/volume/unmute", handlers.UnMute)
	write.GET("/:address/display/blank", handlers.DisplayBlank)
	write.GET("/:address/display/unblank", handlers.DisplayUnBlank)
	write.GET("/:address/input/:port", handlers.SetInputPort)

	//status endpoints
	read.GET("/:address/volume/level", handlers.VolumeLevel)
	read.GET("/:address/volume/mute/status", handlers.MuteStatus)
	read.GET("/:address/power/status", handlers.PowerStatus)
	read.GET("/:address/display/status", handlers.BlankedStatus)
	read.GET("/:address/input/current", handlers.InputStatus)
	read.GET("/:address/active/:port", handlers.HasActiveSignal)
	read.GET("/:address/hardware", handlers.GetHardwareInfo)

	//------------------
	//Pooled endpoints
	//------------------
	write.GET("/pooled/:address/volume/set/:level", handlers.SetVolumePooled)
	write.GET("/pooled/:address/power/on", handlers.PowerOnPooled)
	write.GET("/pooled/:address/power/standby", handlers.PowerStandbyPooled)
	write.GET("/pooled/:address/volume/mute", handlers.MutePooled)
	write.GET("/pooled/:address/volume/unmute", handlers.UnMutePooled)
	write.GET("/pooled/:address/display/blank", handlers.DisplayBlankPooled)
	write.GET("/pooled/:address/display/unblank", handlers.DisplayUnBlankPooled)
	write.GET("/pooled/:address/input/:port", handlers.SetInputPortPooled)

	//status endpoints
	read.GET("/pooled/:address/volume/level", handlers.VolumeLevelPooled)
	read.GET("/pooled/:address/volume/mute/status", handlers.MuteStatusPooled)
	read.GET("/pooled/:address/power/status", handlers.PowerStatusPooled)
	read.GET("/pooled/:address/display/status", handlers.BlankedStatusPooled)
	read.GET("/pooled/:address/input/current", handlers.InputStatusPooled)
	read.GET("/pooled/:address/active/:port", handlers.HasActiveSignalPooled)
	read.GET("/pooled/:address/hardware", handlers.GetHardwareInfoPooled)

	router.PUT("/log-level/:level", log.SetLogLevel)
	router.GET("/log-level", log.GetLogLevel)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
