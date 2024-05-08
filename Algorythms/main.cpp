#include <iostream>
#include <queue>

using namespace std;

class Node
{
public:
    int val;
    Node *left;
    Node *right;
    Node *next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node *_left, Node *_right, Node *_next)
        : val(_val), left(_left), right(_right), next(_next) {}
};

struct status
{
    Node *node;
    int depth;
};

class Solution
{
public:
    Node *connect(Node *root)
    {
        if (root == NULL)
            return NULL;

        queue<status> q;

        q.push({root, 0});

        for (; !q.empty();)
        {
            status top = q.front();
            q.pop();
            if (top.node->left != NULL)
            {
                q.push({top.node->left, top.depth + 1});
            }
            if (top.node->right != NULL)
            {
                q.push({top.node->right, top.depth + 1});
            }

            if (q.empty())
            {
                continue;
            }
            if (q.front().depth != top.depth)
            {
                continue;
            }
            top.node->next = q.front().node;
        }

        return root;
    }
};

int main()
{
    return 0;
}