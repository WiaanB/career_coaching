
fn main() {
    constant_func(&mut[12,1,12,12]);
    linear_func(&mut[12,1,12,12]);
    square(2);
}

// O(1)
// Always executes at exactly the same speed as n does not affect the statement/function
fn constant_func(n: &mut[i32]) {
    println!("Constant:");
    let first = n[0];
    println!("{first}");
    println!("");
}

// O(n)
// O(n), is the worst case scenario, so the function has a Big O of O(n)
fn linear_func(n: &mut[i32]) {
    println!("Linear:");
    for val in n { // O(n), where n is the elements of the arr
        /*
            The following are all O(1) since they have a constant expected execution time
            These are treated as best case scenario statements/functions
        */
        let sum = &1_000 * &200_000;
        println!("{sum}");
        println!("Hello! with value {}", val);
    }
    println!("");
}

// O(n^2)
// looping over n within a loop of n, meaning as n grows, the complexity grows further
fn square(n: i32) {
    println!("Exponential:");
    for i in 0..n {
        for j in 0..n  {
            println!("Outer: {}, Inner: {}", i,j);
        }
    }
    println!("");
}