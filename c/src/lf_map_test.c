#include "lf_map.h"
#include <assert.h>
#include <pthread.h>
#include <stdio.h>

#define NUM_ENTRIES 26
#define NUM_THREADS 16

void test_single_threaded() {
    printf("Testing single threaded...\n");
    lf_map_create();

    hm_entry *entries[NUM_ENTRIES];
    int i;
    for (i = 0; i < NUM_ENTRIES; i++) {
        hm_entry *curr_entry = (hm_entry *)malloc(sizeof(hm_entry));
        char *key = (char *)malloc(2 * sizeof(char));
        key[0] = 'a' + i;
        key[1] = '\0';
        curr_entry->key = key;
        curr_entry->val = i;
        entries[i] = curr_entry;
        lf_map_put(curr_entry);
    }

    for (i = 0; i < NUM_ENTRIES; i++) {
        int curr_val = lf_map_get(entries[i]);
        assert(curr_val == i);
    }

    entries[0]->val = 99;
    lf_map_put(entries[0]);
    assert(lf_map_get(entries[0]) == 99);
    lf_map_destroy();
    printf("Passed single threaded\n");
}

void *thread_start(void *args) {
    int thread_id = (int)args;

    hm_entry *entries[NUM_ENTRIES];
    int i;
    for (i = 0; i < NUM_ENTRIES; i++) {
        hm_entry *curr_entry = (hm_entry *)malloc(sizeof(hm_entry));
        char *key = (char *)malloc(2 * sizeof(char));
        key[0] = 'a' + i;
        key[1] = '\0';
        curr_entry->key = key;
        curr_entry->val = i;
        entries[i] = curr_entry;
        lf_map_put(curr_entry);
    }

    for (i = 0; i < NUM_ENTRIES; i++) {
        int curr_val = lf_map_get(entries[i]);
        if (curr_val != i) {
            printf("i: %d curr_val: %d\n", i, curr_val);
        }
        assert(curr_val == i);
    }

    pthread_exit(0);
}

void test_multi_threaded() {
    printf("Testing multi threaded...\n");
    lf_map_create();
    pthread_t thread_arr[NUM_THREADS];
    int i;
    for (i = 0; i < NUM_THREADS; i++) {
        pthread_create(&thread_arr[i], NULL, thread_start, (void *)i);
    }

    int err = 0;
    bool fail = false;
    for (i = 0; i < NUM_THREADS; i++) {
        pthread_join(thread_arr[i], (void **)&err);
        fail |= err;
    }

    if (fail) {
        printf("Failed multi threaded\n");
    } else {
        printf("Passed multi threaded\n");
    }
    lf_map_destroy();
}

void main() {
    test_single_threaded();
    test_multi_threaded();
}