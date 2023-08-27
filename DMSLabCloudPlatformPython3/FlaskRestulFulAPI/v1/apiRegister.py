from flask import Flask
from FlaskRestulFulAPI import systemsoftwareinformationbp as systemsoftwareinformation_bp

from FlaskRestulFulAPI import create_app

app = create_app()


app.register_blueprint(systemsoftwareinformation_bp,url_prefix="/systemsoftware")


if __name__ == '__main__':
    app.run(debug=True)
