#include <iostream>

struct ListNode {
    int val;
    struct ListNode* next;
};

ListNode* createNode(int val) {
    ListNode* node = new ListNode;

    node->val = val;
    node->next = NULL;

    return node;
}

void addNode(ListNode** curr, int val) {
    ListNode* node = createNode(val);

    if (*curr == NULL) {
        *curr = node;
        return;
    }

    (*curr)->next = node;
    *curr = node;
}

void cleanup(ListNode* head) {
    if (head == NULL) {
        return;
    }

    cleanup(head->next);
    delete head;
}

void traverseSLL(ListNode* head) {
    ListNode* node = head;

    while (node != NULL) {
        printf("%d ", node->val);
        node = node->next;
    }
}

ListNode* createSLL(int arr[], int size) {
    ListNode* head = NULL;
    ListNode* curr = NULL;

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

ListNode* middleOfSLL(ListNode* head) {
    if (head == NULL) {
        return NULL;
    }

    ListNode *slow_p = head, *fast_p = head;

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

    ListNode* l1 = createSLL(nums1, nums1Size);
    traverseSLL(middleOfSLL(l1));

    cleanup(l1);

    return 0;
}