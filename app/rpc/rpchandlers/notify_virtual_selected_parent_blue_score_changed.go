package rpchandlers

import (
	"github.com/ixbasANT/gord/app/appmessage"
	"github.com/ixbasANT/gord/app/rpc/rpccontext"
	"github.com/ixbasANT/gord/infrastructure/network/netadapter/router"
)

// HandleNotifyVirtualSelectedParentBlueScoreChanged handles the respectively named RPC command
func HandleNotifyVirtualSelectedParentBlueScoreChanged(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateVirtualSelectedParentBlueScoreChangedNotifications()

	response := appmessage.NewNotifyVirtualSelectedParentBlueScoreChangedResponseMessage()
	return response, nil
}
