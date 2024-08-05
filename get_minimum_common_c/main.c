#include <stdio.h>

int get_common(int* nums1, int nums1Size, int* nums2, int nums2Size) {
    if (nums1Size == 0 || nums2Size == 0) {
        return -1;
    }

    int i = 0, j = 0;

    while (i < nums1Size && j < nums2Size) {
        if (nums1[i] == nums2[j]) {
            return nums1[i];
        }

        if (nums1[i] < nums2[j]) {
            i++;
        } else {
            j++;
        }
    }

    return -1;
}

int main(int argc, char **argv) {
    int nums1[] = {1, 2, 3};
    int nums1Size = 3;
    
    int nums2[] = {2, 4};
    int nums2Size = 2;

    printf("%d\n", get_common(nums1, nums1Size, nums2, nums2Size));

    return 0;
}