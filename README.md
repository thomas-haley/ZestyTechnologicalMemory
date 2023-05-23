# Thomas_Challenge_2
 Thomas Haley Coding Challenge 2

## Notes
- The "input.json" file is the ingress point for this application, anything in that file will be read as json and converted.
- The "output.json" file is deleted every time the program runs, though removing it manually would lead to no issues.
- Average execution time across 50 runs: 1.06 milliseconds
    - Times in microseconds can be found in "time.txt"
- Estimated time to code: ~4 hours

## Execution
### Local
1. Clone the [repository](https://github.com/thomas-haley/ZestyTechnologicalMemory)

2. Using a CLI of your choosing, navigate to the project's folder. (In the same directory as "main.go")
3. Type the following in the command line and hit Enter.
```sh
go run .
```
4. The results of the conversion will be saved to "output.json" as well as printed to the console.

### Replit
1. Navigate to [Replit](https://replit.com) and login.
2. Once you are at your dashboard, create a new Repl at the top left of the screen.
3. Click "Import from GitHub" at the top right and paste the following url into the field labeled "GitHub URL" and click the "Import from GitHub" button at the bottom right. 
```
https://github.com/thomas-haley/ZestyTechnologicalMemory
```
4. Once imported, at the top of the screen click Run, the converted json will be saved to "output.json" as well as printed to the console.