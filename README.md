# goal

A platform to document and map informal paths created by pedestrians and animals. Users share photos and locations of these paths, building a map of how people prefer to navigate their environment.

## How to start the project

- Use brew to install postgres@16

```bash
brew install postgresql@16
```

- Start the postgres service

```bash
brew services start postgresql@16
```

- Setup the DB for local development

```bash
# From root of project type
./scripts/start_local_db.sh
```

- Make sure tailwind is installed

```bash
npm install tailwindcss@latest
```

- Run makefile

```bash
make run dev
```
