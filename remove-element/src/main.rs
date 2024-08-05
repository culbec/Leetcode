use std::vec::Vec;

fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
    let mut k : i32 = 0;
    let mut index : usize = 0;

    for i in 0..nums.len() {
        if nums[i] != val {
            nums[index] = nums[i];
            k = k + 1;
            index = index + 1;
        }
    }

    k
}

fn main() {
    let mut vec : Vec<i32> = vec![3, 2, 2, 3];

    println!("{}", remove_element(&mut vec, 2));
}
