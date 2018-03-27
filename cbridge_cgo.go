package ubus

/*
#include <libubus.h>
extern int ubusHandlerProxyGO(struct ubus_context *ctx, struct ubus_object *obj, struct ubus_request_data *req, char *method, struct blob_attr *msg);
extern void ubusDataHandlerProxyGO(struct ubus_request *req, int type, struct blob_attr *msg);

int ubus_handler_stub(struct ubus_context *ctx, struct ubus_object *obj, struct ubus_request_data *req, const char *method, struct blob_attr *msg)
{
	return ubusHandlerProxyGO(ctx, obj, req, (char*)method, msg);
}

void ubus_data_handler_stub(struct ubus_request *req, int type, struct blob_attr *msg)
{
	ubusDataHandlerProxyGO(req, type, msg);
}
*/
import "C"
