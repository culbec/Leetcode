#include <string>
using namespace std;

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
    string smallestStr;

    string smallestFromLeaf(TreeNode* root) {
        dfs(root, "");
        return this->smallestStr;
    }

    void dfs(TreeNode* root, string currStr) {
        if (root == nullptr) {
            return;
        }

        currStr = char('a' + root->val) + currStr;
        if (root->left == nullptr && root->right == nullptr) {
            if (this->smallestStr == "" || this->smallestStr > currStr) {
                this->smallestStr = currStr;
            }
        }

        dfs(root->left, currStr);
        dfs(root->right, currStr);
    }
};
