from flask import Flask, request, jsonify
from predict import predict_dropout_probability
from model_training import train_and_save_dropout_model

app = Flask(__name__)


@app.route('/train', methods=['POST'])
def train_model_route():
    try:
        result = train_and_save_dropout_model()
        return jsonify({
            "message": "Model trained successfully",
            "r2_score": round(result["r2_score"], 2),
            "rmse": round(result["rmse"], 2),
            "sample_prediction": round(result["sample_dropout_prediction"], 2)
        }), 200
    except Exception as e:
        return jsonify({"error": str(e)}), 500


@app.route('/predict', methods=['POST'])
def predict_route():
    try:
        data = request.get_json()
        required_fields = [
            "age", "gender", "experience",
            "applications_accepted", "applications_rejected", "ratings_avg"
        ]
        for field in required_fields:
            if field not in data:
                return jsonify({"error": f"Missing field: {field}"}), 400

        prediction = predict_dropout_probability(
            age=data["age"],
            gender=data["gender"],
            experience=data["experience"],
            applications_accepted=data["applications_accepted"],
            applications_rejected=data["applications_rejected"],
            ratings_avg=data["ratings_avg"]
        )
        return jsonify({"dropout_percentage": round(prediction, 2)}), 200

    except Exception as e:
        return jsonify({"error": str(e)}), 500


if __name__ == '__main__':
    app.run(debug=True)
