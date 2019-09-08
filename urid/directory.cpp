
#include "directory.h"
#include <lvtk/uri_directory.hpp>

extern "C" {

LVTK_URI_Directory* lvtk_uri_directory_new() {
    lvtk::URIDirectory* d = new lvtk::URIDirectory();
    return (LVTK_URI_Directory*) d;
}

void lvtk_uri_directory_free (LVTK_URI_Directory* d) {
    delete static_cast<lvtk::URIDirectory*> (d);
}

uint32_t lvtk_uri_directory_map (LVTK_URI_Directory* d, const char* uri) {
    return (static_cast<lvtk::URIDirectory*>(d))->map (uri);
}

const char* lvtk_uri_directory_unmap (LVTK_URI_Directory* d, uint32_t urid) {
    return (static_cast<lvtk::URIDirectory*>(d))->unmap (urid);
}

const LV2_Feature* lvtk_uri_directory_get_map_feature(LVTK_URI_Directory* d) {
    return (static_cast<lvtk::URIDirectory*>(d))->get_map_feature();
}

const LV2_Feature* lvtk_uri_directory_get_unmap_feature(LVTK_URI_Directory* d) {
    return (static_cast<lvtk::URIDirectory*>(d))->get_unmap_feature();
}

}
