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
		0, 0.01, 0, // top
		-0.01, -0.01, 0, // left
		0.01, -0.01, 0, // right
	}
)

type object struct {
	drawable uint32

	body *gravity.Body
}

func newObject(b *gravity.Body, farthest float64, massiest float64) *object {
	points := make([]float32, len(triangle))
	copy(points, triangle)

	//Adjust size
	points[0] *= float32(b.Mass/farthest*massiest)
	points[1] *= float32(b.Mass/farthest*massiest)
	points[2] *= float32(b.Mass/farthest*massiest)

	points[3] *= float32(b.Mass/farthest*massiest)
	points[4] *= float32(b.Mass/farthest*massiest)
	points[5] *= float32(b.Mass/farthest*massiest)

	points[6] *= float32(b.Mass/farthest*massiest)
	points[7] *= float32(b.Mass/farthest*massiest)
	points[8] *= float32(b.Mass/farthest*massiest)

			//Adjust position
			points[0] += float32(b.XPos/farthest*.75)
			points[1] += float32(b.YPos/farthest*.75)
			points[2] += float32(b.ZPos/farthest*.75)

			points[3] += float32(b.XPos/farthest*.75)
			points[4] += float32(b.YPos/farthest*.75)
			points[5] += float32(b.ZPos/farthest*.75)
		       
			points[6] += float32(b.XPos/farthest*.75)
			points[7] += float32(b.YPos/farthest*.75)
			points[8] += float32(b.ZPos/farthest*.75)

	return &object{
		drawable: makeVao(points),
		body: b,
	}
}

func (o *object) draw() {
	gl.BindVertexArray(o.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))
}

func (o *object) delete() {
	gl.DeleteVertexArrays(1,&o.drawable)
	gl.DeleteBuffers(1,&o.drawable)
}

func main() {

	runtime.LockOSThread()

	fmt.Println("go-gravity")
	b1 := gravity.Body{
		XPos: 0,
		YPos: 0,
		ZPos: 0,
		XVel: -.001,
		YVel: 0,
		ZVel: 0,
		Mass: 10,
	}
	b2 := gravity.Body{
		XPos: 0,
		YPos: -10,
		ZPos: 0,
		XVel: .1,
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
		var objects []*object
		farthest := u.FarthestPointFromOrigin()
		massiest := u.LargestMass()
		for i := range u.Bodies {
			objects = append(objects, newObject(&u.Bodies[i],farthest,massiest))
		}
		draw(objects, window, program)
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

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

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

func draw(objects []*object, window *glfw.Window, program uint32) {

	//gl.ClearColor(0, 0.5, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for _, object := range objects {
		object.draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()

	for _, object := range objects {
		object.delete()
	}
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
