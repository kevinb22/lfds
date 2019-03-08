#include "lf_map.h"

#include <stdlib.h>

#define HT_SIZE 97

typedef struct _hm_entry
{
    char *key;
    int val;
} hm_entry;

hm_entry **hash_table;

int lf_map_create()
{
    hash_table = (hm_entry **)calloc(HT_SIZE, sizeof(hm_entry *));
    if (!hash_table)
    {
        return -1;
    }
    return 0;
}

void lf_map_put(hm_entry *entry)
{
    if (entry == NULL)
    {
        return;
    }
    unsigned long hash_key = hash(entry->key);

    unsigned long index = hash_key % HT_SIZE;

    // TODO: Insert into hash table
    if (!hash_table[index])
    {
    }
}

int lf_map_get(hm_entry *entry)
{

    return 0;
}