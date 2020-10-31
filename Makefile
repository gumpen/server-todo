# migrate
db-up:
	sql-migrate up
db-down:
	sql-migrate down -limit=0
db-reset: db-down db-up