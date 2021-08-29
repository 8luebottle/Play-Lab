# [date_replacer.py] Replace Date(YYYY-MM-DD) as ("YYYY-MM-DD")

import logging
import os
import re

logging.basicConfig(level=logging.INFO, format='▶︎ %(message)s')
dir_path = os.path.dirname(os.path.realpath(__file__))

file_extension = '.yaml'
date_format = 'YYYY-MM-DD'
find = r'([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))'


def get_files():
    input_file_name = input("Enter the name of OpenAPI file: ")
    if not input_file_name:
        input_file_name = 'openapi.yaml'
    if not input_file_name.endswith(file_extension):
        input_file_name += file_extension

    input_file = (dir_path, input_file_name)
    output_file = (dir_path, 'openapi-result.yaml')

    logging.info("Name of the input file: {} | Name of the output file: {}".format(input_file[1], output_file[1]))
    return input_file, output_file


def replace_date(input_file, output_file):
    with open('/'.join(input_file), 'r') as original_file:
        with open('/'.join(output_file), 'a+') as new_file:
            contents = original_file.readlines()
            new_file.seek(0)
            for line in contents:
                search = line[(len(date_format) + 1) * -1:]
                found = re.search(find, search)

                if found:
                    logging.info("Found the date to replace: {}".format(search[:-1]))
                    line = re.sub(find, r'"\1"', line)

                new_file.writelines(line)

    logging.info("Date({}) has been replaced successfully".format(date_format))


def main():
    replace_date(*(get_files()))


main()
