import joblib
import pandas as pd

def predict_dropout_probability(
    age: int,
    gender: int,
    experience: int,
    applications_accepted: int,
    applications_rejected: int,
    ratings_avg: float,
    model_path="dropout_model.pkl"
) -> float:
    """
    Predict the dropout percentage for an employee using a stored model.
    
    Returns a float percentage between 0 and 100.
    """
    # Load model from disk
    try:
        model = joblib.load(model_path)
    except FileNotFoundError:
        raise FileNotFoundError(f"Model file '{model_path}' not found.")

    # Prepare input data
    input_data = pd.DataFrame([{
        "age": age,
        "gender": gender,
        "experience": experience,
        "applications_accepted": applications_accepted,
        "applications_rejected": applications_rejected,
        "ratings_avg": ratings_avg
    }])

    # Make prediction
    predicted_dropout = model.predict(input_data)[0]
    print(f"🔮 Predicted dropout percentage: {predicted_dropout:.2f}%")
    return predicted_dropout
