#include <assert.h>
#include <stdio.h>
#include "lf_map.h"


#define NUM_ENTRIES 26

void main() {
    lf_map_create();

    hm_entry *entries[NUM_ENTRIES];
    int i;
    for (i = 0; i < NUM_ENTRIES; i++) {
        hm_entry *curr_entry = (hm_entry *)malloc(sizeof(hm_entry));
        char *key = (char *) malloc(2 * sizeof(char)); 
        key[0]= 'a' + i;
        key[1]= '\0';
        curr_entry->key = key;
        // curr_entry->key[0]++;
        // printf("i: %d key:%s\n", i, curr_entry->key);
        curr_entry->val = i;
        entries[i] = curr_entry;
        lf_map_put(curr_entry);
    }

    for (i = 0; i < NUM_ENTRIES; i++) {
        int curr_val = lf_map_get(entries[i]);
        // printf("curr_val: %d, expected: %d\n", curr_val, i);
        assert(curr_val == i);
    }

    entries[0]->val = 99;
    lf_map_put(entries[0]);
    assert(lf_map_get(entries[0]) == 99);
    lf_map_destroy();
    printf("Passed\n");
}