package bob

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"math"
	"runtime"
	"time"
)

const (
	DefaultWidth  = 800
	DefaultHeight = 600
	DefaultTitle  = "Bob"
)

var (
	clearColor mgl32.Vec4
)

type AppOptions struct {
	Title  string
	Width  uint32
	Height uint32

	Resizable bool

	IsFullscreen bool
	MonitorId    uint32

	ClearColor mgl32.Vec4
}

type App interface {
	Setup()
	Update(float64)
	IsRunning() bool

	SetWindow(*glfw.Window)
}

func init() {
	runtime.LockOSThread()
}

func RunApp(app App, options *AppOptions) {
	log.Println("Initializing OpenGL")
	if err := gl.Init(); err != nil {
		log.Fatalf("Error initializing OpenGL: %s\n", err.Error())
	}

	log.Print("Initializing GLFW")
	if err := glfw.Init(); err != nil {
		log.Fatalf("Error initializing GLFW: %s\n", err.Error())
	}
	defer glfw.Terminate()

	width := DefaultWidth
	height := DefaultHeight
	title := DefaultTitle

	var monitor *glfw.Monitor

	if options != nil {
		if options.Width > 0 {
			width = int(options.Width)
		}

		if options.Height > 0 {
			height = int(options.Height)
		}

		if len(options.Title) > 0 {
			title = options.Title
		}

		if !options.Resizable {
			glfw.WindowHint(glfw.Resizable, glfw.False)
		}

		if options.IsFullscreen {
			monitors := glfw.GetMonitors()

			if options.MonitorId < uint32(len(monitors)) {
				monitor = monitors[options.MonitorId]
			}
		}

		clearColor = options.ClearColor
	}

	if monitor != nil {
		videoMode := monitor.GetVideoMode()
		width = videoMode.Width
		height = videoMode.Height
	}

	window, err := glfw.CreateWindow(width, height, title, monitor, nil)
	if err != nil {
		log.Fatalf("Could not create window: %s\n", err.Error())
	}

	app.SetWindow(window)

	window.MakeContextCurrent()

	SetViewport(0, 0, int32(width), int32(height))

	app.Setup()

	delta := time.Duration(0)
	var deltaSec float64

	for {
		if !app.IsRunning() || window.ShouldClose() {
			return
		}

		start := time.Now()
		gl.ClearColor(clearColor.X(), clearColor.Y(), clearColor.Z(), clearColor.W())
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		if !math.IsInf(deltaSec, 0) && !math.IsInf(deltaSec, -1) {
			app.Update(deltaSec)
		}

		err := gl.GetError()
		if err != 0 {
			log.Fatalf("OpenGL error: %d\n", err)
		}

		window.SwapBuffers()
		delta = time.Since(start)
		deltaSec = delta.Seconds()
		glfw.PollEvents()
	}

	//goland:noinspection GoUnreachableCode
	glfw.Terminate()
}

func SetViewport(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}
