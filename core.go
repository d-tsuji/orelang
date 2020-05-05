package orelang

type Engine struct {
	operators map[interface{}]IOperator
	variables map[interface{}]interface{}
}

func NewEngine() Engine {
	operators := make(map[interface{}]IOperator)
	variables := make(map[interface{}]interface{})

	operators["+"] = AddOperator{}
	operators["*"] = MultiplyOperator{}
	operators["="] = EqualOperator{}
	operators["set"] = SetOperator{}
	operators["get"] = GetOperator{}
	operators["until"] = UntilOperator{}
	operators["step"] = StepOperator{}

	return Engine{operators, variables}
}

func (en Engine) getExpression(script interface{}) IExpression {
	v, ok := script.([]interface{})
	if ok {
		return &CallOperator{
			operator: en.operators[v[0]],
			args:     v[1:],
		}
	}
	return ImmediateValue{value: script}
}

func (en Engine) Eval(script interface{}) interface{} {
	// TODO: It's very Java-like...
	return en.getExpression(script).eval(&en)
}

type IExpression interface {
	eval(engine *Engine) interface{}
}

type CallOperator struct {
	operator IOperator
	args     []interface{}
}

func (c *CallOperator) eval(engine *Engine) interface{} {
	return c.operator.call(engine, c.args)
}

type ImmediateValue struct {
	value interface{}
}

func (im ImmediateValue) eval(engine *Engine) interface{} {
	return im.value
}
