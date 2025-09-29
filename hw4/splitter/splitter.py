from flask import Flask, request, jsonify
import boto3

app = Flask(__name__)
s3 = boto3.client('s3', region_name='us-west-2')

def parse_s3_url(url):
    parts = url.replace("s3://", "").split("/", 1)
    return parts[0], parts[1]

@app.route("/split", methods=["POST"])
def split_file():
    data = request.get_json()
    s3_url = data["s3_url"]
    bucket, key = parse_s3_url(s3_url)

    obj = s3.get_object(Bucket=bucket, Key=key)
    content = obj['Body'].read().decode()
    lines = content.splitlines()

    chunks = [lines[i::3] for i in range(3)]
    urls = []

    for idx, chunk in enumerate(chunks):
        chunk_key = f"chunk_{idx}.txt"
        s3.put_object(Bucket=bucket, Key=chunk_key, Body="\n".join(chunk))
        urls.append(f"s3://{bucket}/{chunk_key}")

    return jsonify(urls)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080, debug=True)