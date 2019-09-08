
#pragma once

#include <lv2/lv2plug.in/ns/lv2core/lv2.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef void LVTK_URI_Directory;

LVTK_URI_Directory* lvtk_uri_directory_new();
uint32_t lvtk_uri_directory_map (LVTK_URI_Directory* directory, const char* uri);
const char* lvtk_uri_directory_unmap (LVTK_URI_Directory* directory, uint32_t urid);
const LV2_Feature* lvtk_uri_directory_get_map_feature(LVTK_URI_Directory*);
const LV2_Feature* lvtk_uri_directory_get_unmap_feature(LVTK_URI_Directory*);
void lvtk_uri_directory_free (LVTK_URI_Directory*);

#ifdef __cplusplus
}
#endif
