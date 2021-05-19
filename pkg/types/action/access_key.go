package action

import (
	"github.com/eteu-technologies/borsh-go"
	"github.com/eteu-technologies/near-api-go/pkg/types"
)

type AccessKeyPermission struct {
	Enum borsh.Enum `borsh_enum:"true"`

	FunctionCallPermission AccessKeyFunctionCallPermission
	FullAccessPermission   struct{}
}

func NewFunctionCallPermission(allowance types.Balance, receiverID types.AccountID, methodNames []string) AccessKeyPermission {
	return AccessKeyPermission{
		Enum: borsh.Enum(0),
		FunctionCallPermission: AccessKeyFunctionCallPermission{
			Allowance:   &allowance,
			ReceiverID:  receiverID,
			MethodNames: methodNames,
		},
	}
}

func NewFunctionCallUnlimitedAllowancePermission(receiverID types.AccountID, methodNames []string) AccessKeyPermission {
	return AccessKeyPermission{
		Enum: borsh.Enum(0),
		FunctionCallPermission: AccessKeyFunctionCallPermission{
			Allowance:   nil,
			ReceiverID:  receiverID,
			MethodNames: methodNames,
		},
	}
}

func NewFullAccessPermission() AccessKeyPermission {
	return AccessKeyPermission{
		Enum: borsh.Enum(0),
	}
}

type AccessKeyFunctionCallPermission struct {
	Allowance   *types.Balance
	ReceiverID  types.AccountID
	MethodNames []string
}
