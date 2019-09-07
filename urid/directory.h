
#pragma once

#include <lv2/lv2plug.in/ns/lv2core/lv2.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef void LVTKURIDirectory;

LVTKURIDirectory* lvtk_uri_directory_new();
uint32_t lvtk_uri_directory_map (LVTKURIDirectory* directory, const char* uri);
const char* lvtk_uri_directory_unmap (LVTKURIDirectory* directory, uint32_t urid);
const LV2_Feature* lvtk_uri_directory_get_map_feature(LVTKURIDirectory*);
const LV2_Feature* lvtk_uri_directory_get_unmap_feature(LVTKURIDirectory*);
void lvtk_uri_directory_free (LVTKURIDirectory*);

#ifdef __cplusplus
}
#endif
