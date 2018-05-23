package main

import (
	"fmt"
	"runtime"
	"strings"
)

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/jotingen/go-gravity/gravity"
)

const (
	windowWidth        = 960
	windowHeight       = 540
	vertexShaderSource = `
    #version 410
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

	fragmentShaderSource = `
    #version 410
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(1, 1, 1, 1);
    }
` + "\x00"
)

var (
	triangle = []float32{
		0, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
	}
)

func main() {

	runtime.LockOSThread()

	fmt.Println("go-gravity")
	b1 := gravity.Body{
		XPos: 0,
		YPos: 0,
		ZPos: 0,
		XVel: 0,
		YVel: 0,
		ZVel: 0,
		Mass: 10000,
	}
	b2 := gravity.Body{
		XPos: 0,
		YPos: 1,
		ZPos: 0,
		XVel: .01,
		YVel: 0,
		ZVel: 0,
		Mass: 1,
	}
	u := gravity.Universe{}
	u.Bodies = append(u.Bodies, b1)
	u.Bodies = append(u.Bodies, b2)

	window := initGlfw()
	defer glfw.Terminate()

	program := initOpenGL()

	for !window.ShouldClose() {
		fmt.Printf("%+v\n", u)
		u.Step()
		triangle = []float32{
			float32(u.Bodies[0].XPos) + 0.0, float32(u.Bodies[0].YPos) + 0.1, float32(u.Bodies[0].ZPos) + 0.0, //top
			float32(u.Bodies[0].XPos) - 0.1, float32(u.Bodies[0].YPos) - 0.1, float32(u.Bodies[0].ZPos) + 0.0, //left
			float32(u.Bodies[0].XPos) + 0.1, float32(u.Bodies[0].YPos) - 0.1, float32(u.Bodies[0].ZPos) + 0.0, //right
		}
		vao := makeVao(triangle)
		draw(vao, window, program)
	}

}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("could not initialize glfw: %v", err))
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(800, 600, "Hello world", nil, nil)
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

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)

	return prog
}

func draw(vao uint32, window *glfw.Window, program uint32) {

	//gl.ClearColor(0, 0.5, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()

}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
