s = [' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ']


def print_grid():
    print('---------')
    print('|', s[0], s[1], s[2], '|')
    print('|', s[3], s[4], s[5], '|')
    print('|', s[6], s[7], s[8], '|')
    print('---------')


def result():
    x_wins = s[0] == s[1] == s[2] == 'X' or s[3] == s[4] == s[5] == 'X' or s[6] == s[7] == s[8] == 'X' or s[0] == s[3] \
             == s[6] == 'X' or s[1] == s[4] == s[7] == 'X' or s[2] == s[5] == s[8] == 'X' or s[0] == s[4] == s[8] == \
             'X' or s[2] == s[4] == s[6] == 'X'

    o_wins = s[0] == s[1] == s[2] == 'O' or s[3] == s[4] == s[5] == 'O' or s[6] == s[7] == s[8] == 'O' or s[0] == s[3] \
             == s[6] == 'O' or s[1] == s[4] == s[7] == 'O' or s[2] == s[5] == s[8] == 'O' or s[0] == s[4] == s[8]\
             == 'O' or s[2] == s[4] == s[6] == 'O'
    empty_count = 0
    for i in range(len(s)):
        if s[i] == ' ':
            empty_count += 1
    if x_wins:
        return 'X wins'
    elif o_wins:
        return 'O wins'
    elif empty_count == 0:
        return 'Draw'
    return ''


def coordinate(x, y):
    if int(x) == 1:
        return int(y) - 1
    elif int(x) == 2:
        return int(y) + 2
    return int(y) + 5


def choose_place(xo):
    x, y = input().split()
    if not x.isdigit() or not y.isdigit():
        print("You should enter numbers!")
        choose_place(xo)
    elif int(x) > 3 or int(y) > 3 or int(x) < 1 or int(y) < 1:
        print("Coordinates should be from 1 to 3!")
        choose_place(xo)
    elif s[coordinate(x, y)] != " ":
        print("This cell is occupied! Choose another one!")
        choose_place(xo)
    else:
        s[coordinate(x, y)] = xo


print_grid()
step = 0
while step <= 9:
    if step % 2 == 0:
        choose_place('X')
    else:
        choose_place('O')
    print_grid()
    result()
    if result() != '':
        print(result())
        break
    step += 1
