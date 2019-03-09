#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

typedef struct _hm_entry
{
    unsigned char *key;
    int val;
} hm_entry;

typedef struct _lf_node {
    void *data;
    struct _lf_node *next;
} lf_node;

typedef struct _hm_bucket {
    int version;
    lf_node *list;
} hm_bucket;

void lf_map_put(hm_entry *entry);

int lf_map_get(hm_entry *entry);

unsigned long
hash(unsigned char *str)
{
    unsigned long hash = 5381;
    int c;

    while (c = *str++)
        hash = ((hash << 5) + hash) + c; /* hash * 33 + c */

    return hash;
}

bool hm_entry_key_equals(hm_entry *e1, hm_entry *e2) {
    if (!e1 && !e2) {
        return true;
    }
    if (!e1 || !e2) {
        return false;
    }

    return strcmp(e1->key, e2->key) == 0;
}

