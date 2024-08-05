#include <stdlib.h>
#include <stdio.h>

struct ListNode {
    int val;
    struct ListNode* next;
};

struct ListNode* createNode(int val) {
    struct ListNode* node = malloc(sizeof(struct ListNode));

    node->val = val;
    node->next = NULL;
}

void addNode(struct ListNode** curr, int val) {
    struct ListNode* node = createNode(val);

    if (*curr == NULL) {
        *curr = node;
        return;
    }

    (*curr)->next = node;
    *curr = node;
}

void cleanup(struct ListNode* head) {
    if (head == NULL) {
        return;
    }

    cleanup(head->next);
    free(head);
}

void traverseSLL(struct ListNode* head) {
    struct ListNode* node = head;

    while (node != NULL) {
        printf("%d ", node->val);
        node = node->next;
    }
}

struct ListNode* createSLL(int arr[], int size) {
    struct ListNode* head = NULL;
    struct ListNode* curr = NULL;

    for(int i = 0; i < size; i++) {
        if (head == NULL) {
            addNode(&head, arr[i]);
            curr = head;
            continue;
        }

        addNode(&curr, arr[i]);
    }

    return head;
}

struct ListNode* middleOfSLL(struct ListNode* head) {
    if (head == NULL) {
        return NULL;
    }

    struct ListNode *slow_p = head, *fast_p = head;

    while(fast_p->next != NULL) {
        fast_p = fast_p->next;
        slow_p = slow_p->next;

        if (fast_p->next == NULL) {
            break;
        }

        fast_p = fast_p->next;
    }

    return slow_p;
}

int main() {
    int nums1[] = {1, 2, 3, 4, 5, 6};
    int nums1Size = 6;

    struct ListNode* l1 = createSLL(nums1, nums1Size);
    traverseSLL(middleOfSLL(l1));

    cleanup(l1);

    return 0;
}