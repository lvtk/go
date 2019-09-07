
#include "directory.h"
#include <lvtk/uri_directory.hpp>

extern "C" {

LVTKURIDirectory* lvtk_uri_directory_new() {
    lvtk::URIDirectory* d = new lvtk::URIDirectory();
    return (LVTKURIDirectory*) d;
}

void lvtk_uri_directory_free (LVTKURIDirectory* d) {
    delete static_cast<lvtk::URIDirectory*> (d);
}

uint32_t lvtk_uri_directory_map (LVTKURIDirectory* d, const char* uri) {
    return (static_cast<lvtk::URIDirectory*>(d))->map (uri);
}

const char* lvtk_uri_directory_unmap (LVTKURIDirectory* d, uint32_t urid) {
    return (static_cast<lvtk::URIDirectory*>(d))->unmap (urid);
}

const LV2_Feature* lvtk_uri_directory_get_map_feature(LVTKURIDirectory* d) {
    return (static_cast<lvtk::URIDirectory*>(d))->get_map_feature();
}

const LV2_Feature* lvtk_uri_directory_get_unmap_feature(LVTKURIDirectory* d) {
    return (static_cast<lvtk::URIDirectory*>(d))->get_unmap_feature();
}

}
