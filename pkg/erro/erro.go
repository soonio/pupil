package erro

import "runtime"

type Erro struct {
	text string // 错误消息
	fn   string // 错误出现的位置
}

var _ error = (*Erro)(nil)

func (t *Erro) Error() string {
	return t.text
}

func (t *Erro) On() string {
	return t.fn
}

func New(text string) error {
	e := &Erro{text: text}

	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	e.fn = f.Name()

	return e
}
