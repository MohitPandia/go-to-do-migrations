# 🚀 Installation Guide

## Prerequisites

Before you start, make sure you have the following installed and configured on your system:

1. **PostgreSQL Server (ver: 14):**

   - Make sure you have PostgreSQL installed and running on your machine. You can download and install PostgreSQL from the [official website](https://www.postgresql.org/).

2. **GoLang:**

   - Install the latest version of Golang from the [official website](https://go.dev/dl/).

3. **Visual Studio Code:**

   - Install Visual Studio Code, a lightweight but powerful source code editor, from the [official website](https://code.visualstudio.com/).

4. **DBeaver:**

   - Install DBeaver, a universal database tool for developers and database administrators, from the [official website](https://dbeaver.io/).

5. **.env Setup:**
   - Run following commmand on root of repo:
   ```
   cp example.env .env
   ```

---

## 📝 Steps

- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Steps](#steps)
    - [Setting up DATABASE](#setting-up-database)
      - [Postgres Start on Local Machine](#postgres-start-on-local-machine)
      - [Postgres Status Check](#postgres-status-check)
      - [Accessing Postgres in Ubuntu](#accessing-postgres-in-ubuntu)
      - [Create User](#create-user)
      - [Create Database](#create-database)
      - [Change User Password](#change-user-password)
      - [Grant User Permissions](#grant-user-permissions)
      - [Other Useful Commands](#other-useful-commands)
    - [Installing and Connecting to Database using DBeaver](#installing-and-connecting-to-database-using-dbeaver)
      - [Steps](#steps-1)
      - [Connecting to a Database](#connecting-to-a-database)
    - [Running this repo](#running-this-repo)
  - [Verification](#verification)
  - [Additional Resources](#additional-resources)

### Setting up DATABASE

#### Connect to postgres server

```bash
# MacOS
$ psql postgres

# Ubuntu
$ sudo -u postgres psql
```

#### Create database user

```sql
CREATE USER postgres;
```

#### Create database

```sql
CREATE DATABASE rule_engine_v2;
```

#### Change password for database user

```sql
ALTER USER postgres PASSWORD 'postgres';
```

#### Grant user permissions 

```sql
-- MacOS
GRANT postgres TO zoop;

-- Ubuntu
GRANT admin TO zoop;
```

#### Other Useful Commands

To list all databases, use the following command in the Postgres command line:

```shell
\l
```

To list all users and their roles, use the following command in the Postgres command line:

```shell
\du
```

### Installing and Connecting to Database using DBeaver

This guide will walk you through the process of installing DBeaver and connecting to a database.

#### Steps

1. Go to the [DBeaver website](https://dbeaver.io/) and download the appropriate version for your operating system.
2. Run the installer and follow the on-screen instructions to complete the installation.

#### Connecting to a Database

1. Launch DBeaver.
2. Click on the "New Connection" button in the toolbar, or go to `Database` -> `New Connection`.

3. In the "New Connection" dialog, select the database type you want to connect to (e.g., PostgreSQL, MySQL, Oracle, etc.) and click "Next".

4. Fill in the connection details:

   - **Host:** The hostname or IP address of the database server.
   - **Port:** The port number on which the database server is running.
   - **Database:** The name of the database you want to connect to.
   - **Username:** Your database username.
   - **Password:** Your database password.

5. Click "Test Connection" to verify that the connection settings are correct and DBeaver can connect to the database successfully. If the test is successful, click "Finish" to save the connection.

6. The new connection will appear in the DBeaver workspace. Double-click on it to establish a connection to the database.

7. You can now explore the database structure, run queries, and perform various database operations using DBeaver.

### Running this repo

Install packages :

```shell
go mod tidy
```

```shell
go mod vendor
```

Start Migrations :

```shell
make run
```

---

## Verification

Following the steps to verify that your migrations was done properly.

- Connect to database in DBeaver
- Verify if the tables and column are created properly
- Compare tables and columns in database with code

---

## Additional Resources

- [Golang Installion](https://www.scaler.com/topics/golang/install-golang/)

- [Postgres Installion](https://www.scaler.com/topics/postgresql/install-postgresql-ubuntu-windows-mac/)
