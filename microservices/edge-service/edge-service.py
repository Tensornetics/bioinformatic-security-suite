from flask import Flask, jsonify, request
from flask_cors import CORS
from blockchain import query_bioinfo, create_bioinfo, change_bioinfo_status

app = Flask(__name__)
CORS(app)


@app.route('/query', methods=['POST'])
def query():
    identifier = request.json['identifier']
    bio_info = query_bioinfo(identifier)
    return jsonify(bio_info)


@app.route('/create', methods=['POST'])
def create():
    bio_info = request.json
    create_bioinfo(bio_info)
    return jsonify({"message": "Bioinfo created"})


@app.route('/change_status', methods=['POST'])
def change_status():
    data = request.json
    identifier = data['identifier']
    status = data['status']
    change_bioinfo_status(identifier, status)
    return jsonify({"message": "Status changed"})


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
