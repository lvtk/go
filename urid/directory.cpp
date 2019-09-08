
#include <map>
#include <string>
#include "directory.h"

namespace lvtk {

/** Maintains a map of Strings/Symbols to integers
    
    This class also implements LV2 URID Map/Unmap features.  Plugin 
    implementations don't need to use this.  You can, however, use this in a 
    LV2 host to easily provide URID map/unmaping features to plugins.

    @headerfile lvtk/uri_directory.hpp
    @ingroup urid
 */
class URIDirectory
{
public:
    /** Create an empty symbol map and initialized LV2 URID features */
    URIDirectory()
    {
        map_feature.URI    = LV2_URID__map;
        map_data.handle    = (void*) this;
        map_data.map       = &URIDirectory::_map;
        map_feature.data   = &map_data;

        unmap_feature.URI  = LV2_URID__unmap;
        unmap_data.handle  = this;
        unmap_data.unmap   = _unmap;
        unmap_feature.data = &unmap_data;
    }

    ~URIDirectory()
    {
        clear();
    }

    /** Map a symbol/uri to an unsigned integer
        @param key The symbol to map
        @returns A mapped URID, a return of 0 indicates failure */
    inline uint32_t map (const char* key)
    {
        if (! contains (key))
        {
            const uint32_t urid (1 + (uint32_t) mapped.size());
            mapped [key] = urid;
            unmapped [urid] = std::string (key);
            return urid;
        }

        return mapped [key];
    }

    /** Containment test of a URI
        
        @param uri The URI to test
        @returns True if found */
    inline bool contains (const char* uri) {
        return mapped.find (uri) != mapped.end();
    }

    /** Containment test of a URID
        
        @param urid The URID to test
        @return True if found */
    inline bool contains (uint32_t urid) {
        return unmapped.find (urid) != unmapped.end();
    }

    /** Unmap an already mapped id to its symbol
        
        @param urid The URID to unmap
        @return The previously mapped symbol or 0 if the urid isn't in the cache
     */
    inline const char* unmap (uint32_t urid) {
        if (contains (urid))
            return (const char*) unmapped [urid].c_str();
        return "";
    }

    /** Clear the URIDirectory */
    inline void clear()
    {
        mapped.clear();
        unmapped.clear();
    }

    /** @returns a LV2_Feature with LV2_URID_Map as the data member */
    const LV2_Feature *const get_map_feature()      const { return &map_feature; }
    /** @returns a LV2_Feature with LV2_URID_Unmap as the data member */
    const LV2_Feature *const get_unmap_feature()    const { return &unmap_feature; }

private:
    std::map<std::string, uint32_t> mapped;
    std::map<uint32_t, std::string> unmapped;

    LV2_Feature         map_feature;
    LV2_URID_Map        map_data;
    LV2_Feature         unmap_feature;
    LV2_URID_Unmap      unmap_data;

    static uint32_t _map (LV2_URID_Map_Handle self, const char* uri) {
        return (static_cast<URIDirectory*> (self))->map (uri);
    }

    static const char* _unmap (LV2_URID_Unmap_Handle self, uint32_t urid) {
        return (static_cast<URIDirectory*> (self))->unmap (urid);
    }
};

}

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

LV2_URID_Unmap* lvtk_uri_directory_get_unmap(LvtkURIDirectory* d) {
    return (LV2_URID_Unmap*) lvtk_uri_directory_get_unmap_feature(d)->data;
}

}
