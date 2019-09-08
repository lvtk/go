
#include "directory.h"
#include <lvtk/uri_directory.hpp>

extern "C" {

LvtkURIDirectory* lvtk_uri_directory_new() {
    lvtk::URIDirectory* d = new lvtk::URIDirectory();
    return (LvtkURIDirectory*) d;
}

void lvtk_uri_directory_free (LvtkURIDirectory* d) {
    delete static_cast<lvtk::URIDirectory*> (d);
}

uint32_t lvtk_uri_directory_map (LvtkURIDirectory* d, const char* uri) {
    return (static_cast<lvtk::URIDirectory*>(d))->map (uri);
}

const char* lvtk_uri_directory_unmap (LvtkURIDirectory* d, uint32_t urid) {
    return (static_cast<lvtk::URIDirectory*>(d))->unmap (urid);
}

const LV2_Feature* lvtk_uri_directory_get_map_feature(LvtkURIDirectory* d) {
    return (static_cast<lvtk::URIDirectory*>(d))->get_map_feature();
}

LV2_URID_Map* lvtk_uri_directory_get_map(LvtkURIDirectory* d) {
    return (LV2_URID_Map*) lvtk_uri_directory_get_map_feature(d)->data;
}

const LV2_Feature* lvtk_uri_directory_get_unmap_feature(LvtkURIDirectory* d) {
    return (static_cast<lvtk::URIDirectory*>(d))->get_unmap_feature();
}

LV2_URID_Unmap* lvtk_uri_directory_get_unmap(LvtkURIDirectory*);
    return (LV2_URID_Unmap*) lvtk_uri_directory_get_unmap_feature(d)->data;
}
