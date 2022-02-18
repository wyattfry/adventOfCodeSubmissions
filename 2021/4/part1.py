#!/usr/bin/env python3
import os

def get_input(input_filename):
    input_file_path = os.path.dirname(os.path.realpath(__file__)) + "/" + input_filename
    return open(input_file_path).readlines()


# board_list=[]
# i=0
# while i < len(input):
#     j=0
#     board_list.append([])
#     while j < 5:
#         board_list[int(i/5)].append(input[i+j])
#         j+=1
#     i+=5
#mic drop!
#whatcha think?
# nicee, very concise

def is_not_newline(input): # using a lambda on line 29 makes it so you don't have to have this function def here
    return input != '\n' #shortcut i didn't know was possible

def parse_boards(raw_boards):
    board = []
    board_list = []
    i = 0
    raw_boards = list(filter(is_not_newline, raw_boards)) #what would happen if you didn't use list function?
    # filter here returns a "filter object"
    # https://docs.python.org/3/library/functions.html#filter
    # "Construct an iterator from those elements of iterable for which function returns true"
    # So filter returns an "iterator", which i believe is an
    # abstract super-category that contains lists among other
    # things that can be iterated over.
    # Then the list() function takes anything that can be iterated over
    # and turns it into a list.
    while True:
        if i > 0 and i % 5 == 0:
            board_list.append(board)
            board = []
        if i >= len(raw_boards):
            break
        line = [ int(n) for n in raw_boards[i].strip().split() ]
        board.append(line)
        i += 1
    return board_list

#look man i think you're making this way harder than it is
#you should realy just filter all the "\n" lines out before you run any loops to create a list of boards
# ah i get it
# booooooom.
# still trying to figure out lamda thing.....
# lambdas took me a while too, they're used a lot
# in javascript, but coming from java i had
# never seen them before.
# They come in handy if you have a function that
# takes a function as an argument
# can you get on voice anytime soon?
# probably not but we have the weekend ;)
# ok, can you stop coding for like 10 min while i try to catch up?

# yes
# thx
# i just re-wrote it without a lambda so you can
# see the non-lambda version. it's just slightly more
# verbose.

def parse_input(string_list):
    # return
    # - numbers drawn (list of integers)
    # - list of boards (2d list of numbers, 5x5)
    numbers_drawn = [ int(n) for n in string_list[0].strip().split(',') ]
    boards = parse_boards(string_list[2:])
    return numbers_drawn, boards

############# main ##############
def main(input_filename):
    unparsed_input = get_input(input_filename)
    numbers_drawn, boards = parse_input(unparsed_input)
    # for num in numbers_drawn:
    #     print(f"number drawn: {num}")

main('test-input.txt')
# main('input.txt')
