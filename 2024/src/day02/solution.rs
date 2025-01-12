use crate::day01::read_input_from_file;

const PATH: &str = "2024/src/day02/input.txt";

pub fn solution1() {
    let input = read_input_from_file(PATH);
    let mut safe = 0;
    for i in 0..input.len() {
        let mut inc = 0;
        let mut dec = 0;
        let mut is_unsafe = false;
        for j in 1..input[i].len() {
            let dif = input[i][j - 1] - input[i][j];
            // checking for two adjacent levels differs by atleast one and at most 3
            if dif.abs() > 3 {
                is_unsafe = true;
                break;
            }
            // check if its inc or dec only
            if dif > 0 {
                inc += 1;
            } else if dif < 0 {
                dec += 1;
            } else {
                is_unsafe = true;
                break;
            }
        }
        if !is_unsafe && (inc == 0 || dec == 0) {
            safe += 1;
        }
    }
    println!("DAY02 Solution 1: {}", safe)
}

pub fn solution2() {
    let input = read_input_from_file(PATH);
    let mut safe = 0;

    for i in 0..input.len() {
        if is_safe(&input[i]) {
            safe += 1;
        } else {
            for j in 0..input[i].len() {
                let mut modified_input = input[i].clone();
                modified_input.remove(j);
                if is_safe(&modified_input) {
                    safe += 1;
                    break;
                }
            }
        }
    }

    println!("DAY02 Solution 2: {}", safe);
}

fn is_safe(report: &Vec<i32>) -> bool {
    let mut inc = 0;
    let mut dec = 0;
    for j in 1..report.len() {
        let dif = report[j - 1] - report[j];
        // checking for two adjacent levels differs by atleast one and at most 3
        if dif.abs() > 3 {
            return false;
        }
        // check if its inc or dec only
        if dif > 0 {
            inc += 1;
        } else if dif < 0 {
            dec += 1;
        } else {
            return false;
        }
    }
    inc == 0 || dec == 0
}
