/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int sumOfLeftLeaves(TreeNode* root) {
        // Adding a flag variable that has the value:
        // * 0: the node is not on the left side.
        // * 1: the node is on the left side.
        return recursiveSum(root, 0);
    }

private:
    int recursiveSum(TreeNode* root, int flag) {
        if (root == nullptr) {
            return 0;
        } 
        
        // Leaf node case.
        if (root->left == nullptr && root->right == nullptr && flag == 1) {
            return root->val;
        }

        // Divide et impera to calculate the sum on both branches.
        return recursiveSum(root->left, 1) + recursiveSum(root->right, 0);
    }
};
