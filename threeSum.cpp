#include <hash_map>

using namespace std;

class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& src) {
		vector<vector<int>> ret;

		vector<int> nums;
		int idxcount = 0;

		hash_map<int,hash_map<int,bool>> nummap;
		for (int i=0; i<src.size(); i++) {
			auto it = nummap.find(src[i]);
			if (it != nummap.end()) {
				nummap[src[i]] = hash_map<int,bool>();
			}
			int n = nummap[src[i]].size();
			if (n < 3) {
				nummap[src[i]][i-idxcount] = true;
				nums.push_back(src[i]);
			} else {
				idxcount++;
			}
		}
        
        long long sum = 0;
        map<vector<int>, bool> hash;
        for (int i=0; i<nums.size(); i++) {
			for (int j=0; j<nums.size(); j++) {
				if (j != i) {
					sum = nums[i]+nums[j];
					hash_map<int,hash_map<int,bool>>::iterator it = nummap.find(-sum);
					if (it != nummap.end()) {
						int n = 0;
						hash_map<int,bool> v = it->second;
						if (v.find(i) != v.end()) {
							n++;
						}
						if (v.find(j) != v.end()) {
							n++;
						}
						if (v.size() > n) {
							vector<int> tmp;
							tmp.push_back(nums[i]);
							tmp.push_back(nums[j]);
							tmp.push_back(-sum);
							if (tmp[0] > tmp[1]) {
								int v = tmp[0];
								tmp[0] = tmp[1];
								tmp[1] = v;
							}
							if (tmp[1] > tmp[2]) {
								int v = tmp[1];
								tmp[1] = tmp[2];
								tmp[2] = v;
								if (tmp[0] > tmp[1]) {
									v = tmp[0];
									tmp[0] = tmp[1];
									tmp[1] = v;
								}
							}
							if (hash.find(tmp) == hash.end()) {
								hash[tmp] = true;
								ret.push_back(tmp);
							}
						}
					}
				}
			}
		}
        return ret;
    }
};
