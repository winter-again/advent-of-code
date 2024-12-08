use std::fs::File;
use std::io::{self, prelude::*, BufReader};

fn main() -> io::Result<()> {
    let file_smpl = File::open("../input_smpl.txt")?;
    let mut reader = BufReader::new(file_smpl);

    let part1_output_smpl = solve_part1(&mut reader);
    dbg!(part1_output_smpl);

    let file = File::open("../input.txt")?;
    let mut reader = BufReader::new(file);

    let part1_output = solve_part1(&mut reader);
    dbg!(part1_output);

    Ok(())
}

fn solve_part1<R: BufRead>(reader: &mut R) -> i32 {
    // let mut a: Vec<i32> = Vec::new();
    // let mut b: Vec<i32> = Vec::new();

    let mut a = vec![];
    let mut b = vec![];

    for line in reader.lines() {
        // let nums: Vec<i32> = line
        //     .unwrap()
        //     .split_whitespace()
        //     .map(|s| s.parse().expect("num parsing err"))
        //     .collect();
        // a.push(nums[0]);
        // b.push(nums[1]);

        let out = line.unwrap();
        let mut nums = out.split_whitespace();
        a.push(nums.next().unwrap().parse::<i32>().unwrap());
        b.push(nums.next().unwrap().parse::<i32>().unwrap());
    }

    a.sort();
    b.sort();

    let ans: i32 = std::iter::zip(a, b).map(|(x, y)| (x - y).abs()).sum();
    return ans;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn correct_on_smpl() {
        let file = match File::open("../input_smpl.txt") {
            Ok(file) => file,
            Err(err) => {
                println!("error reading file: {}", err);
                std::process::exit(1);
            }
        };
        let mut reader = BufReader::new(file);
        let result = solve_part1(&mut reader);
        assert_eq!(result, 11);
    }

    #[test]
    fn correct_on_actual() {
        let file = match File::open("../input.txt") {
            Ok(file) => file,
            Err(err) => {
                println!("error reading file: {}", err);
                std::process::exit(1);
            }
        };
        let mut reader = BufReader::new(file);
        let result = solve_part1(&mut reader);
        assert_eq!(result, 1970720);
    }
}
