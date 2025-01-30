function main() {
  tc1 = [2, 7, 11, 15];
  tc2 = [3, 2, 4];
  tc3 = [3, 3];
  console.log(twoSum(tc1, 9));
  console.log(twoSum(tc2, 6));
  console.log(twoSum(tc3, 6));
}

function twoSum(nums, target) {
  let exist = {};
  for (let i = 0; i < nums.length; i++) {
    if (target - nums[i] in exist) {
      return [i, exist[target - nums[i]]];
    }

    exist[nums[i]] = i;
  }
}

main();
