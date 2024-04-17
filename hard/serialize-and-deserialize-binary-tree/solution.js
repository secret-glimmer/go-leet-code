let MARKER = "M",
    m = 1;

function serializeRec(node, stream) {
    // Adding marker to stream if the node is null
    if (node == null) {
        stream.push(MARKER + m);
        m++;
        return stream;
    }
    // Adding node to stream
    stream.push(node.val);

    // Doing a pre-order tree traversal for serialization
    stream = serializeRec(node.left, stream);
    return serializeRec(node.right, stream);
}

var serialize = function (root) {
    let stream = [];
    stream = serializeRec(root, stream);
    return stream;
};

/**
 * Decodes your encoded data to tree.
 *
 * @param {string} data
 * @return {TreeNode}
 */


function deserialize_helper(stream) {
    // dequeue last element from list
    let val = stream.pop();

    // Return null when a marker is encountered
    if (typeof val == "string" && val[0] == MARKER)
        return null;

    // Creating new Binary Tree Node for current value from stream
    let node = new TreeNode(val);

    // Doing a pre-order tree traversal for deserialization
    node.left = deserialize_helper(stream);
    node.right = deserialize_helper(stream);

    // return node if it exists
    return node;
}

var deserialize = function (stream) {
    stream.reverse()
    let node = deserialize_helper(stream)
    return node
};

/**
 * Your functions will be called as such:
 * deserialize(serialize(root));
 */