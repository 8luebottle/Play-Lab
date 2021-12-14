# [date_replacer.py] Replace Date(YYYY-MM-DD) as ("YYYY-MM-DD")
#                    Replace Datetime(YYYY-MM-DD hh:mm:ss) as ("YYYY-MM-DD hh:mm:ss")

import logging
import os
import re

logging.basicConfig(level=logging.INFO, format='▶︎ %(message)s')
dir_path = os.path.dirname(os.path.realpath(__file__))


class Replace:
    def __init__(self):
        self.file_extension = '.yaml'
        self.file_name = 'openapi.yaml'
        self.output_file_name = 'openapi-result.yaml'

        self.date_format = 'YYYY-MM-DD'
        self.datetime_format = 'YYYY-MM-DD hh:mm:ss'
        self.find_date = r'([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))'
        self.find_datetime = r'([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]) ([0-1][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]))'

    def get_files(self):
        input_file_name = input("Enter the name of OpenAPI file: ")
        if not input_file_name:
            input_file_name = self.file_name
        if not input_file_name.endswith(self.file_extension):
            input_file_name += self.file_extension

        input_file = (dir_path, input_file_name)
        output_file = (dir_path, self.output_file_name)

        logging.info(f"Name of the input file: {input_file[1]} | Name of the output file: {output_file[1]}")
        return input_file, output_file

    @staticmethod
    def line(line, data_format, find):
        """
        :param line: replacement line.
        :param data_format: date or date time format.
        :param find: regular expression to match a line.
        :return: replaced line
        """
        search = line[(len(data_format) + 1) * -1:]
        found = re.search(find, search)

        if found:
            logging.info(f"Found the data to replace: {search[:-1]}")
            line = re.sub(find, r'"\1"', line)

        return line

    def date(self, input_file, output_file):
        with open('/'.join(input_file), 'r') as original_file:
            with open('/'.join(output_file), 'a+') as new_file:
                contents = original_file.readlines()
                new_file.seek(0)

                for line in contents:

                    line = Replace.line(line, self.datetime_format, self.find_datetime)
                    line = Replace.line(line, self.date_format, self.find_date)

                    new_file.writelines(line)

        logging.info(f"Date({self.date_format}) and Datetime({self.datetime_format}) have been replaced successfully")


def main():
    replace = Replace()
    replace.date(*(replace.get_files()))


main()
