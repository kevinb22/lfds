#include "lf_map.h"

#define HT_SIZE 97

hm_bucket *hash_table;

int lf_map_create()
{
    hash_table = (hm_bucket *)calloc(HT_SIZE, sizeof(hm_bucket));
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
    // printf("PUT: key: %s, hash_val: %lu\n", entry->key, hash_key);

    unsigned long index = hash_key % HT_SIZE;

    // TODO: Insert into hash table
    // atomic
    hm_bucket *bucket = &hash_table[index];
    // Look at the version...

    lf_node *head = bucket->list;
    lf_node *next = head;
    while (next) {
        if (hm_entry_key_equals(entry->key, ((hm_entry *) next->data)->key)) {

            // check version number
            ((hm_entry *) next->data)->val = entry->val;
            return;
        }
        next = next->next;
    }

    lf_node *new_head = (lf_node *) malloc(sizeof(lf_node));
    new_head->data = entry;
    new_head->next = head;
    bucket->list = new_head;
}

int lf_map_get(hm_entry *entry)
{
    if (!entry) {
        return -1;
    }

    unsigned long hash_key = hash(entry->key);
    // printf("GET: key: %s, hash_val: %lu\n", entry->key, hash_key);
    unsigned long index = hash_key % HT_SIZE;

    // atomic
    hm_bucket bucket = hash_table[index];

    lf_node *curr = bucket.list;

    while (curr) {
        if (hm_entry_key_equals(entry->key, ((hm_entry *) curr->data)->key)) {
            // check version number
            return ((hm_entry *) curr->data)->val;
        }
        curr = curr->next;
    }

    return -1;
}

void lf_map_destroy() {
    hm_bucket *curr_bucket = hash_table;
    int i;
    for (i = 0; i < HT_SIZE; i++, curr_bucket++) {
        lf_node *curr_node = curr_bucket->list;
        while (curr_node) {
            free(curr_node->data);
            lf_node *old_node = curr_node;
            curr_node = curr_node->next;
            free (old_node);
        }
    }
    free(hash_table);
}

unsigned long hash(unsigned char *str)
{
    unsigned long hash_code = 5381;
    int c;

    while (c = *str++)
        hash_code = ((hash_code << 5) + hash_code) + c; /* hash * 33 + c */

    return hash_code;
}

bool hm_entry_key_equals(char *k1, char *k2) {
    return strcmp(k1, k2) == 0;
}
