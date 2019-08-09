package minheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func toEl(i int) interface{} {
	return &i
}

func fromEl(i interface{}) int {
	return *(i.(*int))
}

func TestPeek(t *testing.T) {
	mh := NewMinHeap()

	el := &Element{
		Value:    toEl(1),
		Priority: 5,
	}

	mh.PushEl(el)
	assert.Equal(t, 1, fromEl(mh.PeekEl().Value))
	assert.Equal(t, 1, mh.Len())

	el = &Element{
		Value:    toEl(2),
		Priority: 1,
	}
	mh.PushEl(el)
	assert.Equal(t, 2, mh.Len())
	assert.Equal(t, 2, fromEl(mh.PeekEl().Value))
	assert.Equal(t, 2, fromEl(mh.PeekEl().Value))
	assert.Equal(t, 2, mh.Len())

	el = mh.PopEl()

	assert.Equal(t, 2, fromEl(el.Value))
	assert.Equal(t, 1, mh.Len())
	assert.Equal(t, 1, fromEl(mh.PeekEl().Value))

	mh.PopEl()
	assert.Equal(t, 0, mh.Len())
}

func TestUpdate(t *testing.T) {
	mh := NewMinHeap()
	x := &Element{
		Value:    toEl(1),
		Priority: 4,
	}
	y := &Element{
		Value:    toEl(2),
		Priority: 3,
	}
	z := &Element{
		Value:    toEl(3),
		Priority: 8,
	}
	mh.PushEl(x)
	mh.PushEl(y)
	mh.PushEl(z)
	assert.Equal(t, 2, fromEl(mh.PeekEl().Value))

	mh.UpdateEl(z, 1)
	assert.Equal(t, 3, fromEl(mh.PeekEl().Value))

	mh.UpdateEl(x, 0)
	assert.Equal(t, 1, fromEl(mh.PeekEl().Value))
}
