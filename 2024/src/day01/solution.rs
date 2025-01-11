use std::{collections::HashMap, fs::read_to_string};

pub fn read_input_from_file() -> Vec<Vec<i32>> {
    let file = read_to_string("2024/src/day01/input.txt").expect("failed to read file");
    file.lines()
        .map(|line| {
            line.split_whitespace()
                .map(|x| x.parse::<i32>().expect("Failed to parse number"))
                .collect::<Vec<i32>>()
        })
        .collect::<Vec<Vec<i32>>>()
}

pub fn solution1() {
    let input = read_input_from_file();
    let mut vec1 = Vec::new();
    let mut vec2 = Vec::new();
    for i in 0..input.len() {
        vec1.push(input[i][0]);
        vec2.push(input[i][1]);
    }
    vec1.sort();
    vec2.sort();

    let mut sum = 0;
    for i in 0..vec1.len() {
        let diff = vec1[i] - vec2[i];
        sum += diff.abs();
    }

    println!("DAY01 Solution 1: {}", sum)
}

pub fn solution2() {
    let input = read_input_from_file();
    let mut vec1 = Vec::new();
    let mut vec2 = Vec::new();
    for i in 0..input.len() {
        vec1.push(input[i][0]);
        vec2.push(input[i][1]);
    }
    let mut hm: HashMap<i32, i32> = HashMap::new();
    for i in 0..vec2.len() {
        let result = hm.insert(vec2[i], 1);
        if let Some(count) = result {
            hm.insert(vec2[i], count + 1);
        }
    }
    let mut sum = 0;
    for i in 0..vec1.len() {
        let result = hm.get(&vec1[i]);
        if let Some(count) = result {
            sum += vec1[i] * count;
        }
    }

    println!("DAY01 Solution 2: {}", sum);
}
