class Solution {
public:
    int matrixScore(vector<vector<int>>& grid) {
        int n = grid.size(), m = grid[0].size();

        // Processing the rows.
        for(int i = 0; i < n; i++) {
            // Ensuring we have the most representative bit 1 to ensure the biggest number.
            if (grid[i][0] == 0) {
                flipRow(grid[i]);
            }
        }

        // Processing the columns.
        for(int i = 0; i < m; i++) {
            int ones = countOnes(grid, i, n);

            if (ones <= n / 2) {
                flipCol(grid, i, n);
            }
        }
        
        // Processing the bit representation as decimal.

        int sum = 0;

        for (const auto& row: grid) {
            string numString = convertRowToString(row);
            int num = binaryToDecimal(numString);

            sum += num;
        }

        return sum;
    }  
private:
    int countOnes(const vector<vector<int>> &grid, int colNum, int rowSize) {
        int count = 0;

        for(int i = 0; i < rowSize; i++) {
            if (grid[i][colNum] == 1) {
                count++;
            }
        }

        return count;
    }

    void flipRow(vector<int>& row) {
        for (auto& num: row) {
            num ^= 1;
        }
    }

    void flipCol(vector<vector<int>> &grid, int colNum, int rowSize) {
        for (int i = 0; i < rowSize; i++) {
            grid[i][colNum] ^= 1;
        }
    }

    string convertRowToString(const vector<int>& row) {
        string numString = "";

        for (const auto& num: row) {
            numString += to_string(num);
        }

        return numString;
    }

    int binaryToDecimal(string n) { 
        string num = n; 
        int dec_value = 0; 
    
        // Initializing base value to 1, i.e 2^0 
        int base = 1; 
    
        int len = num.length(); 
        for (int i = len - 1; i >= 0; i--) { 
            if (num[i] == '1') 
                dec_value += base; 
            base = base * 2; 
        } 
    
        return dec_value; 
    } 
};
