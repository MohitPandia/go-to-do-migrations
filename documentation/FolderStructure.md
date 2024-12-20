# 📂 Folder Structure

| Directory/Files                  | Description                                              |
|----------------------------------|----------------------------------------------------------|
| `.github`                        | GitHub-specific files and configurations                 |
| `.vscode`                        | VSCode-specific settings                                 |
| `app`                            | Main application code                                    |
| `commands`                       | Command script for drop, seed & migrate                  |
| `configs`                        | Configuration related code                               |
| `constants`                      | Constants used across the application                    |
| `database`                       | Database connection initalization code                   |
| `documentation`                  | Documentation for repo                                   |
| `tables`                         | Table definitions (e.g. database tables)                 |
| `utils`                          | Utility functions and helpers                            |
| `.gitignore`                     | Specifies files and directories to be ignored by Git     |
| `README.md`                      | Readme file                                              |
| `Makefile`                       | Contains script to run migrate, drop and seed            |
| `example.env`                    | Example environment variables file                       |
| `go.mod`                         | Go module file                                           |
| `go.sum`                         | Go dependencies file                                     |
| `main.go`                        | Main application entry point                             |

## Explanation

- **.github**: Contains GitHub-specific files and configurations such as actions or issue templates.
- **.vscode**: Contains VSCode-specific settings.
- **app**: The main application code.
- **commands**: The main application code.
  - **commands/drop_tables.go**: script to drop tables
  - **commands/migrate.go**: script to migrate tables and columns
  - **commands/seed.go**: script to add seeds to tables
- **configs**: Configuration files for the application.
- **constants**: Application-wide constants.
- **database**: Contains database connection and initialization code.
- **documentation**: Contains documentation for this repo.
- **tables**: Definitions of tables, reflecting the structure of the database.
- **utils**: Utility functions and helpers used across the application.
- **.gitignore**: Specifies files and directories to be ignored by Git.
- **README.md**: A markdown file containing information about the project.
- **example.env**: Example file for environment variables.
- **go.mod**: Specifies the module's dependencies.
- **go.sum**: Contains checksums for module dependencies.
- **main.go**: The main entry point of the application.
