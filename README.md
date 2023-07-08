# spe-skill-test
add config.toml file

copy & paste config.example to config.toml

make init

make tidy

make run


APIs

1. PING

curl --location --request GET 'http://127.0.0.1:10001/ping'

3. Naricissistic Number

curl --location --request GET 'http://127.0.0.1:10001/narcissistic-number?number=111'

4. Parity Outlier

curl --location --request GET 'http://127.0.0.1:10001/parity-outlier?numbers=12&numbers=20&numbers=11&numbers=13&numbers=100&numbers=13'

5. Needle In Haystack

curl --location --request POST 'http://127.0.0.1:10001/needle-in-haystack' \
--header 'Content-Type: application/json' \
--data-raw '{
    "haystack": ["a","a","blue","a"],
    "needle": "blue"
}'

6. Blue Ocean

curl --location --request POST 'http://127.0.0.1:10001/blue-ocean' \
--header 'Content-Type: application/json' \
--data-raw '{
    "blue_ocean": [1,5,5,5,5,5,3],
    "remove": [1]
}'
