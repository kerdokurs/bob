package bob

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Vao struct {
	id uint32
}

func NewVao() *Vao {
	vao := &Vao{}
	gl.GenVertexArrays(1, &vao.id)
	return vao
}

func (v *Vao) Drop() {
	gl.DeleteVertexArrays(1, &v.id)
}

func (v *Vao) Bind() {
	gl.BindVertexArray(v.id)
}

func (v *Vao) Unbind() {
	gl.BindVertexArray(0)
}

func (v *Vao) Id() uint32 {
	return v.id
}
