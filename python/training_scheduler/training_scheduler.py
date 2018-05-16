#!/bin/python3

import os
import random


class TrainingScheduler:

    @staticmethod
    def get_files(path):
        """build distinct list of random numbers"""
        files = []
        for d in os.walk(path):
            for file in d[2]:
                "print(os.path.join(dir[0], file), sep='\n')"
                files.append(os.path.join(d[0], file))
        return set(files)

    @staticmethod
    def load(filename):
        with open(filename, mode='rt', encoding='utf-8') as f:
            lines = f.readlines()
        return set(lines)

    @staticmethod
    def take(file_set, number):
        random_set = []
        file_list = list(file_set)
        while len(set(random_set)) < number:
            random_set.append(file_list[random.randint(0, len(file_set)-1)])
        return set(random_set)

    @staticmethod
    def append(file_set, file_name):
        with open(file_name, mode='at', encoding='utf-8') as f:
            f.writelines(file_set)

    @staticmethod
    def write(file_set, file_name):
        with open(file_name, mode='wt', encoding='utf-8') as f:
            f.writelines(file_set)


def run_training_scheduler():
    scheduler = TrainingScheduler()
    loaded_path = 'C:\\Users\\garre\\Documents\\Pluralsight\\practice'
    loaded_set = scheduler.get_files(loaded_path)
    completed_path = 'C:\\Users\\garre\\Documents\\Pluralsight\\completed.txt'
    if not os.path.isfile(completed_path):
        scheduler.write([], completed_path)
    completed_set = scheduler.load(completed_path)
    current_set = loaded_set.difference(completed_set)
    new_set = scheduler.take(current_set, 10)
    print(*new_set, sep='\n')
    tasks_path = 'C:\\Users\\garre\\Documents\\Pluralsight\\tasks.txt'
    scheduler.write([line + '\n' for line in list(new_set)], tasks_path)

if __name__ == '__main__':
    run_training_scheduler()
