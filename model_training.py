

import pandas as pd
from sklearn.ensemble import RandomForestRegressor
from sklearn.model_selection import train_test_split
from sklearn.metrics import mean_squared_error, r2_score
import joblib

def train_and_save_dropout_model(csv_path="employee_data.csv", model_path="dropout_model.pkl"):
    # Step 1: Load dataset
    df = pd.read_csv(csv_path)

    # Step 2: Validate columns
    required_columns = [
        "age", "gender", "experience", "applications_accepted",
        "applications_rejected", "ratings_avg", "dropout_percentage"
    ]
    for col in required_columns:
        if col not in df.columns:
            raise ValueError(f"Missing required column: {col}")

    # Step 3: Split data
    X = df[["age", "gender", "experience", "applications_accepted", "applications_rejected", "ratings_avg"]]
    y = df["dropout_percentage"]

    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

    # Step 4: Train model
    model = RandomForestRegressor(n_estimators=100, random_state=42)
    model.fit(X_train, y_train)

    # Step 5: Evaluate model
    y_pred = model.predict(X_test)
    r2 = r2_score(y_test, y_pred)
    rmse = mean_squared_error(y_test, y_pred)

    print(f"✅ Model trained successfully.")
    print(f"📈 R² Score: {r2:.2f}")
    print(f"📉 RMSE: {rmse:.2f}")

    # Step 6: Save the model
    joblib.dump(model, model_path)
    print(f"💾 Model saved to {model_path}")

    # Step 7: Sample prediction
    sample = pd.DataFrame([{
        "age": 25,
        "gender": 1,
        "experience": 3,
        "applications_accepted": 1,
        "applications_rejected": 0,
        "ratings_avg": 4.5
    }])

    predicted_dropout = model.predict(sample)[0]
    print(f"🔮 Sample dropout prediction: {predicted_dropout:.2f}%")

    return {
        "r2_score": r2,
        "rmse": rmse,
        "sample_dropout_prediction": predicted_dropout
    }

# Uncomment to run:
# train_and_save_dropout_model()
