import os

home_dir = os.path.expanduser("~")

desktop_dir = os.path.join(home_dir, "Desktop")

file_path = os.path.join(desktop_dir, 'Captura.png')


def file_head(file_name, size):
    with open(file_name, 'rb') as file:
        data = file.read(size)

    if len(data) < size:
        raise ValueError(f'{file_name!r} too small')

    return data


if __name__ == '__main__':
    breakpoint()
    data = file_head(file_path, 8)
    print(data)
