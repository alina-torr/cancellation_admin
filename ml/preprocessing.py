import os
import numpy as np
import pandas as pd
from sklearn.compose import ColumnTransformer
from category_encoders import TargetEncoder as TE
from sklearn.preprocessing import TargetEncoder
from joblib import load, dump
from sklearn.impute import SimpleImputer

folder = "../files"

num_columns = ["lead_time", "stays_in_weekend_nights", "stays_in_week_nights", "adults", "children", "babies", "previous_cancellations", "previous_bookings_not_canceled", "booking_changes", "days_in_waiting_list", "adr", "required_car_parking_spaces", "total_of_special_requests"]

def preprocess(df, hotel_id):
    pipeline = load(os.path.join(folder, f'pipeline_{hotel_id}.joblib'))
    df_encoded = pipeline.transform(df)
    print(pipeline.get_feature_names_out())
    return pd.DataFrame(df_encoded, columns=pipeline.get_feature_names_out())

def create_pipeline(bookings, predicts, hotel_id):
    CategEncoder = TE(handle_missing='return_nan')
    CategEncoderNanAsCategory = TargetEncoder()
    new_pipeline = ColumnTransformer([
        ('imputer', SimpleImputer(missing_values=np.nan, strategy='mean', keep_empty_features=True), num_columns),
        ('categorical_', CategEncoder, ["market_segment", "distribution_channel", "reserved_room_type", "assigned_room_type", "customer_type", "arrival_date_month"]),
        ('categorical', CategEncoderNanAsCategory, ["agent", "company", "meal", "country"]),
    ], remainder='passthrough')
    new_pipeline.fit(bookings, y=predicts)
    dump(new_pipeline, os.path.join(folder, f'pipeline_{hotel_id}.joblib'))
    return new_pipeline
