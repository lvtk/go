#include <stdlib.h>
#include <stdio.h>
#include <memory.h>
#include <lv2/lv2plug.in/ns/extensions/ui/ui.h>
#include "goui.h"

extern int lvtk_go_ui_instantiate (GoUI*, void* p0, void* p1, void* p2, void* p3, void* p4);
extern void lvtk_go_ui_port_event (GoUI* p0);
extern void lvtk_go_ui_cleanup (GoUI* p0);
extern LV2UI_Widget lvtk_go_ui_get_widget (GoUI* p0);
extern void* lvtk_go_ui_extension_data (void* p0);
extern char* lvtk_go_ui_get_uri();
extern int lvtk_go_ui_idle(GoUI*);

static LV2UI_Handle instantiate (const struct _LV2UI_Descriptor* descriptor,
                                 const char*                     plugin_uri,
                                 const char*                     bundle_path,
                                 LV2UI_Write_Function            write_function,
                                 LV2UI_Controller                controller,
                                 LV2UI_Widget*                   widget,
                                 const LV2_Feature* const*       features)
{
    GoUI* ui = (GoUI*) malloc (sizeof(GoUI));
    if (0 != lvtk_go_ui_instantiate (ui, (void*)plugin_uri, (void*)bundle_path, 
                                     write_function, controller, (void*)features))
    {
        free (ui);
        return NULL;
    }

    *widget = lvtk_go_ui_get_widget (ui);
    return ui;
}

static void cleanup (LV2UI_Handle ui) {
    lvtk_go_ui_cleanup (ui);
    free (ui);
}

static void port_event (LV2UI_Handle ui,
                        uint32_t     port_index,
                        uint32_t     buffer_size,
                        uint32_t     format,
                        const void*  buffer)
{
    lvtk_go_ui_port_event (ui);
}

static int idle (LV2UI_Handle handle) {
    return lvtk_go_ui_idle (handle);
}

static const void* extension_data (const char* uri) {
    const void* data = lvtk_go_ui_extension_data ((void*) uri);
    if (data != NULL)
        return data;

    printf("GoUI: internal extension data: %s\n", uri);
    if (0 == strcmp(uri, LV2_UI__idleInterface)) {
        printf("GoUI: wants idle\n");
        static const LV2UI_Idle_Interface idle_interface = { idle };
        return &idle_interface;
    }

    return NULL;
}

#define GOUI_URI_MAX_LENGTH 256
static char goui_uri[GOUI_URI_MAX_LENGTH] = { '\0' };
static const LV2UI_Descriptor _descriptor = {
    NULL,
    instantiate,
    cleanup,
    port_event,
    extension_data
};

LV2_SYMBOL_EXPORT
const LV2UI_Descriptor* lv2ui_descriptor (uint32_t index) {
    if (_descriptor.URI == NULL) {
        memset(goui_uri, 0, GOUI_URI_MAX_LENGTH);
        char* uri = lvtk_go_ui_get_uri();
        size_t len = strlen (uri);
        if (len > GOUI_URI_MAX_LENGTH)
            len = GOUI_URI_MAX_LENGTH;
        memcpy(goui_uri, uri, len);
        ((LV2UI_Descriptor*)&_descriptor)->URI = goui_uri;
        free(uri);
    }

    return index == 0 ? &_descriptor : NULL;
}
