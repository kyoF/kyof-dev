def main():
    year = int(input('year:'))
    
    if year % 400 == 0:
        print('うるう年')
    elif year % 100 == 0:
        print('Not うるう年')
    elif year % 4 == 0:
        print('うるう年')
    else:
        print('Not うるう年')

if __name__ == '__main__':
    main()
