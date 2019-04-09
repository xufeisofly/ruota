package ruota

import (
	"errors"
	"fmt"
)

type RClient struct {
	Trans      RTransport
	Serializer RSerializer
}

type RClientArg struct {
	Name string
}

type RClientResult struct {
	Result []string
}

func NewRClient(trans RTransport, serializer RSerializer) (RClient, error) {
	return RClient{
		Trans:      trans,
		Serializer: serializer,
	}, nil
}

// IDL 生成的代码
func (p *RClient) FunCall(name string) ([]string, error) {
	var arg RClientArg
	var result RClientResult

	arg.Name = name

	p.Call("FunCall", &arg, &result)
	return result.Result, nil
}

func (p *RClient) Call(funName string, arg *RClientArg, result *RClientResult) error {
	if err := p.Send(funName, arg.Name); err != nil {
		return err
	}
	result.Result, _ = p.Recv(funName)
	return nil
}

func (p *RClient) Send(funName string, arg string) error {
	sFunName, _ := p.Serializer.SerializeString(funName)
	if err := p.Trans.WriteFunName(sFunName); err != nil {
		return err
	}
	sArg, _ := p.Serializer.SerializeString(arg)
	err := p.Trans.WriteArg(sArg)
	p.Trans.Flush()
	return err
}

func (p *RClient) Recv(funName string) ([]string, error) {
	// TODO should be a loop, 参考 thrift demo Result Read
	// 接收函数名
	var dFunName string
	for {
		sFunName, err := p.Trans.ReadFunName()
		if err != nil {
			return []string{}, err
		}
		dFunName, err := p.Serializer.DeserializeString(sFunName)
		if err != nil {
			return []string{}, err
		}
		// fmt.Println(dFunName)

		if dFunName == "STOP" {
			break
		}
	}

	// 验证函数名是否正确
	if dFunName != funName {
		return []string{}, errors.New("Method not same")
	}
	fmt.Printf("Valid Result of FunName: %s", dFunName)

	// 接受结果（数组）
	sResult, _, err := p.Trans.ReadList()
	if err != nil {
		return []string{}, err
	}
	dResult, err := p.Serializer.DeserializeList(sResult, STRING)
	var ret []string
	for _, v := range dResult {
		ret = append(ret, v.(string))
	}
	return ret, err
}
