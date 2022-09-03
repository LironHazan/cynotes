# cynotes
Save encrypted file revisions 
WIP

### Installation:
(TBD using goreleaser)
- clone the repo 
- ``go install``

### Commands: 

- New - ``Add a new secrete note``
Will generate an encrypted revision of the current file under /Users/user/.cynotes
- Edit - ``cynotes edit`` Edit a selected note.
- Read - ``cynotes read`` Choose a revision to read.
- Browse - ``cynotes browse`` Open the repo in the browser
- List - ``cynotes list``

### Development:

Add new command template: ``cobra-cli add [command]``


Took inspiration from [Eureka](https://github.com/simeg/eureka)