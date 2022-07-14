package bob

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

type Vbo struct {
	id     uint32
	target uint32
	size   int32
}

func NewVbo(target uint32) *Vbo {
	vbo := &Vbo{target: target}
	gl.GenBuffers(1, &vbo.id)
	return vbo
}

func (v *Vbo) Drop() {
	gl.DeleteBuffers(1, &v.id)
}

func (v *Vbo) Bind() {
	gl.BindBuffer(v.target, v.id)
}

func (v *Vbo) Unbind() {
	gl.BindBuffer(v.target, 0)
}

func (v *Vbo) Fill(data unsafe.Pointer, elements, size int, use uint32) {
	v.size = int32(size)

	v.Bind()
	gl.BufferData(v.target, elements*size, data, use)
	v.Unbind()
}

func (v *Vbo) Update(data unsafe.Pointer, elements, offset, size int) {
	v.size = int32(size)

	v.Bind()
	gl.BufferSubData(v.target, offset, elements*size, data)
	v.Unbind()
}

func (v *Vbo) AttribPointer(attribLocation int32, size int32, btype uint32, normalized bool, stride int32) {
	gl.VertexAttribPointer(uint32(attribLocation), size, btype, normalized, stride, nil)
}

func (v *Vbo) Id() uint32 {
	return v.id
}

func (v *Vbo) Target() uint32 {
	return v.target
}

func (v *Vbo) Size() int32 {
	return v.size
}
