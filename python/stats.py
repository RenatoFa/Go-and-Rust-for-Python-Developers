class StatsError(Exception):
    pass


def mean(values: list) -> int:
    if not values:
        raise StatsError('"mean" with no values')
    return sum(values) / len(values)


if __name__ == '__main__':
    values = [2, 4, 8]
    try:
        print(mean(values))
    except StatsError as error:
        print(f'error: {error}')
