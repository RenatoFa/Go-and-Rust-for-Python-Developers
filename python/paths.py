def split_ext(path: str) -> tuple[str, str]:
    i = path.rfind('.')
    if i == -1:
        return path, ''
    return path[:i], path[i:]


if __name__ == '__main__':
    root, extension = split_ext('app.py')
    print(f'{root=}, {extension=}')
