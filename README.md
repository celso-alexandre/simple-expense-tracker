# Building API+WEB Simple Template
The API uses a simple way of defining entities and methods.

Eg.:
Entity ExpensesPlan - method: create: route "ALL /expenses-plan/create"
Entity ExpensesPlan - method: list: route "ALL /expenses-plan/list"

## Choices
api - api subdirectory
language: go(lang)
database: sqlc+pgx v5

web - web subdirectory
framework: vite react
query api:
- orval (generated types based on swagger api schema)
- react-query (provides hooks to avoid useEffect + useState when querying api; caches results in client-side)
- components: antd (a very powerfull chinese library that may not be the prettiest one, but one of the most powerful and pratical) + tailwindcss
- query params: use-query-params - a very useful (despite outdated) library that leverages serializing/desserializing query params of any kind
