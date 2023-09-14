package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	fmt.Println("Hello World")

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
