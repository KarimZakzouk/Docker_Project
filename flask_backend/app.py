from flask import Flask, jsonify
import mysql.connector

app = Flask(__name__)

def get_db_connection():
    connection = mysql.connector.connect(
        host='mysql',
        user='user',
        password='userpassword',
        database='mydatabase'
    )
    return connection

@app.route('/')
def get_countries():
    connection = get_db_connection()
    cursor = connection.cursor(dictionary=True)
    cursor.execute('SELECT * FROM countries;')
    countries = cursor.fetchall()
    cursor.close()
    connection.close()
    return jsonify(countries)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
