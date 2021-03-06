package main

import (
	"github.com/xtwoend/whatsapp-rest/ctl"
	"github.com/xtwoend/whatsapp-rest/hlp/auth"
	"github.com/xtwoend/whatsapp-rest/hlp/router"
)

// Initialize Function in Main Route
func init() {
	// Set Endpoint for Root Functions
	router.Router.Get(router.RouterBasePath, ctl.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", ctl.GetHealth)

	// Set Endpoint for Authorization Functions
	router.Router.With(auth.Basic).Get(router.RouterBasePath+"/auth", ctl.GetAuth)

	// Set Endpoint for WhatsApp Functions
	router.Router.With(auth.JWT).Post("/whatsapp-login", ctl.WhatsAppLogin)
	router.Router.With(auth.JWT).Post("/whatsapp-send-text", ctl.WhatsAppSendText)
	router.Router.With(auth.JWT).Post("/whatsapp-send-image", ctl.WhatsAppSendImage)
	router.Router.With(auth.JWT).Post("/whatsapp-send-doc", ctl.WhatsAppSendDocument)
	router.Router.With(auth.JWT).Post("/whatsapp-logout", ctl.WhatsAppLogout)

	
}
