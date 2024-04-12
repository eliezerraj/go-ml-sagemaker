# go-ml-sagemaker

POC for test purposes

## Endpoints

+ POST /fraudPredict

        {
            "card_number":"111.111.000.001",
            "terminal_name": "TERM-1",
            "coord_x": 30,
            "coord_y": 30,
            "card_type":"CREDIT",
            "card_model":"VIRTUAL",
            "mcc":"COMPUTE",
            "status":"OK",
            "currency":"BRL",
            "amount": 300.55,
            "payment_at":"2024-02-14T22:59:01.859507132-03:00",
            "tx_1d": 2,
            "avg_1d": 300.57,
            "tx_7d": 3,
            "avg_7d": 300.12,
            "tx_30d": 6,
            "avg_30d": 900.82,
            "time_btw_cc_tx": 60
        }
