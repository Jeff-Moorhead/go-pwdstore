go-pwdstore
=========

This is a simple commandline password manager implemented in Go. It stores passwords as AES-256 encrypted
values in a JSON file either at the default location of ~/.pwdmgr/store.json (currently untested on Windows),
or at a location of the user's choosing by passing the --file option. The manager also stores the encryption
key in a file at the default location of ~/.pwdmgr/store.key, or somewhere else by specifying the --key option.

The Elephant in the Room
------------------------

This app was built as a side project to explore cryptography and how to use symmetric encryption in Go. *IT HAS
NOT BEEN TESTED AGAINST VULNERABILITIES AND SHOULD NOT BE USED TO STORE SENSITIVE INFORMATION*. One obvious
vulnerability is that it keeps the encryption key alongside the encrypted data. The manager writes data and key
files with 600 permissions so only your user can access the files, but if anybody gains access to your computer
and account information, your information will be vulnerable. Overall, there are much better commercial password
managers available. Use those instead for any real data you need to keep secret.

Usage
-----

To initialize a new password manager, run `go-pwdstore --init`. From here, you can add, remove, change, and view
passwords. To add a new password, use the `--title` and `--password` options. If `--add`, `--remove`, and `--set`
are omitted, the default behavior is to display the decrypted password for the given title.

The order of priority for options is: 

- `--init`, execution stops after a new manager is initialized. This option will overwrite any existing data unless
different data and key file locations are given via `--file` and `--key`.
- `--all`, view all titles available, no passwords are displayed. Execution stops after displaying titles.
- `--add`, requires `--title` and `--password` options.
- `--set`, requires `--title` and `--password` options.
- `--remove`, requires `--title` option.
- No flag given: displays password, requires `--title` option.

Contributing and Issues
-----------------------

If you are interested in contributing to this project, have any questions about my implementation, or would like to
report an issue, please email me at jeff.moorhead1@gmail.com.
