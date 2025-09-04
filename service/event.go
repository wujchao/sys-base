package service

import (
	"context"
	v1 "github.com/wujchao/sys-base/api/sys_login/v1"
	"github.com/wujchao/sys-base/internal/model"
)

type (
	ILoginEvent interface {
		Before(ctx context.Context) (user v1.LoginDoReq)
		After(ctx context.Context) (user *model.LoginUserRes, err error)
	}
	ILogoutEvent interface {
		Before(ctx context.Context) (userId string)
		After(ctx context.Context) (userId string, err error)
	}
	IUserCreateEvent interface {
		Before(ctx context.Context) (user *model.UserInput)
		After(ctx context.Context) (user *model.UserOutput, err error)
	}
	IUserDeleteEvent interface {
		Before(ctx context.Context) (userId string)
		After(ctx context.Context) (userId string, err error)
	}
)

var (
	localLoginEvent      ILoginEvent
	localLogoutEvent     ILogoutEvent
	localUserCreateEvent IUserCreateEvent
	localUserDeleteEvent IUserDeleteEvent
)

type DefaultLoginEventImpl struct {
}

func (d *DefaultLoginEventImpl) Before(ctx context.Context) (user v1.LoginDoReq) {

	return
}

func (d *DefaultLoginEventImpl) After(ctx context.Context) (user *model.LoginUserRes, err error) {

	return
}

type DefaultLogoutEventImpl struct {
}

func (d *DefaultLogoutEventImpl) Before(ctx context.Context) (userId string) {

	return
}

func (d *DefaultLogoutEventImpl) After(ctx context.Context) (userId string, err error) {

	return
}

type DefaultUserCreateEventImpl struct {
}

func (d *DefaultUserCreateEventImpl) Before(ctx context.Context) (user *model.UserInput) {

	return
}

func (d *DefaultUserCreateEventImpl) After(ctx context.Context) (user *model.UserOutput, err error) {

	return
}

type DefaultUserDeleteEventImpl struct {
}

func (d *DefaultUserDeleteEventImpl) Before(ctx context.Context) (userId string) {

	return
}

func (d *DefaultUserDeleteEventImpl) After(ctx context.Context) (userId string, err error) {

	return
}

func LoginEvent() ILoginEvent {
	if localLoginEvent == nil {
		localLoginEvent = &DefaultLoginEventImpl{}
	}
	return localLoginEvent
}

func RegisterLoginEvent(i ILoginEvent) {
	localLoginEvent = i
}

func LogoutEvent() ILogoutEvent {
	if localLogoutEvent == nil {
		localLogoutEvent = &DefaultLogoutEventImpl{}
	}
	return localLogoutEvent
}

func RegisterLogoutEvent(i ILogoutEvent) {
	localLogoutEvent = i
}

func UserCreateEvent() IUserCreateEvent {
	if localUserCreateEvent == nil {
		localUserCreateEvent = &DefaultUserCreateEventImpl{}
	}
	return localUserCreateEvent
}

func RegisterUserCreateEvent(i IUserCreateEvent) {
	localUserCreateEvent = i
}

func UserDeleteEvent() IUserDeleteEvent {
	if localUserDeleteEvent == nil {
		localUserDeleteEvent = &DefaultUserDeleteEventImpl{}
	}
	return localUserDeleteEvent
}

func RegisterUserDeleteEvent(i IUserDeleteEvent) {
	localUserDeleteEvent = i
}
