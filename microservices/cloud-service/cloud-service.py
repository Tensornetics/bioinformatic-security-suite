from flask import Flask, request, jsonify
from flask_cors import CORS
from blockchain_sdk import Blockchain

app = Flask(__name__)
CORS(app)
blockchain = Blockchain()


@app.route('/query', methods=['GET'])
def query():
    identifier = request.args.get('identifier')
    result = blockchain.query(f'queryBioInfo({identifier})')
    return jsonify(result)


@app.route('/create', methods=['POST'])
def create():
    data = request.get_json()
    result = blockchain.invoke(
        f'createBioInfo({data["identifier"]}, {data["name"]}, {data["owner"]}, {data["hash"]}, {data["status"]})'
    )
    return jsonify(result)


@app.route('/update', methods=['POST'])
def update():
    data = request.get_json()
    result = blockchain.invoke(
        f'changeBioInfoStatus({data["identifier"]}, {data["status"]})'
    )
    return jsonify(result)


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
