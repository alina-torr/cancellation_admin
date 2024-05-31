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
category_columns = ["market_segment", "distribution_channel", "reserved_room_type", "assigned_room_type", "customer_type", "arrival_date_month"]

def preprocess(df, hotel_id):
    pipeline = load(os.path.join(folder, f'pipeline_{hotel_id}.joblib'))
    df_encoded = pipeline.transform(df)
    return pd.DataFrame(df_encoded, columns=pipeline.get_feature_names_out())

def create_pipeline(bookings, predicts, hotel_id):
    CategEncoder = TE(handle_missing='return_nan')
    CategEncoderNanAsCategory = TargetEncoder()
    new_pipeline = ColumnTransformer([
        ('imputer', SimpleImputer(missing_values=np.nan, strategy='median', keep_empty_features=True), num_columns),
        ('categorical_', CategEncoder, category_columns),
        ('categorical', CategEncoderNanAsCategory, ["agent", "company", "meal", "country"]),
    ], remainder='passthrough')
    new_pipeline.fit(bookings, y=predicts)
    dump(new_pipeline, os.path.join(folder, f'pipeline_{hotel_id}.joblib'))
    return new_pipeline

def remove_incorrect_data(df):
    df = df = df.drop(df[df['adr'] < 0].index)
    df = df.drop(df[(df['adults'] + df['children'] + df['babies'] == 0)].index)
    df = df.drop(df[(df['stays_in_week_nights'] + df['stays_in_weekend_nights'] == 0)].index)
    df = df.drop(df[(df['lead_time'] < 0)].index)
    df = df.drop(df[(df['adults'] < 0)].index)
    df = df.drop(df[(df['children'] < 0)].index)
    df = df.drop(df[(df['babies'] < 0)].index)
    df = df.drop(df[(df['required_car_parking_spaces'] < 0)].index)
    df = df.drop(df[(df['total_of_special_requests'] < 0)].index)
    df = df.drop(df[(df['previous_cancellations'] < 0)].index)
    df = df.drop(df[(df['previous_bookings_not_canceled'] < 0)].index)
    df = df.drop(df[(df['stays_in_weekend_nights'] < 0)].index)
    df = df.drop(df[(df['stays_in_week_nights'] < 0)].index)
    df = df.drop(df[(df['booking_changes'] < 0)].index)
    df = df.drop(df[(df['days_in_waiting_list'] < 0)].index)
    return df
