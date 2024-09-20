var isSelfCrossing = function (distance) {
    if (distance.length <= 3) return false;
    let i = 2;
    while (i < distance.length && distance[i] > distance[i - 2]) i++;
    if (i >= distance.length) return false;
    if ((i >= 4 && distance[i] >= distance[i - 2] - distance[i - 4]) ||
        (i === 3 && distance[i] === distance[i - 2])) {
        distance[i - 1] -= distance[i - 3];
    }
    i++;
    while (i < distance.length) {
        if (distance[i] >= distance[i - 2]) return true;
        i++;
    }
    return false;
};