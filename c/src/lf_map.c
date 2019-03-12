#include "lf_map.h"
#include <stdio.h>

#define HT_SIZE 97

hm_bucket *hash_table;

int lf_map_create() {
    hash_table = (hm_bucket *)calloc(HT_SIZE, sizeof(hm_bucket));
    if (!hash_table) {
        return -1;
    }
    int i;
    for (i = 0; i < HT_SIZE; i++) {
        hash_table[i].version = 1;
        __sync_lock_release(&hash_table[i].spin_flag);
    }
    return 0;
}

void lf_map_put(hm_entry *entry) {
    if (entry == NULL) {
        return;
    }
    int version_old;
    unsigned long hash_key = hash(entry->key);

    unsigned long index = hash_key % HT_SIZE;

    hm_bucket *bucket = &hash_table[index];

LOOP:
    version_old = bucket->version;

    lf_node *head = bucket->list;
    lf_node *next = head;
    while (next) {
        if (hm_entry_key_equals(entry->key, ((hm_entry *)next->data)->key)) {
            if (!__sync_bool_compare_and_swap(&bucket->spin_flag, 0, 1)) {
                goto LOOP;
            }
            if (!__sync_bool_compare_and_swap(&bucket->version, version_old, version_old + 1)) {
                if (!__sync_bool_compare_and_swap(&bucket->spin_flag, 1, 0)) {
                    printf("Trying to release unheld lock...exiting\n");
                    exit(1);
                }
                goto LOOP;
            }

            ((hm_entry *)next->data)->val = entry->val;
            if (!__sync_bool_compare_and_swap(&bucket->spin_flag, 1, 0)) {
                printf("Trying to release unheld lock...exiting\n");
                exit(1);
            }
            return;
        }
        next = next->next;
    }

    lf_node *new_head = (lf_node *)malloc(sizeof(lf_node));
    new_head->data = entry;
    new_head->next = head;
    if (!__sync_bool_compare_and_swap(&bucket->spin_flag, 0, 1)) {
        free(new_head);
        goto LOOP;
    }
    if (!__sync_bool_compare_and_swap(&bucket->version, version_old, version_old + 1)) {
        if (!__sync_bool_compare_and_swap(&bucket->spin_flag, 1, 0)) {
            printf("Trying to release unheld lock...exiting\n");
            free(new_head);
            exit(1);
        }
        free(new_head);
        goto LOOP;
    }
    bucket->list = new_head;
    if (!__sync_bool_compare_and_swap(&bucket->spin_flag, 1, 0)) {
        printf("Trying to release unheld lock...exiting\n");
        free(new_head);
        exit(1);
    }
}

int lf_map_get(hm_entry *entry) {
    if (!entry) {
        return -1;
    }
    int version_old;
    int version_new;
    unsigned long hash_key = hash(entry->key);
    unsigned long index = hash_key % HT_SIZE;

    hm_bucket *bucket = &hash_table[index];

LOOP:
    version_old = bucket->version;

    lf_node *curr = bucket->list;

    while (curr) {
        __atomic_load(&bucket->version, &version_new, __ATOMIC_RELAXED);
        if (version_old != version_new) {
            goto LOOP;
        }
        if (hm_entry_key_equals(entry->key, ((hm_entry *)curr->data)->key)) {
            return ((hm_entry *)curr->data)->val;
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
            free(old_node);
        }
    }
    free(hash_table);
}

unsigned long hash(unsigned char *str) {
    unsigned long hash_code = 5381;
    int c;

    while (c = *str++)
        hash_code = ((hash_code << 5) + hash_code) + c; /* hash * 33 + c */

    return hash_code;
}

bool hm_entry_key_equals(char *k1, char *k2) {
    return strcmp(k1, k2) == 0;
}
