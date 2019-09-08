
#ifndef LVTK_GO_URI_DIRECTORY
#define LVTK_GO_URI_DIRECTORY

#include <lv2/lv2plug.in/ns/lv2core/lv2.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef void LvtkURIDirectory;

LvtkURIDirectory* lvtk_uri_directory_new();
uint32_t lvtk_uri_directory_map (LvtkURIDirectory*, const char*);
const char* lvtk_uri_directory_unmap (LvtkURIDirectory*, uint32_t);
const LV2_Feature* lvtk_uri_directory_get_map_feature(LvtkURIDirectory*);
const LV2_Feature* lvtk_uri_directory_get_unmap_feature(LvtkURIDirectory*);
void lvtk_uri_directory_free (LvtkURIDirectory*);

#ifdef __cplusplus
}
#endif

#endif
