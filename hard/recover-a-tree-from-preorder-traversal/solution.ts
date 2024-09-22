class TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
    constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.left = (left===undefined ? null : left)
        this.right = (right===undefined ? null : right)
    }
}

function recoverFromPreorder(traversal: string): TreeNode | null {
    const stack: TreeNode[] = [];
    let i = 0;

    while (i < traversal.length) {
        let level = 0;

        // Count the number of dashes to determine the node level
        while (i < traversal.length && traversal[i] === '-') {
            level++;
            i++;
        }

        // Parse the node value
        let value = 0;
        while (i < traversal.length && traversal[i] >= '0' && traversal[i] <= '9') {
            value = value * 10 + (traversal[i].charCodeAt(0) - '0'.charCodeAt(0));
            i++;
        }

        const node = new TreeNode(value);

        // Adjust the stack to match the current node's level
        while (stack.length > level) {
            stack.pop();
        }

        // Attach the node to its parent
        if (stack.length > 0) {
            if (!stack[stack.length - 1].left) {
                stack[stack.length - 1].left = node;
            } else {
                stack[stack.length - 1].right = node;
            }
        }

        // Push the current node to the stack
        stack.push(node);
    }

    // The root node is the bottom-most element in the stack
    return stack.length > 0 ? stack[0] : null;
}

// Optional helper function to serialize the tree (for testing)
function serialize(root: TreeNode | null): (number | null)[] {
    const result: (number | null)[] = [];
    if (!root) return result;

    const queue: (TreeNode | null)[] = [root];

    while (queue.length > 0) {
        const node = queue.shift()!;
        if (node) {
            result.push(node.val);
            queue.push(node.left);
            queue.push(node.right);
        } else {
            result.push(null);
        }
    }

    // Remove trailing nulls for cleaner output
    while (result.length > 0 && result[result.length - 1] === null) {
        result.pop();
    }

    return result;
}