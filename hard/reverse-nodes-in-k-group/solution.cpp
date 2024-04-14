/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* reverse(ListNode* head) {

         ListNode* prev = NULL;  
         ListNode* curr = head;

         while(curr) {
             ListNode* tmp = curr->next;
             curr->next = prev;
             prev = curr; 
             curr = tmp;
         } 

         return prev; 
    }

    ListNode* reverseKGroup(ListNode* head, int k) {
        
        if(!head || !head->next || k == 0 || k == 1) return head;

        ListNode* head1 = head;
        ListNode* tail1 = head;

        ListNode* tmp = head;
        int len = 0; 

        while(tmp) {
            len++;
            tmp = tmp->next;
        }

        if(k > len) return head;

        int l = 1;

        while(l < k && tail1->next) {
            tail1 = tail1->next;
            l++;
        } 

        ListNode* head2 = tail1->next;
        tail1->next = NULL;

        head2 = reverseKGroup(head2, k); 

        head1 = reverse(head);
        tail1 = head1;

        while(tail1->next) tail1 = tail1->next;

        tail1->next = head2;

        return head1;
    }
};