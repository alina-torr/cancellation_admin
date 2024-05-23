import pandas as pd
from sklearn.compose import ColumnTransformer
from category_encoders import TargetEncoder as TE
from sklearn.preprocessing import TargetEncoder
from joblib import load, dump

# pipeline = load('pipeline.joblib')

categorical_features = ["meal", "country", "market_segment", "distribution_channel", "reserved_room_type", "assigned_room_type", "customer_type", "arrival_date_month"]

def handle_categorical_feature(df, hotel_id):
    pipeline = load(f'pipeline_{hotel_id}.joblib')
    df_encoded = pipeline.transform(df)
    print(pipeline.get_feature_names_out())
    return pd.DataFrame(df_encoded, columns=pipeline.get_feature_names_out())

def preprocess(df, hotel_id):
    return handle_categorical_feature(df, hotel_id)

def create_pipeline(bookings, predicts, hotel_id):
    CategEncoderNan = TE(handle_missing='return_nan')
    CategEncoder = TargetEncoder()
    new_pipeline = ColumnTransformer([
        ('categorical_nan', CategEncoderNan, categorical_features),
        ('categorical', CategEncoder, ["agent", "company"]),
    ], remainder='passthrough')
    print(predicts)
    new_pipeline.fit(bookings, y=predicts)
    print('dsdsdds')
    dump(new_pipeline, f'pipeline_{hotel_id}.joblib')
    return new_pipeline
