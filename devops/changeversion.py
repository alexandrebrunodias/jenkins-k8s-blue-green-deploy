import sys

template = ''
with open('app/template.yml', 'r+') as file:
    template = file.read()
with open('app/deployment.yml', 'w') as file:
    file.write(template.replace('NEWVERSION', sys.argv[1]))