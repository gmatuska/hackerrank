from unittest import TestCase
from training_scheduler import TrainingScheduler
import os
import random


class TestRun_training_scheduler(TestCase):
    def test_list_files_recursively(self):
        files = []
        path = "C:\\Users\\garre\\Documents\\Pluralsight\\practice"
        for d in os.walk(path):
            for file in d[2]:
                # "print(os.path.join(dir[0], file), sep='\n')"
                files.append(file)
        self.assertEqual(15, len(files))

    def test_distinct(self):
        "build distinct list of random numbers"
        scheduler = TrainingScheduler()
        path = "C:\\Users\\garre\\Documents\\Pluralsight\\practice"
        self.assertEqual(15, len(scheduler.get_files(path)))

    def test_load_from_file(self):
        with open('loaded.txt', mode='rt', encoding='utf-8') as f:
            lines = f.readlines()
        self.assertEqual(15, len(lines))

    def test_set_difference(self):
        scheduler = TrainingScheduler()
        lset = scheduler.load("loaded.txt")
        cset = scheduler.load("completed.txt")
        self.assertEqual(3, len(cset.difference(lset)))

    def test_take_random_from_list(self):
        randomSet = [1, 2, 3, 4, 5, 6, 7, 8, 9]
        otherSet = set()
        while len(otherSet) < 3:
            otherSet.add(randomSet[random.randint(0, len(randomSet)-1)])
        print(otherSet,sep='\n')
        self.assertEqual(3, len(otherSet))
