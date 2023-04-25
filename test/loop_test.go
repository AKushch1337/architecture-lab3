package test

import (
	"github.com/AKushch1337/architecture-lab3/painter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/exp/shiny/screen"
	"image"
	"image/color"
	"image/draw"
	"testing"
	"time"
)

type MockReceiver struct {
	mock.Mock
}

func (r *MockReceiver) Update(t screen.Texture) {
	r.Called(t)
}

type MockScreen struct {
	mock.Mock
}

func (s *MockScreen) NewBuffer(size image.Point) (screen.Buffer, error) {
	return nil, nil
}

func (s *MockScreen) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
	return nil, nil
}

func (s *MockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	args := s.Called(size)
	return args.Get(0).(screen.Texture), args.Error(1)
}

type MockTexture struct {
	mock.Mock
}

func (t *MockTexture) Release() {
	t.Called()
}

func (t *MockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	t.Called(dp, src, sr)
}

func (t *MockTexture) Bounds() image.Rectangle {
	args := t.Called()
	return args.Get(0).(image.Rectangle)
}

func (t *MockTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	t.Called(dr, src, op)
}

func (t *MockTexture) Size() image.Point {
	args := t.Called()
	return args.Get(0).(image.Point)
}

func TestLoop(t *testing.T) {
	screenMock := new(MockScreen)
	textureMock := new(MockTexture)
	receiverMock := new(MockReceiver)
	texture := image.Pt(800, 800)
	screenMock.On("NewTexture", texture).Return(textureMock, nil)
	receiverMock.On("Update", textureMock).Return()
	loop := painter.Loop{
		Receiver: receiverMock,
	}

	loop.Start(screenMock)

	op1 := new(MockOperation)
	op2 := new(MockOperation)

	textureMock.On("Bounds").Return(image.Rectangle{})
	op1.On("Do", textureMock).Return(false)
	op2.On("Do", textureMock).Return(true)

	assert.Equal(t, len(loop.MsgQueue.Queue), 0)
	loop.Post(op1)
	loop.Post(op2)
	time.Sleep(1 * time.Second)
	assert.Equal(t, len(loop.MsgQueue.Queue), 0) //Повідомлення запостились і тепер черга знову пуста

	op1.AssertCalled(t, "Do", textureMock)
	op2.AssertCalled(t, "Do", textureMock)
	receiverMock.AssertCalled(t, "Update", textureMock)
	screenMock.AssertCalled(t, "NewTexture", image.Pt(800, 800))
}

type MockOperation struct {
	mock.Mock
}

func (o *MockOperation) Do(t screen.Texture) bool {
	args := o.Called(t)
	return args.Bool(0)
}
