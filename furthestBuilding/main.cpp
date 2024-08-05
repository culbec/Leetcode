#include <queue>

int furthestBuilding(std::vector<int> heights, int bricks, int ladders) {
    std::priority_queue<int, std::vector<int>, std::less<>> gaps;
    int currHeight = heights.at(0);
    int noBuildings = 0;

    for (int i = 0 ; i < heights.size(); i++) {
        int height = heights.at(i);

        if (currHeight > height) {
            noBuildings++;
        } else {
            if (ladders == 0 && bricks == 0) {
                break;
            } else if (ladders > 0) {
                int gap = height - currHeight;
                gaps.push(gap);
                ladders--;
                noBuildings++;
            } else {
                if (gaps.empty()) {
                    if (height > bricks) {
                        break;
                    } else {
                        bricks -= height;
                        noBuildings++;
                    }
                } else {
                    int minGap = gaps.top();
                    gaps.pop();

                    if (minGap > bricks) {
                        break;
                    } else {
                        bricks -= minGap;
                        noBuildings++;
                    }
                }
            }
        }

        currHeight = height;
    }

    return noBuildings;
}