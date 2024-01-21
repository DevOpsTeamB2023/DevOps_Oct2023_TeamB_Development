# Store this code in 'app.py' file

from flask import Flask, render_template, request, redirect, url_for, session
from flask_mysqldb import MySQL
import MySQLdb.cursors
import re


app = Flask(__name__)


app.secret_key = 'your secret key'

app.config['MYSQL_HOST'] = 'localhost'
app.config['MYSQL_USER'] = 'record_system'
app.config['MYSQL_PASSWORD'] = 'dopasgpwd'
app.config['MYSQL_DB'] = 'record_db'

mysql = MySQL(app)

@app.route('/')

@app.route('/register', methods =['GET', 'POST'])
def register():
	msg = ''
	if request.method == 'POST' and 'username' in request.form and 'password' in request.form and 'email' in request.form :
		username = request.form['username']
		password = request.form['password']
		email = request.form['email']
		cursor = mysql.connection.cursor(MySQLdb.cursors.DictCursor)
		cursor.execute('SELECT * FROM Account WHERE username = % s', (username, ))
		account = cursor.fetchone()
		if account:
			msg = 'Username already exists. Try another username.'
		elif not re.match(r'[^@]+@[^@]+\.[^@]+', email):
			msg = 'Invalid email address!'
		elif not re.match(r'[A-Za-z0-9]+', username):
			msg = 'Username must contain only characters and numbers!'
		elif not username or not password or not email:
			msg = 'Please fill out the form!'
		else:
			#Account values: AccID, Username, Password, UserType, AccStatus
			cursor.execute('INSERT INTO account VALUES (NULL, % s, % s, % s, "User", "Waiting for approval")', (username, password, email))
			mysql.connection.commit()
			msg = 'Account request sent. Please wait for admin approval.'
	elif request.method == 'POST':
		msg = 'Please fill out the form!'
	return render_template('register.html', msg = msg)

if __name__ == "__main__":
    app.run(debug=True)