
#ifndef LVTK_GO_URI_DIRECTORY
#define LVTK_GO_URI_DIRECTORY

#include <lv2/lv2plug.in/ns/lv2core/lv2.h>
#include <lv2/lv2plug.in/ns/ext/urid/urid.h>
#include <stdint.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef void LvtkURIDirectory;

LvtkURIDirectory* lvtk_uri_directory_new();
uint32_t lvtk_uri_directory_map (LvtkURIDirectory*, const char*);
const char* lvtk_uri_directory_unmap (LvtkURIDirectory*, uint32_t);
const LV2_Feature* lvtk_uri_directory_get_map_feature(LvtkURIDirectory*);
LV2_URID_Map* lvtk_uri_directory_get_map(LvtkURIDirectory*);
const LV2_Feature* lvtk_uri_directory_get_unmap_feature(LvtkURIDirectory*);
LV2_URID_Unmap* lvtk_uri_directory_get_unmap(LvtkURIDirectory*);
void lvtk_uri_directory_free (LvtkURIDirectory*);

#ifdef __cplusplus
}
#endif

#endif
