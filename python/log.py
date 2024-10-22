import logging
from os import remove

logging.basicConfig(
    format=(
        '%(asctime)s - %(name)s:%(levelname)s - %(filename)s:%(lineno)d - %(message)s'
    ),
    level=logging.INFO
)

logger = logging.getLogger('shutdown')


def shutdown_server(pid_file):
    try:
        with open(pid_file) as fp:
            pid = int(fp.read())
    except (OSError, ValueError) as err:
        logger.error('%r: cannot get pid - %s', pid_file, err)
        raise

    logger.info('shutting down %s', pid)

    try:
        remove(pid_file)
    except OSError as err:
        logger.warning('cannot delete %r - %s', pid_file, err)


shutdown_server('oi')
