import subprocess
import sys

if len(sys.argv) < 2:
    print("\033[31mUsage: python done.py <commit message> \033[0m")
    sys.exit(1)

message = " ".join(sys.argv[1:])

commands = [
    ["git","pull"],
    ["git", "add", "."],
    ["git", "commit", "-m", message],
    ["git", "push"]
]

for command in commands:
    subprocess.run(command)

print("\033[32mOperation was a success.\033[0m")

# | Color   | Code |
# | ------- | ---- |
# | Red     | 31   |
# | Green   | 32   |
# | Yellow  | 33   |
# | Blue    | 34   |
# | Magenta | 35   |
# | Cyan    | 36   |