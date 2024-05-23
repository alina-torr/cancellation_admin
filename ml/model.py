import grpc
import ml_pb2_grpc as pb2_grpc
import ml_pb2 as pb2
from concurrent import futures

import numpy as np
from preprocessing import create_pipeline, handle_categorical_feature, preprocess
import xgboost as xgb

from joblib import dump, load
import pandas as pd
pd.set_option('display.max_columns', None)


class MlService(pb2_grpc.MlServicer):

    def __init__(self, *args, **kwargs):
        pass

    def TrainModel(self, request, context):
        print('aaaaaaaaaaaaaaaaaaa')
        bookings = request.bookings
        df = pd.DataFrame(data=self.toModelForm(bookings))
        predicts = pd.DataFrame(data=[c for c in request.isCanceled])
        pipeline = create_pipeline(df, predicts, request.hotelId)
        print(pipeline)

        x = pipeline.transform(df)
        y = predicts
        print(x)
        train_x = xgb.DMatrix(x, label=y)

        model = xgb.train({'colsample_bytree': 0.8, 'gamma': 0.1, 'learning_rate': 0.2, 'max_depth': 9, 'min_child_weight': 1, 'reg_lambda': 1.5, 'subsample': 0.9, 'objective': 'binary:logistic'}, train_x, num_boost_round=250)
        print(model)
        dump(model, f'model_{request.hotelId}.joblib')
        return pb2.IsTrained(isTrained=True)

    def GetPredictions(self, request, context):

        bookings = request.bookings
        hotelId = request.hotelId

        df = pd.DataFrame(data=self.toModelForm(bookings))
        result_df = preprocess(df, hotelId)

        model = load(f'model_{request.hotelId}.joblib')
        
        probas = model.predict(xgb.DMatrix(result_df))
        return pb2.IsCanceledResultResponse(**{"predictions": probas})

    def toModelForm(self, bookings):
        return {
            'lead_time': [b.leadtime for b in bookings],
            'arrival_date_month': [b.arrivalDateMonth for b in bookings],
            'arrival_date_week_number': [b.arrivalDateWeekNumber for b in bookings],
            'arrival_date_day_of_month': [b.arrivalDayOfMonth for b in bookings],
            'stays_in_weekend_nights': [b.staysInWeekendNights for b in bookings],
            'stays_in_week_nights': [b.staysInWeekNights for b in bookings],
            'adults': [b.adults for b in bookings],
            'children': [b.children for b in bookings],
            'babies': [b.babies for b in bookings],
            'meal': [b.meal for b in bookings],
            'market_segment': [b.marketSegment for b in bookings],
            'distribution_channel': [b.distributionChannel for b in bookings],
            'previous_cancellations': [b.previousCancellations for b in bookings],
            'previous_bookings_not_canceled': [b.previousBookingsNotCanceled for b in bookings],
            'reserved_room_type': [b.reservedRoomType for b in bookings],
            'assigned_room_type': [b.assignedRoomType for b in bookings],
            'booking_changes': [b.bookingChanges for b in bookings],
            'agent': [b.agent for b in bookings],
            'company': [b.company for b in bookings],
            'days_in_waiting_list': [b.daysInWaitingList for b in bookings],
            'customer_type': [b.customerType for b in bookings],
            'adr': [b.adr for b in bookings],
            'required_car_parking_spaces': [b.requiredCarParkingSpaces for b in bookings],
            'total_of_special_requests': [b.totalOfSpecialRequests for b in bookings],
            'country': [b.country for b in bookings],        
        }

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_MlServicer_to_server(MlService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()