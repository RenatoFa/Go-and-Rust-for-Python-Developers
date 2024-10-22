mod mutable;

fn main() {
    let x = 14;
    println!("{}", x);

    let x: f64 = 3.14159;
    println!("{}", x);

    let x;
    x = 0;
    println!("{}", x);

    let y;
    y = "Hello World";

    println!("{}", y);
    mutable::test_mutable()
}
