class Solution {
public:
    double mincostToHireWorkers(vector<int>& quality, vector<int>& wage, int k) {
        int n = quality.size();

        double totalCost = numeric_limits<double>::max();
        double currentTotalQuality = 0;

        // Keeping track of the workers wage based on its quality.
        vector<pair<double, int>> wageToQualityRatio;
        for (int i = 0; i < n; i++) {
            wageToQualityRatio.push_back({(double) wage[i] / quality[i], quality[i]});
        }

        // Sorting the workers based on their wage to quality ratio.
        sort(wageToQualityRatio.begin(), wageToQualityRatio.end());

        // Initializing the priority queue to keep track of the most 'valuable' workers.
        priority_queue<int> highestQualityWorkers;

        for (int i = 0; i < n; i++) {
            highestQualityWorkers.push(wageToQualityRatio[i].second);
            currentTotalQuality += wageToQualityRatio[i].second;

            if (highestQualityWorkers.size() > k) {
                currentTotalQuality -= highestQualityWorkers.top();
                highestQualityWorkers.pop();
            }

            if (highestQualityWorkers.size() == k) {
                totalCost = min(totalCost, currentTotalQuality * wageToQualityRatio[i].first);
            }
        }

        return totalCost;
    }
};
