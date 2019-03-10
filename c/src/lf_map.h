#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

#ifndef LF_MAP
#define LF_MAP

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

unsigned long hash(unsigned char *str);

bool hm_entry_key_equals(char *k1, char *k2);

int lf_map_create();

int lf_map_get(hm_entry *entry);

void lf_map_put(hm_entry *entry);

void lf_map_destroy();

#endif