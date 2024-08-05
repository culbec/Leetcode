/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
struct TreeNode* removeLeafNodes(struct TreeNode* root, int target) {
    // Base case: null node.
    if (root == NULL) {
        return NULL;
    }

    // Postorder traversal of the tree.
    root->left = removeLeafNodes(root->left, target);
    root->right = removeLeafNodes(root->right, target);

    // Checking if the current node is a leaf and has the value of target.
    if (root->left == NULL && root->right == NULL && root->val == target) {
        return NULL;
    }

    return root;
}
