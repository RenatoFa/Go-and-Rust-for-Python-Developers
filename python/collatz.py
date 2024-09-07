def collatz_step(number: int) -> int:
    """Compute one step in Collatz conjecture"""
    if number % 2 == 0:
        return number // 2
    return number * 3 + 1


if __name__ == '__main__':
    print(collatz_step(4))
    print(collatz_step(5))
