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
    class Condition {
    public:
        int sum;
        int min_n;
        int max_n;
        bool isBst;

        Condition()
        {
            sum = 0;
            min_n = INT_MAX;
            max_n = INT_MIN;
            isBst = true;
        }
    };

    Condition check(TreeNode* root, int& ans) 
    {
        if (root == NULL) {
            return Condition();  
        }

        Condition left = check(root->left, ans);
        Condition right = check(root->right, ans);

        Condition current;
        
        if (left.isBst && right.isBst && left.max_n < root->val && root->val < right.min_n) 
        {
            current.isBst = true;
            current.sum = left.sum + right.sum + root->val;
            current.min_n = min(left.min_n, root->val);
            current.max_n = max(right.max_n, root->val);

            ans = max(ans, current.sum);
        } 
        else 
        {
            current.isBst = false;
        }

        return current;
    }

    int maxSumBST(TreeNode* root) {
        int ans = 0;

        check(root, ans);

        return ans;
    }
};