package main

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
)

func buildWindow() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("couldn't initialize glfw: %v", err))
	}

	setWindowHints()

	window, err := glfw.CreateWindow(800, 600, "Graphics", nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create opengl renderer: %v", err))
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	program := gl.CreateProgram()
	gl.LinkProgram(program)
	return program
}

func setWindowHints() {
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	//glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)   Only works on <= OpenGL 3.2
	//glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)      Only works on <= OpenGL 3.0
}
