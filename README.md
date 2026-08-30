# Sift

Sift is an open-source desktop database tool built with Go, Wails, React, TypeScript, and SQLite.

The project started as a small SQLite SQL CLI and is being developed into a desktop application focused on helping developers understand, inspect, and debug databases.

## Current Status

**Development stage:** early desktop prototype

### Working now

- Go project and SQLite database layer
- SQL execution against `sift.db`
- Multiline SQL input in the original CLI
- CLI commands: `.help`, `.tables`, `.schema`, `.clear`, `.cancel`, `.exit`
- SQL error handling
- Wails desktop application
- React + TypeScript frontend
- Initial Sift desktop layout
- SQL editor
- Run button
- `Ctrl + Enter` to execute a query
- Query results displayed in a table
- Query errors displayed inside the application
- Wails frontend hot reload during development

### Current checkpoint

```text
Write SQL
   ↓
Run / Ctrl + Enter
   ↓
Wails
   ↓
Go backend
   ↓
SQLite
   ↓
Return query result
   ↓
Display result in Sift
```

The current database explorer is still mostly UI scaffolding; it is not yet dynamically populated from the database.

---

# What Sift Is

Sift is not intended to be just another general-purpose SQL client.

The current direction is:

> **Sift through your database and find what matters.**

The SQL editor and database browsing features are the foundation. The main goal is to eventually help developers understand and debug their databases.

Potential areas include:

- Understanding database structure
- Inspecting tables and schemas
- Understanding how queries execute
- Finding inefficient queries
- Finding possible database problems
- Comparing database schemas

These are planned directions, not features that are already implemented.

---

# Roadmap

## 1. Database Explorer

**Next major task.**

Replace the current hardcoded database explorer with information retrieved from SQLite.

Target:

```text
DATABASE

▾ Tables
   users
   orders
   products

▾ Views
▾ Indexes
```

The explorer should eventually allow database objects to be selected and inspected.

## 2. Table Inspection

Add a proper table view.

Possible sections:

```text
Data
Structure
Indexes
```

The data view should show the actual rows in a selected table.

The structure view should show information such as:

```text
Column
Type
Nullable
Primary Key
```

## 3. Query Analysis

This is the main feature intended to differentiate Sift from a basic SQL client.

Example:

```sql
SELECT *
FROM users
WHERE email = 'bob@example.com';
```

Sift should eventually be able to show useful information about how the database executes the query.

Possible output:

```text
Query Analysis

Execution time
Rows scanned
Rows returned

Potential issue:
Full table scan

Possible improvement:
Consider an index on users(email)
```

The first implementation will focus on SQLite.

## 4. Execution Plan

Build a readable representation of a query's execution plan.

The goal is to make database execution easier to understand instead of only exposing raw `EXPLAIN` output.

## 5. Database Health

Investigate useful checks that can identify potential database problems.

Possible checks:

- Missing indexes
- Full table scans
- Duplicate data
- Suspicious relationships
- Other detectable schema or data problems

The exact checks will be decided while implementing the feature.

## 6. Database Diff

Compare two databases and show differences.

Potential output:

```text
Schema differences

users
  + phone

orders
  + created_at
```

A later version could generate SQL for applying schema changes.

## 7. Query History

Store executed queries locally.

Potential functionality:

- Search query history
- Restore previous queries
- Save queries
- Favorite queries

## 8. SQL Editor Improvements

The current SQL editor is intentionally basic.

Potential improvements:

- Syntax highlighting
- SQL autocomplete
- SQL formatting
- Multiple query tabs
- Better keyboard shortcuts
- Query error locations

## 9. Additional Database Support

SQLite is the first database supported by Sift.

After the SQLite functionality is solid, other databases can be considered.

Potential candidates:

- PostgreSQL
- MySQL / MariaDB

This is a later goal.

## 10. Distribution

The eventual goal is for Sift to be downloadable as a desktop application.

Initial target:

- Windows

Later:

- macOS
- Linux

The project should eventually have GitHub releases and downloadable builds.

---

# Development History

## Stage 1 — SQLite CLI

Sift originally started as a Go command-line application.

The CLI implemented:

- SQLite connection
- SQL execution
- Multiline SQL
- SQL error handling
- `.help`
- `.tables`
- `.schema`
- `.clear`
- `.cancel`
- `.exit`

## Stage 2 — Wails Desktop Application

The project was converted into a Wails desktop application.

Current frontend stack:

```text
React
TypeScript
Vite
```

Current backend stack:

```text
Go
SQLite
Wails
```

The frontend communicates with Go through Wails-generated bindings.

## Stage 3 — Current Desktop Prototype

The current prototype contains:

- Sift desktop window
- Database sidebar
- SQL editor
- Run button
- Query execution
- Results table
- Error display
- Status bar

The next development step is to replace the placeholder database explorer with real database metadata.

---

# Architecture

```text
┌─────────────────────────┐
│      React Frontend     │
│                         │
│  Explorer / Editor /    │
│  Results                │
└────────────┬────────────┘
             │
             │ Wails bindings
             ▼
┌─────────────────────────┐
│       Go Backend        │
│                         │
│       App               │
└────────────┬────────────┘
             │
             ▼
┌─────────────────────────┐
│     Database Layer      │
└────────────┬────────────┘
             │
             ▼
┌─────────────────────────┐
│         SQLite          │
│                         │
│        sift.db          │
└─────────────────────────┘
```

Wails provides the bridge between the Go backend and web frontend.

---

# Project Structure

```text
sift/
├── build/
├── database/
├── frontend/
├── app.go
├── main.go
├── go.mod
├── go.sum
├── sift.db
└── wails.json
```

The exact internal structure will evolve as the project grows.

---

# Development

Start Sift in development mode:

```powershell
wails dev
```

Format Go code:

```powershell
gofmt -w main.go app.go database\*.go
```

Build the application:

```powershell
wails build
```

## Design Direction

Sift should feel like a developer tool.

Current design direction:

- Dark-first
- Minimal
- Dense but readable
- Keyboard-friendly
- Low visual clutter
- SQL editor as the primary workspace
- Database explorer on the left
- Results below the editor

The UI should prioritize useful information over decorative elements.

---

# Immediate Next Steps

1. Make the database explorer dynamic.
2. Populate the table list from SQLite.
3. Allow a table to be selected.
4. Add table data inspection.
5. Add table structure inspection.
6. Start implementing SQLite query analysis.

---

# Long-Term Direction

The long-term goal is for Sift to move from:

```text
SQL editor
     +
database browser
```

toward:

```text
Database
    ↓
Understand
    ↓
Inspect
    ↓
Analyze
    ↓
Find problems
    ↓
Debug
```

The project should remain focused on those goals rather than trying to reproduce every feature of established database applications.

---

# License

License: TBD
