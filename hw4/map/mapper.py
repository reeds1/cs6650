from flask import Flask, request, jsonify
import boto3
import json
from collections import Counter

app = Flask(__name__)
s3 = boto3.client('s3', region_name='us-west-2')

def parse_s3_url(url):
    parts = url.replace("s3://", "").split("/", 1)
    return parts[0], parts[1]

@app.route("/map", methods=["POST"])
def map_file():
    data = request.get_json()
    s3_url = data["s3_url"]
    bucket, key = parse_s3_url(s3_url)

    obj = s3.get_object(Bucket=bucket, Key=key)
    content = obj['Body'].read().decode()
    words = content.split()
    counts = Counter(words)

    result_key = f"{key}_count.json"
    s3.put_object(Bucket=bucket, Key=result_key, Body=json.dumps(counts))

    return jsonify(f"s3://{bucket}/{result_key}")

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080, debug=True)
