#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>

struct ListNode {
    int val;
    struct ListNode* next;
};

bool rec(struct ListNode* head, struct ListNode* tail) {
    if (tail == NULL) {
        return true;
    }

    bool r = rec(head, tail->next);
    bool ok = head->val == tail->val;

    if (head == tail || head->next == tail) {
        return ok;
    }

    head = head->next;
    return ok && r;
}

bool isPalindrome(struct ListNode* head) {
    return rec(head, head->next);
}

struct ListNode* createSLL(int arr[], int n) {
    struct ListNode* head = NULL;
    struct ListNode* curr = NULL;

    for (int i = 0; i < n; i++) {
        if (head == NULL) {
            head = (struct ListNode*) malloc (sizeof(struct ListNode));
            head->val = arr[i];
            head->next = NULL;

            curr = head;
        } else {
            struct ListNode* node = (struct ListNode*) malloc (sizeof(struct ListNode));
            node->val = arr[i];
            node->next = NULL;

            curr->next = node;
            curr = node;
        }
    }

    return head;
}

void clear(struct ListNode* head) {
    if (head == NULL) {
        return;
    }

    clear(head->next);
    free(head);
}

void main() {
    int arr[] = {1, 2, 2, 1};
    int n = 4;

    struct ListNode* head = createSLL(arr, n);

    printf("%d\n", isPalindrome(head));
    
    clear(head);
}